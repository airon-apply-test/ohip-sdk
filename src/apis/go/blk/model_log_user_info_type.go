/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the LogUserInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogUserInfoType{}

// LogUserInfoType struct for LogUserInfoType
type LogUserInfoType struct {
	UserInfo *UserInfoType `json:"userInfo,omitempty"`
	// User Log Changes Time Stamp details
	LogDateTime *time.Time `json:"logDateTime,omitempty"`
}

// NewLogUserInfoType instantiates a new LogUserInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogUserInfoType() *LogUserInfoType {
	this := LogUserInfoType{}
	return &this
}

// NewLogUserInfoTypeWithDefaults instantiates a new LogUserInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogUserInfoTypeWithDefaults() *LogUserInfoType {
	this := LogUserInfoType{}
	return &this
}

// GetUserInfo returns the UserInfo field value if set, zero value otherwise.
func (o *LogUserInfoType) GetUserInfo() UserInfoType {
	if o == nil || IsNil(o.UserInfo) {
		var ret UserInfoType
		return ret
	}
	return *o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogUserInfoType) GetUserInfoOk() (*UserInfoType, bool) {
	if o == nil || IsNil(o.UserInfo) {
		return nil, false
	}
	return o.UserInfo, true
}

// HasUserInfo returns a boolean if a field has been set.
func (o *LogUserInfoType) HasUserInfo() bool {
	if o != nil && !IsNil(o.UserInfo) {
		return true
	}

	return false
}

// SetUserInfo gets a reference to the given UserInfoType and assigns it to the UserInfo field.
func (o *LogUserInfoType) SetUserInfo(v UserInfoType) {
	o.UserInfo = &v
}

// GetLogDateTime returns the LogDateTime field value if set, zero value otherwise.
func (o *LogUserInfoType) GetLogDateTime() time.Time {
	if o == nil || IsNil(o.LogDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LogDateTime
}

// GetLogDateTimeOk returns a tuple with the LogDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogUserInfoType) GetLogDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LogDateTime) {
		return nil, false
	}
	return o.LogDateTime, true
}

// HasLogDateTime returns a boolean if a field has been set.
func (o *LogUserInfoType) HasLogDateTime() bool {
	if o != nil && !IsNil(o.LogDateTime) {
		return true
	}

	return false
}

// SetLogDateTime gets a reference to the given time.Time and assigns it to the LogDateTime field.
func (o *LogUserInfoType) SetLogDateTime(v time.Time) {
	o.LogDateTime = &v
}

func (o LogUserInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogUserInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserInfo) {
		toSerialize["userInfo"] = o.UserInfo
	}
	if !IsNil(o.LogDateTime) {
		toSerialize["logDateTime"] = o.LogDateTime
	}
	return toSerialize, nil
}

type NullableLogUserInfoType struct {
	value *LogUserInfoType
	isSet bool
}

func (v NullableLogUserInfoType) Get() *LogUserInfoType {
	return v.value
}

func (v *NullableLogUserInfoType) Set(val *LogUserInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogUserInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogUserInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogUserInfoType(val *LogUserInfoType) *NullableLogUserInfoType {
	return &NullableLogUserInfoType{value: val, isSet: true}
}

func (v NullableLogUserInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogUserInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


