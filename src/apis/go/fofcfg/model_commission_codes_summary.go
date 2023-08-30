/*
OPERA Cloud Front Desk Configuration API

APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fofcfg

import (
	"encoding/json"
)

// checks if the CommissionCodesSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommissionCodesSummary{}

// CommissionCodesSummary Response object for fetching commission codes.
type CommissionCodesSummary struct {
	// Commission code details.
	CommissionCodesSummary []CommissionCodeSummaryInfoType `json:"commissionCodesSummary,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCommissionCodesSummary instantiates a new CommissionCodesSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommissionCodesSummary() *CommissionCodesSummary {
	this := CommissionCodesSummary{}
	return &this
}

// NewCommissionCodesSummaryWithDefaults instantiates a new CommissionCodesSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommissionCodesSummaryWithDefaults() *CommissionCodesSummary {
	this := CommissionCodesSummary{}
	return &this
}

// GetCommissionCodesSummary returns the CommissionCodesSummary field value if set, zero value otherwise.
func (o *CommissionCodesSummary) GetCommissionCodesSummary() []CommissionCodeSummaryInfoType {
	if o == nil || IsNil(o.CommissionCodesSummary) {
		var ret []CommissionCodeSummaryInfoType
		return ret
	}
	return o.CommissionCodesSummary
}

// GetCommissionCodesSummaryOk returns a tuple with the CommissionCodesSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodesSummary) GetCommissionCodesSummaryOk() ([]CommissionCodeSummaryInfoType, bool) {
	if o == nil || IsNil(o.CommissionCodesSummary) {
		return nil, false
	}
	return o.CommissionCodesSummary, true
}

// HasCommissionCodesSummary returns a boolean if a field has been set.
func (o *CommissionCodesSummary) HasCommissionCodesSummary() bool {
	if o != nil && !IsNil(o.CommissionCodesSummary) {
		return true
	}

	return false
}

// SetCommissionCodesSummary gets a reference to the given []CommissionCodeSummaryInfoType and assigns it to the CommissionCodesSummary field.
func (o *CommissionCodesSummary) SetCommissionCodesSummary(v []CommissionCodeSummaryInfoType) {
	o.CommissionCodesSummary = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CommissionCodesSummary) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodesSummary) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CommissionCodesSummary) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CommissionCodesSummary) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CommissionCodesSummary) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodesSummary) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CommissionCodesSummary) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CommissionCodesSummary) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CommissionCodesSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommissionCodesSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommissionCodesSummary) {
		toSerialize["commissionCodesSummary"] = o.CommissionCodesSummary
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCommissionCodesSummary struct {
	value *CommissionCodesSummary
	isSet bool
}

func (v NullableCommissionCodesSummary) Get() *CommissionCodesSummary {
	return v.value
}

func (v *NullableCommissionCodesSummary) Set(val *CommissionCodesSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCommissionCodesSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCommissionCodesSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommissionCodesSummary(val *CommissionCodesSummary) *NullableCommissionCodesSummary {
	return &NullableCommissionCodesSummary{value: val, isSet: true}
}

func (v NullableCommissionCodesSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommissionCodesSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


