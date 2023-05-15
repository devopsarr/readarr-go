# \MetadataProviderConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetadataProviderConfig**](MetadataProviderConfigAPI.md#GetMetadataProviderConfig) | **Get** /api/v1/config/metadataprovider | 
[**GetMetadataProviderConfigById**](MetadataProviderConfigAPI.md#GetMetadataProviderConfigById) | **Get** /api/v1/config/metadataprovider/{id} | 
[**UpdateMetadataProviderConfig**](MetadataProviderConfigAPI.md#UpdateMetadataProviderConfig) | **Put** /api/v1/config/metadataprovider/{id} | 



## GetMetadataProviderConfig

> MetadataProviderConfigResource GetMetadataProviderConfig(ctx).Execute()



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
    resp, r, err := apiClient.MetadataProviderConfigAPI.GetMetadataProviderConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigAPI.GetMetadataProviderConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataProviderConfig`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigAPI.GetMetadataProviderConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataProviderConfigRequest struct via the builder pattern


### Return type

[**MetadataProviderConfigResource**](MetadataProviderConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataProviderConfigById

> MetadataProviderConfigResource GetMetadataProviderConfigById(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataProviderConfigAPI.GetMetadataProviderConfigById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigAPI.GetMetadataProviderConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataProviderConfigById`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigAPI.GetMetadataProviderConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataProviderConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataProviderConfigResource**](MetadataProviderConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetadataProviderConfig

> MetadataProviderConfigResource UpdateMetadataProviderConfig(ctx, id).MetadataProviderConfigResource(metadataProviderConfigResource).Execute()



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
    metadataProviderConfigResource := *readarrClient.NewMetadataProviderConfigResource() // MetadataProviderConfigResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProviderConfigAPI.UpdateMetadataProviderConfig(context.Background(), id).MetadataProviderConfigResource(metadataProviderConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigAPI.UpdateMetadataProviderConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetadataProviderConfig`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigAPI.UpdateMetadataProviderConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetadataProviderConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataProviderConfigResource** | [**MetadataProviderConfigResource**](MetadataProviderConfigResource.md) |  | 

### Return type

[**MetadataProviderConfigResource**](MetadataProviderConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

