BEGIN;

CREATE TABLE IF NOT EXISTS projects (
    id SERIAL Primary Key,
    project_id uuid UNIQUE NOT NULL,
    project_name VARCHAR(255) NOT NULL
);

CREATE INDEX ON "projects" ("project_id");

COMMIT;
