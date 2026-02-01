-- name: GetSubject :one
SELECT * FROM subjects
WHERE id = $1 LIMIT 1;

-- name: ListSubjects :many
SELECT * FROM subjects
ORDER BY year, name;

-- name: ListSubjectsByStudy :many
SELECT * FROM subjects
WHERE study_id = $1
ORDER BY year, name;

-- name: CreateSubject :one
INSERT INTO subjects (
    study_id, name, year
) VALUES (
    $1, $2, $3
)
RETURNING *;
