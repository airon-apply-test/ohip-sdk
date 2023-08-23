/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MembershipTransactionTierPointsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipTransactionTierPointsType{}

// MembershipTransactionTierPointsType Details associated with tier points.
type MembershipTransactionTierPointsType struct {
	// Total number of base stay tier points for a membership points transaction.
	BaseStay *int32 `json:"baseStay,omitempty"`
	// Total number of base nights tier points for a membership points transaction.
	BaseNights *int32 `json:"baseNights,omitempty"`
	// Total number of base revenue tier points for a membership points transaction.
	BaseRevenue *int32 `json:"baseRevenue,omitempty"`
	// The total number of bonus stay tier points.
	BonusStay *int32 `json:"bonusStay,omitempty"`
	// The total number of bonus nights tier points.
	BonusNights *int32 `json:"bonusNights,omitempty"`
	// The total number of bonus revenue tier points.
	BonusRevenue *int32 `json:"bonusRevenue,omitempty"`
}

// NewMembershipTransactionTierPointsType instantiates a new MembershipTransactionTierPointsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipTransactionTierPointsType() *MembershipTransactionTierPointsType {
	this := MembershipTransactionTierPointsType{}
	return &this
}

// NewMembershipTransactionTierPointsTypeWithDefaults instantiates a new MembershipTransactionTierPointsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipTransactionTierPointsTypeWithDefaults() *MembershipTransactionTierPointsType {
	this := MembershipTransactionTierPointsType{}
	return &this
}

// GetBaseStay returns the BaseStay field value if set, zero value otherwise.
func (o *MembershipTransactionTierPointsType) GetBaseStay() int32 {
	if o == nil || IsNil(o.BaseStay) {
		var ret int32
		return ret
	}
	return *o.BaseStay
}

// GetBaseStayOk returns a tuple with the BaseStay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTransactionTierPointsType) GetBaseStayOk() (*int32, bool) {
	if o == nil || IsNil(o.BaseStay) {
		return nil, false
	}
	return o.BaseStay, true
}

// HasBaseStay returns a boolean if a field has been set.
func (o *MembershipTransactionTierPointsType) HasBaseStay() bool {
	if o != nil && !IsNil(o.BaseStay) {
		return true
	}

	return false
}

// SetBaseStay gets a reference to the given int32 and assigns it to the BaseStay field.
func (o *MembershipTransactionTierPointsType) SetBaseStay(v int32) {
	o.BaseStay = &v
}

// GetBaseNights returns the BaseNights field value if set, zero value otherwise.
func (o *MembershipTransactionTierPointsType) GetBaseNights() int32 {
	if o == nil || IsNil(o.BaseNights) {
		var ret int32
		return ret
	}
	return *o.BaseNights
}

// GetBaseNightsOk returns a tuple with the BaseNights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTransactionTierPointsType) GetBaseNightsOk() (*int32, bool) {
	if o == nil || IsNil(o.BaseNights) {
		return nil, false
	}
	return o.BaseNights, true
}

// HasBaseNights returns a boolean if a field has been set.
func (o *MembershipTransactionTierPointsType) HasBaseNights() bool {
	if o != nil && !IsNil(o.BaseNights) {
		return true
	}

	return false
}

// SetBaseNights gets a reference to the given int32 and assigns it to the BaseNights field.
func (o *MembershipTransactionTierPointsType) SetBaseNights(v int32) {
	o.BaseNights = &v
}

// GetBaseRevenue returns the BaseRevenue field value if set, zero value otherwise.
func (o *MembershipTransactionTierPointsType) GetBaseRevenue() int32 {
	if o == nil || IsNil(o.BaseRevenue) {
		var ret int32
		return ret
	}
	return *o.BaseRevenue
}

// GetBaseRevenueOk returns a tuple with the BaseRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTransactionTierPointsType) GetBaseRevenueOk() (*int32, bool) {
	if o == nil || IsNil(o.BaseRevenue) {
		return nil, false
	}
	return o.BaseRevenue, true
}

// HasBaseRevenue returns a boolean if a field has been set.
func (o *MembershipTransactionTierPointsType) HasBaseRevenue() bool {
	if o != nil && !IsNil(o.BaseRevenue) {
		return true
	}

	return false
}

// SetBaseRevenue gets a reference to the given int32 and assigns it to the BaseRevenue field.
func (o *MembershipTransactionTierPointsType) SetBaseRevenue(v int32) {
	o.BaseRevenue = &v
}

