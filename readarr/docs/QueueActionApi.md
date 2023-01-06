# \QueueActionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1QueueGrabBulk**](QueueActionApi.md#CreateApiV1QueueGrabBulk) | **Post** /api/v1/queue/grab/bulk | 
[**CreateApiV1QueueGrabById**](QueueActionApi.md#CreateApiV1QueueGrabById) | **Post** /api/v1/queue/grab/{id} | 



## CreateApiV1QueueGrabBulk

> CreateApiV1QueueGrabBulk(ctx).QueueBulkResource(queueBulkResource).Execute()



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
    queueBulkResource := *readarrClient.NewQueueBulkResource() // QueueBulkResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueActionApi.CreateApiV1QueueGrabBulk(context.Background()).QueueBulkResource(queueBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueActionApi.CreateApiV1QueueGrabBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1QueueGrabBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queueBulkResource** | [**QueueBulkResource**](QueueBulkResource.md) |  | 

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


## CreateApiV1QueueGrabById

> CreateApiV1QueueGrabById(ctx, id).Execute()



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
    resp, r, err := apiClient.QueueActionApi.CreateApiV1QueueGrabById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueActionApi.CreateApiV1QueueGrabById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateApiV1QueueGrabByIdRequest struct via the builder pattern


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

