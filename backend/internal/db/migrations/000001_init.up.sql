CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    google_sub TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    full_name TEXT,
    pfp TEXT,
    hd TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE studies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT UNIQUE NOT NULL,
    type TEXT,
    duration INT
);

CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    study_id INT REFERENCES studies(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    code TEXT NOT NULL,
    year TEXT
);

CREATE TABLE resources (
    id SERIAL PRIMARY KEY,
    owner_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    file_url TEXT NOT NULL,
    file_size BIGINT,
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_subjects_study_id ON subjects(study_id);
CREATE INDEX idx_subjects_code ON subjects(code);

CREATE INDEX idx_resources_subject_id ON resources(subject_id);
CREATE INDEX idx_resources_owner_id ON resources(owner_id);
CREATE INDEX idx_resources_subject_created ON resources(subject_id, created_at DESC);
