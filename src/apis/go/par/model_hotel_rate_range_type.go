/*
OPERA Cloud Price Availability Rate API

APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HotelRateRangeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelRateRangeType{}

// HotelRateRangeType The rate rage information of the hotel.
type HotelRateRangeType struct {
	TimeSpan *TimeSpanType `json:"timeSpan,omitempty"`
	// Minimum Rate offered by the hotel.
	MinRate *float32 `json:"minRate,omitempty"`
	// Maximum Rate offered by the hotel.
	MaxRate *float32 `json:"maxRate,omitempty"`
	// The base currency code for rate range(The currency code used by the hotel which the rate range belongs to).
	CurrencyCode *string `json:"currencyCode,omitempty"`
	HotelId *string `json:"hotelId,omitempty"`
}

// NewHotelRateRangeType instantiates a new HotelRateRangeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelRateRangeType() *HotelRateRangeType {
	this := HotelRateRangeType{}
	return &this
}

// NewHotelRateRangeTypeWithDefaults instantiates a new HotelRateRangeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelRateRangeTypeWithDefaults() *HotelRateRangeType {
	this := HotelRateRangeType{}
	return &this
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *HotelRateRangeType) GetTimeSpan() TimeSpanType {
	if o == nil || IsNil(o.TimeSpan) {
		var ret TimeSpanType
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelRateRangeType) GetTimeSpanOk() (*TimeSpanType, bool) {
	if o == nil || IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *HotelRateRangeType) HasTimeSpan() bool {
	if o != nil && !IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given TimeSpanType and assigns it to the TimeSpan field.
func (o *HotelRateRangeType) SetTimeSpan(v TimeSpanType) {
	o.TimeSpan = &v
}

// GetMinRate returns the MinRate field value if set, zero value otherwise.
func (o *HotelRateRangeType) GetMinRate() float32 {
	if o == nil || IsNil(o.MinRate) {
		var ret float32
		return ret
	}
	return *o.MinRate
}

// GetMinRateOk returns a tuple with the MinRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelRateRangeType) GetMinRateOk() (*float32, bool) {
	if o == nil || IsNil(o.MinRate) {
		return nil, false
	}
	return o.MinRate, true
}

// HasMinRate returns a boolean if a field has been set.
func (o *HotelRateRangeType) HasMinRate() bool {
	if o != nil && !IsNil(o.MinRate) {
		return true
	}

	return false
}

// SetMinRate gets a reference to the given float32 and assigns it to the MinRate field.
func (o *HotelRateRangeType) SetMinRate(v float32) {
	o.MinRate = &v
}

// GetMaxRate returns the MaxRate field value if set, zero value otherwise.
func (o *HotelRateRangeType) GetMaxRate() float32 {
	if o == nil || IsNil(o.MaxRate) {
		var ret float32
		return ret
	}
	return *o.MaxRate
}

// GetMaxRateOk returns a tuple with the MaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelRateRangeType) GetMaxRateOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRate) {
		return nil, false
	}
	return o.MaxRate, true
}

// HasMaxRate returns a boolean if a field has been set.
func (o *HotelRateRangeType) HasMaxRate() bool {
	if o != nil && !IsNil(o.MaxRate) {
		return true
	}

	return false
}

// SetMaxRate gets a reference to the given float32 and assigns it to the MaxRate field.
func (o *HotelRateRangeType) SetMaxRate(v float32) {
	o.MaxRate = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *HotelRateRangeType) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelRateRangeType) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *HotelRateRangeType) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *HotelRateRangeType) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *HotelRateRangeType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelRateRangeType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *HotelRateRangeType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *HotelRateRangeType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o HotelRateRangeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelRateRangeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeSpan) {
		toSerialize["timeSpan"] = o.TimeSpan
	}
	if !IsNil(o.MinRate) {
		toSerialize["minRate"] = o.MinRate
	}
	if !IsNil(o.MaxRate) {
		toSerialize["maxRate"] = o.MaxRate
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableHotelRateRangeType struct {
	value *HotelRateRangeType
	isSet bool
}

func (v NullableHotelRateRangeType) Get() *HotelRateRangeType {
	return v.value
}

func (v *NullableHotelRateRangeType) Set(val *HotelRateRangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelRateRangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelRateRangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelRateRangeType(val *HotelRateRangeType) *NullableHotelRateRangeType {
	return &NullableHotelRateRangeType{value: val, isSet: true}
}

func (v NullableHotelRateRangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelRateRangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


