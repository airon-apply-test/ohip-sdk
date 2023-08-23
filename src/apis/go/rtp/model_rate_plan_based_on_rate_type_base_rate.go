/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RatePlanBasedOnRateTypeBaseRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatePlanBasedOnRateTypeBaseRate{}

// RatePlanBasedOnRateTypeBaseRate Base Rate type
type RatePlanBasedOnRateTypeBaseRate struct {
	// Rate plan code used to base the rate on.
	BasedOnRatePlan *string `json:"basedOnRatePlan,omitempty"`
	// Base Amount used for base rate calculation.
	BaseAmount *float32 `json:"baseAmount,omitempty"`
	// Flat or Percentage (FLT/PCT) indicator.
	FlatOrPercentage *string `json:"flatOrPercentage,omitempty"`
	// Rounding style used for the calculated rate amounts. Valid values are U,D,N,C,F which means Up, Down, None, Up-Keep Decimal, Down - Keep Decimal.
	Rounding *string `json:"rounding,omitempty"`
	// Rate Plan code.
	DependentRatePlans []string `json:"dependentRatePlans,omitempty"`
}

// NewRatePlanBasedOnRateTypeBaseRate instantiates a new RatePlanBasedOnRateTypeBaseRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatePlanBasedOnRateTypeBaseRate() *RatePlanBasedOnRateTypeBaseRate {
	this := RatePlanBasedOnRateTypeBaseRate{}
	return &this
}

// NewRatePlanBasedOnRateTypeBaseRateWithDefaults instantiates a new RatePlanBasedOnRateTypeBaseRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatePlanBasedOnRateTypeBaseRateWithDefaults() *RatePlanBasedOnRateTypeBaseRate {
	this := RatePlanBasedOnRateTypeBaseRate{}
	return &this
}

// GetBasedOnRatePlan returns the BasedOnRatePlan field value if set, zero value otherwise.
func (o *RatePlanBasedOnRateTypeBaseRate) GetBasedOnRatePlan() string {
	if o == nil || IsNil(o.BasedOnRatePlan) {
		var ret string
		return ret
	}
	return *o.BasedOnRatePlan
}

// GetBasedOnRatePlanOk returns a tuple with the BasedOnRatePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) GetBasedOnRatePlanOk() (*string, bool) {
	if o == nil || IsNil(o.BasedOnRatePlan) {
		return nil, false
	}
	return o.BasedOnRatePlan, true
}

// HasBasedOnRatePlan returns a boolean if a field has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) HasBasedOnRatePlan() bool {
	if o != nil && !IsNil(o.BasedOnRatePlan) {
		return true
	}

	return false
}

// SetBasedOnRatePlan gets a reference to the given string and assigns it to the BasedOnRatePlan field.
func (o *RatePlanBasedOnRateTypeBaseRate) SetBasedOnRatePlan(v string) {
	o.BasedOnRatePlan = &v
}

// GetBaseAmount returns the BaseAmount field value if set, zero value otherwise.
func (o *RatePlanBasedOnRateTypeBaseRate) GetBaseAmount() float32 {
	if o == nil || IsNil(o.BaseAmount) {
		var ret float32
		return ret
	}
	return *o.BaseAmount
}

// GetBaseAmountOk returns a tuple with the BaseAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) GetBaseAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseAmount) {
		return nil, false
	}
	return o.BaseAmount, true
}

// HasBaseAmount returns a boolean if a field has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) HasBaseAmount() bool {
	if o != nil && !IsNil(o.BaseAmount) {
		return true
	}

	return false
}

// SetBaseAmount gets a reference to the given float32 and assigns it to the BaseAmount field.
func (o *RatePlanBasedOnRateTypeBaseRate) SetBaseAmount(v float32) {
	o.BaseAmount = &v
}

// GetFlatOrPercentage returns the FlatOrPercentage field value if set, zero value otherwise.
func (o *RatePlanBasedOnRateTypeBaseRate) GetFlatOrPercentage() string {
	if o == nil || IsNil(o.FlatOrPercentage) {
		var ret string
		return ret
	}
	return *o.FlatOrPercentage
}

