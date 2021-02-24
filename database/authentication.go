package database

type AuthenticationData struct {
	accountID    string
	sessionID    string
	authToken    string
	refreshtoken string
}

var authenticationTable []AuthenticationData

//StoreSession Stores the session in the database
func StoreSession(accountID string, sessionID string, authToken string, refreshToken string) error {
	authenticationTable = append(authenticationTable, AuthenticationData{accountID, sessionID, authToken, refreshToken})
	return nil
}
