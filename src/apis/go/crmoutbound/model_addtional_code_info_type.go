/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the AddtionalCodeInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddtionalCodeInfoType{}

// AddtionalCodeInfoType Allows to supply additional information in the form of name value pairs in Code block, for a given MasterInfoType.
type AddtionalCodeInfoType struct {
	// Holds name of additional code information
	Code []AddtionalCodeInfoTypeCodeInner `json:"code,omitempty"`
}

// NewAddtionalCodeInfoType instantiates a new AddtionalCodeInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddtionalCodeInfoType() *AddtionalCodeInfoType {
	this := AddtionalCodeInfoType{}
	return &this
}

// NewAddtionalCodeInfoTypeWithDefaults instantiates a new AddtionalCodeInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddtionalCodeInfoTypeWithDefaults() *AddtionalCodeInfoType {
	this := AddtionalCodeInfoType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AddtionalCodeInfoType) GetCode() []AddtionalCodeInfoTypeCodeInner {
	if o == nil || IsNil(o.Code) {
		var ret []AddtionalCodeInfoTypeCodeInner
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddtionalCodeInfoType) GetCodeOk() ([]AddtionalCodeInfoTypeCodeInner, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AddtionalCodeInfoType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given []AddtionalCodeInfoTypeCodeInner and assigns it to the Code field.
func (o *AddtionalCodeInfoType) SetCode(v []AddtionalCodeInfoTypeCodeInner) {
	o.Code = v
}

func (o AddtionalCodeInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddtionalCodeInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableAddtionalCodeInfoType struct {
	value *AddtionalCodeInfoType
	isSet bool
}

func (v NullableAddtionalCodeInfoType) Get() *AddtionalCodeInfoType {
	return v.value
}

func (v *NullableAddtionalCodeInfoType) Set(val *AddtionalCodeInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableAddtionalCodeInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableAddtionalCodeInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddtionalCodeInfoType(val *AddtionalCodeInfoType) *NullableAddtionalCodeInfoType {
	return &NullableAddtionalCodeInfoType{value: val, isSet: true}
}

func (v NullableAddtionalCodeInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddtionalCodeInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


