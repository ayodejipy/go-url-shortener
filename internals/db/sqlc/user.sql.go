// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, first_name, last_name, password, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) 
RETURNING id, email, first_name, last_name, password, role, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Password  string      `json:"password"`
	Role      pgtype.Text `json:"role"`
}

type CreateUserRow struct {
	ID        pgtype.UUID      `json:"id"`
	Email     string           `json:"email"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Password  string           `json:"password"`
	Role      pgtype.Text      `json:"role"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.Password,
		arg.Role,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users 
SET is_deleted = $2, deleted_at = NOW()
WHERE id = $1
RETURNING id, email, first_name, last_name, password, role, created_at, updated_at, deleted_at
`

type DeleteUserParams struct {
	ID        pgtype.UUID `json:"id"`
	IsDeleted pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) error {
	_, err := q.db.Exec(ctx, deleteUser, arg.ID, arg.IsDeleted)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, email, first_name, last_name, password, role, created_at, updated_at, deleted_at 
FROM users 
WHERE id = $1 AND is_deleted = FALSE LIMIT 1
`

type GetUserRow struct {
	ID        pgtype.UUID      `json:"id"`
	Email     string           `json:"email"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Password  string           `json:"password"`
	Role      pgtype.Text      `json:"role"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) GetUser(ctx context.Context, id pgtype.UUID) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, first_name, last_name, password, role, created_at, updated_at, deleted_at 
FROM users 
WHERE email = $1 AND is_deleted = FALSE LIMIT 1
`

type GetUserByEmailRow struct {
	ID        pgtype.UUID      `json:"id"`
	Email     string           `json:"email"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Password  string           `json:"password"`
	Role      pgtype.Text      `json:"role"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, email, first_name, last_name, password, role, created_at, updated_at, deleted_at 
FROM users 
ORDER BY id
`

type GetUsersRow struct {
	ID        pgtype.UUID      `json:"id"`
	Email     string           `json:"email"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Password  string           `json:"password"`
	Role      pgtype.Text      `json:"role"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) GetUsers(ctx context.Context) ([]GetUsersRow, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.Password,
			&i.Role,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET email = $2, first_name = $3, last_name = $4, password = $5, role = $6, updated_at = NOW()
WHERE id = $1
RETURNING id, email, first_name, last_name, password, role, created_at, updated_at, deleted_at
`

type UpdateUserParams struct {
	ID        pgtype.UUID `json:"id"`
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Password  string      `json:"password"`
	Role      pgtype.Text `json:"role"`
}

type UpdateUserRow struct {
	ID        pgtype.UUID      `json:"id"`
	Email     string           `json:"email"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Password  string           `json:"password"`
	Role      pgtype.Text      `json:"role"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.Password,
		arg.Role,
	)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
