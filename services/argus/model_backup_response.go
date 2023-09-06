/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type BackupResponse struct {
	// REQUIRED
	AlertConfigBackups *[]string `json:"alertConfigBackups"`
	// REQUIRED
	AlertRulesBackups *[]string `json:"alertRulesBackups"`
	// REQUIRED
	GrafanaBackups *[]string `json:"grafanaBackups"`
	// REQUIRED
	Message *string `json:"message"`
	// REQUIRED
	ScrapeConfigBackups *[]string `json:"scrapeConfigBackups"`
}
