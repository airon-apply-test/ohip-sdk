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

// checks if the PrepaidCardTransactionInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrepaidCardTransactionInfoType{}

// PrepaidCardTransactionInfoType Prepaid Card Transactions
type PrepaidCardTransactionInfoType struct {
	Amount *CurrencyAmountType `json:"amount,omitempty"`
	Type *PrepaidCardTrxTypeType `json:"type,omitempty"`
	// Prepaid card transaction date.
	Date *string `json:"date,omitempty"`
	// Opera transaction number.
	TransactionNo *int32 `json:"transactionNo,omitempty"`
	// Vendor transaction number.
	VendorTransactionNo *string `json:"vendorTransactionNo,omitempty"`
	ProfileId *UniqueIDType `json:"profileId,omitempty"`
	Source *PrepaidCardTransactionSourceType `json:"source,omitempty"`
	// Indicate if the transaction is cancellable or not.
	Cancellable *bool `json:"cancellable,omitempty"`
}

// NewPrepaidCardTransactionInfoType instantiates a new PrepaidCardTransactionInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepaidCardTransactionInfoType() *PrepaidCardTransactionInfoType {
	this := PrepaidCardTransactionInfoType{}
	return &this
}

// NewPrepaidCardTransactionInfoTypeWithDefaults instantiates a new PrepaidCardTransactionInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepaidCardTransactionInfoTypeWithDefaults() *PrepaidCardTransactionInfoType {
	this := PrepaidCardTransactionInfoType{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PrepaidCardTransactionInfoType) GetAmount() CurrencyAmountType {
	if o == nil || IsNil(o.Amount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepaidCardTransactionInfoType) GetAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PrepaidCardTransactionInfoType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given CurrencyAmountType and assigns it to the Amount field.
func (o *PrepaidCardTransactionInfoType) SetAmount(v CurrencyAmountType) {
	o.Amount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PrepaidCardTransactionInfoType) GetType() PrepaidCardTrxTypeType {
	if o == nil || IsNil(o.Type) {
		var ret PrepaidCardTrxTypeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepaidCardTransactionInfoType) GetTypeOk() (*PrepaidCardTrxTypeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PrepaidCardTransactionInfoType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PrepaidCardTrxTypeType and assigns it to the Type field.
func (o *PrepaidCardTransactionInfoType) SetType(v PrepaidCardTrxTypeType) {
	o.Type = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *PrepaidCardTransactionInfoType) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepaidCardTransactionInfoType) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *PrepaidCardTransactionInfoType) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *PrepaidCardTransactionInfoType) SetDate(v string) {
	o.Date = &v
}

// GetTransactionNo returns the TransactionNo field value if set, zero value otherwise.
func (o *PrepaidCardTransactionInfoType) GetTransactionNo() int32 {
	if o == nil || IsNil(o.TransactionNo) {
		var ret int32
		return ret
	}
	return *o.TransactionNo
}

// GetTransactionNoOk returns a tuple with the TransactionNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepaidCardTransactionInfoType) GetTransactionNoOk() (*int32, bool) {
	if o == nil || IsNil(o.TransactionNo) {
		return nil, false
	}
	return o.TransactionNo, true
}

// HasTransactionNo returns a boolean if a field has been set.
func (o *PrepaidCardTransactionInfoType) HasTransactionNo() bool {
	if o != nil && !IsNil(o.TransactionNo) {
		return true
	}

	return false
}

// SetTransactionNo gets a reference to the given int32 and assigns it to the TransactionNo field.
func (o *PrepaidCardTransactionInfoType) SetTransactionNo(v int32) {
	o.TransactionNo = &v
}

// GetVendorTransactionNo returns the VendorTransactionNo field value if set, zero value otherwise.
func (o *PrepaidCardTransactionInfoType) GetVendorTransactionNo() string {
	if o == nil || IsNil(o.VendorTransactionNo) {
		var ret string
		return ret
	}
	return *o.VendorTransactionNo
}

