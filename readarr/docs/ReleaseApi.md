# \ReleaseApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Release**](ReleaseApi.md#CreateApiV1Release) | **Post** /api/v1/release | 
[**GetApiV1ReleaseById**](ReleaseApi.md#GetApiV1ReleaseById) | **Get** /api/v1/release/{id} | 
[**ListApiV1Release**](ReleaseApi.md#ListApiV1Release) | **Get** /api/v1/release | 



## CreateApiV1Release

> ReleaseResource CreateApiV1Release(ctx).ReleaseResource(releaseResource).Execute()



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
    resp, r, err := apiClient.ReleaseApi.CreateApiV1Release(context.Background()).ReleaseResource(releaseResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.CreateApiV1Release``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Release`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.CreateApiV1Release`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ReleaseRequest struct via the builder pattern


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


## GetApiV1ReleaseById

> ReleaseResource GetApiV1ReleaseById(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleaseApi.GetApiV1ReleaseById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetApiV1ReleaseById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ReleaseById`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetApiV1ReleaseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ReleaseByIdRequest struct via the builder pattern


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


## ListApiV1Release

> []ReleaseResource ListApiV1Release(ctx).BookId(bookId).AuthorId(authorId).Execute()



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
    bookId := int32(56) // int32 |  (optional)
    authorId := int32(56) // int32 |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.ListApiV1Release(context.Background()).BookId(bookId).AuthorId(authorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.ListApiV1Release``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Release`: []ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.ListApiV1Release`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookId** | **int32** |  | 
 **authorId** | **int32** |  | 

### Return type

[**[]ReleaseResource**](ReleaseResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

