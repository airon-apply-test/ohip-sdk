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

// checks if the ChannelSellLimitsByDateRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelSellLimitsByDateRange{}

// ChannelSellLimitsByDateRange Request object to create or update sell limit schedules for the channel or channel room type within a date range. Existing schedules can be split as needed to account for overlapping schedules.
type ChannelSellLimitsByDateRange struct {
	// Details about a sell limit schedule for a channel or channel room type.
	SellLimits []ChannelSellLimitScheduleType `json:"sellLimits,omitempty"`
	// Flag to indicate whether any overlapping schedules should be automatically adjusted (split, truncated, etc.) as needed.
	AdjustOverlappingSchedules *bool `json:"adjustOverlappingSchedules,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChannelSellLimitsByDateRange instantiates a new ChannelSellLimitsByDateRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSellLimitsByDateRange() *ChannelSellLimitsByDateRange {
	this := ChannelSellLimitsByDateRange{}
	return &this
}

// NewChannelSellLimitsByDateRangeWithDefaults instantiates a new ChannelSellLimitsByDateRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSellLimitsByDateRangeWithDefaults() *ChannelSellLimitsByDateRange {
	this := ChannelSellLimitsByDateRange{}
	return &this
}

// GetSellLimits returns the SellLimits field value if set, zero value otherwise.
func (o *ChannelSellLimitsByDateRange) GetSellLimits() []ChannelSellLimitScheduleType {
	if o == nil || IsNil(o.SellLimits) {
		var ret []ChannelSellLimitScheduleType
		return ret
	}
	return o.SellLimits
}

// GetSellLimitsOk returns a tuple with the SellLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitsByDateRange) GetSellLimitsOk() ([]ChannelSellLimitScheduleType, bool) {
	if o == nil || IsNil(o.SellLimits) {
		return nil, false
	}
	return o.SellLimits, true
}

// HasSellLimits returns a boolean if a field has been set.
func (o *ChannelSellLimitsByDateRange) HasSellLimits() bool {
	if o != nil && !IsNil(o.SellLimits) {
		return true
	}

	return false
}

// SetSellLimits gets a reference to the given []ChannelSellLimitScheduleType and assigns it to the SellLimits field.
func (o *ChannelSellLimitsByDateRange) SetSellLimits(v []ChannelSellLimitScheduleType) {
	o.SellLimits = v
}

// GetAdjustOverlappingSchedules returns the AdjustOverlappingSchedules field value if set, zero value otherwise.
func (o *ChannelSellLimitsByDateRange) GetAdjustOverlappingSchedules() bool {
	if o == nil || IsNil(o.AdjustOverlappingSchedules) {
		var ret bool
		return ret
	}
	return *o.AdjustOverlappingSchedules
}

// GetAdjustOverlappingSchedulesOk returns a tuple with the AdjustOverlappingSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitsByDateRange) GetAdjustOverlappingSchedulesOk() (*bool, bool) {
	if o == nil || IsNil(o.AdjustOverlappingSchedules) {
		return nil, false
	}
	return o.AdjustOverlappingSchedules, true
}

// HasAdjustOverlappingSchedules returns a boolean if a field has been set.
func (o *ChannelSellLimitsByDateRange) HasAdjustOverlappingSchedules() bool {
	if o != nil && !IsNil(o.AdjustOverlappingSchedules) {
		return true
	}

	return false
}

// SetAdjustOverlappingSchedules gets a reference to the given bool and assigns it to the AdjustOverlappingSchedules field.
func (o *ChannelSellLimitsByDateRange) SetAdjustOverlappingSchedules(v bool) {
	o.AdjustOverlappingSchedules = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChannelSellLimitsByDateRange) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSellLimitsByDateRange) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChannelSellLimitsByDateRange) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChannelSellLimitsByDateRange) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChannelSellLimitsByDateRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelSellLimitsByDateRange) ToMap() (map[string]interface{}, error) {
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

type NullableChannelSellLimitsByDateRange struct {
	value *ChannelSellLimitsByDateRange
	isSet bool
}

func (v NullableChannelSellLimitsByDateRange) Get() *ChannelSellLimitsByDateRange {
	return v.value
}

func (v *NullableChannelSellLimitsByDateRange) Set(val *ChannelSellLimitsByDateRange) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSellLimitsByDateRange) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSellLimitsByDateRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSellLimitsByDateRange(val *ChannelSellLimitsByDateRange) *NullableChannelSellLimitsByDateRange {
	return &NullableChannelSellLimitsByDateRange{value: val, isSet: true}
}

func (v NullableChannelSellLimitsByDateRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSellLimitsByDateRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


