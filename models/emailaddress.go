package models

import "github.com/badoux/checkmail"

type EmailAddress struct {
	EmailAddress string
}

func (e *EmailAddress) IsValidEmailAddress() error {

	err1 := checkmail.ValidateFormat(e.EmailAddress)

	if err1 != nil {
		return err1
	}

	return checkmail.ValidateHost(e.EmailAddress)
}
