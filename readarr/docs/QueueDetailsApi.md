# \QueueDetailsApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQueueDetails**](QueueDetailsApi.md#ListQueueDetails) | **Get** /api/v1/queue/details | 



## ListQueueDetails

> []QueueResource ListQueueDetails(ctx).AuthorId(authorId).BookIds(bookIds).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



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
    resp, r, err := apiClient.QueueDetailsApi.ListQueueDetails(context.Background()).AuthorId(authorId).BookIds(bookIds).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsApi.ListQueueDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueueDetails`: []QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueDetailsApi.ListQueueDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQueueDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookIds** | **[]int32** |  | 
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to true]

### Return type

[**[]QueueResource**](QueueResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

