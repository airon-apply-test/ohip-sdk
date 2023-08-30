/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the LinkedAllowanceReservationInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkedAllowanceReservationInfoType{}

// LinkedAllowanceReservationInfoType Linked Allowance Reservation Information.
type LinkedAllowanceReservationInfoType struct {
	ReservationId *ReservationId `json:"reservationId,omitempty"`
	// Reservation confirmation number.
	ConfirmationNo *string `json:"confirmationNo,omitempty"`
	GuestNameId *UniqueIDType `json:"guestNameId,omitempty"`
	// Display Name for the guest.
	GuestDisplayName *string `json:"guestDisplayName,omitempty"`
	// Guest Room number.
	RoomId *string `json:"roomId,omitempty"`
	// Indicates that guest is allowed to consumed shared allowances from others.
	ConsumeSharedAllowances *bool `json:"consumeSharedAllowances,omitempty"`
}

// NewLinkedAllowanceReservationInfoType instantiates a new LinkedAllowanceReservationInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedAllowanceReservationInfoType() *LinkedAllowanceReservationInfoType {
	this := LinkedAllowanceReservationInfoType{}
	return &this
}

// NewLinkedAllowanceReservationInfoTypeWithDefaults instantiates a new LinkedAllowanceReservationInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedAllowanceReservationInfoTypeWithDefaults() *LinkedAllowanceReservationInfoType {
	this := LinkedAllowanceReservationInfoType{}
	return &this
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *LinkedAllowanceReservationInfoType) GetReservationId() ReservationId {
	if o == nil || IsNil(o.ReservationId) {
		var ret ReservationId
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedAllowanceReservationInfoType) GetReservationIdOk() (*ReservationId, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *LinkedAllowanceReservationInfoType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given ReservationId and assigns it to the ReservationId field.
func (o *LinkedAllowanceReservationInfoType) SetReservationId(v ReservationId) {
	o.ReservationId = &v
}

// GetConfirmationNo returns the ConfirmationNo field value if set, zero value otherwise.
func (o *LinkedAllowanceReservationInfoType) GetConfirmationNo() string {
	if o == nil || IsNil(o.ConfirmationNo) {
		var ret string
		return ret
	}
	return *o.ConfirmationNo
}

// GetConfirmationNoOk returns a tuple with the ConfirmationNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedAllowanceReservationInfoType) GetConfirmationNoOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmationNo) {
		return nil, false
	}
	return o.ConfirmationNo, true
}

// HasConfirmationNo returns a boolean if a field has been set.
func (o *LinkedAllowanceReservationInfoType) HasConfirmationNo() bool {
	if o != nil && !IsNil(o.ConfirmationNo) {
		return true
	}

	return false
}

// SetConfirmationNo gets a reference to the given string and assigns it to the ConfirmationNo field.
func (o *LinkedAllowanceReservationInfoType) SetConfirmationNo(v string) {
	o.ConfirmationNo = &v
}

// GetGuestNameId returns the GuestNameId field value if set, zero value otherwise.
func (o *LinkedAllowanceReservationInfoType) GetGuestNameId() UniqueIDType {
	if o == nil || IsNil(o.GuestNameId) {
		var ret UniqueIDType
		return ret
	}
	return *o.GuestNameId
}

// GetGuestNameIdOk returns a tuple with the GuestNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedAllowanceReservationInfoType) GetGuestNameIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.GuestNameId) {
		return nil, false
	}
	return o.GuestNameId, true
}

// HasGuestNameId returns a boolean if a field has been set.
func (o *LinkedAllowanceReservationInfoType) HasGuestNameId() bool {
	if o != nil && !IsNil(o.GuestNameId) {
		return true
	}

	return false
}

// SetGuestNameId gets a reference to the given UniqueIDType and assigns it to the GuestNameId field.
func (o *LinkedAllowanceReservationInfoType) SetGuestNameId(v UniqueIDType) {
	o.GuestNameId = &v
}

// GetGuestDisplayName returns the GuestDisplayName field value if set, zero value otherwise.
func (o *LinkedAllowanceReservationInfoType) GetGuestDisplayName() string {
	if o == nil || IsNil(o.GuestDisplayName) {
		var ret string
		return ret
	}
	return *o.GuestDisplayName
}

