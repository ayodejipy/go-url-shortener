-- name: CreateVerifyCode :one
INSERT INTO verification_codes (code, expires_at, is_active, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING id, code, expires_at, is_active, user_id created_at, updated_at;

-- name: GetVerifyCode :one
SELECT * FROM verification_codes 
WHERE code = $1 LIMIT 1;

-- name: UpdateVerifyCode :one
UPDATE verification_codes 
SET expires_at = $2, is_active = $3
WHERE code = $1
RETURNING id, code, expires_at, is_active, user_id created_at, updated_at;

-- name: DeleteVerifyCode :exec
DELETE FROM verification_codes WHERE id = $1;

-- name: DeleteExpiredCodes :exec
DELETE FROM verification_codes
WHERE expires_at < NOW();