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

// checks if the ARCreditCardPaymentInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARCreditCardPaymentInfoType{}

// ARCreditCardPaymentInfoType AR Credit Card payment information.
type ARCreditCardPaymentInfoType struct {
	// Guest Name .
	GuestName *string `json:"guestName,omitempty"`
	ProfileId *ProfileId `json:"profileId,omitempty"`
	Amount *CurrencyAmountType `json:"amount,omitempty"`
	// Date of posting.
	PostingDate *string `json:"postingDate,omitempty"`
	// Reference Text for the payment.
	Reference *string `json:"reference,omitempty"`
	// Remarks for payment.
	Remark *string `json:"remark,omitempty"`
	// The Folio number of this posting, if there was a Folio already generated.
	FolioNo *float32 `json:"folioNo,omitempty"`
	// Bill Number returned by the Fiscal Printer.
	FiscalFolioNo *float32 `json:"fiscalFolioNo,omitempty"`
	// Transaction number of the payment.
	TransactionNo *float32 `json:"transactionNo,omitempty"`
	// Transaction Date of the payment.
	TransactionDate *string `json:"transactionDate,omitempty"`
}

// NewARCreditCardPaymentInfoType instantiates a new ARCreditCardPaymentInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARCreditCardPaymentInfoType() *ARCreditCardPaymentInfoType {
	this := ARCreditCardPaymentInfoType{}
	return &this
}

// NewARCreditCardPaymentInfoTypeWithDefaults instantiates a new ARCreditCardPaymentInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARCreditCardPaymentInfoTypeWithDefaults() *ARCreditCardPaymentInfoType {
	this := ARCreditCardPaymentInfoType{}
	return &this
}

// GetGuestName returns the GuestName field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetGuestName() string {
	if o == nil || IsNil(o.GuestName) {
		var ret string
		return ret
	}
	return *o.GuestName
}

// GetGuestNameOk returns a tuple with the GuestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetGuestNameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestName) {
		return nil, false
	}
	return o.GuestName, true
}

// HasGuestName returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasGuestName() bool {
	if o != nil && !IsNil(o.GuestName) {
		return true
	}

	return false
}

// SetGuestName gets a reference to the given string and assigns it to the GuestName field.
func (o *ARCreditCardPaymentInfoType) SetGuestName(v string) {
	o.GuestName = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetProfileId() ProfileId {
	if o == nil || IsNil(o.ProfileId) {
		var ret ProfileId
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetProfileIdOk() (*ProfileId, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given ProfileId and assigns it to the ProfileId field.
func (o *ARCreditCardPaymentInfoType) SetProfileId(v ProfileId) {
	o.ProfileId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetAmount() CurrencyAmountType {
	if o == nil || IsNil(o.Amount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given CurrencyAmountType and assigns it to the Amount field.
func (o *ARCreditCardPaymentInfoType) SetAmount(v CurrencyAmountType) {
	o.Amount = &v
}

// GetPostingDate returns the PostingDate field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetPostingDate() string {
	if o == nil || IsNil(o.PostingDate) {
		var ret string
		return ret
	}
	return *o.PostingDate
}

// GetPostingDateOk returns a tuple with the PostingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetPostingDateOk() (*string, bool) {
	if o == nil || IsNil(o.PostingDate) {
		return nil, false
	}
	return o.PostingDate, true
}

// HasPostingDate returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasPostingDate() bool {
	if o != nil && !IsNil(o.PostingDate) {
		return true
	}

	return false
}

// SetPostingDate gets a reference to the given string and assigns it to the PostingDate field.
func (o *ARCreditCardPaymentInfoType) SetPostingDate(v string) {
	o.PostingDate = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ARCreditCardPaymentInfoType) SetReference(v string) {
	o.Reference = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *ARCreditCardPaymentInfoType) SetRemark(v string) {
	o.Remark = &v
}

// GetFolioNo returns the FolioNo field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetFolioNo() float32 {
	if o == nil || IsNil(o.FolioNo) {
		var ret float32
		return ret
	}
	return *o.FolioNo
}

// GetFolioNoOk returns a tuple with the FolioNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetFolioNoOk() (*float32, bool) {
	if o == nil || IsNil(o.FolioNo) {
		return nil, false
	}
	return o.FolioNo, true
}

// HasFolioNo returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasFolioNo() bool {
	if o != nil && !IsNil(o.FolioNo) {
		return true
	}

	return false
}

// SetFolioNo gets a reference to the given float32 and assigns it to the FolioNo field.
func (o *ARCreditCardPaymentInfoType) SetFolioNo(v float32) {
	o.FolioNo = &v
}

// GetFiscalFolioNo returns the FiscalFolioNo field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetFiscalFolioNo() float32 {
	if o == nil || IsNil(o.FiscalFolioNo) {
		var ret float32
		return ret
	}
	return *o.FiscalFolioNo
}

