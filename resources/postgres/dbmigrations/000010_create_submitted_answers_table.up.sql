BEGIN;

CREATE TABLE IF NOT EXISTS submitted_answers (
    id SERIAL Primary Key,
    submitted_answers_id uuid UNIQUE NOT NULL,
    project_id uuid REFERENCES projects NOT NULL,
    generic_questionnaire_id uuid REFERENCES generic_questionnaire NOT NULL,
    account_id uuid REFERENCES accounts NOT NULL
);

CREATE INDEX ON "submitted_answers" ("submitted_answers_id");
CREATE INDEX ON "submitted_answers" ("project_id");
CREATE INDEX ON "submitted_answers" ("generic_questionnaire_id");
CREATE INDEX ON "submitted_answers" ("account_id");

COMMIT;