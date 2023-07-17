# \BookApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBook**](BookApi.md#CreateBook) | **Post** /api/v1/book | 
[**DeleteBook**](BookApi.md#DeleteBook) | **Delete** /api/v1/book/{id} | 
[**GetBookById**](BookApi.md#GetBookById) | **Get** /api/v1/book/{id} | 
[**GetBookidOverview**](BookApi.md#GetBookidOverview) | **Get** /api/v1/book/{id}/overview | 
[**ListBook**](BookApi.md#ListBook) | **Get** /api/v1/book | 
[**PutBookMonitor**](BookApi.md#PutBookMonitor) | **Put** /api/v1/book/monitor | 
[**UpdateBook**](BookApi.md#UpdateBook) | **Put** /api/v1/book/{id} | 



## CreateBook

> BookResource CreateBook(ctx).BookResource(bookResource).Execute()



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
    resp, r, err := apiClient.BookApi.CreateBook(context.Background()).BookResource(bookResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.CreateBook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBook`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `BookApi.CreateBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookResource** | [**BookResource**](BookResource.md) |  | 

### Return type

[**BookResource**](BookResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBook

> DeleteBook(ctx, id).DeleteFiles(deleteFiles).AddImportListExclusion(addImportListExclusion).Execute()



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
    deleteFiles := true // bool |  (optional) (default to false)
    addImportListExclusion := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookApi.DeleteBook(context.Background(), id).DeleteFiles(deleteFiles).AddImportListExclusion(addImportListExclusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.DeleteBook``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteFiles** | **bool** |  | [default to false]
 **addImportListExclusion** | **bool** |  | [default to false]

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


## GetBookById

> BookResource GetBookById(ctx, id).Execute()



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
    resp, r, err := apiClient.BookApi.GetBookById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.GetBookById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookById`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `BookApi.GetBookById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookResource**](BookResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookidOverview

> GetBookidOverview(ctx, id).Execute()



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
    resp, r, err := apiClient.BookApi.GetBookidOverview(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.GetBookidOverview``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetBookidOverviewRequest struct via the builder pattern


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


## ListBook

> []BookResource ListBook(ctx).AuthorId(authorId).BookIds(bookIds).TitleSlug(titleSlug).IncludeAllAuthorBooks(includeAllAuthorBooks).Execute()



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
    resp, r, err := apiClient.BookApi.ListBook(context.Background()).AuthorId(authorId).BookIds(bookIds).TitleSlug(titleSlug).IncludeAllAuthorBooks(includeAllAuthorBooks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.ListBook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBook`: []BookResource
    fmt.Fprintf(os.Stdout, "Response from `BookApi.ListBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookIds** | **[]int32** |  | 
 **titleSlug** | **string** |  | 
 **includeAllAuthorBooks** | **bool** |  | [default to false]

### Return type

[**[]BookResource**](BookResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBookMonitor

> PutBookMonitor(ctx).BooksMonitoredResource(booksMonitoredResource).Execute()



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
    resp, r, err := apiClient.BookApi.PutBookMonitor(context.Background()).BooksMonitoredResource(booksMonitoredResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.PutBookMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBookMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **booksMonitoredResource** | [**BooksMonitoredResource**](BooksMonitoredResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBook

> BookResource UpdateBook(ctx, id).BookResource(bookResource).Execute()



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
    resp, r, err := apiClient.BookApi.UpdateBook(context.Background(), id).BookResource(bookResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.UpdateBook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBook`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `BookApi.UpdateBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookResource** | [**BookResource**](BookResource.md) |  | 

### Return type

[**BookResource**](BookResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

