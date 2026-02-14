CREATE TABLE resource_files (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    resource_id INT NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    s3_key TEXT NOT NULL,
    file_name TEXT NOT NULL,
    file_size BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_resource_files_resource_id ON resource_files(resource_id);

ALTER TABLE resources DROP COLUMN file_url;
ALTER TABLE resources DROP COLUMN file_size;
