/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

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
	"strings"
)


// ProfileApiService ProfileApi service
type ProfileApiService service

type ApiPostMembershipNumberRequest struct {
	ctx context.Context
	ApiService *ProfileApiService
	extSystemCode string
	membershipType string
	authorization *string
	xAppKey *string
	membershipNumber *PostMembershipNumberRequest
	xTransactionId *string
	acceptLanguage *string
}

// Bearer token that needs to be passed which is generated post user authentication
func (r ApiPostMembershipNumberRequest) Authorization(authorization string) ApiPostMembershipNumberRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r ApiPostMembershipNumberRequest) XAppKey(xAppKey string) ApiPostMembershipNumberRequest {
	r.xAppKey = &xAppKey
	return r
}

// Request object for generating membership number.
func (r ApiPostMembershipNumberRequest) MembershipNumber(membershipNumber PostMembershipNumberRequest) ApiPostMembershipNumberRequest {
	r.membershipNumber = &membershipNumber
	return r
}

// Transaction Id
func (r ApiPostMembershipNumberRequest) XTransactionId(xTransactionId string) ApiPostMembershipNumberRequest {
	r.xTransactionId = &xTransactionId
	return r
}

// Language code
func (r ApiPostMembershipNumberRequest) AcceptLanguage(acceptLanguage string) ApiPostMembershipNumberRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiPostMembershipNumberRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.PostMembershipNumberExecute(r)
}

/*
PostMembershipNumber Generate membership number

This API will generate a new membership number. <p><strong>OperationId:</strong>postMembershipNumber</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param extSystemCode Profile will be downloaded from this External System.
 @param membershipType
 @return ApiPostMembershipNumberRequest
*/
func (a *ProfileApiService) PostMembershipNumber(ctx context.Context, extSystemCode string, membershipType string) ApiPostMembershipNumberRequest {
	return ApiPostMembershipNumberRequest{
		ApiService: a,
		ctx: ctx,
		extSystemCode: extSystemCode,
		membershipType: membershipType,
	}
}

// Execute executes the request
//  @return Status
func (a *ProfileApiService) PostMembershipNumberExecute(r ApiPostMembershipNumberRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfileApiService.PostMembershipNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/externalSystems/{extSystemCode}/memberships/{membershipType}/membershipIDs"
	localVarPath = strings.Replace(localVarPath, "{"+"extSystemCode"+"}", url.PathEscape(parameterValueToString(r.extSystemCode, "extSystemCode")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"membershipType"+"}", url.PathEscape(parameterValueToString(r.membershipType, "membershipType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.extSystemCode) < 1 {
		return localVarReturnValue, nil, reportError("extSystemCode must have at least 1 elements")
	}
	if strlen(r.extSystemCode) > 2000 {
		return localVarReturnValue, nil, reportError("extSystemCode must have less than 2000 elements")
	}
	if strlen(r.membershipType) < 1 {
		return localVarReturnValue, nil, reportError("membershipType must have at least 1 elements")
	}
	if strlen(r.membershipType) > 2000 {
		return localVarReturnValue, nil, reportError("membershipType must have less than 2000 elements")
	}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.xAppKey == nil {
		return localVarReturnValue, nil, reportError("xAppKey is required and must be specified")
	}
	if r.membershipNumber == nil {
		return localVarReturnValue, nil, reportError("membershipNumber is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-app-key", r.xAppKey, "")
	if r.xTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-transactionId", r.xTransactionId, "")
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	// body params
	localVarPostBody = r.membershipNumber
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
