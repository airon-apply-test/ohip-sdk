/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SystemType Defines values for System Type
type SystemType string

// List of systemType
const (
	DEV SystemType = "Dev"
	UAT SystemType = "Uat"
	STAGE SystemType = "Stage"
	PROD SystemType = "Prod"
)

// All allowed values of SystemType enum
var AllowedSystemTypeEnumValues = []SystemType{
	"Dev",
	"Uat",
	"Stage",
	"Prod",
}

func (v *SystemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SystemType(value)
	for _, existing := range AllowedSystemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SystemType", value)
}

// NewSystemTypeFromValue returns a pointer to a valid SystemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSystemTypeFromValue(v string) (*SystemType, error) {
	ev := SystemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SystemType: valid values are %v", v, AllowedSystemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SystemType) IsValid() bool {
	for _, existing := range AllowedSystemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to systemType value
func (v SystemType) Ptr() *SystemType {
	return &v
}

type NullableSystemType struct {
	value *SystemType
	isSet bool
}

func (v NullableSystemType) Get() *SystemType {
	return v.value
}

func (v *NullableSystemType) Set(val *SystemType) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemType) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemType(val *SystemType) *NullableSystemType {
	return &NullableSystemType{value: val, isSet: true}
}

func (v NullableSystemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

