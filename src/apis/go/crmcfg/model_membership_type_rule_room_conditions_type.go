/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmcfg

import (
	"encoding/json"
)

// checks if the MembershipTypeRuleRoomConditionsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipTypeRuleRoomConditionsType{}

// MembershipTypeRuleRoomConditionsType Membership Type Rules room conditions.
type MembershipTypeRuleRoomConditionsType struct {
	// Room type/label for which the rule is applied.
	RoomType *string `json:"roomType,omitempty"`
	// Membership room group for which the rule is applied.
	RoomGroup *string `json:"roomGroup,omitempty"`
	// Room class for the room for which the rule is applied.
	RoomClass *string `json:"roomClass,omitempty"`
	// Indicates whether to use Room Type or Booked Room Type for the rule. If Y then Booked Room Type will be used for points calculation otherwise Room Type to be used.
	RoomsToCharge *bool `json:"roomsToCharge,omitempty"`
}

// NewMembershipTypeRuleRoomConditionsType instantiates a new MembershipTypeRuleRoomConditionsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipTypeRuleRoomConditionsType() *MembershipTypeRuleRoomConditionsType {
	this := MembershipTypeRuleRoomConditionsType{}
	return &this
}

// NewMembershipTypeRuleRoomConditionsTypeWithDefaults instantiates a new MembershipTypeRuleRoomConditionsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipTypeRuleRoomConditionsTypeWithDefaults() *MembershipTypeRuleRoomConditionsType {
	this := MembershipTypeRuleRoomConditionsType{}
	return &this
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *MembershipTypeRuleRoomConditionsType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTypeRuleRoomConditionsType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *MembershipTypeRuleRoomConditionsType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *MembershipTypeRuleRoomConditionsType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetRoomGroup returns the RoomGroup field value if set, zero value otherwise.
func (o *MembershipTypeRuleRoomConditionsType) GetRoomGroup() string {
	if o == nil || IsNil(o.RoomGroup) {
		var ret string
		return ret
	}
	return *o.RoomGroup
}

// GetRoomGroupOk returns a tuple with the RoomGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTypeRuleRoomConditionsType) GetRoomGroupOk() (*string, bool) {
	if o == nil || IsNil(o.RoomGroup) {
		return nil, false
	}
	return o.RoomGroup, true
}

// HasRoomGroup returns a boolean if a field has been set.
func (o *MembershipTypeRuleRoomConditionsType) HasRoomGroup() bool {
	if o != nil && !IsNil(o.RoomGroup) {
		return true
	}

	return false
}

// SetRoomGroup gets a reference to the given string and assigns it to the RoomGroup field.
func (o *MembershipTypeRuleRoomConditionsType) SetRoomGroup(v string) {
	o.RoomGroup = &v
}

// GetRoomClass returns the RoomClass field value if set, zero value otherwise.
func (o *MembershipTypeRuleRoomConditionsType) GetRoomClass() string {
	if o == nil || IsNil(o.RoomClass) {
		var ret string
		return ret
	}
	return *o.RoomClass
}

// GetRoomClassOk returns a tuple with the RoomClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTypeRuleRoomConditionsType) GetRoomClassOk() (*string, bool) {
	if o == nil || IsNil(o.RoomClass) {
		return nil, false
	}
	return o.RoomClass, true
}

// HasRoomClass returns a boolean if a field has been set.
func (o *MembershipTypeRuleRoomConditionsType) HasRoomClass() bool {
	if o != nil && !IsNil(o.RoomClass) {
		return true
	}

	return false
}

// SetRoomClass gets a reference to the given string and assigns it to the RoomClass field.
func (o *MembershipTypeRuleRoomConditionsType) SetRoomClass(v string) {
	o.RoomClass = &v
}

// GetRoomsToCharge returns the RoomsToCharge field value if set, zero value otherwise.
func (o *MembershipTypeRuleRoomConditionsType) GetRoomsToCharge() bool {
	if o == nil || IsNil(o.RoomsToCharge) {
		var ret bool
		return ret
	}
	return *o.RoomsToCharge
}

// GetRoomsToChargeOk returns a tuple with the RoomsToCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipTypeRuleRoomConditionsType) GetRoomsToChargeOk() (*bool, bool) {
	if o == nil || IsNil(o.RoomsToCharge) {
		return nil, false
	}
	return o.RoomsToCharge, true
}

// HasRoomsToCharge returns a boolean if a field has been set.
func (o *MembershipTypeRuleRoomConditionsType) HasRoomsToCharge() bool {
	if o != nil && !IsNil(o.RoomsToCharge) {
		return true
	}

	return false
}

// SetRoomsToCharge gets a reference to the given bool and assigns it to the RoomsToCharge field.
func (o *MembershipTypeRuleRoomConditionsType) SetRoomsToCharge(v bool) {
	o.RoomsToCharge = &v
}

func (o MembershipTypeRuleRoomConditionsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipTypeRuleRoomConditionsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.RoomGroup) {
		toSerialize["roomGroup"] = o.RoomGroup
	}
	if !IsNil(o.RoomClass) {
		toSerialize["roomClass"] = o.RoomClass
	}
	if !IsNil(o.RoomsToCharge) {
		toSerialize["roomsToCharge"] = o.RoomsToCharge
	}
	return toSerialize, nil
}

type NullableMembershipTypeRuleRoomConditionsType struct {
	value *MembershipTypeRuleRoomConditionsType
	isSet bool
}

func (v NullableMembershipTypeRuleRoomConditionsType) Get() *MembershipTypeRuleRoomConditionsType {
	return v.value
}

func (v *NullableMembershipTypeRuleRoomConditionsType) Set(val *MembershipTypeRuleRoomConditionsType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipTypeRuleRoomConditionsType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipTypeRuleRoomConditionsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipTypeRuleRoomConditionsType(val *MembershipTypeRuleRoomConditionsType) *NullableMembershipTypeRuleRoomConditionsType {
	return &NullableMembershipTypeRuleRoomConditionsType{value: val, isSet: true}
}

func (v NullableMembershipTypeRuleRoomConditionsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipTypeRuleRoomConditionsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


