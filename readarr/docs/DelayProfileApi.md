# \DelayProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Delayprofile**](DelayProfileApi.md#CreateApiV1Delayprofile) | **Post** /api/v1/delayprofile | 
[**DeleteApiV1Delayprofile**](DelayProfileApi.md#DeleteApiV1Delayprofile) | **Delete** /api/v1/delayprofile/{id} | 
[**GetApiV1DelayprofileById**](DelayProfileApi.md#GetApiV1DelayprofileById) | **Get** /api/v1/delayprofile/{id} | 
[**ListApiV1Delayprofile**](DelayProfileApi.md#ListApiV1Delayprofile) | **Get** /api/v1/delayprofile | 
[**UpdateApiV1Delayprofile**](DelayProfileApi.md#UpdateApiV1Delayprofile) | **Put** /api/v1/delayprofile/{id} | 
[**UpdateApiV1DelayprofileReorder**](DelayProfileApi.md#UpdateApiV1DelayprofileReorder) | **Put** /api/v1/delayprofile/reorder/{id} | 



## CreateApiV1Delayprofile

> DelayProfileResource CreateApiV1Delayprofile(ctx).DelayProfileResource(delayProfileResource).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.CreateApiV1Delayprofile(context.Background()).DelayProfileResource(delayProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.CreateApiV1Delayprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Delayprofile`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.CreateApiV1Delayprofile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1DelayprofileRequest struct via the builder pattern


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


## DeleteApiV1Delayprofile

> DeleteApiV1Delayprofile(ctx, id).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.DeleteApiV1Delayprofile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.DeleteApiV1Delayprofile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1DelayprofileRequest struct via the builder pattern


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


## GetApiV1DelayprofileById

> DelayProfileResource GetApiV1DelayprofileById(ctx, id).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.GetApiV1DelayprofileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.GetApiV1DelayprofileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1DelayprofileById`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.GetApiV1DelayprofileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1DelayprofileByIdRequest struct via the builder pattern


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


## ListApiV1Delayprofile

> []DelayProfileResource ListApiV1Delayprofile(ctx).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.ListApiV1Delayprofile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ListApiV1Delayprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Delayprofile`: []DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.ListApiV1Delayprofile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1DelayprofileRequest struct via the builder pattern


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


## UpdateApiV1Delayprofile

> DelayProfileResource UpdateApiV1Delayprofile(ctx, id).DelayProfileResource(delayProfileResource).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.UpdateApiV1Delayprofile(context.Background(), id).DelayProfileResource(delayProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.UpdateApiV1Delayprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Delayprofile`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.UpdateApiV1Delayprofile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1DelayprofileRequest struct via the builder pattern


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


## UpdateApiV1DelayprofileReorder

> UpdateApiV1DelayprofileReorder(ctx, id).AfterId(afterId).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.UpdateApiV1DelayprofileReorder(context.Background(), id).AfterId(afterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.UpdateApiV1DelayprofileReorder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateApiV1DelayprofileReorderRequest struct via the builder pattern


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

