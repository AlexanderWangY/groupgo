-- name: CreateRefreshToken :one
INSERT INTO auth.refresh_tokens (
    session_id,
    token,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetRefreshTokenByToken :one
SELECT * FROM auth.refresh_tokens
WHERE token = $1
LIMIT 1;

-- name: RevokeRefreshToken :one
UPDATE auth.refresh_tokens
SET is_revoked = true
WHERE id = $1
RETURNING *;

-- name: RevokeAllRefreshTokensForSession :exec
UPDATE auth.refresh_tokens
SET is_revoked = true
WHERE session_id = $1;

-- name: DeleteRefreshTokenByID :exec
DELETE FROM auth.refresh_tokens
WHERE id = $1;

-- name: DeleteRefreshTokenByToken :exec
DELETE FROM auth.refresh_tokens
WHERE token = $1;
