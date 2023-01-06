# \QualityProfileSchemaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1QualityprofileSchema**](QualityProfileSchemaApi.md#GetApiV1QualityprofileSchema) | **Get** /api/v1/qualityprofile/schema | 



## GetApiV1QualityprofileSchema

> QualityProfileResource GetApiV1QualityprofileSchema(ctx).Execute()



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
    resp, r, err := apiClient.QualityProfileSchemaApi.GetApiV1QualityprofileSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileSchemaApi.GetApiV1QualityprofileSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QualityprofileSchema`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileSchemaApi.GetApiV1QualityprofileSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QualityprofileSchemaRequest struct via the builder pattern


### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

