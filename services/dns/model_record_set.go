/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

type RecordSet struct {
	// if the record set is active or not
	Active *bool `json:"active,omitempty"`
	// comment
	Comment *string `json:"comment,omitempty"`
	// when record set creation finished
	// REQUIRED
	CreationFinished *string `json:"creationFinished"`
	// when record set creation started
	// REQUIRED
	CreationStarted *string `json:"creationStarted"`
	// Error shows error in case create/update/delete failed
	Error *string `json:"error,omitempty"`
	// rr set id
	// REQUIRED
	Id *string `json:"id"`
	// name of the record which should be a valid domain according to rfc1035 Section 2.3.4
	// REQUIRED
	Name *string `json:"name"`
	// records
	// REQUIRED
	Records *[]Record `json:"records"`
	// record set state
	// REQUIRED
	State *string `json:"state"`
	// time to live
	// REQUIRED
	Ttl *int32 `json:"ttl"`
	// record set type
	// REQUIRED
	Type *string `json:"type"`
	// when record set update/deletion finished
	// REQUIRED
	UpdateFinished *string `json:"updateFinished"`
	// when record set update/deletion started
	// REQUIRED
	UpdateStarted *string `json:"updateStarted"`
}
