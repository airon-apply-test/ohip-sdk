/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
	"fmt"
)

// CardProcessingType the model 'CardProcessingType'
type CardProcessingType string

// List of cardProcessingType
const (
	CARDPROCESSINGTYPE_EFT CardProcessingType = "Eft"
	CARDPROCESSINGTYPE_MANUAL CardProcessingType = "Manual"
)

// All allowed values of CardProcessingType enum
var AllowedCardProcessingTypeEnumValues = []CardProcessingType{
	"Eft",
	"Manual",
}

func (v *CardProcessingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardProcessingType(value)
	for _, existing := range AllowedCardProcessingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardProcessingType", value)
}

// NewCardProcessingTypeFromValue returns a pointer to a valid CardProcessingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardProcessingTypeFromValue(v string) (*CardProcessingType, error) {
	ev := CardProcessingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardProcessingType: valid values are %v", v, AllowedCardProcessingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardProcessingType) IsValid() bool {
	for _, existing := range AllowedCardProcessingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cardProcessingType value
func (v CardProcessingType) Ptr() *CardProcessingType {
	return &v
}

type NullableCardProcessingType struct {
	value *CardProcessingType
	isSet bool
}

func (v NullableCardProcessingType) Get() *CardProcessingType {
	return v.value
}

func (v *NullableCardProcessingType) Set(val *CardProcessingType) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProcessingType) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProcessingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProcessingType(val *CardProcessingType) *NullableCardProcessingType {
	return &NullableCardProcessingType{value: val, isSet: true}
}

func (v NullableCardProcessingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProcessingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

