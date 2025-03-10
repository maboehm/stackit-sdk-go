/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

type CreateUserPayload struct {
	// A user chosen description to differentiate between multiple users.
	// REQUIRED
	Description *string `json:"description"`
	// Is true if the user has write access to the secrets engine. Is false for a read-only user.
	// REQUIRED
	Write *bool `json:"write"`
}
