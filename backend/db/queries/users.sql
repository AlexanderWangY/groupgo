-- name: GetUser :one
SELECT * FROM auth.users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM auth.users
ORDER BY created_at;

-- name: CreateUser :one
INSERT INTO auth.users (email, password_hash, first_name, last_name, is_email_verified)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
