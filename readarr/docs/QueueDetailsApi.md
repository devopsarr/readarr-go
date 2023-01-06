# \QueueDetailsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1QueueDetailsById**](QueueDetailsApi.md#GetApiV1QueueDetailsById) | **Get** /api/v1/queue/details/{id} | 
[**ListApiV1QueueDetails**](QueueDetailsApi.md#ListApiV1QueueDetails) | **Get** /api/v1/queue/details | 



## GetApiV1QueueDetailsById

> QueueResource GetApiV1QueueDetailsById(ctx, id).Execute()



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
    resp, r, err := apiClient.QueueDetailsApi.GetApiV1QueueDetailsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsApi.GetApiV1QueueDetailsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QueueDetailsById`: QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueDetailsApi.GetApiV1QueueDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QueueDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueueResource**](QueueResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1QueueDetails

> []QueueResource ListApiV1QueueDetails(ctx).AuthorId(authorId).BookIds(bookIds).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



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
    authorId := int32(56) // int32 |  (optional)
    bookIds := []int32{int32(123)} // []int32 |  (optional)
    includeAuthor := true // bool |  (optional) (default to false)
    includeBook := true // bool |  (optional) (default to true)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueDetailsApi.ListApiV1QueueDetails(context.Background()).AuthorId(authorId).BookIds(bookIds).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsApi.ListApiV1QueueDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1QueueDetails`: []QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueDetailsApi.ListApiV1QueueDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1QueueDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookIds** | **[]int32** |  | 
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to true]

### Return type

[**[]QueueResource**](QueueResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

