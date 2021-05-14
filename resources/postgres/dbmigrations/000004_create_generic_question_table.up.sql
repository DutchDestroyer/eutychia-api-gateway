BEGIN;

CREATE TYPE enum_question_type AS ENUM (
    'OPEN_QUESTION',
    'MULTIPLE_CHOICE',
    'SLIDER'
);

CREATE TABLE IF NOT EXISTS generic_question (
    id SERIAL Primary Key,
    generic_question_id uuid UNIQUE NOT NULL,
    question TEXT NOT NULL,
    question_type enum_question_type NOT NULL,
    generic_questionnaire_id uuid REFERENCES generic_questionnaire (generic_questionnaire_id) NOT NULL
);

CREATE INDEX ON "generic_question" ("generic_questionnaire_id");
CREATE INDEX ON "generic_question" ("generic_question_id");

COMMIT;