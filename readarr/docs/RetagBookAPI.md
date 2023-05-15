# \RetagBookAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRetag**](RetagBookAPI.md#ListRetag) | **Get** /api/v1/retag | 



## ListRetag

> []RetagBookResource ListRetag(ctx).AuthorId(authorId).BookId(bookId).Execute()



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
    bookId := int32(56) // int32 |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetagBookAPI.ListRetag(context.Background()).AuthorId(authorId).BookId(bookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetagBookAPI.ListRetag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRetag`: []RetagBookResource
    fmt.Fprintf(os.Stdout, "Response from `RetagBookAPI.ListRetag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRetagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookId** | **int32** |  | 

### Return type

[**[]RetagBookResource**](RetagBookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

