# \BookApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Book**](BookApi.md#CreateApiV1Book) | **Post** /api/v1/book | 
[**DeleteApiV1Book**](BookApi.md#DeleteApiV1Book) | **Delete** /api/v1/book/{id} | 
[**GetApiV1BookById**](BookApi.md#GetApiV1BookById) | **Get** /api/v1/book/{id} | 
[**ListApiV1Book**](BookApi.md#ListApiV1Book) | **Get** /api/v1/book | 
[**PutApiV1BookMonitor**](BookApi.md#PutApiV1BookMonitor) | **Put** /api/v1/book/monitor | 
[**UpdateApiV1Book**](BookApi.md#UpdateApiV1Book) | **Put** /api/v1/book/{id} | 



## CreateApiV1Book

> BookResource CreateApiV1Book(ctx).BookResource(bookResource).Execute()



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
    bookResource := *readarrClient.NewBookResource() // BookResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookApi.CreateApiV1Book(context.Background()).BookResource(bookResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.CreateApiV1Book``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Book`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `BookApi.CreateApiV1Book`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1BookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookResource** | [**BookResource**](BookResource.md) |  | 

### Return type

[**BookResource**](BookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1Book

> DeleteApiV1Book(ctx, id).Execute()



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
    resp, r, err := apiClient.BookApi.DeleteApiV1Book(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.DeleteApiV1Book``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1BookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1BookById

> BookResource GetApiV1BookById(ctx, id).Execute()



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
    resp, r, err := apiClient.BookApi.GetApiV1BookById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.GetApiV1BookById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1BookById`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `BookApi.GetApiV1BookById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1BookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookResource**](BookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1Book

> []BookResource ListApiV1Book(ctx).AuthorId(authorId).BookIds(bookIds).TitleSlug(titleSlug).IncludeAllAuthorBooks(includeAllAuthorBooks).Execute()



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
    bookIds := []int32{int32(123)} // []int32 |  (optional)
    titleSlug := "titleSlug_example" // string |  (optional)
    includeAllAuthorBooks := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookApi.ListApiV1Book(context.Background()).AuthorId(authorId).BookIds(bookIds).TitleSlug(titleSlug).IncludeAllAuthorBooks(includeAllAuthorBooks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.ListApiV1Book``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Book`: []BookResource
    fmt.Fprintf(os.Stdout, "Response from `BookApi.ListApiV1Book`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1BookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookIds** | **[]int32** |  | 
 **titleSlug** | **string** |  | 
 **includeAllAuthorBooks** | **bool** |  | [default to false]

### Return type

[**[]BookResource**](BookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1BookMonitor

> PutApiV1BookMonitor(ctx).BooksMonitoredResource(booksMonitoredResource).Execute()



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
    booksMonitoredResource := *readarrClient.NewBooksMonitoredResource() // BooksMonitoredResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookApi.PutApiV1BookMonitor(context.Background()).BooksMonitoredResource(booksMonitoredResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.PutApiV1BookMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1BookMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **booksMonitoredResource** | [**BooksMonitoredResource**](BooksMonitoredResource.md) |  | 

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


## UpdateApiV1Book

> BookResource UpdateApiV1Book(ctx, id).BookResource(bookResource).Execute()



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
    id := "id_example" // string | 
    bookResource := *readarrClient.NewBookResource() // BookResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookApi.UpdateApiV1Book(context.Background(), id).BookResource(bookResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.UpdateApiV1Book``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Book`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `BookApi.UpdateApiV1Book`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1BookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookResource** | [**BookResource**](BookResource.md) |  | 

### Return type

[**BookResource**](BookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

