# \SeriesAPI

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSeries**](SeriesAPI.md#ListSeries) | **Get** /api/v1/series | 



## ListSeries

> []SeriesResource ListSeries(ctx).AuthorId(authorId).Execute()



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
	authorId := int32(56) // int32 |  (optional)

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeriesAPI.ListSeries(context.Background()).AuthorId(authorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.ListSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSeries`: []SeriesResource
	fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.ListSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 

### Return type

[**[]SeriesResource**](SeriesResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

