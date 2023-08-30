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

// checks if the AROldBalanceChargeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AROldBalanceChargeType{}

// AROldBalanceChargeType AR Old Balances Single Posting Type.
type AROldBalanceChargeType struct {
	// User-defined posting reference.
	PostingReference *string `json:"postingReference,omitempty"`
	// User-defined Supplement.
	PostingRemark *string `json:"postingRemark,omitempty"`
	// Date of the Posting.
	Date *string `json:"date,omitempty"`
	// The Folio number of this posting, if there was a Folio entered.
	FolioNo *float32 `json:"folioNo,omitempty"`
	// The Fiscal Bill number of this posting
	FiscalBillNo *string `json:"fiscalBillNo,omitempty"`
	Amount *CurrencyAmountType `json:"amount,omitempty"`
	// Values of atmost 20 Taxes entered.
	TaxCodes []ARTaxCodeType `json:"taxCodes,omitempty"`
	Paid *CurrencyAmountType `json:"paid,omitempty"`
}

// NewAROldBalanceChargeType instantiates a new AROldBalanceChargeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAROldBalanceChargeType() *AROldBalanceChargeType {
	this := AROldBalanceChargeType{}
	return &this
}

// NewAROldBalanceChargeTypeWithDefaults instantiates a new AROldBalanceChargeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAROldBalanceChargeTypeWithDefaults() *AROldBalanceChargeType {
	this := AROldBalanceChargeType{}
	return &this
}

// GetPostingReference returns the PostingReference field value if set, zero value otherwise.
func (o *AROldBalanceChargeType) GetPostingReference() string {
	if o == nil || IsNil(o.PostingReference) {
		var ret string
		return ret
	}
	return *o.PostingReference
}

// GetPostingReferenceOk returns a tuple with the PostingReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AROldBalanceChargeType) GetPostingReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.PostingReference) {
		return nil, false
	}
	return o.PostingReference, true
}

// HasPostingReference returns a boolean if a field has been set.
func (o *AROldBalanceChargeType) HasPostingReference() bool {
	if o != nil && !IsNil(o.PostingReference) {
		return true
	}

	return false
}

// SetPostingReference gets a reference to the given string and assigns it to the PostingReference field.
func (o *AROldBalanceChargeType) SetPostingReference(v string) {
	o.PostingReference = &v
}

// GetPostingRemark returns the PostingRemark field value if set, zero value otherwise.
func (o *AROldBalanceChargeType) GetPostingRemark() string {
	if o == nil || IsNil(o.PostingRemark) {
		var ret string
		return ret
	}
	return *o.PostingRemark
}

// GetPostingRemarkOk returns a tuple with the PostingRemark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AROldBalanceChargeType) GetPostingRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.PostingRemark) {
		return nil, false
	}
	return o.PostingRemark, true
}

// HasPostingRemark returns a boolean if a field has been set.
func (o *AROldBalanceChargeType) HasPostingRemark() bool {
	if o != nil && !IsNil(o.PostingRemark) {
		return true
	}

	return false
}

// SetPostingRemark gets a reference to the given string and assigns it to the PostingRemark field.
func (o *AROldBalanceChargeType) SetPostingRemark(v string) {
	o.PostingRemark = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AROldBalanceChargeType) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AROldBalanceChargeType) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AROldBalanceChargeType) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AROldBalanceChargeType) SetDate(v string) {
	o.Date = &v
}

// GetFolioNo returns the FolioNo field value if set, zero value otherwise.
func (o *AROldBalanceChargeType) GetFolioNo() float32 {
	if o == nil || IsNil(o.FolioNo) {
		var ret float32
		return ret
	}
	return *o.FolioNo
}

// GetFolioNoOk returns a tuple with the FolioNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AROldBalanceChargeType) GetFolioNoOk() (*float32, bool) {
	if o == nil || IsNil(o.FolioNo) {
		return nil, false
	}
	return o.FolioNo, true
}

// HasFolioNo returns a boolean if a field has been set.
func (o *AROldBalanceChargeType) HasFolioNo() bool {
	if o != nil && !IsNil(o.FolioNo) {
		return true
	}

	return false
}

// SetFolioNo gets a reference to the given float32 and assigns it to the FolioNo field.
func (o *AROldBalanceChargeType) SetFolioNo(v float32) {
	o.FolioNo = &v
}

