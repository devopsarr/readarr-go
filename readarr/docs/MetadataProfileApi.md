# \MetadataProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Metadataprofile**](MetadataProfileApi.md#CreateApiV1Metadataprofile) | **Post** /api/v1/metadataprofile | 
[**DeleteApiV1Metadataprofile**](MetadataProfileApi.md#DeleteApiV1Metadataprofile) | **Delete** /api/v1/metadataprofile/{id} | 
[**GetApiV1MetadataprofileById**](MetadataProfileApi.md#GetApiV1MetadataprofileById) | **Get** /api/v1/metadataprofile/{id} | 
[**ListApiV1Metadataprofile**](MetadataProfileApi.md#ListApiV1Metadataprofile) | **Get** /api/v1/metadataprofile | 
[**UpdateApiV1Metadataprofile**](MetadataProfileApi.md#UpdateApiV1Metadataprofile) | **Put** /api/v1/metadataprofile/{id} | 



## CreateApiV1Metadataprofile

> MetadataProfileResource CreateApiV1Metadataprofile(ctx).MetadataProfileResource(metadataProfileResource).Execute()



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
    resp, r, err := apiClient.MetadataProfileApi.CreateApiV1Metadataprofile(context.Background()).MetadataProfileResource(metadataProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.CreateApiV1Metadataprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Metadataprofile`: MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.CreateApiV1Metadataprofile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1MetadataprofileRequest struct via the builder pattern


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


## DeleteApiV1Metadataprofile

> DeleteApiV1Metadataprofile(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataProfileApi.DeleteApiV1Metadataprofile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.DeleteApiV1Metadataprofile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1MetadataprofileRequest struct via the builder pattern


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


## GetApiV1MetadataprofileById

> MetadataProfileResource GetApiV1MetadataprofileById(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataProfileApi.GetApiV1MetadataprofileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.GetApiV1MetadataprofileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1MetadataprofileById`: MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.GetApiV1MetadataprofileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MetadataprofileByIdRequest struct via the builder pattern


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


## ListApiV1Metadataprofile

> []MetadataProfileResource ListApiV1Metadataprofile(ctx).Execute()



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
    resp, r, err := apiClient.MetadataProfileApi.ListApiV1Metadataprofile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.ListApiV1Metadataprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Metadataprofile`: []MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.ListApiV1Metadataprofile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1MetadataprofileRequest struct via the builder pattern


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


## UpdateApiV1Metadataprofile

> MetadataProfileResource UpdateApiV1Metadataprofile(ctx, id).MetadataProfileResource(metadataProfileResource).Execute()



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
    resp, r, err := apiClient.MetadataProfileApi.UpdateApiV1Metadataprofile(context.Background(), id).MetadataProfileResource(metadataProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.UpdateApiV1Metadataprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Metadataprofile`: MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.UpdateApiV1Metadataprofile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1MetadataprofileRequest struct via the builder pattern


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

