/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type UpdateInstanceGrafanaConfigsPayloadGenericOauth struct {
	// Set api_url to the resource that returns OpenID UserInfo compatible information.
	// REQUIRED
	ApiUrl *string `json:"apiUrl"`
	// Authentication endpoint of idp.
	// REQUIRED
	AuthUrl *string `json:"authUrl"`
	// enable or disable generic oauth login
	// REQUIRED
	Enabled *bool `json:"enabled"`
	// Display name for the oAuth provider
	Name *string `json:"name,omitempty"`
	// Oauth client id for auth endpoint.
	// REQUIRED
	OauthClientId *string `json:"oauthClientId"`
	// Oauth client secret for auth endpoint.
	// REQUIRED
	OauthClientSecret *string `json:"oauthClientSecret"`
	// Grafana checks for the presence of a role using the JMESPath specified via the role_attribute_path configuration option. The JMESPath is applied to the id_token first. If there is no match, then the UserInfo endpoint specified via the api_url configuration option is tried next. The result after evaluation of the role_attribute_path JMESPath expression should be a valid Grafana role, for example, Viewer, Editor or Admin For example: contains(roles[\\*], 'grafana-admin') && 'Admin' || contains(roles[\\*], 'grafana-editor') && 'Editor' || contains(roles[\\*], 'grafana-viewer') && 'Viewer'
	// REQUIRED
	RoleAttributePath *string `json:"roleAttributePath"`
	// If  therole_attribute_path property does not return a role, then the user is assigned the Viewer role by default. You can disable the role assignment by setting role_attribute_strict = true. It denies user access if no role or an invalid role is returned.
	RoleAttributeStrict *bool `json:"roleAttributeStrict,omitempty"`
	// Space seperated list of scopes of the token
	Scopes *string `json:"scopes,omitempty"`
	// Token endpoint of the idp.
	// REQUIRED
	TokenUrl *string `json:"tokenUrl"`
	// enable or disable Proof Key for Code Exchange
	UsePkce *bool `json:"usePkce,omitempty"`
}
