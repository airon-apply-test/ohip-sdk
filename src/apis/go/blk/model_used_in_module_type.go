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

// UsedInModuleType Simple type for indicating the module for which user wants to create the SellMessage. Valid module types are Look To Book Sales,Reservations,Blocks,Function Diary and Group Rooms Control.
type UsedInModuleType string

// List of usedInModuleType
const (
	BLOCKS UsedInModuleType = "Blocks"
	RESERVATIONS UsedInModuleType = "Reservations"
	FUNCTION_DIARY UsedInModuleType = "FunctionDiary"
	LOOK_TO_BOOK_SALES UsedInModuleType = "LookToBookSales"
	GROUP_ROOMS_CONTROL UsedInModuleType = "GroupRoomsControl"
)

// All allowed values of UsedInModuleType enum
var AllowedUsedInModuleTypeEnumValues = []UsedInModuleType{
	"Blocks",
	"Reservations",
	"FunctionDiary",
	"LookToBookSales",
	"GroupRoomsControl",
}

func (v *UsedInModuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsedInModuleType(value)
	for _, existing := range AllowedUsedInModuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsedInModuleType", value)
}

// NewUsedInModuleTypeFromValue returns a pointer to a valid UsedInModuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsedInModuleTypeFromValue(v string) (*UsedInModuleType, error) {
	ev := UsedInModuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsedInModuleType: valid values are %v", v, AllowedUsedInModuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsedInModuleType) IsValid() bool {
	for _, existing := range AllowedUsedInModuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to usedInModuleType value
func (v UsedInModuleType) Ptr() *UsedInModuleType {
	return &v
}

type NullableUsedInModuleType struct {
	value *UsedInModuleType
	isSet bool
}

func (v NullableUsedInModuleType) Get() *UsedInModuleType {
	return v.value
}

func (v *NullableUsedInModuleType) Set(val *UsedInModuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableUsedInModuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableUsedInModuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsedInModuleType(val *UsedInModuleType) *NullableUsedInModuleType {
	return &NullableUsedInModuleType{value: val, isSet: true}
}

func (v NullableUsedInModuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsedInModuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

