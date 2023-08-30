/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// OperaExchangeContentApiService OperaExchangeContentApi service
type OperaExchangeContentApiService service

type OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest struct {
	ctx context.Context
	ApiService *OperaExchangeContentApiService
	messageId string
	authorization *string
	xAppKey *string
	xHotelid *string
	integrationMessageType *string
	integrationSystem *string
	xExternalsystem *string
	acceptLanguage *string
}

// Bearer token that needs to be passed which is generated post user authentication
func (r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) Authorization(authorization string) OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) XAppKey(xAppKey string) OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest {
	r.xAppKey = &xAppKey
	return r
}

// Mandatory parameter to identify the hotel code where the end user is logged in
func (r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) XHotelid(xHotelid string) OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest {
	r.xHotelid = &xHotelid
	return r
}

// Determines whether message is an inbound message or a outbound message
func (r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) IntegrationMessageType(integrationMessageType string) OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest {
	r.integrationMessageType = &integrationMessageType
	return r
}

// Represents both OXI and OXI HUB are installed.
func (r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) IntegrationSystem(integrationSystem string) OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest {
	r.integrationSystem = &integrationSystem
	return r
}

// External system code.
func (r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) XExternalsystem(xExternalsystem string) OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest {
	r.xExternalsystem = &xExternalsystem
	return r
}

// Language code
func (r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) AcceptLanguage(acceptLanguage string) OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) Execute() (*FetchIntegrationMessageAttachments, *http.Response, error) {
	return r.ApiService.GetIntegrationMessageAttachmentsExecute(r)
}

/*
GetIntegrationMessageAttachments Fetch integration message attachments

API to Fetch OXI Inbound/Outbound XML Message Content by Message Id. <p><strong>OperationId:</strong>getIntegrationMessageAttachments</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param messageId Unique OPERA internal ID used to find inbound/outbound Message in OPERA.
 @return OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest
*/
func (a *OperaExchangeContentApiService) GetIntegrationMessageAttachments(ctx context.Context, messageId string) OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest {
	return OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest{
		ApiService: a,
		ctx: ctx,
		messageId: messageId,
	}
}

// Execute executes the request
//  @return FetchIntegrationMessageAttachments
func (a *OperaExchangeContentApiService) GetIntegrationMessageAttachmentsExecute(r OperaExchangeContentApiGetIntegrationMessageAttachmentsRequest) (*FetchIntegrationMessageAttachments, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FetchIntegrationMessageAttachments
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperaExchangeContentApiService.GetIntegrationMessageAttachments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages/{messageId}/integrationMessageAttachments"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.messageId) < 1 {
		return localVarReturnValue, nil, reportError("messageId must have at least 1 elements")
	}
	if strlen(r.messageId) > 2000 {
		return localVarReturnValue, nil, reportError("messageId must have less than 2000 elements")
	}

	if r.integrationMessageType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "integrationMessageType", r.integrationMessageType, "")
	}
	if r.integrationSystem != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "integrationSystem", r.integrationSystem, "")
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
