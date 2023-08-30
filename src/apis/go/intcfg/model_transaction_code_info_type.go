/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intcfg

import (
	"encoding/json"
)

// checks if the TransactionCodeInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionCodeInfoType{}

// TransactionCodeInfoType struct for TransactionCodeInfoType
type TransactionCodeInfoType struct {
	// All charges code of transaction code setup
	AllChargesCode *string `json:"allChargesCode,omitempty"`
	// Minibar charges code of transaction code setup
	MiniBarCode *string `json:"miniBarCode,omitempty"`
	// Cash or credit card posting of transaction code setup
	CashCreditCardPosting *string `json:"cashCreditCardPosting,omitempty"`
	// Subtotal posting of transaction code setup
	SubtotalPosting *string `json:"subtotalPosting,omitempty"`
	// Difference posting of transaction code setup
	DifferencePosting *string `json:"differencePosting,omitempty"`
	// Calculated transaction code posting of transaction code setup
	CalcTrxCodePosting *string `json:"calcTrxCodePosting,omitempty"`
	// Number of dialed digits transaction code setup
	NumberOfDialedDigits *string `json:"numberOfDialedDigits,omitempty"`
	// Postings information of transaction codes setup.
	PostingAccounts []PostingAccountType `json:"postingAccounts,omitempty"`
	// Split information of transaction code setup.
	Itemizers []ItemizerType `json:"itemizers,omitempty"`
	// Split information of transaction code setup.
	TransactionCodes []TransactionCodeDetailType `json:"transactionCodes,omitempty"`
}

// NewTransactionCodeInfoType instantiates a new TransactionCodeInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCodeInfoType() *TransactionCodeInfoType {
	this := TransactionCodeInfoType{}
	return &this
}

// NewTransactionCodeInfoTypeWithDefaults instantiates a new TransactionCodeInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCodeInfoTypeWithDefaults() *TransactionCodeInfoType {
	this := TransactionCodeInfoType{}
	return &this
}

// GetAllChargesCode returns the AllChargesCode field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetAllChargesCode() string {
	if o == nil || IsNil(o.AllChargesCode) {
		var ret string
		return ret
	}
	return *o.AllChargesCode
}

// GetAllChargesCodeOk returns a tuple with the AllChargesCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetAllChargesCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AllChargesCode) {
		return nil, false
	}
	return o.AllChargesCode, true
}

// HasAllChargesCode returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasAllChargesCode() bool {
	if o != nil && !IsNil(o.AllChargesCode) {
		return true
	}

	return false
}

// SetAllChargesCode gets a reference to the given string and assigns it to the AllChargesCode field.
func (o *TransactionCodeInfoType) SetAllChargesCode(v string) {
	o.AllChargesCode = &v
}

// GetMiniBarCode returns the MiniBarCode field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetMiniBarCode() string {
	if o == nil || IsNil(o.MiniBarCode) {
		var ret string
		return ret
	}
	return *o.MiniBarCode
}

// GetMiniBarCodeOk returns a tuple with the MiniBarCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetMiniBarCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MiniBarCode) {
		return nil, false
	}
	return o.MiniBarCode, true
}

// HasMiniBarCode returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasMiniBarCode() bool {
	if o != nil && !IsNil(o.MiniBarCode) {
		return true
	}

	return false
}

// SetMiniBarCode gets a reference to the given string and assigns it to the MiniBarCode field.
func (o *TransactionCodeInfoType) SetMiniBarCode(v string) {
	o.MiniBarCode = &v
}

// GetCashCreditCardPosting returns the CashCreditCardPosting field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetCashCreditCardPosting() string {
	if o == nil || IsNil(o.CashCreditCardPosting) {
		var ret string
		return ret
	}
	return *o.CashCreditCardPosting
}

// GetCashCreditCardPostingOk returns a tuple with the CashCreditCardPosting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetCashCreditCardPostingOk() (*string, bool) {
	if o == nil || IsNil(o.CashCreditCardPosting) {
		return nil, false
	}
	return o.CashCreditCardPosting, true
}

// HasCashCreditCardPosting returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasCashCreditCardPosting() bool {
	if o != nil && !IsNil(o.CashCreditCardPosting) {
		return true
	}

	return false
}

