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

// checks if the ARYearViewInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARYearViewInfoType{}

// ARYearViewInfoType Information regarding Year View balances for an account.
type ARYearViewInfoType struct {
	TotalOutstanding *ARBalanceType `json:"totalOutstanding,omitempty"`
	// The debit and credit balance per account.
	YearView []ARYearViewType `json:"yearView,omitempty"`
}

// NewARYearViewInfoType instantiates a new ARYearViewInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARYearViewInfoType() *ARYearViewInfoType {
	this := ARYearViewInfoType{}
	return &this
}

// NewARYearViewInfoTypeWithDefaults instantiates a new ARYearViewInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARYearViewInfoTypeWithDefaults() *ARYearViewInfoType {
	this := ARYearViewInfoType{}
	return &this
}

// GetTotalOutstanding returns the TotalOutstanding field value if set, zero value otherwise.
func (o *ARYearViewInfoType) GetTotalOutstanding() ARBalanceType {
	if o == nil || IsNil(o.TotalOutstanding) {
		var ret ARBalanceType
		return ret
	}
	return *o.TotalOutstanding
}

// GetTotalOutstandingOk returns a tuple with the TotalOutstanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARYearViewInfoType) GetTotalOutstandingOk() (*ARBalanceType, bool) {
	if o == nil || IsNil(o.TotalOutstanding) {
		return nil, false
	}
	return o.TotalOutstanding, true
}

// HasTotalOutstanding returns a boolean if a field has been set.
func (o *ARYearViewInfoType) HasTotalOutstanding() bool {
	if o != nil && !IsNil(o.TotalOutstanding) {
		return true
	}

	return false
}

// SetTotalOutstanding gets a reference to the given ARBalanceType and assigns it to the TotalOutstanding field.
func (o *ARYearViewInfoType) SetTotalOutstanding(v ARBalanceType) {
	o.TotalOutstanding = &v
}

// GetYearView returns the YearView field value if set, zero value otherwise.
func (o *ARYearViewInfoType) GetYearView() []ARYearViewType {
	if o == nil || IsNil(o.YearView) {
		var ret []ARYearViewType
		return ret
	}
	return o.YearView
}

// GetYearViewOk returns a tuple with the YearView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARYearViewInfoType) GetYearViewOk() ([]ARYearViewType, bool) {
	if o == nil || IsNil(o.YearView) {
		return nil, false
	}
	return o.YearView, true
}

// HasYearView returns a boolean if a field has been set.
func (o *ARYearViewInfoType) HasYearView() bool {
	if o != nil && !IsNil(o.YearView) {
		return true
	}

	return false
}

// SetYearView gets a reference to the given []ARYearViewType and assigns it to the YearView field.
func (o *ARYearViewInfoType) SetYearView(v []ARYearViewType) {
	o.YearView = v
}

func (o ARYearViewInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARYearViewInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalOutstanding) {
		toSerialize["totalOutstanding"] = o.TotalOutstanding
	}
	if !IsNil(o.YearView) {
		toSerialize["yearView"] = o.YearView
	}
	return toSerialize, nil
}

type NullableARYearViewInfoType struct {
	value *ARYearViewInfoType
	isSet bool
}

func (v NullableARYearViewInfoType) Get() *ARYearViewInfoType {
	return v.value
}

func (v *NullableARYearViewInfoType) Set(val *ARYearViewInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableARYearViewInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableARYearViewInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARYearViewInfoType(val *ARYearViewInfoType) *NullableARYearViewInfoType {
	return &NullableARYearViewInfoType{value: val, isSet: true}
}

func (v NullableARYearViewInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARYearViewInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


