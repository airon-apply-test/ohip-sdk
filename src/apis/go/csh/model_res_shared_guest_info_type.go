/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResSharedGuestInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResSharedGuestInfoType{}

// ResSharedGuestInfoType Contains information regarding the share reservation.
type ResSharedGuestInfoType struct {
	ProfileId *ProfileId `json:"profileId,omitempty"`
	// Given name, first name or names
	FirstName *string `json:"firstName,omitempty"`
	// Family name, last name.
	LastName *string `json:"lastName,omitempty"`
	// String representation of the full name
	FullName *string `json:"fullName,omitempty"`
}

// NewResSharedGuestInfoType instantiates a new ResSharedGuestInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResSharedGuestInfoType() *ResSharedGuestInfoType {
	this := ResSharedGuestInfoType{}
	return &this
}

// NewResSharedGuestInfoTypeWithDefaults instantiates a new ResSharedGuestInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResSharedGuestInfoTypeWithDefaults() *ResSharedGuestInfoType {
	this := ResSharedGuestInfoType{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ResSharedGuestInfoType) GetProfileId() ProfileId {
	if o == nil || IsNil(o.ProfileId) {
		var ret ProfileId
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResSharedGuestInfoType) GetProfileIdOk() (*ProfileId, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ResSharedGuestInfoType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given ProfileId and assigns it to the ProfileId field.
func (o *ResSharedGuestInfoType) SetProfileId(v ProfileId) {
	o.ProfileId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ResSharedGuestInfoType) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResSharedGuestInfoType) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResSharedGuestInfoType) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ResSharedGuestInfoType) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ResSharedGuestInfoType) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResSharedGuestInfoType) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ResSharedGuestInfoType) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ResSharedGuestInfoType) SetLastName(v string) {
	o.LastName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *ResSharedGuestInfoType) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResSharedGuestInfoType) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *ResSharedGuestInfoType) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *ResSharedGuestInfoType) SetFullName(v string) {
	o.FullName = &v
}

func (o ResSharedGuestInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResSharedGuestInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	return toSerialize, nil
}

type NullableResSharedGuestInfoType struct {
	value *ResSharedGuestInfoType
	isSet bool
}

func (v NullableResSharedGuestInfoType) Get() *ResSharedGuestInfoType {
	return v.value
}

func (v *NullableResSharedGuestInfoType) Set(val *ResSharedGuestInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableResSharedGuestInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableResSharedGuestInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResSharedGuestInfoType(val *ResSharedGuestInfoType) *NullableResSharedGuestInfoType {
	return &NullableResSharedGuestInfoType{value: val, isSet: true}
}

func (v NullableResSharedGuestInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResSharedGuestInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


