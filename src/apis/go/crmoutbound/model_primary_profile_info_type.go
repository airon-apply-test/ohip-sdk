/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PrimaryProfileInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrimaryProfileInfoType{}

// PrimaryProfileInfoType Contains basic information of profile.
type PrimaryProfileInfoType struct {
	ProfileId *UniqueIDType `json:"profileId,omitempty"`
	ProfileType *ProfileTypeType `json:"profileType,omitempty"`
	// Name of the account.
	ProfileName *string `json:"profileName,omitempty"`
}

// NewPrimaryProfileInfoType instantiates a new PrimaryProfileInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrimaryProfileInfoType() *PrimaryProfileInfoType {
	this := PrimaryProfileInfoType{}
	return &this
}

// NewPrimaryProfileInfoTypeWithDefaults instantiates a new PrimaryProfileInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrimaryProfileInfoTypeWithDefaults() *PrimaryProfileInfoType {
	this := PrimaryProfileInfoType{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *PrimaryProfileInfoType) GetProfileId() UniqueIDType {
	if o == nil || IsNil(o.ProfileId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryProfileInfoType) GetProfileIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *PrimaryProfileInfoType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given UniqueIDType and assigns it to the ProfileId field.
func (o *PrimaryProfileInfoType) SetProfileId(v UniqueIDType) {
	o.ProfileId = &v
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *PrimaryProfileInfoType) GetProfileType() ProfileTypeType {
	if o == nil || IsNil(o.ProfileType) {
		var ret ProfileTypeType
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryProfileInfoType) GetProfileTypeOk() (*ProfileTypeType, bool) {
	if o == nil || IsNil(o.ProfileType) {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *PrimaryProfileInfoType) HasProfileType() bool {
	if o != nil && !IsNil(o.ProfileType) {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given ProfileTypeType and assigns it to the ProfileType field.
func (o *PrimaryProfileInfoType) SetProfileType(v ProfileTypeType) {
	o.ProfileType = &v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *PrimaryProfileInfoType) GetProfileName() string {
	if o == nil || IsNil(o.ProfileName) {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryProfileInfoType) GetProfileNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileName) {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *PrimaryProfileInfoType) HasProfileName() bool {
	if o != nil && !IsNil(o.ProfileName) {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *PrimaryProfileInfoType) SetProfileName(v string) {
	o.ProfileName = &v
}

func (o PrimaryProfileInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrimaryProfileInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.ProfileType) {
		toSerialize["profileType"] = o.ProfileType
	}
	if !IsNil(o.ProfileName) {
		toSerialize["profileName"] = o.ProfileName
	}
	return toSerialize, nil
}

type NullablePrimaryProfileInfoType struct {
	value *PrimaryProfileInfoType
	isSet bool
}

func (v NullablePrimaryProfileInfoType) Get() *PrimaryProfileInfoType {
	return v.value
}

func (v *NullablePrimaryProfileInfoType) Set(val *PrimaryProfileInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrimaryProfileInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrimaryProfileInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrimaryProfileInfoType(val *PrimaryProfileInfoType) *NullablePrimaryProfileInfoType {
	return &NullablePrimaryProfileInfoType{value: val, isSet: true}
}

func (v NullablePrimaryProfileInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrimaryProfileInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


