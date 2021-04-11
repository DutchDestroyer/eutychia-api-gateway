package services

import (
	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
)

// AddNewProject does all logic to add a new project to be added to the db
func AddNewProject(projectName string, tests []string, researcher string, participants []models.Participant) error {

	var participantIDs []string

	for i := range participants {
		participantID, err := linkParticipantToAccount(
			participants[i].EmailAddress, participants[i].FirstName, participants[i].LastName)

		if err != nil {
			return err
		}

		participantIDs = append(participantIDs, participantID)
	}

	return database.AddNewProject(projectName, tests, []string{researcher}, participantIDs)
}

//GetProjectsAsParticipantForAccount gets all the projects of the specific accountID where this account is a participant
func GetProjectsAsParticipantForAccount(accountID string) ([]models.Project, error) {
	projectIDs, errDbAccount := database.GetProjectIDsAsParticipantForAccount(accountID)

	if errDbAccount != nil {
		return []models.Project{}, errDbAccount
	}

	projects, errDBProjects := database.GetProjects(projectIDs)

	if errDBProjects != nil {
		return []models.Project{}, errDBProjects
	}

	var projectsToReturn []models.Project

	for i := range projects {
		projectsToReturn = append(projectsToReturn, models.Project{ID: projects[i].ID, Name: projects[i].Name})
	}

	return projectsToReturn, nil
}

//GetProjectsAsResearcherForAccount gets all the projects of the specific accountID where this account is a researcher
func GetProjectsAsResearcherForAccount(accountID string) {

}

func StoreTestAnswers(projectID string, testID string, accountID string, answers []models.SubmittedAnswers) error {

	return database.StoreAnswers(projectID, testID, accountID, answers)
}
