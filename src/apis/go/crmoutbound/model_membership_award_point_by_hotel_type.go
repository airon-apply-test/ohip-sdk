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

// checks if the MembershipAwardPointByHotelType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipAwardPointByHotelType{}

// MembershipAwardPointByHotelType Award points information group by Hotel.
type MembershipAwardPointByHotelType struct {
	// Award points issued at Property level.
	HotelId *string `json:"hotelId,omitempty"`
	// Sum of total positive and negative points.
	TotalPoints *float32 `json:"totalPoints,omitempty"`
}

// NewMembershipAwardPointByHotelType instantiates a new MembershipAwardPointByHotelType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipAwardPointByHotelType() *MembershipAwardPointByHotelType {
	this := MembershipAwardPointByHotelType{}
	return &this
}

// NewMembershipAwardPointByHotelTypeWithDefaults instantiates a new MembershipAwardPointByHotelType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipAwardPointByHotelTypeWithDefaults() *MembershipAwardPointByHotelType {
	this := MembershipAwardPointByHotelType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *MembershipAwardPointByHotelType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwardPointByHotelType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *MembershipAwardPointByHotelType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *MembershipAwardPointByHotelType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetTotalPoints returns the TotalPoints field value if set, zero value otherwise.
func (o *MembershipAwardPointByHotelType) GetTotalPoints() float32 {
	if o == nil || IsNil(o.TotalPoints) {
		var ret float32
		return ret
	}
	return *o.TotalPoints
}

// GetTotalPointsOk returns a tuple with the TotalPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwardPointByHotelType) GetTotalPointsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalPoints) {
		return nil, false
	}
	return o.TotalPoints, true
}

// HasTotalPoints returns a boolean if a field has been set.
func (o *MembershipAwardPointByHotelType) HasTotalPoints() bool {
	if o != nil && !IsNil(o.TotalPoints) {
		return true
	}

	return false
}

// SetTotalPoints gets a reference to the given float32 and assigns it to the TotalPoints field.
func (o *MembershipAwardPointByHotelType) SetTotalPoints(v float32) {
	o.TotalPoints = &v
}

func (o MembershipAwardPointByHotelType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipAwardPointByHotelType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.TotalPoints) {
		toSerialize["totalPoints"] = o.TotalPoints
	}
	return toSerialize, nil
}

type NullableMembershipAwardPointByHotelType struct {
	value *MembershipAwardPointByHotelType
	isSet bool
}

func (v NullableMembershipAwardPointByHotelType) Get() *MembershipAwardPointByHotelType {
	return v.value
}

func (v *NullableMembershipAwardPointByHotelType) Set(val *MembershipAwardPointByHotelType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipAwardPointByHotelType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipAwardPointByHotelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipAwardPointByHotelType(val *MembershipAwardPointByHotelType) *NullableMembershipAwardPointByHotelType {
	return &NullableMembershipAwardPointByHotelType{value: val, isSet: true}
}

func (v NullableMembershipAwardPointByHotelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipAwardPointByHotelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


