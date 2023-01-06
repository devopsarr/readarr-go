# \QueueStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1QueueStatus**](QueueStatusApi.md#GetApiV1QueueStatus) | **Get** /api/v1/queue/status | 
[**GetApiV1QueueStatusById**](QueueStatusApi.md#GetApiV1QueueStatusById) | **Get** /api/v1/queue/status/{id} | 



## GetApiV1QueueStatus

> QueueStatusResource GetApiV1QueueStatus(ctx).Execute()



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

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueStatusApi.GetApiV1QueueStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueStatusApi.GetApiV1QueueStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QueueStatus`: QueueStatusResource
    fmt.Fprintf(os.Stdout, "Response from `QueueStatusApi.GetApiV1QueueStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QueueStatusRequest struct via the builder pattern


### Return type

[**QueueStatusResource**](QueueStatusResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1QueueStatusById

> QueueStatusResource GetApiV1QueueStatusById(ctx, id).Execute()



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
    resp, r, err := apiClient.QueueStatusApi.GetApiV1QueueStatusById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueStatusApi.GetApiV1QueueStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QueueStatusById`: QueueStatusResource
    fmt.Fprintf(os.Stdout, "Response from `QueueStatusApi.GetApiV1QueueStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QueueStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueueStatusResource**](QueueStatusResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

