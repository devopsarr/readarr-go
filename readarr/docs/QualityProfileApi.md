# \QualityProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Qualityprofile**](QualityProfileApi.md#CreateApiV1Qualityprofile) | **Post** /api/v1/qualityprofile | 
[**DeleteApiV1Qualityprofile**](QualityProfileApi.md#DeleteApiV1Qualityprofile) | **Delete** /api/v1/qualityprofile/{id} | 
[**GetApiV1QualityprofileById**](QualityProfileApi.md#GetApiV1QualityprofileById) | **Get** /api/v1/qualityprofile/{id} | 
[**ListApiV1Qualityprofile**](QualityProfileApi.md#ListApiV1Qualityprofile) | **Get** /api/v1/qualityprofile | 
[**UpdateApiV1Qualityprofile**](QualityProfileApi.md#UpdateApiV1Qualityprofile) | **Put** /api/v1/qualityprofile/{id} | 



## CreateApiV1Qualityprofile

> QualityProfileResource CreateApiV1Qualityprofile(ctx).QualityProfileResource(qualityProfileResource).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.CreateApiV1Qualityprofile(context.Background()).QualityProfileResource(qualityProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.CreateApiV1Qualityprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Qualityprofile`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.CreateApiV1Qualityprofile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1QualityprofileRequest struct via the builder pattern


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


## DeleteApiV1Qualityprofile

> DeleteApiV1Qualityprofile(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.DeleteApiV1Qualityprofile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.DeleteApiV1Qualityprofile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1QualityprofileRequest struct via the builder pattern


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


## GetApiV1QualityprofileById

> QualityProfileResource GetApiV1QualityprofileById(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.GetApiV1QualityprofileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.GetApiV1QualityprofileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QualityprofileById`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.GetApiV1QualityprofileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QualityprofileByIdRequest struct via the builder pattern


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


## ListApiV1Qualityprofile

> []QualityProfileResource ListApiV1Qualityprofile(ctx).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.ListApiV1Qualityprofile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.ListApiV1Qualityprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Qualityprofile`: []QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.ListApiV1Qualityprofile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1QualityprofileRequest struct via the builder pattern


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


## UpdateApiV1Qualityprofile

> QualityProfileResource UpdateApiV1Qualityprofile(ctx, id).QualityProfileResource(qualityProfileResource).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.UpdateApiV1Qualityprofile(context.Background(), id).QualityProfileResource(qualityProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.UpdateApiV1Qualityprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Qualityprofile`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.UpdateApiV1Qualityprofile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1QualityprofileRequest struct via the builder pattern


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

