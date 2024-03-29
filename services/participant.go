package services

import (
	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
)

type IParticipantService interface {
	CreateParticipant(string, string, string) (*models.Participant, error)
	LinkParticipantToAccount(string, string, string) (string, error)
	getAccountDBService() database.IAccountDBService
}

type ParticipantService struct{}

func (p *ParticipantService) CreateParticipant(firstName string, lastName string, emailAddress string) (*models.Participant, error) {

	email := models.EmailAddress{EmailAddress: emailAddress}

	validationError := email.IsValidEmailAddress()

	if validationError != nil {
		return &models.Participant{}, validationError
	}

	return &models.Participant{
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: email,
		AccountID:    "",
	}, nil
}

func (p *ParticipantService) LinkParticipantToAccount(emailAddress string, firstName string, lastName string) (string, error) {
	account, err1 := p.getAccountDBService().GetDatabaseEntryBasedOnMail(emailAddress)

	if err1 != nil {
		if account.AccountID == "" {
			// participant is not yet known in database
			// Create the account and send the participant an email so they can sign up
			accountID, err := p.getAccountDBService().CreateInitialParticipantAccount(firstName, lastName, emailAddress)
			// Include the newly created accountID in the email, so when the participant signs up via email, it can be linked to the account
			if err != nil {
				return "", err
			}

			// TODO send email that the user should create an account to participate in research
			sendEmailErr := SendEmail(emailAddress, "test1")
			if sendEmailErr != nil {
				// Do something to notify the researcher
			}

			return accountID, nil
		} else {
			// something went wrong, return error
			return "", err1
		}
	}

	// if account is not null, send an email to the participant that they have been invited for the research
	// send email
	sendEmailErr := SendEmail(emailAddress, "test2")

	if sendEmailErr != nil {
		// Do something to notify the researcher
	}

	// If account already exists, return the AccountID
	return account.AccountID, nil
}

func (p *ParticipantService) getAccountDBService() database.IAccountDBService {
	return &database.AccountDBService{}
}
