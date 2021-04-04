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
		projectsToReturn = append(projectsToReturn, models.Project{ID: projects[i].ID, Name: projects[i].Name})
	}

	return projectsToReturn, nil
}

//GetProjectsAsResearcherForAccount gets all the projects of the specific accountID where this account is a researcher
func GetProjectsAsResearcherForAccount(accountID string) {

}

// GetTestsOfProject gets all the tests of a project
func GetTestsOfProject(projectID string) ([]models.GenericTestOverview, error) {
	projects, errGetProjects := database.GetProjects([]string{projectID})

	if errGetProjects != nil {
		return []models.GenericTestOverview{}, errGetProjects
	}

	var testIDs []string

	for i := range projects {
		testIDs = append(testIDs, projects[i].TestIDs...)
	}

	tests, errGetTests := database.GetTestsOfIDs(testIDs)

	if errGetTests != nil {
		return []models.GenericTestOverview{}, errGetTests
	}

	var modTest []models.GenericTestOverview

	for i := range tests {
		modTest = append(modTest, models.GenericTestOverview{ID: tests[i].ID, Name: tests[i].Name, Type: tests[i].Type})
	}

	return modTest, nil

}