// SetCashCreditCardPosting gets a reference to the given string and assigns it to the CashCreditCardPosting field.
func (o *TransactionCodeInfoType) SetCashCreditCardPosting(v string) {
	o.CashCreditCardPosting = &v
}

// GetSubtotalPosting returns the SubtotalPosting field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetSubtotalPosting() string {
	if o == nil || IsNil(o.SubtotalPosting) {
		var ret string
		return ret
	}
	return *o.SubtotalPosting
}

// GetSubtotalPostingOk returns a tuple with the SubtotalPosting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetSubtotalPostingOk() (*string, bool) {
	if o == nil || IsNil(o.SubtotalPosting) {
		return nil, false
	}
	return o.SubtotalPosting, true
}

// HasSubtotalPosting returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasSubtotalPosting() bool {
	if o != nil && !IsNil(o.SubtotalPosting) {
		return true
	}

	return false
}

// SetSubtotalPosting gets a reference to the given string and assigns it to the SubtotalPosting field.
func (o *TransactionCodeInfoType) SetSubtotalPosting(v string) {
	o.SubtotalPosting = &v
}

// GetDifferencePosting returns the DifferencePosting field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetDifferencePosting() string {
	if o == nil || IsNil(o.DifferencePosting) {
		var ret string
		return ret
	}
	return *o.DifferencePosting
}

// GetDifferencePostingOk returns a tuple with the DifferencePosting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetDifferencePostingOk() (*string, bool) {
	if o == nil || IsNil(o.DifferencePosting) {
		return nil, false
	}
	return o.DifferencePosting, true
}

// HasDifferencePosting returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasDifferencePosting() bool {
	if o != nil && !IsNil(o.DifferencePosting) {
		return true
	}

	return false
}

// SetDifferencePosting gets a reference to the given string and assigns it to the DifferencePosting field.
func (o *TransactionCodeInfoType) SetDifferencePosting(v string) {
	o.DifferencePosting = &v
}

// GetCalcTrxCodePosting returns the CalcTrxCodePosting field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetCalcTrxCodePosting() string {
	if o == nil || IsNil(o.CalcTrxCodePosting) {
		var ret string
		return ret
	}
	return *o.CalcTrxCodePosting
}

// GetCalcTrxCodePostingOk returns a tuple with the CalcTrxCodePosting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetCalcTrxCodePostingOk() (*string, bool) {
	if o == nil || IsNil(o.CalcTrxCodePosting) {
		return nil, false
	}
	return o.CalcTrxCodePosting, true
}

// HasCalcTrxCodePosting returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasCalcTrxCodePosting() bool {
	if o != nil && !IsNil(o.CalcTrxCodePosting) {
		return true
	}

	return false
}

// SetCalcTrxCodePosting gets a reference to the given string and assigns it to the CalcTrxCodePosting field.
func (o *TransactionCodeInfoType) SetCalcTrxCodePosting(v string) {
	o.CalcTrxCodePosting = &v
}

// GetNumberOfDialedDigits returns the NumberOfDialedDigits field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetNumberOfDialedDigits() string {
	if o == nil || IsNil(o.NumberOfDialedDigits) {
		var ret string
		return ret
	}
	return *o.NumberOfDialedDigits
}

// GetNumberOfDialedDigitsOk returns a tuple with the NumberOfDialedDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetNumberOfDialedDigitsOk() (*string, bool) {
	if o == nil || IsNil(o.NumberOfDialedDigits) {
		return nil, false
	}
	return o.NumberOfDialedDigits, true
}

// HasNumberOfDialedDigits returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasNumberOfDialedDigits() bool {
	if o != nil && !IsNil(o.NumberOfDialedDigits) {
		return true
	}

	return false
}

// SetNumberOfDialedDigits gets a reference to the given string and assigns it to the NumberOfDialedDigits field.
func (o *TransactionCodeInfoType) SetNumberOfDialedDigits(v string) {
	o.NumberOfDialedDigits = &v
}

// GetPostingAccounts returns the PostingAccounts field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetPostingAccounts() []PostingAccountType {
	if o == nil || IsNil(o.PostingAccounts) {
		var ret []PostingAccountType
		return ret
	}
	return o.PostingAccounts
}

// GetPostingAccountsOk returns a tuple with the PostingAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetPostingAccountsOk() ([]PostingAccountType, bool) {
	if o == nil || IsNil(o.PostingAccounts) {
		return nil, false
	}
	return o.PostingAccounts, true
}

