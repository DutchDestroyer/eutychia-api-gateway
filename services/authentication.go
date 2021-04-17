package services

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"time"

	"errors"

	"github.com/google/uuid"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/DutchDestroyer/eutychia-api-gateway/models"

	"github.com/DutchDestroyer/eutychia-api-gateway/database"
)

// CreateAccountAuthentication create authentication of the account
func CreateAccountAuthentication(account *models.Account) error {
	sessionID := uuid.New().String()

	authToken, authTokenKey, authErr := createAuthToken(account.AccountID, sessionID)
	if authErr != nil {
		return authErr
	}

	refreshToken, refreshTokenKey, refreshErr := createRefreshToken(account.AccountID, sessionID)
	if refreshErr != nil {
		return refreshErr
	}

	dbErr := database.StoreSession(account.AccountID, sessionID, authToken, authTokenKey, refreshToken, refreshTokenKey)

	if dbErr != nil {
		return dbErr
	}

	account.SessionID = sessionID
	account.AuthToken = authToken
	account.RefreshToken = refreshToken

	return nil
}

// UpdateAccountAuthentication create authentication of the account when logging in with refreshtoken
func UpdateAccountAuthentication(accountID string, sessionID string) (string, error) {

	authToken, authTokenKey, authErr := createAuthToken(accountID, sessionID)
	if authErr != nil {
		return "", authErr
	}

	dbErr := database.UpdateSessionAuthToken(accountID, sessionID, authToken, authTokenKey)
	if dbErr != nil {
		return "", dbErr
	}

	return authToken, nil
}

//LogOutWithAccount logs out the provided session of the provided account
func LogOutWithAccount(sessionID string, accountID string, accessToken string) error {
	return database.RemoveSession(accountID, sessionID)
}

// IsValidPasswordLogin validates password
func IsValidPasswordLogin(acc models.Account) (database.AccountDAO, error) {

	emailAddress := models.EmailAddress(acc.Username)

	emailValidation := emailAddress.IsValidEmailAddress()

	if emailValidation != nil {
		return database.AccountDAO{}, emailValidation
	}

	// Find email address in db
	accountDAO, errDAO := database.GetDatabaseEntryBasedOnMail(emailAddress.EmailAddress)

	if errDAO != nil {
		return database.AccountDAO{}, errDAO
	}

	log.Printf(accountDAO.Password)

	if isValidPassword([]byte(accountDAO.Password), []byte(acc.Password)) {
		return accountDAO, nil
	}

	return database.AccountDAO{}, errors.New("Invalid email password combination")
}

func RefreshAccessToken(accountID string, sessionID string, refreshToken string) (string, error) {
	newAuthToken, err2 := UpdateAccountAuthentication(accountID, sessionID)

	if err2 != nil {
		return "", nil
	}

	return newAuthToken, nil
}

// IsValidTokenLogin validates the token
func IsValidTokenLogin(token string, accountID string, sessionID string, tokenType string) error {

	sessionData, errDB := database.GetSessionData(accountID, sessionID)

	if errDB != nil {
		return errDB
	}

	var tokenKey rsa.PublicKey

	if tokenType == "refreshToken" {
		tokenKey = sessionData.RefreshTokenKey
	} else if tokenType == "authToken" {
		tokenKey = sessionData.AuthTokenKey
	} else {
		return errors.New("invalid token")
	}

	_, err := jwt.Parse([]byte(token), jwt.WithValidate(true), jwt.WithVerify(jwa.RS256, tokenKey))

	return err
}

func createAuthToken(accountID string, sessionID string) (string, rsa.PublicKey, error) {
	return createToken(accountID, sessionID, time.Duration(time.Duration.Minutes(15)))
}

func createRefreshToken(accountID string, sessionID string) (string, rsa.PublicKey, error) {
	return createToken(accountID, sessionID, time.Duration(time.Duration.Hours(24)))
}

// createToken Creates a token with a specific time
func createToken(accountID string, sessionID string, timespan time.Duration) (string, rsa.PublicKey, error) {
	alg := jwa.RS256
	key, errGenerate := rsa.GenerateKey(rand.Reader, 2048)
	if errGenerate != nil {
		return "", rsa.PublicKey{}, errGenerate
	}

	// store the public key in the db

	token := jwt.New()
	token.Set("accountID", accountID)
	token.Set("sessionID", sessionID)
	token.Set("expiryDate", time.Now().UTC().Add(timespan))
	signed, errSign := jwt.Sign(token, alg, key)

	if errSign != nil {
		return "", rsa.PublicKey{}, errSign
	}

	return string(signed), key.PublicKey, nil
}

//encryptPassword encrypts a password
func encryptPassword(password string) ([]byte, error) {
	pw, err1 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err1 != nil {
		return []byte{}, err1
	}

	return pw, nil
}

func isValidPassword(dbPassword []byte, givenPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(dbPassword, givenPassword)

	return err == nil
}
