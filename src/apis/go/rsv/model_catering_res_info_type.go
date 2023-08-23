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

// checks if the CateringResInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CateringResInfoType{}

// CateringResInfoType Information regarding catering event and catering revenue type associated to the reservation.
type CateringResInfoType struct {
	EventId *EventId `json:"eventId,omitempty"`
	// Catering revenue type associated to the reservation.
	RevenueType *string `json:"revenueType,omitempty"`
}

// NewCateringResInfoType instantiates a new CateringResInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCateringResInfoType() *CateringResInfoType {
	this := CateringResInfoType{}
	return &this
}

// NewCateringResInfoTypeWithDefaults instantiates a new CateringResInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCateringResInfoTypeWithDefaults() *CateringResInfoType {
	this := CateringResInfoType{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *CateringResInfoType) GetEventId() EventId {
	if o == nil || IsNil(o.EventId) {
		var ret EventId
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringResInfoType) GetEventIdOk() (*EventId, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *CateringResInfoType) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given EventId and assigns it to the EventId field.
func (o *CateringResInfoType) SetEventId(v EventId) {
	o.EventId = &v
}

// GetRevenueType returns the RevenueType field value if set, zero value otherwise.
func (o *CateringResInfoType) GetRevenueType() string {
	if o == nil || IsNil(o.RevenueType) {
		var ret string
		return ret
	}
	return *o.RevenueType
}

// GetRevenueTypeOk returns a tuple with the RevenueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringResInfoType) GetRevenueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RevenueType) {
		return nil, false
	}
	return o.RevenueType, true
}

// HasRevenueType returns a boolean if a field has been set.
func (o *CateringResInfoType) HasRevenueType() bool {
	if o != nil && !IsNil(o.RevenueType) {
		return true
	}

	return false
}

// SetRevenueType gets a reference to the given string and assigns it to the RevenueType field.
func (o *CateringResInfoType) SetRevenueType(v string) {
	o.RevenueType = &v
}

func (o CateringResInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CateringResInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.RevenueType) {
		toSerialize["revenueType"] = o.RevenueType
	}
	return toSerialize, nil
}

type NullableCateringResInfoType struct {
	value *CateringResInfoType
	isSet bool
}

func (v NullableCateringResInfoType) Get() *CateringResInfoType {
	return v.value
}

func (v *NullableCateringResInfoType) Set(val *CateringResInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableCateringResInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableCateringResInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCateringResInfoType(val *CateringResInfoType) *NullableCateringResInfoType {
	return &NullableCateringResInfoType{value: val, isSet: true}
}

func (v NullableCateringResInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCateringResInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


