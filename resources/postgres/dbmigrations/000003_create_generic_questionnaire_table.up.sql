BEGIN;

CREATE TABLE IF NOT EXISTS generic_questionnaire (
    id SERIAL Primary Key,
    generic_questionnaire_id uuid UNIQUE NOT NULL,
    questionnaire_name VARCHAR(255) NOT NULL,
    questionnaire_description TEXT NOT NULL,
    final_remark TEXT NOT NULL
);

CREATE INDEX ON "generic_questionnaire" ("generic_questionnaire_id");

COMMIT;