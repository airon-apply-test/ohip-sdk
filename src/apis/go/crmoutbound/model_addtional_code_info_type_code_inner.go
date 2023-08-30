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

// checks if the AddtionalCodeInfoTypeCodeInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddtionalCodeInfoTypeCodeInner{}

// AddtionalCodeInfoTypeCodeInner struct for AddtionalCodeInfoTypeCodeInner
type AddtionalCodeInfoTypeCodeInner struct {
	Name *MasterInfoCodeDetailType `json:"name,omitempty"`
	// Holds value of additional code information
	Value *string `json:"value,omitempty"`
}

// NewAddtionalCodeInfoTypeCodeInner instantiates a new AddtionalCodeInfoTypeCodeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddtionalCodeInfoTypeCodeInner() *AddtionalCodeInfoTypeCodeInner {
	this := AddtionalCodeInfoTypeCodeInner{}
	return &this
}

// NewAddtionalCodeInfoTypeCodeInnerWithDefaults instantiates a new AddtionalCodeInfoTypeCodeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddtionalCodeInfoTypeCodeInnerWithDefaults() *AddtionalCodeInfoTypeCodeInner {
	this := AddtionalCodeInfoTypeCodeInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddtionalCodeInfoTypeCodeInner) GetName() MasterInfoCodeDetailType {
	if o == nil || IsNil(o.Name) {
		var ret MasterInfoCodeDetailType
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddtionalCodeInfoTypeCodeInner) GetNameOk() (*MasterInfoCodeDetailType, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddtionalCodeInfoTypeCodeInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given MasterInfoCodeDetailType and assigns it to the Name field.
func (o *AddtionalCodeInfoTypeCodeInner) SetName(v MasterInfoCodeDetailType) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AddtionalCodeInfoTypeCodeInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddtionalCodeInfoTypeCodeInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AddtionalCodeInfoTypeCodeInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AddtionalCodeInfoTypeCodeInner) SetValue(v string) {
	o.Value = &v
}

func (o AddtionalCodeInfoTypeCodeInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddtionalCodeInfoTypeCodeInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAddtionalCodeInfoTypeCodeInner struct {
	value *AddtionalCodeInfoTypeCodeInner
	isSet bool
}

func (v NullableAddtionalCodeInfoTypeCodeInner) Get() *AddtionalCodeInfoTypeCodeInner {
	return v.value
}

func (v *NullableAddtionalCodeInfoTypeCodeInner) Set(val *AddtionalCodeInfoTypeCodeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAddtionalCodeInfoTypeCodeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAddtionalCodeInfoTypeCodeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddtionalCodeInfoTypeCodeInner(val *AddtionalCodeInfoTypeCodeInner) *NullableAddtionalCodeInfoTypeCodeInner {
	return &NullableAddtionalCodeInfoTypeCodeInner{value: val, isSet: true}
}

func (v NullableAddtionalCodeInfoTypeCodeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddtionalCodeInfoTypeCodeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


