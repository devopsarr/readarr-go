# \ImportListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1ImportList**](ImportListApi.md#CreateApiV1ImportList) | **Post** /api/v1/importlist | 
[**CreateApiV1ImportListActionByName**](ImportListApi.md#CreateApiV1ImportListActionByName) | **Post** /api/v1/importlist/action/{name} | 
[**DeleteApiV1ImportList**](ImportListApi.md#DeleteApiV1ImportList) | **Delete** /api/v1/importlist/{id} | 
[**GetApiV1ImportListById**](ImportListApi.md#GetApiV1ImportListById) | **Get** /api/v1/importlist/{id} | 
[**ListApiV1ImportList**](ImportListApi.md#ListApiV1ImportList) | **Get** /api/v1/importlist | 
[**ListApiV1ImportListSchema**](ImportListApi.md#ListApiV1ImportListSchema) | **Get** /api/v1/importlist/schema | 
[**TestApiV1ImportList**](ImportListApi.md#TestApiV1ImportList) | **Post** /api/v1/importlist/test | 
[**TestallApiV1ImportList**](ImportListApi.md#TestallApiV1ImportList) | **Post** /api/v1/importlist/testall | 
[**UpdateApiV1ImportList**](ImportListApi.md#UpdateApiV1ImportList) | **Put** /api/v1/importlist/{id} | 



## CreateApiV1ImportList

> ImportListResource CreateApiV1ImportList(ctx).ImportListResource(importListResource).Execute()



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
    importListResource := *readarrClient.NewImportListResource() // ImportListResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListApi.CreateApiV1ImportList(context.Background()).ImportListResource(importListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.CreateApiV1ImportList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1ImportList`: ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.CreateApiV1ImportList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ImportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListResource** | [**ImportListResource**](ImportListResource.md) |  | 

### Return type

[**ImportListResource**](ImportListResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiV1ImportListActionByName

> CreateApiV1ImportListActionByName(ctx, name).ImportListResource(importListResource).Execute()



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
    importListResource := *readarrClient.NewImportListResource() // ImportListResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListApi.CreateApiV1ImportListActionByName(context.Background(), name).ImportListResource(importListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.CreateApiV1ImportListActionByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateApiV1ImportListActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importListResource** | [**ImportListResource**](ImportListResource.md) |  | 

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


## DeleteApiV1ImportList

> DeleteApiV1ImportList(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListApi.DeleteApiV1ImportList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.DeleteApiV1ImportList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1ImportListRequest struct via the builder pattern


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


## GetApiV1ImportListById

> ImportListResource GetApiV1ImportListById(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListApi.GetApiV1ImportListById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.GetApiV1ImportListById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ImportListById`: ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.GetApiV1ImportListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ImportListByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportListResource**](ImportListResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1ImportList

> []ImportListResource ListApiV1ImportList(ctx).Execute()



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
    resp, r, err := apiClient.ImportListApi.ListApiV1ImportList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.ListApiV1ImportList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1ImportList`: []ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.ListApiV1ImportList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ImportListRequest struct via the builder pattern


### Return type

[**[]ImportListResource**](ImportListResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1ImportListSchema

> []ImportListResource ListApiV1ImportListSchema(ctx).Execute()



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
    resp, r, err := apiClient.ImportListApi.ListApiV1ImportListSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.ListApiV1ImportListSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1ImportListSchema`: []ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.ListApiV1ImportListSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ImportListSchemaRequest struct via the builder pattern


### Return type

[**[]ImportListResource**](ImportListResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestApiV1ImportList

> TestApiV1ImportList(ctx).ImportListResource(importListResource).Execute()



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
    importListResource := *readarrClient.NewImportListResource() // ImportListResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListApi.TestApiV1ImportList(context.Background()).ImportListResource(importListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.TestApiV1ImportList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestApiV1ImportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListResource** | [**ImportListResource**](ImportListResource.md) |  | 

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


## TestallApiV1ImportList

> TestallApiV1ImportList(ctx).Execute()



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
    resp, r, err := apiClient.ImportListApi.TestallApiV1ImportList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.TestallApiV1ImportList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallApiV1ImportListRequest struct via the builder pattern


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


## UpdateApiV1ImportList

> ImportListResource UpdateApiV1ImportList(ctx, id).ImportListResource(importListResource).Execute()



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
    importListResource := *readarrClient.NewImportListResource() // ImportListResource |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListApi.UpdateApiV1ImportList(context.Background(), id).ImportListResource(importListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.UpdateApiV1ImportList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ImportList`: ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.UpdateApiV1ImportList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ImportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importListResource** | [**ImportListResource**](ImportListResource.md) |  | 

### Return type

[**ImportListResource**](ImportListResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

