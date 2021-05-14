BEGIN;

CREATE TABLE IF NOT EXISTS answers (
    id SERIAL Primary Key,
    answer_id uuid UNIQUE NOT NULL,
    answer text NOT NULL,
    generic_question_id uuid REFERENCES generic_question (generic_question_id) NOT NULL
);

CREATE INDEX ON "answers" ("generic_question_id");
CREATE INDEX ON "answers" ("answer_id");

COMMIT;