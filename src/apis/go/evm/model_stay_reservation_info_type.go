/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
)

// checks if the StayReservationInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StayReservationInfoType{}

// StayReservationInfoType The Reservation class contains the current reservation being created or altered.
type StayReservationInfoType struct {
	// Unique Id that references an object uniquely in the system.
	ReservationIdList []UniqueIDType `json:"reservationIdList,omitempty"`
	RoomStay *StayInfoType `json:"roomStay,omitempty"`
	AttachedProfiles []ResAttachedProfileType `json:"attachedProfiles,omitempty"`
	HotelId *string `json:"hotelId,omitempty"`
	ReservationStatus *PMSResStatusType `json:"reservationStatus,omitempty"`
	ComputedReservationStatus *PMSResStatusType `json:"computedReservationStatus,omitempty"`
}

// NewStayReservationInfoType instantiates a new StayReservationInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStayReservationInfoType() *StayReservationInfoType {
	this := StayReservationInfoType{}
	return &this
}

// NewStayReservationInfoTypeWithDefaults instantiates a new StayReservationInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStayReservationInfoTypeWithDefaults() *StayReservationInfoType {
	this := StayReservationInfoType{}
	return &this
}

// GetReservationIdList returns the ReservationIdList field value if set, zero value otherwise.
func (o *StayReservationInfoType) GetReservationIdList() []UniqueIDType {
	if o == nil || IsNil(o.ReservationIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ReservationIdList
}

// GetReservationIdListOk returns a tuple with the ReservationIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayReservationInfoType) GetReservationIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationIdList) {
		return nil, false
	}
	return o.ReservationIdList, true
}

// HasReservationIdList returns a boolean if a field has been set.
func (o *StayReservationInfoType) HasReservationIdList() bool {
	if o != nil && !IsNil(o.ReservationIdList) {
		return true
	}

	return false
}

// SetReservationIdList gets a reference to the given []UniqueIDType and assigns it to the ReservationIdList field.
func (o *StayReservationInfoType) SetReservationIdList(v []UniqueIDType) {
	o.ReservationIdList = v
}

// GetRoomStay returns the RoomStay field value if set, zero value otherwise.
func (o *StayReservationInfoType) GetRoomStay() StayInfoType {
	if o == nil || IsNil(o.RoomStay) {
		var ret StayInfoType
		return ret
	}
	return *o.RoomStay
}

// GetRoomStayOk returns a tuple with the RoomStay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayReservationInfoType) GetRoomStayOk() (*StayInfoType, bool) {
	if o == nil || IsNil(o.RoomStay) {
		return nil, false
	}
	return o.RoomStay, true
}

// HasRoomStay returns a boolean if a field has been set.
func (o *StayReservationInfoType) HasRoomStay() bool {
	if o != nil && !IsNil(o.RoomStay) {
		return true
	}

	return false
}

// SetRoomStay gets a reference to the given StayInfoType and assigns it to the RoomStay field.
func (o *StayReservationInfoType) SetRoomStay(v StayInfoType) {
	o.RoomStay = &v
}

// GetAttachedProfiles returns the AttachedProfiles field value if set, zero value otherwise.
func (o *StayReservationInfoType) GetAttachedProfiles() []ResAttachedProfileType {
	if o == nil || IsNil(o.AttachedProfiles) {
		var ret []ResAttachedProfileType
		return ret
	}
	return o.AttachedProfiles
}

// GetAttachedProfilesOk returns a tuple with the AttachedProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayReservationInfoType) GetAttachedProfilesOk() ([]ResAttachedProfileType, bool) {
	if o == nil || IsNil(o.AttachedProfiles) {
		return nil, false
	}
	return o.AttachedProfiles, true
}

// HasAttachedProfiles returns a boolean if a field has been set.
func (o *StayReservationInfoType) HasAttachedProfiles() bool {
	if o != nil && !IsNil(o.AttachedProfiles) {
		return true
	}

	return false
}

// SetAttachedProfiles gets a reference to the given []ResAttachedProfileType and assigns it to the AttachedProfiles field.
func (o *StayReservationInfoType) SetAttachedProfiles(v []ResAttachedProfileType) {
	o.AttachedProfiles = v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *StayReservationInfoType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayReservationInfoType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *StayReservationInfoType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *StayReservationInfoType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetReservationStatus returns the ReservationStatus field value if set, zero value otherwise.
func (o *StayReservationInfoType) GetReservationStatus() PMSResStatusType {
	if o == nil || IsNil(o.ReservationStatus) {
		var ret PMSResStatusType
		return ret
	}
	return *o.ReservationStatus
}

// GetReservationStatusOk returns a tuple with the ReservationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayReservationInfoType) GetReservationStatusOk() (*PMSResStatusType, bool) {
	if o == nil || IsNil(o.ReservationStatus) {
		return nil, false
	}
	return o.ReservationStatus, true
}

// HasReservationStatus returns a boolean if a field has been set.
func (o *StayReservationInfoType) HasReservationStatus() bool {
	if o != nil && !IsNil(o.ReservationStatus) {
		return true
	}

	return false
}

// SetReservationStatus gets a reference to the given PMSResStatusType and assigns it to the ReservationStatus field.
func (o *StayReservationInfoType) SetReservationStatus(v PMSResStatusType) {
	o.ReservationStatus = &v
}

// GetComputedReservationStatus returns the ComputedReservationStatus field value if set, zero value otherwise.
func (o *StayReservationInfoType) GetComputedReservationStatus() PMSResStatusType {
	if o == nil || IsNil(o.ComputedReservationStatus) {
		var ret PMSResStatusType
		return ret
	}
	return *o.ComputedReservationStatus
}

// GetComputedReservationStatusOk returns a tuple with the ComputedReservationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StayReservationInfoType) GetComputedReservationStatusOk() (*PMSResStatusType, bool) {
	if o == nil || IsNil(o.ComputedReservationStatus) {
		return nil, false
	}
	return o.ComputedReservationStatus, true
}

// HasComputedReservationStatus returns a boolean if a field has been set.
func (o *StayReservationInfoType) HasComputedReservationStatus() bool {
	if o != nil && !IsNil(o.ComputedReservationStatus) {
		return true
	}

	return false
}

// SetComputedReservationStatus gets a reference to the given PMSResStatusType and assigns it to the ComputedReservationStatus field.
func (o *StayReservationInfoType) SetComputedReservationStatus(v PMSResStatusType) {
	o.ComputedReservationStatus = &v
}

func (o StayReservationInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StayReservationInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationIdList) {
		toSerialize["reservationIdList"] = o.ReservationIdList
	}
	if !IsNil(o.RoomStay) {
		toSerialize["roomStay"] = o.RoomStay
	}
	if !IsNil(o.AttachedProfiles) {
		toSerialize["attachedProfiles"] = o.AttachedProfiles
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ReservationStatus) {
		toSerialize["reservationStatus"] = o.ReservationStatus
	}
	if !IsNil(o.ComputedReservationStatus) {
		toSerialize["computedReservationStatus"] = o.ComputedReservationStatus
	}
	return toSerialize, nil
}

type NullableStayReservationInfoType struct {
	value *StayReservationInfoType
	isSet bool
}

func (v NullableStayReservationInfoType) Get() *StayReservationInfoType {
	return v.value
}

func (v *NullableStayReservationInfoType) Set(val *StayReservationInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableStayReservationInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableStayReservationInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStayReservationInfoType(val *StayReservationInfoType) *NullableStayReservationInfoType {
	return &NullableStayReservationInfoType{value: val, isSet: true}
}

func (v NullableStayReservationInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStayReservationInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


