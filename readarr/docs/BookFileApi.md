# \BookFileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1BookFile**](BookFileApi.md#DeleteApiV1BookFile) | **Delete** /api/v1/bookfile/{id} | 
[**DeleteApiV1BookFileBulk**](BookFileApi.md#DeleteApiV1BookFileBulk) | **Delete** /api/v1/bookfile/bulk | 
[**GetApiV1BookFileById**](BookFileApi.md#GetApiV1BookFileById) | **Get** /api/v1/bookfile/{id} | 
[**ListApiV1BookFile**](BookFileApi.md#ListApiV1BookFile) | **Get** /api/v1/bookfile | 
[**PutApiV1BookFileEditor**](BookFileApi.md#PutApiV1BookFileEditor) | **Put** /api/v1/bookfile/editor | 
[**UpdateApiV1BookFile**](BookFileApi.md#UpdateApiV1BookFile) | **Put** /api/v1/bookfile/{id} | 



## DeleteApiV1BookFile

> DeleteApiV1BookFile(ctx, id).Execute()



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
    resp, r, err := apiClient.BookFileApi.DeleteApiV1BookFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.DeleteApiV1BookFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1BookFileRequest struct via the builder pattern


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


## DeleteApiV1BookFileBulk

> DeleteApiV1BookFileBulk(ctx).BookFileListResource(bookFileListResource).Execute()



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
    resp, r, err := apiClient.BookFileApi.DeleteApiV1BookFileBulk(context.Background()).BookFileListResource(bookFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.DeleteApiV1BookFileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1BookFileBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookFileListResource** | [**BookFileListResource**](BookFileListResource.md) |  | 

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


## GetApiV1BookFileById

> BookFileResource GetApiV1BookFileById(ctx, id).Execute()



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
    resp, r, err := apiClient.BookFileApi.GetApiV1BookFileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.GetApiV1BookFileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1BookFileById`: BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.GetApiV1BookFileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1BookFileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookFileResource**](BookFileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1BookFile

> []BookFileResource ListApiV1BookFile(ctx).AuthorId(authorId).BookFileIds(bookFileIds).BookId(bookId).Unmapped(unmapped).Execute()



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
    resp, r, err := apiClient.BookFileApi.ListApiV1BookFile(context.Background()).AuthorId(authorId).BookFileIds(bookFileIds).BookId(bookId).Unmapped(unmapped).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.ListApiV1BookFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1BookFile`: []BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.ListApiV1BookFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1BookFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **int32** |  | 
 **bookFileIds** | **[]int32** |  | 
 **bookId** | **[]int32** |  | 
 **unmapped** | **bool** |  | 

### Return type

[**[]BookFileResource**](BookFileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1BookFileEditor

> PutApiV1BookFileEditor(ctx).BookFileListResource(bookFileListResource).Execute()



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
    resp, r, err := apiClient.BookFileApi.PutApiV1BookFileEditor(context.Background()).BookFileListResource(bookFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.PutApiV1BookFileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1BookFileEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookFileListResource** | [**BookFileListResource**](BookFileListResource.md) |  | 

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


## UpdateApiV1BookFile

> BookFileResource UpdateApiV1BookFile(ctx, id).BookFileResource(bookFileResource).Execute()



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
    resp, r, err := apiClient.BookFileApi.UpdateApiV1BookFile(context.Background(), id).BookFileResource(bookFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.UpdateApiV1BookFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1BookFile`: BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.UpdateApiV1BookFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1BookFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookFileResource** | [**BookFileResource**](BookFileResource.md) |  | 

### Return type

[**BookFileResource**](BookFileResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

