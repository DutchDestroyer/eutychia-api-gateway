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
	accountServices "github.com/DutchDestroyer/eutychia-api-gateway/services"
	authenticationServices "github.com/DutchDestroyer/eutychia-api-gateway/services"
	gentestServices "github.com/DutchDestroyer/eutychia-api-gateway/services"
	projectServices "github.com/DutchDestroyer/eutychia-api-gateway/services"
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

//GetAllTests
func (s *DefaultApiService) GetAllTests(ctx context.Context, accountID string) (ImplResponse, error) {
	if !services.IsCorrectUUID(accountID) {
		return Response(http.StatusBadRequest, nil), errors.New("Incorrect data provided by client")
	}

	isResearcher, err1 := accountServices.IsResearcherAccount(accountID)

	if err1 != nil {
		return Response(http.StatusInternalServerError, nil), err1
	}

	if !isResearcher {
		return Response(http.StatusForbidden, nil), errors.New("account doesn't have right permissions")
	}

	genericTests, err2 := gentestServices.GetAllGenericTests()

	if err2 != nil {
		return Response(http.StatusInternalServerError, nil), err2
	}

	var allTests []TestsForAccount

	for i := range genericTests {
		allTests = append(allTests, TestsForAccount{TestID: genericTests[i].ID, TestName: genericTests[i].Name})
	}

	return Response(http.StatusOK, allTests), nil
}

//CreatesNewProject
func (s *DefaultApiService) CreatesNewProject(ctx context.Context, accountID string, createProject CreateProject) (ImplResponse, error) {
	// TODO - update CreateNewAccount with the required logic for this service method.
	if !services.IsCorrectUUID(accountID) {
		return Response(http.StatusBadRequest, nil), errors.New("Incorrect data provided by client")
	}

	isResearcher, err1 := accountServices.IsResearcherAccount(accountID)

	if err1 != nil {
		return Response(http.StatusInternalServerError, nil), err1
	}

	if !isResearcher {
		return Response(http.StatusForbidden, nil), errors.New("account doesn't have right permissions")
	}

	var participants []models.Participant

	for i := range createProject.Participants {
		p := createProject.Participants[i]
		participant, partError := services.CreateParticipant(p.Firstame, p.Lastname, p.EmailAddress)

		if partError != nil {
			return Response(http.StatusBadRequest, nil), partError
		}

		participants = append(participants, *participant)
	}

	err2 := projectServices.AddNewProject(createProject.ProjectTitle, createProject.Tests, accountID, participants)

	if err2 != nil {
		return Response(http.StatusInternalServerError, nil), err2
	}

	return Response(http.StatusOK, nil), nil
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

	isNewAccount, err1 := accountServices.FinaleAccountCreation(accountID, aCF.EmailAddress, aCF.Password, aCF.FirstName, aCF.LastName)

	if err1 != nil {
		if isNewAccount {
			return Response(http.StatusInternalServerError, nil), err1
		} else {
			return Response(http.StatusUnauthorized, nil), err1
		}
	}

	if !isNewAccount {
		// This should never happen!!!
		return Response(http.StatusUnauthorized, nil), errors.New("Invalid")
	}

	return Response(http.StatusOK, nil), nil
}

// GetAccountByID -
func (s *DefaultApiService) GetAccountByID(ctx context.Context, accountID string) (ImplResponse, error) {
	// TODO - update GetAccountByID with the required logic for this service method.

	if !services.IsCorrectUUID(accountID) {
		return Response(http.StatusBadRequest, nil), errors.New("Incorrect data provided by client")
	}

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusOK, GetAccountIdResponse{ID: "7b43fcf0-be12-4f91-8baa-fcdcac8118d5", Name: "Mark Wijnbergen", Email: "markwijnbergen@hey.com"}), nil
}

// GetGenericTestOfProject -
func (s *DefaultApiService) GetGenericTestOfProject(ctx context.Context, projectID string, testID string) (ImplResponse, error) {
	// TODO - update GetGenericTestOfProject with the required logic for this service method.

	if !services.IsCorrectUUID(projectID) || !services.IsCorrectUUID(testID) {
		return Response(http.StatusBadRequest, nil), errors.New("Incorrect data provided by client")
	}

	test, err := gentestServices.GetTestData(projectID, testID)

	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

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

	return Response(http.StatusOK, genericTest), nil
}

// GetProjectsOfAccount -
func (s *DefaultApiService) GetProjectsOfAccount(ctx context.Context, accountID string) (ImplResponse, error) {
	// TODO - update GetProjectsOfAccount with the required logic for this service method.

	if !services.IsCorrectUUID(accountID) {
		return Response(http.StatusBadRequest, nil), errors.New("Incorrect data provided by client")
	}

	projects, err := projectServices.GetProjectsAsParticipantForAccount(accountID)

	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	var projectsToReturn []Project

	for i := range projects {
		projectsToReturn = append(projectsToReturn, Project{projects[i].ID, projects[i].Name})
	}

	return Response(http.StatusOK, ProjectsAccountId{Projects: projectsToReturn}), nil
}

// GetTestsToPerformByAccount -
func (s *DefaultApiService) GetTestsToPerformByAccount(ctx context.Context, projectID string, accountID string) (ImplResponse, error) {

	if !services.IsCorrectUUID(projectID) || !services.IsCorrectUUID(accountID) {
		return Response(http.StatusBadRequest, nil), errors.New("Incorrect data provided by client")
	}

	tests, errTests := gentestServices.GetTestsOfProject(projectID)

	if errTests != nil {
		return Response(http.StatusInternalServerError, nil), errTests
	}

	var testProjects []Test

	for i := range tests {
		testProjects = append(testProjects, Test{tests[i].ID, tests[i].Name, tests[i].Type})
	}

	return Response(http.StatusOK, TestsProject{TestsToPerform: testProjects}), nil
}

