package services

import (
	"database/sql"

	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
)

type IProjectService interface {
	AddNewProject(string, []string, string, []models.Participant, *sql.Tx) error
	GetProjectsAsParticipantForAccount(string, *sql.Tx) ([]models.Project, error)
	GetProjectsAsResearcherForAccount(string, *sql.Tx)
	StoreTestAnswers(string, string, string, []models.SubmittedAnswers, *sql.Tx) error
}

type ProjectService struct {
	ParticipantService    IParticipantService
	AccountDBService      database.IAccountDBService
	ProjectDBService      database.IProjectDBService
	StoredAnswerDBService database.ISubmittedAnswerDBService
}

// AddNewProject does all logic to add a new project to be added to the db
func (p *ProjectService) AddNewProject(projectName string, tests []string, researcher string, participants []models.Participant, tx *sql.Tx) error {

	var participantIDs []string

	for i := range participants {
		participantID, err := p.ParticipantService.LinkParticipantToAccount(
			participants[i].EmailAddress.EmailAddress, participants[i].FirstName, participants[i].LastName, tx)

		if err != nil {
			return err
		}

		participantIDs = append(participantIDs, participantID)
	}

	return p.ProjectDBService.AddNewProject(projectName, tests, []string{researcher}, participantIDs)
}

//GetProjectsAsParticipantForAccount gets all the projects of the specific accountID where this account is a participant
func (p *ProjectService) GetProjectsAsParticipantForAccount(accountID string, tx *sql.Tx) ([]models.Project, error) {
	projectIDs, errDbAccount := p.AccountDBService.GetProjectIDsAsParticipantForAccount(accountID, tx)

	if errDbAccount != nil {
		return []models.Project{}, errDbAccount
	}

	projects, errDBProjects := p.ProjectDBService.GetProjects(projectIDs)

	if errDBProjects != nil {
		return []models.Project{}, errDBProjects
	}

	var projectsToReturn []models.Project

	for i := range projects {
		projectsToReturn = append(projectsToReturn, models.Project{ID: projects[i].ID, Title: projects[i].Name})
	}

	return projectsToReturn, nil
}

//GetProjectsAsResearcherForAccount gets all the projects of the specific accountID where this account is a researcher
func (p *ProjectService) GetProjectsAsResearcherForAccount(accountID string, tx *sql.Tx) {

}

func (p *ProjectService) StoreTestAnswers(projectID string, testID string, accountID string, answers []models.SubmittedAnswers, tx *sql.Tx) error {

	return p.StoredAnswerDBService.StoreAnswers(projectID, testID, accountID, answers)
}
