/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type NetworkCheckResponse struct {
	// REQUIRED
	Message *string `json:"message"`
	// REQUIRED
	NetworkChecks *[]NetworkCheckChildResponse `json:"networkChecks"`
}
