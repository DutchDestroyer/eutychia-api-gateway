/*
 * Test api for project
 *
 * This api is a test version to connect the backend and frontend
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

//Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{Code: code, Body: body}
}

