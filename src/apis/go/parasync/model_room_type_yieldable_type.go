/*
OPERA Cloud Price Availability Rate Async API

APIs to cater for Price and Rate Availability Asynchronous functionality in OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RoomTypeYieldableType Simple type for Yieldable or Non Yieldable
type RoomTypeYieldableType string

// List of roomTypeYieldableType
const (
	YIELDABLE RoomTypeYieldableType = "Yieldable"
	NON_YIELDABLE RoomTypeYieldableType = "NonYieldable"
)

// All allowed values of RoomTypeYieldableType enum
var AllowedRoomTypeYieldableTypeEnumValues = []RoomTypeYieldableType{
	"Yieldable",
	"NonYieldable",
}

func (v *RoomTypeYieldableType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoomTypeYieldableType(value)
	for _, existing := range AllowedRoomTypeYieldableTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoomTypeYieldableType", value)
}

// NewRoomTypeYieldableTypeFromValue returns a pointer to a valid RoomTypeYieldableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoomTypeYieldableTypeFromValue(v string) (*RoomTypeYieldableType, error) {
	ev := RoomTypeYieldableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoomTypeYieldableType: valid values are %v", v, AllowedRoomTypeYieldableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoomTypeYieldableType) IsValid() bool {
	for _, existing := range AllowedRoomTypeYieldableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to roomTypeYieldableType value
func (v RoomTypeYieldableType) Ptr() *RoomTypeYieldableType {
	return &v
}

type NullableRoomTypeYieldableType struct {
	value *RoomTypeYieldableType
	isSet bool
}

func (v NullableRoomTypeYieldableType) Get() *RoomTypeYieldableType {
	return v.value
}

func (v *NullableRoomTypeYieldableType) Set(val *RoomTypeYieldableType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomTypeYieldableType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomTypeYieldableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomTypeYieldableType(val *RoomTypeYieldableType) *NullableRoomTypeYieldableType {
	return &NullableRoomTypeYieldableType{value: val, isSet: true}
}

func (v NullableRoomTypeYieldableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomTypeYieldableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

