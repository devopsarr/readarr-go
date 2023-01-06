# \RenameBookApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApiV1Rename**](RenameBookApi.md#ListApiV1Rename) | **Get** /api/v1/rename | 



## ListApiV1Rename

> []RenameBookResource ListApiV1Rename(ctx).AuthorId(authorId).BookId(bookId).Execute()



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
    resp, r, err := apiClient.RenameBookApi.ListApiV1Rename(context.Background()).AuthorId(authorId).BookId(bookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenameBookApi.ListApiV1Rename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Rename`: []RenameBookResource
    fmt.Fprintf(os.Stdout, "Response from `RenameBookApi.ListApiV1Rename`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1RenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookId** | **int32** |  | 

### Return type

[**[]RenameBookResource**](RenameBookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

