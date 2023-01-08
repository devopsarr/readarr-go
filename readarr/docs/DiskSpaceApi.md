# \DiskSpaceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApiV1DiskSpace**](DiskSpaceApi.md#ListApiV1DiskSpace) | **Get** /api/v1/diskspace | 



## ListApiV1DiskSpace

> []DiskSpaceResource ListApiV1DiskSpace(ctx).Execute()



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
    resp, r, err := apiClient.DiskSpaceApi.ListApiV1DiskSpace(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiskSpaceApi.ListApiV1DiskSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1DiskSpace`: []DiskSpaceResource
    fmt.Fprintf(os.Stdout, "Response from `DiskSpaceApi.ListApiV1DiskSpace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1DiskSpaceRequest struct via the builder pattern


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

