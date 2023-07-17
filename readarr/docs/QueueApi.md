# \QueueApi

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteQueue**](QueueApi.md#DeleteQueue) | **Delete** /api/v1/queue/{id} | 
[**DeleteQueueBulk**](QueueApi.md#DeleteQueueBulk) | **Delete** /api/v1/queue/bulk | 
[**GetQueue**](QueueApi.md#GetQueue) | **Get** /api/v1/queue | 



## DeleteQueue

> DeleteQueue(ctx, id).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipReDownload(skipReDownload).Execute()



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
    removeFromClient := true // bool |  (optional) (default to true)
    blocklist := true // bool |  (optional) (default to false)
    skipReDownload := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.DeleteQueue(context.Background(), id).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipReDownload(skipReDownload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.DeleteQueue``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeFromClient** | **bool** |  | [default to true]
 **blocklist** | **bool** |  | [default to false]
 **skipReDownload** | **bool** |  | [default to false]

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


## DeleteQueueBulk

> DeleteQueueBulk(ctx).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipReDownload(skipReDownload).QueueBulkResource(queueBulkResource).Execute()



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
    removeFromClient := true // bool |  (optional) (default to true)
    blocklist := true // bool |  (optional) (default to false)
    skipReDownload := true // bool |  (optional) (default to false)
    queueBulkResource := *readarrClient.NewQueueBulkResource() // QueueBulkResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.DeleteQueueBulk(context.Background()).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipReDownload(skipReDownload).QueueBulkResource(queueBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.DeleteQueueBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeFromClient** | **bool** |  | [default to true]
 **blocklist** | **bool** |  | [default to false]
 **skipReDownload** | **bool** |  | [default to false]
 **queueBulkResource** | [**QueueBulkResource**](QueueBulkResource.md) |  | 

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


## GetQueue

> QueueResourcePagingResource GetQueue(ctx).IncludeUnknownAuthorItems(includeUnknownAuthorItems).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()



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
    includeUnknownAuthorItems := true // bool |  (optional) (default to false)
    includeAuthor := true // bool |  (optional) (default to false)
    includeBook := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.GetQueue(context.Background()).IncludeUnknownAuthorItems(includeUnknownAuthorItems).IncludeAuthor(includeAuthor).IncludeBook(includeBook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.GetQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueue`: QueueResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `QueueApi.GetQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUnknownAuthorItems** | **bool** |  | [default to false]
 **includeAuthor** | **bool** |  | [default to false]
 **includeBook** | **bool** |  | [default to false]

### Return type

[**QueueResourcePagingResource**](QueueResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

