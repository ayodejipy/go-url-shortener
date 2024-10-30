package dto

import "github.com/jackc/pgx/v5/pgtype"

type TokenPayload struct {
	ID    pgtype.UUID `json:"id"`
	Email string      `json:"email"`
	Role  pgtype.Text `json:"role"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyEmailPayload struct {
	Code string `json:"code"`
}

type RequestVerificationCodePayload struct {
	Email string `json:"email"`
}
type ForgotPasswordPayload struct {
	Email string `json:"email"`
}

type ResetPasswordPayload struct {
	Token           string `json:"token"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CreateUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password,omitempty"`
}
