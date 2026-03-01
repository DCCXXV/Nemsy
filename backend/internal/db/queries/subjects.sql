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

-- name: ListSubjectsByStudyWithPinned :many
SELECT s.*,
       EXISTS(
           SELECT 1 FROM pinned_subjects ps
           WHERE ps.user_id = $2 AND ps.subject_id = s.id
       ) AS pinned
FROM subjects s
WHERE s.study_id = $1
ORDER BY s.year, s.name;

-- name: PinSubject :exec
INSERT INTO pinned_subjects (user_id, subject_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: UnpinSubject :exec
DELETE FROM pinned_subjects WHERE user_id = $1 AND subject_id = $2;
