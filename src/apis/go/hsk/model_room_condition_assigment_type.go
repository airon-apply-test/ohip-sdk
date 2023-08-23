/*
OPERA Cloud Housekeeping Service API

APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RoomConditionAssigmentType Simple type for Room Condition assignment type. Valid values are Available and NotAvailable Only.
type RoomConditionAssigmentType string

// List of roomConditionAssigmentType
const (
	AVAILABLE RoomConditionAssigmentType = "Available"
	NOT_AVAILABLE RoomConditionAssigmentType = "NotAvailable"
)

// All allowed values of RoomConditionAssigmentType enum
var AllowedRoomConditionAssigmentTypeEnumValues = []RoomConditionAssigmentType{
	"Available",
	"NotAvailable",
}

func (v *RoomConditionAssigmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoomConditionAssigmentType(value)
	for _, existing := range AllowedRoomConditionAssigmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoomConditionAssigmentType", value)
}

// NewRoomConditionAssigmentTypeFromValue returns a pointer to a valid RoomConditionAssigmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoomConditionAssigmentTypeFromValue(v string) (*RoomConditionAssigmentType, error) {
	ev := RoomConditionAssigmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoomConditionAssigmentType: valid values are %v", v, AllowedRoomConditionAssigmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoomConditionAssigmentType) IsValid() bool {
	for _, existing := range AllowedRoomConditionAssigmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to roomConditionAssigmentType value
func (v RoomConditionAssigmentType) Ptr() *RoomConditionAssigmentType {
	return &v
}

type NullableRoomConditionAssigmentType struct {
	value *RoomConditionAssigmentType
	isSet bool
}

func (v NullableRoomConditionAssigmentType) Get() *RoomConditionAssigmentType {
	return v.value
}

func (v *NullableRoomConditionAssigmentType) Set(val *RoomConditionAssigmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomConditionAssigmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomConditionAssigmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomConditionAssigmentType(val *RoomConditionAssigmentType) *NullableRoomConditionAssigmentType {
	return &NullableRoomConditionAssigmentType{value: val, isSet: true}
}

func (v NullableRoomConditionAssigmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomConditionAssigmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

