/*
OPI Token Exchange Service API

Oracle Token Exchange Service Specification<br /><br /> Compatible with OPERA Cloud release 1.0.1.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 1.0.1
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokenexchange

import (
	"encoding/json"
)

// checks if the OpenPaymentTokenExchange400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenPaymentTokenExchange400Response{}

// OpenPaymentTokenExchange400Response Problem Details for HTTP APIs
type OpenPaymentTokenExchange400Response struct {
	// Unique identifier value that is attached to the request that allows reference to a particular transaction
	CorrelationId *string `json:"correlationId,omitempty"`
	// Detailed error message
	Detail *string `json:"detail,omitempty"`
	// Application specific error code
	OerrorCode *string `json:"o:errorCode,omitempty"`
	// Drill down to the details
	OerrorDetails []ProblemDetail `json:"o:errorDetails,omitempty"`
	// JSON path to indicate where the error occurs
	OerrorPath *string `json:"o:errorPath,omitempty"`
	// HTTP status code
	Status *int32 `json:"status,omitempty"`
	// Summary error message
	Title string `json:"title"`
	// HTTP error code page
	Type string `json:"type"`
}

// NewOpenPaymentTokenExchange400Response instantiates a new OpenPaymentTokenExchange400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenPaymentTokenExchange400Response(title string, type_ string) *OpenPaymentTokenExchange400Response {
	this := OpenPaymentTokenExchange400Response{}
	this.Title = title
	this.Type = type_
	return &this
}

// NewOpenPaymentTokenExchange400ResponseWithDefaults instantiates a new OpenPaymentTokenExchange400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenPaymentTokenExchange400ResponseWithDefaults() *OpenPaymentTokenExchange400Response {
	this := OpenPaymentTokenExchange400Response{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *OpenPaymentTokenExchange400Response) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange400Response) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *OpenPaymentTokenExchange400Response) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *OpenPaymentTokenExchange400Response) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *OpenPaymentTokenExchange400Response) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange400Response) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *OpenPaymentTokenExchange400Response) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *OpenPaymentTokenExchange400Response) SetDetail(v string) {
	o.Detail = &v
}

// GetOerrorCode returns the OerrorCode field value if set, zero value otherwise.
func (o *OpenPaymentTokenExchange400Response) GetOerrorCode() string {
	if o == nil || IsNil(o.OerrorCode) {
		var ret string
		return ret
	}
	return *o.OerrorCode
}

// GetOerrorCodeOk returns a tuple with the OerrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange400Response) GetOerrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OerrorCode) {
		return nil, false
	}
	return o.OerrorCode, true
}

// HasOerrorCode returns a boolean if a field has been set.
func (o *OpenPaymentTokenExchange400Response) HasOerrorCode() bool {
	if o != nil && !IsNil(o.OerrorCode) {
		return true
	}

	return false
}

// SetOerrorCode gets a reference to the given string and assigns it to the OerrorCode field.
func (o *OpenPaymentTokenExchange400Response) SetOerrorCode(v string) {
	o.OerrorCode = &v
}

// GetOerrorDetails returns the OerrorDetails field value if set, zero value otherwise.
func (o *OpenPaymentTokenExchange400Response) GetOerrorDetails() []ProblemDetail {
	if o == nil || IsNil(o.OerrorDetails) {
		var ret []ProblemDetail
		return ret
	}
	return o.OerrorDetails
}

// GetOerrorDetailsOk returns a tuple with the OerrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange400Response) GetOerrorDetailsOk() ([]ProblemDetail, bool) {
	if o == nil || IsNil(o.OerrorDetails) {
		return nil, false
	}
	return o.OerrorDetails, true
}

// HasOerrorDetails returns a boolean if a field has been set.
func (o *OpenPaymentTokenExchange400Response) HasOerrorDetails() bool {
	if o != nil && !IsNil(o.OerrorDetails) {
		return true
	}

	return false
}

// SetOerrorDetails gets a reference to the given []ProblemDetail and assigns it to the OerrorDetails field.
func (o *OpenPaymentTokenExchange400Response) SetOerrorDetails(v []ProblemDetail) {
	o.OerrorDetails = v
}

// GetOerrorPath returns the OerrorPath field value if set, zero value otherwise.
func (o *OpenPaymentTokenExchange400Response) GetOerrorPath() string {
	if o == nil || IsNil(o.OerrorPath) {
		var ret string
		return ret
	}
	return *o.OerrorPath
}

// GetOerrorPathOk returns a tuple with the OerrorPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange400Response) GetOerrorPathOk() (*string, bool) {
	if o == nil || IsNil(o.OerrorPath) {
		return nil, false
	}
	return o.OerrorPath, true
}

// HasOerrorPath returns a boolean if a field has been set.
func (o *OpenPaymentTokenExchange400Response) HasOerrorPath() bool {
	if o != nil && !IsNil(o.OerrorPath) {
		return true
	}

	return false
}

// SetOerrorPath gets a reference to the given string and assigns it to the OerrorPath field.
func (o *OpenPaymentTokenExchange400Response) SetOerrorPath(v string) {
	o.OerrorPath = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OpenPaymentTokenExchange400Response) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange400Response) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OpenPaymentTokenExchange400Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *OpenPaymentTokenExchange400Response) SetStatus(v int32) {
	o.Status = &v
}

// GetTitle returns the Title field value
func (o *OpenPaymentTokenExchange400Response) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange400Response) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OpenPaymentTokenExchange400Response) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *OpenPaymentTokenExchange400Response) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange400Response) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OpenPaymentTokenExchange400Response) SetType(v string) {
	o.Type = v
}

func (o OpenPaymentTokenExchange400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenPaymentTokenExchange400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.OerrorCode) {
		toSerialize["o:errorCode"] = o.OerrorCode
	}
	if !IsNil(o.OerrorDetails) {
		toSerialize["o:errorDetails"] = o.OerrorDetails
	}
	if !IsNil(o.OerrorPath) {
		toSerialize["o:errorPath"] = o.OerrorPath
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableOpenPaymentTokenExchange400Response struct {
	value *OpenPaymentTokenExchange400Response
	isSet bool
}

func (v NullableOpenPaymentTokenExchange400Response) Get() *OpenPaymentTokenExchange400Response {
	return v.value
}

func (v *NullableOpenPaymentTokenExchange400Response) Set(val *OpenPaymentTokenExchange400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenPaymentTokenExchange400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenPaymentTokenExchange400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenPaymentTokenExchange400Response(val *OpenPaymentTokenExchange400Response) *NullableOpenPaymentTokenExchange400Response {
	return &NullableOpenPaymentTokenExchange400Response{value: val, isSet: true}
}

func (v NullableOpenPaymentTokenExchange400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenPaymentTokenExchange400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


