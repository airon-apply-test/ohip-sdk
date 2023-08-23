/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// BEProcessorApiService BEProcessorApi service
type BEProcessorApiService service

type ApiClearCacheBeProcessorRequest struct {
	ctx context.Context
	ApiService *BEProcessorApiService
	authorization *string
	xAppKey *string
	xHotelid *string
	xExternalsystem *string
	acceptLanguage *string
}

// Bearer token that needs to be passed which is generated post user authentication
func (r ApiClearCacheBeProcessorRequest) Authorization(authorization string) ApiClearCacheBeProcessorRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r ApiClearCacheBeProcessorRequest) XAppKey(xAppKey string) ApiClearCacheBeProcessorRequest {
	r.xAppKey = &xAppKey
	return r
}

// Mandatory parameter to identify the hotel code where the end user is logged in
func (r ApiClearCacheBeProcessorRequest) XHotelid(xHotelid string) ApiClearCacheBeProcessorRequest {
	r.xHotelid = &xHotelid
	return r
}

// External system code.
func (r ApiClearCacheBeProcessorRequest) XExternalsystem(xExternalsystem string) ApiClearCacheBeProcessorRequest {
	r.xExternalsystem = &xExternalsystem
	return r
}

// Language code
func (r ApiClearCacheBeProcessorRequest) AcceptLanguage(acceptLanguage string) ApiClearCacheBeProcessorRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiClearCacheBeProcessorRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.ClearCacheBeProcessorExecute(r)
}

/*
ClearCacheBeProcessor Operation to clear cache.

 <p><strong>OperationId:</strong>clearCacheBeProcessor</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiClearCacheBeProcessorRequest
*/
func (a *BEProcessorApiService) ClearCacheBeProcessor(ctx context.Context) ApiClearCacheBeProcessorRequest {
	return ApiClearCacheBeProcessorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Status
func (a *BEProcessorApiService) ClearCacheBeProcessorExecute(r ApiClearCacheBeProcessorRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BEProcessorApiService.ClearCacheBeProcessor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/config/oxi/services/beProcessor/cache"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.xAppKey == nil {
		return localVarReturnValue, nil, reportError("xAppKey is required and must be specified")
	}
	if r.xHotelid == nil {
		return localVarReturnValue, nil, reportError("xHotelid is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-app-key", r.xAppKey, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-hotelid", r.xHotelid, "")
	if r.xExternalsystem != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-externalsystem", r.xExternalsystem, "")
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExceptionDetailType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ExceptionDetailType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ExceptionDetailType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiPingBeProcessorRequest struct {
	ctx context.Context
	ApiService *BEProcessorApiService
	authorization *string
	xAppKey *string
	xHotelid *string
	xExternalsystem *string
	acceptLanguage *string
}

// Bearer token that needs to be passed which is generated post user authentication
func (r ApiPingBeProcessorRequest) Authorization(authorization string) ApiPingBeProcessorRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r ApiPingBeProcessorRequest) XAppKey(xAppKey string) ApiPingBeProcessorRequest {
	r.xAppKey = &xAppKey
	return r
}

// Mandatory parameter to identify the hotel code where the end user is logged in
func (r ApiPingBeProcessorRequest) XHotelid(xHotelid string) ApiPingBeProcessorRequest {
	r.xHotelid = &xHotelid
	return r
}

// External system code.
func (r ApiPingBeProcessorRequest) XExternalsystem(xExternalsystem string) ApiPingBeProcessorRequest {
	r.xExternalsystem = &xExternalsystem
	return r
}

// Language code
func (r ApiPingBeProcessorRequest) AcceptLanguage(acceptLanguage string) ApiPingBeProcessorRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiPingBeProcessorRequest) Execute() (*OperaVersion, *http.Response, error) {
	return r.ApiService.PingBeProcessorExecute(r)
}

/*
PingBeProcessor Operation to ping.

 <p><strong>OperationId:</strong>pingBeProcessor</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPingBeProcessorRequest
*/
func (a *BEProcessorApiService) PingBeProcessor(ctx context.Context) ApiPingBeProcessorRequest {
	return ApiPingBeProcessorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OperaVersion
func (a *BEProcessorApiService) PingBeProcessorExecute(r ApiPingBeProcessorRequest) (*OperaVersion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OperaVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BEProcessorApiService.PingBeProcessor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/config/oxi/services/beProcessor/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.xAppKey == nil {
		return localVarReturnValue, nil, reportError("xAppKey is required and must be specified")
	}
	if r.xHotelid == nil {
		return localVarReturnValue, nil, reportError("xHotelid is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-app-key", r.xAppKey, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-hotelid", r.xHotelid, "")
	if r.xExternalsystem != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-externalsystem", r.xExternalsystem, "")
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExceptionDetailType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ExceptionDetailType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ExceptionDetailType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
