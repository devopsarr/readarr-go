# \ImportListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Importlist**](ImportListApi.md#CreateApiV1Importlist) | **Post** /api/v1/importlist | 
[**CreateApiV1ImportlistActionByName**](ImportListApi.md#CreateApiV1ImportlistActionByName) | **Post** /api/v1/importlist/action/{name} | 
[**DeleteApiV1Importlist**](ImportListApi.md#DeleteApiV1Importlist) | **Delete** /api/v1/importlist/{id} | 
[**GetApiV1ImportlistById**](ImportListApi.md#GetApiV1ImportlistById) | **Get** /api/v1/importlist/{id} | 
[**ListApiV1Importlist**](ImportListApi.md#ListApiV1Importlist) | **Get** /api/v1/importlist | 
[**ListApiV1ImportlistSchema**](ImportListApi.md#ListApiV1ImportlistSchema) | **Get** /api/v1/importlist/schema | 
[**TestApiV1Importlist**](ImportListApi.md#TestApiV1Importlist) | **Post** /api/v1/importlist/test | 
[**TestallApiV1Importlist**](ImportListApi.md#TestallApiV1Importlist) | **Post** /api/v1/importlist/testall | 
[**UpdateApiV1Importlist**](ImportListApi.md#UpdateApiV1Importlist) | **Put** /api/v1/importlist/{id} | 



## CreateApiV1Importlist

> ImportListResource CreateApiV1Importlist(ctx).ImportListResource(importListResource).Execute()



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
    resp, r, err := apiClient.ImportListApi.CreateApiV1Importlist(context.Background()).ImportListResource(importListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.CreateApiV1Importlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Importlist`: ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.CreateApiV1Importlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ImportlistRequest struct via the builder pattern


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


## CreateApiV1ImportlistActionByName

> CreateApiV1ImportlistActionByName(ctx, name).ImportListResource(importListResource).Execute()



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
    resp, r, err := apiClient.ImportListApi.CreateApiV1ImportlistActionByName(context.Background(), name).ImportListResource(importListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.CreateApiV1ImportlistActionByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateApiV1ImportlistActionByNameRequest struct via the builder pattern


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


## DeleteApiV1Importlist

> DeleteApiV1Importlist(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListApi.DeleteApiV1Importlist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.DeleteApiV1Importlist``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1ImportlistRequest struct via the builder pattern


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


## GetApiV1ImportlistById

> ImportListResource GetApiV1ImportlistById(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListApi.GetApiV1ImportlistById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.GetApiV1ImportlistById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ImportlistById`: ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.GetApiV1ImportlistById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ImportlistByIdRequest struct via the builder pattern


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


## ListApiV1Importlist

> []ImportListResource ListApiV1Importlist(ctx).Execute()



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
    resp, r, err := apiClient.ImportListApi.ListApiV1Importlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.ListApiV1Importlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Importlist`: []ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.ListApiV1Importlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ImportlistRequest struct via the builder pattern


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


## ListApiV1ImportlistSchema

> []ImportListResource ListApiV1ImportlistSchema(ctx).Execute()



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
    resp, r, err := apiClient.ImportListApi.ListApiV1ImportlistSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.ListApiV1ImportlistSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1ImportlistSchema`: []ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.ListApiV1ImportlistSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ImportlistSchemaRequest struct via the builder pattern


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


## TestApiV1Importlist

> TestApiV1Importlist(ctx).ImportListResource(importListResource).Execute()



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
    resp, r, err := apiClient.ImportListApi.TestApiV1Importlist(context.Background()).ImportListResource(importListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.TestApiV1Importlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestApiV1ImportlistRequest struct via the builder pattern


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


## TestallApiV1Importlist

> TestallApiV1Importlist(ctx).Execute()



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
    resp, r, err := apiClient.ImportListApi.TestallApiV1Importlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.TestallApiV1Importlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallApiV1ImportlistRequest struct via the builder pattern


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


## UpdateApiV1Importlist

> ImportListResource UpdateApiV1Importlist(ctx, id).ImportListResource(importListResource).Execute()



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
    resp, r, err := apiClient.ImportListApi.UpdateApiV1Importlist(context.Background(), id).ImportListResource(importListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListApi.UpdateApiV1Importlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Importlist`: ImportListResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListApi.UpdateApiV1Importlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ImportlistRequest struct via the builder pattern


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

