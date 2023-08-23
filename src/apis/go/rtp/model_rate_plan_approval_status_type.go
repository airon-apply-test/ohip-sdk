/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RatePlanApprovalStatusType Indicates that the Rate Plan is approved for selling.
type RatePlanApprovalStatusType string

// List of ratePlanApprovalStatusType
const (
	NEW_UNAPPROVED RatePlanApprovalStatusType = "NewUnapproved"
	CHANGED_UNAPPROVED RatePlanApprovalStatusType = "ChangedUnapproved"
	REJECTED RatePlanApprovalStatusType = "Rejected"
	APPROVED RatePlanApprovalStatusType = "Approved"
)

// All allowed values of RatePlanApprovalStatusType enum
var AllowedRatePlanApprovalStatusTypeEnumValues = []RatePlanApprovalStatusType{
	"NewUnapproved",
	"ChangedUnapproved",
	"Rejected",
	"Approved",
}

func (v *RatePlanApprovalStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RatePlanApprovalStatusType(value)
	for _, existing := range AllowedRatePlanApprovalStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RatePlanApprovalStatusType", value)
}

// NewRatePlanApprovalStatusTypeFromValue returns a pointer to a valid RatePlanApprovalStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRatePlanApprovalStatusTypeFromValue(v string) (*RatePlanApprovalStatusType, error) {
	ev := RatePlanApprovalStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RatePlanApprovalStatusType: valid values are %v", v, AllowedRatePlanApprovalStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RatePlanApprovalStatusType) IsValid() bool {
	for _, existing := range AllowedRatePlanApprovalStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ratePlanApprovalStatusType value
func (v RatePlanApprovalStatusType) Ptr() *RatePlanApprovalStatusType {
	return &v
}

type NullableRatePlanApprovalStatusType struct {
	value *RatePlanApprovalStatusType
	isSet bool
}

func (v NullableRatePlanApprovalStatusType) Get() *RatePlanApprovalStatusType {
	return v.value
}

func (v *NullableRatePlanApprovalStatusType) Set(val *RatePlanApprovalStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePlanApprovalStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePlanApprovalStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePlanApprovalStatusType(val *RatePlanApprovalStatusType) *NullableRatePlanApprovalStatusType {
	return &NullableRatePlanApprovalStatusType{value: val, isSet: true}
}

func (v NullableRatePlanApprovalStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePlanApprovalStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

