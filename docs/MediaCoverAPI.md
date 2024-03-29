# \MediaCoverAPI

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMediaCoverAuthorByFilename**](MediaCoverAPI.md#GetMediaCoverAuthorByFilename) | **Get** /api/v1/mediacover/author/{authorId}/{filename} | 
[**GetMediaCoverBookByFilename**](MediaCoverAPI.md#GetMediaCoverBookByFilename) | **Get** /api/v1/mediacover/book/{bookId}/{filename} | 



## GetMediaCoverAuthorByFilename

> GetMediaCoverAuthorByFilename(ctx, authorId, filename).Execute()



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
	authorId := int32(56) // int32 | 
	filename := "filename_example" // string | 

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	r, err := apiClient.MediaCoverAPI.GetMediaCoverAuthorByFilename(context.Background(), authorId, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverAPI.GetMediaCoverAuthorByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorId** | **int32** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaCoverAuthorByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetMediaCoverBookByFilename

> GetMediaCoverBookByFilename(ctx, bookId, filename).Execute()



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
	bookId := int32(56) // int32 | 
	filename := "filename_example" // string | 

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	r, err := apiClient.MediaCoverAPI.GetMediaCoverBookByFilename(context.Background(), bookId, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverAPI.GetMediaCoverBookByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **int32** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaCoverBookByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

