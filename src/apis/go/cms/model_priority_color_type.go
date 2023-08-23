/*
OPERA Cloud API for Customer Management Service

This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PriorityColorType Priority code color configuration type.
type PriorityColorType string

// List of priorityColorType
const (
	RED PriorityColorType = "Red"
	DARK_RED PriorityColorType = "DarkRed"
	BLUE PriorityColorType = "Blue"
	DARK_BLUE PriorityColorType = "DarkBlue"
	CYAN PriorityColorType = "Cyan"
	DARK_CYAN PriorityColorType = "DarkCyan"
	MAGENTA PriorityColorType = "Magenta"
	DARK_MAGENTA PriorityColorType = "DarkMagenta"
	GREEN PriorityColorType = "Green"
	DARK_GREEN PriorityColorType = "DarkGreen"
	BLACK PriorityColorType = "Black"
	YELLOW PriorityColorType = "Yellow"
	DARK_YELLOW PriorityColorType = "DarkYellow"
)

// All allowed values of PriorityColorType enum
var AllowedPriorityColorTypeEnumValues = []PriorityColorType{
	"Red",
	"DarkRed",
	"Blue",
	"DarkBlue",
	"Cyan",
	"DarkCyan",
	"Magenta",
	"DarkMagenta",
	"Green",
	"DarkGreen",
	"Black",
	"Yellow",
	"DarkYellow",
}

func (v *PriorityColorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriorityColorType(value)
	for _, existing := range AllowedPriorityColorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriorityColorType", value)
}

// NewPriorityColorTypeFromValue returns a pointer to a valid PriorityColorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriorityColorTypeFromValue(v string) (*PriorityColorType, error) {
	ev := PriorityColorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriorityColorType: valid values are %v", v, AllowedPriorityColorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriorityColorType) IsValid() bool {
	for _, existing := range AllowedPriorityColorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to priorityColorType value
func (v PriorityColorType) Ptr() *PriorityColorType {
	return &v
}

type NullablePriorityColorType struct {
	value *PriorityColorType
	isSet bool
}

func (v NullablePriorityColorType) Get() *PriorityColorType {
	return v.value
}

func (v *NullablePriorityColorType) Set(val *PriorityColorType) {
	v.value = val
	v.isSet = true
}

func (v NullablePriorityColorType) IsSet() bool {
	return v.isSet
}

func (v *NullablePriorityColorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriorityColorType(val *PriorityColorType) *NullablePriorityColorType {
	return &NullablePriorityColorType{value: val, isSet: true}
}

func (v NullablePriorityColorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriorityColorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

