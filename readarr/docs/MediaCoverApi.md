# \MediaCoverApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1MediaCoverAuthorauthorIdByFilename**](MediaCoverApi.md#GetApiV1MediaCoverAuthorauthorIdByFilename) | **Get** /api/v1/mediacover/author/{authorId}/{filename} | 
[**GetApiV1MediaCoverBookbookIdByFilename**](MediaCoverApi.md#GetApiV1MediaCoverBookbookIdByFilename) | **Get** /api/v1/mediacover/book/{bookId}/{filename} | 



## GetApiV1MediaCoverAuthorauthorIdByFilename

> GetApiV1MediaCoverAuthorauthorIdByFilename(ctx, authorId, filename).Execute()



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
    resp, r, err := apiClient.MediaCoverApi.GetApiV1MediaCoverAuthorauthorIdByFilename(context.Background(), authorId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.GetApiV1MediaCoverAuthorauthorIdByFilename``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiV1MediaCoverAuthorauthorIdByFilenameRequest struct via the builder pattern


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


## GetApiV1MediaCoverBookbookIdByFilename

> GetApiV1MediaCoverBookbookIdByFilename(ctx, bookId, filename).Execute()



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
    resp, r, err := apiClient.MediaCoverApi.GetApiV1MediaCoverBookbookIdByFilename(context.Background(), bookId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.GetApiV1MediaCoverBookbookIdByFilename``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiV1MediaCoverBookbookIdByFilenameRequest struct via the builder pattern


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

