# \QualityProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1QualityProfile**](QualityProfileApi.md#CreateApiV1QualityProfile) | **Post** /api/v1/qualityprofile | 
[**DeleteApiV1QualityProfile**](QualityProfileApi.md#DeleteApiV1QualityProfile) | **Delete** /api/v1/qualityprofile/{id} | 
[**GetApiV1QualityProfileById**](QualityProfileApi.md#GetApiV1QualityProfileById) | **Get** /api/v1/qualityprofile/{id} | 
[**ListApiV1QualityProfile**](QualityProfileApi.md#ListApiV1QualityProfile) | **Get** /api/v1/qualityprofile | 
[**UpdateApiV1QualityProfile**](QualityProfileApi.md#UpdateApiV1QualityProfile) | **Put** /api/v1/qualityprofile/{id} | 



## CreateApiV1QualityProfile

> QualityProfileResource CreateApiV1QualityProfile(ctx).QualityProfileResource(qualityProfileResource).Execute()



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
    qualityProfileResource := *readarrClient.NewQualityProfileResource() // QualityProfileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityProfileApi.CreateApiV1QualityProfile(context.Background()).QualityProfileResource(qualityProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.CreateApiV1QualityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1QualityProfile`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.CreateApiV1QualityProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1QualityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityProfileResource** | [**QualityProfileResource**](QualityProfileResource.md) |  | 

### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1QualityProfile

> DeleteApiV1QualityProfile(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.DeleteApiV1QualityProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.DeleteApiV1QualityProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1QualityProfileRequest struct via the builder pattern


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


## GetApiV1QualityProfileById

> QualityProfileResource GetApiV1QualityProfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.GetApiV1QualityProfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.GetApiV1QualityProfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QualityProfileById`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.GetApiV1QualityProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QualityProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1QualityProfile

> []QualityProfileResource ListApiV1QualityProfile(ctx).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.ListApiV1QualityProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.ListApiV1QualityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1QualityProfile`: []QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.ListApiV1QualityProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1QualityProfileRequest struct via the builder pattern


### Return type

[**[]QualityProfileResource**](QualityProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1QualityProfile

> QualityProfileResource UpdateApiV1QualityProfile(ctx, id).QualityProfileResource(qualityProfileResource).Execute()



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
    qualityProfileResource := *readarrClient.NewQualityProfileResource() // QualityProfileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityProfileApi.UpdateApiV1QualityProfile(context.Background(), id).QualityProfileResource(qualityProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.UpdateApiV1QualityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1QualityProfile`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.UpdateApiV1QualityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1QualityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qualityProfileResource** | [**QualityProfileResource**](QualityProfileResource.md) |  | 

### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

