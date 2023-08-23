/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResAccompanyGuestInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResAccompanyGuestInfoType{}

// ResAccompanyGuestInfoType struct for ResAccompanyGuestInfoType
type ResAccompanyGuestInfoType struct {
	// Given name, first name or names
	FirstName *string `json:"firstName,omitempty"`
	// Family name, last name.
	LastName *string `json:"lastName,omitempty"`
	// String representation of the full name
	FullName *string `json:"fullName,omitempty"`
	// Unique identifier of the police registration card number.
	RegistrationCardNo *string `json:"registrationCardNo,omitempty"`
	// Unique Id that references an object uniquely in the system.
	ProfileIdList []UniqueIDType `json:"profileIdList,omitempty"`
}

// NewResAccompanyGuestInfoType instantiates a new ResAccompanyGuestInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResAccompanyGuestInfoType() *ResAccompanyGuestInfoType {
	this := ResAccompanyGuestInfoType{}
	return &this
}

// NewResAccompanyGuestInfoTypeWithDefaults instantiates a new ResAccompanyGuestInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResAccompanyGuestInfoTypeWithDefaults() *ResAccompanyGuestInfoType {
	this := ResAccompanyGuestInfoType{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ResAccompanyGuestInfoType) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResAccompanyGuestInfoType) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResAccompanyGuestInfoType) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ResAccompanyGuestInfoType) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ResAccompanyGuestInfoType) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResAccompanyGuestInfoType) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ResAccompanyGuestInfoType) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ResAccompanyGuestInfoType) SetLastName(v string) {
	o.LastName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *ResAccompanyGuestInfoType) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResAccompanyGuestInfoType) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *ResAccompanyGuestInfoType) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *ResAccompanyGuestInfoType) SetFullName(v string) {
	o.FullName = &v
}

// GetRegistrationCardNo returns the RegistrationCardNo field value if set, zero value otherwise.
func (o *ResAccompanyGuestInfoType) GetRegistrationCardNo() string {
	if o == nil || IsNil(o.RegistrationCardNo) {
		var ret string
		return ret
	}
	return *o.RegistrationCardNo
}

// GetRegistrationCardNoOk returns a tuple with the RegistrationCardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResAccompanyGuestInfoType) GetRegistrationCardNoOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationCardNo) {
		return nil, false
	}
	return o.RegistrationCardNo, true
}

// HasRegistrationCardNo returns a boolean if a field has been set.
func (o *ResAccompanyGuestInfoType) HasRegistrationCardNo() bool {
	if o != nil && !IsNil(o.RegistrationCardNo) {
		return true
	}

	return false
}

// SetRegistrationCardNo gets a reference to the given string and assigns it to the RegistrationCardNo field.
func (o *ResAccompanyGuestInfoType) SetRegistrationCardNo(v string) {
	o.RegistrationCardNo = &v
}

// GetProfileIdList returns the ProfileIdList field value if set, zero value otherwise.
func (o *ResAccompanyGuestInfoType) GetProfileIdList() []UniqueIDType {
	if o == nil || IsNil(o.ProfileIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ProfileIdList
}

// GetProfileIdListOk returns a tuple with the ProfileIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResAccompanyGuestInfoType) GetProfileIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileIdList) {
		return nil, false
	}
	return o.ProfileIdList, true
}

// HasProfileIdList returns a boolean if a field has been set.
func (o *ResAccompanyGuestInfoType) HasProfileIdList() bool {
	if o != nil && !IsNil(o.ProfileIdList) {
		return true
	}

	return false
}

// SetProfileIdList gets a reference to the given []UniqueIDType and assigns it to the ProfileIdList field.
func (o *ResAccompanyGuestInfoType) SetProfileIdList(v []UniqueIDType) {
	o.ProfileIdList = v
}

func (o ResAccompanyGuestInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResAccompanyGuestInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.RegistrationCardNo) {
		toSerialize["registrationCardNo"] = o.RegistrationCardNo
	}
	if !IsNil(o.ProfileIdList) {
		toSerialize["profileIdList"] = o.ProfileIdList
	}
	return toSerialize, nil
}

type NullableResAccompanyGuestInfoType struct {
	value *ResAccompanyGuestInfoType
	isSet bool
}

func (v NullableResAccompanyGuestInfoType) Get() *ResAccompanyGuestInfoType {
	return v.value
}

func (v *NullableResAccompanyGuestInfoType) Set(val *ResAccompanyGuestInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableResAccompanyGuestInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableResAccompanyGuestInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResAccompanyGuestInfoType(val *ResAccompanyGuestInfoType) *NullableResAccompanyGuestInfoType {
	return &NullableResAccompanyGuestInfoType{value: val, isSet: true}
}

func (v NullableResAccompanyGuestInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResAccompanyGuestInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


