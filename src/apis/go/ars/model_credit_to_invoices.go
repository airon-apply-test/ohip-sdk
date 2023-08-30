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

// checks if the CreditToInvoices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditToInvoices{}

// CreditToInvoices struct for CreditToInvoices
type CreditToInvoices struct {
	Criteria *ARApplyPaymentCriteriaType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCreditToInvoices instantiates a new CreditToInvoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditToInvoices() *CreditToInvoices {
	this := CreditToInvoices{}
	return &this
}

// NewCreditToInvoicesWithDefaults instantiates a new CreditToInvoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditToInvoicesWithDefaults() *CreditToInvoices {
	this := CreditToInvoices{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *CreditToInvoices) GetCriteria() ARApplyPaymentCriteriaType {
	if o == nil || IsNil(o.Criteria) {
		var ret ARApplyPaymentCriteriaType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditToInvoices) GetCriteriaOk() (*ARApplyPaymentCriteriaType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *CreditToInvoices) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ARApplyPaymentCriteriaType and assigns it to the Criteria field.
func (o *CreditToInvoices) SetCriteria(v ARApplyPaymentCriteriaType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreditToInvoices) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditToInvoices) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreditToInvoices) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CreditToInvoices) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreditToInvoices) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditToInvoices) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreditToInvoices) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CreditToInvoices) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CreditToInvoices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditToInvoices) ToMap() (map[string]interface{}, error) {
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

type NullableCreditToInvoices struct {
	value *CreditToInvoices
	isSet bool
}

func (v NullableCreditToInvoices) Get() *CreditToInvoices {
	return v.value
}

func (v *NullableCreditToInvoices) Set(val *CreditToInvoices) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditToInvoices) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditToInvoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditToInvoices(val *CreditToInvoices) *NullableCreditToInvoices {
	return &NullableCreditToInvoices{value: val, isSet: true}
}

func (v NullableCreditToInvoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditToInvoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


