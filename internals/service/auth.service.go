package service

import (
	"context"
	"errors"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/dto"
	"rest/api/internals/utils"
)


type AuthService struct {
	Store db.Store
	Config *config.AppConfig
}

func (s *AuthService) Login() {}

func (s *AuthService) Register(ctx context.Context, userParams db.CreateUserParams) (string, error) {
	auth := &utils.Auth{}

	// hash the password
	hash, err := auth.HashPassword(userParams.Password)
	if err != nil {
		return "", errors.New("unable to hash password")
	}

	// save hashed password
	userParams.Password = hash

	// save the record to the database
	createdUser, err := s.Store.CreateUser(ctx, userParams)
	if err != nil {
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
		return "", errors.New("unable to generate token")
	}
	
	return token, nil
}

func (s *AuthService) ResetPassword() {}

func (s *AuthService) ForgotPassword() {}