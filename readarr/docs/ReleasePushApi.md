# \ReleasePushApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1ReleasePush**](ReleasePushApi.md#CreateApiV1ReleasePush) | **Post** /api/v1/release/push | 
[**GetApiV1ReleasePushById**](ReleasePushApi.md#GetApiV1ReleasePushById) | **Get** /api/v1/release/push/{id} | 



## CreateApiV1ReleasePush

> ReleaseResource CreateApiV1ReleasePush(ctx).ReleaseResource(releaseResource).Execute()



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
    releaseResource := *readarrClient.NewReleaseResource() // ReleaseResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasePushApi.CreateApiV1ReleasePush(context.Background()).ReleaseResource(releaseResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePushApi.CreateApiV1ReleasePush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1ReleasePush`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleasePushApi.CreateApiV1ReleasePush`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ReleasePushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseResource** | [**ReleaseResource**](ReleaseResource.md) |  | 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1ReleasePushById

> ReleaseResource GetApiV1ReleasePushById(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleasePushApi.GetApiV1ReleasePushById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePushApi.GetApiV1ReleasePushById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ReleasePushById`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleasePushApi.GetApiV1ReleasePushById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ReleasePushByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

