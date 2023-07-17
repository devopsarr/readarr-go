# \BookFileApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBookFile**](BookFileApi.md#DeleteBookFile) | **Delete** /api/v1/bookfile/{id} | 
[**DeleteBookFileBulk**](BookFileApi.md#DeleteBookFileBulk) | **Delete** /api/v1/bookfile/bulk | 
[**GetBookFileById**](BookFileApi.md#GetBookFileById) | **Get** /api/v1/bookfile/{id} | 
[**ListBookFile**](BookFileApi.md#ListBookFile) | **Get** /api/v1/bookfile | 
[**PutBookFileEditor**](BookFileApi.md#PutBookFileEditor) | **Put** /api/v1/bookfile/editor | 
[**UpdateBookFile**](BookFileApi.md#UpdateBookFile) | **Put** /api/v1/bookfile/{id} | 



## DeleteBookFile

> DeleteBookFile(ctx, id).Execute()



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
    resp, r, err := apiClient.BookFileApi.DeleteBookFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.DeleteBookFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteBookFileRequest struct via the builder pattern


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


## DeleteBookFileBulk

> DeleteBookFileBulk(ctx).BookFileListResource(bookFileListResource).Execute()



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
    bookFileListResource := *readarrClient.NewBookFileListResource() // BookFileListResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookFileApi.DeleteBookFileBulk(context.Background()).BookFileListResource(bookFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.DeleteBookFileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookFileBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookFileListResource** | [**BookFileListResource**](BookFileListResource.md) |  | 

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


## GetBookFileById

> BookFileResource GetBookFileById(ctx, id).Execute()



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
    resp, r, err := apiClient.BookFileApi.GetBookFileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.GetBookFileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookFileById`: BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.GetBookFileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookFileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookFileResource**](BookFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBookFile

> []BookFileResource ListBookFile(ctx).AuthorId(authorId).BookFileIds(bookFileIds).BookId(bookId).Unmapped(unmapped).Execute()



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
    bookFileIds := []int32{int32(123)} // []int32 |  (optional)
    bookId := []int32{int32(123)} // []int32 |  (optional)
    unmapped := true // bool |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookFileApi.ListBookFile(context.Background()).AuthorId(authorId).BookFileIds(bookFileIds).BookId(bookId).Unmapped(unmapped).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.ListBookFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBookFile`: []BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.ListBookFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBookFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookFileIds** | **[]int32** |  | 
 **bookId** | **[]int32** |  | 
 **unmapped** | **bool** |  | 

### Return type

[**[]BookFileResource**](BookFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBookFileEditor

> PutBookFileEditor(ctx).BookFileListResource(bookFileListResource).Execute()



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
    bookFileListResource := *readarrClient.NewBookFileListResource() // BookFileListResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookFileApi.PutBookFileEditor(context.Background()).BookFileListResource(bookFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.PutBookFileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBookFileEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookFileListResource** | [**BookFileListResource**](BookFileListResource.md) |  | 

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


## UpdateBookFile

> BookFileResource UpdateBookFile(ctx, id).BookFileResource(bookFileResource).Execute()



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
    bookFileResource := *readarrClient.NewBookFileResource() // BookFileResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookFileApi.UpdateBookFile(context.Background(), id).BookFileResource(bookFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.UpdateBookFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBookFile`: BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.UpdateBookFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBookFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookFileResource** | [**BookFileResource**](BookFileResource.md) |  | 

### Return type

[**BookFileResource**](BookFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

