/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResvRoutingInfoTypeRoom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResvRoutingInfoTypeRoom{}

// ResvRoutingInfoTypeRoom Room routing type.
type ResvRoutingInfoTypeRoom struct {
	// Room number to route the instructions.
	RoomId *string `json:"roomId,omitempty"`
	GuestNameId *UniqueIDType `json:"guestNameId,omitempty"`
	// Display Name for the guest.
	GuestDisplayName *string `json:"guestDisplayName,omitempty"`
	ReservationNameId *UniqueIDType `json:"reservationNameId,omitempty"`
	// Set of routing instructions associated to this routing type.
	Instructions []RoutingInstructionType `json:"instructions,omitempty"`
}

// NewResvRoutingInfoTypeRoom instantiates a new ResvRoutingInfoTypeRoom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResvRoutingInfoTypeRoom() *ResvRoutingInfoTypeRoom {
	this := ResvRoutingInfoTypeRoom{}
	return &this
}

// NewResvRoutingInfoTypeRoomWithDefaults instantiates a new ResvRoutingInfoTypeRoom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResvRoutingInfoTypeRoomWithDefaults() *ResvRoutingInfoTypeRoom {
	this := ResvRoutingInfoTypeRoom{}
	return &this
}

// GetRoomId returns the RoomId field value if set, zero value otherwise.
func (o *ResvRoutingInfoTypeRoom) GetRoomId() string {
	if o == nil || IsNil(o.RoomId) {
		var ret string
		return ret
	}
	return *o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResvRoutingInfoTypeRoom) GetRoomIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoomId) {
		return nil, false
	}
	return o.RoomId, true
}

// HasRoomId returns a boolean if a field has been set.
func (o *ResvRoutingInfoTypeRoom) HasRoomId() bool {
	if o != nil && !IsNil(o.RoomId) {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given string and assigns it to the RoomId field.
func (o *ResvRoutingInfoTypeRoom) SetRoomId(v string) {
	o.RoomId = &v
}

// GetGuestNameId returns the GuestNameId field value if set, zero value otherwise.
func (o *ResvRoutingInfoTypeRoom) GetGuestNameId() UniqueIDType {
	if o == nil || IsNil(o.GuestNameId) {
		var ret UniqueIDType
		return ret
	}
	return *o.GuestNameId
}

// GetGuestNameIdOk returns a tuple with the GuestNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResvRoutingInfoTypeRoom) GetGuestNameIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.GuestNameId) {
		return nil, false
	}
	return o.GuestNameId, true
}

// HasGuestNameId returns a boolean if a field has been set.
func (o *ResvRoutingInfoTypeRoom) HasGuestNameId() bool {
	if o != nil && !IsNil(o.GuestNameId) {
		return true
	}

	return false
}

// SetGuestNameId gets a reference to the given UniqueIDType and assigns it to the GuestNameId field.
func (o *ResvRoutingInfoTypeRoom) SetGuestNameId(v UniqueIDType) {
	o.GuestNameId = &v
}

// GetGuestDisplayName returns the GuestDisplayName field value if set, zero value otherwise.
func (o *ResvRoutingInfoTypeRoom) GetGuestDisplayName() string {
	if o == nil || IsNil(o.GuestDisplayName) {
		var ret string
		return ret
	}
	return *o.GuestDisplayName
}

// GetGuestDisplayNameOk returns a tuple with the GuestDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResvRoutingInfoTypeRoom) GetGuestDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestDisplayName) {
		return nil, false
	}
	return o.GuestDisplayName, true
}

// HasGuestDisplayName returns a boolean if a field has been set.
func (o *ResvRoutingInfoTypeRoom) HasGuestDisplayName() bool {
	if o != nil && !IsNil(o.GuestDisplayName) {
		return true
	}

	return false
}

// SetGuestDisplayName gets a reference to the given string and assigns it to the GuestDisplayName field.
func (o *ResvRoutingInfoTypeRoom) SetGuestDisplayName(v string) {
	o.GuestDisplayName = &v
}

// GetReservationNameId returns the ReservationNameId field value if set, zero value otherwise.
func (o *ResvRoutingInfoTypeRoom) GetReservationNameId() UniqueIDType {
	if o == nil || IsNil(o.ReservationNameId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ReservationNameId
}

// GetReservationNameIdOk returns a tuple with the ReservationNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResvRoutingInfoTypeRoom) GetReservationNameIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationNameId) {
		return nil, false
	}
	return o.ReservationNameId, true
}

// HasReservationNameId returns a boolean if a field has been set.
func (o *ResvRoutingInfoTypeRoom) HasReservationNameId() bool {
	if o != nil && !IsNil(o.ReservationNameId) {
		return true
	}

	return false
}

// SetReservationNameId gets a reference to the given UniqueIDType and assigns it to the ReservationNameId field.
func (o *ResvRoutingInfoTypeRoom) SetReservationNameId(v UniqueIDType) {
	o.ReservationNameId = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *ResvRoutingInfoTypeRoom) GetInstructions() []RoutingInstructionType {
	if o == nil || IsNil(o.Instructions) {
		var ret []RoutingInstructionType
		return ret
	}
	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResvRoutingInfoTypeRoom) GetInstructionsOk() ([]RoutingInstructionType, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *ResvRoutingInfoTypeRoom) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []RoutingInstructionType and assigns it to the Instructions field.
func (o *ResvRoutingInfoTypeRoom) SetInstructions(v []RoutingInstructionType) {
	o.Instructions = v
}

func (o ResvRoutingInfoTypeRoom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResvRoutingInfoTypeRoom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomId) {
		toSerialize["roomId"] = o.RoomId
	}
	if !IsNil(o.GuestNameId) {
		toSerialize["guestNameId"] = o.GuestNameId
	}
	if !IsNil(o.GuestDisplayName) {
		toSerialize["guestDisplayName"] = o.GuestDisplayName
	}
	if !IsNil(o.ReservationNameId) {
		toSerialize["reservationNameId"] = o.ReservationNameId
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	return toSerialize, nil
}

type NullableResvRoutingInfoTypeRoom struct {
	value *ResvRoutingInfoTypeRoom
	isSet bool
}

func (v NullableResvRoutingInfoTypeRoom) Get() *ResvRoutingInfoTypeRoom {
	return v.value
}

func (v *NullableResvRoutingInfoTypeRoom) Set(val *ResvRoutingInfoTypeRoom) {
	v.value = val
	v.isSet = true
}

func (v NullableResvRoutingInfoTypeRoom) IsSet() bool {
	return v.isSet
}

func (v *NullableResvRoutingInfoTypeRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResvRoutingInfoTypeRoom(val *ResvRoutingInfoTypeRoom) *NullableResvRoutingInfoTypeRoom {
	return &NullableResvRoutingInfoTypeRoom{value: val, isSet: true}
}

func (v NullableResvRoutingInfoTypeRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResvRoutingInfoTypeRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


