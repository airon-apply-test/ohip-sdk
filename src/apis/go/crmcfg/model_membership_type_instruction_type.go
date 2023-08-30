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

// MembershipTypeInstructionType Membership levels.
type MembershipTypeInstructionType string

// List of membershipTypeInstructionType
const (
	MEMBERSHIPTYPEINSTRUCTIONTYPE_PRIMARY_DETAILS MembershipTypeInstructionType = "PrimaryDetails"
	MEMBERSHIPTYPEINSTRUCTIONTYPE_CARD_NUMBER_DETAILS MembershipTypeInstructionType = "CardNumberDetails"
	MEMBERSHIPTYPEINSTRUCTIONTYPE_POINTS_DETAILS MembershipTypeInstructionType = "PointsDetails"
	MEMBERSHIPTYPEINSTRUCTIONTYPE_EXCEPTION_CRITERIA_DETAILS MembershipTypeInstructionType = "ExceptionCriteriaDetails"
	MEMBERSHIPTYPEINSTRUCTIONTYPE_ADDITIONAL_DETAILS MembershipTypeInstructionType = "AdditionalDetails"
	MEMBERSHIPTYPEINSTRUCTIONTYPE_ENROLLMENT_DETAILS MembershipTypeInstructionType = "EnrollmentDetails"
	MEMBERSHIPTYPEINSTRUCTIONTYPE_LEVELS MembershipTypeInstructionType = "Levels"
)

// All allowed values of MembershipTypeInstructionType enum
var AllowedMembershipTypeInstructionTypeEnumValues = []MembershipTypeInstructionType{
	"PrimaryDetails",
	"CardNumberDetails",
	"PointsDetails",
	"ExceptionCriteriaDetails",
	"AdditionalDetails",
	"EnrollmentDetails",
	"Levels",
}

func (v *MembershipTypeInstructionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MembershipTypeInstructionType(value)
	for _, existing := range AllowedMembershipTypeInstructionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MembershipTypeInstructionType", value)
}

// NewMembershipTypeInstructionTypeFromValue returns a pointer to a valid MembershipTypeInstructionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMembershipTypeInstructionTypeFromValue(v string) (*MembershipTypeInstructionType, error) {
	ev := MembershipTypeInstructionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MembershipTypeInstructionType: valid values are %v", v, AllowedMembershipTypeInstructionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MembershipTypeInstructionType) IsValid() bool {
	for _, existing := range AllowedMembershipTypeInstructionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to membershipTypeInstructionType value
func (v MembershipTypeInstructionType) Ptr() *MembershipTypeInstructionType {
	return &v
}

type NullableMembershipTypeInstructionType struct {
	value *MembershipTypeInstructionType
	isSet bool
}

func (v NullableMembershipTypeInstructionType) Get() *MembershipTypeInstructionType {
	return v.value
}

func (v *NullableMembershipTypeInstructionType) Set(val *MembershipTypeInstructionType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipTypeInstructionType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipTypeInstructionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipTypeInstructionType(val *MembershipTypeInstructionType) *NullableMembershipTypeInstructionType {
	return &NullableMembershipTypeInstructionType{value: val, isSet: true}
}

func (v NullableMembershipTypeInstructionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipTypeInstructionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

