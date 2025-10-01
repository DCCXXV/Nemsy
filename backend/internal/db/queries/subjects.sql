-- name: GetSubject :one
SELECT * FROM subjects
WHERE id = $1 LIMIT 1;

-- name: ListSubjects :many
SELECT * FROM subjects
GROUP BY year
ORDER BY name;

-- name: CreateSubject :one
INSERT INTO subjects (
    name, code, year
) VALUES (
    $1, $2, $3
)
RETURNING *;
