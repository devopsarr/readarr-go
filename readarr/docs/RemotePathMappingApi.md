# \RemotePathMappingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Remotepathmapping**](RemotePathMappingApi.md#CreateApiV1Remotepathmapping) | **Post** /api/v1/remotepathmapping | 
[**DeleteApiV1Remotepathmapping**](RemotePathMappingApi.md#DeleteApiV1Remotepathmapping) | **Delete** /api/v1/remotepathmapping/{id} | 
[**GetApiV1RemotepathmappingById**](RemotePathMappingApi.md#GetApiV1RemotepathmappingById) | **Get** /api/v1/remotepathmapping/{id} | 
[**ListApiV1Remotepathmapping**](RemotePathMappingApi.md#ListApiV1Remotepathmapping) | **Get** /api/v1/remotepathmapping | 
[**UpdateApiV1Remotepathmapping**](RemotePathMappingApi.md#UpdateApiV1Remotepathmapping) | **Put** /api/v1/remotepathmapping/{id} | 



## CreateApiV1Remotepathmapping

> RemotePathMappingResource CreateApiV1Remotepathmapping(ctx).RemotePathMappingResource(remotePathMappingResource).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.CreateApiV1Remotepathmapping(context.Background()).RemotePathMappingResource(remotePathMappingResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.CreateApiV1Remotepathmapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Remotepathmapping`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.CreateApiV1Remotepathmapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1RemotepathmappingRequest struct via the builder pattern


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


## DeleteApiV1Remotepathmapping

> DeleteApiV1Remotepathmapping(ctx, id).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.DeleteApiV1Remotepathmapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.DeleteApiV1Remotepathmapping``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1RemotepathmappingRequest struct via the builder pattern


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


## GetApiV1RemotepathmappingById

> RemotePathMappingResource GetApiV1RemotepathmappingById(ctx, id).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.GetApiV1RemotepathmappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.GetApiV1RemotepathmappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1RemotepathmappingById`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.GetApiV1RemotepathmappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1RemotepathmappingByIdRequest struct via the builder pattern


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


## ListApiV1Remotepathmapping

> []RemotePathMappingResource ListApiV1Remotepathmapping(ctx).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.ListApiV1Remotepathmapping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.ListApiV1Remotepathmapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Remotepathmapping`: []RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.ListApiV1Remotepathmapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1RemotepathmappingRequest struct via the builder pattern


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


## UpdateApiV1Remotepathmapping

> RemotePathMappingResource UpdateApiV1Remotepathmapping(ctx, id).RemotePathMappingResource(remotePathMappingResource).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.UpdateApiV1Remotepathmapping(context.Background(), id).RemotePathMappingResource(remotePathMappingResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.UpdateApiV1Remotepathmapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Remotepathmapping`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.UpdateApiV1Remotepathmapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1RemotepathmappingRequest struct via the builder pattern


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

