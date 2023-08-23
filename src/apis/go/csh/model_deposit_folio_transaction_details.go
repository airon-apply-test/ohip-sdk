/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DepositFolioTransactionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositFolioTransactionDetails{}

// DepositFolioTransactionDetails Response for the fetch transaction details request. Detail information regarding the folio transaction will be returned
type DepositFolioTransactionDetails struct {
	// List of Deposit Postings with details.
	Transactions []DepositDetailPostingType `json:"transactions,omitempty"`
	// List of Transaction codes info.
	TrxCodesInfo []TrxInfoType `json:"trxCodesInfo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewDepositFolioTransactionDetails instantiates a new DepositFolioTransactionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositFolioTransactionDetails() *DepositFolioTransactionDetails {
	this := DepositFolioTransactionDetails{}
	return &this
}

// NewDepositFolioTransactionDetailsWithDefaults instantiates a new DepositFolioTransactionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositFolioTransactionDetailsWithDefaults() *DepositFolioTransactionDetails {
	this := DepositFolioTransactionDetails{}
	return &this
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *DepositFolioTransactionDetails) GetTransactions() []DepositDetailPostingType {
	if o == nil || IsNil(o.Transactions) {
		var ret []DepositDetailPostingType
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositFolioTransactionDetails) GetTransactionsOk() ([]DepositDetailPostingType, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *DepositFolioTransactionDetails) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []DepositDetailPostingType and assigns it to the Transactions field.
func (o *DepositFolioTransactionDetails) SetTransactions(v []DepositDetailPostingType) {
	o.Transactions = v
}

// GetTrxCodesInfo returns the TrxCodesInfo field value if set, zero value otherwise.
func (o *DepositFolioTransactionDetails) GetTrxCodesInfo() []TrxInfoType {
	if o == nil || IsNil(o.TrxCodesInfo) {
		var ret []TrxInfoType
		return ret
	}
	return o.TrxCodesInfo
}

// GetTrxCodesInfoOk returns a tuple with the TrxCodesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositFolioTransactionDetails) GetTrxCodesInfoOk() ([]TrxInfoType, bool) {
	if o == nil || IsNil(o.TrxCodesInfo) {
		return nil, false
	}
	return o.TrxCodesInfo, true
}

// HasTrxCodesInfo returns a boolean if a field has been set.
func (o *DepositFolioTransactionDetails) HasTrxCodesInfo() bool {
	if o != nil && !IsNil(o.TrxCodesInfo) {
		return true
	}

	return false
}

// SetTrxCodesInfo gets a reference to the given []TrxInfoType and assigns it to the TrxCodesInfo field.
func (o *DepositFolioTransactionDetails) SetTrxCodesInfo(v []TrxInfoType) {
	o.TrxCodesInfo = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DepositFolioTransactionDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositFolioTransactionDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DepositFolioTransactionDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *DepositFolioTransactionDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DepositFolioTransactionDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositFolioTransactionDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DepositFolioTransactionDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *DepositFolioTransactionDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o DepositFolioTransactionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositFolioTransactionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
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

type NullableDepositFolioTransactionDetails struct {
	value *DepositFolioTransactionDetails
	isSet bool
}

func (v NullableDepositFolioTransactionDetails) Get() *DepositFolioTransactionDetails {
	return v.value
}

func (v *NullableDepositFolioTransactionDetails) Set(val *DepositFolioTransactionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositFolioTransactionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositFolioTransactionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositFolioTransactionDetails(val *DepositFolioTransactionDetails) *NullableDepositFolioTransactionDetails {
	return &NullableDepositFolioTransactionDetails{value: val, isSet: true}
}

func (v NullableDepositFolioTransactionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositFolioTransactionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


