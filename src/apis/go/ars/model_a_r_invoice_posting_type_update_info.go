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

// checks if the ARInvoicePostingTypeUpdateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARInvoicePostingTypeUpdateInfo{}

// ARInvoicePostingTypeUpdateInfo Update info associated to this transaction.
type ARInvoicePostingTypeUpdateInfo struct {
	UpdateDate *string `json:"updateDate,omitempty"`
	UpdateBy *string `json:"updateBy,omitempty"`
}

// NewARInvoicePostingTypeUpdateInfo instantiates a new ARInvoicePostingTypeUpdateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARInvoicePostingTypeUpdateInfo() *ARInvoicePostingTypeUpdateInfo {
	this := ARInvoicePostingTypeUpdateInfo{}
	return &this
}

// NewARInvoicePostingTypeUpdateInfoWithDefaults instantiates a new ARInvoicePostingTypeUpdateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARInvoicePostingTypeUpdateInfoWithDefaults() *ARInvoicePostingTypeUpdateInfo {
	this := ARInvoicePostingTypeUpdateInfo{}
	return &this
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *ARInvoicePostingTypeUpdateInfo) GetUpdateDate() string {
	if o == nil || IsNil(o.UpdateDate) {
		var ret string
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARInvoicePostingTypeUpdateInfo) GetUpdateDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateDate) {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *ARInvoicePostingTypeUpdateInfo) HasUpdateDate() bool {
	if o != nil && !IsNil(o.UpdateDate) {
		return true
	}

	return false
}

// SetUpdateDate gets a reference to the given string and assigns it to the UpdateDate field.
func (o *ARInvoicePostingTypeUpdateInfo) SetUpdateDate(v string) {
	o.UpdateDate = &v
}

// GetUpdateBy returns the UpdateBy field value if set, zero value otherwise.
func (o *ARInvoicePostingTypeUpdateInfo) GetUpdateBy() string {
	if o == nil || IsNil(o.UpdateBy) {
		var ret string
		return ret
	}
	return *o.UpdateBy
}

// GetUpdateByOk returns a tuple with the UpdateBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARInvoicePostingTypeUpdateInfo) GetUpdateByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateBy) {
		return nil, false
	}
	return o.UpdateBy, true
}

// HasUpdateBy returns a boolean if a field has been set.
func (o *ARInvoicePostingTypeUpdateInfo) HasUpdateBy() bool {
	if o != nil && !IsNil(o.UpdateBy) {
		return true
	}

	return false
}

// SetUpdateBy gets a reference to the given string and assigns it to the UpdateBy field.
func (o *ARInvoicePostingTypeUpdateInfo) SetUpdateBy(v string) {
	o.UpdateBy = &v
}

func (o ARInvoicePostingTypeUpdateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARInvoicePostingTypeUpdateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdateDate) {
		toSerialize["updateDate"] = o.UpdateDate
	}
	if !IsNil(o.UpdateBy) {
		toSerialize["updateBy"] = o.UpdateBy
	}
	return toSerialize, nil
}

type NullableARInvoicePostingTypeUpdateInfo struct {
	value *ARInvoicePostingTypeUpdateInfo
	isSet bool
}

func (v NullableARInvoicePostingTypeUpdateInfo) Get() *ARInvoicePostingTypeUpdateInfo {
	return v.value
}

func (v *NullableARInvoicePostingTypeUpdateInfo) Set(val *ARInvoicePostingTypeUpdateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableARInvoicePostingTypeUpdateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableARInvoicePostingTypeUpdateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARInvoicePostingTypeUpdateInfo(val *ARInvoicePostingTypeUpdateInfo) *NullableARInvoicePostingTypeUpdateInfo {
	return &NullableARInvoicePostingTypeUpdateInfo{value: val, isSet: true}
}

func (v NullableARInvoicePostingTypeUpdateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARInvoicePostingTypeUpdateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


