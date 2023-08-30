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

// checks if the ForeignCurrencyExchangeTransactionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForeignCurrencyExchangeTransactionType{}

// ForeignCurrencyExchangeTransactionType Details about posting and generated folio on currency exchange.
type ForeignCurrencyExchangeTransactionType struct {
	Posting *SummaryPostingType `json:"posting,omitempty"`
	Folio *FolioType `json:"folio,omitempty"`
}

// NewForeignCurrencyExchangeTransactionType instantiates a new ForeignCurrencyExchangeTransactionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForeignCurrencyExchangeTransactionType() *ForeignCurrencyExchangeTransactionType {
	this := ForeignCurrencyExchangeTransactionType{}
	return &this
}

// NewForeignCurrencyExchangeTransactionTypeWithDefaults instantiates a new ForeignCurrencyExchangeTransactionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForeignCurrencyExchangeTransactionTypeWithDefaults() *ForeignCurrencyExchangeTransactionType {
	this := ForeignCurrencyExchangeTransactionType{}
	return &this
}

// GetPosting returns the Posting field value if set, zero value otherwise.
func (o *ForeignCurrencyExchangeTransactionType) GetPosting() SummaryPostingType {
	if o == nil || IsNil(o.Posting) {
		var ret SummaryPostingType
		return ret
	}
	return *o.Posting
}

// GetPostingOk returns a tuple with the Posting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForeignCurrencyExchangeTransactionType) GetPostingOk() (*SummaryPostingType, bool) {
	if o == nil || IsNil(o.Posting) {
		return nil, false
	}
	return o.Posting, true
}

// HasPosting returns a boolean if a field has been set.
func (o *ForeignCurrencyExchangeTransactionType) HasPosting() bool {
	if o != nil && !IsNil(o.Posting) {
		return true
	}

	return false
}

// SetPosting gets a reference to the given SummaryPostingType and assigns it to the Posting field.
func (o *ForeignCurrencyExchangeTransactionType) SetPosting(v SummaryPostingType) {
	o.Posting = &v
}

// GetFolio returns the Folio field value if set, zero value otherwise.
func (o *ForeignCurrencyExchangeTransactionType) GetFolio() FolioType {
	if o == nil || IsNil(o.Folio) {
		var ret FolioType
		return ret
	}
	return *o.Folio
}

// GetFolioOk returns a tuple with the Folio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForeignCurrencyExchangeTransactionType) GetFolioOk() (*FolioType, bool) {
	if o == nil || IsNil(o.Folio) {
		return nil, false
	}
	return o.Folio, true
}

// HasFolio returns a boolean if a field has been set.
func (o *ForeignCurrencyExchangeTransactionType) HasFolio() bool {
	if o != nil && !IsNil(o.Folio) {
		return true
	}

	return false
}

// SetFolio gets a reference to the given FolioType and assigns it to the Folio field.
func (o *ForeignCurrencyExchangeTransactionType) SetFolio(v FolioType) {
	o.Folio = &v
}

func (o ForeignCurrencyExchangeTransactionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForeignCurrencyExchangeTransactionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Posting) {
		toSerialize["posting"] = o.Posting
	}
	if !IsNil(o.Folio) {
		toSerialize["folio"] = o.Folio
	}
	return toSerialize, nil
}

type NullableForeignCurrencyExchangeTransactionType struct {
	value *ForeignCurrencyExchangeTransactionType
	isSet bool
}

func (v NullableForeignCurrencyExchangeTransactionType) Get() *ForeignCurrencyExchangeTransactionType {
	return v.value
}

func (v *NullableForeignCurrencyExchangeTransactionType) Set(val *ForeignCurrencyExchangeTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableForeignCurrencyExchangeTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableForeignCurrencyExchangeTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForeignCurrencyExchangeTransactionType(val *ForeignCurrencyExchangeTransactionType) *NullableForeignCurrencyExchangeTransactionType {
	return &NullableForeignCurrencyExchangeTransactionType{value: val, isSet: true}
}

func (v NullableForeignCurrencyExchangeTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForeignCurrencyExchangeTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


