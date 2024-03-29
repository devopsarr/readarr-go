# \CalendarFeedAPI

All URIs are relative to *http://localhost:8787*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedV1CalendarReadarrIcs**](CalendarFeedAPI.md#GetFeedV1CalendarReadarrIcs) | **Get** /feed/v1/calendar/readarr.ics | 



## GetFeedV1CalendarReadarrIcs

> GetFeedV1CalendarReadarrIcs(ctx).PastDays(pastDays).FutureDays(futureDays).TagList(tagList).Unmonitored(unmonitored).Execute()



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
	pastDays := int32(56) // int32 |  (optional) (default to 7)
	futureDays := int32(56) // int32 |  (optional) (default to 28)
	tagList := "tagList_example" // string |  (optional) (default to "")
	unmonitored := true // bool |  (optional) (default to false)

	configuration := readarrClient.NewConfiguration()
	apiClient := readarrClient.NewAPIClient(configuration)
	r, err := apiClient.CalendarFeedAPI.GetFeedV1CalendarReadarrIcs(context.Background()).PastDays(pastDays).FutureDays(futureDays).TagList(tagList).Unmonitored(unmonitored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarFeedAPI.GetFeedV1CalendarReadarrIcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedV1CalendarReadarrIcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pastDays** | **int32** |  | [default to 7]
 **futureDays** | **int32** |  | [default to 28]
 **tagList** | **string** |  | [default to &quot;&quot;]
 **unmonitored** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

