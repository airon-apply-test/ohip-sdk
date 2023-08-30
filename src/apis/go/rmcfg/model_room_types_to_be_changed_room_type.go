/*
OPERA Cloud Room Configuration API

APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rmcfg

import (
	"encoding/json"
)

// checks if the RoomTypesToBeChangedRoomType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomTypesToBeChangedRoomType{}

// RoomTypesToBeChangedRoomType Room Type details to be changed.
type RoomTypesToBeChangedRoomType struct {
	RoomTypeDetails *RoomTypeType `json:"roomTypeDetails,omitempty"`
	// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
	HotelId *string `json:"hotelId,omitempty"`
}

// NewRoomTypesToBeChangedRoomType instantiates a new RoomTypesToBeChangedRoomType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomTypesToBeChangedRoomType() *RoomTypesToBeChangedRoomType {
	this := RoomTypesToBeChangedRoomType{}
	return &this
}

// NewRoomTypesToBeChangedRoomTypeWithDefaults instantiates a new RoomTypesToBeChangedRoomType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomTypesToBeChangedRoomTypeWithDefaults() *RoomTypesToBeChangedRoomType {
	this := RoomTypesToBeChangedRoomType{}
	return &this
}

// GetRoomTypeDetails returns the RoomTypeDetails field value if set, zero value otherwise.
func (o *RoomTypesToBeChangedRoomType) GetRoomTypeDetails() RoomTypeType {
	if o == nil || IsNil(o.RoomTypeDetails) {
		var ret RoomTypeType
		return ret
	}
	return *o.RoomTypeDetails
}

// GetRoomTypeDetailsOk returns a tuple with the RoomTypeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypesToBeChangedRoomType) GetRoomTypeDetailsOk() (*RoomTypeType, bool) {
	if o == nil || IsNil(o.RoomTypeDetails) {
		return nil, false
	}
	return o.RoomTypeDetails, true
}

// HasRoomTypeDetails returns a boolean if a field has been set.
func (o *RoomTypesToBeChangedRoomType) HasRoomTypeDetails() bool {
	if o != nil && !IsNil(o.RoomTypeDetails) {
		return true
	}

	return false
}

// SetRoomTypeDetails gets a reference to the given RoomTypeType and assigns it to the RoomTypeDetails field.
func (o *RoomTypesToBeChangedRoomType) SetRoomTypeDetails(v RoomTypeType) {
	o.RoomTypeDetails = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *RoomTypesToBeChangedRoomType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypesToBeChangedRoomType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *RoomTypesToBeChangedRoomType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *RoomTypesToBeChangedRoomType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o RoomTypesToBeChangedRoomType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomTypesToBeChangedRoomType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomTypeDetails) {
		toSerialize["roomTypeDetails"] = o.RoomTypeDetails
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableRoomTypesToBeChangedRoomType struct {
	value *RoomTypesToBeChangedRoomType
	isSet bool
}

func (v NullableRoomTypesToBeChangedRoomType) Get() *RoomTypesToBeChangedRoomType {
	return v.value
}

func (v *NullableRoomTypesToBeChangedRoomType) Set(val *RoomTypesToBeChangedRoomType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomTypesToBeChangedRoomType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomTypesToBeChangedRoomType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomTypesToBeChangedRoomType(val *RoomTypesToBeChangedRoomType) *NullableRoomTypesToBeChangedRoomType {
	return &NullableRoomTypesToBeChangedRoomType{value: val, isSet: true}
}

func (v NullableRoomTypesToBeChangedRoomType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomTypesToBeChangedRoomType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


