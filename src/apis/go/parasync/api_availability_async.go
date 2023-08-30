/*
OPERA Cloud Price Availability Rate Async API

APIs to cater for Price and Rate Availability Asynchronous functionality in OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package parasync

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AvailabilityAsyncApiService AvailabilityAsyncApi service
type AvailabilityAsyncApiService service

type AvailabilityAsyncApiGetRestrictionsRequest struct {
	ctx context.Context
	ApiService *AvailabilityAsyncApiService
	hotelId string
	extSystemCode string
	requestId string
	authorization *string
	xAppKey *string
	xHotelid *string
	acceptLanguage *string
}

// Bearer token that needs to be passed which is generated post user authentication
func (r AvailabilityAsyncApiGetRestrictionsRequest) Authorization(authorization string) AvailabilityAsyncApiGetRestrictionsRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r AvailabilityAsyncApiGetRestrictionsRequest) XAppKey(xAppKey string) AvailabilityAsyncApiGetRestrictionsRequest {
	r.xAppKey = &xAppKey
	return r
}

// Mandatory parameter to identify the hotel code where the end user is logged in
func (r AvailabilityAsyncApiGetRestrictionsRequest) XHotelid(xHotelid string) AvailabilityAsyncApiGetRestrictionsRequest {
	r.xHotelid = &xHotelid
	return r
}

// Language code
func (r AvailabilityAsyncApiGetRestrictionsRequest) AcceptLanguage(acceptLanguage string) AvailabilityAsyncApiGetRestrictionsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r AvailabilityAsyncApiGetRestrictionsRequest) Execute() (*RestrictionsStatus, *http.Response, error) {
	return r.ApiService.GetRestrictionsExecute(r)
}

/*
GetRestrictions Get status for created restrictions.

Use this to check whether the restrictions sent have been processed. <p><strong>OperationId:</strong>getRestrictions</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param hotelId Unique ID of Hotel in OPERA.
 @param extSystemCode Unique code of the external system.
 @param requestId Unique ID to fetch status of created restrictions.
 @return AvailabilityAsyncApiGetRestrictionsRequest
*/
func (a *AvailabilityAsyncApiService) GetRestrictions(ctx context.Context, hotelId string, extSystemCode string, requestId string) AvailabilityAsyncApiGetRestrictionsRequest {
	return AvailabilityAsyncApiGetRestrictionsRequest{
		ApiService: a,
		ctx: ctx,
		hotelId: hotelId,
		extSystemCode: extSystemCode,
		requestId: requestId,
	}
}

