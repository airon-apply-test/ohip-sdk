/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cshoutbound

import (
	"encoding/json"
)

// checks if the ErrorInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorInstance{}

// ErrorInstance Complex type that contains error instance details for a REST call.
type ErrorInstance struct {
	// Absolute URI [RFC3986] that identifies the problem type.  When dereferenced, it SHOULD provide a human-readable summary of the problem (for example, using HTML).
	Type string `json:"type"`
	// Short, human-readable summary of the problem.  The summary SHOULD NOT change for subsequent occurrences of the problem, except for purposes of localization.
	Title string `json:"title"`
	// HTTP status code for this occurrence of the problem, set by the origin server.
	Status *int32 `json:"status,omitempty"`
	// Human-readable description specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// Absolute URI that identifies the specific occurrence of the problem.  It may or may not provide additional information if dereferenced.
	Instance *string `json:"instance,omitempty"`
	// Application error code, which is different from HTTP error code.
	OerrorCode *string `json:"o:errorCode,omitempty"`
	// Path to the problem at the resource or property level.
	OerrorPath *string `json:"o:errorPath,omitempty"`
}

// NewErrorInstance instantiates a new ErrorInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorInstance(type_ string, title string) *ErrorInstance {
	this := ErrorInstance{}
	this.Type = type_
	this.Title = title
	return &this
}

// NewErrorInstanceWithDefaults instantiates a new ErrorInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorInstanceWithDefaults() *ErrorInstance {
	this := ErrorInstance{}
	return &this
}

// GetType returns the Type field value
func (o *ErrorInstance) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ErrorInstance) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ErrorInstance) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *ErrorInstance) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ErrorInstance) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ErrorInstance) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ErrorInstance) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInstance) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ErrorInstance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ErrorInstance) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorInstance) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInstance) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorInstance) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorInstance) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ErrorInstance) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInstance) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ErrorInstance) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ErrorInstance) SetInstance(v string) {
	o.Instance = &v
}

// GetOerrorCode returns the OerrorCode field value if set, zero value otherwise.
func (o *ErrorInstance) GetOerrorCode() string {
	if o == nil || IsNil(o.OerrorCode) {
		var ret string
		return ret
	}
	return *o.OerrorCode
}

// GetOerrorCodeOk returns a tuple with the OerrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInstance) GetOerrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OerrorCode) {
		return nil, false
	}
	return o.OerrorCode, true
}

// HasOerrorCode returns a boolean if a field has been set.
func (o *ErrorInstance) HasOerrorCode() bool {
	if o != nil && !IsNil(o.OerrorCode) {
		return true
	}

	return false
}

// SetOerrorCode gets a reference to the given string and assigns it to the OerrorCode field.
func (o *ErrorInstance) SetOerrorCode(v string) {
	o.OerrorCode = &v
}

// GetOerrorPath returns the OerrorPath field value if set, zero value otherwise.
func (o *ErrorInstance) GetOerrorPath() string {
	if o == nil || IsNil(o.OerrorPath) {
		var ret string
		return ret
	}
	return *o.OerrorPath
}

// GetOerrorPathOk returns a tuple with the OerrorPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInstance) GetOerrorPathOk() (*string, bool) {
	if o == nil || IsNil(o.OerrorPath) {
		return nil, false
	}
	return o.OerrorPath, true
}

// HasOerrorPath returns a boolean if a field has been set.
func (o *ErrorInstance) HasOerrorPath() bool {
	if o != nil && !IsNil(o.OerrorPath) {
		return true
	}

	return false
}

// SetOerrorPath gets a reference to the given string and assigns it to the OerrorPath field.
func (o *ErrorInstance) SetOerrorPath(v string) {
	o.OerrorPath = &v
}

func (o ErrorInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.OerrorCode) {
		toSerialize["o:errorCode"] = o.OerrorCode
	}
	if !IsNil(o.OerrorPath) {
		toSerialize["o:errorPath"] = o.OerrorPath
	}
	return toSerialize, nil
}

type NullableErrorInstance struct {
	value *ErrorInstance
	isSet bool
}

func (v NullableErrorInstance) Get() *ErrorInstance {
	return v.value
}

func (v *NullableErrorInstance) Set(val *ErrorInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorInstance(val *ErrorInstance) *NullableErrorInstance {
	return &NullableErrorInstance{value: val, isSet: true}
}

func (v NullableErrorInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


