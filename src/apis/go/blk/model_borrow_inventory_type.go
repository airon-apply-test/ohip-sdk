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

// checks if the BorrowInventoryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BorrowInventoryType{}

// BorrowInventoryType This holds the number of rooms for each date that will be borrowed from the candidate room type provided.
type BorrowInventoryType struct {
	BlockId *BlockId `json:"blockId,omitempty"`
	// The date on which rooms need to be borrowed for the block either from another room type in the block or from House if the block is elastic.
	BorrowDate *string `json:"borrowDate,omitempty"`
	// Specifies the number of rooms to be borrowed from the room type or House.
	BorrowRooms []BorrowRoomType `json:"borrowRooms,omitempty"`
}

// NewBorrowInventoryType instantiates a new BorrowInventoryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBorrowInventoryType() *BorrowInventoryType {
	this := BorrowInventoryType{}
	return &this
}

// NewBorrowInventoryTypeWithDefaults instantiates a new BorrowInventoryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBorrowInventoryTypeWithDefaults() *BorrowInventoryType {
	this := BorrowInventoryType{}
	return &this
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *BorrowInventoryType) GetBlockId() BlockId {
	if o == nil || IsNil(o.BlockId) {
		var ret BlockId
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowInventoryType) GetBlockIdOk() (*BlockId, bool) {
	if o == nil || IsNil(o.BlockId) {
		return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *BorrowInventoryType) HasBlockId() bool {
	if o != nil && !IsNil(o.BlockId) {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given BlockId and assigns it to the BlockId field.
func (o *BorrowInventoryType) SetBlockId(v BlockId) {
	o.BlockId = &v
}

// GetBorrowDate returns the BorrowDate field value if set, zero value otherwise.
func (o *BorrowInventoryType) GetBorrowDate() string {
	if o == nil || IsNil(o.BorrowDate) {
		var ret string
		return ret
	}
	return *o.BorrowDate
}

// GetBorrowDateOk returns a tuple with the BorrowDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowInventoryType) GetBorrowDateOk() (*string, bool) {
	if o == nil || IsNil(o.BorrowDate) {
		return nil, false
	}
	return o.BorrowDate, true
}

// HasBorrowDate returns a boolean if a field has been set.
func (o *BorrowInventoryType) HasBorrowDate() bool {
	if o != nil && !IsNil(o.BorrowDate) {
		return true
	}

	return false
}

// SetBorrowDate gets a reference to the given string and assigns it to the BorrowDate field.
func (o *BorrowInventoryType) SetBorrowDate(v string) {
	o.BorrowDate = &v
}

// GetBorrowRooms returns the BorrowRooms field value if set, zero value otherwise.
func (o *BorrowInventoryType) GetBorrowRooms() []BorrowRoomType {
	if o == nil || IsNil(o.BorrowRooms) {
		var ret []BorrowRoomType
		return ret
	}
	return o.BorrowRooms
}

// GetBorrowRoomsOk returns a tuple with the BorrowRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowInventoryType) GetBorrowRoomsOk() ([]BorrowRoomType, bool) {
	if o == nil || IsNil(o.BorrowRooms) {
		return nil, false
	}
	return o.BorrowRooms, true
}

// HasBorrowRooms returns a boolean if a field has been set.
func (o *BorrowInventoryType) HasBorrowRooms() bool {
	if o != nil && !IsNil(o.BorrowRooms) {
		return true
	}

	return false
}

// SetBorrowRooms gets a reference to the given []BorrowRoomType and assigns it to the BorrowRooms field.
func (o *BorrowInventoryType) SetBorrowRooms(v []BorrowRoomType) {
	o.BorrowRooms = v
}

func (o BorrowInventoryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BorrowInventoryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockId) {
		toSerialize["blockId"] = o.BlockId
	}
	if !IsNil(o.BorrowDate) {
		toSerialize["borrowDate"] = o.BorrowDate
	}
	if !IsNil(o.BorrowRooms) {
		toSerialize["borrowRooms"] = o.BorrowRooms
	}
	return toSerialize, nil
}

type NullableBorrowInventoryType struct {
	value *BorrowInventoryType
	isSet bool
}

func (v NullableBorrowInventoryType) Get() *BorrowInventoryType {
	return v.value
}

func (v *NullableBorrowInventoryType) Set(val *BorrowInventoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableBorrowInventoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableBorrowInventoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBorrowInventoryType(val *BorrowInventoryType) *NullableBorrowInventoryType {
	return &NullableBorrowInventoryType{value: val, isSet: true}
}

func (v NullableBorrowInventoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBorrowInventoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


