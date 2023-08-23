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

// checks if the PostDepositToGuestFolioType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostDepositToGuestFolioType{}

// PostDepositToGuestFolioType Criteria for transferring the Deposit Amount to the Guest Folio
type PostDepositToGuestFolioType struct {
	// Resort for which the reservation defined.
	HotelId *string `json:"hotelId,omitempty"`
	ReservationId *ReservationId `json:"reservationId,omitempty"`
	// The Cashier ID of the Cashier who is currently processing the transaction(s).
	CashierId *float32 `json:"cashierId,omitempty"`
}

// NewPostDepositToGuestFolioType instantiates a new PostDepositToGuestFolioType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostDepositToGuestFolioType() *PostDepositToGuestFolioType {
	this := PostDepositToGuestFolioType{}
	return &this
}

// NewPostDepositToGuestFolioTypeWithDefaults instantiates a new PostDepositToGuestFolioType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostDepositToGuestFolioTypeWithDefaults() *PostDepositToGuestFolioType {
	this := PostDepositToGuestFolioType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *PostDepositToGuestFolioType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDepositToGuestFolioType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *PostDepositToGuestFolioType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *PostDepositToGuestFolioType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *PostDepositToGuestFolioType) GetReservationId() ReservationId {
	if o == nil || IsNil(o.ReservationId) {
		var ret ReservationId
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDepositToGuestFolioType) GetReservationIdOk() (*ReservationId, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *PostDepositToGuestFolioType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given ReservationId and assigns it to the ReservationId field.
func (o *PostDepositToGuestFolioType) SetReservationId(v ReservationId) {
	o.ReservationId = &v
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *PostDepositToGuestFolioType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDepositToGuestFolioType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *PostDepositToGuestFolioType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *PostDepositToGuestFolioType) SetCashierId(v float32) {
	o.CashierId = &v
}

func (o PostDepositToGuestFolioType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostDepositToGuestFolioType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.CashierId) {
		toSerialize["cashierId"] = o.CashierId
	}
	return toSerialize, nil
}

type NullablePostDepositToGuestFolioType struct {
	value *PostDepositToGuestFolioType
	isSet bool
}

func (v NullablePostDepositToGuestFolioType) Get() *PostDepositToGuestFolioType {
	return v.value
}

func (v *NullablePostDepositToGuestFolioType) Set(val *PostDepositToGuestFolioType) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDepositToGuestFolioType) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDepositToGuestFolioType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDepositToGuestFolioType(val *PostDepositToGuestFolioType) *NullablePostDepositToGuestFolioType {
	return &NullablePostDepositToGuestFolioType{value: val, isSet: true}
}

func (v NullablePostDepositToGuestFolioType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDepositToGuestFolioType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


