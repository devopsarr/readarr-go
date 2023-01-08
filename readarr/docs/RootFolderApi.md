# \RootFolderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1RootFolder**](RootFolderApi.md#CreateApiV1RootFolder) | **Post** /api/v1/rootfolder | 
[**DeleteApiV1RootFolder**](RootFolderApi.md#DeleteApiV1RootFolder) | **Delete** /api/v1/rootfolder/{id} | 
[**GetApiV1RootFolderById**](RootFolderApi.md#GetApiV1RootFolderById) | **Get** /api/v1/rootfolder/{id} | 
[**ListApiV1RootFolder**](RootFolderApi.md#ListApiV1RootFolder) | **Get** /api/v1/rootfolder | 
[**UpdateApiV1RootFolder**](RootFolderApi.md#UpdateApiV1RootFolder) | **Put** /api/v1/rootfolder/{id} | 



## CreateApiV1RootFolder

> RootFolderResource CreateApiV1RootFolder(ctx).RootFolderResource(rootFolderResource).Execute()



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
    rootFolderResource := *readarrClient.NewRootFolderResource() // RootFolderResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RootFolderApi.CreateApiV1RootFolder(context.Background()).RootFolderResource(rootFolderResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.CreateApiV1RootFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1RootFolder`: RootFolderResource
    fmt.Fprintf(os.Stdout, "Response from `RootFolderApi.CreateApiV1RootFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1RootFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rootFolderResource** | [**RootFolderResource**](RootFolderResource.md) |  | 

### Return type

[**RootFolderResource**](RootFolderResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1RootFolder

> DeleteApiV1RootFolder(ctx, id).Execute()



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
    resp, r, err := apiClient.RootFolderApi.DeleteApiV1RootFolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.DeleteApiV1RootFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1RootFolderRequest struct via the builder pattern


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


## GetApiV1RootFolderById

> RootFolderResource GetApiV1RootFolderById(ctx, id).Execute()



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
    resp, r, err := apiClient.RootFolderApi.GetApiV1RootFolderById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.GetApiV1RootFolderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1RootFolderById`: RootFolderResource
    fmt.Fprintf(os.Stdout, "Response from `RootFolderApi.GetApiV1RootFolderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1RootFolderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RootFolderResource**](RootFolderResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1RootFolder

> []RootFolderResource ListApiV1RootFolder(ctx).Execute()



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
    resp, r, err := apiClient.RootFolderApi.ListApiV1RootFolder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.ListApiV1RootFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1RootFolder`: []RootFolderResource
    fmt.Fprintf(os.Stdout, "Response from `RootFolderApi.ListApiV1RootFolder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1RootFolderRequest struct via the builder pattern


### Return type

[**[]RootFolderResource**](RootFolderResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1RootFolder

> RootFolderResource UpdateApiV1RootFolder(ctx, id).RootFolderResource(rootFolderResource).Execute()



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
    rootFolderResource := *readarrClient.NewRootFolderResource() // RootFolderResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RootFolderApi.UpdateApiV1RootFolder(context.Background(), id).RootFolderResource(rootFolderResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.UpdateApiV1RootFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1RootFolder`: RootFolderResource
    fmt.Fprintf(os.Stdout, "Response from `RootFolderApi.UpdateApiV1RootFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1RootFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rootFolderResource** | [**RootFolderResource**](RootFolderResource.md) |  | 

### Return type

[**RootFolderResource**](RootFolderResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

