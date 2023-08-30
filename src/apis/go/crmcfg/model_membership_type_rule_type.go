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
	"fmt"
)

// MembershipTypeRuleType Membership Type/Program rule types.
type MembershipTypeRuleType string

// List of membershipTypeRuleType
const (
	MEMBERSHIPTYPERULETYPE_ALL MembershipTypeRuleType = "All"
	MEMBERSHIPTYPERULETYPE_POINTS MembershipTypeRuleType = "Points"
	MEMBERSHIPTYPERULETYPE_TIER MembershipTypeRuleType = "Tier"
)

// All allowed values of MembershipTypeRuleType enum
var AllowedMembershipTypeRuleTypeEnumValues = []MembershipTypeRuleType{
	"All",
	"Points",
	"Tier",
}

func (v *MembershipTypeRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MembershipTypeRuleType(value)
	for _, existing := range AllowedMembershipTypeRuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MembershipTypeRuleType", value)
}

// NewMembershipTypeRuleTypeFromValue returns a pointer to a valid MembershipTypeRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMembershipTypeRuleTypeFromValue(v string) (*MembershipTypeRuleType, error) {
	ev := MembershipTypeRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MembershipTypeRuleType: valid values are %v", v, AllowedMembershipTypeRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MembershipTypeRuleType) IsValid() bool {
	for _, existing := range AllowedMembershipTypeRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to membershipTypeRuleType value
func (v MembershipTypeRuleType) Ptr() *MembershipTypeRuleType {
	return &v
}

type NullableMembershipTypeRuleType struct {
	value *MembershipTypeRuleType
	isSet bool
}

func (v NullableMembershipTypeRuleType) Get() *MembershipTypeRuleType {
	return v.value
}

func (v *NullableMembershipTypeRuleType) Set(val *MembershipTypeRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipTypeRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipTypeRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipTypeRuleType(val *MembershipTypeRuleType) *NullableMembershipTypeRuleType {
	return &NullableMembershipTypeRuleType{value: val, isSet: true}
}

func (v NullableMembershipTypeRuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipTypeRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

