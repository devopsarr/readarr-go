/*
Readarr

Readarr API docs

API version: v0.4.10.2734
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
	"time"
)


// HistoryAPIService HistoryAPI service
type HistoryAPIService service

type ApiCreateHistoryFailedByIdRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	id int32
}

func (r ApiCreateHistoryFailedByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateHistoryFailedByIdExecute(r)
}

/*
CreateHistoryFailedById Method for CreateHistoryFailedById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCreateHistoryFailedByIdRequest
*/
func (a *HistoryAPIService) CreateHistoryFailedById(ctx context.Context, id int32) ApiCreateHistoryFailedByIdRequest {
	return ApiCreateHistoryFailedByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *HistoryAPIService) CreateHistoryFailedByIdExecute(r ApiCreateHistoryFailedByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.CreateHistoryFailedById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/history/failed/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetHistoryRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	page *int32
	pageSize *int32
	sortKey *string
	sortDirection *SortDirection
	includeAuthor *bool
	includeBook *bool
	eventType *[]int32
	bookId *int32
	downloadId *string
}

func (r ApiGetHistoryRequest) Page(page int32) ApiGetHistoryRequest {
	r.page = &page
	return r
}

func (r ApiGetHistoryRequest) PageSize(pageSize int32) ApiGetHistoryRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetHistoryRequest) SortKey(sortKey string) ApiGetHistoryRequest {
	r.sortKey = &sortKey
	return r
}

func (r ApiGetHistoryRequest) SortDirection(sortDirection SortDirection) ApiGetHistoryRequest {
	r.sortDirection = &sortDirection
	return r
}

func (r ApiGetHistoryRequest) IncludeAuthor(includeAuthor bool) ApiGetHistoryRequest {
	r.includeAuthor = &includeAuthor
	return r
}

func (r ApiGetHistoryRequest) IncludeBook(includeBook bool) ApiGetHistoryRequest {
	r.includeBook = &includeBook
	return r
}

func (r ApiGetHistoryRequest) EventType(eventType []int32) ApiGetHistoryRequest {
	r.eventType = &eventType
	return r
}

func (r ApiGetHistoryRequest) BookId(bookId int32) ApiGetHistoryRequest {
	r.bookId = &bookId
	return r
}

func (r ApiGetHistoryRequest) DownloadId(downloadId string) ApiGetHistoryRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetHistoryRequest) Execute() (*HistoryResourcePagingResource, *http.Response, error) {
	return r.ApiService.GetHistoryExecute(r)
}

/*
GetHistory Method for GetHistory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHistoryRequest
*/
func (a *HistoryAPIService) GetHistory(ctx context.Context) ApiGetHistoryRequest {
	return ApiGetHistoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistoryResourcePagingResource
func (a *HistoryAPIService) GetHistoryExecute(r ApiGetHistoryRequest) (*HistoryResourcePagingResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistoryResourcePagingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.GetHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 10
		r.pageSize = &defaultValue
	}
	if r.sortKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortKey", r.sortKey, "form", "")
	}
	if r.sortDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortDirection", r.sortDirection, "form", "")
	}
	if r.includeAuthor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeAuthor", r.includeAuthor, "form", "")
	}
	if r.includeBook != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeBook", r.includeBook, "form", "")
	}
	if r.eventType != nil {
		t := *r.eventType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", t, "form", "multi")
		}
	}
	if r.bookId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bookId", r.bookId, "form", "")
	}
	if r.downloadId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListHistoryAuthorRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	authorId *int32
	bookId *int32
	eventType *EntityHistoryEventType
	includeAuthor *bool
	includeBook *bool
}

func (r ApiListHistoryAuthorRequest) AuthorId(authorId int32) ApiListHistoryAuthorRequest {
	r.authorId = &authorId
	return r
}

func (r ApiListHistoryAuthorRequest) BookId(bookId int32) ApiListHistoryAuthorRequest {
	r.bookId = &bookId
	return r
}

func (r ApiListHistoryAuthorRequest) EventType(eventType EntityHistoryEventType) ApiListHistoryAuthorRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListHistoryAuthorRequest) IncludeAuthor(includeAuthor bool) ApiListHistoryAuthorRequest {
	r.includeAuthor = &includeAuthor
	return r
}

func (r ApiListHistoryAuthorRequest) IncludeBook(includeBook bool) ApiListHistoryAuthorRequest {
	r.includeBook = &includeBook
	return r
}

func (r ApiListHistoryAuthorRequest) Execute() ([]HistoryResource, *http.Response, error) {
	return r.ApiService.ListHistoryAuthorExecute(r)
}

/*
ListHistoryAuthor Method for ListHistoryAuthor

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHistoryAuthorRequest
*/
func (a *HistoryAPIService) ListHistoryAuthor(ctx context.Context) ApiListHistoryAuthorRequest {
	return ApiListHistoryAuthorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []HistoryResource
func (a *HistoryAPIService) ListHistoryAuthorExecute(r ApiListHistoryAuthorRequest) ([]HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []HistoryResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.ListHistoryAuthor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/history/author"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.authorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "authorId", r.authorId, "form", "")
	}
	if r.bookId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bookId", r.bookId, "form", "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "form", "")
	}
	if r.includeAuthor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeAuthor", r.includeAuthor, "form", "")
	} else {
		var defaultValue bool = false
		r.includeAuthor = &defaultValue
	}
	if r.includeBook != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeBook", r.includeBook, "form", "")
	} else {
		var defaultValue bool = false
		r.includeBook = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListHistorySinceRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	date *time.Time
	eventType *EntityHistoryEventType
	includeAuthor *bool
	includeBook *bool
}

func (r ApiListHistorySinceRequest) Date(date time.Time) ApiListHistorySinceRequest {
	r.date = &date
	return r
}

func (r ApiListHistorySinceRequest) EventType(eventType EntityHistoryEventType) ApiListHistorySinceRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListHistorySinceRequest) IncludeAuthor(includeAuthor bool) ApiListHistorySinceRequest {
	r.includeAuthor = &includeAuthor
	return r
}

func (r ApiListHistorySinceRequest) IncludeBook(includeBook bool) ApiListHistorySinceRequest {
	r.includeBook = &includeBook
	return r
}

func (r ApiListHistorySinceRequest) Execute() ([]HistoryResource, *http.Response, error) {
	return r.ApiService.ListHistorySinceExecute(r)
}

/*
ListHistorySince Method for ListHistorySince

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHistorySinceRequest
*/
func (a *HistoryAPIService) ListHistorySince(ctx context.Context) ApiListHistorySinceRequest {
	return ApiListHistorySinceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []HistoryResource
func (a *HistoryAPIService) ListHistorySinceExecute(r ApiListHistorySinceRequest) ([]HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []HistoryResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.ListHistorySince")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/history/since"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "form", "")
	}
	if r.includeAuthor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeAuthor", r.includeAuthor, "form", "")
	} else {
		var defaultValue bool = false
		r.includeAuthor = &defaultValue
	}
	if r.includeBook != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeBook", r.includeBook, "form", "")
	} else {
		var defaultValue bool = false
		r.includeBook = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
