/*
 * Test api for project
 *
 * This api is a test version to connect the backend and frontend
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"

	"github.com/DutchDestroyer/eutychia-api-gateway/models"
	"github.com/DutchDestroyer/eutychia-api-gateway/services"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

func (s *DefaultApiService) getResponseCreator() services.IResponseCreator {
	return &services.ResponseCreator{}
}

//GetAllTests
func (s *DefaultApiService) GetAllTests(ctx context.Context, accountID string) (ImplResponse, error) {

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.GetAllTests, nil,
		services.Identifiers{AccountID: accountID, ProjectID: "", TestID: ""})

	var genericTests = body.([]models.GenericTestOverview)

	var allTests []TestsForAccount

	for i := range genericTests {
		allTests = append(allTests, TestsForAccount{TestID: genericTests[i].ID, TestName: genericTests[i].Name})
	}

	return Response(httpStatusCode, allTests), err
}

//CreatesNewProject
func (s *DefaultApiService) CreatesNewProject(ctx context.Context, accountID string, createProject CreateProject) (ImplResponse, error) {

	var participants []models.Participant

	for i := range createProject.Participants {
		participants = append(participants, models.Participant{
			FirstName:    createProject.Participants[i].Firstame,
			LastName:     createProject.Participants[i].Lastname,
			EmailAddress: models.EmailAddress{createProject.Participants[i].EmailAddress},
			AccountID:    "",
		})
	}

	projectCreation := models.ProjectCreation{
		Project: models.Project{
			ID:    "",
			Title: createProject.ProjectTitle,
		},
		Participants: participants,
		Tests:        createProject.Tests,
	}

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.CreatesNewProject, projectCreation,
		services.Identifiers{AccountID: accountID, ProjectID: "", TestID: ""})

	return Response(httpStatusCode, body), err
}

// DeleteAccountByID -
func (s *DefaultApiService) DeleteAccountByID(ctx context.Context, accountID string) (ImplResponse, error) {
	// TODO - update DeleteAccountByID with the required logic for this service method.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteAccountByID method not implemented")
}

// SendEmailForSignUp -
func (s *DefaultApiService) FinalizeAccountCreation(ctx context.Context, accountID string, aCF AccountCreationFinalize) (ImplResponse, error) {

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.FinalizeAccountCreation, aCF.Password,
		services.Identifiers{AccountID: accountID, ProjectID: "", TestID: ""})

	return Response(httpStatusCode, body), err
}

// GetAccountByID -
func (s *DefaultApiService) GetAccountByID(ctx context.Context, accountID string) (ImplResponse, error) {
	return Response(http.StatusNotImplemented, nil), errors.New("GetAccountByID method not implemented")
}

// GetGenericTestOfProject -
func (s *DefaultApiService) GetGenericTestOfProject(ctx context.Context, projectID string, testID string) (ImplResponse, error) {

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.GetGenericTestOfProject, nil,
		services.Identifiers{AccountID: "", ProjectID: projectID, TestID: testID})

	test := body.(models.GenericTestData)

	var questions []GenericTestQuestions

	for i := range test.Questions {
		questions = append(questions, GenericTestQuestions{test.Questions[i].Question, test.Questions[i].QuestionType, test.Questions[i].Answers})
	}

	genericTest := GenericTest{
		Title:          test.Title,
		Description:    test.Description,
		DisplayAnswers: test.DisplayAnswers,
		FinalRemark:    test.FinalRemark,
		Questions:      questions,
	}

	return Response(httpStatusCode, genericTest), err
}

// GetProjectsOfAccount -
func (s *DefaultApiService) GetProjectsOfAccount(ctx context.Context, accountID string) (ImplResponse, error) {

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.GetProjectsOfAccount, nil,
		services.Identifiers{AccountID: accountID, ProjectID: "", TestID: ""})

	projects := body.([]models.Project)
	var projectsToReturn []Project

	for i := range projects {
		projectsToReturn = append(projectsToReturn, Project{projects[i].ID, projects[i].Title})
	}

	return Response(httpStatusCode, ProjectsAccountId{Projects: projectsToReturn}), err
}

// GetTestsToPerformByAccount -
func (s *DefaultApiService) GetTestsToPerformByAccount(ctx context.Context, projectID string, accountID string) (ImplResponse, error) {
	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.GetTestsToPerformByAccount, nil,
		services.Identifiers{AccountID: accountID, ProjectID: projectID, TestID: ""})

	tests := body.([]models.GenericTestOverview)
	var testProjects []Test

	for i := range tests {
		testProjects = append(testProjects, Test{tests[i].ID, tests[i].Name, tests[i].Type})
	}

	return Response(httpStatusCode, TestsProject{TestsToPerform: testProjects}), err
}

// LogInWithAccount -
func (s *DefaultApiService) LogInWithAccount(ctx context.Context, loginAccount LoginAccount) (ImplResponse, error) {

	account := models.Account{
		Username:     &models.EmailAddress{EmailAddress: loginAccount.EmailAddress},
		Password:     loginAccount.Password,
		AuthToken:    "",
		RefreshToken: loginAccount.RefreshToken,
		AccountID:    loginAccount.AccountID,
		AccountType:  "",
		SessionID:    loginAccount.SessionID,
		GrantType:    loginAccount.GrantType,
	}

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.LogInWithAccount, account,
		services.Identifiers{AccountID: "", ProjectID: "", TestID: ""})

	finalAccount := body.(models.Account)

	return Response(httpStatusCode, AccountDetails{finalAccount.AccountID, finalAccount.SessionID, finalAccount.AuthToken, finalAccount.RefreshToken, finalAccount.AccountType}), err
}

// LogOutWithAccount -
func (s *DefaultApiService) LogOutWithAccount(ctx context.Context, logoutAccount LogoutAccount) (ImplResponse, error) {
	account := models.Account{
		Username:     &models.EmailAddress{EmailAddress: ""},
		Password:     "",
		AuthToken:    logoutAccount.AccessToken,
		RefreshToken: "",
		AccountID:    logoutAccount.AccountID,
		AccountType:  "",
		SessionID:    logoutAccount.SessionID,
		GrantType:    "",
	}

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.LogOutWithAccount, account,
		services.Identifiers{AccountID: "", ProjectID: "", TestID: ""})

	return Response(httpStatusCode, body), err
}

// RefreshAccessToken -
func (s *DefaultApiService) RefreshAccessToken(ctx context.Context, refreshDetails RefreshDetails) (ImplResponse, error) {
	account := models.Account{
		Username:     &models.EmailAddress{EmailAddress: ""},
		Password:     "",
		AuthToken:    "",
		RefreshToken: refreshDetails.RefreshToken,
		AccountID:    refreshDetails.AccountID,
		AccountType:  "",
		SessionID:    refreshDetails.SessionID,
		GrantType:    "",
	}

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.RefreshAccessToken, account,
		services.Identifiers{AccountID: "", ProjectID: "", TestID: ""})

	return Response(httpStatusCode, body.(string)), err
}

// SubmitAnswerToTest -
func (s *DefaultApiService) SubmitAnswerToTest(ctx context.Context, projectID string, testID string, genericTestAnswers GenericTestAnswers) (ImplResponse, error) {

	var submittedAnswers []models.SubmittedAnswers

	for i := range genericTestAnswers.Answers {
		submittedAnswers = append(submittedAnswers,
			models.SubmittedAnswers{
				QuestionNumber: genericTestAnswers.Answers[i].Question,
				Answer:         genericTestAnswers.Answers[i].Answer,
				TimeToAnswer:   genericTestAnswers.Answers[i].TimeToAnswer,
			})
	}

	httpStatusCode, body, err := s.getResponseCreator().ApiCallFactory(services.SubmitAnswerToTest, submittedAnswers,
		services.Identifiers{AccountID: genericTestAnswers.AccountID, ProjectID: projectID, TestID: testID})

	return Response(httpStatusCode, body), err

}
