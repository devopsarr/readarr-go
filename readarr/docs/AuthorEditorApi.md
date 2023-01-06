# \AuthorEditorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1AuthorEditor**](AuthorEditorApi.md#DeleteApiV1AuthorEditor) | **Delete** /api/v1/author/editor | 
[**PutApiV1AuthorEditor**](AuthorEditorApi.md#PutApiV1AuthorEditor) | **Put** /api/v1/author/editor | 



## DeleteApiV1AuthorEditor

> DeleteApiV1AuthorEditor(ctx).AuthorEditorResource(authorEditorResource).Execute()



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
    resp, r, err := apiClient.AuthorEditorApi.DeleteApiV1AuthorEditor(context.Background()).AuthorEditorResource(authorEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorEditorApi.DeleteApiV1AuthorEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1AuthorEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorEditorResource** | [**AuthorEditorResource**](AuthorEditorResource.md) |  | 

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


## PutApiV1AuthorEditor

> PutApiV1AuthorEditor(ctx).AuthorEditorResource(authorEditorResource).Execute()



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
    resp, r, err := apiClient.AuthorEditorApi.PutApiV1AuthorEditor(context.Background()).AuthorEditorResource(authorEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorEditorApi.PutApiV1AuthorEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1AuthorEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorEditorResource** | [**AuthorEditorResource**](AuthorEditorResource.md) |  | 

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