// GetBonusStay returns the BonusStay field value if set, zero value otherwise.
func (o *MembershipTransactionTierPointsType) GetBonusStay() int32 {
	if o == nil || IsNil(o.BonusStay) {
		var ret int32
		return ret
	}
	return *o.BonusStay
}

// GetBonusStayOk returns a tuple with the BonusStay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTransactionTierPointsType) GetBonusStayOk() (*int32, bool) {
	if o == nil || IsNil(o.BonusStay) {
		return nil, false
	}
	return o.BonusStay, true
}

// HasBonusStay returns a boolean if a field has been set.
func (o *MembershipTransactionTierPointsType) HasBonusStay() bool {
	if o != nil && !IsNil(o.BonusStay) {
		return true
	}

	return false
}

// SetBonusStay gets a reference to the given int32 and assigns it to the BonusStay field.
func (o *MembershipTransactionTierPointsType) SetBonusStay(v int32) {
	o.BonusStay = &v
}

// GetBonusNights returns the BonusNights field value if set, zero value otherwise.
func (o *MembershipTransactionTierPointsType) GetBonusNights() int32 {
	if o == nil || IsNil(o.BonusNights) {
		var ret int32
		return ret
	}
	return *o.BonusNights
}

// GetBonusNightsOk returns a tuple with the BonusNights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTransactionTierPointsType) GetBonusNightsOk() (*int32, bool) {
	if o == nil || IsNil(o.BonusNights) {
		return nil, false
	}
	return o.BonusNights, true
}

// HasBonusNights returns a boolean if a field has been set.
func (o *MembershipTransactionTierPointsType) HasBonusNights() bool {
	if o != nil && !IsNil(o.BonusNights) {
		return true
	}

	return false
}

// SetBonusNights gets a reference to the given int32 and assigns it to the BonusNights field.
func (o *MembershipTransactionTierPointsType) SetBonusNights(v int32) {
	o.BonusNights = &v
}

// GetBonusRevenue returns the BonusRevenue field value if set, zero value otherwise.
func (o *MembershipTransactionTierPointsType) GetBonusRevenue() int32 {
	if o == nil || IsNil(o.BonusRevenue) {
		var ret int32
		return ret
	}
	return *o.BonusRevenue
}

// GetBonusRevenueOk returns a tuple with the BonusRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTransactionTierPointsType) GetBonusRevenueOk() (*int32, bool) {
	if o == nil || IsNil(o.BonusRevenue) {
		return nil, false
	}
	return o.BonusRevenue, true
}

// HasBonusRevenue returns a boolean if a field has been set.
func (o *MembershipTransactionTierPointsType) HasBonusRevenue() bool {
	if o != nil && !IsNil(o.BonusRevenue) {
		return true
	}

	return false
}

// SetBonusRevenue gets a reference to the given int32 and assigns it to the BonusRevenue field.
func (o *MembershipTransactionTierPointsType) SetBonusRevenue(v int32) {
	o.BonusRevenue = &v
}

func (o MembershipTransactionTierPointsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipTransactionTierPointsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseStay) {
		toSerialize["baseStay"] = o.BaseStay
	}
	if !IsNil(o.BaseNights) {
		toSerialize["baseNights"] = o.BaseNights
	}
	if !IsNil(o.BaseRevenue) {
		toSerialize["baseRevenue"] = o.BaseRevenue
	}
	if !IsNil(o.BonusStay) {
		toSerialize["bonusStay"] = o.BonusStay
	}
	if !IsNil(o.BonusNights) {
		toSerialize["bonusNights"] = o.BonusNights
	}
	if !IsNil(o.BonusRevenue) {
		toSerialize["bonusRevenue"] = o.BonusRevenue
	}
	return toSerialize, nil
}

type NullableMembershipTransactionTierPointsType struct {
	value *MembershipTransactionTierPointsType
	isSet bool
}

func (v NullableMembershipTransactionTierPointsType) Get() *MembershipTransactionTierPointsType {
	return v.value
}

func (v *NullableMembershipTransactionTierPointsType) Set(val *MembershipTransactionTierPointsType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipTransactionTierPointsType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipTransactionTierPointsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipTransactionTierPointsType(val *MembershipTransactionTierPointsType) *NullableMembershipTransactionTierPointsType {
	return &NullableMembershipTransactionTierPointsType{value: val, isSet: true}
}

func (v NullableMembershipTransactionTierPointsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipTransactionTierPointsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


