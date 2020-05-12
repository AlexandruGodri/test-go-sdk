# ServerProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] [default to null]
**Cores** | **int32** | The total number of cores for the server | [default to null]
**Ram** | **int32** | The amount of memory for the server in MB, e.g. 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB. | [default to null]
**AvailabilityZone** | **string** | The availability zone in which the server should exist | [optional] [default to null]
**VmState** | **string** | Status of the virtual Machine | [optional] [default to null]
**BootCdrom** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**BootVolume** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**CpuFamily** | **string** | Cpu family of pserver | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

