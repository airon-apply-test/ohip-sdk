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

// checks if the ARCompressInvoicesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARCompressInvoicesType{}

// ARCompressInvoicesType Criteria type compressing invoices i.e grouping multiple invoices into one,for an Account.
type ARCompressInvoicesType struct {
	Account *ARAccountCriteriaType `json:"account,omitempty"`
	// A collection of AR Invoices.
	Invoices []ARInvoiceType `json:"invoices,omitempty"`
	// Reference Text for the Master Invoice which will be created.
	Reference *string `json:"reference,omitempty"`
	// Remarks for the Master Invoice which will be created.
	Remark *string `json:"remark,omitempty"`
	// The Cashier ID of the Cashier who is currently processing the transaction(s).
	CashierId *float32 `json:"cashierId,omitempty"`
	// Folio Type for the Master Invoice which will be created.
	FolioTypeName *string `json:"folioTypeName,omitempty"`
}

// NewARCompressInvoicesType instantiates a new ARCompressInvoicesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARCompressInvoicesType() *ARCompressInvoicesType {
	this := ARCompressInvoicesType{}
	return &this
}

// NewARCompressInvoicesTypeWithDefaults instantiates a new ARCompressInvoicesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARCompressInvoicesTypeWithDefaults() *ARCompressInvoicesType {
	this := ARCompressInvoicesType{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ARCompressInvoicesType) GetAccount() ARAccountCriteriaType {
	if o == nil || IsNil(o.Account) {
		var ret ARAccountCriteriaType
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCompressInvoicesType) GetAccountOk() (*ARAccountCriteriaType, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ARCompressInvoicesType) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ARAccountCriteriaType and assigns it to the Account field.
func (o *ARCompressInvoicesType) SetAccount(v ARAccountCriteriaType) {
	o.Account = &v
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *ARCompressInvoicesType) GetInvoices() []ARInvoiceType {
	if o == nil || IsNil(o.Invoices) {
		var ret []ARInvoiceType
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCompressInvoicesType) GetInvoicesOk() ([]ARInvoiceType, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *ARCompressInvoicesType) HasInvoices() bool {
	if o != nil && !IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []ARInvoiceType and assigns it to the Invoices field.
func (o *ARCompressInvoicesType) SetInvoices(v []ARInvoiceType) {
	o.Invoices = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ARCompressInvoicesType) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCompressInvoicesType) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ARCompressInvoicesType) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ARCompressInvoicesType) SetReference(v string) {
	o.Reference = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *ARCompressInvoicesType) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCompressInvoicesType) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *ARCompressInvoicesType) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *ARCompressInvoicesType) SetRemark(v string) {
	o.Remark = &v
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *ARCompressInvoicesType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCompressInvoicesType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *ARCompressInvoicesType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *ARCompressInvoicesType) SetCashierId(v float32) {
	o.CashierId = &v
}

// GetFolioTypeName returns the FolioTypeName field value if set, zero value otherwise.
func (o *ARCompressInvoicesType) GetFolioTypeName() string {
	if o == nil || IsNil(o.FolioTypeName) {
		var ret string
		return ret
	}
	return *o.FolioTypeName
}

// GetFolioTypeNameOk returns a tuple with the FolioTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCompressInvoicesType) GetFolioTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.FolioTypeName) {
		return nil, false
	}
	return o.FolioTypeName, true
}

// HasFolioTypeName returns a boolean if a field has been set.
func (o *ARCompressInvoicesType) HasFolioTypeName() bool {
	if o != nil && !IsNil(o.FolioTypeName) {
		return true
	}

	return false
}

// SetFolioTypeName gets a reference to the given string and assigns it to the FolioTypeName field.
func (o *ARCompressInvoicesType) SetFolioTypeName(v string) {
	o.FolioTypeName = &v
}

func (o ARCompressInvoicesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARCompressInvoicesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Invoices) {
		toSerialize["invoices"] = o.Invoices
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.CashierId) {
		toSerialize["cashierId"] = o.CashierId
	}
	if !IsNil(o.FolioTypeName) {
		toSerialize["folioTypeName"] = o.FolioTypeName
	}
	return toSerialize, nil
}

type NullableARCompressInvoicesType struct {
	value *ARCompressInvoicesType
	isSet bool
}

func (v NullableARCompressInvoicesType) Get() *ARCompressInvoicesType {
	return v.value
}

func (v *NullableARCompressInvoicesType) Set(val *ARCompressInvoicesType) {
	v.value = val
	v.isSet = true
}

func (v NullableARCompressInvoicesType) IsSet() bool {
	return v.isSet
}

func (v *NullableARCompressInvoicesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARCompressInvoicesType(val *ARCompressInvoicesType) *NullableARCompressInvoicesType {
	return &NullableARCompressInvoicesType{value: val, isSet: true}
}

func (v NullableARCompressInvoicesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARCompressInvoicesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


