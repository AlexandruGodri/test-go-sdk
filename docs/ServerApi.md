# {{classname}}

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersDatacenterIdServersGet**](ServerApi.md#DatacentersDatacenterIdServersGet) | **Get** /datacenters/{datacenterId}/servers | List Servers 
[**DatacentersDatacenterIdServersPost**](ServerApi.md#DatacentersDatacenterIdServersPost) | **Post** /datacenters/{datacenterId}/servers | Create a Server
[**DatacentersDatacenterIdServersServerIdCdromsCdromIdDelete**](ServerApi.md#DatacentersDatacenterIdServersServerIdCdromsCdromIdDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/cdroms/{cdromId} | Detach a CD-ROM
[**DatacentersDatacenterIdServersServerIdCdromsCdromIdGet**](ServerApi.md#DatacentersDatacenterIdServersServerIdCdromsCdromIdGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/cdroms/{cdromId} | Retrieve an attached CD-ROM
[**DatacentersDatacenterIdServersServerIdCdromsGet**](ServerApi.md#DatacentersDatacenterIdServersServerIdCdromsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/cdroms | List attached CD-ROMs 
[**DatacentersDatacenterIdServersServerIdCdromsPost**](ServerApi.md#DatacentersDatacenterIdServersServerIdCdromsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/cdroms | Attach a CD-ROM
[**DatacentersDatacenterIdServersServerIdDelete**](ServerApi.md#DatacentersDatacenterIdServersServerIdDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId} | Delete a Server
[**DatacentersDatacenterIdServersServerIdGet**](ServerApi.md#DatacentersDatacenterIdServersServerIdGet) | **Get** /datacenters/{datacenterId}/servers/{serverId} | Retrieve a Server
[**DatacentersDatacenterIdServersServerIdPatch**](ServerApi.md#DatacentersDatacenterIdServersServerIdPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId} | Partially modify a Server
[**DatacentersDatacenterIdServersServerIdPut**](ServerApi.md#DatacentersDatacenterIdServersServerIdPut) | **Put** /datacenters/{datacenterId}/servers/{serverId} | Modify a Server
[**DatacentersDatacenterIdServersServerIdRebootPost**](ServerApi.md#DatacentersDatacenterIdServersServerIdRebootPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/reboot | Reboot a Server
[**DatacentersDatacenterIdServersServerIdStartPost**](ServerApi.md#DatacentersDatacenterIdServersServerIdStartPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/start | Start a Server
[**DatacentersDatacenterIdServersServerIdStopPost**](ServerApi.md#DatacentersDatacenterIdServersServerIdStopPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/stop | Stop a Server
[**DatacentersDatacenterIdServersServerIdUpgradePost**](ServerApi.md#DatacentersDatacenterIdServersServerIdUpgradePost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/upgrade | Upgrade a Server
[**DatacentersDatacenterIdServersServerIdVolumesGet**](ServerApi.md#DatacentersDatacenterIdServersServerIdVolumesGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/volumes | List Attached Volumes
[**DatacentersDatacenterIdServersServerIdVolumesPost**](ServerApi.md#DatacentersDatacenterIdServersServerIdVolumesPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/volumes | Attach a volume
[**DatacentersDatacenterIdServersServerIdVolumesVolumeIdDelete**](ServerApi.md#DatacentersDatacenterIdServersServerIdVolumesVolumeIdDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/volumes/{volumeId} | Detach a volume
[**DatacentersDatacenterIdServersServerIdVolumesVolumeIdGet**](ServerApi.md#DatacentersDatacenterIdServersServerIdVolumesVolumeIdGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/volumes/{volumeId} | Retrieve an attached volume

# **DatacentersDatacenterIdServersGet**
> Servers DatacentersDatacenterIdServersGet(ctx, datacenterId, optional)
List Servers 

You can retrieve a list of servers within a datacenter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **upgradeNeeded** | **optional.Bool**| It can be used to filter which servers can be upgraded which can not be upgraded. | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Servers**](Servers.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersPost**
> Server DatacentersDatacenterIdServersPost(ctx, body, datacenterId, optional)
Create a Server

Creates a server within an existing datacenter. You can configure the boot volume and connect the server to an existing LAN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Server**](Server.md)| Server to be created | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Server**](Server.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdCdromsCdromIdDelete**
> interface{} DatacentersDatacenterIdServersServerIdCdromsCdromIdDelete(ctx, datacenterId, serverId, cdromId, optional)
Detach a CD-ROM

This will detach a CD-ROM from the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **cdromId** | **string**| The unique ID of the CD-ROM | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdCdromsCdromIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdCdromsCdromIdDeleteOpts struct
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

# **DatacentersDatacenterIdServersServerIdCdromsCdromIdGet**
> Image DatacentersDatacenterIdServersServerIdCdromsCdromIdGet(ctx, datacenterId, serverId, cdromId, optional)
Retrieve an attached CD-ROM

You can retrieve a specific CD-ROM attached to the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **cdromId** | **string**| The unique ID of the CD-ROM | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdCdromsCdromIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdCdromsCdromIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Image**](Image.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdCdromsGet**
> Cdroms DatacentersDatacenterIdServersServerIdCdromsGet(ctx, datacenterId, serverId, optional)
List attached CD-ROMs 

You can retrieve a list of CD-ROMs attached to the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdCdromsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdCdromsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Cdroms**](Cdroms.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdCdromsPost**
> Image DatacentersDatacenterIdServersServerIdCdromsPost(ctx, body, datacenterId, serverId, optional)
Attach a CD-ROM

You can attach a CD-ROM to an existing server. You can attach up to 2 CD-ROMs to one server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Image**](Image.md)| CD-ROM to be attached | 
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdCdromsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdCdromsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Image**](Image.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdDelete**
> interface{} DatacentersDatacenterIdServersServerIdDelete(ctx, datacenterId, serverId, optional)
Delete a Server

This will remove a server from your datacenter; however, it will not remove the storage volumes attached to the server. You will need to make a separate API call to perform that action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdDeleteOpts struct
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

# **DatacentersDatacenterIdServersServerIdGet**
> Server DatacentersDatacenterIdServersServerIdGet(ctx, datacenterId, serverId, optional)
Retrieve a Server

Returns information about a server such as its configuration, provisioning status, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Server**](Server.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdPatch**
> Server DatacentersDatacenterIdServersServerIdPatch(ctx, body, datacenterId, serverId, optional)
Partially modify a Server

You can use update attributes of a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerProperties**](ServerProperties.md)| Modified properties of Server | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Server**](Server.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdPut**
> Server DatacentersDatacenterIdServersServerIdPut(ctx, body, datacenterId, serverId, optional)
Modify a Server

Allows to modify the attributes of a Server. From v5 onwards 'allowReboot' attribute will no longer be available. For certain server property change it was earlier forced to be provided. Now this behaviour is implicit and backend will do this automatically e.g. in earlier versions, when CPU family changes, the 'allowReboot' property was required to be set to true which will no longer be the case and the server will be rebooted automatically

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Server**](Server.md)| Modified Server | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Server**](Server.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdRebootPost**
> interface{} DatacentersDatacenterIdServersServerIdRebootPost(ctx, datacenterId, serverId, optional)
Reboot a Server

This will force a hard reboot of the server. Do not use this method if you want to gracefully reboot the machine. This is the equivalent of powering off the machine and turning it back on.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdRebootPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdRebootPostOpts struct
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

# **DatacentersDatacenterIdServersServerIdStartPost**
> interface{} DatacentersDatacenterIdServersServerIdStartPost(ctx, datacenterId, serverId, optional)
Start a Server

This will start a server. If the server's public IP was deallocated then a new IP will be assigned

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdStartPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdStartPostOpts struct
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

# **DatacentersDatacenterIdServersServerIdStopPost**
> interface{} DatacentersDatacenterIdServersServerIdStopPost(ctx, datacenterId, serverId, optional)
Stop a Server

This will stop a server. The machine will be forcefully powered off, billing will cease, and the public IP, if one is allocated, will be deallocated. The operation is not supported for CoreVPS servers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdStopPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdStopPostOpts struct
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

# **DatacentersDatacenterIdServersServerIdUpgradePost**
> interface{} DatacentersDatacenterIdServersServerIdUpgradePost(ctx, datacenterId, serverId, optional)
Upgrade a Server

This will upgrade the version of the server, if needed. To verify if there is an upgrade available for a server, call '/datacenters/{datacenterId}/servers?upgradeNeeded=true'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdUpgradePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdUpgradePostOpts struct
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

# **DatacentersDatacenterIdServersServerIdVolumesGet**
> AttachedVolumes DatacentersDatacenterIdServersServerIdVolumesGet(ctx, datacenterId, serverId, optional)
List Attached Volumes

You can retrieve a list of volumes attached to the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdVolumesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdVolumesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**AttachedVolumes**](AttachedVolumes.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdVolumesPost**
> Volume DatacentersDatacenterIdServersServerIdVolumesPost(ctx, body, datacenterId, serverId, optional)
Attach a volume

This will attach a pre-existing storage volume to the server. It is also possible to create and attach a volume in one step just by providing a new volume description as payload. Combine count of Nics and volumes attached to the server should not exceed size 24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Volume**](Volume.md)| Volume to be attached (created and attached) | 
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdVolumesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdVolumesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdVolumesVolumeIdDelete**
> interface{} DatacentersDatacenterIdServersServerIdVolumesVolumeIdDelete(ctx, datacenterId, serverId, volumeId, optional)
Detach a volume

This will detach the volume from the server. This will not delete the volume from your datacenter. You will need to make a separate request to perform a deletion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdVolumesVolumeIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdVolumesVolumeIdDeleteOpts struct
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

# **DatacentersDatacenterIdServersServerIdVolumesVolumeIdGet**
> Volume DatacentersDatacenterIdServersServerIdVolumesVolumeIdGet(ctx, datacenterId, serverId, volumeId, optional)
Retrieve an attached volume

This will retrieve the properties of an attached volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***ServerApiDatacentersDatacenterIdServersServerIdVolumesVolumeIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDatacentersDatacenterIdServersServerIdVolumesVolumeIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

