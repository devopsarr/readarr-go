# \HistoryApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHistoryFailedById**](HistoryApi.md#CreateHistoryFailedById) | **Post** /api/v1/history/failed/{id} | 
[**GetHistory**](HistoryApi.md#GetHistory) | **Get** /api/v1/history | 
[**ListHistoryAuthor**](HistoryApi.md#ListHistoryAuthor) | **Get** /api/v1/history/author | 
[**ListHistorySince**](HistoryApi.md#ListHistorySince) | **Get** /api/v1/history/since | 



## CreateHistoryFailedById

> CreateHistoryFailedById(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.CreateHistoryFailedById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.CreateHistoryFailedById``: %v\n", err)
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

> HistoryResourcePagingResource GetHistory(ctx).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



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
    includeAuthor := true // bool |  (optional) (default to false)
    includeBook := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.GetHistory(context.Background()).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistory`: HistoryResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to false]

### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

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
    readarrClient "./openapi"
)

func main() {
    authorId := int32(56) // int32 |  (optional)
    bookId := int32(56) // int32 |  (optional)
    eventType := readarrClient.EntityHistoryEventType("unknown") // EntityHistoryEventType |  (optional)
    includeAuthor := true // bool |  (optional) (default to false)
    includeBook := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ListHistoryAuthor(context.Background()).AuthorId(authorId).BookId(bookId).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListHistoryAuthor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryAuthor`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListHistoryAuthor`: %v\n", resp)
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
    readarrClient "./openapi"
)

func main() {
    date := time.Now() // time.Time |  (optional)
    eventType := readarrClient.EntityHistoryEventType("unknown") // EntityHistoryEventType |  (optional)
    includeAuthor := true // bool |  (optional) (default to false)
    includeBook := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ListHistorySince(context.Background()).Date(date).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListHistorySince``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistorySince`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListHistorySince`: %v\n", resp)
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
