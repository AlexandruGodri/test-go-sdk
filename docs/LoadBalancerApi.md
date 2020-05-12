# {{classname}}

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersDatacenterIdLoadbalancersGet**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersGet) | **Get** /datacenters/{datacenterId}/loadbalancers | List Load Balancers
[**DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsGet**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsGet) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics | List Load Balancer Members 
[**DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdDelete**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdDelete) | **Delete** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics/{nicId} | Detach a nic from loadbalancer
[**DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdGet**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdGet) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics/{nicId} | Retrieve a nic attached to Load Balancer
[**DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsPost**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsPost) | **Post** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics | Attach a nic to Load Balancer
[**DatacentersDatacenterIdLoadbalancersLoadbalancerIdDelete**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersLoadbalancerIdDelete) | **Delete** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Delete a Loadbalancer.
[**DatacentersDatacenterIdLoadbalancersLoadbalancerIdGet**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersLoadbalancerIdGet) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Retrieve a loadbalancer
[**DatacentersDatacenterIdLoadbalancersLoadbalancerIdPatch**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersLoadbalancerIdPatch) | **Patch** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Partially modify a Loadbalancer
[**DatacentersDatacenterIdLoadbalancersLoadbalancerIdPut**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersLoadbalancerIdPut) | **Put** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Modify a Load Balancer
[**DatacentersDatacenterIdLoadbalancersPost**](LoadBalancerApi.md#DatacentersDatacenterIdLoadbalancersPost) | **Post** /datacenters/{datacenterId}/loadbalancers | Create a Load Balancer

# **DatacentersDatacenterIdLoadbalancersGet**
> Loadbalancers DatacentersDatacenterIdLoadbalancersGet(ctx, datacenterId, optional)
List Load Balancers

Retrieve a list of Load Balancers within the datacenter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancers**](Loadbalancers.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsGet**
> BalancedNics DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsGet(ctx, datacenterId, loadbalancerId, optional)
List Load Balancer Members 

You can retrieve a list of nics attached to a Load Balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**BalancedNics**](BalancedNics.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdDelete**
> interface{} DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdDelete(ctx, datacenterId, loadbalancerId, nicId, optional)
Detach a nic from loadbalancer

This will remove a nic from Load Balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **loadbalancerId** | **string**| The unique ID of the Load Balancer | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdDeleteOpts struct
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

# **DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdGet**
> Nic DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdGet(ctx, datacenterId, loadbalancerId, nicId, optional)
Retrieve a nic attached to Load Balancer

This will retrieve the properties of an attached nic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **loadbalancerId** | **string**| The unique ID of the Load Balancer | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsNicIdGetOpts struct
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

# **DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsPost**
> Nic DatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsPost(ctx, body, datacenterId, loadbalancerId, optional)
Attach a nic to Load Balancer

This will attach a pre-existing nic to a Load Balancer. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Nic**](Nic.md)| Nic id to be attached | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdBalancednicsPostOpts struct
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

# **DatacentersDatacenterIdLoadbalancersLoadbalancerIdDelete**
> interface{} DatacentersDatacenterIdLoadbalancersLoadbalancerIdDelete(ctx, datacenterId, loadbalancerId, optional)
Delete a Loadbalancer.

Removes the specific Loadbalancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdDeleteOpts struct
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

# **DatacentersDatacenterIdLoadbalancersLoadbalancerIdGet**
> Loadbalancer DatacentersDatacenterIdLoadbalancersLoadbalancerIdGet(ctx, datacenterId, loadbalancerId, optional)
Retrieve a loadbalancer

Retrieves the attributes of a given Loadbalancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancer**](Loadbalancer.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLoadbalancersLoadbalancerIdPatch**
> Loadbalancer DatacentersDatacenterIdLoadbalancersLoadbalancerIdPatch(ctx, body, datacenterId, loadbalancerId, optional)
Partially modify a Loadbalancer

You can use update attributes of a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoadbalancerProperties**](LoadbalancerProperties.md)| Modified Loadbalancer | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancer**](Loadbalancer.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLoadbalancersLoadbalancerIdPut**
> Loadbalancer DatacentersDatacenterIdLoadbalancersLoadbalancerIdPut(ctx, body, datacenterId, loadbalancerId, optional)
Modify a Load Balancer

You can use update attributes of a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Loadbalancer**](Loadbalancer.md)| Modified Loadbalancer | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersLoadbalancerIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancer**](Loadbalancer.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLoadbalancersPost**
> Loadbalancer DatacentersDatacenterIdLoadbalancersPost(ctx, body, datacenterId, optional)
Create a Load Balancer

Creates a Loadbalancer within the datacenter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Loadbalancer**](Loadbalancer.md)| Loadbalancer to be created | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***LoadBalancerApiDatacentersDatacenterIdLoadbalancersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerApiDatacentersDatacenterIdLoadbalancersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancer**](Loadbalancer.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

