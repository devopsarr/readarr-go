# \AuthorLookupAPI

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthorLookup**](AuthorLookupAPI.md#GetAuthorLookup) | **Get** /api/v1/author/lookup | 



## GetAuthorLookup

> GetAuthorLookup(ctx).Term(term).Execute()



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
	term := "term_example" // string |  (optional)

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	r, err := apiClient.AuthorLookupAPI.GetAuthorLookup(context.Background()).Term(term).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorLookupAPI.GetAuthorLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

