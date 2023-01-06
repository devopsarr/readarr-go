# \MissingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1WantedMissing**](MissingApi.md#GetApiV1WantedMissing) | **Get** /api/v1/wanted/missing | 
[**GetApiV1WantedMissingById**](MissingApi.md#GetApiV1WantedMissingById) | **Get** /api/v1/wanted/missing/{id} | 



## GetApiV1WantedMissing

> BookResourcePagingResource GetApiV1WantedMissing(ctx).IncludeAuthor(includeAuthor).Execute()



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
    resp, r, err := apiClient.MissingApi.GetApiV1WantedMissing(context.Background()).IncludeAuthor(includeAuthor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MissingApi.GetApiV1WantedMissing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1WantedMissing`: BookResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `MissingApi.GetApiV1WantedMissing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1WantedMissingRequest struct via the builder pattern


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


## GetApiV1WantedMissingById

> BookResource GetApiV1WantedMissingById(ctx, id).Execute()



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
    resp, r, err := apiClient.MissingApi.GetApiV1WantedMissingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MissingApi.GetApiV1WantedMissingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1WantedMissingById`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `MissingApi.GetApiV1WantedMissingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1WantedMissingByIdRequest struct via the builder pattern


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

