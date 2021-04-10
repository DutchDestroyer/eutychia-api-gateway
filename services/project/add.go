package projects

import (
	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
	"github.com/DutchDestroyer/eutychia-api-gateway/services/participant"
)

// AddNewProject does all logic to add a new project to be added to the db
func AddNewProject(projectName string, tests []string, researcher string, participants []models.Participant) error {

	var participantIDs []string

	for i := range participants {
		participantID, err := participant.LinkParticipantToAccount(
			participants[i].EmailAddress, participants[i].FirstName, participants[i].LastName)

		if err != nil {
			return err
		}

		participantIDs = append(participantIDs, participantID)
	}

	return database.AddNewProject(projectName, tests, []string{researcher}, participantIDs)
}
