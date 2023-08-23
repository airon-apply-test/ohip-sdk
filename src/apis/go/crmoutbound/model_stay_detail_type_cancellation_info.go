/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StayDetailTypeCancellationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StayDetailTypeCancellationInfo{}

// StayDetailTypeCancellationInfo Information regarding why reservation has been/was cancelled.
type StayDetailTypeCancellationInfo struct {
	Description *string `json:"description,omitempty"`
	Code *string `json:"code,omitempty"`
	// Date when reservation was last cancelled.
	Date *string `json:"date,omitempty"`
}

// NewStayDetailTypeCancellationInfo instantiates a new StayDetailTypeCancellationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStayDetailTypeCancellationInfo() *StayDetailTypeCancellationInfo {
	this := StayDetailTypeCancellationInfo{}
	return &this
}

// NewStayDetailTypeCancellationInfoWithDefaults instantiates a new StayDetailTypeCancellationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStayDetailTypeCancellationInfoWithDefaults() *StayDetailTypeCancellationInfo {
	this := StayDetailTypeCancellationInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StayDetailTypeCancellationInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayDetailTypeCancellationInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StayDetailTypeCancellationInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StayDetailTypeCancellationInfo) SetDescription(v string) {
	o.Description = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *StayDetailTypeCancellationInfo) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayDetailTypeCancellationInfo) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *StayDetailTypeCancellationInfo) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *StayDetailTypeCancellationInfo) SetCode(v string) {
	o.Code = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *StayDetailTypeCancellationInfo) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayDetailTypeCancellationInfo) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *StayDetailTypeCancellationInfo) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *StayDetailTypeCancellationInfo) SetDate(v string) {
	o.Date = &v
}

func (o StayDetailTypeCancellationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StayDetailTypeCancellationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

type NullableStayDetailTypeCancellationInfo struct {
	value *StayDetailTypeCancellationInfo
	isSet bool
}

func (v NullableStayDetailTypeCancellationInfo) Get() *StayDetailTypeCancellationInfo {
	return v.value
}

func (v *NullableStayDetailTypeCancellationInfo) Set(val *StayDetailTypeCancellationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStayDetailTypeCancellationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStayDetailTypeCancellationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStayDetailTypeCancellationInfo(val *StayDetailTypeCancellationInfo) *NullableStayDetailTypeCancellationInfo {
	return &NullableStayDetailTypeCancellationInfo{value: val, isSet: true}
}

func (v NullableStayDetailTypeCancellationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStayDetailTypeCancellationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


