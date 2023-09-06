/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type UpdateScrapeConfigPayloadStaticConfigsInner struct {
	// Labels assigned to all metrics scraped from the targets. `Additional Validators:` * should not contain more than 5 keys * each key and value should not be longer than 200 characters
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The targets specified by the static config.
	// REQUIRED
	Targets *[]string `json:"targets"`
}
