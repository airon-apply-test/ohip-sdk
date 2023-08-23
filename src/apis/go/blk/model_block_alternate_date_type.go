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

// checks if the BlockAlternateDateType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockAlternateDateType{}

// BlockAlternateDateType Information for Alternate date types
type BlockAlternateDateType struct {
	// Alternate Begin date.
	AlternateBeginDate *string `json:"alternateBeginDate,omitempty"`
	BlockRates *BlockGridRatesType `json:"blockRates,omitempty"`
	// The Room Category Label
	RoomCategory *string `json:"roomCategory,omitempty"`
	// The Priority given to these alternate dates.
	Priority *int32 `json:"priority,omitempty"`
}

// NewBlockAlternateDateType instantiates a new BlockAlternateDateType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockAlternateDateType() *BlockAlternateDateType {
	this := BlockAlternateDateType{}
	return &this
}

// NewBlockAlternateDateTypeWithDefaults instantiates a new BlockAlternateDateType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockAlternateDateTypeWithDefaults() *BlockAlternateDateType {
	this := BlockAlternateDateType{}
	return &this
}

// GetAlternateBeginDate returns the AlternateBeginDate field value if set, zero value otherwise.
func (o *BlockAlternateDateType) GetAlternateBeginDate() string {
	if o == nil || IsNil(o.AlternateBeginDate) {
		var ret string
		return ret
	}
	return *o.AlternateBeginDate
}

// GetAlternateBeginDateOk returns a tuple with the AlternateBeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAlternateDateType) GetAlternateBeginDateOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateBeginDate) {
		return nil, false
	}
	return o.AlternateBeginDate, true
}

// HasAlternateBeginDate returns a boolean if a field has been set.
func (o *BlockAlternateDateType) HasAlternateBeginDate() bool {
	if o != nil && !IsNil(o.AlternateBeginDate) {
		return true
	}

	return false
}

// SetAlternateBeginDate gets a reference to the given string and assigns it to the AlternateBeginDate field.
func (o *BlockAlternateDateType) SetAlternateBeginDate(v string) {
	o.AlternateBeginDate = &v
}

// GetBlockRates returns the BlockRates field value if set, zero value otherwise.
func (o *BlockAlternateDateType) GetBlockRates() BlockGridRatesType {
	if o == nil || IsNil(o.BlockRates) {
		var ret BlockGridRatesType
		return ret
	}
	return *o.BlockRates
}

// GetBlockRatesOk returns a tuple with the BlockRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAlternateDateType) GetBlockRatesOk() (*BlockGridRatesType, bool) {
	if o == nil || IsNil(o.BlockRates) {
		return nil, false
	}
	return o.BlockRates, true
}

// HasBlockRates returns a boolean if a field has been set.
func (o *BlockAlternateDateType) HasBlockRates() bool {
	if o != nil && !IsNil(o.BlockRates) {
		return true
	}

	return false
}

// SetBlockRates gets a reference to the given BlockGridRatesType and assigns it to the BlockRates field.
func (o *BlockAlternateDateType) SetBlockRates(v BlockGridRatesType) {
	o.BlockRates = &v
}

// GetRoomCategory returns the RoomCategory field value if set, zero value otherwise.
func (o *BlockAlternateDateType) GetRoomCategory() string {
	if o == nil || IsNil(o.RoomCategory) {
		var ret string
		return ret
	}
	return *o.RoomCategory
}

// GetRoomCategoryOk returns a tuple with the RoomCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAlternateDateType) GetRoomCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.RoomCategory) {
		return nil, false
	}
	return o.RoomCategory, true
}

// HasRoomCategory returns a boolean if a field has been set.
func (o *BlockAlternateDateType) HasRoomCategory() bool {
	if o != nil && !IsNil(o.RoomCategory) {
		return true
	}

	return false
}

// SetRoomCategory gets a reference to the given string and assigns it to the RoomCategory field.
func (o *BlockAlternateDateType) SetRoomCategory(v string) {
	o.RoomCategory = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *BlockAlternateDateType) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAlternateDateType) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *BlockAlternateDateType) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *BlockAlternateDateType) SetPriority(v int32) {
	o.Priority = &v
}

func (o BlockAlternateDateType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockAlternateDateType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlternateBeginDate) {
		toSerialize["alternateBeginDate"] = o.AlternateBeginDate
	}
	if !IsNil(o.BlockRates) {
		toSerialize["blockRates"] = o.BlockRates
	}
	if !IsNil(o.RoomCategory) {
		toSerialize["roomCategory"] = o.RoomCategory
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return toSerialize, nil
}

type NullableBlockAlternateDateType struct {
	value *BlockAlternateDateType
	isSet bool
}

func (v NullableBlockAlternateDateType) Get() *BlockAlternateDateType {
	return v.value
}

func (v *NullableBlockAlternateDateType) Set(val *BlockAlternateDateType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockAlternateDateType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockAlternateDateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockAlternateDateType(val *BlockAlternateDateType) *NullableBlockAlternateDateType {
	return &NullableBlockAlternateDateType{value: val, isSet: true}
}

func (v NullableBlockAlternateDateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockAlternateDateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