// GetFlatOrPercentageOk returns a tuple with the FlatOrPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) GetFlatOrPercentageOk() (*string, bool) {
	if o == nil || IsNil(o.FlatOrPercentage) {
		return nil, false
	}
	return o.FlatOrPercentage, true
}

// HasFlatOrPercentage returns a boolean if a field has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) HasFlatOrPercentage() bool {
	if o != nil && !IsNil(o.FlatOrPercentage) {
		return true
	}

	return false
}

// SetFlatOrPercentage gets a reference to the given string and assigns it to the FlatOrPercentage field.
func (o *RatePlanBasedOnRateTypeBaseRate) SetFlatOrPercentage(v string) {
	o.FlatOrPercentage = &v
}

// GetRounding returns the Rounding field value if set, zero value otherwise.
func (o *RatePlanBasedOnRateTypeBaseRate) GetRounding() string {
	if o == nil || IsNil(o.Rounding) {
		var ret string
		return ret
	}
	return *o.Rounding
}

// GetRoundingOk returns a tuple with the Rounding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) GetRoundingOk() (*string, bool) {
	if o == nil || IsNil(o.Rounding) {
		return nil, false
	}
	return o.Rounding, true
}

// HasRounding returns a boolean if a field has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) HasRounding() bool {
	if o != nil && !IsNil(o.Rounding) {
		return true
	}

	return false
}

// SetRounding gets a reference to the given string and assigns it to the Rounding field.
func (o *RatePlanBasedOnRateTypeBaseRate) SetRounding(v string) {
	o.Rounding = &v
}

// GetDependentRatePlans returns the DependentRatePlans field value if set, zero value otherwise.
func (o *RatePlanBasedOnRateTypeBaseRate) GetDependentRatePlans() []string {
	if o == nil || IsNil(o.DependentRatePlans) {
		var ret []string
		return ret
	}
	return o.DependentRatePlans
}

// GetDependentRatePlansOk returns a tuple with the DependentRatePlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) GetDependentRatePlansOk() ([]string, bool) {
	if o == nil || IsNil(o.DependentRatePlans) {
		return nil, false
	}
	return o.DependentRatePlans, true
}

// HasDependentRatePlans returns a boolean if a field has been set.
func (o *RatePlanBasedOnRateTypeBaseRate) HasDependentRatePlans() bool {
	if o != nil && !IsNil(o.DependentRatePlans) {
		return true
	}

	return false
}

// SetDependentRatePlans gets a reference to the given []string and assigns it to the DependentRatePlans field.
func (o *RatePlanBasedOnRateTypeBaseRate) SetDependentRatePlans(v []string) {
	o.DependentRatePlans = v
}

func (o RatePlanBasedOnRateTypeBaseRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatePlanBasedOnRateTypeBaseRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasedOnRatePlan) {
		toSerialize["basedOnRatePlan"] = o.BasedOnRatePlan
	}
	if !IsNil(o.BaseAmount) {
		toSerialize["baseAmount"] = o.BaseAmount
	}
	if !IsNil(o.FlatOrPercentage) {
		toSerialize["flatOrPercentage"] = o.FlatOrPercentage
	}
	if !IsNil(o.Rounding) {
		toSerialize["rounding"] = o.Rounding
	}
	if !IsNil(o.DependentRatePlans) {
		toSerialize["dependentRatePlans"] = o.DependentRatePlans
	}
	return toSerialize, nil
}

type NullableRatePlanBasedOnRateTypeBaseRate struct {
	value *RatePlanBasedOnRateTypeBaseRate
	isSet bool
}

func (v NullableRatePlanBasedOnRateTypeBaseRate) Get() *RatePlanBasedOnRateTypeBaseRate {
	return v.value
}

func (v *NullableRatePlanBasedOnRateTypeBaseRate) Set(val *RatePlanBasedOnRateTypeBaseRate) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePlanBasedOnRateTypeBaseRate) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePlanBasedOnRateTypeBaseRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePlanBasedOnRateTypeBaseRate(val *RatePlanBasedOnRateTypeBaseRate) *NullableRatePlanBasedOnRateTypeBaseRate {
	return &NullableRatePlanBasedOnRateTypeBaseRate{value: val, isSet: true}
}

func (v NullableRatePlanBasedOnRateTypeBaseRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePlanBasedOnRateTypeBaseRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


