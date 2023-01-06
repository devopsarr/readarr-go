# \SeriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApiV1Series**](SeriesApi.md#ListApiV1Series) | **Get** /api/v1/series | 



## ListApiV1Series

> []SeriesResource ListApiV1Series(ctx).AuthorId(authorId).Execute()



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

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesApi.ListApiV1Series(context.Background()).AuthorId(authorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesApi.ListApiV1Series``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Series`: []SeriesResource
    fmt.Fprintf(os.Stdout, "Response from `SeriesApi.ListApiV1Series`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1SeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 

### Return type

[**[]SeriesResource**](SeriesResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

