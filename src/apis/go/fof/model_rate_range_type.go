/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the RateRangeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateRangeType{}

// RateRangeType Rate Range details like maximum rate amount and minimum rate amount in each available rate category.
type RateRangeType struct {
	// The base amount charged for the accommodation or service.
	Base []TotalType `json:"base,omitempty"`
	// Rate Change Indicator.
	RateChange *bool `json:"rateChange,omitempty"`
}

// NewRateRangeType instantiates a new RateRangeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateRangeType() *RateRangeType {
	this := RateRangeType{}
	return &this
}

// NewRateRangeTypeWithDefaults instantiates a new RateRangeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateRangeTypeWithDefaults() *RateRangeType {
	this := RateRangeType{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *RateRangeType) GetBase() []TotalType {
	if o == nil || IsNil(o.Base) {
		var ret []TotalType
		return ret
	}
	return o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateRangeType) GetBaseOk() ([]TotalType, bool) {
	if o == nil || IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *RateRangeType) HasBase() bool {
	if o != nil && !IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given []TotalType and assigns it to the Base field.
func (o *RateRangeType) SetBase(v []TotalType) {
	o.Base = v
}

// GetRateChange returns the RateChange field value if set, zero value otherwise.
func (o *RateRangeType) GetRateChange() bool {
	if o == nil || IsNil(o.RateChange) {
		var ret bool
		return ret
	}
	return *o.RateChange
}

// GetRateChangeOk returns a tuple with the RateChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateRangeType) GetRateChangeOk() (*bool, bool) {
	if o == nil || IsNil(o.RateChange) {
		return nil, false
	}
	return o.RateChange, true
}

// HasRateChange returns a boolean if a field has been set.
func (o *RateRangeType) HasRateChange() bool {
	if o != nil && !IsNil(o.RateChange) {
		return true
	}

	return false
}

// SetRateChange gets a reference to the given bool and assigns it to the RateChange field.
func (o *RateRangeType) SetRateChange(v bool) {
	o.RateChange = &v
}

func (o RateRangeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateRangeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !IsNil(o.RateChange) {
		toSerialize["rateChange"] = o.RateChange
	}
	return toSerialize, nil
}

type NullableRateRangeType struct {
	value *RateRangeType
	isSet bool
}

func (v NullableRateRangeType) Get() *RateRangeType {
	return v.value
}

func (v *NullableRateRangeType) Set(val *RateRangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableRateRangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableRateRangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateRangeType(val *RateRangeType) *NullableRateRangeType {
	return &NullableRateRangeType{value: val, isSet: true}
}

func (v NullableRateRangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateRangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


