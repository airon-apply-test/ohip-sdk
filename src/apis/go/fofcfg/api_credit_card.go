/*
OPERA Cloud Front Desk Configuration API

APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fofcfg

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// CreditCardApiService CreditCardApi service
type CreditCardApiService service

type CreditCardApiGetCreditCardInfoRequest struct {
	ctx context.Context
	ApiService *CreditCardApiService
	hotelId *string
	authorization *string
	xAppKey *string
	xHotelid *string
	reservationId *string
	idExtension *int32
	idContext *string
	type_ *string
	cardId *string
	cardIdExtension *int32
	cardIdContext *string
	cardIdType *string
	accessTransactionType *string
	xExternalsystem *string
	acceptLanguage *string
}

// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
func (r CreditCardApiGetCreditCardInfoRequest) HotelId(hotelId string) CreditCardApiGetCreditCardInfoRequest {
	r.hotelId = &hotelId
	return r
}

// Bearer token that needs to be passed which is generated post user authentication
func (r CreditCardApiGetCreditCardInfoRequest) Authorization(authorization string) CreditCardApiGetCreditCardInfoRequest {
	r.authorization = &authorization
	return r
}

// Client or Partner’s Application Key
func (r CreditCardApiGetCreditCardInfoRequest) XAppKey(xAppKey string) CreditCardApiGetCreditCardInfoRequest {
	r.xAppKey = &xAppKey
	return r
}

// Mandatory parameter to identify the hotel code where the end user is logged in
func (r CreditCardApiGetCreditCardInfoRequest) XHotelid(xHotelid string) CreditCardApiGetCreditCardInfoRequest {
	r.xHotelid = &xHotelid
	return r
}

// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
func (r CreditCardApiGetCreditCardInfoRequest) ReservationId(reservationId string) CreditCardApiGetCreditCardInfoRequest {
	r.reservationId = &reservationId
	return r
}

// Additional identifying value assigned by the creating system.
func (r CreditCardApiGetCreditCardInfoRequest) IdExtension(idExtension int32) CreditCardApiGetCreditCardInfoRequest {
	r.idExtension = &idExtension
	return r
}

// Used to identify the source of the identifier (e.g., IATA, ABTA).
func (r CreditCardApiGetCreditCardInfoRequest) IdContext(idContext string) CreditCardApiGetCreditCardInfoRequest {
	r.idContext = &idContext
	return r
}

// A reference to the type of object defined by the UniqueID element.
func (r CreditCardApiGetCreditCardInfoRequest) Type_(type_ string) CreditCardApiGetCreditCardInfoRequest {
	r.type_ = &type_
	return r
}

// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
func (r CreditCardApiGetCreditCardInfoRequest) CardId(cardId string) CreditCardApiGetCreditCardInfoRequest {
	r.cardId = &cardId
	return r
}

// Additional identifying value assigned by the creating system.
func (r CreditCardApiGetCreditCardInfoRequest) CardIdExtension(cardIdExtension int32) CreditCardApiGetCreditCardInfoRequest {
	r.cardIdExtension = &cardIdExtension
	return r
}

// Used to identify the source of the identifier (e.g., IATA, ABTA).
func (r CreditCardApiGetCreditCardInfoRequest) CardIdContext(cardIdContext string) CreditCardApiGetCreditCardInfoRequest {
	r.cardIdContext = &cardIdContext
	return r
}

// A reference to the type of object defined by the UniqueID element.
func (r CreditCardApiGetCreditCardInfoRequest) CardIdType(cardIdType string) CreditCardApiGetCreditCardInfoRequest {
	r.cardIdType = &cardIdType
	return r
}

// Currently supported transaction type categories used for credit card authorization.
func (r CreditCardApiGetCreditCardInfoRequest) AccessTransactionType(accessTransactionType string) CreditCardApiGetCreditCardInfoRequest {
	r.accessTransactionType = &accessTransactionType
	return r
}

// External system code.
func (r CreditCardApiGetCreditCardInfoRequest) XExternalsystem(xExternalsystem string) CreditCardApiGetCreditCardInfoRequest {
	r.xExternalsystem = &xExternalsystem
	return r
}

// Language code
func (r CreditCardApiGetCreditCardInfoRequest) AcceptLanguage(acceptLanguage string) CreditCardApiGetCreditCardInfoRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r CreditCardApiGetCreditCardInfoRequest) Execute() (*CreditCardInfo, *http.Response, error) {
	return r.ApiService.GetCreditCardInfoExecute(r)
}

/*
GetCreditCardInfo Get credit card token information

Use this API to return the token inclusive cardType and expirationDate for the credit card used in a specific reservation. Oracle does not provide any credit card numbers. Include the hotelId and the cardId in the parameters. The cardId can be returned with the getReservation operation. <p><strong>OperationId:</strong>getCreditCardInfo</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CreditCardApiGetCreditCardInfoRequest
*/
func (a *CreditCardApiService) GetCreditCardInfo(ctx context.Context) CreditCardApiGetCreditCardInfoRequest {
	return CreditCardApiGetCreditCardInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreditCardInfo
func (a *CreditCardApiService) GetCreditCardInfoExecute(r CreditCardApiGetCreditCardInfoRequest) (*CreditCardInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditCardInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditCardApiService.GetCreditCardInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/creditCardInfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	parameterAddToHeaderOrQuery(localVarQueryParams, "hotelId", r.hotelId, "")
	if r.reservationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reservationId", r.reservationId, "")
	}
	if r.idExtension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "idExtension", r.idExtension, "")
	}
	if r.idContext != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "idContext", r.idContext, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.cardId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cardId", r.cardId, "")
	}
	if r.cardIdExtension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cardIdExtension", r.cardIdExtension, "")
	}
	if r.cardIdContext != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cardIdContext", r.cardIdContext, "")
	}
	if r.cardIdType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cardIdType", r.cardIdType, "")
	}
	if r.accessTransactionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accessTransactionType", r.accessTransactionType, "")
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
