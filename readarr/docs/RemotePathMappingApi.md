# \RemotePathMappingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1RemotePathMapping**](RemotePathMappingApi.md#CreateApiV1RemotePathMapping) | **Post** /api/v1/remotepathmapping | 
[**DeleteApiV1RemotePathMapping**](RemotePathMappingApi.md#DeleteApiV1RemotePathMapping) | **Delete** /api/v1/remotepathmapping/{id} | 
[**GetApiV1RemotePathMappingById**](RemotePathMappingApi.md#GetApiV1RemotePathMappingById) | **Get** /api/v1/remotepathmapping/{id} | 
[**ListApiV1RemotePathMapping**](RemotePathMappingApi.md#ListApiV1RemotePathMapping) | **Get** /api/v1/remotepathmapping | 
[**UpdateApiV1RemotePathMapping**](RemotePathMappingApi.md#UpdateApiV1RemotePathMapping) | **Put** /api/v1/remotepathmapping/{id} | 



## CreateApiV1RemotePathMapping

> RemotePathMappingResource CreateApiV1RemotePathMapping(ctx).RemotePathMappingResource(remotePathMappingResource).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    readarrClient "./openapi"
)

func main() {
    remotePathMappingResource := *readarrClient.NewRemotePathMappingResource() // RemotePathMappingResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.CreateApiV1RemotePathMapping(context.Background()).RemotePathMappingResource(remotePathMappingResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.CreateApiV1RemotePathMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1RemotePathMapping`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.CreateApiV1RemotePathMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1RemotePathMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remotePathMappingResource** | [**RemotePathMappingResource**](RemotePathMappingResource.md) |  | 

### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1RemotePathMapping

> DeleteApiV1RemotePathMapping(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    readarrClient "./openapi"
)

func main() {
    id := int32(56) // int32 | 

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.DeleteApiV1RemotePathMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.DeleteApiV1RemotePathMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1RemotePathMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1RemotePathMappingById

> RemotePathMappingResource GetApiV1RemotePathMappingById(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    readarrClient "./openapi"
)

func main() {
    id := int32(56) // int32 | 

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.GetApiV1RemotePathMappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.GetApiV1RemotePathMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1RemotePathMappingById`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.GetApiV1RemotePathMappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1RemotePathMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1RemotePathMapping

> []RemotePathMappingResource ListApiV1RemotePathMapping(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    readarrClient "./openapi"
)

func main() {

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.ListApiV1RemotePathMapping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.ListApiV1RemotePathMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1RemotePathMapping`: []RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.ListApiV1RemotePathMapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1RemotePathMappingRequest struct via the builder pattern


### Return type

[**[]RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1RemotePathMapping

> RemotePathMappingResource UpdateApiV1RemotePathMapping(ctx, id).RemotePathMappingResource(remotePathMappingResource).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    readarrClient "./openapi"
)

func main() {
    id := "id_example" // string | 
    remotePathMappingResource := *readarrClient.NewRemotePathMappingResource() // RemotePathMappingResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.UpdateApiV1RemotePathMapping(context.Background(), id).RemotePathMappingResource(remotePathMappingResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.UpdateApiV1RemotePathMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1RemotePathMapping`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.UpdateApiV1RemotePathMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1RemotePathMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remotePathMappingResource** | [**RemotePathMappingResource**](RemotePathMappingResource.md) |  | 

### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

