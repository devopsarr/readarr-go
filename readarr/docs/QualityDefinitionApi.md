# \QualityDefinitionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1QualitydefinitionById**](QualityDefinitionApi.md#GetApiV1QualitydefinitionById) | **Get** /api/v1/qualitydefinition/{id} | 
[**ListApiV1Qualitydefinition**](QualityDefinitionApi.md#ListApiV1Qualitydefinition) | **Get** /api/v1/qualitydefinition | 
[**PutApiV1QualitydefinitionUpdate**](QualityDefinitionApi.md#PutApiV1QualitydefinitionUpdate) | **Put** /api/v1/qualitydefinition/update | 
[**UpdateApiV1Qualitydefinition**](QualityDefinitionApi.md#UpdateApiV1Qualitydefinition) | **Put** /api/v1/qualitydefinition/{id} | 



## GetApiV1QualitydefinitionById

> QualityDefinitionResource GetApiV1QualitydefinitionById(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.GetApiV1QualitydefinitionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.GetApiV1QualitydefinitionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QualitydefinitionById`: QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.GetApiV1QualitydefinitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QualitydefinitionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1Qualitydefinition

> []QualityDefinitionResource ListApiV1Qualitydefinition(ctx).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.ListApiV1Qualitydefinition(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.ListApiV1Qualitydefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Qualitydefinition`: []QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.ListApiV1Qualitydefinition`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1QualitydefinitionRequest struct via the builder pattern


### Return type

[**[]QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1QualitydefinitionUpdate

> PutApiV1QualitydefinitionUpdate(ctx).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
    qualityDefinitionResource := []readarrClient.QualityDefinitionResource{*readarrClient.NewQualityDefinitionResource()} // []QualityDefinitionResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityDefinitionApi.PutApiV1QualitydefinitionUpdate(context.Background()).QualityDefinitionResource(qualityDefinitionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.PutApiV1QualitydefinitionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1QualitydefinitionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityDefinitionResource** | [**[]QualityDefinitionResource**](QualityDefinitionResource.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1Qualitydefinition

> QualityDefinitionResource UpdateApiV1Qualitydefinition(ctx, id).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
    qualityDefinitionResource := *readarrClient.NewQualityDefinitionResource() // QualityDefinitionResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityDefinitionApi.UpdateApiV1Qualitydefinition(context.Background(), id).QualityDefinitionResource(qualityDefinitionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.UpdateApiV1Qualitydefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Qualitydefinition`: QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.UpdateApiV1Qualitydefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1QualitydefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qualityDefinitionResource** | [**QualityDefinitionResource**](QualityDefinitionResource.md) |  | 

### Return type

[**QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

