/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

type ClusterResponse struct {
	Extensions  *Extension   `json:"extensions,omitempty"`
	Hibernation *Hibernation `json:"hibernation,omitempty"`
	// REQUIRED
	Kubernetes  *Kubernetes  `json:"kubernetes"`
	Maintenance *Maintenance `json:"maintenance,omitempty"`
	Name        *string      `json:"name,omitempty"`
	// REQUIRED
	Nodepools *[]Nodepool    `json:"nodepools"`
	Status    *ClusterStatus `json:"status,omitempty"`
}
