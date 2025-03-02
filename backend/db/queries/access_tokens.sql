-- name: CreateAccessToken :one
INSERT INTO auth.access_tokens (session_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAccessTokenByToken :one
SELECT * FROM auth.access_tokens
WHERE token = $1
LIMIT 1;

-- name: DeleteAccessToken :exec
DELETE FROM auth.access_tokens
WHERE id = $1;

-- name: DeleteAccessTokenBySessionID :exec
DELETE FROM auth.access_tokens
WHERE session_id = $1;
