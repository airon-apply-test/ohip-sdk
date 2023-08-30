/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
)

// checks if the ChannelSellLimitsByDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelSellLimitsByDate{}

// ChannelSellLimitsByDate Request object to create or update sell limits for the channel or channel room type by day. Advanced logic is implemented to combine consecutive blocks of sell limits into a single schedule which have the same limits configured. Existing schedules can also be split as needed to account for overlapping schedules.
type ChannelSellLimitsByDate struct {
	SellLimits *ChannelSellLimitsByDateType `json:"sellLimits,omitempty"`
	// Flag to indicate whether any overlapping schedules should be automatically adjusted (split, truncated, etc.) as needed.
	AdjustOverlappingSchedules *bool `json:"adjustOverlappingSchedules,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChannelSellLimitsByDate instantiates a new ChannelSellLimitsByDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSellLimitsByDate() *ChannelSellLimitsByDate {
	this := ChannelSellLimitsByDate{}
	return &this
}

// NewChannelSellLimitsByDateWithDefaults instantiates a new ChannelSellLimitsByDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSellLimitsByDateWithDefaults() *ChannelSellLimitsByDate {
	this := ChannelSellLimitsByDate{}
	return &this
}

// GetSellLimits returns the SellLimits field value if set, zero value otherwise.
func (o *ChannelSellLimitsByDate) GetSellLimits() ChannelSellLimitsByDateType {
	if o == nil || IsNil(o.SellLimits) {
		var ret ChannelSellLimitsByDateType
		return ret
	}
	return *o.SellLimits
}

// GetSellLimitsOk returns a tuple with the SellLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitsByDate) GetSellLimitsOk() (*ChannelSellLimitsByDateType, bool) {
	if o == nil || IsNil(o.SellLimits) {
		return nil, false
	}
	return o.SellLimits, true
}

// HasSellLimits returns a boolean if a field has been set.
func (o *ChannelSellLimitsByDate) HasSellLimits() bool {
	if o != nil && !IsNil(o.SellLimits) {
		return true
	}

	return false
}

// SetSellLimits gets a reference to the given ChannelSellLimitsByDateType and assigns it to the SellLimits field.
func (o *ChannelSellLimitsByDate) SetSellLimits(v ChannelSellLimitsByDateType) {
	o.SellLimits = &v
}

// GetAdjustOverlappingSchedules returns the AdjustOverlappingSchedules field value if set, zero value otherwise.
func (o *ChannelSellLimitsByDate) GetAdjustOverlappingSchedules() bool {
	if o == nil || IsNil(o.AdjustOverlappingSchedules) {
		var ret bool
		return ret
	}
	return *o.AdjustOverlappingSchedules
}

// GetAdjustOverlappingSchedulesOk returns a tuple with the AdjustOverlappingSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitsByDate) GetAdjustOverlappingSchedulesOk() (*bool, bool) {
	if o == nil || IsNil(o.AdjustOverlappingSchedules) {
		return nil, false
	}
	return o.AdjustOverlappingSchedules, true
}

// HasAdjustOverlappingSchedules returns a boolean if a field has been set.
func (o *ChannelSellLimitsByDate) HasAdjustOverlappingSchedules() bool {
	if o != nil && !IsNil(o.AdjustOverlappingSchedules) {
		return true
	}

	return false
}

// SetAdjustOverlappingSchedules gets a reference to the given bool and assigns it to the AdjustOverlappingSchedules field.
func (o *ChannelSellLimitsByDate) SetAdjustOverlappingSchedules(v bool) {
	o.AdjustOverlappingSchedules = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChannelSellLimitsByDate) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitsByDate) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChannelSellLimitsByDate) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChannelSellLimitsByDate) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChannelSellLimitsByDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelSellLimitsByDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SellLimits) {
		toSerialize["sellLimits"] = o.SellLimits
	}
	if !IsNil(o.AdjustOverlappingSchedules) {
		toSerialize["adjustOverlappingSchedules"] = o.AdjustOverlappingSchedules
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChannelSellLimitsByDate struct {
	value *ChannelSellLimitsByDate
	isSet bool
}

func (v NullableChannelSellLimitsByDate) Get() *ChannelSellLimitsByDate {
	return v.value
}

func (v *NullableChannelSellLimitsByDate) Set(val *ChannelSellLimitsByDate) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSellLimitsByDate) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSellLimitsByDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSellLimitsByDate(val *ChannelSellLimitsByDate) *NullableChannelSellLimitsByDate {
	return &NullableChannelSellLimitsByDate{value: val, isSet: true}
}

func (v NullableChannelSellLimitsByDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSellLimitsByDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


