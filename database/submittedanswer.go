package database

import (
	"time"

	"github.com/DutchDestroyer/eutychia-api-gateway/models"
)

type SubmittedAnswersOfTestDAO struct {
	ProjectID string
	TestID    string
	AccountID string
	Answers   []SubmittedAnswerDAO
}

type SubmittedAnswerDAO struct {
	QuestionNumber int32
	Answer         string
	TimeToAnswer   float32
	TimeOfAnswer   time.Time
}

var submittedAnswersOfTest []SubmittedAnswersOfTestDAO

func StoreAnswers(projectID string, testID string, accountID string, answers []models.SubmittedAnswers) error {

	timenow := time.Now().UTC()

	var submittedAnswers []SubmittedAnswerDAO

	for i := range answers {
		submittedAnswers = append(submittedAnswers,
			SubmittedAnswerDAO{answers[i].QuestionNumber, answers[i].Answer, answers[i].TimeToAnswer, timenow})
	}

	submittedAnswersOfTest = append(submittedAnswersOfTest,
		SubmittedAnswersOfTestDAO{projectID, testID, accountID, submittedAnswers})

	return nil
}
