/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
)

// checks if the ChannelInventorySnapshotInventoriesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelInventorySnapshotInventoriesType{}

// ChannelInventorySnapshotInventoriesType Collection of inventory statistics and restrictions for a hotel in relation to booking channels.
type ChannelInventorySnapshotInventoriesType struct {
	// Details of inventory statistics and restrictions.
	Inventory []ChannelInventoryType `json:"inventory,omitempty"`
	// Hotel that the inventory belongs to.
	HotelId *string `json:"hotelId,omitempty"`
}

// NewChannelInventorySnapshotInventoriesType instantiates a new ChannelInventorySnapshotInventoriesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelInventorySnapshotInventoriesType() *ChannelInventorySnapshotInventoriesType {
	this := ChannelInventorySnapshotInventoriesType{}
	return &this
}

// NewChannelInventorySnapshotInventoriesTypeWithDefaults instantiates a new ChannelInventorySnapshotInventoriesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelInventorySnapshotInventoriesTypeWithDefaults() *ChannelInventorySnapshotInventoriesType {
	this := ChannelInventorySnapshotInventoriesType{}
	return &this
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *ChannelInventorySnapshotInventoriesType) GetInventory() []ChannelInventoryType {
	if o == nil || IsNil(o.Inventory) {
		var ret []ChannelInventoryType
		return ret
	}
	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelInventorySnapshotInventoriesType) GetInventoryOk() ([]ChannelInventoryType, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *ChannelInventorySnapshotInventoriesType) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given []ChannelInventoryType and assigns it to the Inventory field.
func (o *ChannelInventorySnapshotInventoriesType) SetInventory(v []ChannelInventoryType) {
	o.Inventory = v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ChannelInventorySnapshotInventoriesType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelInventorySnapshotInventoriesType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ChannelInventorySnapshotInventoriesType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ChannelInventorySnapshotInventoriesType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o ChannelInventorySnapshotInventoriesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelInventorySnapshotInventoriesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableChannelInventorySnapshotInventoriesType struct {
	value *ChannelInventorySnapshotInventoriesType
	isSet bool
}

func (v NullableChannelInventorySnapshotInventoriesType) Get() *ChannelInventorySnapshotInventoriesType {
	return v.value
}

func (v *NullableChannelInventorySnapshotInventoriesType) Set(val *ChannelInventorySnapshotInventoriesType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelInventorySnapshotInventoriesType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelInventorySnapshotInventoriesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelInventorySnapshotInventoriesType(val *ChannelInventorySnapshotInventoriesType) *NullableChannelInventorySnapshotInventoriesType {
	return &NullableChannelInventorySnapshotInventoriesType{value: val, isSet: true}
}

func (v NullableChannelInventorySnapshotInventoriesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelInventorySnapshotInventoriesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


