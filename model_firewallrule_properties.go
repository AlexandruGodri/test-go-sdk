/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FirewallruleProperties struct {
	// A name of that resource
	Name string `json:"name,omitempty"`
	// The protocol for the rule. Property cannot be modified after creation (disallowed in update requests)
	Protocol string `json:"protocol"`
	// Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address
	SourceMac string `json:"sourceMac,omitempty"`
	// Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs
	SourceIp string `json:"sourceIp,omitempty"`
	// In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs
	TargetIp string `json:"targetIp,omitempty"`
	// Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes
	IcmpCode int32 `json:"icmpCode,omitempty"`
	// Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen. Value null allows all types
	IcmpType int32 `json:"icmpType,omitempty"`
	// Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd value null to allow all ports
	PortRangeStart int32 `json:"portRangeStart,omitempty"`
	// Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports
	PortRangeEnd int32 `json:"portRangeEnd,omitempty"`
}
