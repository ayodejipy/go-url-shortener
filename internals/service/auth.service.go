package service

import (
	"context"
	"errors"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/dto"
	"rest/api/internals/logger"
	"rest/api/internals/utils"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)


type AuthService struct {
	Store db.Store
	Config *config.AppConfig
	Logger *logger.Logger
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

func (s *AuthService) Login(ctx context.Context, params dto.LoginPayload) (string, error) {
	auth := &utils.Auth{}

	// find user by email
	user, err := s.GetUserByEmail(ctx, params.Email)
	if err != nil {
		s.Logger.Error("Error occurred: %v", err)
		return "", errors.New("invalid user")
	}

	// compare whether password is correct
	err = auth.ComparePassword(user.Password, params.Password)
	if err != nil {
		s.Logger.Error("auth.ComparePassword: %v", err)
		return "", errors.New("credentials mismatch")
	}

	payload := dto.TokenPayload{
		ID: user.ID,
		Email: user.Email,
		Role: user.Role,
	}

	token, err := auth.GenerateToken(payload, s.Config.JwtSecret)
	if err != nil {
		s.Logger.Error("auth.GenerateToken: %v", err)
		return "", errors.New("token generation failed")
	}

	return token, nil
}

func (s *AuthService) Register(ctx context.Context, userParams db.CreateUserParams) (string, error) {
	auth := &utils.Auth{}

	// hash the password
	hash, err := auth.HashPassword(userParams.Password)
	if err != nil {
		s.Logger.Error("auth.HashPassword: %v", err)
		return "", errors.New("unable to hash password")
	}

	// save hashed password
	userParams.Password = hash

	// save the record to the database
	createdUser, err := s.Store.CreateUser(ctx, userParams)
	if err != nil {
		s.Logger.Error("s.Store.CreateUser: %v", err)
		return "", errors.New("failed to create new user")
	}

	// Generate token for user
	secret := s.Config.JwtSecret
	payload := dto.TokenPayload{
		ID: createdUser.ID,
		Email: createdUser.Email,
		Role: createdUser.Role,
	}

	token, err := auth.GenerateToken(payload, secret)
	if err != nil {
		s.Logger.Error("auth.GenerateToken: %v", err)
		return "", errors.New("unable to generate token")
	}
	
	return token, nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, params dto.ForgotPasswordPayload) (string, error) {
	auth := &utils.Auth{}

	// find user by email
	user, err := s.GetUserByEmail(ctx, params.Email)
	if err != nil {
		s.Logger.Error("[s.GetUserByEmail:] %v", err)
		return "", errors.New("invalid user")
	}
	// Generate token and set the expiry time
	expiresAt := time.Now().Add(1 * time.Hour).UTC()
	val, err := auth.GenerateRandomCode(32)
	if err != nil {
		s.Logger.Error("Error [auth.GenerateRandomCode]: %v", err)
		return "", errors.New("invalid credentials")
	}

	resetTokenPayload := db.CreatePasswordTokenParams{
		Token: val,
		IsActive: pgtype.Bool{Bool: true, Valid: true},
		UserID: user.ID,
		ExpiresAt: pgtype.Timestamp{Time: expiresAt, Valid: true},

	}

	record, err := s.Store.CreatePasswordToken(ctx, resetTokenPayload)
	if err != nil {
		s.Logger.Error("[s.Store.CreatePasswordToken]: %v", err)
		return "", errors.New("something went wrong")
	}
	// TODO: trigger email to user account 
	return record.Token, nil
}

func (s *AuthService) ResetPassword() {}