BEGIN;

CREATE TABLE IF NOT EXISTS projects_participants_junction (
    project_id uuid REFERENCES projects (project_id) NOT NULL,
    participant_id uuid REfERENCES accounts (account_id) NOT NULL,
    UNIQUE(project_id, participant_id)
);

CREATE INDEX ON "projects_participants_junction" ("project_id");
CREATE INDEX ON "projects_participants_junction" ("participant_id");

COMMIT;