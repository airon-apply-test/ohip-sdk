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
	"time"
)

// checks if the CheckoutInstructionsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutInstructionsType{}

// CheckoutInstructionsType struct for CheckoutInstructionsType
type CheckoutInstructionsType struct {
	// Housekeeping Status to be assigned to the room after checkout.
	RoomStatus *string `json:"roomStatus,omitempty"`
	// Date and time for scheduled check out which could be performed when SCHEDULED CHECKOUT is active. Room status will be assigned according to SCHEDULED CHECKOUT ROOM STATUS parameter value.
	ScheduleOn *time.Time `json:"scheduleOn,omitempty"`
	// Currency Code which the Guest preferred to use and to be stored on the reservation of the Guest.
	GuestPreferredCurrency *string `json:"guestPreferredCurrency,omitempty"`
	// Dictates whether to ignore warnings before check out.
	IgnoreWarnings *bool `json:"ignoreWarnings,omitempty"`
}

// NewCheckoutInstructionsType instantiates a new CheckoutInstructionsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutInstructionsType() *CheckoutInstructionsType {
	this := CheckoutInstructionsType{}
	return &this
}

// NewCheckoutInstructionsTypeWithDefaults instantiates a new CheckoutInstructionsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutInstructionsTypeWithDefaults() *CheckoutInstructionsType {
	this := CheckoutInstructionsType{}
	return &this
}

// GetRoomStatus returns the RoomStatus field value if set, zero value otherwise.
func (o *CheckoutInstructionsType) GetRoomStatus() string {
	if o == nil || IsNil(o.RoomStatus) {
		var ret string
		return ret
	}
	return *o.RoomStatus
}

// GetRoomStatusOk returns a tuple with the RoomStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutInstructionsType) GetRoomStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RoomStatus) {
		return nil, false
	}
	return o.RoomStatus, true
}

// HasRoomStatus returns a boolean if a field has been set.
func (o *CheckoutInstructionsType) HasRoomStatus() bool {
	if o != nil && !IsNil(o.RoomStatus) {
		return true
	}

	return false
}

// SetRoomStatus gets a reference to the given string and assigns it to the RoomStatus field.
func (o *CheckoutInstructionsType) SetRoomStatus(v string) {
	o.RoomStatus = &v
}

// GetScheduleOn returns the ScheduleOn field value if set, zero value otherwise.
func (o *CheckoutInstructionsType) GetScheduleOn() time.Time {
	if o == nil || IsNil(o.ScheduleOn) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleOn
}

// GetScheduleOnOk returns a tuple with the ScheduleOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutInstructionsType) GetScheduleOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduleOn) {
		return nil, false
	}
	return o.ScheduleOn, true
}

// HasScheduleOn returns a boolean if a field has been set.
func (o *CheckoutInstructionsType) HasScheduleOn() bool {
	if o != nil && !IsNil(o.ScheduleOn) {
		return true
	}

	return false
}

// SetScheduleOn gets a reference to the given time.Time and assigns it to the ScheduleOn field.
func (o *CheckoutInstructionsType) SetScheduleOn(v time.Time) {
	o.ScheduleOn = &v
}

// GetGuestPreferredCurrency returns the GuestPreferredCurrency field value if set, zero value otherwise.
func (o *CheckoutInstructionsType) GetGuestPreferredCurrency() string {
	if o == nil || IsNil(o.GuestPreferredCurrency) {
		var ret string
		return ret
	}
	return *o.GuestPreferredCurrency
}

// GetGuestPreferredCurrencyOk returns a tuple with the GuestPreferredCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutInstructionsType) GetGuestPreferredCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.GuestPreferredCurrency) {
		return nil, false
	}
	return o.GuestPreferredCurrency, true
}

// HasGuestPreferredCurrency returns a boolean if a field has been set.
func (o *CheckoutInstructionsType) HasGuestPreferredCurrency() bool {
	if o != nil && !IsNil(o.GuestPreferredCurrency) {
		return true
	}

	return false
}

// SetGuestPreferredCurrency gets a reference to the given string and assigns it to the GuestPreferredCurrency field.
func (o *CheckoutInstructionsType) SetGuestPreferredCurrency(v string) {
	o.GuestPreferredCurrency = &v
}

// GetIgnoreWarnings returns the IgnoreWarnings field value if set, zero value otherwise.
func (o *CheckoutInstructionsType) GetIgnoreWarnings() bool {
	if o == nil || IsNil(o.IgnoreWarnings) {
		var ret bool
		return ret
	}
	return *o.IgnoreWarnings
}

// GetIgnoreWarningsOk returns a tuple with the IgnoreWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutInstructionsType) GetIgnoreWarningsOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreWarnings) {
		return nil, false
	}
	return o.IgnoreWarnings, true
}

// HasIgnoreWarnings returns a boolean if a field has been set.
func (o *CheckoutInstructionsType) HasIgnoreWarnings() bool {
	if o != nil && !IsNil(o.IgnoreWarnings) {
		return true
	}

	return false
}

// SetIgnoreWarnings gets a reference to the given bool and assigns it to the IgnoreWarnings field.
func (o *CheckoutInstructionsType) SetIgnoreWarnings(v bool) {
	o.IgnoreWarnings = &v
}

func (o CheckoutInstructionsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutInstructionsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomStatus) {
		toSerialize["roomStatus"] = o.RoomStatus
	}
	if !IsNil(o.ScheduleOn) {
		toSerialize["scheduleOn"] = o.ScheduleOn
	}
	if !IsNil(o.GuestPreferredCurrency) {
		toSerialize["guestPreferredCurrency"] = o.GuestPreferredCurrency
	}
	if !IsNil(o.IgnoreWarnings) {
		toSerialize["ignoreWarnings"] = o.IgnoreWarnings
	}
	return toSerialize, nil
}

type NullableCheckoutInstructionsType struct {
	value *CheckoutInstructionsType
	isSet bool
}

func (v NullableCheckoutInstructionsType) Get() *CheckoutInstructionsType {
	return v.value
}

func (v *NullableCheckoutInstructionsType) Set(val *CheckoutInstructionsType) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutInstructionsType) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutInstructionsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutInstructionsType(val *CheckoutInstructionsType) *NullableCheckoutInstructionsType {
	return &NullableCheckoutInstructionsType{value: val, isSet: true}
}

func (v NullableCheckoutInstructionsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutInstructionsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


