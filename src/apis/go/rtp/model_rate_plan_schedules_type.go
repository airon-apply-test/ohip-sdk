/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtp

import (
	"encoding/json"
)

// checks if the RatePlanSchedulesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatePlanSchedulesType{}

// RatePlanSchedulesType struct for RatePlanSchedulesType
type RatePlanSchedulesType struct {
	RatePlanScheduleId *UniqueIDType `json:"ratePlanScheduleId,omitempty"`
	RatePlanScheduleDetail *RatePlanScheduleDetailType `json:"ratePlanScheduleDetail,omitempty"`
}

// NewRatePlanSchedulesType instantiates a new RatePlanSchedulesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatePlanSchedulesType() *RatePlanSchedulesType {
	this := RatePlanSchedulesType{}
	return &this
}

// NewRatePlanSchedulesTypeWithDefaults instantiates a new RatePlanSchedulesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatePlanSchedulesTypeWithDefaults() *RatePlanSchedulesType {
	this := RatePlanSchedulesType{}
	return &this
}

// GetRatePlanScheduleId returns the RatePlanScheduleId field value if set, zero value otherwise.
func (o *RatePlanSchedulesType) GetRatePlanScheduleId() UniqueIDType {
	if o == nil || IsNil(o.RatePlanScheduleId) {
		var ret UniqueIDType
		return ret
	}
	return *o.RatePlanScheduleId
}

// GetRatePlanScheduleIdOk returns a tuple with the RatePlanScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanSchedulesType) GetRatePlanScheduleIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.RatePlanScheduleId) {
		return nil, false
	}
	return o.RatePlanScheduleId, true
}

// HasRatePlanScheduleId returns a boolean if a field has been set.
func (o *RatePlanSchedulesType) HasRatePlanScheduleId() bool {
	if o != nil && !IsNil(o.RatePlanScheduleId) {
		return true
	}

	return false
}

// SetRatePlanScheduleId gets a reference to the given UniqueIDType and assigns it to the RatePlanScheduleId field.
func (o *RatePlanSchedulesType) SetRatePlanScheduleId(v UniqueIDType) {
	o.RatePlanScheduleId = &v
}

// GetRatePlanScheduleDetail returns the RatePlanScheduleDetail field value if set, zero value otherwise.
func (o *RatePlanSchedulesType) GetRatePlanScheduleDetail() RatePlanScheduleDetailType {
	if o == nil || IsNil(o.RatePlanScheduleDetail) {
		var ret RatePlanScheduleDetailType
		return ret
	}
	return *o.RatePlanScheduleDetail
}

// GetRatePlanScheduleDetailOk returns a tuple with the RatePlanScheduleDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanSchedulesType) GetRatePlanScheduleDetailOk() (*RatePlanScheduleDetailType, bool) {
	if o == nil || IsNil(o.RatePlanScheduleDetail) {
		return nil, false
	}
	return o.RatePlanScheduleDetail, true
}

// HasRatePlanScheduleDetail returns a boolean if a field has been set.
func (o *RatePlanSchedulesType) HasRatePlanScheduleDetail() bool {
	if o != nil && !IsNil(o.RatePlanScheduleDetail) {
		return true
	}

	return false
}

// SetRatePlanScheduleDetail gets a reference to the given RatePlanScheduleDetailType and assigns it to the RatePlanScheduleDetail field.
func (o *RatePlanSchedulesType) SetRatePlanScheduleDetail(v RatePlanScheduleDetailType) {
	o.RatePlanScheduleDetail = &v
}

func (o RatePlanSchedulesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatePlanSchedulesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RatePlanScheduleId) {
		toSerialize["ratePlanScheduleId"] = o.RatePlanScheduleId
	}
	if !IsNil(o.RatePlanScheduleDetail) {
		toSerialize["ratePlanScheduleDetail"] = o.RatePlanScheduleDetail
	}
	return toSerialize, nil
}

type NullableRatePlanSchedulesType struct {
	value *RatePlanSchedulesType
	isSet bool
}

func (v NullableRatePlanSchedulesType) Get() *RatePlanSchedulesType {
	return v.value
}

func (v *NullableRatePlanSchedulesType) Set(val *RatePlanSchedulesType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePlanSchedulesType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePlanSchedulesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePlanSchedulesType(val *RatePlanSchedulesType) *NullableRatePlanSchedulesType {
	return &NullableRatePlanSchedulesType{value: val, isSet: true}
}

func (v NullableRatePlanSchedulesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePlanSchedulesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


