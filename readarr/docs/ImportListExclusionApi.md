# \ImportListExclusionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Importlistexclusion**](ImportListExclusionApi.md#CreateApiV1Importlistexclusion) | **Post** /api/v1/importlistexclusion | 
[**DeleteApiV1Importlistexclusion**](ImportListExclusionApi.md#DeleteApiV1Importlistexclusion) | **Delete** /api/v1/importlistexclusion/{id} | 
[**GetApiV1ImportlistexclusionById**](ImportListExclusionApi.md#GetApiV1ImportlistexclusionById) | **Get** /api/v1/importlistexclusion/{id} | 
[**ListApiV1Importlistexclusion**](ImportListExclusionApi.md#ListApiV1Importlistexclusion) | **Get** /api/v1/importlistexclusion | 
[**UpdateApiV1Importlistexclusion**](ImportListExclusionApi.md#UpdateApiV1Importlistexclusion) | **Put** /api/v1/importlistexclusion/{id} | 



## CreateApiV1Importlistexclusion

> ImportListExclusionResource CreateApiV1Importlistexclusion(ctx).ImportListExclusionResource(importListExclusionResource).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.CreateApiV1Importlistexclusion(context.Background()).ImportListExclusionResource(importListExclusionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.CreateApiV1Importlistexclusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Importlistexclusion`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.CreateApiV1Importlistexclusion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ImportlistexclusionRequest struct via the builder pattern


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


## DeleteApiV1Importlistexclusion

> DeleteApiV1Importlistexclusion(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.DeleteApiV1Importlistexclusion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.DeleteApiV1Importlistexclusion``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1ImportlistexclusionRequest struct via the builder pattern


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


## GetApiV1ImportlistexclusionById

> ImportListExclusionResource GetApiV1ImportlistexclusionById(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.GetApiV1ImportlistexclusionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.GetApiV1ImportlistexclusionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ImportlistexclusionById`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.GetApiV1ImportlistexclusionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ImportlistexclusionByIdRequest struct via the builder pattern


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


## ListApiV1Importlistexclusion

> []ImportListExclusionResource ListApiV1Importlistexclusion(ctx).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.ListApiV1Importlistexclusion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.ListApiV1Importlistexclusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Importlistexclusion`: []ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.ListApiV1Importlistexclusion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ImportlistexclusionRequest struct via the builder pattern


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


## UpdateApiV1Importlistexclusion

> ImportListExclusionResource UpdateApiV1Importlistexclusion(ctx, id).ImportListExclusionResource(importListExclusionResource).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.UpdateApiV1Importlistexclusion(context.Background(), id).ImportListExclusionResource(importListExclusionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.UpdateApiV1Importlistexclusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Importlistexclusion`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.UpdateApiV1Importlistexclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ImportlistexclusionRequest struct via the builder pattern


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

