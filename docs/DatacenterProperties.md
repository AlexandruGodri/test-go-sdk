# DatacenterProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] [default to null]
**Description** | **string** | A description for the datacenter, e.g. staging, production | [optional] [default to null]
**Location** | **string** | The physical location where the datacenter will be created. This will be where all of your servers live. Property cannot be modified after datacenter creation (disallowed in update requests) | [default to null]
**Version** | **int32** | The version of that Data Center. Gets incremented with every change | [optional] [default to null]
**Features** | **[]string** | List of features supported by the location this data center is part of | [optional] [default to null]
**SecAuthProtection** | **bool** | Boolean value representing if the data center requires extra protection e.g. two factor protection | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

