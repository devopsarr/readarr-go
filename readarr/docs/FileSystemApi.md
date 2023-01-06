# \FileSystemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1Filesystem**](FileSystemApi.md#GetApiV1Filesystem) | **Get** /api/v1/filesystem | 
[**GetApiV1FilesystemMediafiles**](FileSystemApi.md#GetApiV1FilesystemMediafiles) | **Get** /api/v1/filesystem/mediafiles | 
[**GetApiV1FilesystemType**](FileSystemApi.md#GetApiV1FilesystemType) | **Get** /api/v1/filesystem/type | 



## GetApiV1Filesystem

> GetApiV1Filesystem(ctx).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()



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
    resp, r, err := apiClient.FileSystemApi.GetApiV1Filesystem(context.Background()).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.GetApiV1Filesystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1FilesystemRequest struct via the builder pattern


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


## GetApiV1FilesystemMediafiles

> GetApiV1FilesystemMediafiles(ctx).Path(path).Execute()



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
    resp, r, err := apiClient.FileSystemApi.GetApiV1FilesystemMediafiles(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.GetApiV1FilesystemMediafiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1FilesystemMediafilesRequest struct via the builder pattern


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


## GetApiV1FilesystemType

> GetApiV1FilesystemType(ctx).Path(path).Execute()



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
    resp, r, err := apiClient.FileSystemApi.GetApiV1FilesystemType(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.GetApiV1FilesystemType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1FilesystemTypeRequest struct via the builder pattern


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

