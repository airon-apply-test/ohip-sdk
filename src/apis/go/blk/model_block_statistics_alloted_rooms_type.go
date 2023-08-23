/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BlockStatisticsAllotedRoomsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockStatisticsAllotedRoomsType{}

// BlockStatisticsAllotedRoomsType List of all room types with alloted rooms.
type BlockStatisticsAllotedRoomsType struct {
	RoomTypes []StatisticsRoomTypeInfoType `json:"roomTypes,omitempty"`
}

// NewBlockStatisticsAllotedRoomsType instantiates a new BlockStatisticsAllotedRoomsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockStatisticsAllotedRoomsType() *BlockStatisticsAllotedRoomsType {
	this := BlockStatisticsAllotedRoomsType{}
	return &this
}

// NewBlockStatisticsAllotedRoomsTypeWithDefaults instantiates a new BlockStatisticsAllotedRoomsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockStatisticsAllotedRoomsTypeWithDefaults() *BlockStatisticsAllotedRoomsType {
	this := BlockStatisticsAllotedRoomsType{}
	return &this
}

// GetRoomTypes returns the RoomTypes field value if set, zero value otherwise.
func (o *BlockStatisticsAllotedRoomsType) GetRoomTypes() []StatisticsRoomTypeInfoType {
	if o == nil || IsNil(o.RoomTypes) {
		var ret []StatisticsRoomTypeInfoType
		return ret
	}
	return o.RoomTypes
}

// GetRoomTypesOk returns a tuple with the RoomTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsAllotedRoomsType) GetRoomTypesOk() ([]StatisticsRoomTypeInfoType, bool) {
	if o == nil || IsNil(o.RoomTypes) {
		return nil, false
	}
	return o.RoomTypes, true
}

// HasRoomTypes returns a boolean if a field has been set.
func (o *BlockStatisticsAllotedRoomsType) HasRoomTypes() bool {
	if o != nil && !IsNil(o.RoomTypes) {
		return true
	}

	return false
}

// SetRoomTypes gets a reference to the given []StatisticsRoomTypeInfoType and assigns it to the RoomTypes field.
func (o *BlockStatisticsAllotedRoomsType) SetRoomTypes(v []StatisticsRoomTypeInfoType) {
	o.RoomTypes = v
}

func (o BlockStatisticsAllotedRoomsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockStatisticsAllotedRoomsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomTypes) {
		toSerialize["roomTypes"] = o.RoomTypes
	}
	return toSerialize, nil
}

type NullableBlockStatisticsAllotedRoomsType struct {
	value *BlockStatisticsAllotedRoomsType
	isSet bool
}

func (v NullableBlockStatisticsAllotedRoomsType) Get() *BlockStatisticsAllotedRoomsType {
	return v.value
}

func (v *NullableBlockStatisticsAllotedRoomsType) Set(val *BlockStatisticsAllotedRoomsType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockStatisticsAllotedRoomsType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockStatisticsAllotedRoomsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockStatisticsAllotedRoomsType(val *BlockStatisticsAllotedRoomsType) *NullableBlockStatisticsAllotedRoomsType {
	return &NullableBlockStatisticsAllotedRoomsType{value: val, isSet: true}
}

func (v NullableBlockStatisticsAllotedRoomsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockStatisticsAllotedRoomsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


