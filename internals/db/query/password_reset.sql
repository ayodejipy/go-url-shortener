-- name: CreatePasswordToken :one
INSERT INTO password_reset (token, expires_at, is_active, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING id, token, expires_at, is_active, user_id created_at, updated_at;

-- name: GetPasswordToken :one
SELECT * FROM password_reset 
WHERE token = $1 LIMIT 1;

-- name: UpdatePasswordToken :one
UPDATE password_reset 
SET expires_at = $2, is_active = $3
WHERE token = $1
RETURNING id, token, expires_at, is_active, user_id created_at, updated_at;

-- name: DeletePasswordToken :exec
DELETE FROM password_reset WHERE id = $1;

-- name: DeleteExpiredTokens :exec
DELETE FROM password_reset
WHERE expires_at < NOW();