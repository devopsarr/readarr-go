# \DiskSpaceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDiskSpace**](DiskSpaceApi.md#ListDiskSpace) | **Get** /api/v1/diskspace | 



## ListDiskSpace

> []DiskSpaceResource ListDiskSpace(ctx).Execute()



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

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiskSpaceApi.ListDiskSpace(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiskSpaceApi.ListDiskSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiskSpace`: []DiskSpaceResource
    fmt.Fprintf(os.Stdout, "Response from `DiskSpaceApi.ListDiskSpace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDiskSpaceRequest struct via the builder pattern


### Return type

[**[]DiskSpaceResource**](DiskSpaceResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

