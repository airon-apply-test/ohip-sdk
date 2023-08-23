/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BaseType Simple type for base type, flat or differential.
type BaseType string

// List of baseType
const (
	FLAT BaseType = "Flat"
	DIFFERENTIAL BaseType = "Differential"
)

// All allowed values of BaseType enum
var AllowedBaseTypeEnumValues = []BaseType{
	"Flat",
	"Differential",
}

func (v *BaseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BaseType(value)
	for _, existing := range AllowedBaseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BaseType", value)
}

// NewBaseTypeFromValue returns a pointer to a valid BaseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBaseTypeFromValue(v string) (*BaseType, error) {
	ev := BaseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BaseType: valid values are %v", v, AllowedBaseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BaseType) IsValid() bool {
	for _, existing := range AllowedBaseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to baseType value
func (v BaseType) Ptr() *BaseType {
	return &v
}

type NullableBaseType struct {
	value *BaseType
	isSet bool
}

func (v NullableBaseType) Get() *BaseType {
	return v.value
}

func (v *NullableBaseType) Set(val *BaseType) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseType) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseType(val *BaseType) *NullableBaseType {
	return &NullableBaseType{value: val, isSet: true}
}

func (v NullableBaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

