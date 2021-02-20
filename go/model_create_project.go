/*
 * Test api for project
 *
 * This api is a test version to connect the backend and frontend
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateProject struct {

	// title of project
	ProjectTitle string `json:"projectTitle"`

	// participants of project
	Participants []Participant `json:"participants"`

	// tests to perform as part of this project
	Tests []string `json:"tests"`
}