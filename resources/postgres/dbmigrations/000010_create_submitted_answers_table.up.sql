BEGIN;

CREATE TABLE IF NOT EXISTS submitted_answers (
    id SERIAL Primary Key,
    submitted_answers_id uuid UNIQUE NOT NULL
);

CREATE INDEX ON "submitted_answers" ("submitted_answers_id");

COMMIT;