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
INSERT INTO auth.users (email, password_hash, first_name, last_name, is_email_verified)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUser :one
UPDATE auth.users
SET
    email = COALESCE(sqlc.narg('email'), email),
    password_hash = COALESCE(sqlc.narg('password_hash'), password_hash),
    first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name = COALESCE(sqlc.narg('last_name'), last_name),
    is_email_verified = COALESCE(sqlc.narg('is_email_verified'), is_email_verified)
WHERE id = sqlc.arg('id')
RETURNING id, email, password_hash, first_name, last_name, is_email_verified, created_at, updated_at;
