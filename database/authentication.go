package database

import (
	"crypto/rsa"
	"errors"
)

// AuthenticationDAO is the database model of the Authentication data
type AuthenticationDAO struct {
	AccountID       string
	SessionID       string
	AuthToken       string
	AuthTokenKey    rsa.PublicKey
	Refreshtoken    string
	RefreshTokenKey rsa.PublicKey
}

var authenticationTable []AuthenticationDAO

type IAuthenticationDBService interface {
	StoreSession(string, string, string, rsa.PublicKey, string, rsa.PublicKey) error
	UpdateSessionAuthToken(string, string, string, rsa.PublicKey) error
	GetSessionData(string, string) (AuthenticationDAO, error)
	RemoveSession(string, string) error
}

type AuthenticationDBService struct{}

//StoreSession Stores the session in the database
func (a *AuthenticationDBService) StoreSession(accountID string, sessionID string, authToken string, authPK rsa.PublicKey, refreshToken string, refreshPK rsa.PublicKey) error {
	authenticationTable = append(authenticationTable, AuthenticationDAO{accountID, sessionID, authToken, authPK, refreshToken, refreshPK})
	return nil
}

//UpdateSessionAuthToken Updates the authtoken of the session in the database
func (a *AuthenticationDBService) UpdateSessionAuthToken(accountID string, sessionID string, authToken string, authPK rsa.PublicKey) error {
	err := errors.New("could not find session")

	for i := range authenticationTable {
		if authenticationTable[i].AccountID == accountID {
			authenticationTable[i].AuthToken = authToken
			authenticationTable[i].AuthTokenKey = authPK
			return nil
		}
	}
	return err
}

// GetSessionData retrieves a session from the database
func (a *AuthenticationDBService) GetSessionData(accountID string, sessionID string) (AuthenticationDAO, error) {
	err := errors.New("could not find session")

	for i := range authenticationTable {
		if authenticationTable[i].AccountID == accountID && authenticationTable[i].SessionID == sessionID {
			return authenticationTable[i], nil
		}
	}
	return AuthenticationDAO{}, err
}

func (a *AuthenticationDBService) RemoveSession(accountID string, sessionID string) error {
	err := errors.New("could not find session")

	for i := range authenticationTable {
		if authenticationTable[i].AccountID == accountID && authenticationTable[i].SessionID == sessionID {
			authenticationTable[i] = authenticationTable[len(authenticationTable)-1]
			authenticationTable = authenticationTable[:len(authenticationTable)-1]
			return nil
		}
	}
	return err
}
