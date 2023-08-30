/*
OPERA Cloud Front Desk Configuration API

APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fofcfg

import (
	"encoding/json"
)

// checks if the ReservationBasedCommissionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationBasedCommissionType{}

// ReservationBasedCommissionType Reservation based commission details.
type ReservationBasedCommissionType struct {
	CommissionPaidDetails *CommissionPaidDetailsType `json:"commissionPaidDetails,omitempty"`
}

// NewReservationBasedCommissionType instantiates a new ReservationBasedCommissionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationBasedCommissionType() *ReservationBasedCommissionType {
	this := ReservationBasedCommissionType{}
	return &this
}

// NewReservationBasedCommissionTypeWithDefaults instantiates a new ReservationBasedCommissionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationBasedCommissionTypeWithDefaults() *ReservationBasedCommissionType {
	this := ReservationBasedCommissionType{}
	return &this
}

// GetCommissionPaidDetails returns the CommissionPaidDetails field value if set, zero value otherwise.
func (o *ReservationBasedCommissionType) GetCommissionPaidDetails() CommissionPaidDetailsType {
	if o == nil || IsNil(o.CommissionPaidDetails) {
		var ret CommissionPaidDetailsType
		return ret
	}
	return *o.CommissionPaidDetails
}

// GetCommissionPaidDetailsOk returns a tuple with the CommissionPaidDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationBasedCommissionType) GetCommissionPaidDetailsOk() (*CommissionPaidDetailsType, bool) {
	if o == nil || IsNil(o.CommissionPaidDetails) {
		return nil, false
	}
	return o.CommissionPaidDetails, true
}

// HasCommissionPaidDetails returns a boolean if a field has been set.
func (o *ReservationBasedCommissionType) HasCommissionPaidDetails() bool {
	if o != nil && !IsNil(o.CommissionPaidDetails) {
		return true
	}

	return false
}

// SetCommissionPaidDetails gets a reference to the given CommissionPaidDetailsType and assigns it to the CommissionPaidDetails field.
func (o *ReservationBasedCommissionType) SetCommissionPaidDetails(v CommissionPaidDetailsType) {
	o.CommissionPaidDetails = &v
}

func (o ReservationBasedCommissionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationBasedCommissionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommissionPaidDetails) {
		toSerialize["commissionPaidDetails"] = o.CommissionPaidDetails
	}
	return toSerialize, nil
}

type NullableReservationBasedCommissionType struct {
	value *ReservationBasedCommissionType
	isSet bool
}

func (v NullableReservationBasedCommissionType) Get() *ReservationBasedCommissionType {
	return v.value
}

func (v *NullableReservationBasedCommissionType) Set(val *ReservationBasedCommissionType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationBasedCommissionType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationBasedCommissionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationBasedCommissionType(val *ReservationBasedCommissionType) *NullableReservationBasedCommissionType {
	return &NullableReservationBasedCommissionType{value: val, isSet: true}
}

func (v NullableReservationBasedCommissionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationBasedCommissionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