// Execute executes the request
//  @return RestrictionsStatus
func (a *AvailabilityAsyncApiService) GetRestrictionsExecute(r AvailabilityAsyncApiGetRestrictionsRequest) (*RestrictionsStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestrictionsStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityAsyncApiService.GetRestrictions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/externalSystems/{extSystemCode}/hotels/{hotelId}/restrictions/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"hotelId"+"}", url.PathEscape(parameterValueToString(r.hotelId, "hotelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extSystemCode"+"}", url.PathEscape(parameterValueToString(r.extSystemCode, "extSystemCode")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.hotelId) < 1 {
		return localVarReturnValue, nil, reportError("hotelId must have at least 1 elements")
	}
	if strlen(r.hotelId) > 2000 {
		return localVarReturnValue, nil, reportError("hotelId must have less than 2000 elements")
	}
	if strlen(r.extSystemCode) < 1 {
		return localVarReturnValue, nil, reportError("extSystemCode must have at least 1 elements")
	}
	if strlen(r.extSystemCode) > 2000 {
		return localVarReturnValue, nil, reportError("extSystemCode must have less than 2000 elements")
	}
	if strlen(r.requestId) < 1 {
		return localVarReturnValue, nil, reportError("requestId must have at least 1 elements")
	}
	if strlen(r.requestId) > 2000 {
		return localVarReturnValue, nil, reportError("requestId must have less than 2000 elements")
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

type AvailabilityAsyncApiGetRestrictionsProcessStatusRequest struct {
	ctx context.Context
	ApiService *AvailabilityAsyncApiService
	requestId string
	extSystemCode string
	hotelId string
	authorization *string
	xAppKey *string
	xHotelid *string
	acceptLanguage *string
}

// Bearer token that needs to be passed which is generated post user authentication
func (r AvailabilityAsyncApiGetRestrictionsProcessStatusRequest) Authorization(authorization string) AvailabilityAsyncApiGetRestrictionsProcessStatusRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r AvailabilityAsyncApiGetRestrictionsProcessStatusRequest) XAppKey(xAppKey string) AvailabilityAsyncApiGetRestrictionsProcessStatusRequest {
	r.xAppKey = &xAppKey
	return r
}

// Mandatory parameter to identify the hotel code where the end user is logged in
func (r AvailabilityAsyncApiGetRestrictionsProcessStatusRequest) XHotelid(xHotelid string) AvailabilityAsyncApiGetRestrictionsProcessStatusRequest {
	r.xHotelid = &xHotelid
	return r
}

// Language code
func (r AvailabilityAsyncApiGetRestrictionsProcessStatusRequest) AcceptLanguage(acceptLanguage string) AvailabilityAsyncApiGetRestrictionsProcessStatusRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r AvailabilityAsyncApiGetRestrictionsProcessStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetRestrictionsProcessStatusExecute(r)
}

/*
GetRestrictionsProcessStatus Check status of Restrictions

Use this to check whether the restrictions sent have been processed. <p><strong>OperationId:</strong>getRestrictionsProcessStatus</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param requestId Unique ID to fetch status of created restrictions.
 @param extSystemCode Unique code of the external system.
 @param hotelId Unique ID of Hotel in OPERA.
 @return AvailabilityAsyncApiGetRestrictionsProcessStatusRequest
*/
func (a *AvailabilityAsyncApiService) GetRestrictionsProcessStatus(ctx context.Context, requestId string, extSystemCode string, hotelId string) AvailabilityAsyncApiGetRestrictionsProcessStatusRequest {
	return AvailabilityAsyncApiGetRestrictionsProcessStatusRequest{
		ApiService: a,
		ctx: ctx,
		requestId: requestId,
		extSystemCode: extSystemCode,
		hotelId: hotelId,
	}
}

// Execute executes the request
func (a *AvailabilityAsyncApiService) GetRestrictionsProcessStatusExecute(r AvailabilityAsyncApiGetRestrictionsProcessStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodHead
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityAsyncApiService.GetRestrictionsProcessStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/externalSystems/{extSystemCode}/hotels/{hotelId}/restrictions/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extSystemCode"+"}", url.PathEscape(parameterValueToString(r.extSystemCode, "extSystemCode")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hotelId"+"}", url.PathEscape(parameterValueToString(r.hotelId, "hotelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.requestId) < 1 {
		return nil, reportError("requestId must have at least 1 elements")
	}
	if strlen(r.requestId) > 2000 {
		return nil, reportError("requestId must have less than 2000 elements")
	}
	if strlen(r.extSystemCode) < 1 {
		return nil, reportError("extSystemCode must have at least 1 elements")
	}
	if strlen(r.extSystemCode) > 2000 {
		return nil, reportError("extSystemCode must have less than 2000 elements")
	}
	if strlen(r.hotelId) < 1 {
		return nil, reportError("hotelId must have at least 1 elements")
	}
	if strlen(r.hotelId) > 2000 {
		return nil, reportError("hotelId must have less than 2000 elements")
	}

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
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
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

type AvailabilityAsyncApiPostRestrictionsProcessRequest struct {
	ctx context.Context
	ApiService *AvailabilityAsyncApiService
	hotelId string
	extSystemCode string
	authorization *string
	xAppKey *string
	xHotelid *string
	restrictions *Restrictions
	acceptLanguage *string
}

// Bearer token that needs to be passed which is generated post user authentication
func (r AvailabilityAsyncApiPostRestrictionsProcessRequest) Authorization(authorization string) AvailabilityAsyncApiPostRestrictionsProcessRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r AvailabilityAsyncApiPostRestrictionsProcessRequest) XAppKey(xAppKey string) AvailabilityAsyncApiPostRestrictionsProcessRequest {
	r.xAppKey = &xAppKey
	return r
}

// Mandatory parameter to identify the hotel code where the end user is logged in
func (r AvailabilityAsyncApiPostRestrictionsProcessRequest) XHotelid(xHotelid string) AvailabilityAsyncApiPostRestrictionsProcessRequest {
	r.xHotelid = &xHotelid
	return r
}

// Request for configuring restrictions.
func (r AvailabilityAsyncApiPostRestrictionsProcessRequest) Restrictions(restrictions Restrictions) AvailabilityAsyncApiPostRestrictionsProcessRequest {
	r.restrictions = &restrictions
	return r
}

// Language code
func (r AvailabilityAsyncApiPostRestrictionsProcessRequest) AcceptLanguage(acceptLanguage string) AvailabilityAsyncApiPostRestrictionsProcessRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r AvailabilityAsyncApiPostRestrictionsProcessRequest) Execute() ([]InstanceLink, *http.Response, error) {
	return r.ApiService.PostRestrictionsProcessExecute(r)
}

/*
PostRestrictionsProcess Create restrictions in OPERA Cloud.

A user can send various restrictions to OPERA by specifying restriction details in the request. <p><strong>OperationId:</strong>postRestrictionsProcess</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param hotelId Unique ID of the hotel in OPERA
 @param extSystemCode Unique code of the external system
 @return AvailabilityAsyncApiPostRestrictionsProcessRequest
*/
func (a *AvailabilityAsyncApiService) PostRestrictionsProcess(ctx context.Context, hotelId string, extSystemCode string) AvailabilityAsyncApiPostRestrictionsProcessRequest {
	return AvailabilityAsyncApiPostRestrictionsProcessRequest{
		ApiService: a,
		ctx: ctx,
		hotelId: hotelId,
		extSystemCode: extSystemCode,
	}
}

// Execute executes the request
//  @return []InstanceLink
func (a *AvailabilityAsyncApiService) PostRestrictionsProcessExecute(r AvailabilityAsyncApiPostRestrictionsProcessRequest) ([]InstanceLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InstanceLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityAsyncApiService.PostRestrictionsProcess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/externalSystems/{extSystemCode}/hotels/{hotelId}/restrictions"
	localVarPath = strings.Replace(localVarPath, "{"+"hotelId"+"}", url.PathEscape(parameterValueToString(r.hotelId, "hotelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extSystemCode"+"}", url.PathEscape(parameterValueToString(r.extSystemCode, "extSystemCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.hotelId) < 1 {
		return localVarReturnValue, nil, reportError("hotelId must have at least 1 elements")
	}
	if strlen(r.hotelId) > 2000 {
		return localVarReturnValue, nil, reportError("hotelId must have less than 2000 elements")
	}
	if strlen(r.extSystemCode) < 1 {
		return localVarReturnValue, nil, reportError("extSystemCode must have at least 1 elements")
	}
	if strlen(r.extSystemCode) > 2000 {
		return localVarReturnValue, nil, reportError("extSystemCode must have less than 2000 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8"}

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
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	// body params
	localVarPostBody = r.restrictions
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
