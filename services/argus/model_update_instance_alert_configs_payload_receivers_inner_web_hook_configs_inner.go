/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type UpdateInstanceAlertConfigsPayloadReceiversInnerWebHookConfigsInner struct {
	// Microsoft Teams webhooks require special handling. If you set this property to true, it is treated as such
	MsTeams *bool `json:"msTeams,omitempty"`
	// The endpoint to send HTTP POST requests to. `Additional Validators:` * must be a syntactically valid url address
	Url *string `json:"url,omitempty"`
}
