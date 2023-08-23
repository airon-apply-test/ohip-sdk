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

// checks if the TraceResvInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceResvInfoType{}

// TraceResvInfoType Reservation information related to the trace.
type TraceResvInfoType struct {
	// Unique Id that references an object uniquely in the system.
	ReservationIdList []UniqueIDType `json:"reservationIdList,omitempty"`
	// Hotel Code where trace is set.
	HotelId *string `json:"hotelId,omitempty"`
	TimeSpan *TimeSpanType `json:"timeSpan,omitempty"`
	// Room Id
	RoomId *string `json:"roomId,omitempty"`
	// Current Room Status.
	RoomStatus *string `json:"roomStatus,omitempty"`
	// Current Reservation Status.
	ReservationStatus *string `json:"reservationStatus,omitempty"`
	// Collection of guests associated with the reservation.
	ReservationGuests []ResGuestType `json:"reservationGuests,omitempty"`
}

// NewTraceResvInfoType instantiates a new TraceResvInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceResvInfoType() *TraceResvInfoType {
	this := TraceResvInfoType{}
	return &this
}

// NewTraceResvInfoTypeWithDefaults instantiates a new TraceResvInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceResvInfoTypeWithDefaults() *TraceResvInfoType {
	this := TraceResvInfoType{}
	return &this
}

// GetReservationIdList returns the ReservationIdList field value if set, zero value otherwise.
func (o *TraceResvInfoType) GetReservationIdList() []UniqueIDType {
	if o == nil || IsNil(o.ReservationIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ReservationIdList
}

// GetReservationIdListOk returns a tuple with the ReservationIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResvInfoType) GetReservationIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationIdList) {
		return nil, false
	}
	return o.ReservationIdList, true
}

// HasReservationIdList returns a boolean if a field has been set.
func (o *TraceResvInfoType) HasReservationIdList() bool {
	if o != nil && !IsNil(o.ReservationIdList) {
		return true
	}

	return false
}

// SetReservationIdList gets a reference to the given []UniqueIDType and assigns it to the ReservationIdList field.
func (o *TraceResvInfoType) SetReservationIdList(v []UniqueIDType) {
	o.ReservationIdList = v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *TraceResvInfoType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResvInfoType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *TraceResvInfoType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *TraceResvInfoType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *TraceResvInfoType) GetTimeSpan() TimeSpanType {
	if o == nil || IsNil(o.TimeSpan) {
		var ret TimeSpanType
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResvInfoType) GetTimeSpanOk() (*TimeSpanType, bool) {
	if o == nil || IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *TraceResvInfoType) HasTimeSpan() bool {
	if o != nil && !IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given TimeSpanType and assigns it to the TimeSpan field.
func (o *TraceResvInfoType) SetTimeSpan(v TimeSpanType) {
	o.TimeSpan = &v
}

// GetRoomId returns the RoomId field value if set, zero value otherwise.
func (o *TraceResvInfoType) GetRoomId() string {
	if o == nil || IsNil(o.RoomId) {
		var ret string
		return ret
	}
	return *o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResvInfoType) GetRoomIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoomId) {
		return nil, false
	}
	return o.RoomId, true
}

// HasRoomId returns a boolean if a field has been set.
func (o *TraceResvInfoType) HasRoomId() bool {
	if o != nil && !IsNil(o.RoomId) {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given string and assigns it to the RoomId field.
func (o *TraceResvInfoType) SetRoomId(v string) {
	o.RoomId = &v
}

// GetRoomStatus returns the RoomStatus field value if set, zero value otherwise.
func (o *TraceResvInfoType) GetRoomStatus() string {
	if o == nil || IsNil(o.RoomStatus) {
		var ret string
		return ret
	}
	return *o.RoomStatus
}

// GetRoomStatusOk returns a tuple with the RoomStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResvInfoType) GetRoomStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RoomStatus) {
		return nil, false
	}
	return o.RoomStatus, true
}

// HasRoomStatus returns a boolean if a field has been set.
func (o *TraceResvInfoType) HasRoomStatus() bool {
	if o != nil && !IsNil(o.RoomStatus) {
		return true
	}

	return false
}

// SetRoomStatus gets a reference to the given string and assigns it to the RoomStatus field.
func (o *TraceResvInfoType) SetRoomStatus(v string) {
	o.RoomStatus = &v
}

// GetReservationStatus returns the ReservationStatus field value if set, zero value otherwise.
func (o *TraceResvInfoType) GetReservationStatus() string {
	if o == nil || IsNil(o.ReservationStatus) {
		var ret string
		return ret
	}
	return *o.ReservationStatus
}

// GetReservationStatusOk returns a tuple with the ReservationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResvInfoType) GetReservationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReservationStatus) {
		return nil, false
	}
	return o.ReservationStatus, true
}

// HasReservationStatus returns a boolean if a field has been set.
func (o *TraceResvInfoType) HasReservationStatus() bool {
	if o != nil && !IsNil(o.ReservationStatus) {
		return true
	}

	return false
}

// SetReservationStatus gets a reference to the given string and assigns it to the ReservationStatus field.
func (o *TraceResvInfoType) SetReservationStatus(v string) {
	o.ReservationStatus = &v
}

// GetReservationGuests returns the ReservationGuests field value if set, zero value otherwise.
func (o *TraceResvInfoType) GetReservationGuests() []ResGuestType {
	if o == nil || IsNil(o.ReservationGuests) {
		var ret []ResGuestType
		return ret
	}
	return o.ReservationGuests
}

// GetReservationGuestsOk returns a tuple with the ReservationGuests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResvInfoType) GetReservationGuestsOk() ([]ResGuestType, bool) {
	if o == nil || IsNil(o.ReservationGuests) {
		return nil, false
	}
	return o.ReservationGuests, true
}

// HasReservationGuests returns a boolean if a field has been set.
func (o *TraceResvInfoType) HasReservationGuests() bool {
	if o != nil && !IsNil(o.ReservationGuests) {
		return true
	}

	return false
}

// SetReservationGuests gets a reference to the given []ResGuestType and assigns it to the ReservationGuests field.
func (o *TraceResvInfoType) SetReservationGuests(v []ResGuestType) {
	o.ReservationGuests = v
}

func (o TraceResvInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceResvInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationIdList) {
		toSerialize["reservationIdList"] = o.ReservationIdList
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.TimeSpan) {
		toSerialize["timeSpan"] = o.TimeSpan
	}
	if !IsNil(o.RoomId) {
		toSerialize["roomId"] = o.RoomId
	}
	if !IsNil(o.RoomStatus) {
		toSerialize["roomStatus"] = o.RoomStatus
	}
	if !IsNil(o.ReservationStatus) {
		toSerialize["reservationStatus"] = o.ReservationStatus
	}
	if !IsNil(o.ReservationGuests) {
		toSerialize["reservationGuests"] = o.ReservationGuests
	}
	return toSerialize, nil
}

type NullableTraceResvInfoType struct {
	value *TraceResvInfoType
	isSet bool
}

func (v NullableTraceResvInfoType) Get() *TraceResvInfoType {
	return v.value
}

func (v *NullableTraceResvInfoType) Set(val *TraceResvInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceResvInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceResvInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceResvInfoType(val *TraceResvInfoType) *NullableTraceResvInfoType {
	return &NullableTraceResvInfoType{value: val, isSet: true}
}

func (v NullableTraceResvInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceResvInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


