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

// StatusColorType Status code color configuration type.
type StatusColorType string

// List of statusColorType
const (
	WHITE StatusColorType = "White"
	RED StatusColorType = "Red"
	BLUE StatusColorType = "Blue"
	CYAN StatusColorType = "Cyan"
	MAGENTA StatusColorType = "Magenta"
	GREEN StatusColorType = "Green"
	BROWN StatusColorType = "Brown"
	BLACK StatusColorType = "Black"
	YELLOW StatusColorType = "Yellow"
	GRAY StatusColorType = "Gray"
)

// All allowed values of StatusColorType enum
var AllowedStatusColorTypeEnumValues = []StatusColorType{
	"White",
	"Red",
	"Blue",
	"Cyan",
	"Magenta",
	"Green",
	"Brown",
	"Black",
	"Yellow",
	"Gray",
}

func (v *StatusColorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusColorType(value)
	for _, existing := range AllowedStatusColorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusColorType", value)
}

// NewStatusColorTypeFromValue returns a pointer to a valid StatusColorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusColorTypeFromValue(v string) (*StatusColorType, error) {
	ev := StatusColorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusColorType: valid values are %v", v, AllowedStatusColorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusColorType) IsValid() bool {
	for _, existing := range AllowedStatusColorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to statusColorType value
func (v StatusColorType) Ptr() *StatusColorType {
	return &v
}

type NullableStatusColorType struct {
	value *StatusColorType
	isSet bool
}

func (v NullableStatusColorType) Get() *StatusColorType {
	return v.value
}

func (v *NullableStatusColorType) Set(val *StatusColorType) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusColorType) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusColorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusColorType(val *StatusColorType) *NullableStatusColorType {
	return &NullableStatusColorType{value: val, isSet: true}
}

func (v NullableStatusColorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusColorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

