# \ParseApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1Parse**](ParseApi.md#GetApiV1Parse) | **Get** /api/v1/parse | 



## GetApiV1Parse

> ParseResource GetApiV1Parse(ctx).Title(title).Execute()



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
    title := "title_example" // string |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParseApi.GetApiV1Parse(context.Background()).Title(title).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParseApi.GetApiV1Parse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1Parse`: ParseResource
    fmt.Fprintf(os.Stdout, "Response from `ParseApi.GetApiV1Parse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ParseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** |  | 

### Return type

[**ParseResource**](ParseResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

