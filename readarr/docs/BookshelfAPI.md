# \BookshelfAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBookshelf**](BookshelfAPI.md#CreateBookshelf) | **Post** /api/v1/bookshelf | 



## CreateBookshelf

> CreateBookshelf(ctx).BookshelfResource(bookshelfResource).Execute()



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
    resp, r, err := apiClient.BookshelfAPI.CreateBookshelf(context.Background()).BookshelfResource(bookshelfResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookshelfAPI.CreateBookshelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBookshelfRequest struct via the builder pattern


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

