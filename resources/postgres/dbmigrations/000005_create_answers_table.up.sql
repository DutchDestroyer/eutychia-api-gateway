BEGIN;

CREATE TABLE IF NOT EXISTS answers (
    answer_id uuid Primary Key UNIQUE NOT NULL,
    answer text NOT NULL,
    generic_question_id uuid REFERENCES generic_question NOT NULL
);

CREATE INDEX ON "answers" ("generic_question_id");

COMMIT;