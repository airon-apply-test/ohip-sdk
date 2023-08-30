/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the ReservationBlockType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationBlockType{}

// ReservationBlockType Key information about the block for a reservation.
type ReservationBlockType struct {
	// Unique Id that references an object uniquely in the system.
	BlockIdList []UniqueIDType `json:"blockIdList,omitempty"`
	// The Name of the block that is attached to the reservation.
	BlockName *string `json:"blockName,omitempty"`
	// This is the HotelCode of the Block.
	HotelId *string `json:"hotelId,omitempty"`
}

// NewReservationBlockType instantiates a new ReservationBlockType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationBlockType() *ReservationBlockType {
	this := ReservationBlockType{}
	return &this
}

// NewReservationBlockTypeWithDefaults instantiates a new ReservationBlockType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationBlockTypeWithDefaults() *ReservationBlockType {
	this := ReservationBlockType{}
	return &this
}

// GetBlockIdList returns the BlockIdList field value if set, zero value otherwise.
func (o *ReservationBlockType) GetBlockIdList() []UniqueIDType {
	if o == nil || IsNil(o.BlockIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.BlockIdList
}

// GetBlockIdListOk returns a tuple with the BlockIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationBlockType) GetBlockIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.BlockIdList) {
		return nil, false
	}
	return o.BlockIdList, true
}

// HasBlockIdList returns a boolean if a field has been set.
func (o *ReservationBlockType) HasBlockIdList() bool {
	if o != nil && !IsNil(o.BlockIdList) {
		return true
	}

	return false
}

// SetBlockIdList gets a reference to the given []UniqueIDType and assigns it to the BlockIdList field.
func (o *ReservationBlockType) SetBlockIdList(v []UniqueIDType) {
	o.BlockIdList = v
}

// GetBlockName returns the BlockName field value if set, zero value otherwise.
func (o *ReservationBlockType) GetBlockName() string {
	if o == nil || IsNil(o.BlockName) {
		var ret string
		return ret
	}
	return *o.BlockName
}

// GetBlockNameOk returns a tuple with the BlockName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationBlockType) GetBlockNameOk() (*string, bool) {
	if o == nil || IsNil(o.BlockName) {
		return nil, false
	}
	return o.BlockName, true
}

// HasBlockName returns a boolean if a field has been set.
func (o *ReservationBlockType) HasBlockName() bool {
	if o != nil && !IsNil(o.BlockName) {
		return true
	}

	return false
}

// SetBlockName gets a reference to the given string and assigns it to the BlockName field.
func (o *ReservationBlockType) SetBlockName(v string) {
	o.BlockName = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ReservationBlockType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationBlockType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ReservationBlockType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ReservationBlockType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o ReservationBlockType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationBlockType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockIdList) {
		toSerialize["blockIdList"] = o.BlockIdList
	}
	if !IsNil(o.BlockName) {
		toSerialize["blockName"] = o.BlockName
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableReservationBlockType struct {
	value *ReservationBlockType
	isSet bool
}

func (v NullableReservationBlockType) Get() *ReservationBlockType {
	return v.value
}

func (v *NullableReservationBlockType) Set(val *ReservationBlockType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationBlockType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationBlockType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationBlockType(val *ReservationBlockType) *NullableReservationBlockType {
	return &NullableReservationBlockType{value: val, isSet: true}
}

func (v NullableReservationBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationBlockType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