// GetGuestDisplayNameOk returns a tuple with the GuestDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedAllowanceReservationInfoType) GetGuestDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestDisplayName) {
		return nil, false
	}
	return o.GuestDisplayName, true
}

// HasGuestDisplayName returns a boolean if a field has been set.
func (o *LinkedAllowanceReservationInfoType) HasGuestDisplayName() bool {
	if o != nil && !IsNil(o.GuestDisplayName) {
		return true
	}

	return false
}

// SetGuestDisplayName gets a reference to the given string and assigns it to the GuestDisplayName field.
func (o *LinkedAllowanceReservationInfoType) SetGuestDisplayName(v string) {
	o.GuestDisplayName = &v
}

// GetRoomId returns the RoomId field value if set, zero value otherwise.
func (o *LinkedAllowanceReservationInfoType) GetRoomId() string {
	if o == nil || IsNil(o.RoomId) {
		var ret string
		return ret
	}
	return *o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedAllowanceReservationInfoType) GetRoomIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoomId) {
		return nil, false
	}
	return o.RoomId, true
}

// HasRoomId returns a boolean if a field has been set.
func (o *LinkedAllowanceReservationInfoType) HasRoomId() bool {
	if o != nil && !IsNil(o.RoomId) {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given string and assigns it to the RoomId field.
func (o *LinkedAllowanceReservationInfoType) SetRoomId(v string) {
	o.RoomId = &v
}

// GetConsumeSharedAllowances returns the ConsumeSharedAllowances field value if set, zero value otherwise.
func (o *LinkedAllowanceReservationInfoType) GetConsumeSharedAllowances() bool {
	if o == nil || IsNil(o.ConsumeSharedAllowances) {
		var ret bool
		return ret
	}
	return *o.ConsumeSharedAllowances
}

// GetConsumeSharedAllowancesOk returns a tuple with the ConsumeSharedAllowances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedAllowanceReservationInfoType) GetConsumeSharedAllowancesOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsumeSharedAllowances) {
		return nil, false
	}
	return o.ConsumeSharedAllowances, true
}

// HasConsumeSharedAllowances returns a boolean if a field has been set.
func (o *LinkedAllowanceReservationInfoType) HasConsumeSharedAllowances() bool {
	if o != nil && !IsNil(o.ConsumeSharedAllowances) {
		return true
	}

	return false
}

// SetConsumeSharedAllowances gets a reference to the given bool and assigns it to the ConsumeSharedAllowances field.
func (o *LinkedAllowanceReservationInfoType) SetConsumeSharedAllowances(v bool) {
	o.ConsumeSharedAllowances = &v
}

func (o LinkedAllowanceReservationInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkedAllowanceReservationInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.ConfirmationNo) {
		toSerialize["confirmationNo"] = o.ConfirmationNo
	}
	if !IsNil(o.GuestNameId) {
		toSerialize["guestNameId"] = o.GuestNameId
	}
	if !IsNil(o.GuestDisplayName) {
		toSerialize["guestDisplayName"] = o.GuestDisplayName
	}
	if !IsNil(o.RoomId) {
		toSerialize["roomId"] = o.RoomId
	}
	if !IsNil(o.ConsumeSharedAllowances) {
		toSerialize["consumeSharedAllowances"] = o.ConsumeSharedAllowances
	}
	return toSerialize, nil
}

type NullableLinkedAllowanceReservationInfoType struct {
	value *LinkedAllowanceReservationInfoType
	isSet bool
}

func (v NullableLinkedAllowanceReservationInfoType) Get() *LinkedAllowanceReservationInfoType {
	return v.value
}

func (v *NullableLinkedAllowanceReservationInfoType) Set(val *LinkedAllowanceReservationInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedAllowanceReservationInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedAllowanceReservationInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedAllowanceReservationInfoType(val *LinkedAllowanceReservationInfoType) *NullableLinkedAllowanceReservationInfoType {
	return &NullableLinkedAllowanceReservationInfoType{value: val, isSet: true}
}

func (v NullableLinkedAllowanceReservationInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedAllowanceReservationInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


