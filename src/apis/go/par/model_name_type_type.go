/*
OPERA Cloud Price Availability Rate API

APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// NameTypeType the model 'NameTypeType'
type NameTypeType string

// List of nameTypeType
const (
	GUEST NameTypeType = "Guest"
	COMPANY NameTypeType = "Company"
	AGENT NameTypeType = "Agent"
	CONTACT NameTypeType = "Contact"
	SOURCE NameTypeType = "Source"
	GROUP NameTypeType = "Group"
	EMPLOYEE NameTypeType = "Employee"
	HOTEL NameTypeType = "Hotel"
	PURGE NameTypeType = "Purge"
)

// All allowed values of NameTypeType enum
var AllowedNameTypeTypeEnumValues = []NameTypeType{
	"Guest",
	"Company",
	"Agent",
	"Contact",
	"Source",
	"Group",
	"Employee",
	"Hotel",
	"Purge",
}

func (v *NameTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NameTypeType(value)
	for _, existing := range AllowedNameTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NameTypeType", value)
}

// NewNameTypeTypeFromValue returns a pointer to a valid NameTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNameTypeTypeFromValue(v string) (*NameTypeType, error) {
	ev := NameTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NameTypeType: valid values are %v", v, AllowedNameTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NameTypeType) IsValid() bool {
	for _, existing := range AllowedNameTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to nameTypeType value
func (v NameTypeType) Ptr() *NameTypeType {
	return &v
}

type NullableNameTypeType struct {
	value *NameTypeType
	isSet bool
}

func (v NullableNameTypeType) Get() *NameTypeType {
	return v.value
}

func (v *NullableNameTypeType) Set(val *NameTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableNameTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableNameTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameTypeType(val *NameTypeType) *NullableNameTypeType {
	return &NullableNameTypeType{value: val, isSet: true}
}

func (v NullableNameTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

