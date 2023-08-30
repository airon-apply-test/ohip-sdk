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

// checks if the Statements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Statements{}

// Statements Details of the Statement to generate.
type Statements struct {
	ARStatements []ARStatementType `json:"aRStatements,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewStatements instantiates a new Statements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatements() *Statements {
	this := Statements{}
	return &this
}

// NewStatementsWithDefaults instantiates a new Statements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementsWithDefaults() *Statements {
	this := Statements{}
	return &this
}

// GetARStatements returns the ARStatements field value if set, zero value otherwise.
func (o *Statements) GetARStatements() []ARStatementType {
	if o == nil || IsNil(o.ARStatements) {
		var ret []ARStatementType
		return ret
	}
	return o.ARStatements
}

// GetARStatementsOk returns a tuple with the ARStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Statements) GetARStatementsOk() ([]ARStatementType, bool) {
	if o == nil || IsNil(o.ARStatements) {
		return nil, false
	}
	return o.ARStatements, true
}

// HasARStatements returns a boolean if a field has been set.
func (o *Statements) HasARStatements() bool {
	if o != nil && !IsNil(o.ARStatements) {
		return true
	}

	return false
}

// SetARStatements gets a reference to the given []ARStatementType and assigns it to the ARStatements field.
func (o *Statements) SetARStatements(v []ARStatementType) {
	o.ARStatements = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Statements) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Statements) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Statements) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *Statements) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Statements) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Statements) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Statements) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *Statements) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o Statements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Statements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ARStatements) {
		toSerialize["aRStatements"] = o.ARStatements
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableStatements struct {
	value *Statements
	isSet bool
}

func (v NullableStatements) Get() *Statements {
	return v.value
}

func (v *NullableStatements) Set(val *Statements) {
	v.value = val
	v.isSet = true
}

func (v NullableStatements) IsSet() bool {
	return v.isSet
}

func (v *NullableStatements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatements(val *Statements) *NullableStatements {
	return &NullableStatements{value: val, isSet: true}
}

func (v NullableStatements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


