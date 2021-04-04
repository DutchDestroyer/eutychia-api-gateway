package database

import "errors"

// AccountDAO data stored in the db
type AccountDAO struct {
	AccountID             string
	Username              string
	Password              string
	AccountType           string
	ProjectsAsResearcher  []string // Only the IDs are stored here
	ProjectsAsParticipant []string // Only the IDs are stored here
}

var accountDatabase []AccountDAO = []AccountDAO{
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72", "mark.wijnbergen@hey.com", "test123", "researcher",
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808", "497aeeaf-0d41-46c4-a5a1-8a88c7b61809"},
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808", "497aeeaf-0d41-46c4-a5a1-8a88c7b61809"}},
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b73", "wijnbergenmark@gmail.com", "test123", "participant",
		[]string{},
		[]string{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "497aeeaf-0d41-46c4-a5a1-8a88c7b61808"}},
}

// CreateDatabaseEntry creates a new entry in the database
func CreateDatabaseEntry(accountID string, username string, password string, accountType string) {
	accountDatabase = append(accountDatabase, AccountDAO{accountID, username, password, accountType, []string{}, []string{}})
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
		if accountDatabase[i].Username == username {
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
