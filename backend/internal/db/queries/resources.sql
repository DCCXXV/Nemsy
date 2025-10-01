-- name: GetResource :one
SELECT * FROM resources
WHERE id = $1 LIMIT 1;

-- name: ListResourcesBySubject :many
SELECT * FROM resources
WHERE subject_id = $1
ORDER BY created_at DESC;

-- name: ListResourcesByOwner :many
SELECT * FROM resources
WHERE owner_id = $1
ORDER BY created_at DESC;

-- name: CreateResource :one
INSERT INTO resources (
    owner_id, subject_id, title, description, file_url, file_size
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;
