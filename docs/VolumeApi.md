# {{classname}}

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersDatacenterIdVolumesGet**](VolumeApi.md#DatacentersDatacenterIdVolumesGet) | **Get** /datacenters/{datacenterId}/volumes | List Volumes 
[**DatacentersDatacenterIdVolumesPost**](VolumeApi.md#DatacentersDatacenterIdVolumesPost) | **Post** /datacenters/{datacenterId}/volumes | Create a Volume
[**DatacentersDatacenterIdVolumesVolumeIdCreateSnapshotPost**](VolumeApi.md#DatacentersDatacenterIdVolumesVolumeIdCreateSnapshotPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/create-snapshot | Create Volume Snapshot
[**DatacentersDatacenterIdVolumesVolumeIdDelete**](VolumeApi.md#DatacentersDatacenterIdVolumesVolumeIdDelete) | **Delete** /datacenters/{datacenterId}/volumes/{volumeId} | Delete a Volume
[**DatacentersDatacenterIdVolumesVolumeIdGet**](VolumeApi.md#DatacentersDatacenterIdVolumesVolumeIdGet) | **Get** /datacenters/{datacenterId}/volumes/{volumeId} | Retrieve a Volume
[**DatacentersDatacenterIdVolumesVolumeIdPatch**](VolumeApi.md#DatacentersDatacenterIdVolumesVolumeIdPatch) | **Patch** /datacenters/{datacenterId}/volumes/{volumeId} | Partially modify a Volume
[**DatacentersDatacenterIdVolumesVolumeIdPut**](VolumeApi.md#DatacentersDatacenterIdVolumesVolumeIdPut) | **Put** /datacenters/{datacenterId}/volumes/{volumeId} | Modify a Volume
[**DatacentersDatacenterIdVolumesVolumeIdRestoreSnapshotPost**](VolumeApi.md#DatacentersDatacenterIdVolumesVolumeIdRestoreSnapshotPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/restore-snapshot | Restore Volume Snapshot

# **DatacentersDatacenterIdVolumesGet**
> Volumes DatacentersDatacenterIdVolumesGet(ctx, datacenterId, optional)
List Volumes 

Retrieves a list of Volumes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***VolumeApiDatacentersDatacenterIdVolumesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiDatacentersDatacenterIdVolumesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volumes**](Volumes.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesPost**
> Volume DatacentersDatacenterIdVolumesPost(ctx, body, datacenterId, optional)
Create a Volume

Creates a volume within the datacenter. This will not attach the volume to a server. Please see the Servers section for details on how to attach storage volumes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Volume**](Volume.md)| Volume to be created | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***VolumeApiDatacentersDatacenterIdVolumesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiDatacentersDatacenterIdVolumesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesVolumeIdCreateSnapshotPost**
> interface{} DatacentersDatacenterIdVolumesVolumeIdCreateSnapshotPost(ctx, datacenterId, volumeId, optional)
Create Volume Snapshot

Creates a snapshot of a volume within the datacenter. You can use a snapshot to create a new storage volume or to restore a storage volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***VolumeApiDatacentersDatacenterIdVolumesVolumeIdCreateSnapshotPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiDatacentersDatacenterIdVolumesVolumeIdCreateSnapshotPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.**|  | 
 **description** | **optional.**|  | 
 **secAuthProtection** | **optional.**|  | 
 **licenceType** | **optional.**|  | 
 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesVolumeIdDelete**
> interface{} DatacentersDatacenterIdVolumesVolumeIdDelete(ctx, datacenterId, volumeId, optional)
Delete a Volume

Deletes the specified volume. This will result in the volume being removed from your datacenter. Use this with caution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***VolumeApiDatacentersDatacenterIdVolumesVolumeIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiDatacentersDatacenterIdVolumesVolumeIdDeleteOpts struct
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

# **DatacentersDatacenterIdVolumesVolumeIdGet**
> Volume DatacentersDatacenterIdVolumesVolumeIdGet(ctx, datacenterId, volumeId, optional)
Retrieve a Volume

Retrieves the attributes of a given Volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***VolumeApiDatacentersDatacenterIdVolumesVolumeIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiDatacentersDatacenterIdVolumesVolumeIdGetOpts struct
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

# **DatacentersDatacenterIdVolumesVolumeIdPatch**
> Volume DatacentersDatacenterIdVolumesVolumeIdPatch(ctx, body, datacenterId, volumeId, optional)
Partially modify a Volume

You can use update attributes of a Volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VolumeProperties**](VolumeProperties.md)| Modified properties of Volume | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***VolumeApiDatacentersDatacenterIdVolumesVolumeIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiDatacentersDatacenterIdVolumesVolumeIdPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesVolumeIdPut**
> Volume DatacentersDatacenterIdVolumesVolumeIdPut(ctx, body, datacenterId, volumeId, optional)
Modify a Volume

You can use update attributes of a Volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Volume**](Volume.md)| Modified Volume | 
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***VolumeApiDatacentersDatacenterIdVolumesVolumeIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiDatacentersDatacenterIdVolumesVolumeIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatacentersDatacenterIdVolumesVolumeIdRestoreSnapshotPost**
> interface{} DatacentersDatacenterIdVolumesVolumeIdRestoreSnapshotPost(ctx, datacenterId, volumeId, optional)
Restore Volume Snapshot

This will restore a snapshot onto a volume. A snapshot is created as just another image that can be used to create subsequent volumes if you want or to restore an existing volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterId** | **string**| The unique ID of the datacenter | 
  **volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***VolumeApiDatacentersDatacenterIdVolumesVolumeIdRestoreSnapshotPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiDatacentersDatacenterIdVolumesVolumeIdRestoreSnapshotPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.**|  | 
 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

