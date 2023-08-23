/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ReservationLocatorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationLocatorType{}

// ReservationLocatorType Holds the information for a Reservation Guest Locator
type ReservationLocatorType struct {
	DateSpan *DateRangeType `json:"dateSpan,omitempty"`
	TimeSpan *DateTimeSpanType `json:"timeSpan,omitempty"`
	// The Locator Text for the guest.
	LocatorText *string `json:"locatorText,omitempty"`
	// Date and time of the Guest Locator.
	LocatorOn *time.Time `json:"locatorOn,omitempty"`
	// User that entered this Guest Locator.
	LocatorBy *string `json:"locatorBy,omitempty"`
	LocatorId *UniqueIDType `json:"locatorId,omitempty"`
}

// NewReservationLocatorType instantiates a new ReservationLocatorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationLocatorType() *ReservationLocatorType {
	this := ReservationLocatorType{}
	return &this
}

// NewReservationLocatorTypeWithDefaults instantiates a new ReservationLocatorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationLocatorTypeWithDefaults() *ReservationLocatorType {
	this := ReservationLocatorType{}
	return &this
}

// GetDateSpan returns the DateSpan field value if set, zero value otherwise.
func (o *ReservationLocatorType) GetDateSpan() DateRangeType {
	if o == nil || IsNil(o.DateSpan) {
		var ret DateRangeType
		return ret
	}
	return *o.DateSpan
}

// GetDateSpanOk returns a tuple with the DateSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationLocatorType) GetDateSpanOk() (*DateRangeType, bool) {
	if o == nil || IsNil(o.DateSpan) {
		return nil, false
	}
	return o.DateSpan, true
}

// HasDateSpan returns a boolean if a field has been set.
func (o *ReservationLocatorType) HasDateSpan() bool {
	if o != nil && !IsNil(o.DateSpan) {
		return true
	}

	return false
}

// SetDateSpan gets a reference to the given DateRangeType and assigns it to the DateSpan field.
func (o *ReservationLocatorType) SetDateSpan(v DateRangeType) {
	o.DateSpan = &v
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *ReservationLocatorType) GetTimeSpan() DateTimeSpanType {
	if o == nil || IsNil(o.TimeSpan) {
		var ret DateTimeSpanType
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationLocatorType) GetTimeSpanOk() (*DateTimeSpanType, bool) {
	if o == nil || IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *ReservationLocatorType) HasTimeSpan() bool {
	if o != nil && !IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given DateTimeSpanType and assigns it to the TimeSpan field.
func (o *ReservationLocatorType) SetTimeSpan(v DateTimeSpanType) {
	o.TimeSpan = &v
}

// GetLocatorText returns the LocatorText field value if set, zero value otherwise.
func (o *ReservationLocatorType) GetLocatorText() string {
	if o == nil || IsNil(o.LocatorText) {
		var ret string
		return ret
	}
	return *o.LocatorText
}

// GetLocatorTextOk returns a tuple with the LocatorText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationLocatorType) GetLocatorTextOk() (*string, bool) {
	if o == nil || IsNil(o.LocatorText) {
		return nil, false
	}
	return o.LocatorText, true
}

// HasLocatorText returns a boolean if a field has been set.
func (o *ReservationLocatorType) HasLocatorText() bool {
	if o != nil && !IsNil(o.LocatorText) {
		return true
	}

	return false
}

// SetLocatorText gets a reference to the given string and assigns it to the LocatorText field.
func (o *ReservationLocatorType) SetLocatorText(v string) {
	o.LocatorText = &v
}

// GetLocatorOn returns the LocatorOn field value if set, zero value otherwise.
func (o *ReservationLocatorType) GetLocatorOn() time.Time {
	if o == nil || IsNil(o.LocatorOn) {
		var ret time.Time
		return ret
	}
	return *o.LocatorOn
}

// GetLocatorOnOk returns a tuple with the LocatorOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationLocatorType) GetLocatorOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LocatorOn) {
		return nil, false
	}
	return o.LocatorOn, true
}

// HasLocatorOn returns a boolean if a field has been set.
func (o *ReservationLocatorType) HasLocatorOn() bool {
	if o != nil && !IsNil(o.LocatorOn) {
		return true
	}

	return false
}

// SetLocatorOn gets a reference to the given time.Time and assigns it to the LocatorOn field.
func (o *ReservationLocatorType) SetLocatorOn(v time.Time) {
	o.LocatorOn = &v
}

// GetLocatorBy returns the LocatorBy field value if set, zero value otherwise.
func (o *ReservationLocatorType) GetLocatorBy() string {
	if o == nil || IsNil(o.LocatorBy) {
		var ret string
		return ret
	}
	return *o.LocatorBy
}

// GetLocatorByOk returns a tuple with the LocatorBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationLocatorType) GetLocatorByOk() (*string, bool) {
	if o == nil || IsNil(o.LocatorBy) {
		return nil, false
	}
	return o.LocatorBy, true
}

// HasLocatorBy returns a boolean if a field has been set.
func (o *ReservationLocatorType) HasLocatorBy() bool {
	if o != nil && !IsNil(o.LocatorBy) {
		return true
	}

	return false
}

// SetLocatorBy gets a reference to the given string and assigns it to the LocatorBy field.
func (o *ReservationLocatorType) SetLocatorBy(v string) {
	o.LocatorBy = &v
}

// GetLocatorId returns the LocatorId field value if set, zero value otherwise.
func (o *ReservationLocatorType) GetLocatorId() UniqueIDType {
	if o == nil || IsNil(o.LocatorId) {
		var ret UniqueIDType
		return ret
	}
	return *o.LocatorId
}

// GetLocatorIdOk returns a tuple with the LocatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationLocatorType) GetLocatorIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.LocatorId) {
		return nil, false
	}
	return o.LocatorId, true
}

// HasLocatorId returns a boolean if a field has been set.
func (o *ReservationLocatorType) HasLocatorId() bool {
	if o != nil && !IsNil(o.LocatorId) {
		return true
	}

	return false
}

// SetLocatorId gets a reference to the given UniqueIDType and assigns it to the LocatorId field.
func (o *ReservationLocatorType) SetLocatorId(v UniqueIDType) {
	o.LocatorId = &v
}

func (o ReservationLocatorType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationLocatorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateSpan) {
		toSerialize["dateSpan"] = o.DateSpan
	}
	if !IsNil(o.TimeSpan) {
		toSerialize["timeSpan"] = o.TimeSpan
	}
	if !IsNil(o.LocatorText) {
		toSerialize["locatorText"] = o.LocatorText
	}
	if !IsNil(o.LocatorOn) {
		toSerialize["locatorOn"] = o.LocatorOn
	}
	if !IsNil(o.LocatorBy) {
		toSerialize["locatorBy"] = o.LocatorBy
	}
	if !IsNil(o.LocatorId) {
		toSerialize["locatorId"] = o.LocatorId
	}
	return toSerialize, nil
}

type NullableReservationLocatorType struct {
	value *ReservationLocatorType
	isSet bool
}

func (v NullableReservationLocatorType) Get() *ReservationLocatorType {
	return v.value
}

func (v *NullableReservationLocatorType) Set(val *ReservationLocatorType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationLocatorType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationLocatorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationLocatorType(val *ReservationLocatorType) *NullableReservationLocatorType {
	return &NullableReservationLocatorType{value: val, isSet: true}
}

func (v NullableReservationLocatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationLocatorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


