package projects

import (
	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
)

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
		projectsToReturn = append(projectsToReturn, models.Project{projects[i].ID, projects[i].Name})
	}

	return projectsToReturn, nil
}

//GetProjectsAsResearcherForAccount gets all the projects of the specific accountID where this account is a researcher
func GetProjectsAsResearcherForAccount(accountID string) {

}
