# {{classname}}

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersDatacenterIdLabelsGet**](LabelApi.md#DatacentersDatacenterIdLabelsGet) | **Get** /datacenters/{datacenterId}/labels | List all Data Center Labels
[**DatacentersDatacenterIdLabelsKeyDelete**](LabelApi.md#DatacentersDatacenterIdLabelsKeyDelete) | **Delete** /datacenters/{datacenterId}/labels/{key} | Delete a Label from Data Center
[**DatacentersDatacenterIdLabelsKeyGet**](LabelApi.md#DatacentersDatacenterIdLabelsKeyGet) | **Get** /datacenters/{datacenterId}/labels/{key} | Retrieve a Label of Data Center
[**DatacentersDatacenterIdLabelsKeyPut**](LabelApi.md#DatacentersDatacenterIdLabelsKeyPut) | **Put** /datacenters/{datacenterId}/labels/{key} | Modify a Label of Data Center
[**DatacentersDatacenterIdLabelsPost**](LabelApi.md#DatacentersDatacenterIdLabelsPost) | **Post** /datacenters/{datacenterId}/labels | Add a Label to Data Center
[**DatacentersDatacenterIdServersServerIdLabelsGet**](LabelApi.md#DatacentersDatacenterIdServersServerIdLabelsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/labels | List all Server Labels
[**DatacentersDatacenterIdServersServerIdLabelsKeyDelete**](LabelApi.md#DatacentersDatacenterIdServersServerIdLabelsKeyDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Delete a Label from Server
[**DatacentersDatacenterIdServersServerIdLabelsKeyGet**](LabelApi.md#DatacentersDatacenterIdServersServerIdLabelsKeyGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Retrieve a Label of Server
[**DatacentersDatacenterIdServersServerIdLabelsKeyPut**](LabelApi.md#DatacentersDatacenterIdServersServerIdLabelsKeyPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Modify a Label of Server
[**DatacentersDatacenterIdServersServerIdLabelsPost**](LabelApi.md#DatacentersDatacenterIdServersServerIdLabelsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/labels | Add a Label to Server
[**DatacentersDatacenterIdVolumesVolumeIdLabelsGet**](LabelApi.md#DatacentersDatacenterIdVolumesVolumeIdLabelsGet) | **Get** /datacenters/{datacenterId}/volumes/{volumeId}/labels | List all Volume Labels
[**DatacentersDatacenterIdVolumesVolumeIdLabelsKeyDelete**](LabelApi.md#DatacentersDatacenterIdVolumesVolumeIdLabelsKeyDelete) | **Delete** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Delete a Label from Volume
[**DatacentersDatacenterIdVolumesVolumeIdLabelsKeyGet**](LabelApi.md#DatacentersDatacenterIdVolumesVolumeIdLabelsKeyGet) | **Get** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Retrieve a Label of Volume
[**DatacentersDatacenterIdVolumesVolumeIdLabelsKeyPut**](LabelApi.md#DatacentersDatacenterIdVolumesVolumeIdLabelsKeyPut) | **Put** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Modify a Label of Volume
[**DatacentersDatacenterIdVolumesVolumeIdLabelsPost**](LabelApi.md#DatacentersDatacenterIdVolumesVolumeIdLabelsPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/labels | Add a Label to Volume
[**IpblocksIpblockIdLabelsGet**](LabelApi.md#IpblocksIpblockIdLabelsGet) | **Get** /ipblocks/{ipblockId}/labels | List all Ip Block Labels
[**IpblocksIpblockIdLabelsKeyDelete**](LabelApi.md#IpblocksIpblockIdLabelsKeyDelete) | **Delete** /ipblocks/{ipblockId}/labels/{key} | Delete a Label from IP Block
[**IpblocksIpblockIdLabelsKeyGet**](LabelApi.md#IpblocksIpblockIdLabelsKeyGet) | **Get** /ipblocks/{ipblockId}/labels/{key} | Retrieve a Label of IP Block
[**IpblocksIpblockIdLabelsKeyPut**](LabelApi.md#IpblocksIpblockIdLabelsKeyPut) | **Put** /ipblocks/{ipblockId}/labels/{key} | Modify a Label of IP Block
[**IpblocksIpblockIdLabelsPost**](LabelApi.md#IpblocksIpblockIdLabelsPost) | **Post** /ipblocks/{ipblockId}/labels | Add a Label to IP Block
[**LabelsGet**](LabelApi.md#LabelsGet) | **Get** /labels | List Labels 
[**LabelsLabelurnGet**](LabelApi.md#LabelsLabelurnGet) | **Get** /labels/{labelurn} | Returns the label by its URN.
[**SnapshotsSnapshotIdLabelsGet**](LabelApi.md#SnapshotsSnapshotIdLabelsGet) | **Get** /snapshots/{snapshotId}/labels | List all Snapshot Labels
[**SnapshotsSnapshotIdLabelsKeyDelete**](LabelApi.md#SnapshotsSnapshotIdLabelsKeyDelete) | **Delete** /snapshots/{snapshotId}/labels/{key} | Delete a Label from Snapshot
[**SnapshotsSnapshotIdLabelsKeyGet**](LabelApi.md#SnapshotsSnapshotIdLabelsKeyGet) | **Get** /snapshots/{snapshotId}/labels/{key} | Retrieve a Label of Snapshot
[**SnapshotsSnapshotIdLabelsKeyPut**](LabelApi.md#SnapshotsSnapshotIdLabelsKeyPut) | **Put** /snapshots/{snapshotId}/labels/{key} | Modify a Label of Snapshot
[**SnapshotsSnapshotIdLabelsPost**](LabelApi.md#SnapshotsSnapshotIdLabelsPost) | **Post** /snapshots/{snapshotId}/labels | Add a Label to Snapshot

# **DatacentersDatacenterIdLabelsGet**
> LabelResources DatacentersDatacenterIdLabelsGet(ctx, datacenterId, optional)
List all Data Center Labels

You can retrieve a list of all labels associated with a data center

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Data Center | 
 **optional** | ***LabelApiDatacentersDatacenterIdLabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdLabelsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLabelsKeyDelete**
> interface{} DatacentersDatacenterIdLabelsKeyDelete(ctx, datacenterId, key, optional)
Delete a Label from Data Center

This will remove a label from the data center.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Data Center | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdLabelsKeyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdLabelsKeyDeleteOpts struct
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

# **DatacentersDatacenterIdLabelsKeyGet**
> LabelResource DatacentersDatacenterIdLabelsKeyGet(ctx, datacenterId, key, optional)
Retrieve a Label of Data Center

This will retrieve the properties of a associated label to a data center.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Data Center | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdLabelsKeyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdLabelsKeyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLabelsKeyPut**
> LabelResource DatacentersDatacenterIdLabelsKeyPut(ctx, body, datacenterId, key, optional)
Modify a Label of Data Center

This will modify the value of the label on a data center.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Modified Label | 
  **datacenterId** | **string**| The unique ID of the Data Center | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdLabelsKeyPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdLabelsKeyPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdLabelsPost**
> LabelResource DatacentersDatacenterIdLabelsPost(ctx, body, datacenterId, optional)
Add a Label to Data Center

This will add a label to the data center.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Label to be added | 
  **datacenterId** | **string**| The unique ID of the Data Center | 
 **optional** | ***LabelApiDatacentersDatacenterIdLabelsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdLabelsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdLabelsGet**
> LabelResources DatacentersDatacenterIdServersServerIdLabelsGet(ctx, datacenterId, serverId, optional)
List all Server Labels

You can retrieve a list of all labels associated with a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***LabelApiDatacentersDatacenterIdServersServerIdLabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdServersServerIdLabelsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdLabelsKeyDelete**
> interface{} DatacentersDatacenterIdServersServerIdLabelsKeyDelete(ctx, datacenterId, serverId, key, optional)
Delete a Label from Server

This will remove a label from the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdServersServerIdLabelsKeyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdServersServerIdLabelsKeyDeleteOpts struct
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

# **DatacentersDatacenterIdServersServerIdLabelsKeyGet**
> LabelResource DatacentersDatacenterIdServersServerIdLabelsKeyGet(ctx, datacenterId, serverId, key, optional)
Retrieve a Label of Server

This will retrieve the properties of a associated label to a server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdServersServerIdLabelsKeyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdServersServerIdLabelsKeyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdLabelsKeyPut**
> LabelResource DatacentersDatacenterIdServersServerIdLabelsKeyPut(ctx, body, datacenterId, serverId, key, optional)
Modify a Label of Server

This will modify the value of the label on a server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Modified Label | 
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdServersServerIdLabelsKeyPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdServersServerIdLabelsKeyPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdServersServerIdLabelsPost**
> LabelResource DatacentersDatacenterIdServersServerIdLabelsPost(ctx, body, datacenterId, serverId, optional)
Add a Label to Server

This will add a label to the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Label to be added | 
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **serverId** | **string**| The unique ID of the Server | 
 **optional** | ***LabelApiDatacentersDatacenterIdServersServerIdLabelsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdServersServerIdLabelsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesVolumeIdLabelsGet**
> LabelResources DatacentersDatacenterIdVolumesVolumeIdLabelsGet(ctx, datacenterId, volumeId, optional)
List all Volume Labels

You can retrieve a list of all labels associated with a volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesVolumeIdLabelsKeyDelete**
> interface{} DatacentersDatacenterIdVolumesVolumeIdLabelsKeyDelete(ctx, datacenterId, volumeId, key, optional)
Delete a Label from Volume

This will remove a label from the volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsKeyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsKeyDeleteOpts struct
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

# **DatacentersDatacenterIdVolumesVolumeIdLabelsKeyGet**
> LabelResource DatacentersDatacenterIdVolumesVolumeIdLabelsKeyGet(ctx, datacenterId, volumeId, key, optional)
Retrieve a Label of Volume

This will retrieve the properties of a associated label to a volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsKeyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsKeyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesVolumeIdLabelsKeyPut**
> LabelResource DatacentersDatacenterIdVolumesVolumeIdLabelsKeyPut(ctx, body, datacenterId, volumeId, key, optional)
Modify a Label of Volume

This will modify the value of the label on a volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Modified Label | 
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsKeyPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsKeyPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesVolumeIdLabelsPost**
> LabelResource DatacentersDatacenterIdVolumesVolumeIdLabelsPost(ctx, body, datacenterId, volumeId, optional)
Add a Label to Volume

This will add a label to the volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Label to be added | 
  **datacenterId** | **string**| The unique ID of the Datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiDatacentersDatacenterIdVolumesVolumeIdLabelsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksIpblockIdLabelsGet**
> LabelResources IpblocksIpblockIdLabelsGet(ctx, ipblockId, optional)
List all Ip Block Labels

You can retrieve a list of all labels associated with a IP Block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipblockId** | **string**| The unique ID of the Ip Block | 
 **optional** | ***LabelApiIpblocksIpblockIdLabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiIpblocksIpblockIdLabelsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksIpblockIdLabelsKeyDelete**
> interface{} IpblocksIpblockIdLabelsKeyDelete(ctx, ipblockId, key, optional)
Delete a Label from IP Block

This will remove a label from the Ip Block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipblockId** | **string**| The unique ID of the Ip Block | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiIpblocksIpblockIdLabelsKeyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiIpblocksIpblockIdLabelsKeyDeleteOpts struct
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

# **IpblocksIpblockIdLabelsKeyGet**
> LabelResource IpblocksIpblockIdLabelsKeyGet(ctx, ipblockId, key, optional)
Retrieve a Label of IP Block

This will retrieve the properties of a associated label to a Ip Block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipblockId** | **string**| The unique ID of the Ip Block | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiIpblocksIpblockIdLabelsKeyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiIpblocksIpblockIdLabelsKeyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksIpblockIdLabelsKeyPut**
> LabelResource IpblocksIpblockIdLabelsKeyPut(ctx, body, ipblockId, key, optional)
Modify a Label of IP Block

This will modify the value of the label on a Ip Block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Modified Label | 
  **ipblockId** | **string**| The unique ID of the Ip Block | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiIpblocksIpblockIdLabelsKeyPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiIpblocksIpblockIdLabelsKeyPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IpblocksIpblockIdLabelsPost**
> LabelResource IpblocksIpblockIdLabelsPost(ctx, body, ipblockId, optional)
Add a Label to IP Block

This will add a label to the Ip Block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Label to be added | 
  **ipblockId** | **string**| The unique ID of the Ip Block | 
 **optional** | ***LabelApiIpblocksIpblockIdLabelsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiIpblocksIpblockIdLabelsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsGet**
> Labels LabelsGet(ctx, optional)
List Labels 

You can retrieve a complete list of labels that you have access to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LabelApiLabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiLabelsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Labels**](Labels.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsLabelurnGet**
> Label LabelsLabelurnGet(ctx, labelurn, optional)
Returns the label by its URN.

You can retrieve the details of a specific label using its URN. A URN is for uniqueness of a Label and composed using urn:label:<resource_type>:<resource_uuid>:<key>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelurn** | **string**| The URN representing the unique ID of the label. A URN is for uniqueness of a Label and composed using urn:label:&lt;resource_type&gt;:&lt;resource_uuid&gt;:&lt;key&gt; | 
 **optional** | ***LabelApiLabelsLabelurnGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiLabelsLabelurnGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Label**](Label.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsSnapshotIdLabelsGet**
> LabelResources SnapshotsSnapshotIdLabelsGet(ctx, snapshotId, optional)
List all Snapshot Labels

You can retrieve a list of all labels associated with a snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| The unique ID of the Snapshot | 
 **optional** | ***LabelApiSnapshotsSnapshotIdLabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiSnapshotsSnapshotIdLabelsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsSnapshotIdLabelsKeyDelete**
> interface{} SnapshotsSnapshotIdLabelsKeyDelete(ctx, snapshotId, key, optional)
Delete a Label from Snapshot

This will remove a label from the snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| The unique ID of the Snapshot | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiSnapshotsSnapshotIdLabelsKeyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiSnapshotsSnapshotIdLabelsKeyDeleteOpts struct
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

# **SnapshotsSnapshotIdLabelsKeyGet**
> LabelResource SnapshotsSnapshotIdLabelsKeyGet(ctx, snapshotId, key, optional)
Retrieve a Label of Snapshot

This will retrieve the properties of a associated label to a snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| The unique ID of the Snapshot | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiSnapshotsSnapshotIdLabelsKeyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiSnapshotsSnapshotIdLabelsKeyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsSnapshotIdLabelsKeyPut**
> LabelResource SnapshotsSnapshotIdLabelsKeyPut(ctx, body, snapshotId, key, optional)
Modify a Label of Snapshot

This will modify the value of the label on a snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Modified Label | 
  **snapshotId** | **string**| The unique ID of the Snapshot | 
  **key** | **string**| The key of the Label | 
 **optional** | ***LabelApiSnapshotsSnapshotIdLabelsKeyPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiSnapshotsSnapshotIdLabelsKeyPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsSnapshotIdLabelsPost**
> LabelResource SnapshotsSnapshotIdLabelsPost(ctx, body, snapshotId, optional)
Add a Label to Snapshot

This will add a label to the snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelResource**](LabelResource.md)| Label to be added | 
  **snapshotId** | **string**| The unique ID of the Snapshot | 
 **optional** | ***LabelApiSnapshotsSnapshotIdLabelsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelApiSnapshotsSnapshotIdLabelsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

