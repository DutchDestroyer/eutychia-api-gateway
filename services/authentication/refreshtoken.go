package authentication

func RefreshAccessToken(accountID string, sessionID string, refreshToken string) (string, error) {
	err1 := ValidateToken(refreshToken, accountID, sessionID, "refreshToken")
	if err1 != nil {
		return "", err1
	}

	newAuthToken, err2 := UpdateAccountAuthentication(accountID, sessionID)

	if err2 != nil {
		return "", nil
	}

	return newAuthToken, nil
}
