/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
)

// checks if the HotelInfoTypePropertyControlsDateTimeFormatting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelInfoTypePropertyControlsDateTimeFormatting{}

// HotelInfoTypePropertyControlsDateTimeFormatting Date Time Formatting information configuration of the hotel
type HotelInfoTypePropertyControlsDateTimeFormatting struct {
	// Long date format of the hotel.
	LongDateFormat *string `json:"longDateFormat,omitempty"`
	// Short date format of the hotel.
	ShortDateFormat *string `json:"shortDateFormat,omitempty"`
	// Time format for the hotel.
	TimeFormat *string `json:"timeFormat,omitempty"`
	// Time zone region of the hotel.
	TimeZoneRegion *string `json:"timeZoneRegion,omitempty"`
}

// NewHotelInfoTypePropertyControlsDateTimeFormatting instantiates a new HotelInfoTypePropertyControlsDateTimeFormatting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelInfoTypePropertyControlsDateTimeFormatting() *HotelInfoTypePropertyControlsDateTimeFormatting {
	this := HotelInfoTypePropertyControlsDateTimeFormatting{}
	return &this
}

// NewHotelInfoTypePropertyControlsDateTimeFormattingWithDefaults instantiates a new HotelInfoTypePropertyControlsDateTimeFormatting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelInfoTypePropertyControlsDateTimeFormattingWithDefaults() *HotelInfoTypePropertyControlsDateTimeFormatting {
	this := HotelInfoTypePropertyControlsDateTimeFormatting{}
	return &this
}

// GetLongDateFormat returns the LongDateFormat field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) GetLongDateFormat() string {
	if o == nil || IsNil(o.LongDateFormat) {
		var ret string
		return ret
	}
	return *o.LongDateFormat
}

// GetLongDateFormatOk returns a tuple with the LongDateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) GetLongDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.LongDateFormat) {
		return nil, false
	}
	return o.LongDateFormat, true
}

// HasLongDateFormat returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) HasLongDateFormat() bool {
	if o != nil && !IsNil(o.LongDateFormat) {
		return true
	}

	return false
}

// SetLongDateFormat gets a reference to the given string and assigns it to the LongDateFormat field.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) SetLongDateFormat(v string) {
	o.LongDateFormat = &v
}

// GetShortDateFormat returns the ShortDateFormat field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) GetShortDateFormat() string {
	if o == nil || IsNil(o.ShortDateFormat) {
		var ret string
		return ret
	}
	return *o.ShortDateFormat
}

// GetShortDateFormatOk returns a tuple with the ShortDateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) GetShortDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDateFormat) {
		return nil, false
	}
	return o.ShortDateFormat, true
}

// HasShortDateFormat returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) HasShortDateFormat() bool {
	if o != nil && !IsNil(o.ShortDateFormat) {
		return true
	}

	return false
}

// SetShortDateFormat gets a reference to the given string and assigns it to the ShortDateFormat field.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) SetShortDateFormat(v string) {
	o.ShortDateFormat = &v
}

// GetTimeFormat returns the TimeFormat field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) GetTimeFormat() string {
	if o == nil || IsNil(o.TimeFormat) {
		var ret string
		return ret
	}
	return *o.TimeFormat
}

// GetTimeFormatOk returns a tuple with the TimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) GetTimeFormatOk() (*string, bool) {
	if o == nil || IsNil(o.TimeFormat) {
		return nil, false
	}
	return o.TimeFormat, true
}

// HasTimeFormat returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) HasTimeFormat() bool {
	if o != nil && !IsNil(o.TimeFormat) {
		return true
	}

	return false
}

// SetTimeFormat gets a reference to the given string and assigns it to the TimeFormat field.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) SetTimeFormat(v string) {
	o.TimeFormat = &v
}

// GetTimeZoneRegion returns the TimeZoneRegion field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) GetTimeZoneRegion() string {
	if o == nil || IsNil(o.TimeZoneRegion) {
		var ret string
		return ret
	}
	return *o.TimeZoneRegion
}

// GetTimeZoneRegionOk returns a tuple with the TimeZoneRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) GetTimeZoneRegionOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZoneRegion) {
		return nil, false
	}
	return o.TimeZoneRegion, true
}

// HasTimeZoneRegion returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) HasTimeZoneRegion() bool {
	if o != nil && !IsNil(o.TimeZoneRegion) {
		return true
	}

	return false
}

// SetTimeZoneRegion gets a reference to the given string and assigns it to the TimeZoneRegion field.
func (o *HotelInfoTypePropertyControlsDateTimeFormatting) SetTimeZoneRegion(v string) {
	o.TimeZoneRegion = &v
}

func (o HotelInfoTypePropertyControlsDateTimeFormatting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelInfoTypePropertyControlsDateTimeFormatting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LongDateFormat) {
		toSerialize["longDateFormat"] = o.LongDateFormat
	}
	if !IsNil(o.ShortDateFormat) {
		toSerialize["shortDateFormat"] = o.ShortDateFormat
	}
	if !IsNil(o.TimeFormat) {
		toSerialize["timeFormat"] = o.TimeFormat
	}
	if !IsNil(o.TimeZoneRegion) {
		toSerialize["timeZoneRegion"] = o.TimeZoneRegion
	}
	return toSerialize, nil
}

type NullableHotelInfoTypePropertyControlsDateTimeFormatting struct {
	value *HotelInfoTypePropertyControlsDateTimeFormatting
	isSet bool
}

func (v NullableHotelInfoTypePropertyControlsDateTimeFormatting) Get() *HotelInfoTypePropertyControlsDateTimeFormatting {
	return v.value
}

func (v *NullableHotelInfoTypePropertyControlsDateTimeFormatting) Set(val *HotelInfoTypePropertyControlsDateTimeFormatting) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelInfoTypePropertyControlsDateTimeFormatting) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelInfoTypePropertyControlsDateTimeFormatting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelInfoTypePropertyControlsDateTimeFormatting(val *HotelInfoTypePropertyControlsDateTimeFormatting) *NullableHotelInfoTypePropertyControlsDateTimeFormatting {
	return &NullableHotelInfoTypePropertyControlsDateTimeFormatting{value: val, isSet: true}
}

func (v NullableHotelInfoTypePropertyControlsDateTimeFormatting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelInfoTypePropertyControlsDateTimeFormatting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


