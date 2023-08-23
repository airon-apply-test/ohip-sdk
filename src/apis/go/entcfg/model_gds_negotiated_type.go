/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GdsNegotiatedType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GdsNegotiatedType{}

// GdsNegotiatedType This holds a list of GdsNegotiatedInfoType.
type GdsNegotiatedType struct {
	// List of channel negotiated rates for the profile.
	GdsNegotiatedInfoList []GdsNegotiatedInfoType `json:"gdsNegotiatedInfoList,omitempty"`
	// Booking Channel Code.
	BookingChannelCode *string `json:"bookingChannelCode,omitempty"`
	// Hotel Code.
	HotelId *string `json:"hotelId,omitempty"`
	// Channel Room Type.
	ChannelRatePlanCode *string `json:"channelRatePlanCode,omitempty"`
}

// NewGdsNegotiatedType instantiates a new GdsNegotiatedType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGdsNegotiatedType() *GdsNegotiatedType {
	this := GdsNegotiatedType{}
	return &this
}

// NewGdsNegotiatedTypeWithDefaults instantiates a new GdsNegotiatedType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGdsNegotiatedTypeWithDefaults() *GdsNegotiatedType {
	this := GdsNegotiatedType{}
	return &this
}

// GetGdsNegotiatedInfoList returns the GdsNegotiatedInfoList field value if set, zero value otherwise.
func (o *GdsNegotiatedType) GetGdsNegotiatedInfoList() []GdsNegotiatedInfoType {
	if o == nil || IsNil(o.GdsNegotiatedInfoList) {
		var ret []GdsNegotiatedInfoType
		return ret
	}
	return o.GdsNegotiatedInfoList
}

// GetGdsNegotiatedInfoListOk returns a tuple with the GdsNegotiatedInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GdsNegotiatedType) GetGdsNegotiatedInfoListOk() ([]GdsNegotiatedInfoType, bool) {
	if o == nil || IsNil(o.GdsNegotiatedInfoList) {
		return nil, false
	}
	return o.GdsNegotiatedInfoList, true
}

// HasGdsNegotiatedInfoList returns a boolean if a field has been set.
func (o *GdsNegotiatedType) HasGdsNegotiatedInfoList() bool {
	if o != nil && !IsNil(o.GdsNegotiatedInfoList) {
		return true
	}

	return false
}

// SetGdsNegotiatedInfoList gets a reference to the given []GdsNegotiatedInfoType and assigns it to the GdsNegotiatedInfoList field.
func (o *GdsNegotiatedType) SetGdsNegotiatedInfoList(v []GdsNegotiatedInfoType) {
	o.GdsNegotiatedInfoList = v
}

// GetBookingChannelCode returns the BookingChannelCode field value if set, zero value otherwise.
func (o *GdsNegotiatedType) GetBookingChannelCode() string {
	if o == nil || IsNil(o.BookingChannelCode) {
		var ret string
		return ret
	}
	return *o.BookingChannelCode
}

// GetBookingChannelCodeOk returns a tuple with the BookingChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GdsNegotiatedType) GetBookingChannelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BookingChannelCode) {
		return nil, false
	}
	return o.BookingChannelCode, true
}

// HasBookingChannelCode returns a boolean if a field has been set.
func (o *GdsNegotiatedType) HasBookingChannelCode() bool {
	if o != nil && !IsNil(o.BookingChannelCode) {
		return true
	}

	return false
}

// SetBookingChannelCode gets a reference to the given string and assigns it to the BookingChannelCode field.
func (o *GdsNegotiatedType) SetBookingChannelCode(v string) {
	o.BookingChannelCode = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *GdsNegotiatedType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GdsNegotiatedType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *GdsNegotiatedType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *GdsNegotiatedType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetChannelRatePlanCode returns the ChannelRatePlanCode field value if set, zero value otherwise.
func (o *GdsNegotiatedType) GetChannelRatePlanCode() string {
	if o == nil || IsNil(o.ChannelRatePlanCode) {
		var ret string
		return ret
	}
	return *o.ChannelRatePlanCode
}

// GetChannelRatePlanCodeOk returns a tuple with the ChannelRatePlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GdsNegotiatedType) GetChannelRatePlanCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelRatePlanCode) {
		return nil, false
	}
	return o.ChannelRatePlanCode, true
}

// HasChannelRatePlanCode returns a boolean if a field has been set.
func (o *GdsNegotiatedType) HasChannelRatePlanCode() bool {
	if o != nil && !IsNil(o.ChannelRatePlanCode) {
		return true
	}

	return false
}

// SetChannelRatePlanCode gets a reference to the given string and assigns it to the ChannelRatePlanCode field.
func (o *GdsNegotiatedType) SetChannelRatePlanCode(v string) {
	o.ChannelRatePlanCode = &v
}

func (o GdsNegotiatedType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GdsNegotiatedType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GdsNegotiatedInfoList) {
		toSerialize["gdsNegotiatedInfoList"] = o.GdsNegotiatedInfoList
	}
	if !IsNil(o.BookingChannelCode) {
		toSerialize["bookingChannelCode"] = o.BookingChannelCode
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ChannelRatePlanCode) {
		toSerialize["channelRatePlanCode"] = o.ChannelRatePlanCode
	}
	return toSerialize, nil
}

type NullableGdsNegotiatedType struct {
	value *GdsNegotiatedType
	isSet bool
}

func (v NullableGdsNegotiatedType) Get() *GdsNegotiatedType {
	return v.value
}

func (v *NullableGdsNegotiatedType) Set(val *GdsNegotiatedType) {
	v.value = val
	v.isSet = true
}

func (v NullableGdsNegotiatedType) IsSet() bool {
	return v.isSet
}

func (v *NullableGdsNegotiatedType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGdsNegotiatedType(val *GdsNegotiatedType) *NullableGdsNegotiatedType {
	return &NullableGdsNegotiatedType{value: val, isSet: true}
}

func (v NullableGdsNegotiatedType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGdsNegotiatedType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


