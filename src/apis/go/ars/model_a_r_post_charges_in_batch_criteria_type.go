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

// checks if the ARPostChargesInBatchCriteriaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARPostChargesInBatchCriteriaType{}

// ARPostChargesInBatchCriteriaType Criteria for posting a charge to list of accounts.
type ARPostChargesInBatchCriteriaType struct {
	// Unique Id that references an object uniquely in the system.
	AccountIdList []UniqueIDType `json:"accountIdList,omitempty"`
	ChargeInfo *ChargeCriteriaType `json:"chargeInfo,omitempty"`
	// The Cashier ID of the Cashier who is currently processing the transaction(s).
	CashierId *float32 `json:"cashierId,omitempty"`
	HotelId *string `json:"hotelId,omitempty"`
}

// NewARPostChargesInBatchCriteriaType instantiates a new ARPostChargesInBatchCriteriaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARPostChargesInBatchCriteriaType() *ARPostChargesInBatchCriteriaType {
	this := ARPostChargesInBatchCriteriaType{}
	return &this
}

// NewARPostChargesInBatchCriteriaTypeWithDefaults instantiates a new ARPostChargesInBatchCriteriaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARPostChargesInBatchCriteriaTypeWithDefaults() *ARPostChargesInBatchCriteriaType {
	this := ARPostChargesInBatchCriteriaType{}
	return &this
}

// GetAccountIdList returns the AccountIdList field value if set, zero value otherwise.
func (o *ARPostChargesInBatchCriteriaType) GetAccountIdList() []UniqueIDType {
	if o == nil || IsNil(o.AccountIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.AccountIdList
}

// GetAccountIdListOk returns a tuple with the AccountIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARPostChargesInBatchCriteriaType) GetAccountIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.AccountIdList) {
		return nil, false
	}
	return o.AccountIdList, true
}

// HasAccountIdList returns a boolean if a field has been set.
func (o *ARPostChargesInBatchCriteriaType) HasAccountIdList() bool {
	if o != nil && !IsNil(o.AccountIdList) {
		return true
	}

	return false
}

// SetAccountIdList gets a reference to the given []UniqueIDType and assigns it to the AccountIdList field.
func (o *ARPostChargesInBatchCriteriaType) SetAccountIdList(v []UniqueIDType) {
	o.AccountIdList = v
}

// GetChargeInfo returns the ChargeInfo field value if set, zero value otherwise.
func (o *ARPostChargesInBatchCriteriaType) GetChargeInfo() ChargeCriteriaType {
	if o == nil || IsNil(o.ChargeInfo) {
		var ret ChargeCriteriaType
		return ret
	}
	return *o.ChargeInfo
}

// GetChargeInfoOk returns a tuple with the ChargeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARPostChargesInBatchCriteriaType) GetChargeInfoOk() (*ChargeCriteriaType, bool) {
	if o == nil || IsNil(o.ChargeInfo) {
		return nil, false
	}
	return o.ChargeInfo, true
}

// HasChargeInfo returns a boolean if a field has been set.
func (o *ARPostChargesInBatchCriteriaType) HasChargeInfo() bool {
	if o != nil && !IsNil(o.ChargeInfo) {
		return true
	}

	return false
}

// SetChargeInfo gets a reference to the given ChargeCriteriaType and assigns it to the ChargeInfo field.
func (o *ARPostChargesInBatchCriteriaType) SetChargeInfo(v ChargeCriteriaType) {
	o.ChargeInfo = &v
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *ARPostChargesInBatchCriteriaType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARPostChargesInBatchCriteriaType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *ARPostChargesInBatchCriteriaType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *ARPostChargesInBatchCriteriaType) SetCashierId(v float32) {
	o.CashierId = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ARPostChargesInBatchCriteriaType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARPostChargesInBatchCriteriaType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ARPostChargesInBatchCriteriaType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ARPostChargesInBatchCriteriaType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o ARPostChargesInBatchCriteriaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARPostChargesInBatchCriteriaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountIdList) {
		toSerialize["accountIdList"] = o.AccountIdList
	}
	if !IsNil(o.ChargeInfo) {
		toSerialize["chargeInfo"] = o.ChargeInfo
	}
	if !IsNil(o.CashierId) {
		toSerialize["cashierId"] = o.CashierId
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableARPostChargesInBatchCriteriaType struct {
	value *ARPostChargesInBatchCriteriaType
	isSet bool
}

func (v NullableARPostChargesInBatchCriteriaType) Get() *ARPostChargesInBatchCriteriaType {
	return v.value
}

func (v *NullableARPostChargesInBatchCriteriaType) Set(val *ARPostChargesInBatchCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableARPostChargesInBatchCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableARPostChargesInBatchCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARPostChargesInBatchCriteriaType(val *ARPostChargesInBatchCriteriaType) *NullableARPostChargesInBatchCriteriaType {
	return &NullableARPostChargesInBatchCriteriaType{value: val, isSet: true}
}

func (v NullableARPostChargesInBatchCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARPostChargesInBatchCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


