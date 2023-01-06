# \MediaManagementConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1ConfigMediamanagement**](MediaManagementConfigApi.md#GetApiV1ConfigMediamanagement) | **Get** /api/v1/config/mediamanagement | 
[**GetApiV1ConfigMediamanagementById**](MediaManagementConfigApi.md#GetApiV1ConfigMediamanagementById) | **Get** /api/v1/config/mediamanagement/{id} | 
[**UpdateApiV1ConfigMediamanagement**](MediaManagementConfigApi.md#UpdateApiV1ConfigMediamanagement) | **Put** /api/v1/config/mediamanagement/{id} | 



## GetApiV1ConfigMediamanagement

> MediaManagementConfigResource GetApiV1ConfigMediamanagement(ctx).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigApi.GetApiV1ConfigMediamanagement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.GetApiV1ConfigMediamanagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigMediamanagement`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.GetApiV1ConfigMediamanagement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigMediamanagementRequest struct via the builder pattern


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1ConfigMediamanagementById

> MediaManagementConfigResource GetApiV1ConfigMediamanagementById(ctx, id).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigApi.GetApiV1ConfigMediamanagementById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.GetApiV1ConfigMediamanagementById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigMediamanagementById`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.GetApiV1ConfigMediamanagementById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigMediamanagementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1ConfigMediamanagement

> MediaManagementConfigResource UpdateApiV1ConfigMediamanagement(ctx, id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()



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
    mediaManagementConfigResource := *readarrClient.NewMediaManagementConfigResource() // MediaManagementConfigResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaManagementConfigApi.UpdateApiV1ConfigMediamanagement(context.Background(), id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.UpdateApiV1ConfigMediamanagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ConfigMediamanagement`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.UpdateApiV1ConfigMediamanagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ConfigMediamanagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaManagementConfigResource** | [**MediaManagementConfigResource**](MediaManagementConfigResource.md) |  | 

### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

