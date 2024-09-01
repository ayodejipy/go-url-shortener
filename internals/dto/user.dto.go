package dto

import "github.com/jackc/pgx/pgtype"


type TokenPayload struct {
	ID pgtype.UUID `json:"id"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type CreateUser struct {
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password,omitempty"`
}