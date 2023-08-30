/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtp

import (
	"encoding/json"
	"fmt"
)

// RatePlanYieldableType Simple type for valid values for Yieldable element for the Rate plan code.
type RatePlanYieldableType string

// List of ratePlanYieldableType
const (
	RATEPLANYIELDABLETYPE_YIELDABLE RatePlanYieldableType = "Yieldable"
	RATEPLANYIELDABLETYPE_NON_YIELDABLE RatePlanYieldableType = "NonYieldable"
	RATEPLANYIELDABLETYPE_STAY_PATTERN RatePlanYieldableType = "StayPattern"
)

// All allowed values of RatePlanYieldableType enum
var AllowedRatePlanYieldableTypeEnumValues = []RatePlanYieldableType{
	"Yieldable",
	"NonYieldable",
	"StayPattern",
}

func (v *RatePlanYieldableType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RatePlanYieldableType(value)
	for _, existing := range AllowedRatePlanYieldableTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RatePlanYieldableType", value)
}

// NewRatePlanYieldableTypeFromValue returns a pointer to a valid RatePlanYieldableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRatePlanYieldableTypeFromValue(v string) (*RatePlanYieldableType, error) {
	ev := RatePlanYieldableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RatePlanYieldableType: valid values are %v", v, AllowedRatePlanYieldableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RatePlanYieldableType) IsValid() bool {
	for _, existing := range AllowedRatePlanYieldableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ratePlanYieldableType value
func (v RatePlanYieldableType) Ptr() *RatePlanYieldableType {
	return &v
}

type NullableRatePlanYieldableType struct {
	value *RatePlanYieldableType
	isSet bool
}

func (v NullableRatePlanYieldableType) Get() *RatePlanYieldableType {
	return v.value
}

func (v *NullableRatePlanYieldableType) Set(val *RatePlanYieldableType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePlanYieldableType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePlanYieldableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePlanYieldableType(val *RatePlanYieldableType) *NullableRatePlanYieldableType {
	return &NullableRatePlanYieldableType{value: val, isSet: true}
}

func (v NullableRatePlanYieldableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePlanYieldableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

