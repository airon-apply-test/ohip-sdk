/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package act

import (
	"encoding/json"
)

// checks if the URLType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &URLType{}

// URLType Web site address, in IETF(The Internet Engineering Task Force) specified format.
type URLType struct {
	// Property Value
	Value *string `json:"value,omitempty"`
	// Defines the purpose of the URL address, such as personal, business, public, etc.
	Type *string `json:"type,omitempty"`
}

// NewURLType instantiates a new URLType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewURLType() *URLType {
	this := URLType{}
	return &this
}

// NewURLTypeWithDefaults instantiates a new URLType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewURLTypeWithDefaults() *URLType {
	this := URLType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *URLType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *URLType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *URLType) SetValue(v string) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *URLType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *URLType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *URLType) SetType(v string) {
	o.Type = &v
}

func (o URLType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o URLType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableURLType struct {
	value *URLType
	isSet bool
}

func (v NullableURLType) Get() *URLType {
	return v.value
}

func (v *NullableURLType) Set(val *URLType) {
	v.value = val
	v.isSet = true
}

func (v NullableURLType) IsSet() bool {
	return v.isSet
}

func (v *NullableURLType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableURLType(val *URLType) *NullableURLType {
	return &NullableURLType{value: val, isSet: true}
}

func (v NullableURLType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableURLType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


