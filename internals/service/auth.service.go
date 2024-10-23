package service

import (
	"context"
	"errors"
	"net/http"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/dto"
	"rest/api/internals/email"
	"rest/api/internals/logger"
	"rest/api/internals/utils"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService struct {
	Store  db.Store
	Config *config.AppConfig
	Logger *logger.Logger
	Utils  utils.Auth
}

// var (
//     ErrInvalidCredentials = errors.New("invalid credentials")
//     ErrUserNotFound       = errors.New("user not found")
//     ErrTokenFailed       = errors.New("token generation failed")
// )

func (s *AuthService) GetUserByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error) {
	user, err := s.Store.GetUserByEmail(ctx, email)
	if err != nil {
		return db.GetUserByEmailRow{}, err
	}

	return user, nil
}

func (s *AuthService) VerifyTokenGenerated(ctx context.Context, token string) (db.VerificationCode, error) {
	record, err := s.Store.GetVerifyCode(ctx, token)
	if err != nil {
		s.Logger.Error("[s.Store.GetPasswordToken:] %v", err)
		return db.VerificationCode{}, errors.New("invalid token")
	}

	if record.ExpiresAt.Valid {
		if time.Now().After(record.ExpiresAt.Time) {
			return db.VerificationCode{}, errors.New("token expired")
		}
	} else {
		return db.VerificationCode{}, errors.New("invalid token timestamp")
	}

	return record, nil
}

func (s *AuthService) SetAuthCookie(w http.ResponseWriter, token string) error {
	if token == "" {
		return errors.New("token is required")
	}
	// set cookie
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    token,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600 * 24 * 30,
		Secure:   false, // set to TRUE in production
		Path:     "/",
		Domain:   "",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return nil
}

func (s *AuthService) isUserVerified(ctx context.Context, id pgtype.UUID) bool {
	currentUser, err := s.Store.GetUser(ctx, id)
	if err != nil {
		s.Logger.Error("[s.Store.GetUser:] %v", err)
		return false
	}
	return currentUser.IsVerified.Bool
}

func (s *AuthService) GetVerificationCode(ctx context.Context, payload dto.RequestVerificationCodePayload) (string, error) {
	// fetch user
	user, err := s.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		s.Logger.Error("[s.GetUserByEmail]: %v", err)
		return "", errors.New("invalid email")
	}

	// check if user is already verified
	if user.IsVerified.Bool {
		return "", errors.New("user already verified")
	}

	// generate verification code
	expiresAt := time.Now().Add(30 * time.Minute).UTC()

	val, err := s.Utils.GenerateOTP(6)
	if err != nil {
		s.Logger.Error("Error [auth.GenerateOTP]: %v", err)
		return "", errors.New("invalid credentials")
	}

	verificationCodePayload := db.CreateVerifyCodeParams{
		Token:     val,
		IsActive:  pgtype.Bool{Bool: true, Valid: true},
		UserID:    user.ID,
		ExpiresAt: pgtype.Timestamp{Time: expiresAt, Valid: true},
	}
	// save the code to db for reference
	record, err := s.Store.CreateVerifyCode(ctx, verificationCodePayload)
	if err != nil {
		s.Logger.Error("[s.Store.CreateVerifyCode]: %v", err)
		return "", errors.New("something went wrong")
	}

	return record.Token, nil
}

