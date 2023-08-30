/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the ResGuestExternalInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResGuestExternalInfoType{}

// ResGuestExternalInfoType Specifies Company or Travel Agent profile using IATA or Corp. No.
type ResGuestExternalInfoType struct {
	// Given name, first name or names
	GivenName *string `json:"givenName,omitempty"`
	// Family name, last name.
	Surname *string `json:"surname,omitempty"`
}

// NewResGuestExternalInfoType instantiates a new ResGuestExternalInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResGuestExternalInfoType() *ResGuestExternalInfoType {
	this := ResGuestExternalInfoType{}
	return &this
}

// NewResGuestExternalInfoTypeWithDefaults instantiates a new ResGuestExternalInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResGuestExternalInfoTypeWithDefaults() *ResGuestExternalInfoType {
	this := ResGuestExternalInfoType{}
	return &this
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *ResGuestExternalInfoType) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResGuestExternalInfoType) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *ResGuestExternalInfoType) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *ResGuestExternalInfoType) SetGivenName(v string) {
	o.GivenName = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *ResGuestExternalInfoType) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResGuestExternalInfoType) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *ResGuestExternalInfoType) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *ResGuestExternalInfoType) SetSurname(v string) {
	o.Surname = &v
}

func (o ResGuestExternalInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResGuestExternalInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	return toSerialize, nil
}

type NullableResGuestExternalInfoType struct {
	value *ResGuestExternalInfoType
	isSet bool
}

func (v NullableResGuestExternalInfoType) Get() *ResGuestExternalInfoType {
	return v.value
}

func (v *NullableResGuestExternalInfoType) Set(val *ResGuestExternalInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableResGuestExternalInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableResGuestExternalInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResGuestExternalInfoType(val *ResGuestExternalInfoType) *NullableResGuestExternalInfoType {
	return &NullableResGuestExternalInfoType{value: val, isSet: true}
}

func (v NullableResGuestExternalInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResGuestExternalInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


