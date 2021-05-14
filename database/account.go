package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

// AccountDAO data stored in the db
type AccountDAO struct {
	AccountID             string
	FirstName             string
	NonceFirstName        string
	LastName              string
	NonceLastName         string
	EmailAddress          string
	NonceEmailAddress     string
	Password              string
	AccountType           string
	ProjectsAsResearcher  []string // Only the IDs are stored here
	ProjectsAsParticipant []string // Only the IDs are stored here
}

var accountDatabase []AccountDAO = []AccountDAO{
	/*{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72", "Mark1", "Wijnbergen1", "mark.wijnbergen@hey.com", "$2a$04$cR6VcDXU4cSk0gSd1Hmr4euIfZxcYWunEUs1iMZu29JXeWOUN5O1m", "researcher",
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808", "497aeeaf-0d41-46c4-a5a1-8a88c7b61809"},
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808", "497aeeaf-0d41-46c4-a5a1-8a88c7b61809"}},
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b73", "Mark2", "Wijnbergen2", "wijnbergenmark@gmail.com", "$2a$04$cR6VcDXU4cSk0gSd1Hmr4euIfZxcYWunEUs1iMZu29JXeWOUN5O1m", "participant",
		[]string{},
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808"},
	},*/
}

type IAccountDBService interface {
	CreateInitialParticipantAccount(string, string, string, string, string, string, *sql.Tx) (string, error)
	FinalizeAccountCreation(string, []byte, *sql.Tx) error
	GetDatabaseEntry(string, *sql.Tx) (AccountDAO, error)
	GetDatabaseEntryBasedOnMail(string, *sql.Tx) (AccountDAO, error)
	GetProjectIDsAsParticipantForAccount(string, *sql.Tx) ([]string, error)
	GetProjectIDsAsResearcherForAccount(string, *sql.Tx) ([]string, error)
}

type AccountDBService struct{}

// CreateInitialParticipantAccount creates the participant when the participant is for the first time added
// by a researcher for a project. This participant hasn't confirmed it's account yet, so has no password
func (a *AccountDBService) CreateInitialParticipantAccount(
	firstName string,
	nonceFirstName string,
	lastName string,
	nonceLastName string,
	emailAddress string,
	nonceEmailAddress string,
	tx *sql.Tx) (string, error) {

	accountID := uuid.New().String()

	accountDatabase = append(accountDatabase, AccountDAO{
		accountID, firstName, nonceFirstName, lastName, nonceLastName,
		emailAddress, nonceEmailAddress, "", "participant", []string{}, []string{},
	})

	return accountID, nil
}

// FinalizeAccountCreation
func (a *AccountDBService) FinalizeAccountCreation(accountID string, encryptedpassword []byte, tx *sql.Tx) error {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			accountDatabase[i].Password = string(encryptedpassword)
			return nil
		}
	}

	return errors.New("account not found")
}

// GetDatabaseEntry gets an entry from the database
func (a *AccountDBService) GetDatabaseEntry(accountID string, tx *sql.Tx) (AccountDAO, error) {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			return accountDatabase[i], nil
		}
	}
	return AccountDAO{}, errors.New("not found")
}

// GetDatabaseEntryBasedOnMail when user logs in, id is not known
func (a *AccountDBService) GetDatabaseEntryBasedOnMail(username string, tx *sql.Tx) (AccountDAO, error) {
	for i := range accountDatabase {
		if accountDatabase[i].EmailAddress == username {
			return accountDatabase[i], nil
		}
	}
	return AccountDAO{}, errors.New("not found")
}

//GetProjectIDsAsParticipantForAccount gets the projects where the account is a participant
func (a *AccountDBService) GetProjectIDsAsParticipantForAccount(accountID string, tx *sql.Tx) ([]string, error) {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			return accountDatabase[i].ProjectsAsParticipant, nil
		}
	}
	return []string{}, nil
}

//GetProjectIDsAsResearcherForAccount gets the projects where the account is a researcher
func (a *AccountDBService) GetProjectIDsAsResearcherForAccount(accountID string, tx *sql.Tx) ([]string, error) {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			return accountDatabase[i].ProjectsAsResearcher, nil
		}
	}
	return []string{}, nil
}
