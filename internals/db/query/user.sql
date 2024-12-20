-- name: CreateUser :one
INSERT INTO users (email, first_name, last_name, password, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) 
RETURNING id, email, first_name, last_name, password, role, created_at, updated_at, deleted_at;

-- name: GetUser :one
SELECT id, email, first_name, last_name, password, role, is_verified, created_at, updated_at, deleted_at 
FROM users 
WHERE id = $1 AND is_deleted = FALSE LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, email, first_name, last_name, password, role, is_verified, created_at, updated_at, deleted_at 
FROM users 
WHERE email = $1 AND is_deleted = FALSE LIMIT 1;

-- name: GetUsers :many
SELECT id, email, first_name, last_name, password, role, created_at, updated_at, deleted_at 
FROM users 
ORDER BY id;

-- name: UpdateUserPassword :one
UPDATE users 
SET password = $2, updated_at = NOW()
WHERE id = $1
RETURNING id, email, first_name, last_name, role, is_verified, created_at, updated_at, deleted_at;

-- name: UpdateUserVerified :one
UPDATE users 
SET is_verified = $2, updated_at = NOW()
WHERE id = $1
RETURNING id, email, first_name, last_name, role, is_verified, created_at, updated_at, deleted_at;

-- name: UpdateUser :one
UPDATE users 
SET email = $2, first_name = $3, last_name = $4, role = $5, is_verified = $6, updated_at = NOW()
WHERE id = $1
RETURNING id, email, first_name, last_name, role, is_verified, created_at, updated_at, deleted_at;

-- name: DeleteUser :exec
UPDATE users 
SET is_deleted = $2, deleted_at = NOW()
WHERE id = $1
RETURNING id, email, first_name, last_name, password, role, is_verified, created_at, updated_at, deleted_at;