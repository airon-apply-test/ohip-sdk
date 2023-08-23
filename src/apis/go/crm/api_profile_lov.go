/*
OPERA Cloud Customer Relationship Management API

APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
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


// ProfileLOVApiService ProfileLOVApi service
type ProfileLOVApiService service

type ApiGetAvailablePreferencesRequest struct {
	ctx context.Context
	ApiService *ProfileLOVApiService
	id *string
	hotelId *string
	authorization *string
	xAppKey *string
	xHotelid *string
	maxFetchSize *int32
	xExternalsystem *string
	acceptLanguage *string
}

// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
func (r ApiGetAvailablePreferencesRequest) Id(id string) ApiGetAvailablePreferencesRequest {
	r.id = &id
	return r
}

// Property for which preferences are to be fetched.
func (r ApiGetAvailablePreferencesRequest) HotelId(hotelId string) ApiGetAvailablePreferencesRequest {
	r.hotelId = &hotelId
	return r
}

// Bearer token that needs to be passed which is generated post user authentication
func (r ApiGetAvailablePreferencesRequest) Authorization(authorization string) ApiGetAvailablePreferencesRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r ApiGetAvailablePreferencesRequest) XAppKey(xAppKey string) ApiGetAvailablePreferencesRequest {
	r.xAppKey = &xAppKey
	return r
}

// Mandatory parameter to identify the hotel code where the end user is logged in
func (r ApiGetAvailablePreferencesRequest) XHotelid(xHotelid string) ApiGetAvailablePreferencesRequest {
	r.xHotelid = &xHotelid
	return r
}

// Maximum records to be fetched.
func (r ApiGetAvailablePreferencesRequest) MaxFetchSize(maxFetchSize int32) ApiGetAvailablePreferencesRequest {
	r.maxFetchSize = &maxFetchSize
	return r
}

// External system code.
func (r ApiGetAvailablePreferencesRequest) XExternalsystem(xExternalsystem string) ApiGetAvailablePreferencesRequest {
	r.xExternalsystem = &xExternalsystem
	return r
}

// Language code
func (r ApiGetAvailablePreferencesRequest) AcceptLanguage(acceptLanguage string) ApiGetAvailablePreferencesRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiGetAvailablePreferencesRequest) Execute() (*AvailablePreferences, *http.Response, error) {
	return r.ApiService.GetAvailablePreferencesExecute(r)
}

/*
GetAvailablePreferences Get available preferences

Use this API to fetch the available preferences List Of Values for profiles <p><strong>OperationId:</strong>getAvailablePreferences</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAvailablePreferencesRequest
*/
func (a *ProfileLOVApiService) GetAvailablePreferences(ctx context.Context) ApiGetAvailablePreferencesRequest {
	return ApiGetAvailablePreferencesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AvailablePreferences
func (a *ProfileLOVApiService) GetAvailablePreferencesExecute(r ApiGetAvailablePreferencesRequest) (*AvailablePreferences, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AvailablePreferences
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfileLOVApiService.GetAvailablePreferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profileListOfValues/availablePreferences"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}
	if r.hotelId == nil {
		return localVarReturnValue, nil, reportError("hotelId is required and must be specified")
	}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.xAppKey == nil {
		return localVarReturnValue, nil, reportError("xAppKey is required and must be specified")
	}
	if r.xHotelid == nil {
		return localVarReturnValue, nil, reportError("xHotelid is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "hotelId", r.hotelId, "")
	if r.maxFetchSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxFetchSize", r.maxFetchSize, "")
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
