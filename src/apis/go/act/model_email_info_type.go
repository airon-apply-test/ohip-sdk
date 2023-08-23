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

// checks if the EmailInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailInfoType{}

// EmailInfoType Information on an email for the customer.
type EmailInfoType struct {
	Email *EmailType `json:"email,omitempty"`
}

// NewEmailInfoType instantiates a new EmailInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailInfoType() *EmailInfoType {
	this := EmailInfoType{}
	return &this
}

// NewEmailInfoTypeWithDefaults instantiates a new EmailInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailInfoTypeWithDefaults() *EmailInfoType {
	this := EmailInfoType{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmailInfoType) GetEmail() EmailType {
	if o == nil || IsNil(o.Email) {
		var ret EmailType
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailInfoType) GetEmailOk() (*EmailType, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailInfoType) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailType and assigns it to the Email field.
func (o *EmailInfoType) SetEmail(v EmailType) {
	o.Email = &v
}

func (o EmailInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableEmailInfoType struct {
	value *EmailInfoType
	isSet bool
}

func (v NullableEmailInfoType) Get() *EmailInfoType {
	return v.value
}

func (v *NullableEmailInfoType) Set(val *EmailInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailInfoType(val *EmailInfoType) *NullableEmailInfoType {
	return &NullableEmailInfoType{value: val, isSet: true}
}

func (v NullableEmailInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


