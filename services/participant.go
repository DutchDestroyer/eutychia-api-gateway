package services

import "github.com/DutchDestroyer/eutychia-api-gateway/database"

func linkParticipantToAccount(emailAddress string, firstName string, lastName string) (string, error) {
	account, err1 := database.GetDatabaseEntryBasedOnMail(emailAddress)

	if err1 != nil {
		if err1.Error() == "not found" {
			// participant is not yet known in database
			// Create the account and send the participant an email so they can sign up
			accountID, err := database.CreateInitialParticipantAccount(firstName, lastName, emailAddress)
			// Include the newly created accountID in the email, so when the participant signs up via email, it can be linked to the account
			if err != nil {
				return "", err
			}

			// TODO send email

			return accountID, nil
		} else {
			// something went wrong, return error
			return "", err1
		}
	}

	// if account is not null, send an email to the participant that they have been invited for the research
	// send email

	// If account already exists, return the AccountID
	return account.AccountID, nil
}
