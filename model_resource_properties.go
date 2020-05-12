/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResourceProperties struct {
	// name of the resource
	Name string `json:"name,omitempty"`
	// Boolean value representing if the resource is multi factor protected or not e.g. using two factor protection. Currently only Data Centers and Snapshots are allowed to be multi factor protected, The value of attribute if null is intentional and it means that the resource doesn't support multi factor protection at all.
	SecAuthProtection bool `json:"secAuthProtection,omitempty"`
}
