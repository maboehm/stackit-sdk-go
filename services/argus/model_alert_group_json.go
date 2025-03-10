/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type AlertGroupJson struct {
	Interval *string `json:"interval,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Rules *[]AlertRuleRecordJson `json:"rules"`
}
