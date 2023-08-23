/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MembershipAwardUpgradeInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipAwardUpgradeInfoType{}

// MembershipAwardUpgradeInfoType Membership Award cancel related details.
type MembershipAwardUpgradeInfoType struct {
	// If a reservation is booked using award points, this is the number of days before the arrival date by which the reservation may be cancelled without penalty. Zero indicates that the reservation may be cancelled any time up to and including the arrival date without incurring a penalty.
	CancelPenaltyDays *float32 `json:"cancelPenaltyDays,omitempty"`
	// The flat number of award points, or the percentage of the award points, that are forfeited if the guest cancels the reservation fewer than the number of days specified in Cancel Penalty Days before the arrival date.
	CancelPenaltyCharge *float32 `json:"cancelPenaltyCharge,omitempty"`
	CancelPenaltyType *MembershipAwardCancelPenaltyType `json:"cancelPenaltyType,omitempty"`
	CancelPolicyType *MembershipAwardCancelPolicyType `json:"cancelPolicyType,omitempty"`
	// Number of nights cancel policy is applicable.
	NumberOfNights *int32 `json:"numberOfNights,omitempty"`
	// Indicates if the membership award upgrade is based on room category (N) or room category group (Y).
	BasedOnRoomGroup *bool `json:"basedOnRoomGroup,omitempty"`
}

// NewMembershipAwardUpgradeInfoType instantiates a new MembershipAwardUpgradeInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipAwardUpgradeInfoType() *MembershipAwardUpgradeInfoType {
	this := MembershipAwardUpgradeInfoType{}
	return &this
}

// NewMembershipAwardUpgradeInfoTypeWithDefaults instantiates a new MembershipAwardUpgradeInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipAwardUpgradeInfoTypeWithDefaults() *MembershipAwardUpgradeInfoType {
	this := MembershipAwardUpgradeInfoType{}
	return &this
}

// GetCancelPenaltyDays returns the CancelPenaltyDays field value if set, zero value otherwise.
func (o *MembershipAwardUpgradeInfoType) GetCancelPenaltyDays() float32 {
	if o == nil || IsNil(o.CancelPenaltyDays) {
		var ret float32
		return ret
	}
	return *o.CancelPenaltyDays
}

// GetCancelPenaltyDaysOk returns a tuple with the CancelPenaltyDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwardUpgradeInfoType) GetCancelPenaltyDaysOk() (*float32, bool) {
	if o == nil || IsNil(o.CancelPenaltyDays) {
		return nil, false
	}
	return o.CancelPenaltyDays, true
}

// HasCancelPenaltyDays returns a boolean if a field has been set.
func (o *MembershipAwardUpgradeInfoType) HasCancelPenaltyDays() bool {
	if o != nil && !IsNil(o.CancelPenaltyDays) {
		return true
	}

	return false
}

// SetCancelPenaltyDays gets a reference to the given float32 and assigns it to the CancelPenaltyDays field.
func (o *MembershipAwardUpgradeInfoType) SetCancelPenaltyDays(v float32) {
	o.CancelPenaltyDays = &v
}

// GetCancelPenaltyCharge returns the CancelPenaltyCharge field value if set, zero value otherwise.
func (o *MembershipAwardUpgradeInfoType) GetCancelPenaltyCharge() float32 {
	if o == nil || IsNil(o.CancelPenaltyCharge) {
		var ret float32
		return ret
	}
	return *o.CancelPenaltyCharge
}

// GetCancelPenaltyChargeOk returns a tuple with the CancelPenaltyCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwardUpgradeInfoType) GetCancelPenaltyChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.CancelPenaltyCharge) {
		return nil, false
	}
	return o.CancelPenaltyCharge, true
}

// HasCancelPenaltyCharge returns a boolean if a field has been set.
func (o *MembershipAwardUpgradeInfoType) HasCancelPenaltyCharge() bool {
	if o != nil && !IsNil(o.CancelPenaltyCharge) {
		return true
	}

	return false
}

// SetCancelPenaltyCharge gets a reference to the given float32 and assigns it to the CancelPenaltyCharge field.
func (o *MembershipAwardUpgradeInfoType) SetCancelPenaltyCharge(v float32) {
	o.CancelPenaltyCharge = &v
}

// GetCancelPenaltyType returns the CancelPenaltyType field value if set, zero value otherwise.
func (o *MembershipAwardUpgradeInfoType) GetCancelPenaltyType() MembershipAwardCancelPenaltyType {
	if o == nil || IsNil(o.CancelPenaltyType) {
		var ret MembershipAwardCancelPenaltyType
		return ret
	}
	return *o.CancelPenaltyType
}

// GetCancelPenaltyTypeOk returns a tuple with the CancelPenaltyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwardUpgradeInfoType) GetCancelPenaltyTypeOk() (*MembershipAwardCancelPenaltyType, bool) {
	if o == nil || IsNil(o.CancelPenaltyType) {
		return nil, false
	}
	return o.CancelPenaltyType, true
}

// HasCancelPenaltyType returns a boolean if a field has been set.
func (o *MembershipAwardUpgradeInfoType) HasCancelPenaltyType() bool {
	if o != nil && !IsNil(o.CancelPenaltyType) {
		return true
	}

	return false
}

