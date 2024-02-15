# \MetadataProfileSchemaAPI

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetadataprofileSchema**](MetadataProfileSchemaAPI.md#GetMetadataprofileSchema) | **Get** /api/v1/metadataprofile/schema | 



## GetMetadataprofileSchema

> MetadataProfileResource GetMetadataprofileSchema(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	readarrClient "github.com/devopsarr/readarr-go/readarr"
)

func main() {

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataProfileSchemaAPI.GetMetadataprofileSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileSchemaAPI.GetMetadataprofileSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataprofileSchema`: MetadataProfileResource
	fmt.Fprintf(os.Stdout, "Response from `MetadataProfileSchemaAPI.GetMetadataprofileSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataprofileSchemaRequest struct via the builder pattern


### Return type

[**MetadataProfileResource**](MetadataProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

