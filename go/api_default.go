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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateNewAccount",
			strings.ToUpper("Post"),
			"/api/signup/create",
			c.CreateNewAccount,
		},
		{
			"DeleteAccountByID",
			strings.ToUpper("Delete"),
			"/api/accounts/{accountID}",
			c.DeleteAccountByID,
		},
		{
			"GetAccountByID",
			strings.ToUpper("Get"),
			"/api/accounts/{accountID}",
			c.GetAccountByID,
		},
		{
			"GetGenericTestOfProject",
			strings.ToUpper("Get"),
			"/api/projects/{projectID}/genericTests/{testID}",
			c.GetGenericTestOfProject,
		},
		{
			"GetProjectsOfAccount",
			strings.ToUpper("Get"),
			"/api/accounts/{accountID}/projects",
			c.GetProjectsOfAccount,
		},
		{
			"GetTestsToPerformByAccount",
			strings.ToUpper("Get"),
			"/api/projects/{projectID}/{accountID}/tests",
			c.GetTestsToPerformByAccount,
		},
		{
			"LogInWithAccount",
			strings.ToUpper("Post"),
			"/api/authentication/login",
			c.LogInWithAccount,
		},
		{
			"LogOutWithAccount",
			strings.ToUpper("Post"),
			"/api/authentication/revoke",
			c.LogOutWithAccount,
		},
		{
			"SendEmailForSignUp",
			strings.ToUpper("Post"),
			"/api/signup",
			c.SendEmailForSignUp,
		},
		{
			"SubmitAnswerToTest",
			strings.ToUpper("Post"),
			"/api/projects/{projectID}/genericTests/{testID}",
			c.SubmitAnswerToTest,
		},
	}
}

// CreateNewAccount - 
func (c *DefaultApiController) CreateNewAccount(w http.ResponseWriter, r *http.Request) { 
	accountCreation := &AccountCreation{}
	if err := json.NewDecoder(r.Body).Decode(&accountCreation); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.CreateNewAccount(r.Context(), *accountCreation)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// DeleteAccountByID - 
func (c *DefaultApiController) DeleteAccountByID(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	accountID := params["accountID"]
	result, err := c.service.DeleteAccountByID(r.Context(), accountID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetAccountByID - 
func (c *DefaultApiController) GetAccountByID(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	accountID := params["accountID"]
	result, err := c.service.GetAccountByID(r.Context(), accountID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetGenericTestOfProject - 
func (c *DefaultApiController) GetGenericTestOfProject(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	projectID := params["projectID"]
	testID := params["testID"]
	result, err := c.service.GetGenericTestOfProject(r.Context(), projectID, testID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetProjectsOfAccount - 
func (c *DefaultApiController) GetProjectsOfAccount(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	accountID := params["accountID"]
	result, err := c.service.GetProjectsOfAccount(r.Context(), accountID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetTestsToPerformByAccount - 
func (c *DefaultApiController) GetTestsToPerformByAccount(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	projectID := params["projectID"]
	accountID := params["accountID"]
	result, err := c.service.GetTestsToPerformByAccount(r.Context(), projectID, accountID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// LogInWithAccount - 
func (c *DefaultApiController) LogInWithAccount(w http.ResponseWriter, r *http.Request) { 
	loginAccount := &LoginAccount{}
	if err := json.NewDecoder(r.Body).Decode(&loginAccount); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.LogInWithAccount(r.Context(), *loginAccount)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// LogOutWithAccount - 
func (c *DefaultApiController) LogOutWithAccount(w http.ResponseWriter, r *http.Request) { 
	logoutAccount := &LogoutAccount{}
	if err := json.NewDecoder(r.Body).Decode(&logoutAccount); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.LogOutWithAccount(r.Context(), *logoutAccount)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// SendEmailForSignUp - 
func (c *DefaultApiController) SendEmailForSignUp(w http.ResponseWriter, r *http.Request) { 
	signUp := &SignUp{}
	if err := json.NewDecoder(r.Body).Decode(&signUp); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.SendEmailForSignUp(r.Context(), *signUp)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// SubmitAnswerToTest - 
func (c *DefaultApiController) SubmitAnswerToTest(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	projectID := params["projectID"]
	testID := params["testID"]
	genericTestAnswers := &GenericTestAnswers{}
	if err := json.NewDecoder(r.Body).Decode(&genericTestAnswers); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.SubmitAnswerToTest(r.Context(), projectID, testID, *genericTestAnswers)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
