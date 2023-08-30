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

// checks if the ArCreditCardPaymentsToTransfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArCreditCardPaymentsToTransfer{}

// ArCreditCardPaymentsToTransfer Request to transfer AR Credit Card payments.
type ArCreditCardPaymentsToTransfer struct {
	Criteria *TransferARCreditCardPaymentsType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewArCreditCardPaymentsToTransfer instantiates a new ArCreditCardPaymentsToTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArCreditCardPaymentsToTransfer() *ArCreditCardPaymentsToTransfer {
	this := ArCreditCardPaymentsToTransfer{}
	return &this
}

// NewArCreditCardPaymentsToTransferWithDefaults instantiates a new ArCreditCardPaymentsToTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArCreditCardPaymentsToTransferWithDefaults() *ArCreditCardPaymentsToTransfer {
	this := ArCreditCardPaymentsToTransfer{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *ArCreditCardPaymentsToTransfer) GetCriteria() TransferARCreditCardPaymentsType {
	if o == nil || IsNil(o.Criteria) {
		var ret TransferARCreditCardPaymentsType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArCreditCardPaymentsToTransfer) GetCriteriaOk() (*TransferARCreditCardPaymentsType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ArCreditCardPaymentsToTransfer) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given TransferARCreditCardPaymentsType and assigns it to the Criteria field.
func (o *ArCreditCardPaymentsToTransfer) SetCriteria(v TransferARCreditCardPaymentsType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ArCreditCardPaymentsToTransfer) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArCreditCardPaymentsToTransfer) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ArCreditCardPaymentsToTransfer) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ArCreditCardPaymentsToTransfer) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ArCreditCardPaymentsToTransfer) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArCreditCardPaymentsToTransfer) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ArCreditCardPaymentsToTransfer) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ArCreditCardPaymentsToTransfer) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ArCreditCardPaymentsToTransfer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArCreditCardPaymentsToTransfer) ToMap() (map[string]interface{}, error) {
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

type NullableArCreditCardPaymentsToTransfer struct {
	value *ArCreditCardPaymentsToTransfer
	isSet bool
}

func (v NullableArCreditCardPaymentsToTransfer) Get() *ArCreditCardPaymentsToTransfer {
	return v.value
}

func (v *NullableArCreditCardPaymentsToTransfer) Set(val *ArCreditCardPaymentsToTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableArCreditCardPaymentsToTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableArCreditCardPaymentsToTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArCreditCardPaymentsToTransfer(val *ArCreditCardPaymentsToTransfer) *NullableArCreditCardPaymentsToTransfer {
	return &NullableArCreditCardPaymentsToTransfer{value: val, isSet: true}
}

func (v NullableArCreditCardPaymentsToTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArCreditCardPaymentsToTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


