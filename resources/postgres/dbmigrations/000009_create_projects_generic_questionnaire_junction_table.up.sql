BEGIN;

CREATE TABLE IF NOT EXISTS projects_generic_questionnaire_junction (
    project_id uuid REFERENCES projects (project_id) NOT NULL,
    generic_questionnaire_id uuid REFERENCES generic_questionnaire (generic_questionnaire_id) NOT NULL,
    UNIQUE(project_id, generic_questionnaire_id)
);

CREATE INDEX ON "projects_generic_questionnaire_junction" ("project_id");
CREATE INDEX ON "projects_generic_questionnaire_junction" ("generic_questionnaire_id");

COMMIT;
