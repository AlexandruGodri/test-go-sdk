# {{classname}}

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersDatacenterIdLansGet**](LanApi.md#DatacentersDatacenterIdLansGet) | **Get** /datacenters/{datacenterId}/lans | List Lans
[**DatacentersDatacenterIdLansLanIdDelete**](LanApi.md#DatacentersDatacenterIdLansLanIdDelete) | **Delete** /datacenters/{datacenterId}/lans/{lanId} | Delete a Lan.
[**DatacentersDatacenterIdLansLanIdGet**](LanApi.md#DatacentersDatacenterIdLansLanIdGet) | **Get** /datacenters/{datacenterId}/lans/{lanId} | Retrieve a Lan
[**DatacentersDatacenterIdLansLanIdNicsGet**](LanApi.md#DatacentersDatacenterIdLansLanIdNicsGet) | **Get** /datacenters/{datacenterId}/lans/{lanId}/nics | List Lan Members 
[**DatacentersDatacenterIdLansLanIdNicsNicIdGet**](LanApi.md#DatacentersDatacenterIdLansLanIdNicsNicIdGet) | **Get** /datacenters/{datacenterId}/lans/{lanId}/nics/{nicId} | Retrieve a nic attached to lan
[**DatacentersDatacenterIdLansLanIdNicsPost**](LanApi.md#DatacentersDatacenterIdLansLanIdNicsPost) | **Post** /datacenters/{datacenterId}/lans/{lanId}/nics | Attach a nic
[**DatacentersDatacenterIdLansLanIdPatch**](LanApi.md#DatacentersDatacenterIdLansLanIdPatch) | **Patch** /datacenters/{datacenterId}/lans/{lanId} | Partially modify a Lan
[**DatacentersDatacenterIdLansLanIdPut**](LanApi.md#DatacentersDatacenterIdLansLanIdPut) | **Put** /datacenters/{datacenterId}/lans/{lanId} | Modify a Lan
[**DatacentersDatacenterIdLansPost**](LanApi.md#DatacentersDatacenterIdLansPost) | **Post** /datacenters/{datacenterId}/lans | Create a Lan

# **DatacentersDatacenterIdLansGet**
> Lans DatacentersDatacenterIdLansGet(ctx, datacenterId, optional)
List Lans

Retrieve a list of LANs within the datacenter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***LanApiDatacentersDatacenterIdLansGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Lans**](Lans.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLansLanIdDelete**
> interface{} DatacentersDatacenterIdLansLanIdDelete(ctx, datacenterId, lanId, optional)
Delete a Lan.

Removes the specific Lan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***LanApiDatacentersDatacenterIdLansLanIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansLanIdDeleteOpts struct
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

# **DatacentersDatacenterIdLansLanIdGet**
> Lan DatacentersDatacenterIdLansLanIdGet(ctx, datacenterId, lanId, optional)
Retrieve a Lan

Retrieves the attributes of a given LAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***LanApiDatacentersDatacenterIdLansLanIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansLanIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Lan**](Lan.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLansLanIdNicsGet**
> LanNics DatacentersDatacenterIdLansLanIdNicsGet(ctx, datacenterId, lanId, optional)
List Lan Members 

You can retrieve a list of nics attached to a lan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***LanApiDatacentersDatacenterIdLansLanIdNicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansLanIdNicsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LanNics**](LanNics.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLansLanIdNicsNicIdGet**
> Nic DatacentersDatacenterIdLansLanIdNicsNicIdGet(ctx, datacenterId, lanId, nicId, optional)
Retrieve a nic attached to lan

This will retrieve the properties of an attached nic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **lanId** | **string**| The unique ID of the LAN | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***LanApiDatacentersDatacenterIdLansLanIdNicsNicIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansLanIdNicsNicIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nic**](Nic.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLansLanIdNicsPost**
> Nic DatacentersDatacenterIdLansLanIdNicsPost(ctx, body, datacenterId, lanId, optional)
Attach a nic

This will attach a pre-existing nic to a lan. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Nic**](Nic.md)| Nic to be attached | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***LanApiDatacentersDatacenterIdLansLanIdNicsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansLanIdNicsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nic**](Nic.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLansLanIdPatch**
> Lan DatacentersDatacenterIdLansLanIdPatch(ctx, body, datacenterId, lanId, optional)
Partially modify a Lan

You can use update attributes of a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LanProperties**](LanProperties.md)| Modified Lan | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***LanApiDatacentersDatacenterIdLansLanIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansLanIdPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Lan**](Lan.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLansLanIdPut**
> Lan DatacentersDatacenterIdLansLanIdPut(ctx, body, datacenterId, lanId, optional)
Modify a Lan

You can use update attributes of a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Lan**](Lan.md)| Modified Lan | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***LanApiDatacentersDatacenterIdLansLanIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansLanIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Lan**](Lan.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLansPost**
> LanPost DatacentersDatacenterIdLansPost(ctx, body, datacenterId, optional)
Create a Lan

Creates a LAN within the datacenter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LanPost**](LanPost.md)| Lan to be created | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***LanApiDatacentersDatacenterIdLansPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LanApiDatacentersDatacenterIdLansPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LanPost**](LanPost.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

