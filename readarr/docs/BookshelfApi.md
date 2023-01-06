# \BookshelfApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Bookshelf**](BookshelfApi.md#CreateApiV1Bookshelf) | **Post** /api/v1/bookshelf | 



## CreateApiV1Bookshelf

> CreateApiV1Bookshelf(ctx).BookshelfResource(bookshelfResource).Execute()



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
    bookshelfResource := *readarrClient.NewBookshelfResource() // BookshelfResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookshelfApi.CreateApiV1Bookshelf(context.Background()).BookshelfResource(bookshelfResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookshelfApi.CreateApiV1Bookshelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1BookshelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookshelfResource** | [**BookshelfResource**](BookshelfResource.md) |  | 

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

