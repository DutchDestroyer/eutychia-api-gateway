package authentication

import (
	"errors"

	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	models "github.com/DutchDestroyer/eutychia-api-gateway/models"
	"golang.org/x/crypto/bcrypt"
)

// IsValidPasswordLogin validates password
func IsValidPasswordLogin(acc models.Account) (database.AccountDAO, error) {
	// Find email address in db
	accountDAO, errDAO := database.GetDatabaseEntryBasedOnMail(acc.Username)

	if errDAO != nil {
		return database.AccountDAO{}, errDAO
	}

	saltedPassword, err := EncryptPassword(accountDAO.Password)

	if err != nil {
		return database.AccountDAO{}, err
	}

	if IsValidPassword(saltedPassword, []byte(acc.Password)) {
		return accountDAO, nil
	}

	return database.AccountDAO{}, errors.New("Invalid email password combination")
}

// IsValidPassword checks whether the password is valid
func IsValidPassword(hashedPwd []byte, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)

	return err == nil
}

func EncryptPassword(password string) ([]byte, error) {
	pw, err1 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err1 != nil {
		return []byte{}, err1
	}

	return pw, nil
}
