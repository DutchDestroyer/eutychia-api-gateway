package models

import "github.com/badoux/checkmail"

type IEmailAddress interface {
	GetEmailAddress() *EmailAddress
	IsValidEmailAddress() error
}

type EmailAddress struct {
	EmailAddress string
}

func (e *EmailAddress) GetEmailAddress() *EmailAddress {
	return e
}

func (e *EmailAddress) IsValidEmailAddress() error {

	err1 := checkmail.ValidateFormat(e.EmailAddress)

	if err1 != nil {
		return err1
	}

	return nil //checkmail.ValidateHost(e.EmailAddress)
}
