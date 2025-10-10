-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    google_sub, study_id, email, full_name, pfp, hd) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateUserStudy :one
UPDATE users
SET study_id = $2
WHERE id = $1
RETURNING *;
