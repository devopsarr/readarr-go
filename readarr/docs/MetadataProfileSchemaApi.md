# \MetadataProfileSchemaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1MetadataprofileSchema**](MetadataProfileSchemaApi.md#GetApiV1MetadataprofileSchema) | **Get** /api/v1/metadataprofile/schema | 



## GetApiV1MetadataprofileSchema

> MetadataProfileResource GetApiV1MetadataprofileSchema(ctx).Execute()



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
    resp, r, err := apiClient.MetadataProfileSchemaApi.GetApiV1MetadataprofileSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileSchemaApi.GetApiV1MetadataprofileSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1MetadataprofileSchema`: MetadataProfileResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileSchemaApi.GetApiV1MetadataprofileSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MetadataprofileSchemaRequest struct via the builder pattern


### Return type

[**MetadataProfileResource**](MetadataProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

