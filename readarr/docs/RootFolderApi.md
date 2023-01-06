# \RootFolderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Rootfolder**](RootFolderApi.md#CreateApiV1Rootfolder) | **Post** /api/v1/rootfolder | 
[**DeleteApiV1Rootfolder**](RootFolderApi.md#DeleteApiV1Rootfolder) | **Delete** /api/v1/rootfolder/{id} | 
[**GetApiV1RootfolderById**](RootFolderApi.md#GetApiV1RootfolderById) | **Get** /api/v1/rootfolder/{id} | 
[**ListApiV1Rootfolder**](RootFolderApi.md#ListApiV1Rootfolder) | **Get** /api/v1/rootfolder | 
[**UpdateApiV1Rootfolder**](RootFolderApi.md#UpdateApiV1Rootfolder) | **Put** /api/v1/rootfolder/{id} | 



## CreateApiV1Rootfolder

> RootFolderResource CreateApiV1Rootfolder(ctx).RootFolderResource(rootFolderResource).Execute()



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
    resp, r, err := apiClient.RootFolderApi.CreateApiV1Rootfolder(context.Background()).RootFolderResource(rootFolderResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.CreateApiV1Rootfolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Rootfolder`: RootFolderResource
    fmt.Fprintf(os.Stdout, "Response from `RootFolderApi.CreateApiV1Rootfolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1RootfolderRequest struct via the builder pattern


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


## DeleteApiV1Rootfolder

> DeleteApiV1Rootfolder(ctx, id).Execute()



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
    resp, r, err := apiClient.RootFolderApi.DeleteApiV1Rootfolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.DeleteApiV1Rootfolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1RootfolderRequest struct via the builder pattern


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


## GetApiV1RootfolderById

> RootFolderResource GetApiV1RootfolderById(ctx, id).Execute()



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
    resp, r, err := apiClient.RootFolderApi.GetApiV1RootfolderById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.GetApiV1RootfolderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1RootfolderById`: RootFolderResource
    fmt.Fprintf(os.Stdout, "Response from `RootFolderApi.GetApiV1RootfolderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1RootfolderByIdRequest struct via the builder pattern


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


## ListApiV1Rootfolder

> []RootFolderResource ListApiV1Rootfolder(ctx).Execute()



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
    resp, r, err := apiClient.RootFolderApi.ListApiV1Rootfolder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.ListApiV1Rootfolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Rootfolder`: []RootFolderResource
    fmt.Fprintf(os.Stdout, "Response from `RootFolderApi.ListApiV1Rootfolder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1RootfolderRequest struct via the builder pattern


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


## UpdateApiV1Rootfolder

> RootFolderResource UpdateApiV1Rootfolder(ctx, id).RootFolderResource(rootFolderResource).Execute()



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
    resp, r, err := apiClient.RootFolderApi.UpdateApiV1Rootfolder(context.Background(), id).RootFolderResource(rootFolderResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootFolderApi.UpdateApiV1Rootfolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Rootfolder`: RootFolderResource
    fmt.Fprintf(os.Stdout, "Response from `RootFolderApi.UpdateApiV1Rootfolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1RootfolderRequest struct via the builder pattern


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

