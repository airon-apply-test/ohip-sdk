/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PersonNameTypeType Person's name in an external system.
type PersonNameTypeType string

// List of personNameTypeType
const (
	PRIMARY PersonNameTypeType = "Primary"
	ALTERNATE PersonNameTypeType = "Alternate"
	INCOGNITO PersonNameTypeType = "Incognito"
	EXTERNAL PersonNameTypeType = "External"
	PHONETIC PersonNameTypeType = "Phonetic"
)

// All allowed values of PersonNameTypeType enum
var AllowedPersonNameTypeTypeEnumValues = []PersonNameTypeType{
	"Primary",
	"Alternate",
	"Incognito",
	"External",
	"Phonetic",
}

func (v *PersonNameTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersonNameTypeType(value)
	for _, existing := range AllowedPersonNameTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersonNameTypeType", value)
}

// NewPersonNameTypeTypeFromValue returns a pointer to a valid PersonNameTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersonNameTypeTypeFromValue(v string) (*PersonNameTypeType, error) {
	ev := PersonNameTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersonNameTypeType: valid values are %v", v, AllowedPersonNameTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersonNameTypeType) IsValid() bool {
	for _, existing := range AllowedPersonNameTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to personNameTypeType value
func (v PersonNameTypeType) Ptr() *PersonNameTypeType {
	return &v
}

type NullablePersonNameTypeType struct {
	value *PersonNameTypeType
	isSet bool
}

func (v NullablePersonNameTypeType) Get() *PersonNameTypeType {
	return v.value
}

func (v *NullablePersonNameTypeType) Set(val *PersonNameTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonNameTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonNameTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonNameTypeType(val *PersonNameTypeType) *NullablePersonNameTypeType {
	return &NullablePersonNameTypeType{value: val, isSet: true}
}

func (v NullablePersonNameTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonNameTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

