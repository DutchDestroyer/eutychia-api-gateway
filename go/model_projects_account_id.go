/*
 * Test api for project
 *
 * This api is a test version to connect the backend and frontend
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ProjectsAccountId - response with the projects of a specific account
type ProjectsAccountId struct {

	// an array with all projects
	Projects []Project `json:"projects,omitempty"`
}
