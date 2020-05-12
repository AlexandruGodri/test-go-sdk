/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Resources assigned to this group.
type ResourceGroups struct {
	// The resource's unique identifier
	Id string `json:"id,omitempty"`
	// The type of the resource
	Type_ *AllOfResourceGroupsType_ `json:"type,omitempty"`
	// URL to the object representation (absolute path)
	Href string `json:"href,omitempty"`
	// Array of items in that collection
	Items []Resource `json:"items,omitempty"`
}
