# \BookFileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1Bookfile**](BookFileApi.md#DeleteApiV1Bookfile) | **Delete** /api/v1/bookfile/{id} | 
[**DeleteApiV1BookfileBulk**](BookFileApi.md#DeleteApiV1BookfileBulk) | **Delete** /api/v1/bookfile/bulk | 
[**GetApiV1BookfileById**](BookFileApi.md#GetApiV1BookfileById) | **Get** /api/v1/bookfile/{id} | 
[**ListApiV1Bookfile**](BookFileApi.md#ListApiV1Bookfile) | **Get** /api/v1/bookfile | 
[**PutApiV1BookfileEditor**](BookFileApi.md#PutApiV1BookfileEditor) | **Put** /api/v1/bookfile/editor | 
[**UpdateApiV1Bookfile**](BookFileApi.md#UpdateApiV1Bookfile) | **Put** /api/v1/bookfile/{id} | 



## DeleteApiV1Bookfile

> DeleteApiV1Bookfile(ctx, id).Execute()



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
    resp, r, err := apiClient.BookFileApi.DeleteApiV1Bookfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.DeleteApiV1Bookfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1BookfileRequest struct via the builder pattern


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


## DeleteApiV1BookfileBulk

> DeleteApiV1BookfileBulk(ctx).BookFileListResource(bookFileListResource).Execute()



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
    resp, r, err := apiClient.BookFileApi.DeleteApiV1BookfileBulk(context.Background()).BookFileListResource(bookFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.DeleteApiV1BookfileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1BookfileBulkRequest struct via the builder pattern


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


## GetApiV1BookfileById

> BookFileResource GetApiV1BookfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.BookFileApi.GetApiV1BookfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.GetApiV1BookfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1BookfileById`: BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.GetApiV1BookfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1BookfileByIdRequest struct via the builder pattern


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


## ListApiV1Bookfile

> []BookFileResource ListApiV1Bookfile(ctx).AuthorId(authorId).BookFileIds(bookFileIds).BookId(bookId).Unmapped(unmapped).Execute()



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
    resp, r, err := apiClient.BookFileApi.ListApiV1Bookfile(context.Background()).AuthorId(authorId).BookFileIds(bookFileIds).BookId(bookId).Unmapped(unmapped).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.ListApiV1Bookfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Bookfile`: []BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.ListApiV1Bookfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1BookfileRequest struct via the builder pattern


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


## PutApiV1BookfileEditor

> PutApiV1BookfileEditor(ctx).BookFileListResource(bookFileListResource).Execute()



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
    resp, r, err := apiClient.BookFileApi.PutApiV1BookfileEditor(context.Background()).BookFileListResource(bookFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.PutApiV1BookfileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1BookfileEditorRequest struct via the builder pattern


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


## UpdateApiV1Bookfile

> BookFileResource UpdateApiV1Bookfile(ctx, id).BookFileResource(bookFileResource).Execute()



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
    resp, r, err := apiClient.BookFileApi.UpdateApiV1Bookfile(context.Background(), id).BookFileResource(bookFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookFileApi.UpdateApiV1Bookfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Bookfile`: BookFileResource
    fmt.Fprintf(os.Stdout, "Response from `BookFileApi.UpdateApiV1Bookfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1BookfileRequest struct via the builder pattern


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

