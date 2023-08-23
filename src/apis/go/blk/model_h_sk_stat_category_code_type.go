/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// HSKStatCategoryCodeType Supported housekeeping statistical category codes.
type HSKStatCategoryCodeType string

// List of hSKStatCategoryCodeType
const (
	HOTEL_CODE HSKStatCategoryCodeType = "HotelCode"
)

// All allowed values of HSKStatCategoryCodeType enum
var AllowedHSKStatCategoryCodeTypeEnumValues = []HSKStatCategoryCodeType{
	"HotelCode",
}

func (v *HSKStatCategoryCodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HSKStatCategoryCodeType(value)
	for _, existing := range AllowedHSKStatCategoryCodeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HSKStatCategoryCodeType", value)
}

// NewHSKStatCategoryCodeTypeFromValue returns a pointer to a valid HSKStatCategoryCodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHSKStatCategoryCodeTypeFromValue(v string) (*HSKStatCategoryCodeType, error) {
	ev := HSKStatCategoryCodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HSKStatCategoryCodeType: valid values are %v", v, AllowedHSKStatCategoryCodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HSKStatCategoryCodeType) IsValid() bool {
	for _, existing := range AllowedHSKStatCategoryCodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to hSKStatCategoryCodeType value
func (v HSKStatCategoryCodeType) Ptr() *HSKStatCategoryCodeType {
	return &v
}

type NullableHSKStatCategoryCodeType struct {
	value *HSKStatCategoryCodeType
	isSet bool
}

func (v NullableHSKStatCategoryCodeType) Get() *HSKStatCategoryCodeType {
	return v.value
}

func (v *NullableHSKStatCategoryCodeType) Set(val *HSKStatCategoryCodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableHSKStatCategoryCodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableHSKStatCategoryCodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHSKStatCategoryCodeType(val *HSKStatCategoryCodeType) *NullableHSKStatCategoryCodeType {
	return &NullableHSKStatCategoryCodeType{value: val, isSet: true}
}

func (v NullableHSKStatCategoryCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHSKStatCategoryCodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

