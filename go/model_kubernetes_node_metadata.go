/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type KubernetesNodeMetadata struct {
	// Resource's Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11 . Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter. 
	Etag string `json:"etag,omitempty"`
	// The last time the resource was created
	CreatedDate time.Time `json:"createdDate,omitempty"`
	// The last time the resource has been modified
	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`
	// State of the resource.
	State string `json:"state,omitempty"`
	// The last time the software updated on node.
	LastSoftwareUpdatedDate time.Time `json:"lastSoftwareUpdatedDate,omitempty"`
}
