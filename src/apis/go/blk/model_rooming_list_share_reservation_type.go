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

// checks if the RoomingListShareReservationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomingListShareReservationType{}

// RoomingListShareReservationType Contains the rooming list reservation that is to be shared and the information about the share type of this reservation.
type RoomingListShareReservationType struct {
	ReservationId *UniqueIDType `json:"reservationId,omitempty"`
	TypeOfRateAmount *ShareDistributionInstructionType `json:"typeOfRateAmount,omitempty"`
	// Indicates if this reservation is a primary sharer.
	PrimaryShare *bool `json:"primaryShare,omitempty"`
	TimeSpan *TimeSpanType `json:"timeSpan,omitempty"`
}

// NewRoomingListShareReservationType instantiates a new RoomingListShareReservationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomingListShareReservationType() *RoomingListShareReservationType {
	this := RoomingListShareReservationType{}
	return &this
}

// NewRoomingListShareReservationTypeWithDefaults instantiates a new RoomingListShareReservationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomingListShareReservationTypeWithDefaults() *RoomingListShareReservationType {
	this := RoomingListShareReservationType{}
	return &this
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *RoomingListShareReservationType) GetReservationId() UniqueIDType {
	if o == nil || IsNil(o.ReservationId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomingListShareReservationType) GetReservationIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *RoomingListShareReservationType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given UniqueIDType and assigns it to the ReservationId field.
func (o *RoomingListShareReservationType) SetReservationId(v UniqueIDType) {
	o.ReservationId = &v
}

// GetTypeOfRateAmount returns the TypeOfRateAmount field value if set, zero value otherwise.
func (o *RoomingListShareReservationType) GetTypeOfRateAmount() ShareDistributionInstructionType {
	if o == nil || IsNil(o.TypeOfRateAmount) {
		var ret ShareDistributionInstructionType
		return ret
	}
	return *o.TypeOfRateAmount
}

// GetTypeOfRateAmountOk returns a tuple with the TypeOfRateAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomingListShareReservationType) GetTypeOfRateAmountOk() (*ShareDistributionInstructionType, bool) {
	if o == nil || IsNil(o.TypeOfRateAmount) {
		return nil, false
	}
	return o.TypeOfRateAmount, true
}

// HasTypeOfRateAmount returns a boolean if a field has been set.
func (o *RoomingListShareReservationType) HasTypeOfRateAmount() bool {
	if o != nil && !IsNil(o.TypeOfRateAmount) {
		return true
	}

	return false
}

// SetTypeOfRateAmount gets a reference to the given ShareDistributionInstructionType and assigns it to the TypeOfRateAmount field.
func (o *RoomingListShareReservationType) SetTypeOfRateAmount(v ShareDistributionInstructionType) {
	o.TypeOfRateAmount = &v
}

// GetPrimaryShare returns the PrimaryShare field value if set, zero value otherwise.
func (o *RoomingListShareReservationType) GetPrimaryShare() bool {
	if o == nil || IsNil(o.PrimaryShare) {
		var ret bool
		return ret
	}
	return *o.PrimaryShare
}

// GetPrimaryShareOk returns a tuple with the PrimaryShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomingListShareReservationType) GetPrimaryShareOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryShare) {
		return nil, false
	}
	return o.PrimaryShare, true
}

// HasPrimaryShare returns a boolean if a field has been set.
func (o *RoomingListShareReservationType) HasPrimaryShare() bool {
	if o != nil && !IsNil(o.PrimaryShare) {
		return true
	}

	return false
}

// SetPrimaryShare gets a reference to the given bool and assigns it to the PrimaryShare field.
func (o *RoomingListShareReservationType) SetPrimaryShare(v bool) {
	o.PrimaryShare = &v
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *RoomingListShareReservationType) GetTimeSpan() TimeSpanType {
	if o == nil || IsNil(o.TimeSpan) {
		var ret TimeSpanType
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomingListShareReservationType) GetTimeSpanOk() (*TimeSpanType, bool) {
	if o == nil || IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *RoomingListShareReservationType) HasTimeSpan() bool {
	if o != nil && !IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given TimeSpanType and assigns it to the TimeSpan field.
func (o *RoomingListShareReservationType) SetTimeSpan(v TimeSpanType) {
	o.TimeSpan = &v
}

func (o RoomingListShareReservationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomingListShareReservationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.TypeOfRateAmount) {
		toSerialize["typeOfRateAmount"] = o.TypeOfRateAmount
	}
	if !IsNil(o.PrimaryShare) {
		toSerialize["primaryShare"] = o.PrimaryShare
	}
	if !IsNil(o.TimeSpan) {
		toSerialize["timeSpan"] = o.TimeSpan
	}
	return toSerialize, nil
}

type NullableRoomingListShareReservationType struct {
	value *RoomingListShareReservationType
	isSet bool
}

func (v NullableRoomingListShareReservationType) Get() *RoomingListShareReservationType {
	return v.value
}

func (v *NullableRoomingListShareReservationType) Set(val *RoomingListShareReservationType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomingListShareReservationType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomingListShareReservationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomingListShareReservationType(val *RoomingListShareReservationType) *NullableRoomingListShareReservationType {
	return &NullableRoomingListShareReservationType{value: val, isSet: true}
}

func (v NullableRoomingListShareReservationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomingListShareReservationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


