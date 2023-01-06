# \MetadataApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Metadata**](MetadataApi.md#CreateApiV1Metadata) | **Post** /api/v1/metadata | 
[**CreateApiV1MetadataActionByName**](MetadataApi.md#CreateApiV1MetadataActionByName) | **Post** /api/v1/metadata/action/{name} | 
[**DeleteApiV1Metadata**](MetadataApi.md#DeleteApiV1Metadata) | **Delete** /api/v1/metadata/{id} | 
[**GetApiV1MetadataById**](MetadataApi.md#GetApiV1MetadataById) | **Get** /api/v1/metadata/{id} | 
[**ListApiV1Metadata**](MetadataApi.md#ListApiV1Metadata) | **Get** /api/v1/metadata | 
[**ListApiV1MetadataSchema**](MetadataApi.md#ListApiV1MetadataSchema) | **Get** /api/v1/metadata/schema | 
[**TestApiV1Metadata**](MetadataApi.md#TestApiV1Metadata) | **Post** /api/v1/metadata/test | 
[**TestallApiV1Metadata**](MetadataApi.md#TestallApiV1Metadata) | **Post** /api/v1/metadata/testall | 
[**UpdateApiV1Metadata**](MetadataApi.md#UpdateApiV1Metadata) | **Put** /api/v1/metadata/{id} | 



## CreateApiV1Metadata

> MetadataResource CreateApiV1Metadata(ctx).MetadataResource(metadataResource).Execute()



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
    metadataResource := *readarrClient.NewMetadataResource() // MetadataResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.CreateApiV1Metadata(context.Background()).MetadataResource(metadataResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.CreateApiV1Metadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Metadata`: MetadataResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.CreateApiV1Metadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1MetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadataResource** | [**MetadataResource**](MetadataResource.md) |  | 

### Return type

[**MetadataResource**](MetadataResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiV1MetadataActionByName

> CreateApiV1MetadataActionByName(ctx, name).MetadataResource(metadataResource).Execute()



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
    name := "name_example" // string | 
    metadataResource := *readarrClient.NewMetadataResource() // MetadataResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.CreateApiV1MetadataActionByName(context.Background(), name).MetadataResource(metadataResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.CreateApiV1MetadataActionByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1MetadataActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataResource** | [**MetadataResource**](MetadataResource.md) |  | 

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


## DeleteApiV1Metadata

> DeleteApiV1Metadata(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataApi.DeleteApiV1Metadata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.DeleteApiV1Metadata``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1MetadataRequest struct via the builder pattern


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


## GetApiV1MetadataById

> MetadataResource GetApiV1MetadataById(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataApi.GetApiV1MetadataById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetApiV1MetadataById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1MetadataById`: MetadataResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetApiV1MetadataById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MetadataByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataResource**](MetadataResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1Metadata

> []MetadataResource ListApiV1Metadata(ctx).Execute()



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
    resp, r, err := apiClient.MetadataApi.ListApiV1Metadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.ListApiV1Metadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Metadata`: []MetadataResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.ListApiV1Metadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1MetadataRequest struct via the builder pattern


### Return type

[**[]MetadataResource**](MetadataResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1MetadataSchema

> []MetadataResource ListApiV1MetadataSchema(ctx).Execute()



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
    resp, r, err := apiClient.MetadataApi.ListApiV1MetadataSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.ListApiV1MetadataSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1MetadataSchema`: []MetadataResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.ListApiV1MetadataSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1MetadataSchemaRequest struct via the builder pattern


### Return type

[**[]MetadataResource**](MetadataResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestApiV1Metadata

> TestApiV1Metadata(ctx).MetadataResource(metadataResource).Execute()



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
    metadataResource := *readarrClient.NewMetadataResource() // MetadataResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.TestApiV1Metadata(context.Background()).MetadataResource(metadataResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.TestApiV1Metadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestApiV1MetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadataResource** | [**MetadataResource**](MetadataResource.md) |  | 

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


## TestallApiV1Metadata

> TestallApiV1Metadata(ctx).Execute()



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
    resp, r, err := apiClient.MetadataApi.TestallApiV1Metadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.TestallApiV1Metadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallApiV1MetadataRequest struct via the builder pattern


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


## UpdateApiV1Metadata

> MetadataResource UpdateApiV1Metadata(ctx, id).MetadataResource(metadataResource).Execute()



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
    metadataResource := *readarrClient.NewMetadataResource() // MetadataResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.UpdateApiV1Metadata(context.Background(), id).MetadataResource(metadataResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateApiV1Metadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Metadata`: MetadataResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.UpdateApiV1Metadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1MetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataResource** | [**MetadataResource**](MetadataResource.md) |  | 

### Return type

[**MetadataResource**](MetadataResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

