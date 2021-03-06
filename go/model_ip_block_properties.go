/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type IpBlockProperties struct {
	// A collection of IPs associated with the IP Block
	Ips []string `json:"ips,omitempty"`
	// Location of that IP Block. Property cannot be modified after creation (disallowed in update requests)
	Location string `json:"location"`
	// The size of the IP block
	Size int32 `json:"size"`
	// A name of that resource
	Name string `json:"name,omitempty"`
	// Read-Only attribute. Lists consumption detail of an individual ip
	IpConsumers []IpConsumer `json:"ipConsumers,omitempty"`
}
