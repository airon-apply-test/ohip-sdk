/*
OPERA Cloud Customer Relationship Management API

APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ResExpectedTimesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResExpectedTimesType{}

// ResExpectedTimesType Holds the Arrival and Departure Time Information
type ResExpectedTimesType struct {
	// Arrival Time
	ReservationExpectedArrivalTime *time.Time `json:"reservationExpectedArrivalTime,omitempty"`
	// Departure Time
	ReservationExpectedDepartureTime *time.Time `json:"reservationExpectedDepartureTime,omitempty"`
}

// NewResExpectedTimesType instantiates a new ResExpectedTimesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResExpectedTimesType() *ResExpectedTimesType {
	this := ResExpectedTimesType{}
	return &this
}

// NewResExpectedTimesTypeWithDefaults instantiates a new ResExpectedTimesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResExpectedTimesTypeWithDefaults() *ResExpectedTimesType {
	this := ResExpectedTimesType{}
	return &this
}

// GetReservationExpectedArrivalTime returns the ReservationExpectedArrivalTime field value if set, zero value otherwise.
func (o *ResExpectedTimesType) GetReservationExpectedArrivalTime() time.Time {
	if o == nil || IsNil(o.ReservationExpectedArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.ReservationExpectedArrivalTime
}

// GetReservationExpectedArrivalTimeOk returns a tuple with the ReservationExpectedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResExpectedTimesType) GetReservationExpectedArrivalTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReservationExpectedArrivalTime) {
		return nil, false
	}
	return o.ReservationExpectedArrivalTime, true
}

// HasReservationExpectedArrivalTime returns a boolean if a field has been set.
func (o *ResExpectedTimesType) HasReservationExpectedArrivalTime() bool {
	if o != nil && !IsNil(o.ReservationExpectedArrivalTime) {
		return true
	}

	return false
}

// SetReservationExpectedArrivalTime gets a reference to the given time.Time and assigns it to the ReservationExpectedArrivalTime field.
func (o *ResExpectedTimesType) SetReservationExpectedArrivalTime(v time.Time) {
	o.ReservationExpectedArrivalTime = &v
}

// GetReservationExpectedDepartureTime returns the ReservationExpectedDepartureTime field value if set, zero value otherwise.
func (o *ResExpectedTimesType) GetReservationExpectedDepartureTime() time.Time {
	if o == nil || IsNil(o.ReservationExpectedDepartureTime) {
		var ret time.Time
		return ret
	}
	return *o.ReservationExpectedDepartureTime
}

// GetReservationExpectedDepartureTimeOk returns a tuple with the ReservationExpectedDepartureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResExpectedTimesType) GetReservationExpectedDepartureTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReservationExpectedDepartureTime) {
		return nil, false
	}
	return o.ReservationExpectedDepartureTime, true
}

// HasReservationExpectedDepartureTime returns a boolean if a field has been set.
func (o *ResExpectedTimesType) HasReservationExpectedDepartureTime() bool {
	if o != nil && !IsNil(o.ReservationExpectedDepartureTime) {
		return true
	}

	return false
}

// SetReservationExpectedDepartureTime gets a reference to the given time.Time and assigns it to the ReservationExpectedDepartureTime field.
func (o *ResExpectedTimesType) SetReservationExpectedDepartureTime(v time.Time) {
	o.ReservationExpectedDepartureTime = &v
}

func (o ResExpectedTimesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResExpectedTimesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationExpectedArrivalTime) {
		toSerialize["reservationExpectedArrivalTime"] = o.ReservationExpectedArrivalTime
	}
	if !IsNil(o.ReservationExpectedDepartureTime) {
		toSerialize["reservationExpectedDepartureTime"] = o.ReservationExpectedDepartureTime
	}
	return toSerialize, nil
}

type NullableResExpectedTimesType struct {
	value *ResExpectedTimesType
	isSet bool
}

func (v NullableResExpectedTimesType) Get() *ResExpectedTimesType {
	return v.value
}

func (v *NullableResExpectedTimesType) Set(val *ResExpectedTimesType) {
	v.value = val
	v.isSet = true
}

func (v NullableResExpectedTimesType) IsSet() bool {
	return v.isSet
}

func (v *NullableResExpectedTimesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResExpectedTimesType(val *ResExpectedTimesType) *NullableResExpectedTimesType {
	return &NullableResExpectedTimesType{value: val, isSet: true}
}

func (v NullableResExpectedTimesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResExpectedTimesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


