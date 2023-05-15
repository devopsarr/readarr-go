# \BookshelfApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBookshelf**](BookshelfApi.md#CreateBookshelf) | **Post** /api/v1/bookshelf | 



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
    resp, r, err := apiClient.BookshelfApi.CreateBookshelf(context.Background()).BookshelfResource(bookshelfResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookshelfApi.CreateBookshelf``: %v\n", err)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

