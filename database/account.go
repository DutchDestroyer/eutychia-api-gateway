package database

import "errors"

// AccountDAO data stored in the db
type AccountDAO struct {
	AccountID             string
	Username              string
	Password              string
	AccountType           string
	ProjectsAsResearcher  []string
	ProjectsAsParticipant []string
}

var accountDatabase []AccountDAO = []AccountDAO{
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72", "mark.wijnbergen@hey.com", "test123", "researcher", []string{}, []string{}},
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b73", "wijnbergenmark@gmail.com", "test123", "participant", []string{}, []string{}},
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
