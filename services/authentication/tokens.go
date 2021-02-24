package authentication

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

// CreateToken Creates a token with a specific time
func CreateToken(accountID string, sessionID string, timespan time.Duration) (string, error) {
	alg := jwa.RS256
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	// store the public key in the db

	token := jwt.New()
	token.Set("accountID", accountID)
	token.Set("sessionID", sessionID)
	token.Set("expiryDate", time.Now().UTC().Add(timespan))
	token.Set("pK", key.PublicKey)
	signed, err := jwt.Sign(token, alg, key)

	return string(signed), err
}

// ValidateToken validates the token
func ValidateToken(token string, accountID string, sessionID string) error {

	publicKey, err := getPublicKey(accountID, sessionID)

	if err != nil {
		return err
	}

	byteToken := []byte(token)
	_, err = jwt.Parse(byteToken, jwt.WithValidate(true), jwt.WithVerify(jwa.RS256, publicKey))

	return err
}

func getPublicKey(accountID string, sessionID string) (rsa.PublicKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	return key.PublicKey, err
}
