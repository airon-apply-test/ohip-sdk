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

// checks if the BlockWashInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockWashInfoType{}

// BlockWashInfoType Block information and allocated room types information for performing a wash operation.
type BlockWashInfoType struct {
	BlockInfo *BlockWashInfoTypeBlockInfo `json:"blockInfo,omitempty"`
	// List of allocated room types for the block.
	RoomTypes []string `json:"roomTypes,omitempty"`
}

// NewBlockWashInfoType instantiates a new BlockWashInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockWashInfoType() *BlockWashInfoType {
	this := BlockWashInfoType{}
	return &this
}

// NewBlockWashInfoTypeWithDefaults instantiates a new BlockWashInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockWashInfoTypeWithDefaults() *BlockWashInfoType {
	this := BlockWashInfoType{}
	return &this
}

// GetBlockInfo returns the BlockInfo field value if set, zero value otherwise.
func (o *BlockWashInfoType) GetBlockInfo() BlockWashInfoTypeBlockInfo {
	if o == nil || IsNil(o.BlockInfo) {
		var ret BlockWashInfoTypeBlockInfo
		return ret
	}
	return *o.BlockInfo
}

// GetBlockInfoOk returns a tuple with the BlockInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockWashInfoType) GetBlockInfoOk() (*BlockWashInfoTypeBlockInfo, bool) {
	if o == nil || IsNil(o.BlockInfo) {
		return nil, false
	}
	return o.BlockInfo, true
}

// HasBlockInfo returns a boolean if a field has been set.
func (o *BlockWashInfoType) HasBlockInfo() bool {
	if o != nil && !IsNil(o.BlockInfo) {
		return true
	}

	return false
}

// SetBlockInfo gets a reference to the given BlockWashInfoTypeBlockInfo and assigns it to the BlockInfo field.
func (o *BlockWashInfoType) SetBlockInfo(v BlockWashInfoTypeBlockInfo) {
	o.BlockInfo = &v
}

// GetRoomTypes returns the RoomTypes field value if set, zero value otherwise.
func (o *BlockWashInfoType) GetRoomTypes() []string {
	if o == nil || IsNil(o.RoomTypes) {
		var ret []string
		return ret
	}
	return o.RoomTypes
}

// GetRoomTypesOk returns a tuple with the RoomTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockWashInfoType) GetRoomTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.RoomTypes) {
		return nil, false
	}
	return o.RoomTypes, true
}

// HasRoomTypes returns a boolean if a field has been set.
func (o *BlockWashInfoType) HasRoomTypes() bool {
	if o != nil && !IsNil(o.RoomTypes) {
		return true
	}

	return false
}

// SetRoomTypes gets a reference to the given []string and assigns it to the RoomTypes field.
func (o *BlockWashInfoType) SetRoomTypes(v []string) {
	o.RoomTypes = v
}

func (o BlockWashInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockWashInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockInfo) {
		toSerialize["blockInfo"] = o.BlockInfo
	}
	if !IsNil(o.RoomTypes) {
		toSerialize["roomTypes"] = o.RoomTypes
	}
	return toSerialize, nil
}

type NullableBlockWashInfoType struct {
	value *BlockWashInfoType
	isSet bool
}

func (v NullableBlockWashInfoType) Get() *BlockWashInfoType {
	return v.value
}

func (v *NullableBlockWashInfoType) Set(val *BlockWashInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockWashInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockWashInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockWashInfoType(val *BlockWashInfoType) *NullableBlockWashInfoType {
	return &NullableBlockWashInfoType{value: val, isSet: true}
}

func (v NullableBlockWashInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockWashInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


