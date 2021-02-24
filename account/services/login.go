package accountservices

import (
	"errors"

	accountmodels "github.com/DutchDestroyer/eutychia-api-gateway/account/models"
)

// IsValidPasswordLogin validates password
func IsValidPasswordLogin(acc accountmodels.Account) error {
	// Find email address in db
	saltedPassword := []byte("test123")
	// Find pw in database
	// Compare data with db

	if IsValidPassword(saltedPassword, []byte(acc.Password)) {
		return nil
	}

	return errors.New("Invalid email password combination")
}

// IsValidToken determines whether the token validation is done correctly
func IsValidToken(acc accountmodels.Account) error {

	return errors.New("Invalid token")
}
