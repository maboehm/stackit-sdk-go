/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type ProjectInstanceFull struct {
	Error *NullableString `json:"error,omitempty"`
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	Instance *string `json:"instance"`
	Name     *string `json:"name,omitempty"`
	// REQUIRED
	PlanName *string `json:"planName"`
	// REQUIRED
	ServiceName *string `json:"serviceName"`
	// REQUIRED
	Status *string `json:"status"`
}
