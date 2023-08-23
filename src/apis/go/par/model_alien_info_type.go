/*
OPERA Cloud Price Availability Rate API

APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AlienInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlienInfoType{}

// AlienInfoType immigration/visa information of a foreign person.
type AlienInfoType struct {
	// Alien Registration Number.
	AlienRegistrationNo *string `json:"alienRegistrationNo,omitempty"`
	// Immigration Status on an Alien.
	ImmigrationStatus *string `json:"immigrationStatus,omitempty"`
	// Visa Type of an Alien.
	VisaValidityType *string `json:"visaValidityType,omitempty"`
}

// NewAlienInfoType instantiates a new AlienInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlienInfoType() *AlienInfoType {
	this := AlienInfoType{}
	return &this
}

// NewAlienInfoTypeWithDefaults instantiates a new AlienInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlienInfoTypeWithDefaults() *AlienInfoType {
	this := AlienInfoType{}
	return &this
}

// GetAlienRegistrationNo returns the AlienRegistrationNo field value if set, zero value otherwise.
func (o *AlienInfoType) GetAlienRegistrationNo() string {
	if o == nil || IsNil(o.AlienRegistrationNo) {
		var ret string
		return ret
	}
	return *o.AlienRegistrationNo
}

// GetAlienRegistrationNoOk returns a tuple with the AlienRegistrationNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlienInfoType) GetAlienRegistrationNoOk() (*string, bool) {
	if o == nil || IsNil(o.AlienRegistrationNo) {
		return nil, false
	}
	return o.AlienRegistrationNo, true
}

// HasAlienRegistrationNo returns a boolean if a field has been set.
func (o *AlienInfoType) HasAlienRegistrationNo() bool {
	if o != nil && !IsNil(o.AlienRegistrationNo) {
		return true
	}

	return false
}

// SetAlienRegistrationNo gets a reference to the given string and assigns it to the AlienRegistrationNo field.
func (o *AlienInfoType) SetAlienRegistrationNo(v string) {
	o.AlienRegistrationNo = &v
}

// GetImmigrationStatus returns the ImmigrationStatus field value if set, zero value otherwise.
func (o *AlienInfoType) GetImmigrationStatus() string {
	if o == nil || IsNil(o.ImmigrationStatus) {
		var ret string
		return ret
	}
	return *o.ImmigrationStatus
}

// GetImmigrationStatusOk returns a tuple with the ImmigrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlienInfoType) GetImmigrationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ImmigrationStatus) {
		return nil, false
	}
	return o.ImmigrationStatus, true
}

// HasImmigrationStatus returns a boolean if a field has been set.
func (o *AlienInfoType) HasImmigrationStatus() bool {
	if o != nil && !IsNil(o.ImmigrationStatus) {
		return true
	}

	return false
}

// SetImmigrationStatus gets a reference to the given string and assigns it to the ImmigrationStatus field.
func (o *AlienInfoType) SetImmigrationStatus(v string) {
	o.ImmigrationStatus = &v
}

// GetVisaValidityType returns the VisaValidityType field value if set, zero value otherwise.
func (o *AlienInfoType) GetVisaValidityType() string {
	if o == nil || IsNil(o.VisaValidityType) {
		var ret string
		return ret
	}
	return *o.VisaValidityType
}

// GetVisaValidityTypeOk returns a tuple with the VisaValidityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlienInfoType) GetVisaValidityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VisaValidityType) {
		return nil, false
	}
	return o.VisaValidityType, true
}

// HasVisaValidityType returns a boolean if a field has been set.
func (o *AlienInfoType) HasVisaValidityType() bool {
	if o != nil && !IsNil(o.VisaValidityType) {
		return true
	}

	return false
}

// SetVisaValidityType gets a reference to the given string and assigns it to the VisaValidityType field.
func (o *AlienInfoType) SetVisaValidityType(v string) {
	o.VisaValidityType = &v
}

func (o AlienInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlienInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlienRegistrationNo) {
		toSerialize["alienRegistrationNo"] = o.AlienRegistrationNo
	}
	if !IsNil(o.ImmigrationStatus) {
		toSerialize["immigrationStatus"] = o.ImmigrationStatus
	}
	if !IsNil(o.VisaValidityType) {
		toSerialize["visaValidityType"] = o.VisaValidityType
	}
	return toSerialize, nil
}

type NullableAlienInfoType struct {
	value *AlienInfoType
	isSet bool
}

func (v NullableAlienInfoType) Get() *AlienInfoType {
	return v.value
}

func (v *NullableAlienInfoType) Set(val *AlienInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableAlienInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableAlienInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlienInfoType(val *AlienInfoType) *NullableAlienInfoType {
	return &NullableAlienInfoType{value: val, isSet: true}
}

func (v NullableAlienInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlienInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


