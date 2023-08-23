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

// checks if the BlockDailyStatisticsInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockDailyStatisticsInfoType{}

// BlockDailyStatisticsInfoType Information about Block daily statistics.
type BlockDailyStatisticsInfoType struct {
	// Pertain value for available rooms for a block.
	Available *int32 `json:"available,omitempty"`
	// Pertain value for blocked rooms for a block.
	Allocated *int32 `json:"allocated,omitempty"`
	// Pertain value for reserved rooms for a block.
	Pickup *int32 `json:"pickup,omitempty"`
}

// NewBlockDailyStatisticsInfoType instantiates a new BlockDailyStatisticsInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockDailyStatisticsInfoType() *BlockDailyStatisticsInfoType {
	this := BlockDailyStatisticsInfoType{}
	return &this
}

// NewBlockDailyStatisticsInfoTypeWithDefaults instantiates a new BlockDailyStatisticsInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockDailyStatisticsInfoTypeWithDefaults() *BlockDailyStatisticsInfoType {
	this := BlockDailyStatisticsInfoType{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *BlockDailyStatisticsInfoType) GetAvailable() int32 {
	if o == nil || IsNil(o.Available) {
		var ret int32
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDailyStatisticsInfoType) GetAvailableOk() (*int32, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *BlockDailyStatisticsInfoType) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int32 and assigns it to the Available field.
func (o *BlockDailyStatisticsInfoType) SetAvailable(v int32) {
	o.Available = &v
}

// GetAllocated returns the Allocated field value if set, zero value otherwise.
func (o *BlockDailyStatisticsInfoType) GetAllocated() int32 {
	if o == nil || IsNil(o.Allocated) {
		var ret int32
		return ret
	}
	return *o.Allocated
}

// GetAllocatedOk returns a tuple with the Allocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDailyStatisticsInfoType) GetAllocatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Allocated) {
		return nil, false
	}
	return o.Allocated, true
}

// HasAllocated returns a boolean if a field has been set.
func (o *BlockDailyStatisticsInfoType) HasAllocated() bool {
	if o != nil && !IsNil(o.Allocated) {
		return true
	}

	return false
}

// SetAllocated gets a reference to the given int32 and assigns it to the Allocated field.
func (o *BlockDailyStatisticsInfoType) SetAllocated(v int32) {
	o.Allocated = &v
}

// GetPickup returns the Pickup field value if set, zero value otherwise.
func (o *BlockDailyStatisticsInfoType) GetPickup() int32 {
	if o == nil || IsNil(o.Pickup) {
		var ret int32
		return ret
	}
	return *o.Pickup
}

// GetPickupOk returns a tuple with the Pickup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDailyStatisticsInfoType) GetPickupOk() (*int32, bool) {
	if o == nil || IsNil(o.Pickup) {
		return nil, false
	}
	return o.Pickup, true
}

// HasPickup returns a boolean if a field has been set.
func (o *BlockDailyStatisticsInfoType) HasPickup() bool {
	if o != nil && !IsNil(o.Pickup) {
		return true
	}

	return false
}

// SetPickup gets a reference to the given int32 and assigns it to the Pickup field.
func (o *BlockDailyStatisticsInfoType) SetPickup(v int32) {
	o.Pickup = &v
}

func (o BlockDailyStatisticsInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockDailyStatisticsInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Allocated) {
		toSerialize["allocated"] = o.Allocated
	}
	if !IsNil(o.Pickup) {
		toSerialize["pickup"] = o.Pickup
	}
	return toSerialize, nil
}

type NullableBlockDailyStatisticsInfoType struct {
	value *BlockDailyStatisticsInfoType
	isSet bool
}

func (v NullableBlockDailyStatisticsInfoType) Get() *BlockDailyStatisticsInfoType {
	return v.value
}

func (v *NullableBlockDailyStatisticsInfoType) Set(val *BlockDailyStatisticsInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockDailyStatisticsInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockDailyStatisticsInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockDailyStatisticsInfoType(val *BlockDailyStatisticsInfoType) *NullableBlockDailyStatisticsInfoType {
	return &NullableBlockDailyStatisticsInfoType{value: val, isSet: true}
}

func (v NullableBlockDailyStatisticsInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockDailyStatisticsInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


