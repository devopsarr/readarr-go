# \MetadataProviderConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1ConfigMetadataprovider**](MetadataProviderConfigApi.md#GetApiV1ConfigMetadataprovider) | **Get** /api/v1/config/metadataprovider | 
[**GetApiV1ConfigMetadataproviderById**](MetadataProviderConfigApi.md#GetApiV1ConfigMetadataproviderById) | **Get** /api/v1/config/metadataprovider/{id} | 
[**UpdateApiV1ConfigMetadataprovider**](MetadataProviderConfigApi.md#UpdateApiV1ConfigMetadataprovider) | **Put** /api/v1/config/metadataprovider/{id} | 



## GetApiV1ConfigMetadataprovider

> MetadataProviderConfigResource GetApiV1ConfigMetadataprovider(ctx).Execute()



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
    resp, r, err := apiClient.MetadataProviderConfigApi.GetApiV1ConfigMetadataprovider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigApi.GetApiV1ConfigMetadataprovider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigMetadataprovider`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigApi.GetApiV1ConfigMetadataprovider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigMetadataproviderRequest struct via the builder pattern


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


## GetApiV1ConfigMetadataproviderById

> MetadataProviderConfigResource GetApiV1ConfigMetadataproviderById(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataProviderConfigApi.GetApiV1ConfigMetadataproviderById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigApi.GetApiV1ConfigMetadataproviderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigMetadataproviderById`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigApi.GetApiV1ConfigMetadataproviderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigMetadataproviderByIdRequest struct via the builder pattern


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


## UpdateApiV1ConfigMetadataprovider

> MetadataProviderConfigResource UpdateApiV1ConfigMetadataprovider(ctx, id).MetadataProviderConfigResource(metadataProviderConfigResource).Execute()



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
    resp, r, err := apiClient.MetadataProviderConfigApi.UpdateApiV1ConfigMetadataprovider(context.Background(), id).MetadataProviderConfigResource(metadataProviderConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigApi.UpdateApiV1ConfigMetadataprovider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ConfigMetadataprovider`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigApi.UpdateApiV1ConfigMetadataprovider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ConfigMetadataproviderRequest struct via the builder pattern


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

