/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// LetterTypeType Represents Inquiry Letter Type.
type LetterTypeType string

// List of letterTypeType
const (
	CONFIRMATION LetterTypeType = "Confirmation"
	CANCELLATION LetterTypeType = "Cancellation"
	BANQUET_EVENT_ORDER LetterTypeType = "BanquetEventOrder"
	CONTRACT LetterTypeType = "Contract"
	INQUIRY LetterTypeType = "Inquiry"
)

// All allowed values of LetterTypeType enum
var AllowedLetterTypeTypeEnumValues = []LetterTypeType{
	"Confirmation",
	"Cancellation",
	"BanquetEventOrder",
	"Contract",
	"Inquiry",
}

func (v *LetterTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LetterTypeType(value)
	for _, existing := range AllowedLetterTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LetterTypeType", value)
}

// NewLetterTypeTypeFromValue returns a pointer to a valid LetterTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLetterTypeTypeFromValue(v string) (*LetterTypeType, error) {
	ev := LetterTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LetterTypeType: valid values are %v", v, AllowedLetterTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LetterTypeType) IsValid() bool {
	for _, existing := range AllowedLetterTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to letterTypeType value
func (v LetterTypeType) Ptr() *LetterTypeType {
	return &v
}

type NullableLetterTypeType struct {
	value *LetterTypeType
	isSet bool
}

func (v NullableLetterTypeType) Get() *LetterTypeType {
	return v.value
}

func (v *NullableLetterTypeType) Set(val *LetterTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableLetterTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableLetterTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLetterTypeType(val *LetterTypeType) *NullableLetterTypeType {
	return &NullableLetterTypeType{value: val, isSet: true}
}

func (v NullableLetterTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLetterTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

