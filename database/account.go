package database

import (
	"errors"

	"github.com/google/uuid"
)

// AccountDAO data stored in the db
type AccountDAO struct {
	AccountID             string
	FirstName             string
	LastName              string
	EmailAddress          string
	Password              string
	AccountType           string
	ProjectsAsResearcher  []string // Only the IDs are stored here
	ProjectsAsParticipant []string // Only the IDs are stored here
}

var accountDatabase []AccountDAO = []AccountDAO{
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72", "Mark1", "Wijnbergen1", "mark.wijnbergen@hey.com", "$2a$04$cR6VcDXU4cSk0gSd1Hmr4euIfZxcYWunEUs1iMZu29JXeWOUN5O1m", "researcher",
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808", "497aeeaf-0d41-46c4-a5a1-8a88c7b61809"},
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808", "497aeeaf-0d41-46c4-a5a1-8a88c7b61809"}},
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b73", "Mark2", "Wijnbergen2", "wijnbergenmark@gmail.com", "$2a$04$cR6VcDXU4cSk0gSd1Hmr4euIfZxcYWunEUs1iMZu29JXeWOUN5O1m", "participant",
		[]string{},
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808"},
	},
}

// CreateInitialParticipantAccount creates the participant when the participant is for the first time added
// by a researcher for a project. This participant hasn't confirmed it's account yet, so has no password
func CreateInitialParticipantAccount(firstName string, lastName string, emailAddress string) (string, error) {
	accountID := uuid.New().String()

	accountDatabase = append(accountDatabase, AccountDAO{
		accountID, firstName, lastName, emailAddress, "", "participant", []string{}, []string{},
	})

	return accountID, nil
}

// FinalizeAccountCreation
func FinalizeAccountCreation(accountID string, encryptedpassword []byte) error {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			accountDatabase[i].Password = string(encryptedpassword)
			return nil
		}
	}

	return errors.New("Account not found")
}

// GetDatabaseEntry gets an entry from the database
func GetDatabaseEntry(accountID string) (AccountDAO, error) {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			return accountDatabase[i], nil
		}
	}
	return AccountDAO{}, errors.New("Not found")
}

// GetDatabaseEntryBasedOnMail when user logs in, id is not known
func GetDatabaseEntryBasedOnMail(username string) (AccountDAO, error) {
	for i := range accountDatabase {
		if accountDatabase[i].EmailAddress == username {
			return accountDatabase[i], nil
		}
	}
	return AccountDAO{}, errors.New("Not found")
}

//GetProjectIDsAsParticipantForAccount gets the projects where the account is a participant
func GetProjectIDsAsParticipantForAccount(accountID string) ([]string, error) {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			return accountDatabase[i].ProjectsAsParticipant, nil
		}
	}
	return []string{}, nil
}

//GetProjectIDsAsResearcherForAccount gets the projects where the account is a researcher
func GetProjectIDsAsResearcherForAccount(accountID string) ([]string, error) {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			return accountDatabase[i].ProjectsAsResearcher, nil
		}
	}
	return []string{}, nil
}
