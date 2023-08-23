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

// checks if the PaymentTaxInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentTaxInfoType{}

// PaymentTaxInfoType Type for Package Tax Information for Thailand Tax functionality.
type PaymentTaxInfoType struct {
	// Hotel context for the Reservation.
	HotelId *string `json:"hotelId,omitempty"`
	ReservationId *ReservationId `json:"reservationId,omitempty"`
	// Payment Tax record.
	Taxes []PaymentTaxType `json:"taxes,omitempty"`
}

// NewPaymentTaxInfoType instantiates a new PaymentTaxInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentTaxInfoType() *PaymentTaxInfoType {
	this := PaymentTaxInfoType{}
	return &this
}

// NewPaymentTaxInfoTypeWithDefaults instantiates a new PaymentTaxInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentTaxInfoTypeWithDefaults() *PaymentTaxInfoType {
	this := PaymentTaxInfoType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *PaymentTaxInfoType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTaxInfoType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *PaymentTaxInfoType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *PaymentTaxInfoType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *PaymentTaxInfoType) GetReservationId() ReservationId {
	if o == nil || IsNil(o.ReservationId) {
		var ret ReservationId
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTaxInfoType) GetReservationIdOk() (*ReservationId, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *PaymentTaxInfoType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given ReservationId and assigns it to the ReservationId field.
func (o *PaymentTaxInfoType) SetReservationId(v ReservationId) {
	o.ReservationId = &v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *PaymentTaxInfoType) GetTaxes() []PaymentTaxType {
	if o == nil || IsNil(o.Taxes) {
		var ret []PaymentTaxType
		return ret
	}
	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTaxInfoType) GetTaxesOk() ([]PaymentTaxType, bool) {
	if o == nil || IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *PaymentTaxInfoType) HasTaxes() bool {
	if o != nil && !IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given []PaymentTaxType and assigns it to the Taxes field.
func (o *PaymentTaxInfoType) SetTaxes(v []PaymentTaxType) {
	o.Taxes = v
}

func (o PaymentTaxInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentTaxInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	return toSerialize, nil
}

type NullablePaymentTaxInfoType struct {
	value *PaymentTaxInfoType
	isSet bool
}

func (v NullablePaymentTaxInfoType) Get() *PaymentTaxInfoType {
	return v.value
}

func (v *NullablePaymentTaxInfoType) Set(val *PaymentTaxInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentTaxInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentTaxInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentTaxInfoType(val *PaymentTaxInfoType) *NullablePaymentTaxInfoType {
	return &NullablePaymentTaxInfoType{value: val, isSet: true}
}

func (v NullablePaymentTaxInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentTaxInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


