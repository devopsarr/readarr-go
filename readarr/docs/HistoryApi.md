# \HistoryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1HistoryFailed**](HistoryApi.md#CreateApiV1HistoryFailed) | **Post** /api/v1/history/failed | 
[**GetApiV1History**](HistoryApi.md#GetApiV1History) | **Get** /api/v1/history | 
[**ListApiV1HistoryAuthor**](HistoryApi.md#ListApiV1HistoryAuthor) | **Get** /api/v1/history/author | 
[**ListApiV1HistorySince**](HistoryApi.md#ListApiV1HistorySince) | **Get** /api/v1/history/since | 



## CreateApiV1HistoryFailed

> CreateApiV1HistoryFailed(ctx).Body(body).Execute()



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
    body := int32(56) // int32 |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.CreateApiV1HistoryFailed(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.CreateApiV1HistoryFailed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1HistoryFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1History

> HistoryResourcePagingResource GetApiV1History(ctx).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



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
    resp, r, err := apiClient.HistoryApi.GetApiV1History(context.Background()).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetApiV1History``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1History`: HistoryResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetApiV1History`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1HistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to false]

### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1HistoryAuthor

> []HistoryResource ListApiV1HistoryAuthor(ctx).AuthorId(authorId).BookId(bookId).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



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
    eventType := readarrClient.HistoryEventType("unknown") // HistoryEventType |  (optional)
    includeAuthor := true // bool |  (optional) (default to false)
    includeBook := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ListApiV1HistoryAuthor(context.Background()).AuthorId(authorId).BookId(bookId).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListApiV1HistoryAuthor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1HistoryAuthor`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListApiV1HistoryAuthor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1HistoryAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookId** | **int32** |  | 
 **eventType** | [**HistoryEventType**](HistoryEventType.md) |  | 
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1HistorySince

> []HistoryResource ListApiV1HistorySince(ctx).Date(date).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



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
    eventType := readarrClient.HistoryEventType("unknown") // HistoryEventType |  (optional)
    includeAuthor := true // bool |  (optional) (default to false)
    includeBook := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ListApiV1HistorySince(context.Background()).Date(date).EventType(eventType).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListApiV1HistorySince``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1HistorySince`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListApiV1HistorySince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1HistorySinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **eventType** | [**HistoryEventType**](HistoryEventType.md) |  | 
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

