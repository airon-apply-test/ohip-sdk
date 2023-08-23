/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ProfileStatusType the model 'ProfileStatusType'
type ProfileStatusType string

// List of profileStatusType
const (
	ACTIVE ProfileStatusType = "Active"
	INACTIVE ProfileStatusType = "Inactive"
)

// All allowed values of ProfileStatusType enum
var AllowedProfileStatusTypeEnumValues = []ProfileStatusType{
	"Active",
	"Inactive",
}

func (v *ProfileStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProfileStatusType(value)
	for _, existing := range AllowedProfileStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProfileStatusType", value)
}

// NewProfileStatusTypeFromValue returns a pointer to a valid ProfileStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileStatusTypeFromValue(v string) (*ProfileStatusType, error) {
	ev := ProfileStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProfileStatusType: valid values are %v", v, AllowedProfileStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileStatusType) IsValid() bool {
	for _, existing := range AllowedProfileStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to profileStatusType value
func (v ProfileStatusType) Ptr() *ProfileStatusType {
	return &v
}

type NullableProfileStatusType struct {
	value *ProfileStatusType
	isSet bool
}

func (v NullableProfileStatusType) Get() *ProfileStatusType {
	return v.value
}

func (v *NullableProfileStatusType) Set(val *ProfileStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileStatusType(val *ProfileStatusType) *NullableProfileStatusType {
	return &NullableProfileStatusType{value: val, isSet: true}
}

func (v NullableProfileStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

