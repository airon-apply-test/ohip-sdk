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

// checks if the StatementsToBeGenerated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatementsToBeGenerated{}

// StatementsToBeGenerated Operation to generate AR Statements. This will validate and check if there exists any invoices to be included in the Statement based on the criteria used. If Statement Numbering is used, this will return a statement number to use in report as well as report sequence id to identify statement's invoices.
type StatementsToBeGenerated struct {
	Criteria *ARGenerateStatementCriteriaType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewStatementsToBeGenerated instantiates a new StatementsToBeGenerated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementsToBeGenerated() *StatementsToBeGenerated {
	this := StatementsToBeGenerated{}
	return &this
}

// NewStatementsToBeGeneratedWithDefaults instantiates a new StatementsToBeGenerated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementsToBeGeneratedWithDefaults() *StatementsToBeGenerated {
	this := StatementsToBeGenerated{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *StatementsToBeGenerated) GetCriteria() ARGenerateStatementCriteriaType {
	if o == nil || IsNil(o.Criteria) {
		var ret ARGenerateStatementCriteriaType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementsToBeGenerated) GetCriteriaOk() (*ARGenerateStatementCriteriaType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *StatementsToBeGenerated) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ARGenerateStatementCriteriaType and assigns it to the Criteria field.
func (o *StatementsToBeGenerated) SetCriteria(v ARGenerateStatementCriteriaType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *StatementsToBeGenerated) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementsToBeGenerated) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StatementsToBeGenerated) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *StatementsToBeGenerated) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *StatementsToBeGenerated) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementsToBeGenerated) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *StatementsToBeGenerated) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *StatementsToBeGenerated) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o StatementsToBeGenerated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatementsToBeGenerated) ToMap() (map[string]interface{}, error) {
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

type NullableStatementsToBeGenerated struct {
	value *StatementsToBeGenerated
	isSet bool
}

func (v NullableStatementsToBeGenerated) Get() *StatementsToBeGenerated {
	return v.value
}

func (v *NullableStatementsToBeGenerated) Set(val *StatementsToBeGenerated) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementsToBeGenerated) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementsToBeGenerated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementsToBeGenerated(val *StatementsToBeGenerated) *NullableStatementsToBeGenerated {
	return &NullableStatementsToBeGenerated{value: val, isSet: true}
}

func (v NullableStatementsToBeGenerated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementsToBeGenerated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


