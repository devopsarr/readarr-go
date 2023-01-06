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
)


// BookshelfApiService BookshelfApi service
type BookshelfApiService service
type ApiCreateApiV1BookshelfRequest struct {
	ctx context.Context
	ApiService *BookshelfApiService
	bookshelfResource *BookshelfResource
}

func (r ApiCreateApiV1BookshelfRequest) BookshelfResource(bookshelfResource BookshelfResource) ApiCreateApiV1BookshelfRequest {
	r.bookshelfResource = &bookshelfResource
	return r
}

func (r ApiCreateApiV1BookshelfRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateApiV1BookshelfExecute(r)
}

/*
CreateApiV1Bookshelf Method for CreateApiV1Bookshelf

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateApiV1BookshelfRequest
*/
func (a *BookshelfApiService) CreateApiV1Bookshelf(ctx context.Context) ApiCreateApiV1BookshelfRequest {
	return ApiCreateApiV1BookshelfRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *BookshelfApiService) CreateApiV1BookshelfExecute(r ApiCreateApiV1BookshelfRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BookshelfApiService.CreateApiV1Bookshelf")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/bookshelf"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

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
	// body params
	localVarPostBody = r.bookshelfResource
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