// GetFiscalBillNo returns the FiscalBillNo field value if set, zero value otherwise.
func (o *AROldBalanceChargeType) GetFiscalBillNo() string {
	if o == nil || IsNil(o.FiscalBillNo) {
		var ret string
		return ret
	}
	return *o.FiscalBillNo
}

// GetFiscalBillNoOk returns a tuple with the FiscalBillNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AROldBalanceChargeType) GetFiscalBillNoOk() (*string, bool) {
	if o == nil || IsNil(o.FiscalBillNo) {
		return nil, false
	}
	return o.FiscalBillNo, true
}

// HasFiscalBillNo returns a boolean if a field has been set.
func (o *AROldBalanceChargeType) HasFiscalBillNo() bool {
	if o != nil && !IsNil(o.FiscalBillNo) {
		return true
	}

	return false
}

// SetFiscalBillNo gets a reference to the given string and assigns it to the FiscalBillNo field.
func (o *AROldBalanceChargeType) SetFiscalBillNo(v string) {
	o.FiscalBillNo = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AROldBalanceChargeType) GetAmount() CurrencyAmountType {
	if o == nil || IsNil(o.Amount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AROldBalanceChargeType) GetAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AROldBalanceChargeType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given CurrencyAmountType and assigns it to the Amount field.
func (o *AROldBalanceChargeType) SetAmount(v CurrencyAmountType) {
	o.Amount = &v
}

// GetTaxCodes returns the TaxCodes field value if set, zero value otherwise.
func (o *AROldBalanceChargeType) GetTaxCodes() []ARTaxCodeType {
	if o == nil || IsNil(o.TaxCodes) {
		var ret []ARTaxCodeType
		return ret
	}
	return o.TaxCodes
}

// GetTaxCodesOk returns a tuple with the TaxCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AROldBalanceChargeType) GetTaxCodesOk() ([]ARTaxCodeType, bool) {
	if o == nil || IsNil(o.TaxCodes) {
		return nil, false
	}
	return o.TaxCodes, true
}

// HasTaxCodes returns a boolean if a field has been set.
func (o *AROldBalanceChargeType) HasTaxCodes() bool {
	if o != nil && !IsNil(o.TaxCodes) {
		return true
	}

	return false
}

// SetTaxCodes gets a reference to the given []ARTaxCodeType and assigns it to the TaxCodes field.
func (o *AROldBalanceChargeType) SetTaxCodes(v []ARTaxCodeType) {
	o.TaxCodes = v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *AROldBalanceChargeType) GetPaid() CurrencyAmountType {
	if o == nil || IsNil(o.Paid) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AROldBalanceChargeType) GetPaidOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *AROldBalanceChargeType) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given CurrencyAmountType and assigns it to the Paid field.
func (o *AROldBalanceChargeType) SetPaid(v CurrencyAmountType) {
	o.Paid = &v
}

func (o AROldBalanceChargeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AROldBalanceChargeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostingReference) {
		toSerialize["postingReference"] = o.PostingReference
	}
	if !IsNil(o.PostingRemark) {
		toSerialize["postingRemark"] = o.PostingRemark
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.FolioNo) {
		toSerialize["folioNo"] = o.FolioNo
	}
	if !IsNil(o.FiscalBillNo) {
		toSerialize["fiscalBillNo"] = o.FiscalBillNo
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.TaxCodes) {
		toSerialize["taxCodes"] = o.TaxCodes
	}
	if !IsNil(o.Paid) {
		toSerialize["paid"] = o.Paid
	}
	return toSerialize, nil
}

type NullableAROldBalanceChargeType struct {
	value *AROldBalanceChargeType
	isSet bool
}

func (v NullableAROldBalanceChargeType) Get() *AROldBalanceChargeType {
	return v.value
}

func (v *NullableAROldBalanceChargeType) Set(val *AROldBalanceChargeType) {
	v.value = val
	v.isSet = true
}

func (v NullableAROldBalanceChargeType) IsSet() bool {
	return v.isSet
}

func (v *NullableAROldBalanceChargeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAROldBalanceChargeType(val *AROldBalanceChargeType) *NullableAROldBalanceChargeType {
	return &NullableAROldBalanceChargeType{value: val, isSet: true}
}

func (v NullableAROldBalanceChargeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAROldBalanceChargeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


