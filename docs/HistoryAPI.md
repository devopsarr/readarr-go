# \HistoryAPI

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHistoryFailedById**](HistoryAPI.md#CreateHistoryFailedById) | **Post** /api/v1/history/failed/{id} | 
[**GetHistory**](HistoryAPI.md#GetHistory) | **Get** /api/v1/history | 
[**ListHistoryAuthor**](HistoryAPI.md#ListHistoryAuthor) | **Get** /api/v1/history/author | 
[**ListHistorySince**](HistoryAPI.md#ListHistorySince) | **Get** /api/v1/history/since | 



## CreateHistoryFailedById

> CreateHistoryFailedById(ctx, id).Execute()



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
	id := int32(56) // int32 | 

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	r, err := apiClient.HistoryAPI.CreateHistoryFailedById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.CreateHistoryFailedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHistoryFailedByIdRequest struct via the builder pattern


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


## GetHistory

> HistoryResourcePagingResource GetHistory(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeAuthor(includeAuthor).IncludeBook(includeBook).EventType(eventType).BookId(bookId).DownloadId(downloadId).Execute()



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	sortKey := "sortKey_example" // string |  (optional)
	sortDirection := readarrClient.SortDirection("default") // SortDirection |  (optional)
	includeAuthor := true // bool |  (optional)
	includeBook := true // bool |  (optional)
	eventType := []int32{int32(123)} // []int32 |  (optional)
	bookId := int32(56) // int32 |  (optional)
	downloadId := "downloadId_example" // string |  (optional)

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryAPI.GetHistory(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeAuthor(includeAuthor).IncludeBook(includeBook).EventType(eventType).BookId(bookId).DownloadId(downloadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.GetHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistory`: HistoryResourcePagingResource
	fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 
 **includeAuthor** | **bool** |  | 
 **includeBook** | **bool** |  | 
 **eventType** | **[]int32** |  | 
 **bookId** | **int32** |  | 
 **downloadId** | **string** |  | 

### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistoryAuthor

> []HistoryResource ListHistoryAuthor(ctx).AuthorId(authorId).BookId(bookId).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



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
	bookId := int32(56) // int32 |  (optional)
	eventType := readarrClient.EntityHistoryEventType("unknown") // EntityHistoryEventType |  (optional)
	includeAuthor := true // bool |  (optional) (default to false)
	includeBook := true // bool |  (optional) (default to false)

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryAPI.ListHistoryAuthor(context.Background()).AuthorId(authorId).BookId(bookId).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.ListHistoryAuthor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHistoryAuthor`: []HistoryResource
	fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.ListHistoryAuthor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookId** | **int32** |  | 
 **eventType** | [**EntityHistoryEventType**](EntityHistoryEventType.md) |  | 
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistorySince

> []HistoryResource ListHistorySince(ctx).Date(date).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	readarrClient "github.com/devopsarr/readarr-go/readarr"
)

func main() {
	date := time.Now() // time.Time |  (optional)
	eventType := readarrClient.EntityHistoryEventType("unknown") // EntityHistoryEventType |  (optional)
	includeAuthor := true // bool |  (optional) (default to false)
	includeBook := true // bool |  (optional) (default to false)

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryAPI.ListHistorySince(context.Background()).Date(date).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.ListHistorySince``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHistorySince`: []HistoryResource
	fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.ListHistorySince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistorySinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **eventType** | [**EntityHistoryEventType**](EntityHistoryEventType.md) |  | 
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

