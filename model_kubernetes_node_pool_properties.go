/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type KubernetesNodePoolProperties struct {
	// A Kubernetes Node Pool Name. Valid Kubernetes Node Pool name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Name string `json:"name"`
	// A valid uuid of the datacenter on which user has access
	DatacenterId string `json:"datacenterId"`
	// Number of nodes part of the Node Pool
	NodeCount int32 `json:"nodeCount"`
	// A valid cpu family name
	CpuFamily string `json:"cpuFamily"`
	// Number of cores for node
	CoresCount int32 `json:"coresCount"`
	// RAM size for node, minimum size 2048MB is recommended. Ram size must be set to multiple of 1024MB.
	RamSize int32 `json:"ramSize"`
	// The availability zone in which the server should exist
	AvailabilityZone string `json:"availabilityZone"`
	// Hardware type of the volume
	StorageType string `json:"storageType"`
	// The size of the volume in GB. The size should be greater than 10GB.
	StorageSize int32 `json:"storageSize"`
	// The kubernetes version in which a nodepool is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions.
	K8sVersion string `json:"k8sVersion,omitempty"`
	MaintenanceWindow *KubernetesMaintenanceWindow `json:"maintenanceWindow,omitempty"`
}
