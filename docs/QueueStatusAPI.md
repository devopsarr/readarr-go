# \QueueStatusAPI

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueueStatus**](QueueStatusAPI.md#GetQueueStatus) | **Get** /api/v1/queue/status | 



## GetQueueStatus

> QueueStatusResource GetQueueStatus(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	readarrClient "github.com/devopsarr/readarr-go/readarr"
)

func main() {

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueStatusAPI.GetQueueStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueStatusAPI.GetQueueStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueueStatus`: QueueStatusResource
	fmt.Fprintf(os.Stdout, "Response from `QueueStatusAPI.GetQueueStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueStatusRequest struct via the builder pattern


### Return type

[**QueueStatusResource**](QueueStatusResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

