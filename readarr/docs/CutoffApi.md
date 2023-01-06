# \CutoffApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1WantedCutoff**](CutoffApi.md#GetApiV1WantedCutoff) | **Get** /api/v1/wanted/cutoff | 
[**GetApiV1WantedCutoffById**](CutoffApi.md#GetApiV1WantedCutoffById) | **Get** /api/v1/wanted/cutoff/{id} | 



## GetApiV1WantedCutoff

> BookResourcePagingResource GetApiV1WantedCutoff(ctx).IncludeAuthor(includeAuthor).Execute()



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
    includeAuthor := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CutoffApi.GetApiV1WantedCutoff(context.Background()).IncludeAuthor(includeAuthor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CutoffApi.GetApiV1WantedCutoff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1WantedCutoff`: BookResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `CutoffApi.GetApiV1WantedCutoff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1WantedCutoffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeAuthor** | **bool** |  | [default to false]

### Return type

[**BookResourcePagingResource**](BookResourcePagingResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1WantedCutoffById

> BookResource GetApiV1WantedCutoffById(ctx, id).Execute()



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
    resp, r, err := apiClient.CutoffApi.GetApiV1WantedCutoffById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CutoffApi.GetApiV1WantedCutoffById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1WantedCutoffById`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `CutoffApi.GetApiV1WantedCutoffById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1WantedCutoffByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookResource**](BookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

