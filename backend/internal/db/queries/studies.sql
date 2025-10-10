-- name: ListStudies :many
SELECT * FROM studies
ORDER BY name;

-- name: GetStudy :one
SELECT * FROM studies
WHERE id = $1 LIMIT 1;
