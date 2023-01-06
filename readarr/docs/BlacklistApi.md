# \BlacklistApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1Blacklist**](BlacklistApi.md#DeleteApiV1Blacklist) | **Delete** /api/v1/blacklist/{id} | 
[**DeleteApiV1BlacklistBulk**](BlacklistApi.md#DeleteApiV1BlacklistBulk) | **Delete** /api/v1/blacklist/bulk | 
[**GetApiV1Blacklist**](BlacklistApi.md#GetApiV1Blacklist) | **Get** /api/v1/blacklist | 



## DeleteApiV1Blacklist

> DeleteApiV1Blacklist(ctx, id).Execute()



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
    resp, r, err := apiClient.BlacklistApi.DeleteApiV1Blacklist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.DeleteApiV1Blacklist``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1BlacklistRequest struct via the builder pattern


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


## DeleteApiV1BlacklistBulk

> DeleteApiV1BlacklistBulk(ctx).BlacklistBulkResource(blacklistBulkResource).Execute()



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
    blacklistBulkResource := *readarrClient.NewBlacklistBulkResource() // BlacklistBulkResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlacklistApi.DeleteApiV1BlacklistBulk(context.Background()).BlacklistBulkResource(blacklistBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.DeleteApiV1BlacklistBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1BlacklistBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blacklistBulkResource** | [**BlacklistBulkResource**](BlacklistBulkResource.md) |  | 

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


## GetApiV1Blacklist

> BlacklistResourcePagingResource GetApiV1Blacklist(ctx).Execute()



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
    resp, r, err := apiClient.BlacklistApi.GetApiV1Blacklist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.GetApiV1Blacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1Blacklist`: BlacklistResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `BlacklistApi.GetApiV1Blacklist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1BlacklistRequest struct via the builder pattern


### Return type

[**BlacklistResourcePagingResource**](BlacklistResourcePagingResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

