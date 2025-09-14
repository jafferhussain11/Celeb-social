-- name: CreateUser :one
INSERT INTO users (id, name, email, password_hash, bio, profile_image_url, domain)
VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE name = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET
    name = COALESCE($2, name),
    email = COALESCE($3, email),
    password_hash = COALESCE($4, password_hash),
    bio = COALESCE($5, bio),
    profile_image_url = COALESCE($6, profile_image_url),
    domain = COALESCE($7, domain),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
    LIMIT $1 OFFSET $2;