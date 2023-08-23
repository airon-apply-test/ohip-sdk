/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
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
	ResExpectedArrivalTime *time.Time `json:"resExpectedArrivalTime,omitempty"`
	// Departure Time
	ResExpectedDepartureTime *time.Time `json:"resExpectedDepartureTime,omitempty"`
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

// GetResExpectedArrivalTime returns the ResExpectedArrivalTime field value if set, zero value otherwise.
func (o *ResExpectedTimesType) GetResExpectedArrivalTime() time.Time {
	if o == nil || IsNil(o.ResExpectedArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.ResExpectedArrivalTime
}

// GetResExpectedArrivalTimeOk returns a tuple with the ResExpectedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResExpectedTimesType) GetResExpectedArrivalTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ResExpectedArrivalTime) {
		return nil, false
	}
	return o.ResExpectedArrivalTime, true
}

// HasResExpectedArrivalTime returns a boolean if a field has been set.
func (o *ResExpectedTimesType) HasResExpectedArrivalTime() bool {
	if o != nil && !IsNil(o.ResExpectedArrivalTime) {
		return true
	}

	return false
}

// SetResExpectedArrivalTime gets a reference to the given time.Time and assigns it to the ResExpectedArrivalTime field.
func (o *ResExpectedTimesType) SetResExpectedArrivalTime(v time.Time) {
	o.ResExpectedArrivalTime = &v
}

// GetResExpectedDepartureTime returns the ResExpectedDepartureTime field value if set, zero value otherwise.
func (o *ResExpectedTimesType) GetResExpectedDepartureTime() time.Time {
	if o == nil || IsNil(o.ResExpectedDepartureTime) {
		var ret time.Time
		return ret
	}
	return *o.ResExpectedDepartureTime
}

// GetResExpectedDepartureTimeOk returns a tuple with the ResExpectedDepartureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResExpectedTimesType) GetResExpectedDepartureTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ResExpectedDepartureTime) {
		return nil, false
	}
	return o.ResExpectedDepartureTime, true
}

// HasResExpectedDepartureTime returns a boolean if a field has been set.
func (o *ResExpectedTimesType) HasResExpectedDepartureTime() bool {
	if o != nil && !IsNil(o.ResExpectedDepartureTime) {
		return true
	}

	return false
}

// SetResExpectedDepartureTime gets a reference to the given time.Time and assigns it to the ResExpectedDepartureTime field.
func (o *ResExpectedTimesType) SetResExpectedDepartureTime(v time.Time) {
	o.ResExpectedDepartureTime = &v
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
	if !IsNil(o.ResExpectedArrivalTime) {
		toSerialize["resExpectedArrivalTime"] = o.ResExpectedArrivalTime
	}
	if !IsNil(o.ResExpectedDepartureTime) {
		toSerialize["resExpectedDepartureTime"] = o.ResExpectedDepartureTime
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


