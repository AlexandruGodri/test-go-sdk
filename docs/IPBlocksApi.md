# {{classname}}

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IpblocksGet**](IPBlocksApi.md#IpblocksGet) | **Get** /ipblocks | List IP Blocks 
[**IpblocksIpblockIdDelete**](IPBlocksApi.md#IpblocksIpblockIdDelete) | **Delete** /ipblocks/{ipblockId} | Delete IP Block
[**IpblocksIpblockIdGet**](IPBlocksApi.md#IpblocksIpblockIdGet) | **Get** /ipblocks/{ipblockId} | Retrieve an IP Block
[**IpblocksIpblockIdPatch**](IPBlocksApi.md#IpblocksIpblockIdPatch) | **Patch** /ipblocks/{ipblockId} | Partially modify IP Block
[**IpblocksIpblockIdPut**](IPBlocksApi.md#IpblocksIpblockIdPut) | **Put** /ipblocks/{ipblockId} | Modify IP Block
[**IpblocksPost**](IPBlocksApi.md#IpblocksPost) | **Post** /ipblocks | Reserve IP Block

# **IpblocksGet**
> IpBlocks IpblocksGet(ctx, optional)
List IP Blocks 

Retrieve a list of all reserved IP Blocks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPBlocksApiIpblocksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPBlocksApiIpblocksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**IpBlocks**](IpBlocks.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksIpblockIdDelete**
> interface{} IpblocksIpblockIdDelete(ctx, ipblockId, optional)
Delete IP Block

Removes the specific IP Block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipblockId** | **string**|  | 
 **optional** | ***IPBlocksApiIpblocksIpblockIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPBlocksApiIpblocksIpblockIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksIpblockIdGet**
> IpBlock IpblocksIpblockIdGet(ctx, ipblockId, optional)
Retrieve an IP Block

Retrieves the attributes of a given IP Block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipblockId** | **string**|  | 
 **optional** | ***IPBlocksApiIpblocksIpblockIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPBlocksApiIpblocksIpblockIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksIpblockIdPatch**
> IpBlock IpblocksIpblockIdPatch(ctx, body, ipblockId, optional)
Partially modify IP Block

You can use update attributes of a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpBlockProperties**](IpBlockProperties.md)| IP Block to be modified | 
  **ipblockId** | **string**|  | 
 **optional** | ***IPBlocksApiIpblocksIpblockIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPBlocksApiIpblocksIpblockIdPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksIpblockIdPut**
> IpBlock IpblocksIpblockIdPut(ctx, body, ipblockId, optional)
Modify IP Block

You can use update attributes of a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpBlock**](IpBlock.md)| IP Block to be modified | 
  **ipblockId** | **string**|  | 
 **optional** | ***IPBlocksApiIpblocksIpblockIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPBlocksApiIpblocksIpblockIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksPost**
> IpBlock IpblocksPost(ctx, body, optional)
Reserve IP Block

This will reserve a new IP Block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpBlock**](IpBlock.md)| IP Block to be reserved | 
 **optional** | ***IPBlocksApiIpblocksPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPBlocksApiIpblocksPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

