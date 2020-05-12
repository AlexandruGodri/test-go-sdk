# KubernetesNodePoolProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A Kubernetes Node Pool Name. Valid Kubernetes Node Pool name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | [default to null]
**DatacenterId** | **string** | A valid uuid of the datacenter on which user has access | [default to null]
**NodeCount** | **int32** | Number of nodes part of the Node Pool | [default to null]
**CpuFamily** | **string** | A valid cpu family name | [default to null]
**CoresCount** | **int32** | Number of cores for node | [default to null]
**RamSize** | **int32** | RAM size for node, minimum size 2048MB is recommended. Ram size must be set to multiple of 1024MB. | [default to null]
**AvailabilityZone** | **string** | The availability zone in which the server should exist | [default to null]
**StorageType** | **string** | Hardware type of the volume | [default to null]
**StorageSize** | **int32** | The size of the volume in GB. The size should be greater than 10GB. | [default to null]
**K8sVersion** | **string** | The kubernetes version in which a nodepool is running. This imposes restrictions on what kubernetes versions can be run in a cluster&#x27;s nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions. | [optional] [default to null]
**MaintenanceWindow** | [***KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

