ALTER TABLE resources ADD COLUMN file_url TEXT NOT NULL DEFAULT '';
ALTER TABLE resources ADD COLUMN file_size BIGINT;

DROP INDEX IF EXISTS idx_resource_files_resource_id;
DROP TABLE IF EXISTS resource_files;
