/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProfileTypeTelephones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileTypeTelephones{}

// ProfileTypeTelephones List of Telephone Number Information
type ProfileTypeTelephones struct {
	// Collection of Detailed information on telephone/fax for the customer.
	TelephoneInfo []TelephoneInfoType `json:"telephoneInfo,omitempty"`
}

// NewProfileTypeTelephones instantiates a new ProfileTypeTelephones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileTypeTelephones() *ProfileTypeTelephones {
	this := ProfileTypeTelephones{}
	return &this
}

// NewProfileTypeTelephonesWithDefaults instantiates a new ProfileTypeTelephones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileTypeTelephonesWithDefaults() *ProfileTypeTelephones {
	this := ProfileTypeTelephones{}
	return &this
}

// GetTelephoneInfo returns the TelephoneInfo field value if set, zero value otherwise.
func (o *ProfileTypeTelephones) GetTelephoneInfo() []TelephoneInfoType {
	if o == nil || IsNil(o.TelephoneInfo) {
		var ret []TelephoneInfoType
		return ret
	}
	return o.TelephoneInfo
}

// GetTelephoneInfoOk returns a tuple with the TelephoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeTelephones) GetTelephoneInfoOk() ([]TelephoneInfoType, bool) {
	if o == nil || IsNil(o.TelephoneInfo) {
		return nil, false
	}
	return o.TelephoneInfo, true
}

// HasTelephoneInfo returns a boolean if a field has been set.
func (o *ProfileTypeTelephones) HasTelephoneInfo() bool {
	if o != nil && !IsNil(o.TelephoneInfo) {
		return true
	}

	return false
}

// SetTelephoneInfo gets a reference to the given []TelephoneInfoType and assigns it to the TelephoneInfo field.
func (o *ProfileTypeTelephones) SetTelephoneInfo(v []TelephoneInfoType) {
	o.TelephoneInfo = v
}

func (o ProfileTypeTelephones) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileTypeTelephones) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TelephoneInfo) {
		toSerialize["telephoneInfo"] = o.TelephoneInfo
	}
	return toSerialize, nil
}

type NullableProfileTypeTelephones struct {
	value *ProfileTypeTelephones
	isSet bool
}

func (v NullableProfileTypeTelephones) Get() *ProfileTypeTelephones {
	return v.value
}

func (v *NullableProfileTypeTelephones) Set(val *ProfileTypeTelephones) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileTypeTelephones) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileTypeTelephones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileTypeTelephones(val *ProfileTypeTelephones) *NullableProfileTypeTelephones {
	return &NullableProfileTypeTelephones{value: val, isSet: true}
}

func (v NullableProfileTypeTelephones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileTypeTelephones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


