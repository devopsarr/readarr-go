# \QualityDefinitionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQualityDefinitionById**](QualityDefinitionApi.md#GetQualityDefinitionById) | **Get** /api/v1/qualitydefinition/{id} | 
[**ListQualityDefinition**](QualityDefinitionApi.md#ListQualityDefinition) | **Get** /api/v1/qualitydefinition | 
[**PutQualityDefinitionUpdate**](QualityDefinitionApi.md#PutQualityDefinitionUpdate) | **Put** /api/v1/qualitydefinition/update | 
[**UpdateQualityDefinition**](QualityDefinitionApi.md#UpdateQualityDefinition) | **Put** /api/v1/qualitydefinition/{id} | 



## GetQualityDefinitionById

> QualityDefinitionResource GetQualityDefinitionById(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.GetQualityDefinitionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.GetQualityDefinitionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQualityDefinitionById`: QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.GetQualityDefinitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQualityDefinitionByIdRequest struct via the builder pattern


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


## ListQualityDefinition

> []QualityDefinitionResource ListQualityDefinition(ctx).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.ListQualityDefinition(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.ListQualityDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQualityDefinition`: []QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.ListQualityDefinition`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListQualityDefinitionRequest struct via the builder pattern


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


## PutQualityDefinitionUpdate

> PutQualityDefinitionUpdate(ctx).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.PutQualityDefinitionUpdate(context.Background()).QualityDefinitionResource(qualityDefinitionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.PutQualityDefinitionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutQualityDefinitionUpdateRequest struct via the builder pattern


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


## UpdateQualityDefinition

> QualityDefinitionResource UpdateQualityDefinition(ctx, id).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.UpdateQualityDefinition(context.Background(), id).QualityDefinitionResource(qualityDefinitionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.UpdateQualityDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQualityDefinition`: QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.UpdateQualityDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQualityDefinitionRequest struct via the builder pattern


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

