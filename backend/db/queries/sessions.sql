-- name: CreateSession :one
INSERT INTO auth.sessions (
    user_id
) VALUES (
    $1
) RETURNING *;

-- name: GetSessionUserByID :one
SELECT
    u.*
FROM
    auth.sessions s
INNER JOIN auth.users u ON s.user_id = u.id
WHERE
    s.id = $1;
