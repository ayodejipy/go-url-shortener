package service

import db "rest/api/internals/db/sqlc"


type AuthService struct {
	Store db.Store
}

func (s *AuthService) login() {}

func (s *AuthService) register() {}

func (s *AuthService) resetPassword() {}

func (s *AuthService) forgotPassword() {}