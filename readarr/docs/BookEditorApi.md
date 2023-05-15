# \BookEditorApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBookEditor**](BookEditorApi.md#DeleteBookEditor) | **Delete** /api/v1/book/editor | 
[**PutBookEditor**](BookEditorApi.md#PutBookEditor) | **Put** /api/v1/book/editor | 



## DeleteBookEditor

> DeleteBookEditor(ctx).BookEditorResource(bookEditorResource).Execute()



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
    bookEditorResource := *readarrClient.NewBookEditorResource() // BookEditorResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookEditorApi.DeleteBookEditor(context.Background()).BookEditorResource(bookEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookEditorApi.DeleteBookEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookEditorResource** | [**BookEditorResource**](BookEditorResource.md) |  | 

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


## PutBookEditor

> PutBookEditor(ctx).BookEditorResource(bookEditorResource).Execute()



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
    bookEditorResource := *readarrClient.NewBookEditorResource() // BookEditorResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookEditorApi.PutBookEditor(context.Background()).BookEditorResource(bookEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookEditorApi.PutBookEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBookEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookEditorResource** | [**BookEditorResource**](BookEditorResource.md) |  | 

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

