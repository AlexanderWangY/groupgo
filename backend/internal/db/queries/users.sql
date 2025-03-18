-- name: GetUserByID :one
SELECT * FROM auth.users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM auth.users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM auth.users
ORDER BY created_at;

-- name: CreateUser :one
INSERT INTO auth.users (email, password_hash)
VALUES ($1, $2)
RETURNING *;