// GetFiscalFolioNoOk returns a tuple with the FiscalFolioNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetFiscalFolioNoOk() (*float32, bool) {
	if o == nil || IsNil(o.FiscalFolioNo) {
		return nil, false
	}
	return o.FiscalFolioNo, true
}

// HasFiscalFolioNo returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasFiscalFolioNo() bool {
	if o != nil && !IsNil(o.FiscalFolioNo) {
		return true
	}

	return false
}

// SetFiscalFolioNo gets a reference to the given float32 and assigns it to the FiscalFolioNo field.
func (o *ARCreditCardPaymentInfoType) SetFiscalFolioNo(v float32) {
	o.FiscalFolioNo = &v
}

// GetTransactionNo returns the TransactionNo field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetTransactionNo() float32 {
	if o == nil || IsNil(o.TransactionNo) {
		var ret float32
		return ret
	}
	return *o.TransactionNo
}

// GetTransactionNoOk returns a tuple with the TransactionNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetTransactionNoOk() (*float32, bool) {
	if o == nil || IsNil(o.TransactionNo) {
		return nil, false
	}
	return o.TransactionNo, true
}

// HasTransactionNo returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasTransactionNo() bool {
	if o != nil && !IsNil(o.TransactionNo) {
		return true
	}

	return false
}

// SetTransactionNo gets a reference to the given float32 and assigns it to the TransactionNo field.
func (o *ARCreditCardPaymentInfoType) SetTransactionNo(v float32) {
	o.TransactionNo = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *ARCreditCardPaymentInfoType) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARCreditCardPaymentInfoType) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *ARCreditCardPaymentInfoType) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *ARCreditCardPaymentInfoType) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

func (o ARCreditCardPaymentInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARCreditCardPaymentInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuestName) {
		toSerialize["guestName"] = o.GuestName
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.PostingDate) {
		toSerialize["postingDate"] = o.PostingDate
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.FolioNo) {
		toSerialize["folioNo"] = o.FolioNo
	}
	if !IsNil(o.FiscalFolioNo) {
		toSerialize["fiscalFolioNo"] = o.FiscalFolioNo
	}
	if !IsNil(o.TransactionNo) {
		toSerialize["transactionNo"] = o.TransactionNo
	}
	if !IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	return toSerialize, nil
}

type NullableARCreditCardPaymentInfoType struct {
	value *ARCreditCardPaymentInfoType
	isSet bool
}

func (v NullableARCreditCardPaymentInfoType) Get() *ARCreditCardPaymentInfoType {
	return v.value
}

func (v *NullableARCreditCardPaymentInfoType) Set(val *ARCreditCardPaymentInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableARCreditCardPaymentInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableARCreditCardPaymentInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARCreditCardPaymentInfoType(val *ARCreditCardPaymentInfoType) *NullableARCreditCardPaymentInfoType {
	return &NullableARCreditCardPaymentInfoType{value: val, isSet: true}
}

func (v NullableARCreditCardPaymentInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARCreditCardPaymentInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


