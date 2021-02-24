/*
 * Test api for project
 *
 * This api is a test version to connect the backend and frontend
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LoginAccount struct {

	// type of grant type to log in with
	GrantType string `json:"grantType"`

	// account email address to log in with
	EmailAddress string `json:"emailAddress"`

	// password of the account to log in with
	Password string `json:"password,omitempty"`

	// refreshToken of the account to log in with
	RefreshToken string `json:"refreshToken,omitempty"`

	// session of account wrl the refreshToken
	SessionID string `json:"sessionID,omitempty"`

	// id of the account
	AccountID string `json:"accountID,omitempty"`
}
