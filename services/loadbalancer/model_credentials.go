/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type Credentials struct {
	// Credential name
	DisplayName *string `json:"displayName,omitempty"`
	// The password used for the ARGUS instance
	Password *string `json:"password,omitempty"`
	// The username used for the ARGUS instance
	Username *string `json:"username,omitempty"`
}
