-- name: CreateSession :one
INSERT INTO auth.sessions (
    user_id,
    token,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetSessionByToken :one
SELECT * FROM auth.sessions
WHERE token = $1 AND expires_at > NOW()
LIMIT 1;

-- name: DeleteSessionByID :exec
DELETE FROM auth.sessions
WHERE id = $1;

-- name: DeleteAllUserSessionsById :exec
DELETE FROM auth.sessions
WHERE user_id = $1;

-- name: GetAllUserSessionToken :many
SELECT token FROM auth.sessions
WHERE user_id = $1 AND expires_at > NOW();

-- name: GetSessionWithUserInformation :one
SELECT
    auth.sessions.token,
    auth.sessions.id,
    auth.sessions.expires_at,
    auth.users.id AS user_id,
    auth.users.payment_plan
FROM auth.sessions
JOIN auth.users ON auth.sessions.user_id = auth.users.id
WHERE auth.sessions.id = $1;
