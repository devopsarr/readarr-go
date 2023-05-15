# \EditionApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEdition**](EditionApi.md#ListEdition) | **Get** /api/v1/edition | 



## ListEdition

> []EditionResource ListEdition(ctx).BookId(bookId).Execute()



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
    bookId := []int32{int32(123)} // []int32 |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.EditionApi.ListEdition(context.Background()).BookId(bookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EditionApi.ListEdition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEdition`: []EditionResource
    fmt.Fprintf(os.Stdout, "Response from `EditionApi.ListEdition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookId** | **[]int32** |  | 

### Return type

[**[]EditionResource**](EditionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

