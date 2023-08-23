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

// checks if the BlockAvailabilityRoomInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockAvailabilityRoomInfoType{}

// BlockAvailabilityRoomInfoType Captures occupancy and rate information for a room type.
type BlockAvailabilityRoomInfoType struct {
	// Total occupancy alloted for the room type and stay date.
	Inventory *int32 `json:"inventory,omitempty"`
	// The derived rate amount for the room type, stay date and number of adults.
	Rate *float32 `json:"rate,omitempty"`
	RoomType *string `json:"roomType,omitempty"`
	// Available Sell Limit rooms for a Sell Limit block.
	SellLimit *int32 `json:"sellLimit,omitempty"`
}

// NewBlockAvailabilityRoomInfoType instantiates a new BlockAvailabilityRoomInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockAvailabilityRoomInfoType() *BlockAvailabilityRoomInfoType {
	this := BlockAvailabilityRoomInfoType{}
	return &this
}

// NewBlockAvailabilityRoomInfoTypeWithDefaults instantiates a new BlockAvailabilityRoomInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockAvailabilityRoomInfoTypeWithDefaults() *BlockAvailabilityRoomInfoType {
	this := BlockAvailabilityRoomInfoType{}
	return &this
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *BlockAvailabilityRoomInfoType) GetInventory() int32 {
	if o == nil || IsNil(o.Inventory) {
		var ret int32
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAvailabilityRoomInfoType) GetInventoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *BlockAvailabilityRoomInfoType) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given int32 and assigns it to the Inventory field.
func (o *BlockAvailabilityRoomInfoType) SetInventory(v int32) {
	o.Inventory = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *BlockAvailabilityRoomInfoType) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAvailabilityRoomInfoType) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *BlockAvailabilityRoomInfoType) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *BlockAvailabilityRoomInfoType) SetRate(v float32) {
	o.Rate = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *BlockAvailabilityRoomInfoType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAvailabilityRoomInfoType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *BlockAvailabilityRoomInfoType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *BlockAvailabilityRoomInfoType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetSellLimit returns the SellLimit field value if set, zero value otherwise.
func (o *BlockAvailabilityRoomInfoType) GetSellLimit() int32 {
	if o == nil || IsNil(o.SellLimit) {
		var ret int32
		return ret
	}
	return *o.SellLimit
}

// GetSellLimitOk returns a tuple with the SellLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAvailabilityRoomInfoType) GetSellLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.SellLimit) {
		return nil, false
	}
	return o.SellLimit, true
}

// HasSellLimit returns a boolean if a field has been set.
func (o *BlockAvailabilityRoomInfoType) HasSellLimit() bool {
	if o != nil && !IsNil(o.SellLimit) {
		return true
	}

	return false
}

// SetSellLimit gets a reference to the given int32 and assigns it to the SellLimit field.
func (o *BlockAvailabilityRoomInfoType) SetSellLimit(v int32) {
	o.SellLimit = &v
}

func (o BlockAvailabilityRoomInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockAvailabilityRoomInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.SellLimit) {
		toSerialize["sellLimit"] = o.SellLimit
	}
	return toSerialize, nil
}

type NullableBlockAvailabilityRoomInfoType struct {
	value *BlockAvailabilityRoomInfoType
	isSet bool
}

func (v NullableBlockAvailabilityRoomInfoType) Get() *BlockAvailabilityRoomInfoType {
	return v.value
}

func (v *NullableBlockAvailabilityRoomInfoType) Set(val *BlockAvailabilityRoomInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockAvailabilityRoomInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockAvailabilityRoomInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockAvailabilityRoomInfoType(val *BlockAvailabilityRoomInfoType) *NullableBlockAvailabilityRoomInfoType {
	return &NullableBlockAvailabilityRoomInfoType{value: val, isSet: true}
}

func (v NullableBlockAvailabilityRoomInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockAvailabilityRoomInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


