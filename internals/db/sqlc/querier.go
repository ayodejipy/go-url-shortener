// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateUrl(ctx context.Context, arg CreateUrlParams) (CreateUrlRow, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	CreateVerifyCode(ctx context.Context, arg CreateVerifyCodeParams) (CreateVerifyCodeRow, error)
	DeleteExpiredCodes(ctx context.Context) error
	DeleteUrl(ctx context.Context, id pgtype.UUID) error
	DeleteUser(ctx context.Context, arg DeleteUserParams) error
	DeleteVerifyCode(ctx context.Context, id pgtype.UUID) error
	GetUrl(ctx context.Context, id pgtype.UUID) (Url, error)
	GetUrlByCode(ctx context.Context, shortCode string) (Url, error)
	GetUrls(ctx context.Context) ([]Url, error)
	GetUrlsByUser(ctx context.Context, userID pgtype.UUID) ([]Url, error)
	GetUser(ctx context.Context, id pgtype.UUID) (GetUserRow, error)
	GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error)
	GetUsers(ctx context.Context) ([]GetUsersRow, error)
	GetVerifyCode(ctx context.Context, token string) (VerificationCode, error)
	UpdateUrlActive(ctx context.Context, arg UpdateUrlActiveParams) (UpdateUrlActiveRow, error)
	UpdateUrlClickCount(ctx context.Context, arg UpdateUrlClickCountParams) (UpdateUrlClickCountRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (UpdateUserPasswordRow, error)
	UpdateUserVerified(ctx context.Context, arg UpdateUserVerifiedParams) (UpdateUserVerifiedRow, error)
	UpdateVerifyCode(ctx context.Context, arg UpdateVerifyCodeParams) (UpdateVerifyCodeRow, error)
}

var _ Querier = (*Queries)(nil)
