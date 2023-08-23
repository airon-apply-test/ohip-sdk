/*
OPI Token Exchange Service API

Oracle Token Exchange Service Specification<br /><br /> Compatible with OPERA Cloud release 1.0.1.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 1.0.1
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


// HotelsApiService HotelsApi service
type HotelsApiService service

type ApiOpenPaymentTokenExchangeRequest struct {
	ctx context.Context
	ApiService *HotelsApiService
	hotelCode string
	authorization *string
	xAppKey *string
	body *OpenPaymentTokenExchangeRequest
}

// Bearer token that needs to be passed which is generated post user authentication
func (r ApiOpenPaymentTokenExchangeRequest) Authorization(authorization string) ApiOpenPaymentTokenExchangeRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner Application Key
func (r ApiOpenPaymentTokenExchangeRequest) XAppKey(xAppKey string) ApiOpenPaymentTokenExchangeRequest {
	r.xAppKey = &xAppKey
	return r
}

// TokenRequest object
func (r ApiOpenPaymentTokenExchangeRequest) Body(body OpenPaymentTokenExchangeRequest) ApiOpenPaymentTokenExchangeRequest {
	r.body = &body
	return r
}

func (r ApiOpenPaymentTokenExchangeRequest) Execute() (*OpenPaymentTokenExchange200Response, *http.Response, error) {
	return r.ApiService.OpenPaymentTokenExchangeExecute(r)
}

/*
OpenPaymentTokenExchange Card Tokenization

Converts Primary Account Number (PAN) into Token issued by Payment Service Providers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param hotelCode OPERA Hotel Code
 @return ApiOpenPaymentTokenExchangeRequest
*/
func (a *HotelsApiService) OpenPaymentTokenExchange(ctx context.Context, hotelCode string) ApiOpenPaymentTokenExchangeRequest {
	return ApiOpenPaymentTokenExchangeRequest{
		ApiService: a,
		ctx: ctx,
		hotelCode: hotelCode,
	}
}

// Execute executes the request
//  @return OpenPaymentTokenExchange200Response
func (a *HotelsApiService) OpenPaymentTokenExchangeExecute(r ApiOpenPaymentTokenExchangeRequest) (*OpenPaymentTokenExchange200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpenPaymentTokenExchange200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HotelsApiService.OpenPaymentTokenExchange")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hotels/{hotelCode}/tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"hotelCode"+"}", url.PathEscape(parameterValueToString(r.hotelCode, "hotelCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.xAppKey == nil {
		return localVarReturnValue, nil, reportError("xAppKey is required and must be specified")
	}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-app-key", r.xAppKey, "")
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["jwt"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v OpenPaymentTokenExchange400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v OpenPaymentTokenExchange422Response
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
			var v OpenPaymentTokenExchange500Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v OpenPaymentTokenExchange504Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
