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

// checks if the ReservationOveragePaymentsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationOveragePaymentsType{}

// ReservationOveragePaymentsType List of Reservation details for payment that has a folio window balance equal or higher to the credit limit set for the credit card payment method of that folio window.
type ReservationOveragePaymentsType struct {
	// Reservation details to initiate the Credit Limit Overage process
	ReservationOveragePayment []ReservationOveragePaymentType `json:"reservationOveragePayment,omitempty"`
	// Identifies the hotel code.
	HotelId *string `json:"hotelId,omitempty"`
	// The Cashier ID of the Cashier who is currently processing the transaction(s).
	CashierId *float32 `json:"cashierId,omitempty"`
}

// NewReservationOveragePaymentsType instantiates a new ReservationOveragePaymentsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationOveragePaymentsType() *ReservationOveragePaymentsType {
	this := ReservationOveragePaymentsType{}
	return &this
}

// NewReservationOveragePaymentsTypeWithDefaults instantiates a new ReservationOveragePaymentsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationOveragePaymentsTypeWithDefaults() *ReservationOveragePaymentsType {
	this := ReservationOveragePaymentsType{}
	return &this
}

// GetReservationOveragePayment returns the ReservationOveragePayment field value if set, zero value otherwise.
func (o *ReservationOveragePaymentsType) GetReservationOveragePayment() []ReservationOveragePaymentType {
	if o == nil || IsNil(o.ReservationOveragePayment) {
		var ret []ReservationOveragePaymentType
		return ret
	}
	return o.ReservationOveragePayment
}

// GetReservationOveragePaymentOk returns a tuple with the ReservationOveragePayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationOveragePaymentsType) GetReservationOveragePaymentOk() ([]ReservationOveragePaymentType, bool) {
	if o == nil || IsNil(o.ReservationOveragePayment) {
		return nil, false
	}
	return o.ReservationOveragePayment, true
}

// HasReservationOveragePayment returns a boolean if a field has been set.
func (o *ReservationOveragePaymentsType) HasReservationOveragePayment() bool {
	if o != nil && !IsNil(o.ReservationOveragePayment) {
		return true
	}

	return false
}

// SetReservationOveragePayment gets a reference to the given []ReservationOveragePaymentType and assigns it to the ReservationOveragePayment field.
func (o *ReservationOveragePaymentsType) SetReservationOveragePayment(v []ReservationOveragePaymentType) {
	o.ReservationOveragePayment = v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ReservationOveragePaymentsType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationOveragePaymentsType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ReservationOveragePaymentsType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ReservationOveragePaymentsType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *ReservationOveragePaymentsType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationOveragePaymentsType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *ReservationOveragePaymentsType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *ReservationOveragePaymentsType) SetCashierId(v float32) {
	o.CashierId = &v
}

func (o ReservationOveragePaymentsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationOveragePaymentsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationOveragePayment) {
		toSerialize["reservationOveragePayment"] = o.ReservationOveragePayment
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.CashierId) {
		toSerialize["cashierId"] = o.CashierId
	}
	return toSerialize, nil
}

type NullableReservationOveragePaymentsType struct {
	value *ReservationOveragePaymentsType
	isSet bool
}

func (v NullableReservationOveragePaymentsType) Get() *ReservationOveragePaymentsType {
	return v.value
}

func (v *NullableReservationOveragePaymentsType) Set(val *ReservationOveragePaymentsType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationOveragePaymentsType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationOveragePaymentsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationOveragePaymentsType(val *ReservationOveragePaymentsType) *NullableReservationOveragePaymentsType {
	return &NullableReservationOveragePaymentsType{value: val, isSet: true}
}

func (v NullableReservationOveragePaymentsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationOveragePaymentsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


