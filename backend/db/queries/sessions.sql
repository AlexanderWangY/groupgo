-- name: CreateSession :one
INSERT INTO auth.sessions (
    user_id
) VALUES (
    $1
) RETURNING *;
