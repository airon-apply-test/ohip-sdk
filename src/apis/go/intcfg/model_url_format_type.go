/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intcfg

import (
	"encoding/json"
	"fmt"
)

// UrlFormatType Url content format
type UrlFormatType string

// List of urlFormatType
const (
	URLFORMATTYPE_XML UrlFormatType = "Xml"
	URLFORMATTYPE_TEXT UrlFormatType = "Text"
)

// All allowed values of UrlFormatType enum
var AllowedUrlFormatTypeEnumValues = []UrlFormatType{
	"Xml",
	"Text",
}

func (v *UrlFormatType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UrlFormatType(value)
	for _, existing := range AllowedUrlFormatTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UrlFormatType", value)
}

// NewUrlFormatTypeFromValue returns a pointer to a valid UrlFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUrlFormatTypeFromValue(v string) (*UrlFormatType, error) {
	ev := UrlFormatType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UrlFormatType: valid values are %v", v, AllowedUrlFormatTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UrlFormatType) IsValid() bool {
	for _, existing := range AllowedUrlFormatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to urlFormatType value
func (v UrlFormatType) Ptr() *UrlFormatType {
	return &v
}

type NullableUrlFormatType struct {
	value *UrlFormatType
	isSet bool
}

func (v NullableUrlFormatType) Get() *UrlFormatType {
	return v.value
}

func (v *NullableUrlFormatType) Set(val *UrlFormatType) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlFormatType) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlFormatType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlFormatType(val *UrlFormatType) *NullableUrlFormatType {
	return &NullableUrlFormatType{value: val, isSet: true}
}

func (v NullableUrlFormatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlFormatType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

