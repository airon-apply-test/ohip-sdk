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

// checks if the ReservationPolicySummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationPolicySummaryType{}

// ReservationPolicySummaryType A collection of reservation deposit and cancellation policies.
type ReservationPolicySummaryType struct {
	// A list of reservation cancellation policies.
	CancellationPolicies []ResCancellationPolicyType `json:"cancellationPolicies,omitempty"`
	// A list of deposit policies attached with the reservation.
	DepositPolicies []ResDepositPolicyType `json:"depositPolicies,omitempty"`
	// Unique Id that references an object uniquely in the system.
	ReservationIdList []UniqueIDType `json:"reservationIdList,omitempty"`
	// Name identifier for the reservation.
	Name *string `json:"name,omitempty"`
}

// NewReservationPolicySummaryType instantiates a new ReservationPolicySummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationPolicySummaryType() *ReservationPolicySummaryType {
	this := ReservationPolicySummaryType{}
	return &this
}

// NewReservationPolicySummaryTypeWithDefaults instantiates a new ReservationPolicySummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationPolicySummaryTypeWithDefaults() *ReservationPolicySummaryType {
	this := ReservationPolicySummaryType{}
	return &this
}

// GetCancellationPolicies returns the CancellationPolicies field value if set, zero value otherwise.
func (o *ReservationPolicySummaryType) GetCancellationPolicies() []ResCancellationPolicyType {
	if o == nil || IsNil(o.CancellationPolicies) {
		var ret []ResCancellationPolicyType
		return ret
	}
	return o.CancellationPolicies
}

// GetCancellationPoliciesOk returns a tuple with the CancellationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPolicySummaryType) GetCancellationPoliciesOk() ([]ResCancellationPolicyType, bool) {
	if o == nil || IsNil(o.CancellationPolicies) {
		return nil, false
	}
	return o.CancellationPolicies, true
}

// HasCancellationPolicies returns a boolean if a field has been set.
func (o *ReservationPolicySummaryType) HasCancellationPolicies() bool {
	if o != nil && !IsNil(o.CancellationPolicies) {
		return true
	}

	return false
}

// SetCancellationPolicies gets a reference to the given []ResCancellationPolicyType and assigns it to the CancellationPolicies field.
func (o *ReservationPolicySummaryType) SetCancellationPolicies(v []ResCancellationPolicyType) {
	o.CancellationPolicies = v
}

// GetDepositPolicies returns the DepositPolicies field value if set, zero value otherwise.
func (o *ReservationPolicySummaryType) GetDepositPolicies() []ResDepositPolicyType {
	if o == nil || IsNil(o.DepositPolicies) {
		var ret []ResDepositPolicyType
		return ret
	}
	return o.DepositPolicies
}

// GetDepositPoliciesOk returns a tuple with the DepositPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPolicySummaryType) GetDepositPoliciesOk() ([]ResDepositPolicyType, bool) {
	if o == nil || IsNil(o.DepositPolicies) {
		return nil, false
	}
	return o.DepositPolicies, true
}

// HasDepositPolicies returns a boolean if a field has been set.
func (o *ReservationPolicySummaryType) HasDepositPolicies() bool {
	if o != nil && !IsNil(o.DepositPolicies) {
		return true
	}

	return false
}

// SetDepositPolicies gets a reference to the given []ResDepositPolicyType and assigns it to the DepositPolicies field.
func (o *ReservationPolicySummaryType) SetDepositPolicies(v []ResDepositPolicyType) {
	o.DepositPolicies = v
}

// GetReservationIdList returns the ReservationIdList field value if set, zero value otherwise.
func (o *ReservationPolicySummaryType) GetReservationIdList() []UniqueIDType {
	if o == nil || IsNil(o.ReservationIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ReservationIdList
}

// GetReservationIdListOk returns a tuple with the ReservationIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPolicySummaryType) GetReservationIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationIdList) {
		return nil, false
	}
	return o.ReservationIdList, true
}

// HasReservationIdList returns a boolean if a field has been set.
func (o *ReservationPolicySummaryType) HasReservationIdList() bool {
	if o != nil && !IsNil(o.ReservationIdList) {
		return true
	}

	return false
}

// SetReservationIdList gets a reference to the given []UniqueIDType and assigns it to the ReservationIdList field.
func (o *ReservationPolicySummaryType) SetReservationIdList(v []UniqueIDType) {
	o.ReservationIdList = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReservationPolicySummaryType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPolicySummaryType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReservationPolicySummaryType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReservationPolicySummaryType) SetName(v string) {
	o.Name = &v
}

func (o ReservationPolicySummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationPolicySummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancellationPolicies) {
		toSerialize["cancellationPolicies"] = o.CancellationPolicies
	}
	if !IsNil(o.DepositPolicies) {
		toSerialize["depositPolicies"] = o.DepositPolicies
	}
	if !IsNil(o.ReservationIdList) {
		toSerialize["reservationIdList"] = o.ReservationIdList
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableReservationPolicySummaryType struct {
	value *ReservationPolicySummaryType
	isSet bool
}

func (v NullableReservationPolicySummaryType) Get() *ReservationPolicySummaryType {
	return v.value
}

func (v *NullableReservationPolicySummaryType) Set(val *ReservationPolicySummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationPolicySummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationPolicySummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationPolicySummaryType(val *ReservationPolicySummaryType) *NullableReservationPolicySummaryType {
	return &NullableReservationPolicySummaryType{value: val, isSet: true}
}

func (v NullableReservationPolicySummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationPolicySummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