func (s *AuthService) VerifyUser(ctx context.Context, code string) error {
	// get user from context
	// user := ctx.Value("user").(db.CreateUserRow)

	// verify the code
	otpCode, err := s.VerifyTokenGenerated(ctx, code)
	if err != nil {
		s.Logger.Error("[s.VerifyTokenGenerated:] %v", err)
		return err
	}

	user, err := s.Store.GetUser(ctx, otpCode.UserID)
	if err != nil {
		s.Logger.Error("[s.Store.GetUser:] %v", err)
		return errors.New("user not found")
	}

	// check if user is already verified
	if s.isUserVerified(ctx, otpCode.UserID) {
		return errors.New("user already verified")
	}

	// update user
	updateUserPayload := db.UpdateUserVerifiedParams{
		ID:         user.ID,
		IsVerified: pgtype.Bool{Bool: true, Valid: true},
	}

	_, err = s.Store.UpdateUserVerified(ctx, updateUserPayload)
	if err != nil {
		s.Logger.Error("[s.Store.UpdateUser:] %v", err)
		return errors.New("error verifying user")
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, params dto.LoginPayload) (string, error) {
	// find user by email
	user, err := s.GetUserByEmail(ctx, params.Email)
	if err != nil {
		s.Logger.Error("Error occurred: %v", err)
		return "", errors.New("invalid user")
	}

	// compare whether password is correct
	err = s.Utils.ComparePassword(user.Password, params.Password)
	if err != nil {
		s.Logger.Error("auth.ComparePassword: %v", err)
		return "", errors.New("credentials mismatch")
	}

	payload := dto.TokenPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	token, err := s.Utils.GenerateToken(payload, s.Config.JwtSecret)
	if err != nil {
		s.Logger.Error("auth.GenerateToken: %v", err)
		return "", errors.New("token generation failed")
	}

	return token, nil
}

func (s *AuthService) Register(ctx context.Context, userParams db.CreateUserParams) error {
	// find user by email
	// TODO: can this user check be better?
	user, _ := s.GetUserByEmail(ctx, userParams.Email)
	if user.Email == userParams.Email {
		return errors.New("user already exists")
	}

	// hash the password
	hash, err := s.Utils.HashPassword(userParams.Password)
	if err != nil {
		s.Logger.Error("auth.HashPassword: %v", err)
		return errors.New("unable to hash password")
	}

	// save hashed password
	userParams.Password = hash

	// save the record to the database
	createdUser, err := s.Store.CreateUser(ctx, userParams)
	if err != nil {
		s.Logger.Error("s.Store.CreateUser: %v", err)
		return errors.New("failed to create new user")
	}

	code, err := s.GetVerificationCode(ctx, dto.RequestVerificationCodePayload{Email: createdUser.Email})
	if err != nil {
		s.Logger.Error("s.GetVerificationCode: %v", err)
		return err
	}

	// Generate token for user
	// secret := s.Config.JwtSecret
	// payload := dto.TokenPayload{
	// 	ID:    createdUser.ID,
	// 	Email: createdUser.Email,
	// 	Role:  createdUser.Role,
	// }

	// token, err := auth.GenerateToken(payload, secret)
	// if err != nil {
	// 	s.Logger.Error("auth.GenerateToken: %v", err)
	// 	return "", errors.New("unable to generate token")
	// }

	// TODO: trigger email to user account
	emailHandler := email.NewSendEmailHandler(s.Config, s.Logger)
	emailHandler.SendOTPEmail(code, createdUser.Email)

	return nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, params dto.ForgotPasswordPayload) error {
	// find user by email
	user, err := s.GetUserByEmail(ctx, params.Email)
	if err != nil {
		s.Logger.Error("[s.GetUserByEmail:] %v", err)
		return errors.New("invalid user")
	}
	// Generate token and set the expiry time
	expiresAt := time.Now().Add(1 * time.Hour).UTC()
	val, err := s.Utils.GenerateRandomCode(32)
	if err != nil {
		s.Logger.Error("Error [auth.GenerateRandomCode]: %v", err)
		return errors.New("invalid credentials")
	}

	resetTokenPayload := db.CreateVerifyCodeParams{
		Token:     val,
		IsActive:  pgtype.Bool{Bool: true, Valid: true},
		UserID:    user.ID,
		ExpiresAt: pgtype.Timestamp{Time: expiresAt, Valid: true},
	}

	record, err := s.Store.CreateVerifyCode(ctx, resetTokenPayload)
	if err != nil {
		s.Logger.Error("[s.Store.CreateVerifyCode]: %v", err)
		return errors.New("something went wrong")
	}
	// TODO: trigger email to user account
	emailHandler := email.NewSendEmailHandler(s.Config, s.Logger)
	emailHandler.SendPasswordToken(record.Token, params.Email)

	return nil
}

func (s *AuthService) ResetPassword(ctx context.Context, params dto.ResetPasswordPayload) error {
	// find user by email
	user, err := s.GetUserByEmail(ctx, params.Email)
	if err != nil {
		s.Logger.Error("[s.GetUserByEmail:] %v", err)
		return errors.New("unknown user")
	}
	// get token and check expiry
	if _, err = s.VerifyTokenGenerated(ctx, params.Token); err != nil {
		s.Logger.Error("[s.VerifyTokenGenerated:] %v", err)
		return err
	}

	// hash new password
	newHash, err := s.Utils.HashPassword(params.Password)
	if err != nil {
		s.Logger.Error("auth.HashPassword: %v", err)
		return errors.New("error handling password")
	}

	updateUserPayload := db.UpdateUserPasswordParams{
		ID:        user.ID,
		Password:  newHash,
	}

	_, err = s.Store.UpdateUserPassword(ctx, updateUserPayload)
	if err != nil {
		s.Logger.Error("[s.Store.UpdateUser:] %v", err)
		return errors.New("error updating user password")
	}

	_, err = s.Store.UpdateVerifyCode(ctx, db.UpdateVerifyCodeParams{
		Token:     params.Token,
		IsActive:  pgtype.Bool{Bool: false, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		s.Logger.Error("[s.Store.UpdatePasswordToken:] %v", err)
		return errors.New("error invalidating token")
	}

	// TODO: trigger email to user account
	emailHandler := email.NewSendEmailHandler(s.Config, s.Logger)
	emailHandler.SendPasswordResetMail("user@example.com")

	return nil
}
