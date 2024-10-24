-- name: CreateUrl :one
INSERT INTO urls (original_url, short_code, user_id, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING id, original_url, short_code, click_count, is_active, user_id, created_at, updated_at;

-- name: GetUrl :one
SELECT * FROM urls 
WHERE id = $1 LIMIT 1;

-- name: GetUrlByCode :one
SELECT * FROM urls 
WHERE short_code = $1 LIMIT 1;

-- name: GetUrlsByUser :many
SELECT * FROM urls 
WHERE user_id = $1;

-- name: GetUrls :many
SELECT * FROM urls ORDER BY id;

-- name: UpdateUrlClickCount :one
UPDATE urls 
SET click_count = $2, is_active = $3
WHERE id = $1
RETURNING id, original_url, short_code, click_count, is_active, user_id, created_at;

-- name: UpdateUrlActive :one
UPDATE urls 
SET is_active = $2
WHERE id = $1
RETURNING id, original_url, short_code, click_count, is_active, user_id, created_at;

-- name: DeleteUrl :exec
DELETE FROM urls WHERE id = $1;