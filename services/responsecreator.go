package services

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	models "github.com/DutchDestroyer/eutychia-api-gateway/models"
	"github.com/DutchDestroyer/eutychia-api-gateway/resources/postgres"
)

type IResponseCreator interface {
	ApiCallFactory(apicall, interface{}, Identifiers) (int, interface{}, error)
}

type ResponseCreator struct {
}

type apicall int

type Identifiers struct {
	AccountID string
	ProjectID string
	TestID    string
}

const (
	GetAllTests apicall = iota
	CreatesNewProject
	FinalizeAccountCreation
	GetGenericTestOfProject
	GetProjectsOfAccount
	GetTestsToPerformByAccount
	LogInWithAccount
	LogOutWithAccount
	RefreshAccessToken
	SubmitAnswerToTest
)

func (r *ResponseCreator) ApiCallFactory(api apicall, httpBody interface{}, identifiers Identifiers) (int, interface{}, error) {

	ctx := context.Background()
	tx, err := postgres.DB.BeginTx(ctx, nil)
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	switch api {
	case GetAllTests:
		return r.getAllTests(identifiers.AccountID, tx)
	case CreatesNewProject:
		return r.createsNewProject(identifiers.AccountID, httpBody, tx)
	case FinalizeAccountCreation:
		return r.finalizeAccountCreation(identifiers.AccountID, httpBody, tx)
	case GetGenericTestOfProject:
		return r.getGenericTestOfProject(identifiers.ProjectID, identifiers.TestID, tx)
	case GetProjectsOfAccount:
		return r.getProjectsOfAccount(identifiers.AccountID, tx)
	case GetTestsToPerformByAccount:
		return r.getTestsToPerformByAccount(identifiers.ProjectID, identifiers.AccountID, tx)
	case LogInWithAccount:
		return r.logInWithAccount(httpBody, tx)
	case LogOutWithAccount:
		return r.logOutWithAccount(httpBody, tx)
	case RefreshAccessToken:
		return r.refreshAccessToken(httpBody, tx)
	case SubmitAnswerToTest:
		return r.submitAnswersToTest(identifiers.ProjectID, identifiers.TestID, identifiers.AccountID, httpBody, tx)
	default:
		return 404, nil, errors.New("unknown api call")
	}
}

