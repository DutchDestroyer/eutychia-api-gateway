BEGIN;

CREATE TABLE IF NOT EXISTS submitted_answer (
    submitted_answer_id uuid Primary Key UNIQUE NOT NULL,
    question_number SMALLINT NOT NULL,
    answer text NOT NULL,
    time_to_answer DOUBLE PRECISION NOT NULL,
    submitted_answers_id uuid REFERENCES submitted_answers NOT NULL
);

CREATE INDEX ON "submitted_answer" ("submitted_answers_id");

COMMIT;