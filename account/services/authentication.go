package accountservices

import (
	"golang.org/x/crypto/bcrypt"
)

// IsValidPassword checks whether the password is valid
func IsValidPassword(hashedPwd []byte, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		return false
	}

	return true
}
