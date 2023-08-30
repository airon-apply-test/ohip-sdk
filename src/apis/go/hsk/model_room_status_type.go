/*
OPERA Cloud Housekeeping Service API

APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hsk

import (
	"encoding/json"
)

// checks if the RoomStatusType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomStatusType{}

// RoomStatusType struct for RoomStatusType
type RoomStatusType struct {
	// List of status of the reservation to which this Room is assigned..
	ReservationStatusList []HousekeepingRoomReservationStatusType `json:"reservationStatusList,omitempty"`
	HousekeepingRoomStatus *HousekeepingRoomStatusType `json:"housekeepingRoomStatus,omitempty"`
	FrontOfficeStatus *FrontOfficeRoomStatusType `json:"frontOfficeStatus,omitempty"`
	HousekeepingStatus *FrontOfficeRoomStatusType `json:"housekeepingStatus,omitempty"`
}

// NewRoomStatusType instantiates a new RoomStatusType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomStatusType() *RoomStatusType {
	this := RoomStatusType{}
	return &this
}

// NewRoomStatusTypeWithDefaults instantiates a new RoomStatusType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomStatusTypeWithDefaults() *RoomStatusType {
	this := RoomStatusType{}
	return &this
}

// GetReservationStatusList returns the ReservationStatusList field value if set, zero value otherwise.
func (o *RoomStatusType) GetReservationStatusList() []HousekeepingRoomReservationStatusType {
	if o == nil || IsNil(o.ReservationStatusList) {
		var ret []HousekeepingRoomReservationStatusType
		return ret
	}
	return o.ReservationStatusList
}

// GetReservationStatusListOk returns a tuple with the ReservationStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomStatusType) GetReservationStatusListOk() ([]HousekeepingRoomReservationStatusType, bool) {
	if o == nil || IsNil(o.ReservationStatusList) {
		return nil, false
	}
	return o.ReservationStatusList, true
}

// HasReservationStatusList returns a boolean if a field has been set.
func (o *RoomStatusType) HasReservationStatusList() bool {
	if o != nil && !IsNil(o.ReservationStatusList) {
		return true
	}

	return false
}

// SetReservationStatusList gets a reference to the given []HousekeepingRoomReservationStatusType and assigns it to the ReservationStatusList field.
func (o *RoomStatusType) SetReservationStatusList(v []HousekeepingRoomReservationStatusType) {
	o.ReservationStatusList = v
}

// GetHousekeepingRoomStatus returns the HousekeepingRoomStatus field value if set, zero value otherwise.
func (o *RoomStatusType) GetHousekeepingRoomStatus() HousekeepingRoomStatusType {
	if o == nil || IsNil(o.HousekeepingRoomStatus) {
		var ret HousekeepingRoomStatusType
		return ret
	}
	return *o.HousekeepingRoomStatus
}

// GetHousekeepingRoomStatusOk returns a tuple with the HousekeepingRoomStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomStatusType) GetHousekeepingRoomStatusOk() (*HousekeepingRoomStatusType, bool) {
	if o == nil || IsNil(o.HousekeepingRoomStatus) {
		return nil, false
	}
	return o.HousekeepingRoomStatus, true
}

// HasHousekeepingRoomStatus returns a boolean if a field has been set.
func (o *RoomStatusType) HasHousekeepingRoomStatus() bool {
	if o != nil && !IsNil(o.HousekeepingRoomStatus) {
		return true
	}

	return false
}

// SetHousekeepingRoomStatus gets a reference to the given HousekeepingRoomStatusType and assigns it to the HousekeepingRoomStatus field.
func (o *RoomStatusType) SetHousekeepingRoomStatus(v HousekeepingRoomStatusType) {
	o.HousekeepingRoomStatus = &v
}

// GetFrontOfficeStatus returns the FrontOfficeStatus field value if set, zero value otherwise.
func (o *RoomStatusType) GetFrontOfficeStatus() FrontOfficeRoomStatusType {
	if o == nil || IsNil(o.FrontOfficeStatus) {
		var ret FrontOfficeRoomStatusType
		return ret
	}
	return *o.FrontOfficeStatus
}

// GetFrontOfficeStatusOk returns a tuple with the FrontOfficeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomStatusType) GetFrontOfficeStatusOk() (*FrontOfficeRoomStatusType, bool) {
	if o == nil || IsNil(o.FrontOfficeStatus) {
		return nil, false
	}
	return o.FrontOfficeStatus, true
}

// HasFrontOfficeStatus returns a boolean if a field has been set.
func (o *RoomStatusType) HasFrontOfficeStatus() bool {
	if o != nil && !IsNil(o.FrontOfficeStatus) {
		return true
	}

	return false
}

// SetFrontOfficeStatus gets a reference to the given FrontOfficeRoomStatusType and assigns it to the FrontOfficeStatus field.
func (o *RoomStatusType) SetFrontOfficeStatus(v FrontOfficeRoomStatusType) {
	o.FrontOfficeStatus = &v
}

// GetHousekeepingStatus returns the HousekeepingStatus field value if set, zero value otherwise.
func (o *RoomStatusType) GetHousekeepingStatus() FrontOfficeRoomStatusType {
	if o == nil || IsNil(o.HousekeepingStatus) {
		var ret FrontOfficeRoomStatusType
		return ret
	}
	return *o.HousekeepingStatus
}

// GetHousekeepingStatusOk returns a tuple with the HousekeepingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomStatusType) GetHousekeepingStatusOk() (*FrontOfficeRoomStatusType, bool) {
	if o == nil || IsNil(o.HousekeepingStatus) {
		return nil, false
	}
	return o.HousekeepingStatus, true
}

// HasHousekeepingStatus returns a boolean if a field has been set.
func (o *RoomStatusType) HasHousekeepingStatus() bool {
	if o != nil && !IsNil(o.HousekeepingStatus) {
		return true
	}

	return false
}

// SetHousekeepingStatus gets a reference to the given FrontOfficeRoomStatusType and assigns it to the HousekeepingStatus field.
func (o *RoomStatusType) SetHousekeepingStatus(v FrontOfficeRoomStatusType) {
	o.HousekeepingStatus = &v
}

func (o RoomStatusType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomStatusType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationStatusList) {
		toSerialize["reservationStatusList"] = o.ReservationStatusList
	}
	if !IsNil(o.HousekeepingRoomStatus) {
		toSerialize["housekeepingRoomStatus"] = o.HousekeepingRoomStatus
	}
	if !IsNil(o.FrontOfficeStatus) {
		toSerialize["frontOfficeStatus"] = o.FrontOfficeStatus
	}
	if !IsNil(o.HousekeepingStatus) {
		toSerialize["housekeepingStatus"] = o.HousekeepingStatus
	}
	return toSerialize, nil
}

type NullableRoomStatusType struct {
	value *RoomStatusType
	isSet bool
}

func (v NullableRoomStatusType) Get() *RoomStatusType {
	return v.value
}

func (v *NullableRoomStatusType) Set(val *RoomStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomStatusType(val *RoomStatusType) *NullableRoomStatusType {
	return &NullableRoomStatusType{value: val, isSet: true}
}

func (v NullableRoomStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


