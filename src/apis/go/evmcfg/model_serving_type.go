/*
OPERA Cloud Event Configuration API

This API caters for Event Configuration in OPERA Cloud. <br /><There are operations to post, update, fetch and delete codes such as item inventory, function spaces, menu items and many more.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evmcfg

import (
	"encoding/json"
	"fmt"
)

// ServingType Serving
type ServingType string

// List of servingType
const (
	SERVINGTYPE_PP ServingType = "Pp"
	SERVINGTYPE_PT ServingType = "Pt"
)

// All allowed values of ServingType enum
var AllowedServingTypeEnumValues = []ServingType{
	"Pp",
	"Pt",
}

func (v *ServingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServingType(value)
	for _, existing := range AllowedServingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServingType", value)
}

// NewServingTypeFromValue returns a pointer to a valid ServingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServingTypeFromValue(v string) (*ServingType, error) {
	ev := ServingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServingType: valid values are %v", v, AllowedServingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServingType) IsValid() bool {
	for _, existing := range AllowedServingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to servingType value
func (v ServingType) Ptr() *ServingType {
	return &v
}

type NullableServingType struct {
	value *ServingType
	isSet bool
}

func (v NullableServingType) Get() *ServingType {
	return v.value
}

func (v *NullableServingType) Set(val *ServingType) {
	v.value = val
	v.isSet = true
}

func (v NullableServingType) IsSet() bool {
	return v.isSet
}

func (v *NullableServingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServingType(val *ServingType) *NullableServingType {
	return &NullableServingType{value: val, isSet: true}
}

func (v NullableServingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

