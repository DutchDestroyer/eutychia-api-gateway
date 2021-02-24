package database

import "errors"

// AccountDatabase data stored in the db
type AccountDatabase struct {
	AccountID   string
	Username    string
	Password    string
	AccountType string
}

var accountDatabase []AccountDatabase = []AccountDatabase{
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72", "mark.wijnbergen@hey.com", "test123", "researcher"},
	{"683c5de1-5172-4a94-bd3b-2d4bf58b6b73", "wijnbergenmark@gmail.com", "test123", "participant"},
}

// CreateDatabaseEntry creates a new entry in the database
func CreateDatabaseEntry(accountID string, username string, password string, accountType string) {
	accountDatabase = append(accountDatabase, AccountDatabase{accountID, username, password, accountType})
}

// GetDatabaseEntry gets an entry from the database
func GetDatabaseEntry(accountID string) (AccountDatabase, error) {
	for i := range accountDatabase {
		if accountDatabase[i].AccountID == accountID {
			return accountDatabase[i], nil
		}
	}
	return AccountDatabase{}, errors.New("Not found")
}

// GetDatabaseEntryBasedOnMail when user logs in, id is not known
func GetDatabaseEntryBasedOnMail(username string) (AccountDatabase, error) {
	for i := range accountDatabase {
		if accountDatabase[i].Username == username {
			return accountDatabase[i], nil
		}
	}
	return AccountDatabase{}, errors.New("Not found")
}
