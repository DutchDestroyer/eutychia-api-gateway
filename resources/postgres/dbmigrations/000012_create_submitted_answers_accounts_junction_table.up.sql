BEGIN;

CREATE TABLE IF NOT EXISTS submitted_answers_accounts_junction (
    submitted_answers_id uuid REFERENCES submitted_answers (submitted_answers_id) NOT NULL,
    account_id uuid REfERENCES accounts (account_id) NOT NULL,
    UNIQUE(submitted_answers_id, account_id)
);

CREATE INDEX ON "submitted_answers_accounts_junction" ("submitted_answers_id");
CREATE INDEX ON "submitted_answers_accounts_junction" ("account_id");

COMMIT;