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

// checks if the BorrowableInventoryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BorrowableInventoryType{}

// BorrowableInventoryType This provides a collection of all borrowable room types by date.
type BorrowableInventoryType struct {
	// Returning an empty element of this type indicates the this is a House type. This is used in conjunction with the Borrow operations for Blocks where rooms are to be borrowed from House.
	House map[string]interface{} `json:"house,omitempty"`
	// The room type that contains rooms that can be borrowed.
	RoomType *string `json:"roomType,omitempty"`
	// The number of rooms that are available to be borrowed from the room type or house.
	AvailableRooms *int32 `json:"availableRooms,omitempty"`
	// The number of sell limit rooms that are available for the room type.
	AvailableSellLimit *int32 `json:"availableSellLimit,omitempty"`
	// An error that occurred during the processing of a message.
	Errors []ErrorType `json:"errors,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warning []WarningType `json:"warning,omitempty"`
}

// NewBorrowableInventoryType instantiates a new BorrowableInventoryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBorrowableInventoryType() *BorrowableInventoryType {
	this := BorrowableInventoryType{}
	return &this
}

// NewBorrowableInventoryTypeWithDefaults instantiates a new BorrowableInventoryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBorrowableInventoryTypeWithDefaults() *BorrowableInventoryType {
	this := BorrowableInventoryType{}
	return &this
}

// GetHouse returns the House field value if set, zero value otherwise.
func (o *BorrowableInventoryType) GetHouse() map[string]interface{} {
	if o == nil || IsNil(o.House) {
		var ret map[string]interface{}
		return ret
	}
	return o.House
}

// GetHouseOk returns a tuple with the House field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowableInventoryType) GetHouseOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.House) {
		return map[string]interface{}{}, false
	}
	return o.House, true
}

// HasHouse returns a boolean if a field has been set.
func (o *BorrowableInventoryType) HasHouse() bool {
	if o != nil && !IsNil(o.House) {
		return true
	}

	return false
}

// SetHouse gets a reference to the given map[string]interface{} and assigns it to the House field.
func (o *BorrowableInventoryType) SetHouse(v map[string]interface{}) {
	o.House = v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *BorrowableInventoryType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowableInventoryType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *BorrowableInventoryType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *BorrowableInventoryType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetAvailableRooms returns the AvailableRooms field value if set, zero value otherwise.
func (o *BorrowableInventoryType) GetAvailableRooms() int32 {
	if o == nil || IsNil(o.AvailableRooms) {
		var ret int32
		return ret
	}
	return *o.AvailableRooms
}

// GetAvailableRoomsOk returns a tuple with the AvailableRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowableInventoryType) GetAvailableRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableRooms) {
		return nil, false
	}
	return o.AvailableRooms, true
}

// HasAvailableRooms returns a boolean if a field has been set.
func (o *BorrowableInventoryType) HasAvailableRooms() bool {
	if o != nil && !IsNil(o.AvailableRooms) {
		return true
	}

	return false
}

// SetAvailableRooms gets a reference to the given int32 and assigns it to the AvailableRooms field.
func (o *BorrowableInventoryType) SetAvailableRooms(v int32) {
	o.AvailableRooms = &v
}

// GetAvailableSellLimit returns the AvailableSellLimit field value if set, zero value otherwise.
func (o *BorrowableInventoryType) GetAvailableSellLimit() int32 {
	if o == nil || IsNil(o.AvailableSellLimit) {
		var ret int32
		return ret
	}
	return *o.AvailableSellLimit
}

// GetAvailableSellLimitOk returns a tuple with the AvailableSellLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowableInventoryType) GetAvailableSellLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableSellLimit) {
		return nil, false
	}
	return o.AvailableSellLimit, true
}

// HasAvailableSellLimit returns a boolean if a field has been set.
func (o *BorrowableInventoryType) HasAvailableSellLimit() bool {
	if o != nil && !IsNil(o.AvailableSellLimit) {
		return true
	}

	return false
}

// SetAvailableSellLimit gets a reference to the given int32 and assigns it to the AvailableSellLimit field.
func (o *BorrowableInventoryType) SetAvailableSellLimit(v int32) {
	o.AvailableSellLimit = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BorrowableInventoryType) GetErrors() []ErrorType {
	if o == nil || IsNil(o.Errors) {
		var ret []ErrorType
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowableInventoryType) GetErrorsOk() ([]ErrorType, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BorrowableInventoryType) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorType and assigns it to the Errors field.
func (o *BorrowableInventoryType) SetErrors(v []ErrorType) {
	o.Errors = v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *BorrowableInventoryType) GetWarning() []WarningType {
	if o == nil || IsNil(o.Warning) {
		var ret []WarningType
		return ret
	}
	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowableInventoryType) GetWarningOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *BorrowableInventoryType) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given []WarningType and assigns it to the Warning field.
func (o *BorrowableInventoryType) SetWarning(v []WarningType) {
	o.Warning = v
}

func (o BorrowableInventoryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BorrowableInventoryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.House) {
		toSerialize["house"] = o.House
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.AvailableRooms) {
		toSerialize["availableRooms"] = o.AvailableRooms
	}
	if !IsNil(o.AvailableSellLimit) {
		toSerialize["availableSellLimit"] = o.AvailableSellLimit
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	return toSerialize, nil
}

type NullableBorrowableInventoryType struct {
	value *BorrowableInventoryType
	isSet bool
}

func (v NullableBorrowableInventoryType) Get() *BorrowableInventoryType {
	return v.value
}

func (v *NullableBorrowableInventoryType) Set(val *BorrowableInventoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableBorrowableInventoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableBorrowableInventoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBorrowableInventoryType(val *BorrowableInventoryType) *NullableBorrowableInventoryType {
	return &NullableBorrowableInventoryType{value: val, isSet: true}
}

func (v NullableBorrowableInventoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBorrowableInventoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


