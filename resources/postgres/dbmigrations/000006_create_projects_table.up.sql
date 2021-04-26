CREATE TABLE IF NOT EXISTS projects (
    project_id uuid Primary Key UNIQUE NOT NULL,
    project_name VARCHAR(255) NOT NULL
);