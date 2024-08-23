-- name: CreateUrl: one
INSERT INTO urls (original_url, short_code, click_count, is_active, user_id, created_at)
VALUES ($1, $2, $3, $4, $5, NOW()) 
RETURNING id, original_url, short_code, click_count, is_active, user_id, created_at;

-- name: GetUrl :one
SELECT * FROM urls 
WHERE id = $1 LIMIT 1;

-- name: GetUrls :many
SELECT * FROM urls ORDER BY id
LIMIT $1, OFFSET $2;

--name: UpdateUrlClickCount :one
UPDATE urls SET click_count = $1
WHERE id = $2
RETURNING id, original_url, short_code, click_count, is_active, user_id, created_at;

-- name: DeleteUrl :exec
DELETE FROM urls WHERE id = $1;