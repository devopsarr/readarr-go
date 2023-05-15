# \AuthorEditorApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthorEditor**](AuthorEditorApi.md#DeleteAuthorEditor) | **Delete** /api/v1/author/editor | 
[**PutAuthorEditor**](AuthorEditorApi.md#PutAuthorEditor) | **Put** /api/v1/author/editor | 



## DeleteAuthorEditor

> DeleteAuthorEditor(ctx).AuthorEditorResource(authorEditorResource).Execute()



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
    authorEditorResource := *readarrClient.NewAuthorEditorResource() // AuthorEditorResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorEditorApi.DeleteAuthorEditor(context.Background()).AuthorEditorResource(authorEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorEditorApi.DeleteAuthorEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorEditorResource** | [**AuthorEditorResource**](AuthorEditorResource.md) |  | 

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


## PutAuthorEditor

> PutAuthorEditor(ctx).AuthorEditorResource(authorEditorResource).Execute()



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
    authorEditorResource := *readarrClient.NewAuthorEditorResource() // AuthorEditorResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorEditorApi.PutAuthorEditor(context.Background()).AuthorEditorResource(authorEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorEditorApi.PutAuthorEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAuthorEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorEditorResource** | [**AuthorEditorResource**](AuthorEditorResource.md) |  | 

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

