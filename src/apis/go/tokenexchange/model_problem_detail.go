/*
OPI Token Exchange Service API

Oracle Token Exchange Service Specification<br /><br /> Compatible with OPERA Cloud release 1.0.1.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 1.0.1
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProblemDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDetail{}

// ProblemDetail Problem Details for HTTP APIs
type ProblemDetail struct {
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

// NewProblemDetail instantiates a new ProblemDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetail(title string, type_ string) *ProblemDetail {
	this := ProblemDetail{}
	this.Title = title
	this.Type = type_
	return &this
}

// NewProblemDetailWithDefaults instantiates a new ProblemDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetailWithDefaults() *ProblemDetail {
	this := ProblemDetail{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *ProblemDetail) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *ProblemDetail) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *ProblemDetail) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ProblemDetail) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ProblemDetail) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ProblemDetail) SetDetail(v string) {
	o.Detail = &v
}

// GetOerrorCode returns the OerrorCode field value if set, zero value otherwise.
func (o *ProblemDetail) GetOerrorCode() string {
	if o == nil || IsNil(o.OerrorCode) {
		var ret string
		return ret
	}
	return *o.OerrorCode
}

// GetOerrorCodeOk returns a tuple with the OerrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetOerrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OerrorCode) {
		return nil, false
	}
	return o.OerrorCode, true
}

// HasOerrorCode returns a boolean if a field has been set.
func (o *ProblemDetail) HasOerrorCode() bool {
	if o != nil && !IsNil(o.OerrorCode) {
		return true
	}

	return false
}

// SetOerrorCode gets a reference to the given string and assigns it to the OerrorCode field.
func (o *ProblemDetail) SetOerrorCode(v string) {
	o.OerrorCode = &v
}

// GetOerrorDetails returns the OerrorDetails field value if set, zero value otherwise.
func (o *ProblemDetail) GetOerrorDetails() []ProblemDetail {
	if o == nil || IsNil(o.OerrorDetails) {
		var ret []ProblemDetail
		return ret
	}
	return o.OerrorDetails
}

// GetOerrorDetailsOk returns a tuple with the OerrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetOerrorDetailsOk() ([]ProblemDetail, bool) {
	if o == nil || IsNil(o.OerrorDetails) {
		return nil, false
	}
	return o.OerrorDetails, true
}

// HasOerrorDetails returns a boolean if a field has been set.
func (o *ProblemDetail) HasOerrorDetails() bool {
	if o != nil && !IsNil(o.OerrorDetails) {
		return true
	}

	return false
}

// SetOerrorDetails gets a reference to the given []ProblemDetail and assigns it to the OerrorDetails field.
func (o *ProblemDetail) SetOerrorDetails(v []ProblemDetail) {
	o.OerrorDetails = v
}

// GetOerrorPath returns the OerrorPath field value if set, zero value otherwise.
func (o *ProblemDetail) GetOerrorPath() string {
	if o == nil || IsNil(o.OerrorPath) {
		var ret string
		return ret
	}
	return *o.OerrorPath
}

// GetOerrorPathOk returns a tuple with the OerrorPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetOerrorPathOk() (*string, bool) {
	if o == nil || IsNil(o.OerrorPath) {
		return nil, false
	}
	return o.OerrorPath, true
}

// HasOerrorPath returns a boolean if a field has been set.
func (o *ProblemDetail) HasOerrorPath() bool {
	if o != nil && !IsNil(o.OerrorPath) {
		return true
	}

	return false
}

// SetOerrorPath gets a reference to the given string and assigns it to the OerrorPath field.
func (o *ProblemDetail) SetOerrorPath(v string) {
	o.OerrorPath = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProblemDetail) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProblemDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ProblemDetail) SetStatus(v int32) {
	o.Status = &v
}

// GetTitle returns the Title field value
func (o *ProblemDetail) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ProblemDetail) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *ProblemDetail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProblemDetail) SetType(v string) {
	o.Type = v
}

func (o ProblemDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDetail) ToMap() (map[string]interface{}, error) {
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

type NullableProblemDetail struct {
	value *ProblemDetail
	isSet bool
}

func (v NullableProblemDetail) Get() *ProblemDetail {
	return v.value
}

func (v *NullableProblemDetail) Set(val *ProblemDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetail(val *ProblemDetail) *NullableProblemDetail {
	return &NullableProblemDetail{value: val, isSet: true}
}

func (v NullableProblemDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


