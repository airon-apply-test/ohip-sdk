/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
)

// checks if the BookingStatusType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookingStatusType{}

// BookingStatusType Booking status code, description and configured color of the status. This will be used for block and event statuses.
type BookingStatusType struct {
	Status *CodeDescriptionType `json:"status,omitempty"`
	Color *StatusColorType `json:"color,omitempty"`
}

// NewBookingStatusType instantiates a new BookingStatusType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookingStatusType() *BookingStatusType {
	this := BookingStatusType{}
	return &this
}

// NewBookingStatusTypeWithDefaults instantiates a new BookingStatusType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingStatusTypeWithDefaults() *BookingStatusType {
	this := BookingStatusType{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BookingStatusType) GetStatus() CodeDescriptionType {
	if o == nil || IsNil(o.Status) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingStatusType) GetStatusOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BookingStatusType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CodeDescriptionType and assigns it to the Status field.
func (o *BookingStatusType) SetStatus(v CodeDescriptionType) {
	o.Status = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *BookingStatusType) GetColor() StatusColorType {
	if o == nil || IsNil(o.Color) {
		var ret StatusColorType
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingStatusType) GetColorOk() (*StatusColorType, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *BookingStatusType) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given StatusColorType and assigns it to the Color field.
func (o *BookingStatusType) SetColor(v StatusColorType) {
	o.Color = &v
}

func (o BookingStatusType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookingStatusType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	return toSerialize, nil
}

type NullableBookingStatusType struct {
	value *BookingStatusType
	isSet bool
}

func (v NullableBookingStatusType) Get() *BookingStatusType {
	return v.value
}

func (v *NullableBookingStatusType) Set(val *BookingStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableBookingStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableBookingStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookingStatusType(val *BookingStatusType) *NullableBookingStatusType {
	return &NullableBookingStatusType{value: val, isSet: true}
}

func (v NullableBookingStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookingStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


