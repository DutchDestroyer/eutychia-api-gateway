package accountservices

import accountmodels "github.com/DutchDestroyer/eutychia-api-gateway/account/models"

func CreateAccount(emailAddress string, password string, token string) *accountmodels.Account {

	return &accountmodels.Account{
		Username: emailAddress,
		Password: password,
		RefreshToken: token,
	}

}
