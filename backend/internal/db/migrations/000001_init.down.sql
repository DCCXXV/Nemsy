DROP INDEX IF EXISTS idx_resources_subject_created;
DROP INDEX IF EXISTS idx_resources_owner_id;
DROP INDEX IF EXISTS idx_resources_subject_id;
DROP INDEX IF EXISTS idx_subjects_code;
DROP INDEX IF EXISTS idx_subjects_study_id;

DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS subjects;
DROP TABLE IF EXISTS studies;
DROP TABLE IF EXISTS users;
