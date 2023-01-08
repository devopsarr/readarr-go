# \ImportListExclusionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1ImportListExclusion**](ImportListExclusionApi.md#CreateApiV1ImportListExclusion) | **Post** /api/v1/importlistexclusion | 
[**DeleteApiV1ImportListExclusion**](ImportListExclusionApi.md#DeleteApiV1ImportListExclusion) | **Delete** /api/v1/importlistexclusion/{id} | 
[**GetApiV1ImportListExclusionById**](ImportListExclusionApi.md#GetApiV1ImportListExclusionById) | **Get** /api/v1/importlistexclusion/{id} | 
[**ListApiV1ImportListExclusion**](ImportListExclusionApi.md#ListApiV1ImportListExclusion) | **Get** /api/v1/importlistexclusion | 
[**UpdateApiV1ImportListExclusion**](ImportListExclusionApi.md#UpdateApiV1ImportListExclusion) | **Put** /api/v1/importlistexclusion/{id} | 



## CreateApiV1ImportListExclusion

> ImportListExclusionResource CreateApiV1ImportListExclusion(ctx).ImportListExclusionResource(importListExclusionResource).Execute()



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
    importListExclusionResource := *readarrClient.NewImportListExclusionResource() // ImportListExclusionResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListExclusionApi.CreateApiV1ImportListExclusion(context.Background()).ImportListExclusionResource(importListExclusionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.CreateApiV1ImportListExclusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1ImportListExclusion`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.CreateApiV1ImportListExclusion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ImportListExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListExclusionResource** | [**ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1ImportListExclusion

> DeleteApiV1ImportListExclusion(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.DeleteApiV1ImportListExclusion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.DeleteApiV1ImportListExclusion``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1ImportListExclusionRequest struct via the builder pattern


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


## GetApiV1ImportListExclusionById

> ImportListExclusionResource GetApiV1ImportListExclusionById(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.GetApiV1ImportListExclusionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.GetApiV1ImportListExclusionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ImportListExclusionById`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.GetApiV1ImportListExclusionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ImportListExclusionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1ImportListExclusion

> []ImportListExclusionResource ListApiV1ImportListExclusion(ctx).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.ListApiV1ImportListExclusion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.ListApiV1ImportListExclusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1ImportListExclusion`: []ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.ListApiV1ImportListExclusion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ImportListExclusionRequest struct via the builder pattern


### Return type

[**[]ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1ImportListExclusion

> ImportListExclusionResource UpdateApiV1ImportListExclusion(ctx, id).ImportListExclusionResource(importListExclusionResource).Execute()



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
    importListExclusionResource := *readarrClient.NewImportListExclusionResource() // ImportListExclusionResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListExclusionApi.UpdateApiV1ImportListExclusion(context.Background(), id).ImportListExclusionResource(importListExclusionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.UpdateApiV1ImportListExclusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ImportListExclusion`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.UpdateApiV1ImportListExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ImportListExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importListExclusionResource** | [**ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

