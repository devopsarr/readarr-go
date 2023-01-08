# \ReleaseProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1ReleaseProfile**](ReleaseProfileApi.md#CreateApiV1ReleaseProfile) | **Post** /api/v1/releaseprofile | 
[**DeleteApiV1ReleaseProfile**](ReleaseProfileApi.md#DeleteApiV1ReleaseProfile) | **Delete** /api/v1/releaseprofile/{id} | 
[**GetApiV1ReleaseProfileById**](ReleaseProfileApi.md#GetApiV1ReleaseProfileById) | **Get** /api/v1/releaseprofile/{id} | 
[**ListApiV1ReleaseProfile**](ReleaseProfileApi.md#ListApiV1ReleaseProfile) | **Get** /api/v1/releaseprofile | 
[**UpdateApiV1ReleaseProfile**](ReleaseProfileApi.md#UpdateApiV1ReleaseProfile) | **Put** /api/v1/releaseprofile/{id} | 



## CreateApiV1ReleaseProfile

> ReleaseProfileResource CreateApiV1ReleaseProfile(ctx).ReleaseProfileResource(releaseProfileResource).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.CreateApiV1ReleaseProfile(context.Background()).ReleaseProfileResource(releaseProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.CreateApiV1ReleaseProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1ReleaseProfile`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.CreateApiV1ReleaseProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ReleaseProfileRequest struct via the builder pattern


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


## DeleteApiV1ReleaseProfile

> DeleteApiV1ReleaseProfile(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.DeleteApiV1ReleaseProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.DeleteApiV1ReleaseProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1ReleaseProfileRequest struct via the builder pattern


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


## GetApiV1ReleaseProfileById

> ReleaseProfileResource GetApiV1ReleaseProfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.GetApiV1ReleaseProfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.GetApiV1ReleaseProfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ReleaseProfileById`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.GetApiV1ReleaseProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ReleaseProfileByIdRequest struct via the builder pattern


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


## ListApiV1ReleaseProfile

> []ReleaseProfileResource ListApiV1ReleaseProfile(ctx).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.ListApiV1ReleaseProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.ListApiV1ReleaseProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1ReleaseProfile`: []ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.ListApiV1ReleaseProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ReleaseProfileRequest struct via the builder pattern


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


## UpdateApiV1ReleaseProfile

> ReleaseProfileResource UpdateApiV1ReleaseProfile(ctx, id).ReleaseProfileResource(releaseProfileResource).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.UpdateApiV1ReleaseProfile(context.Background(), id).ReleaseProfileResource(releaseProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.UpdateApiV1ReleaseProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ReleaseProfile`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.UpdateApiV1ReleaseProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ReleaseProfileRequest struct via the builder pattern


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

