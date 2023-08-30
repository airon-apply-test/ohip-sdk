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

// MembershipActionType Type of action that the user wants to attach membership information to the reservation .
type MembershipActionType string

// List of membershipActionType
const (
	MEMBERSHIPACTIONTYPE_AUTO_POPULATE MembershipActionType = "AutoPopulate"
	MEMBERSHIPACTIONTYPE_PROMPT_TO_POPULATE MembershipActionType = "PromptToPopulate"
	MEMBERSHIPACTIONTYPE_ALWAYS_PROMPT MembershipActionType = "AlwaysPrompt"
	MEMBERSHIPACTIONTYPE_NO_ACTION MembershipActionType = "NoAction"
)

// All allowed values of MembershipActionType enum
var AllowedMembershipActionTypeEnumValues = []MembershipActionType{
	"AutoPopulate",
	"PromptToPopulate",
	"AlwaysPrompt",
	"NoAction",
}

func (v *MembershipActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MembershipActionType(value)
	for _, existing := range AllowedMembershipActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MembershipActionType", value)
}

// NewMembershipActionTypeFromValue returns a pointer to a valid MembershipActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMembershipActionTypeFromValue(v string) (*MembershipActionType, error) {
	ev := MembershipActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MembershipActionType: valid values are %v", v, AllowedMembershipActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MembershipActionType) IsValid() bool {
	for _, existing := range AllowedMembershipActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to membershipActionType value
func (v MembershipActionType) Ptr() *MembershipActionType {
	return &v
}

type NullableMembershipActionType struct {
	value *MembershipActionType
	isSet bool
}

func (v NullableMembershipActionType) Get() *MembershipActionType {
	return v.value
}

func (v *NullableMembershipActionType) Set(val *MembershipActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipActionType(val *MembershipActionType) *NullableMembershipActionType {
	return &NullableMembershipActionType{value: val, isSet: true}
}

func (v NullableMembershipActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

