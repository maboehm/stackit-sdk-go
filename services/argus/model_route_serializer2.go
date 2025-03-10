/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type RouteSerializer2 struct {
	Continue      *bool              `json:"continue,omitempty"`
	GroupBy       *[]string          `json:"groupBy,omitempty"`
	GroupInterval *string            `json:"groupInterval,omitempty"`
	GroupWait     *string            `json:"groupWait,omitempty"`
	Match         *map[string]string `json:"match,omitempty"`
	MatchRe       *map[string]string `json:"matchRe,omitempty"`
	Matchers      *[]string          `json:"matchers,omitempty"`
	// REQUIRED
	Receiver       *string              `json:"receiver"`
	RepeatInterval *string              `json:"repeatInterval,omitempty"`
	Routes         *[]map[string]string `json:"routes,omitempty"`
}