// SetCancelPenaltyType gets a reference to the given MembershipAwardCancelPenaltyType and assigns it to the CancelPenaltyType field.
func (o *MembershipAwardUpgradeInfoType) SetCancelPenaltyType(v MembershipAwardCancelPenaltyType) {
	o.CancelPenaltyType = &v
}

// GetCancelPolicyType returns the CancelPolicyType field value if set, zero value otherwise.
func (o *MembershipAwardUpgradeInfoType) GetCancelPolicyType() MembershipAwardCancelPolicyType {
	if o == nil || IsNil(o.CancelPolicyType) {
		var ret MembershipAwardCancelPolicyType
		return ret
	}
	return *o.CancelPolicyType
}

// GetCancelPolicyTypeOk returns a tuple with the CancelPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwardUpgradeInfoType) GetCancelPolicyTypeOk() (*MembershipAwardCancelPolicyType, bool) {
	if o == nil || IsNil(o.CancelPolicyType) {
		return nil, false
	}
	return o.CancelPolicyType, true
}

// HasCancelPolicyType returns a boolean if a field has been set.
func (o *MembershipAwardUpgradeInfoType) HasCancelPolicyType() bool {
	if o != nil && !IsNil(o.CancelPolicyType) {
		return true
	}

	return false
}

// SetCancelPolicyType gets a reference to the given MembershipAwardCancelPolicyType and assigns it to the CancelPolicyType field.
func (o *MembershipAwardUpgradeInfoType) SetCancelPolicyType(v MembershipAwardCancelPolicyType) {
	o.CancelPolicyType = &v
}

// GetNumberOfNights returns the NumberOfNights field value if set, zero value otherwise.
func (o *MembershipAwardUpgradeInfoType) GetNumberOfNights() int32 {
	if o == nil || IsNil(o.NumberOfNights) {
		var ret int32
		return ret
	}
	return *o.NumberOfNights
}

// GetNumberOfNightsOk returns a tuple with the NumberOfNights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwardUpgradeInfoType) GetNumberOfNightsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfNights) {
		return nil, false
	}
	return o.NumberOfNights, true
}

// HasNumberOfNights returns a boolean if a field has been set.
func (o *MembershipAwardUpgradeInfoType) HasNumberOfNights() bool {
	if o != nil && !IsNil(o.NumberOfNights) {
		return true
	}

	return false
}

// SetNumberOfNights gets a reference to the given int32 and assigns it to the NumberOfNights field.
func (o *MembershipAwardUpgradeInfoType) SetNumberOfNights(v int32) {
	o.NumberOfNights = &v
}

// GetBasedOnRoomGroup returns the BasedOnRoomGroup field value if set, zero value otherwise.
func (o *MembershipAwardUpgradeInfoType) GetBasedOnRoomGroup() bool {
	if o == nil || IsNil(o.BasedOnRoomGroup) {
		var ret bool
		return ret
	}
	return *o.BasedOnRoomGroup
}

// GetBasedOnRoomGroupOk returns a tuple with the BasedOnRoomGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwardUpgradeInfoType) GetBasedOnRoomGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.BasedOnRoomGroup) {
		return nil, false
	}
	return o.BasedOnRoomGroup, true
}

// HasBasedOnRoomGroup returns a boolean if a field has been set.
func (o *MembershipAwardUpgradeInfoType) HasBasedOnRoomGroup() bool {
	if o != nil && !IsNil(o.BasedOnRoomGroup) {
		return true
	}

	return false
}

// SetBasedOnRoomGroup gets a reference to the given bool and assigns it to the BasedOnRoomGroup field.
func (o *MembershipAwardUpgradeInfoType) SetBasedOnRoomGroup(v bool) {
	o.BasedOnRoomGroup = &v
}

func (o MembershipAwardUpgradeInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipAwardUpgradeInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancelPenaltyDays) {
		toSerialize["cancelPenaltyDays"] = o.CancelPenaltyDays
	}
	if !IsNil(o.CancelPenaltyCharge) {
		toSerialize["cancelPenaltyCharge"] = o.CancelPenaltyCharge
	}
	if !IsNil(o.CancelPenaltyType) {
		toSerialize["cancelPenaltyType"] = o.CancelPenaltyType
	}
	if !IsNil(o.CancelPolicyType) {
		toSerialize["cancelPolicyType"] = o.CancelPolicyType
	}
	if !IsNil(o.NumberOfNights) {
		toSerialize["numberOfNights"] = o.NumberOfNights
	}
	if !IsNil(o.BasedOnRoomGroup) {
		toSerialize["basedOnRoomGroup"] = o.BasedOnRoomGroup
	}
	return toSerialize, nil
}

type NullableMembershipAwardUpgradeInfoType struct {
	value *MembershipAwardUpgradeInfoType
	isSet bool
}

func (v NullableMembershipAwardUpgradeInfoType) Get() *MembershipAwardUpgradeInfoType {
	return v.value
}

func (v *NullableMembershipAwardUpgradeInfoType) Set(val *MembershipAwardUpgradeInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipAwardUpgradeInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipAwardUpgradeInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipAwardUpgradeInfoType(val *MembershipAwardUpgradeInfoType) *NullableMembershipAwardUpgradeInfoType {
	return &NullableMembershipAwardUpgradeInfoType{value: val, isSet: true}
}

func (v NullableMembershipAwardUpgradeInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipAwardUpgradeInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


