package projects

import (
	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
)

func StoreTestAnswers(projectID string, testID string, accountID string, answers []models.SubmittedAnswers) error {

	return database.StoreAnswers(projectID, testID, accountID, answers)
}
