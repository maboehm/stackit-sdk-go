/*
STACKIT Opensearch API

The STACKIT Opensearch API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opensearch

type CreateInstancePayload struct {
	// REQUIRED
	InstanceName *string             `json:"instanceName"`
	Parameters   *InstanceParameters `json:"parameters,omitempty"`
	// REQUIRED
	PlanId *string `json:"planId"`
}
