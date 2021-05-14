BEGIN;

CREATE TABLE IF NOT EXISTS submitted_answers_generic_questionnaire_junction (
    submitted_answers_id uuid REFERENCES submitted_answers (submitted_answers_id) NOT NULL,
    generic_questionnaire_id uuid REFERENCES generic_questionnaire (generic_questionnaire_id) NOT NULL,
    UNIQUE(submitted_answers_id, generic_questionnaire_id)
);

CREATE INDEX ON "submitted_answers_generic_questionnaire_junction" ("submitted_answers_id");
CREATE INDEX ON "submitted_answers_generic_questionnaire_junction" ("generic_questionnaire_id");

COMMIT;