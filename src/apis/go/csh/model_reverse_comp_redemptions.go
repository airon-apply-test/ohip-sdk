/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the ReverseCompRedemptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReverseCompRedemptions{}

// ReverseCompRedemptions Request type of complimentary bucket redemptions reversal.
type ReverseCompRedemptions struct {
	Criteria *ReverseCompRedemptionsCriteria `json:"criteria,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewReverseCompRedemptions instantiates a new ReverseCompRedemptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseCompRedemptions() *ReverseCompRedemptions {
	this := ReverseCompRedemptions{}
	return &this
}

// NewReverseCompRedemptionsWithDefaults instantiates a new ReverseCompRedemptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseCompRedemptionsWithDefaults() *ReverseCompRedemptions {
	this := ReverseCompRedemptions{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *ReverseCompRedemptions) GetCriteria() ReverseCompRedemptionsCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret ReverseCompRedemptionsCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseCompRedemptions) GetCriteriaOk() (*ReverseCompRedemptionsCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ReverseCompRedemptions) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ReverseCompRedemptionsCriteria and assigns it to the Criteria field.
func (o *ReverseCompRedemptions) SetCriteria(v ReverseCompRedemptionsCriteria) {
	o.Criteria = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ReverseCompRedemptions) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseCompRedemptions) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ReverseCompRedemptions) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ReverseCompRedemptions) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ReverseCompRedemptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReverseCompRedemptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableReverseCompRedemptions struct {
	value *ReverseCompRedemptions
	isSet bool
}

func (v NullableReverseCompRedemptions) Get() *ReverseCompRedemptions {
	return v.value
}

func (v *NullableReverseCompRedemptions) Set(val *ReverseCompRedemptions) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseCompRedemptions) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseCompRedemptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseCompRedemptions(val *ReverseCompRedemptions) *NullableReverseCompRedemptions {
	return &NullableReverseCompRedemptions{value: val, isSet: true}
}

func (v NullableReverseCompRedemptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseCompRedemptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


