/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the BorrowedInventoryToReturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BorrowedInventoryToReturn{}

// BorrowedInventoryToReturn The standard optional Opera Context element is also included.
type BorrowedInventoryToReturn struct {
	// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
	HotelId *string `json:"hotelId,omitempty"`
	BlockId *BlockId `json:"blockId,omitempty"`
	ExistingReservationId *UniqueIDType `json:"existingReservationId,omitempty"`
	// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
	RoomType *string `json:"roomType,omitempty"`
	// The number of adults that are expected to stay in a room.
	AdultCount *int32 `json:"adultCount,omitempty"`
	// The date and number of rooms to return to either a room type or House.
	InventoryToReturnList []InventoryToReturnType `json:"inventoryToReturnList,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBorrowedInventoryToReturn instantiates a new BorrowedInventoryToReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBorrowedInventoryToReturn() *BorrowedInventoryToReturn {
	this := BorrowedInventoryToReturn{}
	return &this
}

// NewBorrowedInventoryToReturnWithDefaults instantiates a new BorrowedInventoryToReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBorrowedInventoryToReturnWithDefaults() *BorrowedInventoryToReturn {
	this := BorrowedInventoryToReturn{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *BorrowedInventoryToReturn) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowedInventoryToReturn) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *BorrowedInventoryToReturn) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *BorrowedInventoryToReturn) SetHotelId(v string) {
	o.HotelId = &v
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *BorrowedInventoryToReturn) GetBlockId() BlockId {
	if o == nil || IsNil(o.BlockId) {
		var ret BlockId
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowedInventoryToReturn) GetBlockIdOk() (*BlockId, bool) {
	if o == nil || IsNil(o.BlockId) {
		return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *BorrowedInventoryToReturn) HasBlockId() bool {
	if o != nil && !IsNil(o.BlockId) {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given BlockId and assigns it to the BlockId field.
func (o *BorrowedInventoryToReturn) SetBlockId(v BlockId) {
	o.BlockId = &v
}

// GetExistingReservationId returns the ExistingReservationId field value if set, zero value otherwise.
func (o *BorrowedInventoryToReturn) GetExistingReservationId() UniqueIDType {
	if o == nil || IsNil(o.ExistingReservationId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ExistingReservationId
}

// GetExistingReservationIdOk returns a tuple with the ExistingReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowedInventoryToReturn) GetExistingReservationIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ExistingReservationId) {
		return nil, false
	}
	return o.ExistingReservationId, true
}

// HasExistingReservationId returns a boolean if a field has been set.
func (o *BorrowedInventoryToReturn) HasExistingReservationId() bool {
	if o != nil && !IsNil(o.ExistingReservationId) {
		return true
	}

	return false
}

// SetExistingReservationId gets a reference to the given UniqueIDType and assigns it to the ExistingReservationId field.
func (o *BorrowedInventoryToReturn) SetExistingReservationId(v UniqueIDType) {
	o.ExistingReservationId = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *BorrowedInventoryToReturn) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowedInventoryToReturn) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *BorrowedInventoryToReturn) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *BorrowedInventoryToReturn) SetRoomType(v string) {
	o.RoomType = &v
}

// GetAdultCount returns the AdultCount field value if set, zero value otherwise.
func (o *BorrowedInventoryToReturn) GetAdultCount() int32 {
	if o == nil || IsNil(o.AdultCount) {
		var ret int32
		return ret
	}
	return *o.AdultCount
}

// GetAdultCountOk returns a tuple with the AdultCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowedInventoryToReturn) GetAdultCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AdultCount) {
		return nil, false
	}
	return o.AdultCount, true
}

// HasAdultCount returns a boolean if a field has been set.
func (o *BorrowedInventoryToReturn) HasAdultCount() bool {
	if o != nil && !IsNil(o.AdultCount) {
		return true
	}

	return false
}

// SetAdultCount gets a reference to the given int32 and assigns it to the AdultCount field.
func (o *BorrowedInventoryToReturn) SetAdultCount(v int32) {
	o.AdultCount = &v
}

// GetInventoryToReturnList returns the InventoryToReturnList field value if set, zero value otherwise.
func (o *BorrowedInventoryToReturn) GetInventoryToReturnList() []InventoryToReturnType {
	if o == nil || IsNil(o.InventoryToReturnList) {
		var ret []InventoryToReturnType
		return ret
	}
	return o.InventoryToReturnList
}

// GetInventoryToReturnListOk returns a tuple with the InventoryToReturnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowedInventoryToReturn) GetInventoryToReturnListOk() ([]InventoryToReturnType, bool) {
	if o == nil || IsNil(o.InventoryToReturnList) {
		return nil, false
	}
	return o.InventoryToReturnList, true
}

// HasInventoryToReturnList returns a boolean if a field has been set.
func (o *BorrowedInventoryToReturn) HasInventoryToReturnList() bool {
	if o != nil && !IsNil(o.InventoryToReturnList) {
		return true
	}

	return false
}

// SetInventoryToReturnList gets a reference to the given []InventoryToReturnType and assigns it to the InventoryToReturnList field.
func (o *BorrowedInventoryToReturn) SetInventoryToReturnList(v []InventoryToReturnType) {
	o.InventoryToReturnList = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BorrowedInventoryToReturn) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowedInventoryToReturn) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BorrowedInventoryToReturn) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BorrowedInventoryToReturn) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BorrowedInventoryToReturn) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BorrowedInventoryToReturn) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BorrowedInventoryToReturn) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BorrowedInventoryToReturn) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o BorrowedInventoryToReturn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BorrowedInventoryToReturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.BlockId) {
		toSerialize["blockId"] = o.BlockId
	}
	if !IsNil(o.ExistingReservationId) {
		toSerialize["existingReservationId"] = o.ExistingReservationId
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.AdultCount) {
		toSerialize["adultCount"] = o.AdultCount
	}
	if !IsNil(o.InventoryToReturnList) {
		toSerialize["inventoryToReturnList"] = o.InventoryToReturnList
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBorrowedInventoryToReturn struct {
	value *BorrowedInventoryToReturn
	isSet bool
}

func (v NullableBorrowedInventoryToReturn) Get() *BorrowedInventoryToReturn {
	return v.value
}

func (v *NullableBorrowedInventoryToReturn) Set(val *BorrowedInventoryToReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableBorrowedInventoryToReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableBorrowedInventoryToReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBorrowedInventoryToReturn(val *BorrowedInventoryToReturn) *NullableBorrowedInventoryToReturn {
	return &NullableBorrowedInventoryToReturn{value: val, isSet: true}
}

func (v NullableBorrowedInventoryToReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBorrowedInventoryToReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


