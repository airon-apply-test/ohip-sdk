/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResRevenueBalanceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResRevenueBalanceType{}

// ResRevenueBalanceType Revenue and Balance Amount summary for the reservation.
type ResRevenueBalanceType struct {
	TotalFixedCharge *CurrencyAmountType `json:"totalFixedCharge,omitempty"`
	TotalPayment *CurrencyAmountType `json:"totalPayment,omitempty"`
	RoomRevenue *CurrencyAmountType `json:"roomRevenue,omitempty"`
	FoodAndBevRevenue *CurrencyAmountType `json:"foodAndBevRevenue,omitempty"`
	TotalRevenue *CurrencyAmountType `json:"totalRevenue,omitempty"`
	Balance *CurrencyAmountType `json:"balance,omitempty"`
	CompBalance *CurrencyAmountType `json:"compBalance,omitempty"`
}

// NewResRevenueBalanceType instantiates a new ResRevenueBalanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResRevenueBalanceType() *ResRevenueBalanceType {
	this := ResRevenueBalanceType{}
	return &this
}

// NewResRevenueBalanceTypeWithDefaults instantiates a new ResRevenueBalanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResRevenueBalanceTypeWithDefaults() *ResRevenueBalanceType {
	this := ResRevenueBalanceType{}
	return &this
}

// GetTotalFixedCharge returns the TotalFixedCharge field value if set, zero value otherwise.
func (o *ResRevenueBalanceType) GetTotalFixedCharge() CurrencyAmountType {
	if o == nil || IsNil(o.TotalFixedCharge) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.TotalFixedCharge
}

// GetTotalFixedChargeOk returns a tuple with the TotalFixedCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResRevenueBalanceType) GetTotalFixedChargeOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.TotalFixedCharge) {
		return nil, false
	}
	return o.TotalFixedCharge, true
}

// HasTotalFixedCharge returns a boolean if a field has been set.
func (o *ResRevenueBalanceType) HasTotalFixedCharge() bool {
	if o != nil && !IsNil(o.TotalFixedCharge) {
		return true
	}

	return false
}

// SetTotalFixedCharge gets a reference to the given CurrencyAmountType and assigns it to the TotalFixedCharge field.
func (o *ResRevenueBalanceType) SetTotalFixedCharge(v CurrencyAmountType) {
	o.TotalFixedCharge = &v
}

// GetTotalPayment returns the TotalPayment field value if set, zero value otherwise.
func (o *ResRevenueBalanceType) GetTotalPayment() CurrencyAmountType {
	if o == nil || IsNil(o.TotalPayment) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.TotalPayment
}

// GetTotalPaymentOk returns a tuple with the TotalPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResRevenueBalanceType) GetTotalPaymentOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.TotalPayment) {
		return nil, false
	}
	return o.TotalPayment, true
}

// HasTotalPayment returns a boolean if a field has been set.
func (o *ResRevenueBalanceType) HasTotalPayment() bool {
	if o != nil && !IsNil(o.TotalPayment) {
		return true
	}

	return false
}

// SetTotalPayment gets a reference to the given CurrencyAmountType and assigns it to the TotalPayment field.
func (o *ResRevenueBalanceType) SetTotalPayment(v CurrencyAmountType) {
	o.TotalPayment = &v
}

// GetRoomRevenue returns the RoomRevenue field value if set, zero value otherwise.
func (o *ResRevenueBalanceType) GetRoomRevenue() CurrencyAmountType {
	if o == nil || IsNil(o.RoomRevenue) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.RoomRevenue
}

// GetRoomRevenueOk returns a tuple with the RoomRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResRevenueBalanceType) GetRoomRevenueOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.RoomRevenue) {
		return nil, false
	}
	return o.RoomRevenue, true
}

// HasRoomRevenue returns a boolean if a field has been set.
func (o *ResRevenueBalanceType) HasRoomRevenue() bool {
	if o != nil && !IsNil(o.RoomRevenue) {
		return true
	}

	return false
}

// SetRoomRevenue gets a reference to the given CurrencyAmountType and assigns it to the RoomRevenue field.
func (o *ResRevenueBalanceType) SetRoomRevenue(v CurrencyAmountType) {
	o.RoomRevenue = &v
}

// GetFoodAndBevRevenue returns the FoodAndBevRevenue field value if set, zero value otherwise.
func (o *ResRevenueBalanceType) GetFoodAndBevRevenue() CurrencyAmountType {
	if o == nil || IsNil(o.FoodAndBevRevenue) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.FoodAndBevRevenue
}

