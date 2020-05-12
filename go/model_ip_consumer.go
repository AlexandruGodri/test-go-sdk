/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type IpConsumer struct {

	Ip string `json:"ip,omitempty"`

	Mac string `json:"mac,omitempty"`

	NicId string `json:"nicId,omitempty"`

	ServerId string `json:"serverId,omitempty"`

	ServerName string `json:"serverName,omitempty"`

	DatacenterId string `json:"datacenterId,omitempty"`

	DatacenterName string `json:"datacenterName,omitempty"`
}