// LogInWithAccount -
func (s *DefaultApiService) LogInWithAccount(ctx context.Context, loginAccount LoginAccount) (ImplResponse, error) {

	if !services.IsCorrectUUID(loginAccount.AccountID) {
		return Response(http.StatusBadRequest, nil), errors.New("Invalid account ID")
	}

	// create the account
	account, err := accountServices.GetAccount(loginAccount.EmailAddress, loginAccount.Password, loginAccount.RefreshToken, loginAccount.AccountID, loginAccount.SessionID)

	if err != nil {
		return Response(http.StatusBadRequest, nil), err
	}

	// validate the account is correct
	if loginAccount.GrantType == "password" {
		// Validate password and obtain accountID of account
		accountDAO, validationError := authenticationServices.IsValidPasswordLogin(*account)
		if validationError != nil {
			return Response(http.StatusUnauthorized, nil), validationError
		}

		// assign account ID
		account.AccountID = accountDAO.AccountID
		account.AccountType = accountDAO.AccountType

		// Create authentication for account
		authError := authenticationServices.CreateAccountAuthentication(account)

		if authError != nil {
			return Response(http.StatusInternalServerError, nil), authError
		}

		return Response(http.StatusOK,
			AccountDetails{account.AccountID, account.SessionID, account.AuthToken, account.RefreshToken, account.AccountType}), nil
	} else if loginAccount.GrantType == "refreshToken" {
		validationError := authenticationServices.IsValidTokenLogin(account.RefreshToken, account.AccountID, account.SessionID, loginAccount.GrantType)
		if validationError != nil {
			return Response(http.StatusUnauthorized, nil), validationError
		}

		// Create new authtoken for account
		newAuthToken, authError := authenticationServices.UpdateAccountAuthentication(account.AccountID, account.SessionID)
		if authError != nil {
			return Response(http.StatusInternalServerError, nil), authError
		}

		return Response(http.StatusOK,
			AccountDetails{account.AccountID, account.SessionID, newAuthToken, account.RefreshToken, account.AccountType}), nil
	} else {
		return Response(http.StatusBadRequest, nil), errors.New("grant type not recognized")
	}
}

// LogOutWithAccount -
func (s *DefaultApiService) LogOutWithAccount(ctx context.Context, logoutAccount LogoutAccount) (ImplResponse, error) {

	if services.IsCorrectUUID(logoutAccount.AccountID) || services.IsCorrectUUID(logoutAccount.SessionID) {
		return Response(http.StatusBadRequest, nil), errors.New("Invalid uuid")
	}

	err1 := authenticationServices.IsValidTokenLogin(logoutAccount.AccessToken, logoutAccount.AccountID, logoutAccount.SessionID, "authToken")

	if err1 != nil {
		return Response(http.StatusUnauthorized, nil), err1
	}

	err := authenticationServices.LogOutWithAccount(logoutAccount.SessionID, logoutAccount.AccountID, logoutAccount.AccessToken)

	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	return Response(http.StatusOK, nil), nil
}

// RefreshAccessToken -
func (s *DefaultApiService) RefreshAccessToken(ctx context.Context, refreshDetails RefreshDetails) (ImplResponse, error) {

	if !services.IsCorrectUUID(refreshDetails.AccountID) || !services.IsCorrectUUID(refreshDetails.SessionID) {
		return Response(http.StatusBadRequest, nil), errors.New("Invalid uuid")
	}

	err1 := authenticationServices.IsValidTokenLogin(refreshDetails.RefreshToken, refreshDetails.AccountID, refreshDetails.SessionID, "refreshToken")

	if err1 != nil {
		return Response(http.StatusUnauthorized, nil), err1
	}

	newAuthToken, err := authenticationServices.RefreshAccessToken(refreshDetails.AccountID, refreshDetails.SessionID, refreshDetails.RefreshToken)

	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	return Response(http.StatusOK, newAuthToken), nil
}

// SubmitAnswerToTest -
func (s *DefaultApiService) SubmitAnswerToTest(ctx context.Context, projectID string, testID string, genericTestAnswers GenericTestAnswers) (ImplResponse, error) {
	// TODO - update SubmitAnswerToTest with the required logic for this service method.

	if !services.IsCorrectUUID(projectID) || !services.IsCorrectUUID(testID) || !services.IsCorrectUUID(genericTestAnswers.AccountID) {
		return Response(http.StatusBadRequest, nil), errors.New("Incorrect data provided by client")
	}

	if len(genericTestAnswers.Answers) < 1 {
		return Response(http.StatusBadRequest, nil), errors.New("No answers provided")
	}

	var submittedAnswers []models.SubmittedAnswers

	for i := range genericTestAnswers.Answers {
		submittedAnswers = append(submittedAnswers,
			models.SubmittedAnswers{
				QuestionNumber: genericTestAnswers.Answers[i].Question,
				Answer:         genericTestAnswers.Answers[i].Answer,
				TimeToAnswer:   genericTestAnswers.Answers[i].TimeToAnswer,
			})
	}

	err := projectServices.StoreTestAnswers(projectID, testID, genericTestAnswers.AccountID, submittedAnswers)

	if err != nil {
		return Response(http.StatusInternalServerError, nil), nil
	}

	return Response(http.StatusOK, nil), nil
}
