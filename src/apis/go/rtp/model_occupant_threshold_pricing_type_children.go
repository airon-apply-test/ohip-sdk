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

// checks if the OccupantThresholdPricingTypeChildren type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OccupantThresholdPricingTypeChildren{}

// OccupantThresholdPricingTypeChildren Threshold for Children in the room.
type OccupantThresholdPricingTypeChildren struct {
	// Threshold value, after it is reached the corresponding amount will be charged.
	Threshold *int32 `json:"threshold,omitempty"`
	// Amount to be charged after the threshold is reached.
	Amount *float32 `json:"amount,omitempty"`
	// When rates are Defined by Age buckets, should the 1st buckets children count be excluded from threshold pricing.
	ExcludeBucket1 *bool `json:"excludeBucket1,omitempty"`
	// When rates are Defined by Age buckets, should the 2nd buckets children count be excluded from threshold pricing.
	ExcludeBucket2 *bool `json:"excludeBucket2,omitempty"`
	// When rates are Defined by Age buckets, should the 3rd buckets children count be excluded from threshold pricing.
	ExcludeBucket3 *bool `json:"excludeBucket3,omitempty"`
}

// NewOccupantThresholdPricingTypeChildren instantiates a new OccupantThresholdPricingTypeChildren object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOccupantThresholdPricingTypeChildren() *OccupantThresholdPricingTypeChildren {
	this := OccupantThresholdPricingTypeChildren{}
	return &this
}

// NewOccupantThresholdPricingTypeChildrenWithDefaults instantiates a new OccupantThresholdPricingTypeChildren object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOccupantThresholdPricingTypeChildrenWithDefaults() *OccupantThresholdPricingTypeChildren {
	this := OccupantThresholdPricingTypeChildren{}
	return &this
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *OccupantThresholdPricingTypeChildren) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OccupantThresholdPricingTypeChildren) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *OccupantThresholdPricingTypeChildren) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *OccupantThresholdPricingTypeChildren) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OccupantThresholdPricingTypeChildren) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OccupantThresholdPricingTypeChildren) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OccupantThresholdPricingTypeChildren) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *OccupantThresholdPricingTypeChildren) SetAmount(v float32) {
	o.Amount = &v
}

// GetExcludeBucket1 returns the ExcludeBucket1 field value if set, zero value otherwise.
func (o *OccupantThresholdPricingTypeChildren) GetExcludeBucket1() bool {
	if o == nil || IsNil(o.ExcludeBucket1) {
		var ret bool
		return ret
	}
	return *o.ExcludeBucket1
}

// GetExcludeBucket1Ok returns a tuple with the ExcludeBucket1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OccupantThresholdPricingTypeChildren) GetExcludeBucket1Ok() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeBucket1) {
		return nil, false
	}
	return o.ExcludeBucket1, true
}

// HasExcludeBucket1 returns a boolean if a field has been set.
func (o *OccupantThresholdPricingTypeChildren) HasExcludeBucket1() bool {
	if o != nil && !IsNil(o.ExcludeBucket1) {
		return true
	}

	return false
}

// SetExcludeBucket1 gets a reference to the given bool and assigns it to the ExcludeBucket1 field.
func (o *OccupantThresholdPricingTypeChildren) SetExcludeBucket1(v bool) {
	o.ExcludeBucket1 = &v
}

// GetExcludeBucket2 returns the ExcludeBucket2 field value if set, zero value otherwise.
func (o *OccupantThresholdPricingTypeChildren) GetExcludeBucket2() bool {
	if o == nil || IsNil(o.ExcludeBucket2) {
		var ret bool
		return ret
	}
	return *o.ExcludeBucket2
}

// GetExcludeBucket2Ok returns a tuple with the ExcludeBucket2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OccupantThresholdPricingTypeChildren) GetExcludeBucket2Ok() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeBucket2) {
		return nil, false
	}
	return o.ExcludeBucket2, true
}

// HasExcludeBucket2 returns a boolean if a field has been set.
func (o *OccupantThresholdPricingTypeChildren) HasExcludeBucket2() bool {
	if o != nil && !IsNil(o.ExcludeBucket2) {
		return true
	}

	return false
}

// SetExcludeBucket2 gets a reference to the given bool and assigns it to the ExcludeBucket2 field.
func (o *OccupantThresholdPricingTypeChildren) SetExcludeBucket2(v bool) {
	o.ExcludeBucket2 = &v
}

// GetExcludeBucket3 returns the ExcludeBucket3 field value if set, zero value otherwise.
func (o *OccupantThresholdPricingTypeChildren) GetExcludeBucket3() bool {
	if o == nil || IsNil(o.ExcludeBucket3) {
		var ret bool
		return ret
	}
	return *o.ExcludeBucket3
}

// GetExcludeBucket3Ok returns a tuple with the ExcludeBucket3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OccupantThresholdPricingTypeChildren) GetExcludeBucket3Ok() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeBucket3) {
		return nil, false
	}
	return o.ExcludeBucket3, true
}

// HasExcludeBucket3 returns a boolean if a field has been set.
func (o *OccupantThresholdPricingTypeChildren) HasExcludeBucket3() bool {
	if o != nil && !IsNil(o.ExcludeBucket3) {
		return true
	}

	return false
}

// SetExcludeBucket3 gets a reference to the given bool and assigns it to the ExcludeBucket3 field.
func (o *OccupantThresholdPricingTypeChildren) SetExcludeBucket3(v bool) {
	o.ExcludeBucket3 = &v
}

func (o OccupantThresholdPricingTypeChildren) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OccupantThresholdPricingTypeChildren) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ExcludeBucket1) {
		toSerialize["excludeBucket1"] = o.ExcludeBucket1
	}
	if !IsNil(o.ExcludeBucket2) {
		toSerialize["excludeBucket2"] = o.ExcludeBucket2
	}
	if !IsNil(o.ExcludeBucket3) {
		toSerialize["excludeBucket3"] = o.ExcludeBucket3
	}
	return toSerialize, nil
}

type NullableOccupantThresholdPricingTypeChildren struct {
	value *OccupantThresholdPricingTypeChildren
	isSet bool
}

func (v NullableOccupantThresholdPricingTypeChildren) Get() *OccupantThresholdPricingTypeChildren {
	return v.value
}

func (v *NullableOccupantThresholdPricingTypeChildren) Set(val *OccupantThresholdPricingTypeChildren) {
	v.value = val
	v.isSet = true
}

func (v NullableOccupantThresholdPricingTypeChildren) IsSet() bool {
	return v.isSet
}

func (v *NullableOccupantThresholdPricingTypeChildren) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOccupantThresholdPricingTypeChildren(val *OccupantThresholdPricingTypeChildren) *NullableOccupantThresholdPricingTypeChildren {
	return &NullableOccupantThresholdPricingTypeChildren{value: val, isSet: true}
}

func (v NullableOccupantThresholdPricingTypeChildren) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOccupantThresholdPricingTypeChildren) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


