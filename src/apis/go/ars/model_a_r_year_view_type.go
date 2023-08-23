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

// checks if the ARYearViewType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARYearViewType{}

// ARYearViewType Year information for an Year View.
type ARYearViewType struct {
	DateRange *DateRangeType `json:"dateRange,omitempty"`
	BalanceInfo *ARBalanceType `json:"balanceInfo,omitempty"`
	RunningTotal *CurrencyAmountType `json:"runningTotal,omitempty"`
}

// NewARYearViewType instantiates a new ARYearViewType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARYearViewType() *ARYearViewType {
	this := ARYearViewType{}
	return &this
}

// NewARYearViewTypeWithDefaults instantiates a new ARYearViewType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARYearViewTypeWithDefaults() *ARYearViewType {
	this := ARYearViewType{}
	return &this
}

// GetDateRange returns the DateRange field value if set, zero value otherwise.
func (o *ARYearViewType) GetDateRange() DateRangeType {
	if o == nil || IsNil(o.DateRange) {
		var ret DateRangeType
		return ret
	}
	return *o.DateRange
}

// GetDateRangeOk returns a tuple with the DateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARYearViewType) GetDateRangeOk() (*DateRangeType, bool) {
	if o == nil || IsNil(o.DateRange) {
		return nil, false
	}
	return o.DateRange, true
}

// HasDateRange returns a boolean if a field has been set.
func (o *ARYearViewType) HasDateRange() bool {
	if o != nil && !IsNil(o.DateRange) {
		return true
	}

	return false
}

// SetDateRange gets a reference to the given DateRangeType and assigns it to the DateRange field.
func (o *ARYearViewType) SetDateRange(v DateRangeType) {
	o.DateRange = &v
}

// GetBalanceInfo returns the BalanceInfo field value if set, zero value otherwise.
func (o *ARYearViewType) GetBalanceInfo() ARBalanceType {
	if o == nil || IsNil(o.BalanceInfo) {
		var ret ARBalanceType
		return ret
	}
	return *o.BalanceInfo
}

// GetBalanceInfoOk returns a tuple with the BalanceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARYearViewType) GetBalanceInfoOk() (*ARBalanceType, bool) {
	if o == nil || IsNil(o.BalanceInfo) {
		return nil, false
	}
	return o.BalanceInfo, true
}

// HasBalanceInfo returns a boolean if a field has been set.
func (o *ARYearViewType) HasBalanceInfo() bool {
	if o != nil && !IsNil(o.BalanceInfo) {
		return true
	}

	return false
}

// SetBalanceInfo gets a reference to the given ARBalanceType and assigns it to the BalanceInfo field.
func (o *ARYearViewType) SetBalanceInfo(v ARBalanceType) {
	o.BalanceInfo = &v
}

// GetRunningTotal returns the RunningTotal field value if set, zero value otherwise.
func (o *ARYearViewType) GetRunningTotal() CurrencyAmountType {
	if o == nil || IsNil(o.RunningTotal) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.RunningTotal
}

// GetRunningTotalOk returns a tuple with the RunningTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARYearViewType) GetRunningTotalOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.RunningTotal) {
		return nil, false
	}
	return o.RunningTotal, true
}

// HasRunningTotal returns a boolean if a field has been set.
func (o *ARYearViewType) HasRunningTotal() bool {
	if o != nil && !IsNil(o.RunningTotal) {
		return true
	}

	return false
}

// SetRunningTotal gets a reference to the given CurrencyAmountType and assigns it to the RunningTotal field.
func (o *ARYearViewType) SetRunningTotal(v CurrencyAmountType) {
	o.RunningTotal = &v
}

func (o ARYearViewType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARYearViewType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateRange) {
		toSerialize["dateRange"] = o.DateRange
	}
	if !IsNil(o.BalanceInfo) {
		toSerialize["balanceInfo"] = o.BalanceInfo
	}
	if !IsNil(o.RunningTotal) {
		toSerialize["runningTotal"] = o.RunningTotal
	}
	return toSerialize, nil
}

type NullableARYearViewType struct {
	value *ARYearViewType
	isSet bool
}

func (v NullableARYearViewType) Get() *ARYearViewType {
	return v.value
}

func (v *NullableARYearViewType) Set(val *ARYearViewType) {
	v.value = val
	v.isSet = true
}

func (v NullableARYearViewType) IsSet() bool {
	return v.isSet
}

func (v *NullableARYearViewType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARYearViewType(val *ARYearViewType) *NullableARYearViewType {
	return &NullableARYearViewType{value: val, isSet: true}
}

func (v NullableARYearViewType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARYearViewType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


