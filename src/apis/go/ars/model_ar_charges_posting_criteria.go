/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
)

// checks if the ArChargesPostingCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArChargesPostingCriteria{}

// ArChargesPostingCriteria Request to post charges to an existing invoice, when the functionality is available. If the invoice should not be modified or if the folio should not be modified,based on other functionalities, then the charges cannot be posted to the invoice.
type ArChargesPostingCriteria struct {
	Criteria *ARChargesPostingCriteriaType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewArChargesPostingCriteria instantiates a new ArChargesPostingCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArChargesPostingCriteria() *ArChargesPostingCriteria {
	this := ArChargesPostingCriteria{}
	return &this
}

// NewArChargesPostingCriteriaWithDefaults instantiates a new ArChargesPostingCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArChargesPostingCriteriaWithDefaults() *ArChargesPostingCriteria {
	this := ArChargesPostingCriteria{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *ArChargesPostingCriteria) GetCriteria() ARChargesPostingCriteriaType {
	if o == nil || IsNil(o.Criteria) {
		var ret ARChargesPostingCriteriaType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArChargesPostingCriteria) GetCriteriaOk() (*ARChargesPostingCriteriaType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ArChargesPostingCriteria) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ARChargesPostingCriteriaType and assigns it to the Criteria field.
func (o *ArChargesPostingCriteria) SetCriteria(v ARChargesPostingCriteriaType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ArChargesPostingCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArChargesPostingCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ArChargesPostingCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ArChargesPostingCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ArChargesPostingCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArChargesPostingCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ArChargesPostingCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ArChargesPostingCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ArChargesPostingCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArChargesPostingCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableArChargesPostingCriteria struct {
	value *ArChargesPostingCriteria
	isSet bool
}

func (v NullableArChargesPostingCriteria) Get() *ArChargesPostingCriteria {
	return v.value
}

func (v *NullableArChargesPostingCriteria) Set(val *ArChargesPostingCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableArChargesPostingCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableArChargesPostingCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArChargesPostingCriteria(val *ArChargesPostingCriteria) *NullableArChargesPostingCriteria {
	return &NullableArChargesPostingCriteria{value: val, isSet: true}
}

func (v NullableArChargesPostingCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArChargesPostingCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


