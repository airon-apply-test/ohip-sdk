/*
OPERA Cloud Block Reservation Asynchronous API

APIs to cater Block Reservation related asynchronous functionality in OPERA.<br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.4.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blkasync

import (
	"encoding/json"
)

// checks if the RoomAllocationDetailsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomAllocationDetailsType{}

// RoomAllocationDetailsType struct for RoomAllocationDetailsType
type RoomAllocationDetailsType struct {
	// The total original( forecasted ) rooms for the block.
	OriginalRooms *int32 `json:"originalRooms,omitempty"`
	// The room type for which the allocation details are listed.
	RoomType *string `json:"roomType,omitempty"`
	// The total current( projected + pickup ) rooms for the block.
	CurrentRooms *int32 `json:"currentRooms,omitempty"`
	// The total pickup rooms for the block.
	PickupRooms *int32 `json:"pickupRooms,omitempty"`
	Inventory *BlockGridInvType `json:"inventory,omitempty"`
	Rates *BlockGridRatesType `json:"rates,omitempty"`
	ActualRevenue *BlockActualRevenueType `json:"actualRevenue,omitempty"`
	PotentialRevenue *BlockPotenitalRevenueType `json:"potentialRevenue,omitempty"`
}

// NewRoomAllocationDetailsType instantiates a new RoomAllocationDetailsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomAllocationDetailsType() *RoomAllocationDetailsType {
	this := RoomAllocationDetailsType{}
	return &this
}

// NewRoomAllocationDetailsTypeWithDefaults instantiates a new RoomAllocationDetailsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomAllocationDetailsTypeWithDefaults() *RoomAllocationDetailsType {
	this := RoomAllocationDetailsType{}
	return &this
}

// GetOriginalRooms returns the OriginalRooms field value if set, zero value otherwise.
func (o *RoomAllocationDetailsType) GetOriginalRooms() int32 {
	if o == nil || IsNil(o.OriginalRooms) {
		var ret int32
		return ret
	}
	return *o.OriginalRooms
}

// GetOriginalRoomsOk returns a tuple with the OriginalRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomAllocationDetailsType) GetOriginalRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalRooms) {
		return nil, false
	}
	return o.OriginalRooms, true
}

// HasOriginalRooms returns a boolean if a field has been set.
func (o *RoomAllocationDetailsType) HasOriginalRooms() bool {
	if o != nil && !IsNil(o.OriginalRooms) {
		return true
	}

	return false
}

// SetOriginalRooms gets a reference to the given int32 and assigns it to the OriginalRooms field.
func (o *RoomAllocationDetailsType) SetOriginalRooms(v int32) {
	o.OriginalRooms = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *RoomAllocationDetailsType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomAllocationDetailsType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *RoomAllocationDetailsType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *RoomAllocationDetailsType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetCurrentRooms returns the CurrentRooms field value if set, zero value otherwise.
func (o *RoomAllocationDetailsType) GetCurrentRooms() int32 {
	if o == nil || IsNil(o.CurrentRooms) {
		var ret int32
		return ret
	}
	return *o.CurrentRooms
}

// GetCurrentRoomsOk returns a tuple with the CurrentRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomAllocationDetailsType) GetCurrentRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentRooms) {
		return nil, false
	}
	return o.CurrentRooms, true
}

// HasCurrentRooms returns a boolean if a field has been set.
func (o *RoomAllocationDetailsType) HasCurrentRooms() bool {
	if o != nil && !IsNil(o.CurrentRooms) {
		return true
	}

	return false
}

// SetCurrentRooms gets a reference to the given int32 and assigns it to the CurrentRooms field.
func (o *RoomAllocationDetailsType) SetCurrentRooms(v int32) {
	o.CurrentRooms = &v
}

// GetPickupRooms returns the PickupRooms field value if set, zero value otherwise.
func (o *RoomAllocationDetailsType) GetPickupRooms() int32 {
	if o == nil || IsNil(o.PickupRooms) {
		var ret int32
		return ret
	}
	return *o.PickupRooms
}

// GetPickupRoomsOk returns a tuple with the PickupRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomAllocationDetailsType) GetPickupRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.PickupRooms) {
		return nil, false
	}
	return o.PickupRooms, true
}

// HasPickupRooms returns a boolean if a field has been set.
func (o *RoomAllocationDetailsType) HasPickupRooms() bool {
	if o != nil && !IsNil(o.PickupRooms) {
		return true
	}

	return false
}

// SetPickupRooms gets a reference to the given int32 and assigns it to the PickupRooms field.
func (o *RoomAllocationDetailsType) SetPickupRooms(v int32) {
	o.PickupRooms = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *RoomAllocationDetailsType) GetInventory() BlockGridInvType {
	if o == nil || IsNil(o.Inventory) {
		var ret BlockGridInvType
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomAllocationDetailsType) GetInventoryOk() (*BlockGridInvType, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *RoomAllocationDetailsType) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given BlockGridInvType and assigns it to the Inventory field.
func (o *RoomAllocationDetailsType) SetInventory(v BlockGridInvType) {
	o.Inventory = &v
}

// GetRates returns the Rates field value if set, zero value otherwise.
func (o *RoomAllocationDetailsType) GetRates() BlockGridRatesType {
	if o == nil || IsNil(o.Rates) {
		var ret BlockGridRatesType
		return ret
	}
	return *o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomAllocationDetailsType) GetRatesOk() (*BlockGridRatesType, bool) {
	if o == nil || IsNil(o.Rates) {
		return nil, false
	}
	return o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *RoomAllocationDetailsType) HasRates() bool {
	if o != nil && !IsNil(o.Rates) {
		return true
	}

	return false
}

// SetRates gets a reference to the given BlockGridRatesType and assigns it to the Rates field.
func (o *RoomAllocationDetailsType) SetRates(v BlockGridRatesType) {
	o.Rates = &v
}

// GetActualRevenue returns the ActualRevenue field value if set, zero value otherwise.
func (o *RoomAllocationDetailsType) GetActualRevenue() BlockActualRevenueType {
	if o == nil || IsNil(o.ActualRevenue) {
		var ret BlockActualRevenueType
		return ret
	}
	return *o.ActualRevenue
}

// GetActualRevenueOk returns a tuple with the ActualRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomAllocationDetailsType) GetActualRevenueOk() (*BlockActualRevenueType, bool) {
	if o == nil || IsNil(o.ActualRevenue) {
		return nil, false
	}
	return o.ActualRevenue, true
}

// HasActualRevenue returns a boolean if a field has been set.
func (o *RoomAllocationDetailsType) HasActualRevenue() bool {
	if o != nil && !IsNil(o.ActualRevenue) {
		return true
	}

	return false
}

// SetActualRevenue gets a reference to the given BlockActualRevenueType and assigns it to the ActualRevenue field.
func (o *RoomAllocationDetailsType) SetActualRevenue(v BlockActualRevenueType) {
	o.ActualRevenue = &v
}

// GetPotentialRevenue returns the PotentialRevenue field value if set, zero value otherwise.
func (o *RoomAllocationDetailsType) GetPotentialRevenue() BlockPotenitalRevenueType {
	if o == nil || IsNil(o.PotentialRevenue) {
		var ret BlockPotenitalRevenueType
		return ret
	}
	return *o.PotentialRevenue
}

// GetPotentialRevenueOk returns a tuple with the PotentialRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomAllocationDetailsType) GetPotentialRevenueOk() (*BlockPotenitalRevenueType, bool) {
	if o == nil || IsNil(o.PotentialRevenue) {
		return nil, false
	}
	return o.PotentialRevenue, true
}

// HasPotentialRevenue returns a boolean if a field has been set.
func (o *RoomAllocationDetailsType) HasPotentialRevenue() bool {
	if o != nil && !IsNil(o.PotentialRevenue) {
		return true
	}

	return false
}

// SetPotentialRevenue gets a reference to the given BlockPotenitalRevenueType and assigns it to the PotentialRevenue field.
func (o *RoomAllocationDetailsType) SetPotentialRevenue(v BlockPotenitalRevenueType) {
	o.PotentialRevenue = &v
}

func (o RoomAllocationDetailsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomAllocationDetailsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginalRooms) {
		toSerialize["originalRooms"] = o.OriginalRooms
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.CurrentRooms) {
		toSerialize["currentRooms"] = o.CurrentRooms
	}
	if !IsNil(o.PickupRooms) {
		toSerialize["pickupRooms"] = o.PickupRooms
	}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.Rates) {
		toSerialize["rates"] = o.Rates
	}
	if !IsNil(o.ActualRevenue) {
		toSerialize["actualRevenue"] = o.ActualRevenue
	}
	if !IsNil(o.PotentialRevenue) {
		toSerialize["potentialRevenue"] = o.PotentialRevenue
	}
	return toSerialize, nil
}

type NullableRoomAllocationDetailsType struct {
	value *RoomAllocationDetailsType
	isSet bool
}

func (v NullableRoomAllocationDetailsType) Get() *RoomAllocationDetailsType {
	return v.value
}

func (v *NullableRoomAllocationDetailsType) Set(val *RoomAllocationDetailsType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomAllocationDetailsType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomAllocationDetailsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomAllocationDetailsType(val *RoomAllocationDetailsType) *NullableRoomAllocationDetailsType {
	return &NullableRoomAllocationDetailsType{value: val, isSet: true}
}

func (v NullableRoomAllocationDetailsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomAllocationDetailsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


