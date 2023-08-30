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

// checks if the TransactionARInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionARInfoType{}

// TransactionARInfoType Details of the AR account where this transaction was posted to.
type TransactionARInfoType struct {
	// AR Account number.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The Account name where this invoice is sent.
	AccountName *string `json:"accountName,omitempty"`
	// The invoice number which includes this transaction.
	InvoiceNo *float32 `json:"invoiceNo,omitempty"`
}

// NewTransactionARInfoType instantiates a new TransactionARInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionARInfoType() *TransactionARInfoType {
	this := TransactionARInfoType{}
	return &this
}

// NewTransactionARInfoTypeWithDefaults instantiates a new TransactionARInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionARInfoTypeWithDefaults() *TransactionARInfoType {
	this := TransactionARInfoType{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *TransactionARInfoType) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionARInfoType) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *TransactionARInfoType) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *TransactionARInfoType) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *TransactionARInfoType) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionARInfoType) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *TransactionARInfoType) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *TransactionARInfoType) SetAccountName(v string) {
	o.AccountName = &v
}

// GetInvoiceNo returns the InvoiceNo field value if set, zero value otherwise.
func (o *TransactionARInfoType) GetInvoiceNo() float32 {
	if o == nil || IsNil(o.InvoiceNo) {
		var ret float32
		return ret
	}
	return *o.InvoiceNo
}

// GetInvoiceNoOk returns a tuple with the InvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionARInfoType) GetInvoiceNoOk() (*float32, bool) {
	if o == nil || IsNil(o.InvoiceNo) {
		return nil, false
	}
	return o.InvoiceNo, true
}

// HasInvoiceNo returns a boolean if a field has been set.
func (o *TransactionARInfoType) HasInvoiceNo() bool {
	if o != nil && !IsNil(o.InvoiceNo) {
		return true
	}

	return false
}

// SetInvoiceNo gets a reference to the given float32 and assigns it to the InvoiceNo field.
func (o *TransactionARInfoType) SetInvoiceNo(v float32) {
	o.InvoiceNo = &v
}

func (o TransactionARInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionARInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.InvoiceNo) {
		toSerialize["invoiceNo"] = o.InvoiceNo
	}
	return toSerialize, nil
}

type NullableTransactionARInfoType struct {
	value *TransactionARInfoType
	isSet bool
}

func (v NullableTransactionARInfoType) Get() *TransactionARInfoType {
	return v.value
}

func (v *NullableTransactionARInfoType) Set(val *TransactionARInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionARInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionARInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionARInfoType(val *TransactionARInfoType) *NullableTransactionARInfoType {
	return &NullableTransactionARInfoType{value: val, isSet: true}
}

func (v NullableTransactionARInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionARInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


