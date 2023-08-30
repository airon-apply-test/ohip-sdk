/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtp

import (
	"encoding/json"
)

// checks if the RateByAgeBucketType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateByAgeBucketType{}

// RateByAgeBucketType Rate amount by age bucket.
type RateByAgeBucketType struct {
	// Rate amount by age bucket.
	RateAmount *float32 `json:"rateAmount,omitempty"`
	// Minimum age for age bucket.
	MinimumAge *int32 `json:"minimumAge,omitempty"`
	// Maximum age for age bucket.
	MaximumAge *int32 `json:"maximumAge,omitempty"`
}

// NewRateByAgeBucketType instantiates a new RateByAgeBucketType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateByAgeBucketType() *RateByAgeBucketType {
	this := RateByAgeBucketType{}
	return &this
}

// NewRateByAgeBucketTypeWithDefaults instantiates a new RateByAgeBucketType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateByAgeBucketTypeWithDefaults() *RateByAgeBucketType {
	this := RateByAgeBucketType{}
	return &this
}

// GetRateAmount returns the RateAmount field value if set, zero value otherwise.
func (o *RateByAgeBucketType) GetRateAmount() float32 {
	if o == nil || IsNil(o.RateAmount) {
		var ret float32
		return ret
	}
	return *o.RateAmount
}

// GetRateAmountOk returns a tuple with the RateAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateByAgeBucketType) GetRateAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RateAmount) {
		return nil, false
	}
	return o.RateAmount, true
}

// HasRateAmount returns a boolean if a field has been set.
func (o *RateByAgeBucketType) HasRateAmount() bool {
	if o != nil && !IsNil(o.RateAmount) {
		return true
	}

	return false
}

// SetRateAmount gets a reference to the given float32 and assigns it to the RateAmount field.
func (o *RateByAgeBucketType) SetRateAmount(v float32) {
	o.RateAmount = &v
}

// GetMinimumAge returns the MinimumAge field value if set, zero value otherwise.
func (o *RateByAgeBucketType) GetMinimumAge() int32 {
	if o == nil || IsNil(o.MinimumAge) {
		var ret int32
		return ret
	}
	return *o.MinimumAge
}

// GetMinimumAgeOk returns a tuple with the MinimumAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateByAgeBucketType) GetMinimumAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumAge) {
		return nil, false
	}
	return o.MinimumAge, true
}

// HasMinimumAge returns a boolean if a field has been set.
func (o *RateByAgeBucketType) HasMinimumAge() bool {
	if o != nil && !IsNil(o.MinimumAge) {
		return true
	}

	return false
}

// SetMinimumAge gets a reference to the given int32 and assigns it to the MinimumAge field.
func (o *RateByAgeBucketType) SetMinimumAge(v int32) {
	o.MinimumAge = &v
}

// GetMaximumAge returns the MaximumAge field value if set, zero value otherwise.
func (o *RateByAgeBucketType) GetMaximumAge() int32 {
	if o == nil || IsNil(o.MaximumAge) {
		var ret int32
		return ret
	}
	return *o.MaximumAge
}

// GetMaximumAgeOk returns a tuple with the MaximumAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateByAgeBucketType) GetMaximumAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumAge) {
		return nil, false
	}
	return o.MaximumAge, true
}

// HasMaximumAge returns a boolean if a field has been set.
func (o *RateByAgeBucketType) HasMaximumAge() bool {
	if o != nil && !IsNil(o.MaximumAge) {
		return true
	}

	return false
}

// SetMaximumAge gets a reference to the given int32 and assigns it to the MaximumAge field.
func (o *RateByAgeBucketType) SetMaximumAge(v int32) {
	o.MaximumAge = &v
}

func (o RateByAgeBucketType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateByAgeBucketType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RateAmount) {
		toSerialize["rateAmount"] = o.RateAmount
	}
	if !IsNil(o.MinimumAge) {
		toSerialize["minimumAge"] = o.MinimumAge
	}
	if !IsNil(o.MaximumAge) {
		toSerialize["maximumAge"] = o.MaximumAge
	}
	return toSerialize, nil
}

type NullableRateByAgeBucketType struct {
	value *RateByAgeBucketType
	isSet bool
}

func (v NullableRateByAgeBucketType) Get() *RateByAgeBucketType {
	return v.value
}

func (v *NullableRateByAgeBucketType) Set(val *RateByAgeBucketType) {
	v.value = val
	v.isSet = true
}

func (v NullableRateByAgeBucketType) IsSet() bool {
	return v.isSet
}

func (v *NullableRateByAgeBucketType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateByAgeBucketType(val *RateByAgeBucketType) *NullableRateByAgeBucketType {
	return &NullableRateByAgeBucketType{value: val, isSet: true}
}

func (v NullableRateByAgeBucketType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateByAgeBucketType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


