# {{classname}}

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersDatacenterIdServersServerIdNicsGet**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics | List Nics 
[**DatacentersDatacenterIdServersServerIdNicsNicIdDelete**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Delete a Nic
[**DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdDelete**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Delete a Firewall Rule
[**DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdGet**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Retrieve a Firewall Rule
[**DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPatch**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Partially modify a Firewall Rule
[**DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPut**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Modify a Firewall Rule
[**DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesGet**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | List Firewall Rules 
[**DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesPost**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | Create a Firewall Rule
[**DatacentersDatacenterIdServersServerIdNicsNicIdGet**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Retrieve a Nic
[**DatacentersDatacenterIdServersServerIdNicsNicIdPatch**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Partially modify a Nic
[**DatacentersDatacenterIdServersServerIdNicsNicIdPut**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsNicIdPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Modify a Nic
[**DatacentersDatacenterIdServersServerIdNicsPost**](NicApi.md#DatacentersDatacenterIdServersServerIdNicsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics | Create a Nic

# **DatacentersDatacenterIdServersServerIdNicsGet**
> Nics DatacentersDatacenterIdServersServerIdNicsGet(ctx, datacenterId, serverId, optional)
List Nics 

Retrieves a list of NICs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nics**](Nics.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdNicsNicIdDelete**
> interface{} DatacentersDatacenterIdServersServerIdNicsNicIdDelete(ctx, datacenterId, serverId, nicId, optional)
Delete a Nic

Deletes the specified NIC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdDeleteOpts struct
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

# **DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdDelete**
> interface{} DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdDelete(ctx, datacenterId, serverId, nicId, firewallruleId, optional)
Delete a Firewall Rule

Removes the specific Firewall Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
  **firewallruleId** | **string**| The unique ID of the Firewall Rule | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdDeleteOpts struct
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

# **DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdGet**
> FirewallRule DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdGet(ctx, datacenterId, serverId, nicId, firewallruleId, optional)
Retrieve a Firewall Rule

Retrieves the attributes of a given Firewall Rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
  **firewallruleId** | **string**| The unique ID of the Firewall Rule | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPatch**
> FirewallRule DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPatch(ctx, body, datacenterId, serverId, nicId, firewallruleId, optional)
Partially modify a Firewall Rule

You can use update attributes of a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallruleProperties**](FirewallruleProperties.md)| Modified Firewall Rule | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
  **firewallruleId** | **string**| The unique ID of the Firewall Rule | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPut**
> FirewallRule DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPut(ctx, body, datacenterId, serverId, nicId, firewallruleId, optional)
Modify a Firewall Rule

You can use update attributes of a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallRule**](FirewallRule.md)| Modified Firewall Rule | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
  **firewallruleId** | **string**| The unique ID of the Firewall Rule | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesFirewallruleIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesGet**
> FirewallRules DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesGet(ctx, datacenterId, serverId, nicId, optional)
List Firewall Rules 

Retrieves a list of firewall rules associated with a particular NIC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRules**](FirewallRules.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesPost**
> FirewallRule DatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesPost(ctx, body, datacenterId, serverId, nicId, optional)
Create a Firewall Rule

This will add a Firewall Rule to the NIC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallRule**](FirewallRule.md)| Firewall Rule to be created | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the server | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdFirewallrulesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdNicsNicIdGet**
> Nic DatacentersDatacenterIdServersServerIdNicsNicIdGet(ctx, datacenterId, serverId, nicId, optional)
Retrieve a Nic

Retrieves the attributes of a given NIC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdGetOpts struct
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

# **DatacentersDatacenterIdServersServerIdNicsNicIdPatch**
> Nic DatacentersDatacenterIdServersServerIdNicsNicIdPatch(ctx, body, datacenterId, serverId, nicId, optional)
Partially modify a Nic

You can use update attributes of a Nic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NicProperties**](NicProperties.md)| Modified properties of Nic | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdPatchOpts struct
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

# **DatacentersDatacenterIdServersServerIdNicsNicIdPut**
> Nic DatacentersDatacenterIdServersServerIdNicsNicIdPut(ctx, body, datacenterId, serverId, nicId, optional)
Modify a Nic

You can use update attributes of a Nic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Nic**](Nic.md)| Modified Nic | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsNicIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsNicIdPutOpts struct
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

# **DatacentersDatacenterIdServersServerIdNicsPost**
> Nic DatacentersDatacenterIdServersServerIdNicsPost(ctx, body, datacenterId, serverId, optional)
Create a Nic

Adds a NIC to the target server. Combine count of Nics and volumes attached to the server should not exceed size 24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Nic**](Nic.md)| Nic to be created | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***NicApiDatacentersDatacenterIdServersServerIdNicsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NicApiDatacentersDatacenterIdServersServerIdNicsPostOpts struct
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

