/*
Readarr

Readarr API docs

API version: 0.1.0
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
)


// MediaCoverApiService MediaCoverApi service
type MediaCoverApiService service
type ApiGetMediaCoverAuthorauthorIdByFilenameRequest struct {
	ctx context.Context
	ApiService *MediaCoverApiService
	authorId int32
	filename string
}

func (r ApiGetMediaCoverAuthorauthorIdByFilenameRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetMediaCoverAuthorauthorIdByFilenameExecute(r)
}

/*
GetMediaCoverAuthorauthorIdByFilename Method for GetMediaCoverAuthorauthorIdByFilename

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authorId
 @param filename
 @return ApiGetMediaCoverAuthorauthorIdByFilenameRequest
*/
func (a *MediaCoverApiService) GetMediaCoverAuthorauthorIdByFilename(ctx context.Context, authorId int32, filename string) ApiGetMediaCoverAuthorauthorIdByFilenameRequest {
	return ApiGetMediaCoverAuthorauthorIdByFilenameRequest{
		ApiService: a,
		ctx: ctx,
		authorId: authorId,
		filename: filename,
	}
}

// Execute executes the request
func (a *MediaCoverApiService) GetMediaCoverAuthorauthorIdByFilenameExecute(r ApiGetMediaCoverAuthorauthorIdByFilenameRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaCoverApiService.GetMediaCoverAuthorauthorIdByFilename")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/mediacover/author/{authorId}/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"authorId"+"}", url.PathEscape(parameterToString(r.authorId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterToString(r.filename, "")), -1)

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
type ApiGetMediaCoverBookbookIdByFilenameRequest struct {
	ctx context.Context
	ApiService *MediaCoverApiService
	bookId int32
	filename string
}

func (r ApiGetMediaCoverBookbookIdByFilenameRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetMediaCoverBookbookIdByFilenameExecute(r)
}

/*
GetMediaCoverBookbookIdByFilename Method for GetMediaCoverBookbookIdByFilename

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bookId
 @param filename
 @return ApiGetMediaCoverBookbookIdByFilenameRequest
*/
func (a *MediaCoverApiService) GetMediaCoverBookbookIdByFilename(ctx context.Context, bookId int32, filename string) ApiGetMediaCoverBookbookIdByFilenameRequest {
	return ApiGetMediaCoverBookbookIdByFilenameRequest{
		ApiService: a,
		ctx: ctx,
		bookId: bookId,
		filename: filename,
	}
}

// Execute executes the request
func (a *MediaCoverApiService) GetMediaCoverBookbookIdByFilenameExecute(r ApiGetMediaCoverBookbookIdByFilenameRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaCoverApiService.GetMediaCoverBookbookIdByFilename")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/mediacover/book/{bookId}/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"bookId"+"}", url.PathEscape(parameterToString(r.bookId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterToString(r.filename, "")), -1)

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
