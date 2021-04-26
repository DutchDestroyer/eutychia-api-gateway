BEGIN;

CREATE TABLE IF NOT EXISTS projects_researchers_junction (
    project_id uuid REFERENCES projects NOT NULL,
    researcher_id uuid REfERENCES accounts (account_id) NOT NULL,
    UNIQUE(project_id, researcher_id)
);

CREATE INDEX ON "projects_researchers_junction" ("project_id");
CREATE INDEX ON "projects_researchers_junction" ("researcher_id");

COMMIT;