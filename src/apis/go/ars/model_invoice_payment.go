/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InvoicePayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoicePayment{}

// InvoicePayment struct for InvoicePayment
type InvoicePayment struct {
	// Account Invoice information.
	Details []ARAccountInvoicesPaymentsType `json:"details,omitempty"`
	// List of Transaction codes info.
	TrxCodesInfo []TrxInfoType `json:"trxCodesInfo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewInvoicePayment instantiates a new InvoicePayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicePayment() *InvoicePayment {
	this := InvoicePayment{}
	return &this
}

// NewInvoicePaymentWithDefaults instantiates a new InvoicePayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicePaymentWithDefaults() *InvoicePayment {
	this := InvoicePayment{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InvoicePayment) GetDetails() []ARAccountInvoicesPaymentsType {
	if o == nil || IsNil(o.Details) {
		var ret []ARAccountInvoicesPaymentsType
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicePayment) GetDetailsOk() ([]ARAccountInvoicesPaymentsType, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InvoicePayment) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ARAccountInvoicesPaymentsType and assigns it to the Details field.
func (o *InvoicePayment) SetDetails(v []ARAccountInvoicesPaymentsType) {
	o.Details = v
}

// GetTrxCodesInfo returns the TrxCodesInfo field value if set, zero value otherwise.
func (o *InvoicePayment) GetTrxCodesInfo() []TrxInfoType {
	if o == nil || IsNil(o.TrxCodesInfo) {
		var ret []TrxInfoType
		return ret
	}
	return o.TrxCodesInfo
}

// GetTrxCodesInfoOk returns a tuple with the TrxCodesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicePayment) GetTrxCodesInfoOk() ([]TrxInfoType, bool) {
	if o == nil || IsNil(o.TrxCodesInfo) {
		return nil, false
	}
	return o.TrxCodesInfo, true
}

// HasTrxCodesInfo returns a boolean if a field has been set.
func (o *InvoicePayment) HasTrxCodesInfo() bool {
	if o != nil && !IsNil(o.TrxCodesInfo) {
		return true
	}

	return false
}

// SetTrxCodesInfo gets a reference to the given []TrxInfoType and assigns it to the TrxCodesInfo field.
func (o *InvoicePayment) SetTrxCodesInfo(v []TrxInfoType) {
	o.TrxCodesInfo = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InvoicePayment) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicePayment) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InvoicePayment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *InvoicePayment) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *InvoicePayment) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicePayment) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *InvoicePayment) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *InvoicePayment) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o InvoicePayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoicePayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.TrxCodesInfo) {
		toSerialize["trxCodesInfo"] = o.TrxCodesInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableInvoicePayment struct {
	value *InvoicePayment
	isSet bool
}

func (v NullableInvoicePayment) Get() *InvoicePayment {
	return v.value
}

func (v *NullableInvoicePayment) Set(val *InvoicePayment) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicePayment) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicePayment(val *InvoicePayment) *NullableInvoicePayment {
	return &NullableInvoicePayment{value: val, isSet: true}
}

func (v NullableInvoicePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


