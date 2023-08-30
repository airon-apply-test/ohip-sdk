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

// MembershipTypeRuleTransactionCodeType Determines the Transaction Code for Membership program rule.
type MembershipTypeRuleTransactionCodeType string

// List of membershipTypeRuleTransactionCodeType
const (
	MEMBERSHIPTYPERULETRANSACTIONCODETYPE_RESIDENT MembershipTypeRuleTransactionCodeType = "Resident"
	MEMBERSHIPTYPERULETRANSACTIONCODETYPE_NON_RESIDENT MembershipTypeRuleTransactionCodeType = "NonResident"
)

// All allowed values of MembershipTypeRuleTransactionCodeType enum
var AllowedMembershipTypeRuleTransactionCodeTypeEnumValues = []MembershipTypeRuleTransactionCodeType{
	"Resident",
	"NonResident",
}

func (v *MembershipTypeRuleTransactionCodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MembershipTypeRuleTransactionCodeType(value)
	for _, existing := range AllowedMembershipTypeRuleTransactionCodeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MembershipTypeRuleTransactionCodeType", value)
}

// NewMembershipTypeRuleTransactionCodeTypeFromValue returns a pointer to a valid MembershipTypeRuleTransactionCodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMembershipTypeRuleTransactionCodeTypeFromValue(v string) (*MembershipTypeRuleTransactionCodeType, error) {
	ev := MembershipTypeRuleTransactionCodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MembershipTypeRuleTransactionCodeType: valid values are %v", v, AllowedMembershipTypeRuleTransactionCodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MembershipTypeRuleTransactionCodeType) IsValid() bool {
	for _, existing := range AllowedMembershipTypeRuleTransactionCodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to membershipTypeRuleTransactionCodeType value
func (v MembershipTypeRuleTransactionCodeType) Ptr() *MembershipTypeRuleTransactionCodeType {
	return &v
}

type NullableMembershipTypeRuleTransactionCodeType struct {
	value *MembershipTypeRuleTransactionCodeType
	isSet bool
}

func (v NullableMembershipTypeRuleTransactionCodeType) Get() *MembershipTypeRuleTransactionCodeType {
	return v.value
}

func (v *NullableMembershipTypeRuleTransactionCodeType) Set(val *MembershipTypeRuleTransactionCodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipTypeRuleTransactionCodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipTypeRuleTransactionCodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipTypeRuleTransactionCodeType(val *MembershipTypeRuleTransactionCodeType) *NullableMembershipTypeRuleTransactionCodeType {
	return &NullableMembershipTypeRuleTransactionCodeType{value: val, isSet: true}
}

func (v NullableMembershipTypeRuleTransactionCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipTypeRuleTransactionCodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

