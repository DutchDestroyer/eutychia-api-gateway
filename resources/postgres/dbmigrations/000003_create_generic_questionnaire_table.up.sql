CREATE TABLE IF NOT EXISTS generic_questionnaire (
    generic_questionnaire_id uuid  Primary Key NOT NULL,
    test_name VARCHAR(255) NOT NULL,
    test_description TEXT NOT NULL,
    display_answers BOOLEAN NOT NULL,
    final_remark TEXT NOT NULL
);