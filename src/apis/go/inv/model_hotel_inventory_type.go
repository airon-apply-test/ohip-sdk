/*
OPERA Cloud Inventory API

APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inv

import (
	"encoding/json"
)

// checks if the HotelInventoryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelInventoryType{}

// HotelInventoryType A collection of Reservation objects and Unique IDs of Reservation.
type HotelInventoryType struct {
	// Collection of Inventory counts for the date ranges.
	HouseInventory []InventoryCountsType `json:"houseInventory,omitempty"`
	// Collection of Room type level Inventory counts for the date ranges.
	RoomTypeInventories []InventoryLevelCountsListType `json:"roomTypeInventories,omitempty"`
	// Collection of Room Class level Inventory counts for the date ranges.
	RoomClassInventories []InventoryLevelCountsListType `json:"roomClassInventories,omitempty"`
}

// NewHotelInventoryType instantiates a new HotelInventoryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelInventoryType() *HotelInventoryType {
	this := HotelInventoryType{}
	return &this
}

// NewHotelInventoryTypeWithDefaults instantiates a new HotelInventoryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelInventoryTypeWithDefaults() *HotelInventoryType {
	this := HotelInventoryType{}
	return &this
}

// GetHouseInventory returns the HouseInventory field value if set, zero value otherwise.
func (o *HotelInventoryType) GetHouseInventory() []InventoryCountsType {
	if o == nil || IsNil(o.HouseInventory) {
		var ret []InventoryCountsType
		return ret
	}
	return o.HouseInventory
}

// GetHouseInventoryOk returns a tuple with the HouseInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInventoryType) GetHouseInventoryOk() ([]InventoryCountsType, bool) {
	if o == nil || IsNil(o.HouseInventory) {
		return nil, false
	}
	return o.HouseInventory, true
}

// HasHouseInventory returns a boolean if a field has been set.
func (o *HotelInventoryType) HasHouseInventory() bool {
	if o != nil && !IsNil(o.HouseInventory) {
		return true
	}

	return false
}

// SetHouseInventory gets a reference to the given []InventoryCountsType and assigns it to the HouseInventory field.
func (o *HotelInventoryType) SetHouseInventory(v []InventoryCountsType) {
	o.HouseInventory = v
}

// GetRoomTypeInventories returns the RoomTypeInventories field value if set, zero value otherwise.
func (o *HotelInventoryType) GetRoomTypeInventories() []InventoryLevelCountsListType {
	if o == nil || IsNil(o.RoomTypeInventories) {
		var ret []InventoryLevelCountsListType
		return ret
	}
	return o.RoomTypeInventories
}

// GetRoomTypeInventoriesOk returns a tuple with the RoomTypeInventories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInventoryType) GetRoomTypeInventoriesOk() ([]InventoryLevelCountsListType, bool) {
	if o == nil || IsNil(o.RoomTypeInventories) {
		return nil, false
	}
	return o.RoomTypeInventories, true
}

// HasRoomTypeInventories returns a boolean if a field has been set.
func (o *HotelInventoryType) HasRoomTypeInventories() bool {
	if o != nil && !IsNil(o.RoomTypeInventories) {
		return true
	}

	return false
}

// SetRoomTypeInventories gets a reference to the given []InventoryLevelCountsListType and assigns it to the RoomTypeInventories field.
func (o *HotelInventoryType) SetRoomTypeInventories(v []InventoryLevelCountsListType) {
	o.RoomTypeInventories = v
}

// GetRoomClassInventories returns the RoomClassInventories field value if set, zero value otherwise.
func (o *HotelInventoryType) GetRoomClassInventories() []InventoryLevelCountsListType {
	if o == nil || IsNil(o.RoomClassInventories) {
		var ret []InventoryLevelCountsListType
		return ret
	}
	return o.RoomClassInventories
}

// GetRoomClassInventoriesOk returns a tuple with the RoomClassInventories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInventoryType) GetRoomClassInventoriesOk() ([]InventoryLevelCountsListType, bool) {
	if o == nil || IsNil(o.RoomClassInventories) {
		return nil, false
	}
	return o.RoomClassInventories, true
}

// HasRoomClassInventories returns a boolean if a field has been set.
func (o *HotelInventoryType) HasRoomClassInventories() bool {
	if o != nil && !IsNil(o.RoomClassInventories) {
		return true
	}

	return false
}

// SetRoomClassInventories gets a reference to the given []InventoryLevelCountsListType and assigns it to the RoomClassInventories field.
func (o *HotelInventoryType) SetRoomClassInventories(v []InventoryLevelCountsListType) {
	o.RoomClassInventories = v
}

func (o HotelInventoryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelInventoryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HouseInventory) {
		toSerialize["houseInventory"] = o.HouseInventory
	}
	if !IsNil(o.RoomTypeInventories) {
		toSerialize["roomTypeInventories"] = o.RoomTypeInventories
	}
	if !IsNil(o.RoomClassInventories) {
		toSerialize["roomClassInventories"] = o.RoomClassInventories
	}
	return toSerialize, nil
}

type NullableHotelInventoryType struct {
	value *HotelInventoryType
	isSet bool
}

func (v NullableHotelInventoryType) Get() *HotelInventoryType {
	return v.value
}

func (v *NullableHotelInventoryType) Set(val *HotelInventoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelInventoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelInventoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelInventoryType(val *HotelInventoryType) *NullableHotelInventoryType {
	return &NullableHotelInventoryType{value: val, isSet: true}
}

func (v NullableHotelInventoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelInventoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


