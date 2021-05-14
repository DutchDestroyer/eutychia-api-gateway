BEGIN;

CREATE TABLE IF NOT EXISTS submitted_answers_projects_junction (
    submitted_answers_id uuid REFERENCES submitted_answers (submitted_answers_id) NOT NULL,
    project_id uuid REFERENCES projects (project_id) NOT NULL,
    UNIQUE(submitted_answers_id, project_id)
);

CREATE INDEX ON "submitted_answers_projects_junction" ("submitted_answers_id");
CREATE INDEX ON "submitted_answers_projects_junction" ("project_id");

COMMIT;