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

// ScopeType Minimum required and complimentary values are applicable Per Stay.
type ScopeType string

// List of scopeType
const (
	SCOPETYPE_PER_DAY ScopeType = "PerDay"
	SCOPETYPE_PER_STAY ScopeType = "PerStay"
)

// All allowed values of ScopeType enum
var AllowedScopeTypeEnumValues = []ScopeType{
	"PerDay",
	"PerStay",
}

func (v *ScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScopeType(value)
	for _, existing := range AllowedScopeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScopeType", value)
}

// NewScopeTypeFromValue returns a pointer to a valid ScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScopeTypeFromValue(v string) (*ScopeType, error) {
	ev := ScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScopeType: valid values are %v", v, AllowedScopeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScopeType) IsValid() bool {
	for _, existing := range AllowedScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to scopeType value
func (v ScopeType) Ptr() *ScopeType {
	return &v
}

type NullableScopeType struct {
	value *ScopeType
	isSet bool
}

func (v NullableScopeType) Get() *ScopeType {
	return v.value
}

func (v *NullableScopeType) Set(val *ScopeType) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeType) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeType(val *ScopeType) *NullableScopeType {
	return &NullableScopeType{value: val, isSet: true}
}

func (v NullableScopeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

