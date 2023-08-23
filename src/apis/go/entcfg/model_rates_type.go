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

// checks if the RatesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatesType{}

// RatesType Individual rate amount.
type RatesType struct {
	// The Rate contains a collection of elements that define the amount of the rate, associated fees, additional occupant amounts. Taxes can be broken out or included within the various amounts. A currency can be associated to each amount.
	Rate []AmountType `json:"rate,omitempty"`
	// Rate Range details like maximum rate amount and minimum rate amount in each available rate category.
	RateRange []RateRangeType `json:"rateRange,omitempty"`
}

// NewRatesType instantiates a new RatesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatesType() *RatesType {
	this := RatesType{}
	return &this
}

// NewRatesTypeWithDefaults instantiates a new RatesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatesTypeWithDefaults() *RatesType {
	this := RatesType{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *RatesType) GetRate() []AmountType {
	if o == nil || IsNil(o.Rate) {
		var ret []AmountType
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatesType) GetRateOk() ([]AmountType, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *RatesType) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given []AmountType and assigns it to the Rate field.
func (o *RatesType) SetRate(v []AmountType) {
	o.Rate = v
}

// GetRateRange returns the RateRange field value if set, zero value otherwise.
func (o *RatesType) GetRateRange() []RateRangeType {
	if o == nil || IsNil(o.RateRange) {
		var ret []RateRangeType
		return ret
	}
	return o.RateRange
}

// GetRateRangeOk returns a tuple with the RateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatesType) GetRateRangeOk() ([]RateRangeType, bool) {
	if o == nil || IsNil(o.RateRange) {
		return nil, false
	}
	return o.RateRange, true
}

// HasRateRange returns a boolean if a field has been set.
func (o *RatesType) HasRateRange() bool {
	if o != nil && !IsNil(o.RateRange) {
		return true
	}

	return false
}

// SetRateRange gets a reference to the given []RateRangeType and assigns it to the RateRange field.
func (o *RatesType) SetRateRange(v []RateRangeType) {
	o.RateRange = v
}

func (o RatesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.RateRange) {
		toSerialize["rateRange"] = o.RateRange
	}
	return toSerialize, nil
}

type NullableRatesType struct {
	value *RatesType
	isSet bool
}

func (v NullableRatesType) Get() *RatesType {
	return v.value
}

func (v *NullableRatesType) Set(val *RatesType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatesType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatesType(val *RatesType) *NullableRatesType {
	return &NullableRatesType{value: val, isSet: true}
}

func (v NullableRatesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


