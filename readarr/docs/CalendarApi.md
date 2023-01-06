# \CalendarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1CalendarById**](CalendarApi.md#GetApiV1CalendarById) | **Get** /api/v1/calendar/{id} | 
[**ListApiV1Calendar**](CalendarApi.md#ListApiV1Calendar) | **Get** /api/v1/calendar | 



## GetApiV1CalendarById

> BookResource GetApiV1CalendarById(ctx, id).Execute()



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
    resp, r, err := apiClient.CalendarApi.GetApiV1CalendarById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.GetApiV1CalendarById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1CalendarById`: BookResource
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.GetApiV1CalendarById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1CalendarByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookResource**](BookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1Calendar

> []BookResource ListApiV1Calendar(ctx).Start(start).End(end).Unmonitored(unmonitored).IncludeAuthor(includeAuthor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    readarrClient "./openapi"
)

func main() {
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    unmonitored := true // bool |  (optional) (default to false)
    includeAuthor := true // bool |  (optional) (default to false)

    configuration := readarrClient.NewConfiguration()
    apiClient := readarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalendarApi.ListApiV1Calendar(context.Background()).Start(start).End(end).Unmonitored(unmonitored).IncludeAuthor(includeAuthor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.ListApiV1Calendar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Calendar`: []BookResource
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.ListApiV1Calendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1CalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 
 **unmonitored** | **bool** |  | [default to false]
 **includeAuthor** | **bool** |  | [default to false]

### Return type

[**[]BookResource**](BookResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