// HasPostingAccounts returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasPostingAccounts() bool {
	if o != nil && !IsNil(o.PostingAccounts) {
		return true
	}

	return false
}

// SetPostingAccounts gets a reference to the given []PostingAccountType and assigns it to the PostingAccounts field.
func (o *TransactionCodeInfoType) SetPostingAccounts(v []PostingAccountType) {
	o.PostingAccounts = v
}

// GetItemizers returns the Itemizers field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetItemizers() []ItemizerType {
	if o == nil || IsNil(o.Itemizers) {
		var ret []ItemizerType
		return ret
	}
	return o.Itemizers
}

// GetItemizersOk returns a tuple with the Itemizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetItemizersOk() ([]ItemizerType, bool) {
	if o == nil || IsNil(o.Itemizers) {
		return nil, false
	}
	return o.Itemizers, true
}

// HasItemizers returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasItemizers() bool {
	if o != nil && !IsNil(o.Itemizers) {
		return true
	}

	return false
}

// SetItemizers gets a reference to the given []ItemizerType and assigns it to the Itemizers field.
func (o *TransactionCodeInfoType) SetItemizers(v []ItemizerType) {
	o.Itemizers = v
}

// GetTransactionCodes returns the TransactionCodes field value if set, zero value otherwise.
func (o *TransactionCodeInfoType) GetTransactionCodes() []TransactionCodeDetailType {
	if o == nil || IsNil(o.TransactionCodes) {
		var ret []TransactionCodeDetailType
		return ret
	}
	return o.TransactionCodes
}

// GetTransactionCodesOk returns a tuple with the TransactionCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCodeInfoType) GetTransactionCodesOk() ([]TransactionCodeDetailType, bool) {
	if o == nil || IsNil(o.TransactionCodes) {
		return nil, false
	}
	return o.TransactionCodes, true
}

// HasTransactionCodes returns a boolean if a field has been set.
func (o *TransactionCodeInfoType) HasTransactionCodes() bool {
	if o != nil && !IsNil(o.TransactionCodes) {
		return true
	}

	return false
}

// SetTransactionCodes gets a reference to the given []TransactionCodeDetailType and assigns it to the TransactionCodes field.
func (o *TransactionCodeInfoType) SetTransactionCodes(v []TransactionCodeDetailType) {
	o.TransactionCodes = v
}

func (o TransactionCodeInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionCodeInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllChargesCode) {
		toSerialize["allChargesCode"] = o.AllChargesCode
	}
	if !IsNil(o.MiniBarCode) {
		toSerialize["miniBarCode"] = o.MiniBarCode
	}
	if !IsNil(o.CashCreditCardPosting) {
		toSerialize["cashCreditCardPosting"] = o.CashCreditCardPosting
	}
	if !IsNil(o.SubtotalPosting) {
		toSerialize["subtotalPosting"] = o.SubtotalPosting
	}
	if !IsNil(o.DifferencePosting) {
		toSerialize["differencePosting"] = o.DifferencePosting
	}
	if !IsNil(o.CalcTrxCodePosting) {
		toSerialize["calcTrxCodePosting"] = o.CalcTrxCodePosting
	}
	if !IsNil(o.NumberOfDialedDigits) {
		toSerialize["numberOfDialedDigits"] = o.NumberOfDialedDigits
	}
	if !IsNil(o.PostingAccounts) {
		toSerialize["postingAccounts"] = o.PostingAccounts
	}
	if !IsNil(o.Itemizers) {
		toSerialize["itemizers"] = o.Itemizers
	}
	if !IsNil(o.TransactionCodes) {
		toSerialize["transactionCodes"] = o.TransactionCodes
	}
	return toSerialize, nil
}

type NullableTransactionCodeInfoType struct {
	value *TransactionCodeInfoType
	isSet bool
}

func (v NullableTransactionCodeInfoType) Get() *TransactionCodeInfoType {
	return v.value
}

func (v *NullableTransactionCodeInfoType) Set(val *TransactionCodeInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCodeInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCodeInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCodeInfoType(val *TransactionCodeInfoType) *NullableTransactionCodeInfoType {
	return &NullableTransactionCodeInfoType{value: val, isSet: true}
}

func (v NullableTransactionCodeInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCodeInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


