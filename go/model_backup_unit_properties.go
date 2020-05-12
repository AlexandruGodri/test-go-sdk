/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BackupUnitProperties struct {
	// A name of that resource (only alphanumeric characters are acceptable)
	Name string `json:"name"`
	// the password associated to that resource
	Password string `json:"password,omitempty"`
	// The email associated with the backup unit. Bear in mind that this email does not be the same email as of the user.
	Email string `json:"email,omitempty"`
}
