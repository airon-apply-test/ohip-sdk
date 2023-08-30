/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the ReservationFolioWindowType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationFolioWindowType{}

// ReservationFolioWindowType Folio window view which holds the set of folios for a reservation.
type ReservationFolioWindowType struct {
	PayeeInfo *PayeeInfoType `json:"payeeInfo,omitempty"`
	Balance *CurrencyAmountType `json:"balance,omitempty"`
	// Payment Method Type
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	FolioWindowNo *int32 `json:"folioWindowNo,omitempty"`
}

// NewReservationFolioWindowType instantiates a new ReservationFolioWindowType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationFolioWindowType() *ReservationFolioWindowType {
	this := ReservationFolioWindowType{}
	return &this
}

// NewReservationFolioWindowTypeWithDefaults instantiates a new ReservationFolioWindowType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationFolioWindowTypeWithDefaults() *ReservationFolioWindowType {
	this := ReservationFolioWindowType{}
	return &this
}

// GetPayeeInfo returns the PayeeInfo field value if set, zero value otherwise.
func (o *ReservationFolioWindowType) GetPayeeInfo() PayeeInfoType {
	if o == nil || IsNil(o.PayeeInfo) {
		var ret PayeeInfoType
		return ret
	}
	return *o.PayeeInfo
}

// GetPayeeInfoOk returns a tuple with the PayeeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFolioWindowType) GetPayeeInfoOk() (*PayeeInfoType, bool) {
	if o == nil || IsNil(o.PayeeInfo) {
		return nil, false
	}
	return o.PayeeInfo, true
}

// HasPayeeInfo returns a boolean if a field has been set.
func (o *ReservationFolioWindowType) HasPayeeInfo() bool {
	if o != nil && !IsNil(o.PayeeInfo) {
		return true
	}

	return false
}

// SetPayeeInfo gets a reference to the given PayeeInfoType and assigns it to the PayeeInfo field.
func (o *ReservationFolioWindowType) SetPayeeInfo(v PayeeInfoType) {
	o.PayeeInfo = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *ReservationFolioWindowType) GetBalance() CurrencyAmountType {
	if o == nil || IsNil(o.Balance) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFolioWindowType) GetBalanceOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *ReservationFolioWindowType) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given CurrencyAmountType and assigns it to the Balance field.
func (o *ReservationFolioWindowType) SetBalance(v CurrencyAmountType) {
	o.Balance = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ReservationFolioWindowType) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFolioWindowType) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ReservationFolioWindowType) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *ReservationFolioWindowType) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetFolioWindowNo returns the FolioWindowNo field value if set, zero value otherwise.
func (o *ReservationFolioWindowType) GetFolioWindowNo() int32 {
	if o == nil || IsNil(o.FolioWindowNo) {
		var ret int32
		return ret
	}
	return *o.FolioWindowNo
}

// GetFolioWindowNoOk returns a tuple with the FolioWindowNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFolioWindowType) GetFolioWindowNoOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioWindowNo) {
		return nil, false
	}
	return o.FolioWindowNo, true
}

// HasFolioWindowNo returns a boolean if a field has been set.
func (o *ReservationFolioWindowType) HasFolioWindowNo() bool {
	if o != nil && !IsNil(o.FolioWindowNo) {
		return true
	}

	return false
}

// SetFolioWindowNo gets a reference to the given int32 and assigns it to the FolioWindowNo field.
func (o *ReservationFolioWindowType) SetFolioWindowNo(v int32) {
	o.FolioWindowNo = &v
}

func (o ReservationFolioWindowType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationFolioWindowType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayeeInfo) {
		toSerialize["payeeInfo"] = o.PayeeInfo
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.FolioWindowNo) {
		toSerialize["folioWindowNo"] = o.FolioWindowNo
	}
	return toSerialize, nil
}

type NullableReservationFolioWindowType struct {
	value *ReservationFolioWindowType
	isSet bool
}

func (v NullableReservationFolioWindowType) Get() *ReservationFolioWindowType {
	return v.value
}

func (v *NullableReservationFolioWindowType) Set(val *ReservationFolioWindowType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationFolioWindowType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationFolioWindowType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationFolioWindowType(val *ReservationFolioWindowType) *NullableReservationFolioWindowType {
	return &NullableReservationFolioWindowType{value: val, isSet: true}
}

func (v NullableReservationFolioWindowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationFolioWindowType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


