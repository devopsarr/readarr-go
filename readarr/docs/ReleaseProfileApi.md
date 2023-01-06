# \ReleaseProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Releaseprofile**](ReleaseProfileApi.md#CreateApiV1Releaseprofile) | **Post** /api/v1/releaseprofile | 
[**DeleteApiV1Releaseprofile**](ReleaseProfileApi.md#DeleteApiV1Releaseprofile) | **Delete** /api/v1/releaseprofile/{id} | 
[**GetApiV1ReleaseprofileById**](ReleaseProfileApi.md#GetApiV1ReleaseprofileById) | **Get** /api/v1/releaseprofile/{id} | 
[**ListApiV1Releaseprofile**](ReleaseProfileApi.md#ListApiV1Releaseprofile) | **Get** /api/v1/releaseprofile | 
[**UpdateApiV1Releaseprofile**](ReleaseProfileApi.md#UpdateApiV1Releaseprofile) | **Put** /api/v1/releaseprofile/{id} | 



## CreateApiV1Releaseprofile

> ReleaseProfileResource CreateApiV1Releaseprofile(ctx).ReleaseProfileResource(releaseProfileResource).Execute()



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
    releaseProfileResource := *readarrClient.NewReleaseProfileResource() // ReleaseProfileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseProfileApi.CreateApiV1Releaseprofile(context.Background()).ReleaseProfileResource(releaseProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.CreateApiV1Releaseprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Releaseprofile`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.CreateApiV1Releaseprofile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ReleaseprofileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseProfileResource** | [**ReleaseProfileResource**](ReleaseProfileResource.md) |  | 

### Return type

[**ReleaseProfileResource**](ReleaseProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1Releaseprofile

> DeleteApiV1Releaseprofile(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.DeleteApiV1Releaseprofile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.DeleteApiV1Releaseprofile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1ReleaseprofileRequest struct via the builder pattern


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


## GetApiV1ReleaseprofileById

> ReleaseProfileResource GetApiV1ReleaseprofileById(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.GetApiV1ReleaseprofileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.GetApiV1ReleaseprofileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ReleaseprofileById`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.GetApiV1ReleaseprofileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ReleaseprofileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseProfileResource**](ReleaseProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1Releaseprofile

> []ReleaseProfileResource ListApiV1Releaseprofile(ctx).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.ListApiV1Releaseprofile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.ListApiV1Releaseprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Releaseprofile`: []ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.ListApiV1Releaseprofile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ReleaseprofileRequest struct via the builder pattern


### Return type

[**[]ReleaseProfileResource**](ReleaseProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1Releaseprofile

> ReleaseProfileResource UpdateApiV1Releaseprofile(ctx, id).ReleaseProfileResource(releaseProfileResource).Execute()



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
    releaseProfileResource := *readarrClient.NewReleaseProfileResource() // ReleaseProfileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseProfileApi.UpdateApiV1Releaseprofile(context.Background(), id).ReleaseProfileResource(releaseProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.UpdateApiV1Releaseprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Releaseprofile`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.UpdateApiV1Releaseprofile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ReleaseprofileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseProfileResource** | [**ReleaseProfileResource**](ReleaseProfileResource.md) |  | 

### Return type

[**ReleaseProfileResource**](ReleaseProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

