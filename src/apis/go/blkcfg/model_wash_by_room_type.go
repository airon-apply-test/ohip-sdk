/*
OPERA Cloud Block Configuration API

APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blkcfg

import (
	"encoding/json"
)

// checks if the WashByRoomType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WashByRoomType{}

// WashByRoomType A representation of the single entry information for 'Wash by number of rooms' type.
type WashByRoomType struct {
	// Contains the number of 'Occupancy1'. Based on the parameter it should be replaced by the number of 'Rooms'.
	Occupancy1 *int32 `json:"occupancy1,omitempty"`
	// Contains the number of 'Occupancy2'.
	Occupancy2 *int32 `json:"occupancy2,omitempty"`
	// Contains the number of 'Occupancy3'.
	Occupancy3 *int32 `json:"occupancy3,omitempty"`
	// Contains the number of 'Occupancy4'.
	Occupancy4 *int32 `json:"occupancy4,omitempty"`
	// Contains the available Sell Limit rooms for a Sell Limit block.
	Sell *int32 `json:"sell,omitempty"`
}

// NewWashByRoomType instantiates a new WashByRoomType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWashByRoomType() *WashByRoomType {
	this := WashByRoomType{}
	return &this
}

// NewWashByRoomTypeWithDefaults instantiates a new WashByRoomType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWashByRoomTypeWithDefaults() *WashByRoomType {
	this := WashByRoomType{}
	return &this
}

// GetOccupancy1 returns the Occupancy1 field value if set, zero value otherwise.
func (o *WashByRoomType) GetOccupancy1() int32 {
	if o == nil || IsNil(o.Occupancy1) {
		var ret int32
		return ret
	}
	return *o.Occupancy1
}

// GetOccupancy1Ok returns a tuple with the Occupancy1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WashByRoomType) GetOccupancy1Ok() (*int32, bool) {
	if o == nil || IsNil(o.Occupancy1) {
		return nil, false
	}
	return o.Occupancy1, true
}

// HasOccupancy1 returns a boolean if a field has been set.
func (o *WashByRoomType) HasOccupancy1() bool {
	if o != nil && !IsNil(o.Occupancy1) {
		return true
	}

	return false
}

// SetOccupancy1 gets a reference to the given int32 and assigns it to the Occupancy1 field.
func (o *WashByRoomType) SetOccupancy1(v int32) {
	o.Occupancy1 = &v
}

// GetOccupancy2 returns the Occupancy2 field value if set, zero value otherwise.
func (o *WashByRoomType) GetOccupancy2() int32 {
	if o == nil || IsNil(o.Occupancy2) {
		var ret int32
		return ret
	}
	return *o.Occupancy2
}

// GetOccupancy2Ok returns a tuple with the Occupancy2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WashByRoomType) GetOccupancy2Ok() (*int32, bool) {
	if o == nil || IsNil(o.Occupancy2) {
		return nil, false
	}
	return o.Occupancy2, true
}

// HasOccupancy2 returns a boolean if a field has been set.
func (o *WashByRoomType) HasOccupancy2() bool {
	if o != nil && !IsNil(o.Occupancy2) {
		return true
	}

	return false
}

// SetOccupancy2 gets a reference to the given int32 and assigns it to the Occupancy2 field.
func (o *WashByRoomType) SetOccupancy2(v int32) {
	o.Occupancy2 = &v
}

// GetOccupancy3 returns the Occupancy3 field value if set, zero value otherwise.
func (o *WashByRoomType) GetOccupancy3() int32 {
	if o == nil || IsNil(o.Occupancy3) {
		var ret int32
		return ret
	}
	return *o.Occupancy3
}

// GetOccupancy3Ok returns a tuple with the Occupancy3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WashByRoomType) GetOccupancy3Ok() (*int32, bool) {
	if o == nil || IsNil(o.Occupancy3) {
		return nil, false
	}
	return o.Occupancy3, true
}

// HasOccupancy3 returns a boolean if a field has been set.
func (o *WashByRoomType) HasOccupancy3() bool {
	if o != nil && !IsNil(o.Occupancy3) {
		return true
	}

	return false
}

// SetOccupancy3 gets a reference to the given int32 and assigns it to the Occupancy3 field.
func (o *WashByRoomType) SetOccupancy3(v int32) {
	o.Occupancy3 = &v
}

// GetOccupancy4 returns the Occupancy4 field value if set, zero value otherwise.
func (o *WashByRoomType) GetOccupancy4() int32 {
	if o == nil || IsNil(o.Occupancy4) {
		var ret int32
		return ret
	}
	return *o.Occupancy4
}

// GetOccupancy4Ok returns a tuple with the Occupancy4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WashByRoomType) GetOccupancy4Ok() (*int32, bool) {
	if o == nil || IsNil(o.Occupancy4) {
		return nil, false
	}
	return o.Occupancy4, true
}

// HasOccupancy4 returns a boolean if a field has been set.
func (o *WashByRoomType) HasOccupancy4() bool {
	if o != nil && !IsNil(o.Occupancy4) {
		return true
	}

	return false
}

// SetOccupancy4 gets a reference to the given int32 and assigns it to the Occupancy4 field.
func (o *WashByRoomType) SetOccupancy4(v int32) {
	o.Occupancy4 = &v
}

// GetSell returns the Sell field value if set, zero value otherwise.
func (o *WashByRoomType) GetSell() int32 {
	if o == nil || IsNil(o.Sell) {
		var ret int32
		return ret
	}
	return *o.Sell
}

// GetSellOk returns a tuple with the Sell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WashByRoomType) GetSellOk() (*int32, bool) {
	if o == nil || IsNil(o.Sell) {
		return nil, false
	}
	return o.Sell, true
}

// HasSell returns a boolean if a field has been set.
func (o *WashByRoomType) HasSell() bool {
	if o != nil && !IsNil(o.Sell) {
		return true
	}

	return false
}

// SetSell gets a reference to the given int32 and assigns it to the Sell field.
func (o *WashByRoomType) SetSell(v int32) {
	o.Sell = &v
}

func (o WashByRoomType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WashByRoomType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Occupancy1) {
		toSerialize["occupancy1"] = o.Occupancy1
	}
	if !IsNil(o.Occupancy2) {
		toSerialize["occupancy2"] = o.Occupancy2
	}
	if !IsNil(o.Occupancy3) {
		toSerialize["occupancy3"] = o.Occupancy3
	}
	if !IsNil(o.Occupancy4) {
		toSerialize["occupancy4"] = o.Occupancy4
	}
	if !IsNil(o.Sell) {
		toSerialize["sell"] = o.Sell
	}
	return toSerialize, nil
}

type NullableWashByRoomType struct {
	value *WashByRoomType
	isSet bool
}

func (v NullableWashByRoomType) Get() *WashByRoomType {
	return v.value
}

func (v *NullableWashByRoomType) Set(val *WashByRoomType) {
	v.value = val
	v.isSet = true
}

func (v NullableWashByRoomType) IsSet() bool {
	return v.isSet
}

func (v *NullableWashByRoomType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWashByRoomType(val *WashByRoomType) *NullableWashByRoomType {
	return &NullableWashByRoomType{value: val, isSet: true}
}

func (v NullableWashByRoomType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWashByRoomType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


