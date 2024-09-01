package service

import (
	"context"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/utils"
)


type AuthService struct {
	Store db.Store
}

func (s *AuthService) Login() {}

func (s *AuthService) Register(ctx context.Context, userParams db.CreateUserParams) error {
	auth := &utils.Auth{}

	// hash the password
	hash, err := auth.HashPassword(userParams.Password)
	if err != nil {
		return err
	}

	// save hashed password
	userParams.Password = hash

	// save the record to the database
	createdUser, err := s.Store.CreateUser(ctx, userParams)

	// Generate token for user
	
	return nil
}

func (s *AuthService) ResetPassword() {}

func (s *AuthService) ForgotPassword() {}