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

// MembershipTypeRuleBasedOnType Indicates Membership upgrades/downgrades to the next tier level uses RFM (Recency, Frequency, Monetary Value) score.
type MembershipTypeRuleBasedOnType string

// List of membershipTypeRuleBasedOnType
const (
	MEMBERSHIPTYPERULEBASEDONTYPE_REVENUE MembershipTypeRuleBasedOnType = "Revenue"
	MEMBERSHIPTYPERULEBASEDONTYPE_STAY MembershipTypeRuleBasedOnType = "Stay"
	MEMBERSHIPTYPERULEBASEDONTYPE_NIGHTS MembershipTypeRuleBasedOnType = "Nights"
	MEMBERSHIPTYPERULEBASEDONTYPE_ENROLLMENT MembershipTypeRuleBasedOnType = "Enrollment"
	MEMBERSHIPTYPERULEBASEDONTYPE_TIER_UPGRADE MembershipTypeRuleBasedOnType = "TierUpgrade"
	MEMBERSHIPTYPERULEBASEDONTYPE_RENEWAL MembershipTypeRuleBasedOnType = "Renewal"
	MEMBERSHIPTYPERULEBASEDONTYPE_RFM MembershipTypeRuleBasedOnType = "Rfm"
)

// All allowed values of MembershipTypeRuleBasedOnType enum
var AllowedMembershipTypeRuleBasedOnTypeEnumValues = []MembershipTypeRuleBasedOnType{
	"Revenue",
	"Stay",
	"Nights",
	"Enrollment",
	"TierUpgrade",
	"Renewal",
	"Rfm",
}

func (v *MembershipTypeRuleBasedOnType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MembershipTypeRuleBasedOnType(value)
	for _, existing := range AllowedMembershipTypeRuleBasedOnTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MembershipTypeRuleBasedOnType", value)
}

// NewMembershipTypeRuleBasedOnTypeFromValue returns a pointer to a valid MembershipTypeRuleBasedOnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMembershipTypeRuleBasedOnTypeFromValue(v string) (*MembershipTypeRuleBasedOnType, error) {
	ev := MembershipTypeRuleBasedOnType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MembershipTypeRuleBasedOnType: valid values are %v", v, AllowedMembershipTypeRuleBasedOnTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MembershipTypeRuleBasedOnType) IsValid() bool {
	for _, existing := range AllowedMembershipTypeRuleBasedOnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to membershipTypeRuleBasedOnType value
func (v MembershipTypeRuleBasedOnType) Ptr() *MembershipTypeRuleBasedOnType {
	return &v
}

type NullableMembershipTypeRuleBasedOnType struct {
	value *MembershipTypeRuleBasedOnType
	isSet bool
}

func (v NullableMembershipTypeRuleBasedOnType) Get() *MembershipTypeRuleBasedOnType {
	return v.value
}

func (v *NullableMembershipTypeRuleBasedOnType) Set(val *MembershipTypeRuleBasedOnType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipTypeRuleBasedOnType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipTypeRuleBasedOnType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipTypeRuleBasedOnType(val *MembershipTypeRuleBasedOnType) *NullableMembershipTypeRuleBasedOnType {
	return &NullableMembershipTypeRuleBasedOnType{value: val, isSet: true}
}

func (v NullableMembershipTypeRuleBasedOnType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipTypeRuleBasedOnType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

