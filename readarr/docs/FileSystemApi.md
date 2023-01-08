# \FileSystemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1FileSystem**](FileSystemApi.md#GetApiV1FileSystem) | **Get** /api/v1/filesystem | 
[**GetApiV1FileSystemMediafiles**](FileSystemApi.md#GetApiV1FileSystemMediafiles) | **Get** /api/v1/filesystem/mediafiles | 
[**GetApiV1FileSystemType**](FileSystemApi.md#GetApiV1FileSystemType) | **Get** /api/v1/filesystem/type | 



## GetApiV1FileSystem

> GetApiV1FileSystem(ctx).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()



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
    path := "path_example" // string |  (optional)
    includeFiles := true // bool |  (optional) (default to false)
    allowFoldersWithoutTrailingSlashes := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemApi.GetApiV1FileSystem(context.Background()).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.GetApiV1FileSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1FileSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **includeFiles** | **bool** |  | [default to false]
 **allowFoldersWithoutTrailingSlashes** | **bool** |  | [default to false]

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


## GetApiV1FileSystemMediafiles

> GetApiV1FileSystemMediafiles(ctx).Path(path).Execute()



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
    path := "path_example" // string |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemApi.GetApiV1FileSystemMediafiles(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.GetApiV1FileSystemMediafiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1FileSystemMediafilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

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


## GetApiV1FileSystemType

> GetApiV1FileSystemType(ctx).Path(path).Execute()



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
    path := "path_example" // string |  (optional)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemApi.GetApiV1FileSystemType(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.GetApiV1FileSystemType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1FileSystemTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

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

