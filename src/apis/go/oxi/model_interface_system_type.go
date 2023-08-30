/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
	"fmt"
)

// InterfaceSystemType Defines values for System Type of the interface
type InterfaceSystemType string

// List of interfaceSystemType
const (
	INTERFACESYSTEMTYPE_OXI InterfaceSystemType = "Oxi"
	INTERFACESYSTEMTYPE_OXI_HUB InterfaceSystemType = "OxiHub"
	INTERFACESYSTEMTYPE_OG_SPA InterfaceSystemType = "OgSpa"
	INTERFACESYSTEMTYPE_OG_GAMING InterfaceSystemType = "OgGaming"
	INTERFACESYSTEMTYPE_BEXML InterfaceSystemType = "Bexml"
	INTERFACESYSTEMTYPE_IMPORT InterfaceSystemType = "Import"
	INTERFACESYSTEMTYPE_WEB_SERVICES InterfaceSystemType = "WebServices"
)

// All allowed values of InterfaceSystemType enum
var AllowedInterfaceSystemTypeEnumValues = []InterfaceSystemType{
	"Oxi",
	"OxiHub",
	"OgSpa",
	"OgGaming",
	"Bexml",
	"Import",
	"WebServices",
}

func (v *InterfaceSystemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceSystemType(value)
	for _, existing := range AllowedInterfaceSystemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceSystemType", value)
}

// NewInterfaceSystemTypeFromValue returns a pointer to a valid InterfaceSystemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceSystemTypeFromValue(v string) (*InterfaceSystemType, error) {
	ev := InterfaceSystemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceSystemType: valid values are %v", v, AllowedInterfaceSystemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceSystemType) IsValid() bool {
	for _, existing := range AllowedInterfaceSystemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to interfaceSystemType value
func (v InterfaceSystemType) Ptr() *InterfaceSystemType {
	return &v
}

type NullableInterfaceSystemType struct {
	value *InterfaceSystemType
	isSet bool
}

func (v NullableInterfaceSystemType) Get() *InterfaceSystemType {
	return v.value
}

func (v *NullableInterfaceSystemType) Set(val *InterfaceSystemType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceSystemType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceSystemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceSystemType(val *InterfaceSystemType) *NullableInterfaceSystemType {
	return &NullableInterfaceSystemType{value: val, isSet: true}
}

func (v NullableInterfaceSystemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceSystemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

