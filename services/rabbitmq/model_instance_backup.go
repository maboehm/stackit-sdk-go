/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

type InstanceBackup struct {
	Downloadable *bool `json:"downloadable,omitempty"`
	// REQUIRED
	FinishedAt *string `json:"finished_at"`
	// REQUIRED
	Id   *int64 `json:"id"`
	Size *int64 `json:"size,omitempty"`
	// REQUIRED
	Status      *string `json:"status"`
	TriggeredAt *string `json:"triggered_at,omitempty"`
}
