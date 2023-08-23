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
)

// checks if the HotelReservationTypeReservationProfiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelReservationTypeReservationProfiles{}

// HotelReservationTypeReservationProfiles Collection of guests associated with the reservation.
type HotelReservationTypeReservationProfiles struct {
	ReservationProfile []ReservationProfileType `json:"reservationProfile,omitempty"`
	CommissionPayoutTo *CommissionPayoutToType `json:"commissionPayoutTo,omitempty"`
}

// NewHotelReservationTypeReservationProfiles instantiates a new HotelReservationTypeReservationProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelReservationTypeReservationProfiles() *HotelReservationTypeReservationProfiles {
	this := HotelReservationTypeReservationProfiles{}
	return &this
}

// NewHotelReservationTypeReservationProfilesWithDefaults instantiates a new HotelReservationTypeReservationProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelReservationTypeReservationProfilesWithDefaults() *HotelReservationTypeReservationProfiles {
	this := HotelReservationTypeReservationProfiles{}
	return &this
}

// GetReservationProfile returns the ReservationProfile field value if set, zero value otherwise.
func (o *HotelReservationTypeReservationProfiles) GetReservationProfile() []ReservationProfileType {
	if o == nil || IsNil(o.ReservationProfile) {
		var ret []ReservationProfileType
		return ret
	}
	return o.ReservationProfile
}

// GetReservationProfileOk returns a tuple with the ReservationProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelReservationTypeReservationProfiles) GetReservationProfileOk() ([]ReservationProfileType, bool) {
	if o == nil || IsNil(o.ReservationProfile) {
		return nil, false
	}
	return o.ReservationProfile, true
}

// HasReservationProfile returns a boolean if a field has been set.
func (o *HotelReservationTypeReservationProfiles) HasReservationProfile() bool {
	if o != nil && !IsNil(o.ReservationProfile) {
		return true
	}

	return false
}

// SetReservationProfile gets a reference to the given []ReservationProfileType and assigns it to the ReservationProfile field.
func (o *HotelReservationTypeReservationProfiles) SetReservationProfile(v []ReservationProfileType) {
	o.ReservationProfile = v
}

// GetCommissionPayoutTo returns the CommissionPayoutTo field value if set, zero value otherwise.
func (o *HotelReservationTypeReservationProfiles) GetCommissionPayoutTo() CommissionPayoutToType {
	if o == nil || IsNil(o.CommissionPayoutTo) {
		var ret CommissionPayoutToType
		return ret
	}
	return *o.CommissionPayoutTo
}

// GetCommissionPayoutToOk returns a tuple with the CommissionPayoutTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelReservationTypeReservationProfiles) GetCommissionPayoutToOk() (*CommissionPayoutToType, bool) {
	if o == nil || IsNil(o.CommissionPayoutTo) {
		return nil, false
	}
	return o.CommissionPayoutTo, true
}

// HasCommissionPayoutTo returns a boolean if a field has been set.
func (o *HotelReservationTypeReservationProfiles) HasCommissionPayoutTo() bool {
	if o != nil && !IsNil(o.CommissionPayoutTo) {
		return true
	}

	return false
}

// SetCommissionPayoutTo gets a reference to the given CommissionPayoutToType and assigns it to the CommissionPayoutTo field.
func (o *HotelReservationTypeReservationProfiles) SetCommissionPayoutTo(v CommissionPayoutToType) {
	o.CommissionPayoutTo = &v
}

func (o HotelReservationTypeReservationProfiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelReservationTypeReservationProfiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationProfile) {
		toSerialize["reservationProfile"] = o.ReservationProfile
	}
	if !IsNil(o.CommissionPayoutTo) {
		toSerialize["commissionPayoutTo"] = o.CommissionPayoutTo
	}
	return toSerialize, nil
}

type NullableHotelReservationTypeReservationProfiles struct {
	value *HotelReservationTypeReservationProfiles
	isSet bool
}

func (v NullableHotelReservationTypeReservationProfiles) Get() *HotelReservationTypeReservationProfiles {
	return v.value
}

func (v *NullableHotelReservationTypeReservationProfiles) Set(val *HotelReservationTypeReservationProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelReservationTypeReservationProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelReservationTypeReservationProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelReservationTypeReservationProfiles(val *HotelReservationTypeReservationProfiles) *NullableHotelReservationTypeReservationProfiles {
	return &NullableHotelReservationTypeReservationProfiles{value: val, isSet: true}
}

func (v NullableHotelReservationTypeReservationProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelReservationTypeReservationProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


