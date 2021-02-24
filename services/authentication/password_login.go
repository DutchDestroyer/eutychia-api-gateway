package authentication

import (
	"errors"

	accountmodels "github.com/DutchDestroyer/eutychia-api-gateway/account/models"
	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"golang.org/x/crypto/bcrypt"
)

// IsValidPasswordLogin validates password
func IsValidPasswordLogin(acc accountmodels.Account) (string, error) {
	// Find email address in db
	accountDAO, errDAO := database.GetDatabaseEntryBasedOnMail(acc.Username)

	if errDAO != nil {
		return "", errDAO
	}

	saltedPassword, err := bcrypt.GenerateFromPassword([]byte(accountDAO.Password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	if IsValidPassword(saltedPassword, []byte(acc.Password)) {
		return accountDAO.AccountID, nil
	}

	return "", errors.New("Invalid email password combination")
}

// IsValidPassword checks whether the password is valid
func IsValidPassword(hashedPwd []byte, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)

	return err == nil
}
