# {{classname}}

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**K8sGet**](KubernetesApi.md#K8sGet) | **Get** /k8s | List Kubernetes Clusters
[**K8sK8sClusterIdDelete**](KubernetesApi.md#K8sK8sClusterIdDelete) | **Delete** /k8s/{k8sClusterId} | Delete Kubernetes Cluster
[**K8sK8sClusterIdGet**](KubernetesApi.md#K8sK8sClusterIdGet) | **Get** /k8s/{k8sClusterId} | Retrieve Kubernetes Cluster
[**K8sK8sClusterIdKubeconfigGet**](KubernetesApi.md#K8sK8sClusterIdKubeconfigGet) | **Get** /k8s/{k8sClusterId}/kubeconfig | Retrieve Kubernetes Configuration File
[**K8sK8sClusterIdNodepoolsGet**](KubernetesApi.md#K8sK8sClusterIdNodepoolsGet) | **Get** /k8s/{k8sClusterId}/nodepools | List Kubernetes Node Pools
[**K8sK8sClusterIdNodepoolsNodepoolIdDelete**](KubernetesApi.md#K8sK8sClusterIdNodepoolsNodepoolIdDelete) | **Delete** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Delete Kubernetes Node Pool
[**K8sK8sClusterIdNodepoolsNodepoolIdGet**](KubernetesApi.md#K8sK8sClusterIdNodepoolsNodepoolIdGet) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Retrieve Kubernetes Node Pool
[**K8sK8sClusterIdNodepoolsNodepoolIdNodesGet**](KubernetesApi.md#K8sK8sClusterIdNodepoolsNodepoolIdNodesGet) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes | Retrieve Kubernetes nodes.
[**K8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdGet**](KubernetesApi.md#K8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdGet) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId} | Retrieve Kubernetes node
[**K8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdReplacePost**](KubernetesApi.md#K8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdReplacePost) | **Post** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId}/replace | Recreate the Kubernetes node
[**K8sK8sClusterIdNodepoolsNodepoolIdPut**](KubernetesApi.md#K8sK8sClusterIdNodepoolsNodepoolIdPut) | **Put** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Modify Kubernetes Node Pool
[**K8sK8sClusterIdNodepoolsPost**](KubernetesApi.md#K8sK8sClusterIdNodepoolsPost) | **Post** /k8s/{k8sClusterId}/nodepools | Create a Kubernetes Node Pool
[**K8sK8sClusterIdPut**](KubernetesApi.md#K8sK8sClusterIdPut) | **Put** /k8s/{k8sClusterId} | Modify Kubernetes Cluster
[**K8sPost**](KubernetesApi.md#K8sPost) | **Post** /k8s | Create Kubernetes Cluster
[**K8sVersionsClusterVersionCompatibilitiesGet**](KubernetesApi.md#K8sVersionsClusterVersionCompatibilitiesGet) | **Get** /k8s/versions/{clusterVersion}/compatibilities | Retrieves a list of available kubernetes versions for nodepools depending on the given kubernetes version running in the cluster.
[**K8sVersionsDefaultGet**](KubernetesApi.md#K8sVersionsDefaultGet) | **Get** /k8s/versions/default | Retrieve the current default kubernetes version for clusters and nodepools.
[**K8sVersionsGet**](KubernetesApi.md#K8sVersionsGet) | **Get** /k8s/versions | Retrieve available Kubernetes versions

# **K8sGet**
> KubernetesClusters K8sGet(ctx, optional)
List Kubernetes Clusters

You can retrieve a list of all kubernetes clusters associated with a contract

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KubernetesApiK8sGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesClusters**](KubernetesClusters.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdDelete**
> interface{} K8sK8sClusterIdDelete(ctx, k8sClusterId, optional)
Delete Kubernetes Cluster

This will remove a Kubernetes Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***KubernetesApiK8sK8sClusterIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdDeleteOpts struct
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

# **K8sK8sClusterIdGet**
> KubernetesCluster K8sK8sClusterIdGet(ctx, k8sClusterId, optional)
Retrieve Kubernetes Cluster

This will retrieve a single Kubernetes Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***KubernetesApiK8sK8sClusterIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdKubeconfigGet**
> KubernetesConfig K8sK8sClusterIdKubeconfigGet(ctx, k8sClusterId, optional)
Retrieve Kubernetes Configuration File

You can retrieve kubernetes configuration file for the kubernetes cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***KubernetesApiK8sK8sClusterIdKubeconfigGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdKubeconfigGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesConfig**](KubernetesConfig.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdNodepoolsGet**
> KubernetesNodePools K8sK8sClusterIdNodepoolsGet(ctx, k8sClusterId, optional)
List Kubernetes Node Pools

You can retrieve a list of all kubernetes node pools part of kubernetes cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***KubernetesApiK8sK8sClusterIdNodepoolsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdNodepoolsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodePools**](KubernetesNodePools.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdNodepoolsNodepoolIdDelete**
> interface{} K8sK8sClusterIdNodepoolsNodepoolIdDelete(ctx, k8sClusterId, nodepoolId, optional)
Delete Kubernetes Node Pool

This will remove a Kubernetes Node Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
  **nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
 **optional** | ***KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdDeleteOpts struct
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

# **K8sK8sClusterIdNodepoolsNodepoolIdGet**
> KubernetesNodePool K8sK8sClusterIdNodepoolsNodepoolIdGet(ctx, k8sClusterId, nodepoolId, optional)
Retrieve Kubernetes Node Pool

You can retrieve a single Kubernetes Node Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
  **nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
 **optional** | ***KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodePool**](KubernetesNodePool.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdNodepoolsNodepoolIdNodesGet**
> KubernetesNodes K8sK8sClusterIdNodepoolsNodepoolIdNodesGet(ctx, k8sClusterId, nodepoolId, optional)
Retrieve Kubernetes nodes.

You can retrieve all nodes of Kubernetes Node Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
  **nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
 **optional** | ***KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdNodesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdNodesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodes**](KubernetesNodes.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdGet**
> KubernetesNode K8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdGet(ctx, k8sClusterId, nodepoolId, nodeId, optional)
Retrieve Kubernetes node

You can retrieve a single Kubernetes Node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
  **nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
  **nodeId** | **string**| The unique ID of the Kubernetes Node. | 
 **optional** | ***KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNode**](KubernetesNode.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdReplacePost**
> interface{} K8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdReplacePost(ctx, k8sClusterId, nodepoolId, nodeId, optional)
Recreate the Kubernetes node

You can recreate a single Kubernetes Node.  Managed Kubernetes starts a process which based on the nodepool's template creates & configures a new node, waits for status \"ACTIVE\", and migrates all the pods from the faulty node, deleting it once empty. While this operation occurs, the nodepool will have an extra billable \"ACTIVE\" node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
  **nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
  **nodeId** | **string**| The unique ID of the Kubernetes Node. | 
 **optional** | ***KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdReplacePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdNodesNodeIdReplacePostOpts struct
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

# **K8sK8sClusterIdNodepoolsNodepoolIdPut**
> KubernetesNodePool K8sK8sClusterIdNodepoolsNodepoolIdPut(ctx, body, k8sClusterId, nodepoolId, optional)
Modify Kubernetes Node Pool

This will modify the Kubernetes Node Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**KubernetesNodePoolProperties**](KubernetesNodePoolProperties.md)| Details of the Kubernetes Node Pool | 
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
  **nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
 **optional** | ***KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdNodepoolsNodepoolIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodePool**](KubernetesNodePool.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdNodepoolsPost**
> KubernetesNodePool K8sK8sClusterIdNodepoolsPost(ctx, body, k8sClusterId, optional)
Create a Kubernetes Node Pool

This will create a new Kubernetes Node Pool inside a Kubernetes Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**KubernetesNodePoolProperties**](KubernetesNodePoolProperties.md)| Properties of Kubernetes Node Pool | 
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***KubernetesApiK8sK8sClusterIdNodepoolsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdNodepoolsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodePool**](KubernetesNodePool.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sK8sClusterIdPut**
> KubernetesCluster K8sK8sClusterIdPut(ctx, body, k8sClusterId, optional)
Modify Kubernetes Cluster

This will modify the Kubernetes Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**KubernetesCluster**](KubernetesCluster.md)| Details of of the Kubernetes Cluster | 
  **k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***KubernetesApiK8sK8sClusterIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sK8sClusterIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sPost**
> KubernetesCluster K8sPost(ctx, body, optional)
Create Kubernetes Cluster

This will create a new Kubernetes Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**KubernetesClusterProperties**](KubernetesClusterProperties.md)| Properties of the Kubernetes Cluster | 
 **optional** | ***KubernetesApiK8sPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiK8sPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.**| Controls whether response is pretty-printed (with indentation and new lines) | 
 **depth** | **optional.**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#x27;s children are included  - depth&#x3D;... and so on | 
 **xContractNumber** | **optional.**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sVersionsClusterVersionCompatibilitiesGet**
> []string K8sVersionsClusterVersionCompatibilitiesGet(ctx, clusterVersion)
Retrieves a list of available kubernetes versions for nodepools depending on the given kubernetes version running in the cluster.

You can retrieve a list of available kubernetes versions for nodepools depending on the given kubernetes version running in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterVersion** | **string**|  | 

### Return type

**[]string**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sVersionsDefaultGet**
> string K8sVersionsDefaultGet(ctx, )
Retrieve the current default kubernetes version for clusters and nodepools.

You can retrieve the current default kubernetes version for clusters and nodepools.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **K8sVersionsGet**
> []string K8sVersionsGet(ctx, )
Retrieve available Kubernetes versions

You can retrieve a list of available kubernetes versions

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

