# \MediaCoverApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1MediacoverAuthorauthoridByFilename**](MediaCoverApi.md#GetApiV1MediacoverAuthorauthoridByFilename) | **Get** /api/v1/mediacover/author/{authorId}/{filename} | 
[**GetApiV1MediacoverBookbookidByFilename**](MediaCoverApi.md#GetApiV1MediacoverBookbookidByFilename) | **Get** /api/v1/mediacover/book/{bookId}/{filename} | 



## GetApiV1MediacoverAuthorauthoridByFilename

> GetApiV1MediacoverAuthorauthoridByFilename(ctx, authorId, filename).Execute()



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
    authorId := int32(56) // int32 | 
    filename := "filename_example" // string | 

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaCoverApi.GetApiV1MediacoverAuthorauthoridByFilename(context.Background(), authorId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.GetApiV1MediacoverAuthorauthoridByFilename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorId** | **int32** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MediacoverAuthorauthoridByFilenameRequest struct via the builder pattern


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


## GetApiV1MediacoverBookbookidByFilename

> GetApiV1MediacoverBookbookidByFilename(ctx, bookId, filename).Execute()



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
    bookId := int32(56) // int32 | 
    filename := "filename_example" // string | 

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaCoverApi.GetApiV1MediacoverBookbookidByFilename(context.Background(), bookId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.GetApiV1MediacoverBookbookidByFilename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **int32** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MediacoverBookbookidByFilenameRequest struct via the builder pattern


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

