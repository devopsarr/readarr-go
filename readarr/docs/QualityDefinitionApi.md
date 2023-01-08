# \QualityDefinitionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1QualityDefinitionById**](QualityDefinitionApi.md#GetApiV1QualityDefinitionById) | **Get** /api/v1/qualitydefinition/{id} | 
[**ListApiV1QualityDefinition**](QualityDefinitionApi.md#ListApiV1QualityDefinition) | **Get** /api/v1/qualitydefinition | 
[**PutApiV1QualityDefinitionUpdate**](QualityDefinitionApi.md#PutApiV1QualityDefinitionUpdate) | **Put** /api/v1/qualitydefinition/update | 
[**UpdateApiV1QualityDefinition**](QualityDefinitionApi.md#UpdateApiV1QualityDefinition) | **Put** /api/v1/qualitydefinition/{id} | 



## GetApiV1QualityDefinitionById

> QualityDefinitionResource GetApiV1QualityDefinitionById(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.GetApiV1QualityDefinitionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.GetApiV1QualityDefinitionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QualityDefinitionById`: QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.GetApiV1QualityDefinitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QualityDefinitionByIdRequest struct via the builder pattern


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


## ListApiV1QualityDefinition

> []QualityDefinitionResource ListApiV1QualityDefinition(ctx).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.ListApiV1QualityDefinition(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.ListApiV1QualityDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1QualityDefinition`: []QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.ListApiV1QualityDefinition`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1QualityDefinitionRequest struct via the builder pattern


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


## PutApiV1QualityDefinitionUpdate

> PutApiV1QualityDefinitionUpdate(ctx).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.PutApiV1QualityDefinitionUpdate(context.Background()).QualityDefinitionResource(qualityDefinitionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.PutApiV1QualityDefinitionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1QualityDefinitionUpdateRequest struct via the builder pattern


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


## UpdateApiV1QualityDefinition

> QualityDefinitionResource UpdateApiV1QualityDefinition(ctx, id).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.UpdateApiV1QualityDefinition(context.Background(), id).QualityDefinitionResource(qualityDefinitionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.UpdateApiV1QualityDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1QualityDefinition`: QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.UpdateApiV1QualityDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1QualityDefinitionRequest struct via the builder pattern


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

