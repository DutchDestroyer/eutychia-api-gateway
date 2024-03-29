/*
 * Test api for project
 *
 * This api is a test version to connect the backend and frontend
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AccountCreationFinalize struct {

	// email address of the participant
	EmailAddress string `json:"emailAddress"`

	// first name of the user
	FirstName string `json:"firstName"`

	// last name of the user
	LastName string `json:"lastName"`

	// password of the user
	Password string `json:"password"`
}
