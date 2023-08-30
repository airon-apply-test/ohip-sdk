/*
OPERA Cloud Room Configuration API

APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rmcfg

import (
	"encoding/json"
	"fmt"
)

// RatePlanRatingType Possible Rate plan rating values.
type RatePlanRatingType string

// List of ratePlanRatingType
const (
	RATEPLANRATINGTYPE_MOST_IMPORTANT RatePlanRatingType = "MostImportant"
	RATEPLANRATINGTYPE_IMPORTANT RatePlanRatingType = "Important"
	RATEPLANRATINGTYPE_SOMEWHAT_IMPORTANT RatePlanRatingType = "SomewhatImportant"
	RATEPLANRATINGTYPE_LESS_IMPORTANT RatePlanRatingType = "LessImportant"
	RATEPLANRATINGTYPE_LEAST_IMPORTANT RatePlanRatingType = "LeastImportant"
	RATEPLANRATINGTYPE_NO_VALUE_SET RatePlanRatingType = "NoValueSet"
)

// All allowed values of RatePlanRatingType enum
var AllowedRatePlanRatingTypeEnumValues = []RatePlanRatingType{
	"MostImportant",
	"Important",
	"SomewhatImportant",
	"LessImportant",
	"LeastImportant",
	"NoValueSet",
}

func (v *RatePlanRatingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RatePlanRatingType(value)
	for _, existing := range AllowedRatePlanRatingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RatePlanRatingType", value)
}

// NewRatePlanRatingTypeFromValue returns a pointer to a valid RatePlanRatingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRatePlanRatingTypeFromValue(v string) (*RatePlanRatingType, error) {
	ev := RatePlanRatingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RatePlanRatingType: valid values are %v", v, AllowedRatePlanRatingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RatePlanRatingType) IsValid() bool {
	for _, existing := range AllowedRatePlanRatingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ratePlanRatingType value
func (v RatePlanRatingType) Ptr() *RatePlanRatingType {
	return &v
}

type NullableRatePlanRatingType struct {
	value *RatePlanRatingType
	isSet bool
}

func (v NullableRatePlanRatingType) Get() *RatePlanRatingType {
	return v.value
}

func (v *NullableRatePlanRatingType) Set(val *RatePlanRatingType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePlanRatingType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePlanRatingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePlanRatingType(val *RatePlanRatingType) *NullableRatePlanRatingType {
	return &NullableRatePlanRatingType{value: val, isSet: true}
}

func (v NullableRatePlanRatingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePlanRatingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