// GetFoodAndBevRevenueOk returns a tuple with the FoodAndBevRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResRevenueBalanceType) GetFoodAndBevRevenueOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.FoodAndBevRevenue) {
		return nil, false
	}
	return o.FoodAndBevRevenue, true
}

// HasFoodAndBevRevenue returns a boolean if a field has been set.
func (o *ResRevenueBalanceType) HasFoodAndBevRevenue() bool {
	if o != nil && !IsNil(o.FoodAndBevRevenue) {
		return true
	}

	return false
}

// SetFoodAndBevRevenue gets a reference to the given CurrencyAmountType and assigns it to the FoodAndBevRevenue field.
func (o *ResRevenueBalanceType) SetFoodAndBevRevenue(v CurrencyAmountType) {
	o.FoodAndBevRevenue = &v
}

// GetTotalRevenue returns the TotalRevenue field value if set, zero value otherwise.
func (o *ResRevenueBalanceType) GetTotalRevenue() CurrencyAmountType {
	if o == nil || IsNil(o.TotalRevenue) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResRevenueBalanceType) GetTotalRevenueOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.TotalRevenue) {
		return nil, false
	}
	return o.TotalRevenue, true
}

// HasTotalRevenue returns a boolean if a field has been set.
func (o *ResRevenueBalanceType) HasTotalRevenue() bool {
	if o != nil && !IsNil(o.TotalRevenue) {
		return true
	}

	return false
}

// SetTotalRevenue gets a reference to the given CurrencyAmountType and assigns it to the TotalRevenue field.
func (o *ResRevenueBalanceType) SetTotalRevenue(v CurrencyAmountType) {
	o.TotalRevenue = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *ResRevenueBalanceType) GetBalance() CurrencyAmountType {
	if o == nil || IsNil(o.Balance) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResRevenueBalanceType) GetBalanceOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *ResRevenueBalanceType) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given CurrencyAmountType and assigns it to the Balance field.
func (o *ResRevenueBalanceType) SetBalance(v CurrencyAmountType) {
	o.Balance = &v
}

// GetCompBalance returns the CompBalance field value if set, zero value otherwise.
func (o *ResRevenueBalanceType) GetCompBalance() CurrencyAmountType {
	if o == nil || IsNil(o.CompBalance) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.CompBalance
}

// GetCompBalanceOk returns a tuple with the CompBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResRevenueBalanceType) GetCompBalanceOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.CompBalance) {
		return nil, false
	}
	return o.CompBalance, true
}

// HasCompBalance returns a boolean if a field has been set.
func (o *ResRevenueBalanceType) HasCompBalance() bool {
	if o != nil && !IsNil(o.CompBalance) {
		return true
	}

	return false
}

// SetCompBalance gets a reference to the given CurrencyAmountType and assigns it to the CompBalance field.
func (o *ResRevenueBalanceType) SetCompBalance(v CurrencyAmountType) {
	o.CompBalance = &v
}

func (o ResRevenueBalanceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResRevenueBalanceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalFixedCharge) {
		toSerialize["totalFixedCharge"] = o.TotalFixedCharge
	}
	if !IsNil(o.TotalPayment) {
		toSerialize["totalPayment"] = o.TotalPayment
	}
	if !IsNil(o.RoomRevenue) {
		toSerialize["roomRevenue"] = o.RoomRevenue
	}
	if !IsNil(o.FoodAndBevRevenue) {
		toSerialize["foodAndBevRevenue"] = o.FoodAndBevRevenue
	}
	if !IsNil(o.TotalRevenue) {
		toSerialize["totalRevenue"] = o.TotalRevenue
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.CompBalance) {
		toSerialize["compBalance"] = o.CompBalance
	}
	return toSerialize, nil
}

type NullableResRevenueBalanceType struct {
	value *ResRevenueBalanceType
	isSet bool
}

func (v NullableResRevenueBalanceType) Get() *ResRevenueBalanceType {
	return v.value
}

func (v *NullableResRevenueBalanceType) Set(val *ResRevenueBalanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableResRevenueBalanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableResRevenueBalanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResRevenueBalanceType(val *ResRevenueBalanceType) *NullableResRevenueBalanceType {
	return &NullableResRevenueBalanceType{value: val, isSet: true}
}

func (v NullableResRevenueBalanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResRevenueBalanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


