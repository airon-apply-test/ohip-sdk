/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
)

// checks if the ReservationQueueInformationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationQueueInformationType{}

// ReservationQueueInformationType Information regarding the Date, Duration and Priority of the reservation on Queue for Check in.
type ReservationQueueInformationType struct {
	TimeSpan *ReservationQueueInformationTypeTimeSpan `json:"timeSpan,omitempty"`
	GuestTextInfo *QueueTextInfoType `json:"guestTextInfo,omitempty"`
	// The Queue Priority given to this reservation.
	Priority *int32 `json:"priority,omitempty"`
	// The average time, in seconds, a reservation was on queue prior to Check-In.
	AverageQueueTimeToCheckIn *int32 `json:"averageQueueTimeToCheckIn,omitempty"`
	// The average time, in seconds, of the reservations currently in queue.
	AverageQueueTimeCurrentReservations *int32 `json:"averageQueueTimeCurrentReservations,omitempty"`
	// The Business date on which the reservation was due to arrive and is currently placed on Queue for Check In.
	QueueDate *string `json:"queueDate,omitempty"`
}

// NewReservationQueueInformationType instantiates a new ReservationQueueInformationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationQueueInformationType() *ReservationQueueInformationType {
	this := ReservationQueueInformationType{}
	return &this
}

// NewReservationQueueInformationTypeWithDefaults instantiates a new ReservationQueueInformationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationQueueInformationTypeWithDefaults() *ReservationQueueInformationType {
	this := ReservationQueueInformationType{}
	return &this
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *ReservationQueueInformationType) GetTimeSpan() ReservationQueueInformationTypeTimeSpan {
	if o == nil || IsNil(o.TimeSpan) {
		var ret ReservationQueueInformationTypeTimeSpan
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationQueueInformationType) GetTimeSpanOk() (*ReservationQueueInformationTypeTimeSpan, bool) {
	if o == nil || IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *ReservationQueueInformationType) HasTimeSpan() bool {
	if o != nil && !IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given ReservationQueueInformationTypeTimeSpan and assigns it to the TimeSpan field.
func (o *ReservationQueueInformationType) SetTimeSpan(v ReservationQueueInformationTypeTimeSpan) {
	o.TimeSpan = &v
}

// GetGuestTextInfo returns the GuestTextInfo field value if set, zero value otherwise.
func (o *ReservationQueueInformationType) GetGuestTextInfo() QueueTextInfoType {
	if o == nil || IsNil(o.GuestTextInfo) {
		var ret QueueTextInfoType
		return ret
	}
	return *o.GuestTextInfo
}

// GetGuestTextInfoOk returns a tuple with the GuestTextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationQueueInformationType) GetGuestTextInfoOk() (*QueueTextInfoType, bool) {
	if o == nil || IsNil(o.GuestTextInfo) {
		return nil, false
	}
	return o.GuestTextInfo, true
}

// HasGuestTextInfo returns a boolean if a field has been set.
func (o *ReservationQueueInformationType) HasGuestTextInfo() bool {
	if o != nil && !IsNil(o.GuestTextInfo) {
		return true
	}

	return false
}

// SetGuestTextInfo gets a reference to the given QueueTextInfoType and assigns it to the GuestTextInfo field.
func (o *ReservationQueueInformationType) SetGuestTextInfo(v QueueTextInfoType) {
	o.GuestTextInfo = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ReservationQueueInformationType) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationQueueInformationType) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ReservationQueueInformationType) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ReservationQueueInformationType) SetPriority(v int32) {
	o.Priority = &v
}

// GetAverageQueueTimeToCheckIn returns the AverageQueueTimeToCheckIn field value if set, zero value otherwise.
func (o *ReservationQueueInformationType) GetAverageQueueTimeToCheckIn() int32 {
	if o == nil || IsNil(o.AverageQueueTimeToCheckIn) {
		var ret int32
		return ret
	}
	return *o.AverageQueueTimeToCheckIn
}

// GetAverageQueueTimeToCheckInOk returns a tuple with the AverageQueueTimeToCheckIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationQueueInformationType) GetAverageQueueTimeToCheckInOk() (*int32, bool) {
	if o == nil || IsNil(o.AverageQueueTimeToCheckIn) {
		return nil, false
	}
	return o.AverageQueueTimeToCheckIn, true
}

