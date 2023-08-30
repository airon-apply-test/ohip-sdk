/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmcfg

import (
	"encoding/json"
	"fmt"
)

// CommonMasterColorType Simple type enumeration for Common Master Colors.
type CommonMasterColorType string

// List of commonMasterColorType
const (
	COMMONMASTERCOLORTYPE_NONE CommonMasterColorType = "None"
	COMMONMASTERCOLORTYPE_BLUE CommonMasterColorType = "Blue"
	COMMONMASTERCOLORTYPE_RED CommonMasterColorType = "Red"
	COMMONMASTERCOLORTYPE_CYAN CommonMasterColorType = "Cyan"
	COMMONMASTERCOLORTYPE_GREEN CommonMasterColorType = "Green"
	COMMONMASTERCOLORTYPE_BLACK CommonMasterColorType = "Black"
	COMMONMASTERCOLORTYPE_WHITE CommonMasterColorType = "White"
	COMMONMASTERCOLORTYPE_YELLOW CommonMasterColorType = "Yellow"
	COMMONMASTERCOLORTYPE_GRAY CommonMasterColorType = "Gray"
)

// All allowed values of CommonMasterColorType enum
var AllowedCommonMasterColorTypeEnumValues = []CommonMasterColorType{
	"None",
	"Blue",
	"Red",
	"Cyan",
	"Green",
	"Black",
	"White",
	"Yellow",
	"Gray",
}

func (v *CommonMasterColorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommonMasterColorType(value)
	for _, existing := range AllowedCommonMasterColorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommonMasterColorType", value)
}

// NewCommonMasterColorTypeFromValue returns a pointer to a valid CommonMasterColorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommonMasterColorTypeFromValue(v string) (*CommonMasterColorType, error) {
	ev := CommonMasterColorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommonMasterColorType: valid values are %v", v, AllowedCommonMasterColorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommonMasterColorType) IsValid() bool {
	for _, existing := range AllowedCommonMasterColorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to commonMasterColorType value
func (v CommonMasterColorType) Ptr() *CommonMasterColorType {
	return &v
}

type NullableCommonMasterColorType struct {
	value *CommonMasterColorType
	isSet bool
}

func (v NullableCommonMasterColorType) Get() *CommonMasterColorType {
	return v.value
}

func (v *NullableCommonMasterColorType) Set(val *CommonMasterColorType) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonMasterColorType) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonMasterColorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonMasterColorType(val *CommonMasterColorType) *NullableCommonMasterColorType {
	return &NullableCommonMasterColorType{value: val, isSet: true}
}

func (v NullableCommonMasterColorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonMasterColorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

