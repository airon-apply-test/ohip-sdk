/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChannelSellLimitType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelSellLimitType{}

// ChannelSellLimitType struct for ChannelSellLimitType
type ChannelSellLimitType struct {
	// Number of rooms that can be sold for this particular sell limit date.
	NumberOfRooms *int32 `json:"numberOfRooms,omitempty"`
	// Channel for which this sell limit is applicable.
	BookingChannelCode *string `json:"bookingChannelCode,omitempty"`
	// Date on which this sell limit is applicable.
	Date *string `json:"date,omitempty"`
}

// NewChannelSellLimitType instantiates a new ChannelSellLimitType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSellLimitType() *ChannelSellLimitType {
	this := ChannelSellLimitType{}
	return &this
}

// NewChannelSellLimitTypeWithDefaults instantiates a new ChannelSellLimitType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSellLimitTypeWithDefaults() *ChannelSellLimitType {
	this := ChannelSellLimitType{}
	return &this
}

// GetNumberOfRooms returns the NumberOfRooms field value if set, zero value otherwise.
func (o *ChannelSellLimitType) GetNumberOfRooms() int32 {
	if o == nil || IsNil(o.NumberOfRooms) {
		var ret int32
		return ret
	}
	return *o.NumberOfRooms
}

// GetNumberOfRoomsOk returns a tuple with the NumberOfRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitType) GetNumberOfRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfRooms) {
		return nil, false
	}
	return o.NumberOfRooms, true
}

// HasNumberOfRooms returns a boolean if a field has been set.
func (o *ChannelSellLimitType) HasNumberOfRooms() bool {
	if o != nil && !IsNil(o.NumberOfRooms) {
		return true
	}

	return false
}

// SetNumberOfRooms gets a reference to the given int32 and assigns it to the NumberOfRooms field.
func (o *ChannelSellLimitType) SetNumberOfRooms(v int32) {
	o.NumberOfRooms = &v
}

// GetBookingChannelCode returns the BookingChannelCode field value if set, zero value otherwise.
func (o *ChannelSellLimitType) GetBookingChannelCode() string {
	if o == nil || IsNil(o.BookingChannelCode) {
		var ret string
		return ret
	}
	return *o.BookingChannelCode
}

// GetBookingChannelCodeOk returns a tuple with the BookingChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitType) GetBookingChannelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BookingChannelCode) {
		return nil, false
	}
	return o.BookingChannelCode, true
}

// HasBookingChannelCode returns a boolean if a field has been set.
func (o *ChannelSellLimitType) HasBookingChannelCode() bool {
	if o != nil && !IsNil(o.BookingChannelCode) {
		return true
	}

	return false
}

// SetBookingChannelCode gets a reference to the given string and assigns it to the BookingChannelCode field.
func (o *ChannelSellLimitType) SetBookingChannelCode(v string) {
	o.BookingChannelCode = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ChannelSellLimitType) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitType) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ChannelSellLimitType) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ChannelSellLimitType) SetDate(v string) {
	o.Date = &v
}

func (o ChannelSellLimitType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelSellLimitType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfRooms) {
		toSerialize["numberOfRooms"] = o.NumberOfRooms
	}
	if !IsNil(o.BookingChannelCode) {
		toSerialize["bookingChannelCode"] = o.BookingChannelCode
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

type NullableChannelSellLimitType struct {
	value *ChannelSellLimitType
	isSet bool
}

func (v NullableChannelSellLimitType) Get() *ChannelSellLimitType {
	return v.value
}

func (v *NullableChannelSellLimitType) Set(val *ChannelSellLimitType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSellLimitType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSellLimitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSellLimitType(val *ChannelSellLimitType) *NullableChannelSellLimitType {
	return &NullableChannelSellLimitType{value: val, isSet: true}
}

func (v NullableChannelSellLimitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSellLimitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