// HasAverageQueueTimeToCheckIn returns a boolean if a field has been set.
func (o *ReservationQueueInformationType) HasAverageQueueTimeToCheckIn() bool {
	if o != nil && !IsNil(o.AverageQueueTimeToCheckIn) {
		return true
	}

	return false
}

// SetAverageQueueTimeToCheckIn gets a reference to the given int32 and assigns it to the AverageQueueTimeToCheckIn field.
func (o *ReservationQueueInformationType) SetAverageQueueTimeToCheckIn(v int32) {
	o.AverageQueueTimeToCheckIn = &v
}

// GetAverageQueueTimeCurrentReservations returns the AverageQueueTimeCurrentReservations field value if set, zero value otherwise.
func (o *ReservationQueueInformationType) GetAverageQueueTimeCurrentReservations() int32 {
	if o == nil || IsNil(o.AverageQueueTimeCurrentReservations) {
		var ret int32
		return ret
	}
	return *o.AverageQueueTimeCurrentReservations
}

// GetAverageQueueTimeCurrentReservationsOk returns a tuple with the AverageQueueTimeCurrentReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationQueueInformationType) GetAverageQueueTimeCurrentReservationsOk() (*int32, bool) {
	if o == nil || IsNil(o.AverageQueueTimeCurrentReservations) {
		return nil, false
	}
	return o.AverageQueueTimeCurrentReservations, true
}

// HasAverageQueueTimeCurrentReservations returns a boolean if a field has been set.
func (o *ReservationQueueInformationType) HasAverageQueueTimeCurrentReservations() bool {
	if o != nil && !IsNil(o.AverageQueueTimeCurrentReservations) {
		return true
	}

	return false
}

// SetAverageQueueTimeCurrentReservations gets a reference to the given int32 and assigns it to the AverageQueueTimeCurrentReservations field.
func (o *ReservationQueueInformationType) SetAverageQueueTimeCurrentReservations(v int32) {
	o.AverageQueueTimeCurrentReservations = &v
}

// GetQueueDate returns the QueueDate field value if set, zero value otherwise.
func (o *ReservationQueueInformationType) GetQueueDate() string {
	if o == nil || IsNil(o.QueueDate) {
		var ret string
		return ret
	}
	return *o.QueueDate
}

// GetQueueDateOk returns a tuple with the QueueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationQueueInformationType) GetQueueDateOk() (*string, bool) {
	if o == nil || IsNil(o.QueueDate) {
		return nil, false
	}
	return o.QueueDate, true
}

// HasQueueDate returns a boolean if a field has been set.
func (o *ReservationQueueInformationType) HasQueueDate() bool {
	if o != nil && !IsNil(o.QueueDate) {
		return true
	}

	return false
}

// SetQueueDate gets a reference to the given string and assigns it to the QueueDate field.
func (o *ReservationQueueInformationType) SetQueueDate(v string) {
	o.QueueDate = &v
}

func (o ReservationQueueInformationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationQueueInformationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeSpan) {
		toSerialize["timeSpan"] = o.TimeSpan
	}
	if !IsNil(o.GuestTextInfo) {
		toSerialize["guestTextInfo"] = o.GuestTextInfo
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.AverageQueueTimeToCheckIn) {
		toSerialize["averageQueueTimeToCheckIn"] = o.AverageQueueTimeToCheckIn
	}
	if !IsNil(o.AverageQueueTimeCurrentReservations) {
		toSerialize["averageQueueTimeCurrentReservations"] = o.AverageQueueTimeCurrentReservations
	}
	if !IsNil(o.QueueDate) {
		toSerialize["queueDate"] = o.QueueDate
	}
	return toSerialize, nil
}

type NullableReservationQueueInformationType struct {
	value *ReservationQueueInformationType
	isSet bool
}

func (v NullableReservationQueueInformationType) Get() *ReservationQueueInformationType {
	return v.value
}

func (v *NullableReservationQueueInformationType) Set(val *ReservationQueueInformationType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationQueueInformationType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationQueueInformationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationQueueInformationType(val *ReservationQueueInformationType) *NullableReservationQueueInformationType {
	return &NullableReservationQueueInformationType{value: val, isSet: true}
}

func (v NullableReservationQueueInformationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationQueueInformationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