// GetVendorTransactionNoOk returns a tuple with the VendorTransactionNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepaidCardTransactionInfoType) GetVendorTransactionNoOk() (*string, bool) {
	if o == nil || IsNil(o.VendorTransactionNo) {
		return nil, false
	}
	return o.VendorTransactionNo, true
}

// HasVendorTransactionNo returns a boolean if a field has been set.
func (o *PrepaidCardTransactionInfoType) HasVendorTransactionNo() bool {
	if o != nil && !IsNil(o.VendorTransactionNo) {
		return true
	}

	return false
}

// SetVendorTransactionNo gets a reference to the given string and assigns it to the VendorTransactionNo field.
func (o *PrepaidCardTransactionInfoType) SetVendorTransactionNo(v string) {
	o.VendorTransactionNo = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *PrepaidCardTransactionInfoType) GetProfileId() UniqueIDType {
	if o == nil || IsNil(o.ProfileId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepaidCardTransactionInfoType) GetProfileIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *PrepaidCardTransactionInfoType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given UniqueIDType and assigns it to the ProfileId field.
func (o *PrepaidCardTransactionInfoType) SetProfileId(v UniqueIDType) {
	o.ProfileId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PrepaidCardTransactionInfoType) GetSource() PrepaidCardTransactionSourceType {
	if o == nil || IsNil(o.Source) {
		var ret PrepaidCardTransactionSourceType
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepaidCardTransactionInfoType) GetSourceOk() (*PrepaidCardTransactionSourceType, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PrepaidCardTransactionInfoType) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given PrepaidCardTransactionSourceType and assigns it to the Source field.
func (o *PrepaidCardTransactionInfoType) SetSource(v PrepaidCardTransactionSourceType) {
	o.Source = &v
}

// GetCancellable returns the Cancellable field value if set, zero value otherwise.
func (o *PrepaidCardTransactionInfoType) GetCancellable() bool {
	if o == nil || IsNil(o.Cancellable) {
		var ret bool
		return ret
	}
	return *o.Cancellable
}

// GetCancellableOk returns a tuple with the Cancellable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepaidCardTransactionInfoType) GetCancellableOk() (*bool, bool) {
	if o == nil || IsNil(o.Cancellable) {
		return nil, false
	}
	return o.Cancellable, true
}

// HasCancellable returns a boolean if a field has been set.
func (o *PrepaidCardTransactionInfoType) HasCancellable() bool {
	if o != nil && !IsNil(o.Cancellable) {
		return true
	}

	return false
}

// SetCancellable gets a reference to the given bool and assigns it to the Cancellable field.
func (o *PrepaidCardTransactionInfoType) SetCancellable(v bool) {
	o.Cancellable = &v
}

func (o PrepaidCardTransactionInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrepaidCardTransactionInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.TransactionNo) {
		toSerialize["transactionNo"] = o.TransactionNo
	}
	if !IsNil(o.VendorTransactionNo) {
		toSerialize["vendorTransactionNo"] = o.VendorTransactionNo
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Cancellable) {
		toSerialize["cancellable"] = o.Cancellable
	}
	return toSerialize, nil
}

type NullablePrepaidCardTransactionInfoType struct {
	value *PrepaidCardTransactionInfoType
	isSet bool
}

func (v NullablePrepaidCardTransactionInfoType) Get() *PrepaidCardTransactionInfoType {
	return v.value
}

func (v *NullablePrepaidCardTransactionInfoType) Set(val *PrepaidCardTransactionInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepaidCardTransactionInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepaidCardTransactionInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepaidCardTransactionInfoType(val *PrepaidCardTransactionInfoType) *NullablePrepaidCardTransactionInfoType {
	return &NullablePrepaidCardTransactionInfoType{value: val, isSet: true}
}

func (v NullablePrepaidCardTransactionInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrepaidCardTransactionInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


