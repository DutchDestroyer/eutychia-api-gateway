package uuidvalidator

import "regexp"

// IsCorrectUUID checks whether uuid is correct
func IsCorrectUUID(uuid string) bool {
	return regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$").MatchString(uuid)
}
