package database

import "errors"

// GenericQuestionDAO model of the generic questions
type GenericQuestionDAO struct {
	QuestionID   string //UUID
	Question     string
	QuestionType string
	Answers      []string
}

var questions []GenericQuestionDAO = []GenericQuestionDAO{
	{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "Is this a question 0?", "openQuestion", []string{}},
	{"dc919c8d-93fa-4f69-8086-e2395b9d01e1", "Is this  a question 1?", "multipleChoice", []string{"yes", "no", "maybe"}},
	{"dc919c8d-93fa-4f69-8086-e2395b9d01e2", "Is this a question 2?", "slider", []string{"strongly disagree", "disagree", "neutral", "agree", "strongly agree"}},
}

// GetQuestionsPerID returns all questions with a specific ID
func GetQuestionsPerID(questionIDs []string) ([]GenericQuestionDAO, error) {
	var selectedQuestions []GenericQuestionDAO

	for i := range questions {
		for j := range questionIDs {
			if questions[i].QuestionID == questionIDs[j] {
				selectedQuestions = append(selectedQuestions, questions[i])
			}
		}
	}

	if len(selectedQuestions) != len(questionIDs) {
		return []GenericQuestionDAO{}, errors.New("Number of required questions differs from number of expected questions")
	}

	return selectedQuestions, nil
}
