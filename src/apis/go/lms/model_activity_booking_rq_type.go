/*
OPERA Cloud Leisure Management API

APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ActivityBookingRQType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityBookingRQType{}

// ActivityBookingRQType The choice between a reservation header or a profile ID. One or the other is required.
type ActivityBookingRQType struct {
	// This is not required if a Reservation Id is provided
	HotelId *string `json:"hotelId,omitempty"`
	ReservationId *UniqueIDType `json:"reservationId,omitempty"`
	ProfileId *UniqueIDType `json:"profileId,omitempty"`
	// A collection of Activity objects.
	Activities []ActivityListInner `json:"activities,omitempty"`
}

// NewActivityBookingRQType instantiates a new ActivityBookingRQType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityBookingRQType() *ActivityBookingRQType {
	this := ActivityBookingRQType{}
	return &this
}

// NewActivityBookingRQTypeWithDefaults instantiates a new ActivityBookingRQType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityBookingRQTypeWithDefaults() *ActivityBookingRQType {
	this := ActivityBookingRQType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ActivityBookingRQType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityBookingRQType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ActivityBookingRQType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ActivityBookingRQType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *ActivityBookingRQType) GetReservationId() UniqueIDType {
	if o == nil || IsNil(o.ReservationId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityBookingRQType) GetReservationIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *ActivityBookingRQType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given UniqueIDType and assigns it to the ReservationId field.
func (o *ActivityBookingRQType) SetReservationId(v UniqueIDType) {
	o.ReservationId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ActivityBookingRQType) GetProfileId() UniqueIDType {
	if o == nil || IsNil(o.ProfileId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityBookingRQType) GetProfileIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ActivityBookingRQType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given UniqueIDType and assigns it to the ProfileId field.
func (o *ActivityBookingRQType) SetProfileId(v UniqueIDType) {
	o.ProfileId = &v
}

// GetActivities returns the Activities field value if set, zero value otherwise.
func (o *ActivityBookingRQType) GetActivities() []ActivityListInner {
	if o == nil || IsNil(o.Activities) {
		var ret []ActivityListInner
		return ret
	}
	return o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityBookingRQType) GetActivitiesOk() ([]ActivityListInner, bool) {
	if o == nil || IsNil(o.Activities) {
		return nil, false
	}
	return o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *ActivityBookingRQType) HasActivities() bool {
	if o != nil && !IsNil(o.Activities) {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []ActivityListInner and assigns it to the Activities field.
func (o *ActivityBookingRQType) SetActivities(v []ActivityListInner) {
	o.Activities = v
}

func (o ActivityBookingRQType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityBookingRQType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.Activities) {
		toSerialize["activities"] = o.Activities
	}
	return toSerialize, nil
}

type NullableActivityBookingRQType struct {
	value *ActivityBookingRQType
	isSet bool
}

func (v NullableActivityBookingRQType) Get() *ActivityBookingRQType {
	return v.value
}

func (v *NullableActivityBookingRQType) Set(val *ActivityBookingRQType) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityBookingRQType) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityBookingRQType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityBookingRQType(val *ActivityBookingRQType) *NullableActivityBookingRQType {
	return &NullableActivityBookingRQType{value: val, isSet: true}
}

func (v NullableActivityBookingRQType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityBookingRQType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


