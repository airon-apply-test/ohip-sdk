/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FlexibleBenefitPostingsCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlexibleBenefitPostingsCriteria{}

// FlexibleBenefitPostingsCriteria Apply Flexible Benefit Postings.
type FlexibleBenefitPostingsCriteria struct {
	ApplyFlexibleBenefitCriteria *ApplyFlexibleBenefitPostingsCriteriaType `json:"applyFlexibleBenefitCriteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewFlexibleBenefitPostingsCriteria instantiates a new FlexibleBenefitPostingsCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexibleBenefitPostingsCriteria() *FlexibleBenefitPostingsCriteria {
	this := FlexibleBenefitPostingsCriteria{}
	return &this
}

// NewFlexibleBenefitPostingsCriteriaWithDefaults instantiates a new FlexibleBenefitPostingsCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexibleBenefitPostingsCriteriaWithDefaults() *FlexibleBenefitPostingsCriteria {
	this := FlexibleBenefitPostingsCriteria{}
	return &this
}

// GetApplyFlexibleBenefitCriteria returns the ApplyFlexibleBenefitCriteria field value if set, zero value otherwise.
func (o *FlexibleBenefitPostingsCriteria) GetApplyFlexibleBenefitCriteria() ApplyFlexibleBenefitPostingsCriteriaType {
	if o == nil || IsNil(o.ApplyFlexibleBenefitCriteria) {
		var ret ApplyFlexibleBenefitPostingsCriteriaType
		return ret
	}
	return *o.ApplyFlexibleBenefitCriteria
}

// GetApplyFlexibleBenefitCriteriaOk returns a tuple with the ApplyFlexibleBenefitCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleBenefitPostingsCriteria) GetApplyFlexibleBenefitCriteriaOk() (*ApplyFlexibleBenefitPostingsCriteriaType, bool) {
	if o == nil || IsNil(o.ApplyFlexibleBenefitCriteria) {
		return nil, false
	}
	return o.ApplyFlexibleBenefitCriteria, true
}

// HasApplyFlexibleBenefitCriteria returns a boolean if a field has been set.
func (o *FlexibleBenefitPostingsCriteria) HasApplyFlexibleBenefitCriteria() bool {
	if o != nil && !IsNil(o.ApplyFlexibleBenefitCriteria) {
		return true
	}

	return false
}

// SetApplyFlexibleBenefitCriteria gets a reference to the given ApplyFlexibleBenefitPostingsCriteriaType and assigns it to the ApplyFlexibleBenefitCriteria field.
func (o *FlexibleBenefitPostingsCriteria) SetApplyFlexibleBenefitCriteria(v ApplyFlexibleBenefitPostingsCriteriaType) {
	o.ApplyFlexibleBenefitCriteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FlexibleBenefitPostingsCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleBenefitPostingsCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FlexibleBenefitPostingsCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *FlexibleBenefitPostingsCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FlexibleBenefitPostingsCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleBenefitPostingsCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FlexibleBenefitPostingsCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *FlexibleBenefitPostingsCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o FlexibleBenefitPostingsCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlexibleBenefitPostingsCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyFlexibleBenefitCriteria) {
		toSerialize["applyFlexibleBenefitCriteria"] = o.ApplyFlexibleBenefitCriteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableFlexibleBenefitPostingsCriteria struct {
	value *FlexibleBenefitPostingsCriteria
	isSet bool
}

func (v NullableFlexibleBenefitPostingsCriteria) Get() *FlexibleBenefitPostingsCriteria {
	return v.value
}

func (v *NullableFlexibleBenefitPostingsCriteria) Set(val *FlexibleBenefitPostingsCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleBenefitPostingsCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleBenefitPostingsCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleBenefitPostingsCriteria(val *FlexibleBenefitPostingsCriteria) *NullableFlexibleBenefitPostingsCriteria {
	return &NullableFlexibleBenefitPostingsCriteria{value: val, isSet: true}
}

func (v NullableFlexibleBenefitPostingsCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleBenefitPostingsCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


