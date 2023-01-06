# \IndexerConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1ConfigIndexer**](IndexerConfigApi.md#GetApiV1ConfigIndexer) | **Get** /api/v1/config/indexer | 
[**GetApiV1ConfigIndexerById**](IndexerConfigApi.md#GetApiV1ConfigIndexerById) | **Get** /api/v1/config/indexer/{id} | 
[**UpdateApiV1ConfigIndexer**](IndexerConfigApi.md#UpdateApiV1ConfigIndexer) | **Put** /api/v1/config/indexer/{id} | 



## GetApiV1ConfigIndexer

> IndexerConfigResource GetApiV1ConfigIndexer(ctx).Execute()



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
    resp, r, err := apiClient.IndexerConfigApi.GetApiV1ConfigIndexer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerConfigApi.GetApiV1ConfigIndexer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigIndexer`: IndexerConfigResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerConfigApi.GetApiV1ConfigIndexer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigIndexerRequest struct via the builder pattern


### Return type

[**IndexerConfigResource**](IndexerConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1ConfigIndexerById

> IndexerConfigResource GetApiV1ConfigIndexerById(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerConfigApi.GetApiV1ConfigIndexerById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerConfigApi.GetApiV1ConfigIndexerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigIndexerById`: IndexerConfigResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerConfigApi.GetApiV1ConfigIndexerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigIndexerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexerConfigResource**](IndexerConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1ConfigIndexer

> IndexerConfigResource UpdateApiV1ConfigIndexer(ctx, id).IndexerConfigResource(indexerConfigResource).Execute()



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
    id := "id_example" // string | 
    indexerConfigResource := *readarrClient.NewIndexerConfigResource() // IndexerConfigResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerConfigApi.UpdateApiV1ConfigIndexer(context.Background(), id).IndexerConfigResource(indexerConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerConfigApi.UpdateApiV1ConfigIndexer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ConfigIndexer`: IndexerConfigResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerConfigApi.UpdateApiV1ConfigIndexer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ConfigIndexerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexerConfigResource** | [**IndexerConfigResource**](IndexerConfigResource.md) |  | 

### Return type

[**IndexerConfigResource**](IndexerConfigResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

