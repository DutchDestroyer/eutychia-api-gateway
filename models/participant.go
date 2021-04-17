package models

import (
	"errors"
)

type Participant struct {
	FirstName    string
	LastName     string
	EmailAddress EmailAddress
	AccountID    string
}

func (p *Participant) IsValidParticipant() error {

	if p.FirstName != "" || p.LastName != "" {
		return errors.New("name field is empty")
	}

	emailErr := p.EmailAddress.IsValidEmailAddress()
	if emailErr != nil {
		return emailErr
	}

	return nil
}
