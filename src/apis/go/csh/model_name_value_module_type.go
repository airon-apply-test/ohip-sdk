/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
	"fmt"
)

// NameValueModuleType the model 'NameValueModuleType'
type NameValueModuleType string

// List of nameValueModuleType
const (
	NAMEVALUEMODULETYPE_RESERVATION NameValueModuleType = "Reservation"
	NAMEVALUEMODULETYPE_FOLIO NameValueModuleType = "Folio"
	NAMEVALUEMODULETYPE_PROFILE NameValueModuleType = "Profile"
	NAMEVALUEMODULETYPE_FINTRIX NameValueModuleType = "Fintrix"
)

// All allowed values of NameValueModuleType enum
var AllowedNameValueModuleTypeEnumValues = []NameValueModuleType{
	"Reservation",
	"Folio",
	"Profile",
	"Fintrix",
}

func (v *NameValueModuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NameValueModuleType(value)
	for _, existing := range AllowedNameValueModuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NameValueModuleType", value)
}

// NewNameValueModuleTypeFromValue returns a pointer to a valid NameValueModuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNameValueModuleTypeFromValue(v string) (*NameValueModuleType, error) {
	ev := NameValueModuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NameValueModuleType: valid values are %v", v, AllowedNameValueModuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NameValueModuleType) IsValid() bool {
	for _, existing := range AllowedNameValueModuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to nameValueModuleType value
func (v NameValueModuleType) Ptr() *NameValueModuleType {
	return &v
}

type NullableNameValueModuleType struct {
	value *NameValueModuleType
	isSet bool
}

func (v NullableNameValueModuleType) Get() *NameValueModuleType {
	return v.value
}

func (v *NullableNameValueModuleType) Set(val *NameValueModuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableNameValueModuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableNameValueModuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameValueModuleType(val *NameValueModuleType) *NullableNameValueModuleType {
	return &NullableNameValueModuleType{value: val, isSet: true}
}

func (v NullableNameValueModuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameValueModuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

