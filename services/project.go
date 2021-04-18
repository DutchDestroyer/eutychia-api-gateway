package services

import (
	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
)

type IProjectService interface {
	AddNewProject(string, []string, string, []models.Participant) error
	GetProjectsAsParticipantForAccount(string) ([]models.Project, error)
	GetProjectsAsResearcherForAccount(accountID string)
	StoreTestAnswers(string, string, string, []models.SubmittedAnswers) error

	getParticipantService() IParticipantService

	getAccountDBService() database.IAccountDBService
	getProjectDBService() database.IProjectDBService
	getStoredAnswerDBService() database.ISubmittedAnswerDBService
}

type ProjectService struct{}

func (p *ProjectService) getParticipantService() IParticipantService {
	return &ParticipantService{}
}

func (p *ProjectService) getProjectDBService() database.IProjectDBService {
	return &database.ProjectDBService{}
}

func (p *ProjectService) getAccountDBService() database.IAccountDBService {
	return &database.AccountDBService{}
}

func (p *ProjectService) getStoredAnswerDBService() database.ISubmittedAnswerDBService {
	return &database.SubmittedAnswerDBService{}
}

// AddNewProject does all logic to add a new project to be added to the db
func (p *ProjectService) AddNewProject(projectName string, tests []string, researcher string, participants []models.Participant) error {

	var participantIDs []string

	for i := range participants {
		participantID, err := p.getParticipantService().LinkParticipantToAccount(
			participants[i].EmailAddress.EmailAddress, participants[i].FirstName, participants[i].LastName)

		if err != nil {
			return err
		}

		participantIDs = append(participantIDs, participantID)
	}

	return p.getProjectDBService().AddNewProject(projectName, tests, []string{researcher}, participantIDs)
}

//GetProjectsAsParticipantForAccount gets all the projects of the specific accountID where this account is a participant
func (p *ProjectService) GetProjectsAsParticipantForAccount(accountID string) ([]models.Project, error) {
	projectIDs, errDbAccount := p.getAccountDBService().GetProjectIDsAsParticipantForAccount(accountID)

	if errDbAccount != nil {
		return []models.Project{}, errDbAccount
	}

	projects, errDBProjects := p.getProjectDBService().GetProjects(projectIDs)

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
func (p *ProjectService) GetProjectsAsResearcherForAccount(accountID string) {

}

func (p *ProjectService) StoreTestAnswers(projectID string, testID string, accountID string, answers []models.SubmittedAnswers) error {

	return p.getStoredAnswerDBService().StoreAnswers(projectID, testID, accountID, answers)
}
