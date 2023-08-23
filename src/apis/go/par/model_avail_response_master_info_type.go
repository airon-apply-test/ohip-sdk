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

// checks if the AvailResponseMasterInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailResponseMasterInfoType{}

// AvailResponseMasterInfoType Availability master info type
type AvailResponseMasterInfoType struct {
	RoomTypes *RoomTypeMasterInfoType `json:"roomTypes,omitempty"`
	RatePlans *AvailRatePlanInfoType `json:"ratePlans,omitempty"`
}

// NewAvailResponseMasterInfoType instantiates a new AvailResponseMasterInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailResponseMasterInfoType() *AvailResponseMasterInfoType {
	this := AvailResponseMasterInfoType{}
	return &this
}

// NewAvailResponseMasterInfoTypeWithDefaults instantiates a new AvailResponseMasterInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailResponseMasterInfoTypeWithDefaults() *AvailResponseMasterInfoType {
	this := AvailResponseMasterInfoType{}
	return &this
}

// GetRoomTypes returns the RoomTypes field value if set, zero value otherwise.
func (o *AvailResponseMasterInfoType) GetRoomTypes() RoomTypeMasterInfoType {
	if o == nil || IsNil(o.RoomTypes) {
		var ret RoomTypeMasterInfoType
		return ret
	}
	return *o.RoomTypes
}

// GetRoomTypesOk returns a tuple with the RoomTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailResponseMasterInfoType) GetRoomTypesOk() (*RoomTypeMasterInfoType, bool) {
	if o == nil || IsNil(o.RoomTypes) {
		return nil, false
	}
	return o.RoomTypes, true
}

// HasRoomTypes returns a boolean if a field has been set.
func (o *AvailResponseMasterInfoType) HasRoomTypes() bool {
	if o != nil && !IsNil(o.RoomTypes) {
		return true
	}

	return false
}

// SetRoomTypes gets a reference to the given RoomTypeMasterInfoType and assigns it to the RoomTypes field.
func (o *AvailResponseMasterInfoType) SetRoomTypes(v RoomTypeMasterInfoType) {
	o.RoomTypes = &v
}

// GetRatePlans returns the RatePlans field value if set, zero value otherwise.
func (o *AvailResponseMasterInfoType) GetRatePlans() AvailRatePlanInfoType {
	if o == nil || IsNil(o.RatePlans) {
		var ret AvailRatePlanInfoType
		return ret
	}
	return *o.RatePlans
}

// GetRatePlansOk returns a tuple with the RatePlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailResponseMasterInfoType) GetRatePlansOk() (*AvailRatePlanInfoType, bool) {
	if o == nil || IsNil(o.RatePlans) {
		return nil, false
	}
	return o.RatePlans, true
}

// HasRatePlans returns a boolean if a field has been set.
func (o *AvailResponseMasterInfoType) HasRatePlans() bool {
	if o != nil && !IsNil(o.RatePlans) {
		return true
	}

	return false
}

// SetRatePlans gets a reference to the given AvailRatePlanInfoType and assigns it to the RatePlans field.
func (o *AvailResponseMasterInfoType) SetRatePlans(v AvailRatePlanInfoType) {
	o.RatePlans = &v
}

func (o AvailResponseMasterInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailResponseMasterInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomTypes) {
		toSerialize["roomTypes"] = o.RoomTypes
	}
	if !IsNil(o.RatePlans) {
		toSerialize["ratePlans"] = o.RatePlans
	}
	return toSerialize, nil
}

type NullableAvailResponseMasterInfoType struct {
	value *AvailResponseMasterInfoType
	isSet bool
}

func (v NullableAvailResponseMasterInfoType) Get() *AvailResponseMasterInfoType {
	return v.value
}

func (v *NullableAvailResponseMasterInfoType) Set(val *AvailResponseMasterInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailResponseMasterInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailResponseMasterInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailResponseMasterInfoType(val *AvailResponseMasterInfoType) *NullableAvailResponseMasterInfoType {
	return &NullableAvailResponseMasterInfoType{value: val, isSet: true}
}

func (v NullableAvailResponseMasterInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailResponseMasterInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


