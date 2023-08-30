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

// checks if the ApplyFlexibleBenefitPostingsCriteriaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyFlexibleBenefitPostingsCriteriaType{}

// ApplyFlexibleBenefitPostingsCriteriaType Criteria for retrieving one or more guest's folio transactions.
type ApplyFlexibleBenefitPostingsCriteriaType struct {
	// Hotel context for the Reservations.
	HotelId *string `json:"hotelId,omitempty"`
	ReservationId *ReservationId `json:"reservationId,omitempty"`
	// The Cashier ID of the Cashier who is currently processing the transaction(s).
	CashierId *float32 `json:"cashierId,omitempty"`
}

// NewApplyFlexibleBenefitPostingsCriteriaType instantiates a new ApplyFlexibleBenefitPostingsCriteriaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyFlexibleBenefitPostingsCriteriaType() *ApplyFlexibleBenefitPostingsCriteriaType {
	this := ApplyFlexibleBenefitPostingsCriteriaType{}
	return &this
}

// NewApplyFlexibleBenefitPostingsCriteriaTypeWithDefaults instantiates a new ApplyFlexibleBenefitPostingsCriteriaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyFlexibleBenefitPostingsCriteriaTypeWithDefaults() *ApplyFlexibleBenefitPostingsCriteriaType {
	this := ApplyFlexibleBenefitPostingsCriteriaType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) GetReservationId() ReservationId {
	if o == nil || IsNil(o.ReservationId) {
		var ret ReservationId
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) GetReservationIdOk() (*ReservationId, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given ReservationId and assigns it to the ReservationId field.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) SetReservationId(v ReservationId) {
	o.ReservationId = &v
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *ApplyFlexibleBenefitPostingsCriteriaType) SetCashierId(v float32) {
	o.CashierId = &v
}

func (o ApplyFlexibleBenefitPostingsCriteriaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyFlexibleBenefitPostingsCriteriaType) ToMap() (map[string]interface{}, error) {
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

type NullableApplyFlexibleBenefitPostingsCriteriaType struct {
	value *ApplyFlexibleBenefitPostingsCriteriaType
	isSet bool
}

func (v NullableApplyFlexibleBenefitPostingsCriteriaType) Get() *ApplyFlexibleBenefitPostingsCriteriaType {
	return v.value
}

func (v *NullableApplyFlexibleBenefitPostingsCriteriaType) Set(val *ApplyFlexibleBenefitPostingsCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyFlexibleBenefitPostingsCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyFlexibleBenefitPostingsCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyFlexibleBenefitPostingsCriteriaType(val *ApplyFlexibleBenefitPostingsCriteriaType) *NullableApplyFlexibleBenefitPostingsCriteriaType {
	return &NullableApplyFlexibleBenefitPostingsCriteriaType{value: val, isSet: true}
}

func (v NullableApplyFlexibleBenefitPostingsCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyFlexibleBenefitPostingsCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


