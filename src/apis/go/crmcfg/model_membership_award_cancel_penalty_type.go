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
	"fmt"
)

// MembershipAwardCancelPenaltyType Penalty charge is in percentage.
type MembershipAwardCancelPenaltyType string

// List of membershipAwardCancelPenaltyType
const (
	POINTS MembershipAwardCancelPenaltyType = "Points"
	PERCENT MembershipAwardCancelPenaltyType = "Percent"
)

// All allowed values of MembershipAwardCancelPenaltyType enum
var AllowedMembershipAwardCancelPenaltyTypeEnumValues = []MembershipAwardCancelPenaltyType{
	"Points",
	"Percent",
}

func (v *MembershipAwardCancelPenaltyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MembershipAwardCancelPenaltyType(value)
	for _, existing := range AllowedMembershipAwardCancelPenaltyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MembershipAwardCancelPenaltyType", value)
}

// NewMembershipAwardCancelPenaltyTypeFromValue returns a pointer to a valid MembershipAwardCancelPenaltyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMembershipAwardCancelPenaltyTypeFromValue(v string) (*MembershipAwardCancelPenaltyType, error) {
	ev := MembershipAwardCancelPenaltyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MembershipAwardCancelPenaltyType: valid values are %v", v, AllowedMembershipAwardCancelPenaltyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MembershipAwardCancelPenaltyType) IsValid() bool {
	for _, existing := range AllowedMembershipAwardCancelPenaltyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to membershipAwardCancelPenaltyType value
func (v MembershipAwardCancelPenaltyType) Ptr() *MembershipAwardCancelPenaltyType {
	return &v
}

type NullableMembershipAwardCancelPenaltyType struct {
	value *MembershipAwardCancelPenaltyType
	isSet bool
}

func (v NullableMembershipAwardCancelPenaltyType) Get() *MembershipAwardCancelPenaltyType {
	return v.value
}

func (v *NullableMembershipAwardCancelPenaltyType) Set(val *MembershipAwardCancelPenaltyType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipAwardCancelPenaltyType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipAwardCancelPenaltyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipAwardCancelPenaltyType(val *MembershipAwardCancelPenaltyType) *NullableMembershipAwardCancelPenaltyType {
	return &NullableMembershipAwardCancelPenaltyType{value: val, isSet: true}
}

func (v NullableMembershipAwardCancelPenaltyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipAwardCancelPenaltyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

