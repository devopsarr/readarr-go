# \MetadataProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1MetadataProfile**](MetadataProfileApi.md#CreateApiV1MetadataProfile) | **Post** /api/v1/metadataprofile | 
[**DeleteApiV1MetadataProfile**](MetadataProfileApi.md#DeleteApiV1MetadataProfile) | **Delete** /api/v1/metadataprofile/{id} | 
[**GetApiV1MetadataProfileById**](MetadataProfileApi.md#GetApiV1MetadataProfileById) | **Get** /api/v1/metadataprofile/{id} | 
[**ListApiV1MetadataProfile**](MetadataProfileApi.md#ListApiV1MetadataProfile) | **Get** /api/v1/metadataprofile | 
[**UpdateApiV1MetadataProfile**](MetadataProfileApi.md#UpdateApiV1MetadataProfile) | **Put** /api/v1/metadataprofile/{id} | 



## CreateApiV1MetadataProfile

> MetadataProfileResource CreateApiV1MetadataProfile(ctx).MetadataProfileResource(metadataProfileResource).Execute()



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
    metadataProfileResource := *readarrClient.NewMetadataProfileResource() // MetadataProfileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.CreateApiV1MetadataProfile(context.Background()).MetadataProfileResource(metadataProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.CreateApiV1MetadataProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1MetadataProfile`: MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.CreateApiV1MetadataProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1MetadataProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadataProfileResource** | [**MetadataProfileResource**](MetadataProfileResource.md) |  | 

### Return type

[**MetadataProfileResource**](MetadataProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1MetadataProfile

> DeleteApiV1MetadataProfile(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataProfileApi.DeleteApiV1MetadataProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.DeleteApiV1MetadataProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1MetadataProfileRequest struct via the builder pattern


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


## GetApiV1MetadataProfileById

> MetadataProfileResource GetApiV1MetadataProfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataProfileApi.GetApiV1MetadataProfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.GetApiV1MetadataProfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1MetadataProfileById`: MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.GetApiV1MetadataProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MetadataProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataProfileResource**](MetadataProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1MetadataProfile

> []MetadataProfileResource ListApiV1MetadataProfile(ctx).Execute()



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
    resp, r, err := apiClient.MetadataProfileApi.ListApiV1MetadataProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.ListApiV1MetadataProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1MetadataProfile`: []MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.ListApiV1MetadataProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1MetadataProfileRequest struct via the builder pattern


### Return type

[**[]MetadataProfileResource**](MetadataProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1MetadataProfile

> MetadataProfileResource UpdateApiV1MetadataProfile(ctx, id).MetadataProfileResource(metadataProfileResource).Execute()



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
    metadataProfileResource := *readarrClient.NewMetadataProfileResource() // MetadataProfileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.UpdateApiV1MetadataProfile(context.Background(), id).MetadataProfileResource(metadataProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.UpdateApiV1MetadataProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1MetadataProfile`: MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.UpdateApiV1MetadataProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1MetadataProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataProfileResource** | [**MetadataProfileResource**](MetadataProfileResource.md) |  | 

### Return type

[**MetadataProfileResource**](MetadataProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

