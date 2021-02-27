package authentication

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"time"

	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

// CreateToken Creates a token with a specific time
func CreateToken(accountID string, sessionID string, timespan time.Duration) (string, rsa.PublicKey, error) {
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

// ParseToken parses the token
func ParseToken(token string) (jwt.Token, error) {
	return jwt.Parse([]byte(token))
}

// ValidateToken validates the token
func ValidateToken(token string, accountID string, sessionID string, tokenType string) error {

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
