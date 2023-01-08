# \DelayProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1DelayProfile**](DelayProfileApi.md#CreateApiV1DelayProfile) | **Post** /api/v1/delayprofile | 
[**DeleteApiV1DelayProfile**](DelayProfileApi.md#DeleteApiV1DelayProfile) | **Delete** /api/v1/delayprofile/{id} | 
[**GetApiV1DelayProfileById**](DelayProfileApi.md#GetApiV1DelayProfileById) | **Get** /api/v1/delayprofile/{id} | 
[**ListApiV1DelayProfile**](DelayProfileApi.md#ListApiV1DelayProfile) | **Get** /api/v1/delayprofile | 
[**UpdateApiV1DelayProfile**](DelayProfileApi.md#UpdateApiV1DelayProfile) | **Put** /api/v1/delayprofile/{id} | 
[**UpdateApiV1DelayProfileReorder**](DelayProfileApi.md#UpdateApiV1DelayProfileReorder) | **Put** /api/v1/delayprofile/reorder/{id} | 



## CreateApiV1DelayProfile

> DelayProfileResource CreateApiV1DelayProfile(ctx).DelayProfileResource(delayProfileResource).Execute()



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
    delayProfileResource := *readarrClient.NewDelayProfileResource() // DelayProfileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.CreateApiV1DelayProfile(context.Background()).DelayProfileResource(delayProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.CreateApiV1DelayProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1DelayProfile`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.CreateApiV1DelayProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1DelayProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delayProfileResource** | [**DelayProfileResource**](DelayProfileResource.md) |  | 

### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1DelayProfile

> DeleteApiV1DelayProfile(ctx, id).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.DeleteApiV1DelayProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.DeleteApiV1DelayProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1DelayProfileRequest struct via the builder pattern


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


## GetApiV1DelayProfileById

> DelayProfileResource GetApiV1DelayProfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.GetApiV1DelayProfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.GetApiV1DelayProfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1DelayProfileById`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.GetApiV1DelayProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1DelayProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1DelayProfile

> []DelayProfileResource ListApiV1DelayProfile(ctx).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.ListApiV1DelayProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ListApiV1DelayProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1DelayProfile`: []DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.ListApiV1DelayProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1DelayProfileRequest struct via the builder pattern


### Return type

[**[]DelayProfileResource**](DelayProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1DelayProfile

> DelayProfileResource UpdateApiV1DelayProfile(ctx, id).DelayProfileResource(delayProfileResource).Execute()



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
    delayProfileResource := *readarrClient.NewDelayProfileResource() // DelayProfileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.UpdateApiV1DelayProfile(context.Background(), id).DelayProfileResource(delayProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.UpdateApiV1DelayProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1DelayProfile`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.UpdateApiV1DelayProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1DelayProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delayProfileResource** | [**DelayProfileResource**](DelayProfileResource.md) |  | 

### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1DelayProfileReorder

> UpdateApiV1DelayProfileReorder(ctx, id).AfterId(afterId).Execute()



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
    afterId := int32(56) // int32 |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.UpdateApiV1DelayProfileReorder(context.Background(), id).AfterId(afterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.UpdateApiV1DelayProfileReorder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateApiV1DelayProfileReorderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afterId** | **int32** |  | 

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

