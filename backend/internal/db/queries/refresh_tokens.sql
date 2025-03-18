-- name: CreateRefreshToken :one
INSERT INTO auth.refresh_tokens (
    user_id,
    session_id,
    token,
    expires_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetRefreshTokenByID :one
SELECT * FROM auth.refresh_tokens
WHERE id = $1 AND expires_at > NOW()
LIMIT 1;


-- name: GetRefreshTokenByToken :one
SELECT * FROM auth.refresh_tokens
WHERE token = $1 AND expires_at > NOW()
LIMIT 1;

-- name: MarkUsedRefreshTokenByID :exec
UPDATE auth.refresh_tokens
SET used = TRUE
WHERE id = $1 AND expires_at > NOW();

-- name: MarkUsedRefreshTokenByToken :exec
UPDATE auth.refresh_tokens
SET used = TRUE
WHERE token = $1 AND expires_at > NOW();