func (r *ResponseCreator) getAllTests(accountID string, tx *sql.Tx) (int, interface{}, error) {

	if !IsCorrectUUID(accountID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	isResearcher, err := r.getAccountService().IsResearcherAccount(accountID, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	if !isResearcher {
		tx.Rollback()
		return getPermissionErrorResponse("account doesn't have right permissions")
	}

	genericTests, err := r.getGenTestService().GetAllGenericTests(tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return 200, genericTests, nil
}

func (r *ResponseCreator) createsNewProject(accountID string, httpBody interface{}, tx *sql.Tx) (int, interface{}, error) {
	if !IsCorrectUUID(accountID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	isResearcher, err := r.getAccountService().IsResearcherAccount(accountID, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	if !isResearcher {
		return getPermissionErrorResponse("account doesn't have right permissions")
	}

	createProject := httpBody.(models.ProjectCreation)

	for i := range createProject.Participants {
		if createProject.Participants[i].EmailAddress.IsValidEmailAddress() != nil {
			return getIncorrectDataResponse("invalid email address " + createProject.Participants[i].EmailAddress.EmailAddress)
		}
	}

	err = r.getProjectService().AddNewProject(createProject.Project.Title, createProject.Tests, accountID, createProject.Participants, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return 200, nil, nil
}

func (r *ResponseCreator) finalizeAccountCreation(accountID string, password interface{}, tx *sql.Tx) (int, interface{}, error) {
	if !IsCorrectUUID(accountID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	hasNoPassword, err := r.getAccountService().FinaleAccountCreation(accountID, password.(string), tx)

	if err != nil {
		if !hasNoPassword {
			return getPermissionErrorResponse("account already has a password")
		} else {
			tx.Rollback()
			return getDatabaseErrorResponse(err)
		}
	}

	if !hasNoPassword {
		// This should never happen!!!
		tx.Rollback()
		return getPermissionErrorResponse("account already has a password")
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return 200, nil, nil
}

func (r *ResponseCreator) getGenericTestOfProject(projectID string, testID string, tx *sql.Tx) (int, interface{}, error) {
	if !IsCorrectUUID(projectID) || !IsCorrectUUID(testID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	test, err := r.getGenTestService().GetTestData(projectID, testID, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return 200, test, nil
}

func (r *ResponseCreator) getProjectsOfAccount(accountID string, tx *sql.Tx) (int, interface{}, error) {
	if !IsCorrectUUID(accountID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	projects, err := r.getProjectService().GetProjectsAsParticipantForAccount(accountID, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return 200, projects, nil
}

func (r *ResponseCreator) getTestsToPerformByAccount(projectID string, accountID string, tx *sql.Tx) (int, interface{}, error) {
	if !IsCorrectUUID(projectID) || !IsCorrectUUID(accountID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	tests, err := r.getGenTestService().GetTestsOfProject(projectID, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return 200, tests, nil
}

func (r *ResponseCreator) logInWithAccount(account interface{}, tx *sql.Tx) (int, interface{}, error) {

	accountData := account.(models.Account)

	if accountData.Username.IsValidEmailAddress() != nil {
		return getIncorrectDataResponse("invalid email address")
	}

	// validate the account is correct
	if accountData.GrantType == "password" {
		// Validate password and obtain accountID of account
		accountDAO, err := r.getAuthService().IsValidPasswordLogin(accountData, tx)
		if err != nil {
			tx.Rollback()
			return getPermissionErrorResponse("invalid username and/or password")
		}

		// assign account ID
		accountData.AccountID = accountDAO.AccountID
		accountData.AccountType = accountDAO.AccountType

		// Create authentication for account
		err = r.getAuthService().CreateAccountAuthentication(&accountData, tx)

		if err != nil {
			tx.Rollback()
			return getDatabaseErrorResponse(err)
		}

		err = tx.Commit()
		if err != nil {
			return getDatabaseErrorResponse(err)
		}

		return 200, accountData, nil
	} else if accountData.GrantType == "refreshToken" {
		if !IsCorrectUUID(accountData.AccountID) {
			return getIncorrectDataResponse("invalid uuid")
		}

		sessionData, err := r.getAuthService().GetSessionData(accountData.AccountID, accountData.SessionID, tx)

		if err != nil {
			tx.Rollback()
			return getDatabaseErrorResponse(err)
		}

		err = r.getAuthService().IsValidTokenLogin(accountData.RefreshToken, accountData.AccountID, accountData.SessionID, accountData.GrantType, sessionData)
		if err != nil {
			return getPermissionErrorResponse("Invalid token")
		}

		// Create new authtoken for account
		newAuthToken, err := r.getAuthService().UpdateAccountAuthentication(accountData.AccountID, accountData.SessionID, tx)
		if err != nil {
			tx.Rollback()
			return getDatabaseErrorResponse(err)
		}

		accountData.AuthToken = newAuthToken

		err = tx.Commit()
		if err != nil {
			return getDatabaseErrorResponse(err)
		}

		return 200, accountData, nil
	} else {
		return getIncorrectDataResponse("Invalid grant type")
	}
}

func (r *ResponseCreator) logOutWithAccount(account interface{}, tx *sql.Tx) (int, interface{}, error) {

	logoutAccount := account.(models.Account)

	if !IsCorrectUUID(logoutAccount.AccountID) || !IsCorrectUUID(logoutAccount.SessionID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	sessionData, err := r.getAuthService().GetSessionData(logoutAccount.AccountID, logoutAccount.SessionID, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = r.getAuthService().IsValidTokenLogin(logoutAccount.AuthToken, logoutAccount.AccountID, logoutAccount.SessionID, "authToken", sessionData)

	if err != nil {
		return getPermissionErrorResponse("Invalid token")
	}

	err = r.getAuthService().LogOutWithAccount(logoutAccount.SessionID, logoutAccount.AccountID, logoutAccount.AuthToken)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return 200, nil, nil
}

func (r *ResponseCreator) refreshAccessToken(account interface{}, tx *sql.Tx) (int, interface{}, error) {

	refreshDetails := account.(models.Account)

	if !IsCorrectUUID(refreshDetails.AccountID) || !IsCorrectUUID(refreshDetails.SessionID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	sessionData, err := r.getAuthService().GetSessionData(refreshDetails.AccountID, refreshDetails.SessionID, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = r.getAuthService().IsValidTokenLogin(refreshDetails.RefreshToken, refreshDetails.AccountID, refreshDetails.SessionID, "refreshToken", sessionData)

	if err != nil {
		return getPermissionErrorResponse("invalid token")
	}

	newAuthToken, err := r.getAuthService().RefreshAccessToken(refreshDetails.AccountID, refreshDetails.SessionID, refreshDetails.RefreshToken, tx)

	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return http.StatusOK, newAuthToken, nil
}

func (r *ResponseCreator) submitAnswersToTest(projectID string, testID string, accountID string, answers interface{}, tx *sql.Tx) (int, interface{}, error) {

	genericTestAnswers := answers.([]models.SubmittedAnswers)

	if !IsCorrectUUID(projectID) || !IsCorrectUUID(testID) || !IsCorrectUUID(accountID) {
		return getIncorrectDataResponse("invalid uuid")
	}

	if len(genericTestAnswers) < 1 {
		return getIncorrectDataResponse("no answers provided")
	}

	err := r.getProjectService().StoreTestAnswers(projectID, testID, accountID, genericTestAnswers, tx)
	if err != nil {
		tx.Rollback()
		return getDatabaseErrorResponse(err)
	}

	err = tx.Commit()
	if err != nil {
		return getDatabaseErrorResponse(err)
	}

	return 200, nil, nil
}

func (r *ResponseCreator) getAccountService() IAccountService {
	return &AccountService{
		AccDBService: &database.AccountDBService{},
		AuthService:  r.getAuthService(),
	}
}

func (r *ResponseCreator) getAuthService() IAuthenticationService {
	return &AuthenticationService{
		AuthDBService:    &database.AuthenticationDBService{},
		AccountDBService: &database.AccountDBService{},
	}
}

func (r *ResponseCreator) getGenTestService() IGenTestService {
	return &GenTestService{
		GenTestDBService:     &database.GenericTestDBService{},
		ProjectDBService:     &database.ProjectDBService{},
		GenQuestionDBService: &database.GenQuestionDBService{},
	}
}

func (r *ResponseCreator) getParticipantService() IParticipantService {
	return &ParticipantService{
		AccountDBService: &database.AccountDBService{},
	}
}

func (r *ResponseCreator) getProjectService() IProjectService {
	return &ProjectService{
		ParticipantService:    r.getParticipantService(),
		AccountDBService:      &database.AccountDBService{},
		ProjectDBService:      &database.ProjectDBService{},
		StoredAnswerDBService: &database.SubmittedAnswerDBService{},
	}
}

func getIncorrectDataResponse(err string) (int, interface{}, error) {
	return 400, nil, errors.New(err)
}

func getDatabaseErrorResponse(err error) (int, interface{}, error) {
	return 500, nil, err
}

func getPermissionErrorResponse(err string) (int, interface{}, error) {
	return 401, nil, errors.New(err)
}
