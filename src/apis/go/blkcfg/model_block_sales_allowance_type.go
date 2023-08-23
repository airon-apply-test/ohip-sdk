/*
OPERA Cloud Block Configuration API

APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BlockSalesAllowanceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockSalesAllowanceType{}

// BlockSalesAllowanceType Contains single entry information for Block's Sales Allowance.
type BlockSalesAllowanceType struct {
	// Contains Hotel Code.
	HotelId *string `json:"hotelId,omitempty"`
	// Contains sales allowance date.
	SalesAllowanceDate *string `json:"salesAllowanceDate,omitempty"`
	// Contains room type.
	RoomType *string `json:"roomType,omitempty"`
	// Contains room pool code.
	RoomPool *string `json:"roomPool,omitempty"`
	// Contains allocated number of rooms for allowance.
	Allowance *int32 `json:"allowance,omitempty"`
	// Contains number of rooms are booked/consumed from sales allowance.
	Booked *int32 `json:"booked,omitempty"`
	// Contains number of rooms are overbooked from sales allowance.
	OverBooked *int32 `json:"overBooked,omitempty"`
	// Contains cutoff date for sales allowance.
	CutoffDate *string `json:"cutoffDate,omitempty"`
}

// NewBlockSalesAllowanceType instantiates a new BlockSalesAllowanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSalesAllowanceType() *BlockSalesAllowanceType {
	this := BlockSalesAllowanceType{}
	return &this
}

// NewBlockSalesAllowanceTypeWithDefaults instantiates a new BlockSalesAllowanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSalesAllowanceTypeWithDefaults() *BlockSalesAllowanceType {
	this := BlockSalesAllowanceType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *BlockSalesAllowanceType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSalesAllowanceType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *BlockSalesAllowanceType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *BlockSalesAllowanceType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetSalesAllowanceDate returns the SalesAllowanceDate field value if set, zero value otherwise.
func (o *BlockSalesAllowanceType) GetSalesAllowanceDate() string {
	if o == nil || IsNil(o.SalesAllowanceDate) {
		var ret string
		return ret
	}
	return *o.SalesAllowanceDate
}

// GetSalesAllowanceDateOk returns a tuple with the SalesAllowanceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSalesAllowanceType) GetSalesAllowanceDateOk() (*string, bool) {
	if o == nil || IsNil(o.SalesAllowanceDate) {
		return nil, false
	}
	return o.SalesAllowanceDate, true
}

// HasSalesAllowanceDate returns a boolean if a field has been set.
func (o *BlockSalesAllowanceType) HasSalesAllowanceDate() bool {
	if o != nil && !IsNil(o.SalesAllowanceDate) {
		return true
	}

	return false
}

// SetSalesAllowanceDate gets a reference to the given string and assigns it to the SalesAllowanceDate field.
func (o *BlockSalesAllowanceType) SetSalesAllowanceDate(v string) {
	o.SalesAllowanceDate = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *BlockSalesAllowanceType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSalesAllowanceType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *BlockSalesAllowanceType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *BlockSalesAllowanceType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetRoomPool returns the RoomPool field value if set, zero value otherwise.
func (o *BlockSalesAllowanceType) GetRoomPool() string {
	if o == nil || IsNil(o.RoomPool) {
		var ret string
		return ret
	}
	return *o.RoomPool
}

// GetRoomPoolOk returns a tuple with the RoomPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSalesAllowanceType) GetRoomPoolOk() (*string, bool) {
	if o == nil || IsNil(o.RoomPool) {
		return nil, false
	}
	return o.RoomPool, true
}

// HasRoomPool returns a boolean if a field has been set.
func (o *BlockSalesAllowanceType) HasRoomPool() bool {
	if o != nil && !IsNil(o.RoomPool) {
		return true
	}

	return false
}

// SetRoomPool gets a reference to the given string and assigns it to the RoomPool field.
func (o *BlockSalesAllowanceType) SetRoomPool(v string) {
	o.RoomPool = &v
}

// GetAllowance returns the Allowance field value if set, zero value otherwise.
func (o *BlockSalesAllowanceType) GetAllowance() int32 {
	if o == nil || IsNil(o.Allowance) {
		var ret int32
		return ret
	}
	return *o.Allowance
}

// GetAllowanceOk returns a tuple with the Allowance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSalesAllowanceType) GetAllowanceOk() (*int32, bool) {
	if o == nil || IsNil(o.Allowance) {
		return nil, false
	}
	return o.Allowance, true
}

// HasAllowance returns a boolean if a field has been set.
func (o *BlockSalesAllowanceType) HasAllowance() bool {
	if o != nil && !IsNil(o.Allowance) {
		return true
	}

	return false
}

// SetAllowance gets a reference to the given int32 and assigns it to the Allowance field.
func (o *BlockSalesAllowanceType) SetAllowance(v int32) {
	o.Allowance = &v
}

// GetBooked returns the Booked field value if set, zero value otherwise.
func (o *BlockSalesAllowanceType) GetBooked() int32 {
	if o == nil || IsNil(o.Booked) {
		var ret int32
		return ret
	}
	return *o.Booked
}

// GetBookedOk returns a tuple with the Booked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSalesAllowanceType) GetBookedOk() (*int32, bool) {
	if o == nil || IsNil(o.Booked) {
		return nil, false
	}
	return o.Booked, true
}

// HasBooked returns a boolean if a field has been set.
func (o *BlockSalesAllowanceType) HasBooked() bool {
	if o != nil && !IsNil(o.Booked) {
		return true
	}

	return false
}

// SetBooked gets a reference to the given int32 and assigns it to the Booked field.
func (o *BlockSalesAllowanceType) SetBooked(v int32) {
	o.Booked = &v
}

// GetOverBooked returns the OverBooked field value if set, zero value otherwise.
func (o *BlockSalesAllowanceType) GetOverBooked() int32 {
	if o == nil || IsNil(o.OverBooked) {
		var ret int32
		return ret
	}
	return *o.OverBooked
}

// GetOverBookedOk returns a tuple with the OverBooked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSalesAllowanceType) GetOverBookedOk() (*int32, bool) {
	if o == nil || IsNil(o.OverBooked) {
		return nil, false
	}
	return o.OverBooked, true
}

// HasOverBooked returns a boolean if a field has been set.
func (o *BlockSalesAllowanceType) HasOverBooked() bool {
	if o != nil && !IsNil(o.OverBooked) {
		return true
	}

	return false
}

// SetOverBooked gets a reference to the given int32 and assigns it to the OverBooked field.
func (o *BlockSalesAllowanceType) SetOverBooked(v int32) {
	o.OverBooked = &v
}

// GetCutoffDate returns the CutoffDate field value if set, zero value otherwise.
func (o *BlockSalesAllowanceType) GetCutoffDate() string {
	if o == nil || IsNil(o.CutoffDate) {
		var ret string
		return ret
	}
	return *o.CutoffDate
}

// GetCutoffDateOk returns a tuple with the CutoffDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSalesAllowanceType) GetCutoffDateOk() (*string, bool) {
	if o == nil || IsNil(o.CutoffDate) {
		return nil, false
	}
	return o.CutoffDate, true
}

// HasCutoffDate returns a boolean if a field has been set.
func (o *BlockSalesAllowanceType) HasCutoffDate() bool {
	if o != nil && !IsNil(o.CutoffDate) {
		return true
	}

	return false
}

// SetCutoffDate gets a reference to the given string and assigns it to the CutoffDate field.
func (o *BlockSalesAllowanceType) SetCutoffDate(v string) {
	o.CutoffDate = &v
}

func (o BlockSalesAllowanceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockSalesAllowanceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.SalesAllowanceDate) {
		toSerialize["salesAllowanceDate"] = o.SalesAllowanceDate
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.RoomPool) {
		toSerialize["roomPool"] = o.RoomPool
	}
	if !IsNil(o.Allowance) {
		toSerialize["allowance"] = o.Allowance
	}
	if !IsNil(o.Booked) {
		toSerialize["booked"] = o.Booked
	}
	if !IsNil(o.OverBooked) {
		toSerialize["overBooked"] = o.OverBooked
	}
	if !IsNil(o.CutoffDate) {
		toSerialize["cutoffDate"] = o.CutoffDate
	}
	return toSerialize, nil
}

type NullableBlockSalesAllowanceType struct {
	value *BlockSalesAllowanceType
	isSet bool
}

func (v NullableBlockSalesAllowanceType) Get() *BlockSalesAllowanceType {
	return v.value
}

func (v *NullableBlockSalesAllowanceType) Set(val *BlockSalesAllowanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSalesAllowanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSalesAllowanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSalesAllowanceType(val *BlockSalesAllowanceType) *NullableBlockSalesAllowanceType {
	return &NullableBlockSalesAllowanceType{value: val, isSet: true}
}

func (v NullableBlockSalesAllowanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSalesAllowanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


