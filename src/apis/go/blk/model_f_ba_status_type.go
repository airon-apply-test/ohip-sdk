/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
	"fmt"
)

// FBAStatusType Flexible benefit award has been reimbursed.
type FBAStatusType string

// List of fBAStatusType
const (
	FBASTATUSTYPE_ORDERED FBAStatusType = "Ordered"
	FBASTATUSTYPE_ATTACHED FBAStatusType = "Attached"
	FBASTATUSTYPE_POSTED FBAStatusType = "Posted"
	FBASTATUSTYPE_SETTLED FBAStatusType = "Settled"
	FBASTATUSTYPE_REIMBURSED FBAStatusType = "Reimbursed"
)

// All allowed values of FBAStatusType enum
var AllowedFBAStatusTypeEnumValues = []FBAStatusType{
	"Ordered",
	"Attached",
	"Posted",
	"Settled",
	"Reimbursed",
}

func (v *FBAStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FBAStatusType(value)
	for _, existing := range AllowedFBAStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FBAStatusType", value)
}

// NewFBAStatusTypeFromValue returns a pointer to a valid FBAStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFBAStatusTypeFromValue(v string) (*FBAStatusType, error) {
	ev := FBAStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FBAStatusType: valid values are %v", v, AllowedFBAStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FBAStatusType) IsValid() bool {
	for _, existing := range AllowedFBAStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to fBAStatusType value
func (v FBAStatusType) Ptr() *FBAStatusType {
	return &v
}

type NullableFBAStatusType struct {
	value *FBAStatusType
	isSet bool
}

func (v NullableFBAStatusType) Get() *FBAStatusType {
	return v.value
}

func (v *NullableFBAStatusType) Set(val *FBAStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableFBAStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableFBAStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFBAStatusType(val *FBAStatusType) *NullableFBAStatusType {
	return &NullableFBAStatusType{value: val, isSet: true}
}

func (v NullableFBAStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFBAStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

