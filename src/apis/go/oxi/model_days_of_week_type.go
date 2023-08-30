/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
	"fmt"
)

// DaysOfWeekType Allowed values for the days of week type.
type DaysOfWeekType string

// List of daysOfWeekType
const (
	DAYSOFWEEKTYPE_MONDAY DaysOfWeekType = "Monday"
	DAYSOFWEEKTYPE_TUESDAY DaysOfWeekType = "Tuesday"
	DAYSOFWEEKTYPE_WEDNESDAY DaysOfWeekType = "Wednesday"
	DAYSOFWEEKTYPE_THURSDAY DaysOfWeekType = "Thursday"
	DAYSOFWEEKTYPE_FRIDAY DaysOfWeekType = "Friday"
	DAYSOFWEEKTYPE_SATURDAY DaysOfWeekType = "Saturday"
	DAYSOFWEEKTYPE_SUNDAY DaysOfWeekType = "Sunday"
)

// All allowed values of DaysOfWeekType enum
var AllowedDaysOfWeekTypeEnumValues = []DaysOfWeekType{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func (v *DaysOfWeekType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DaysOfWeekType(value)
	for _, existing := range AllowedDaysOfWeekTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DaysOfWeekType", value)
}

// NewDaysOfWeekTypeFromValue returns a pointer to a valid DaysOfWeekType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDaysOfWeekTypeFromValue(v string) (*DaysOfWeekType, error) {
	ev := DaysOfWeekType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DaysOfWeekType: valid values are %v", v, AllowedDaysOfWeekTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DaysOfWeekType) IsValid() bool {
	for _, existing := range AllowedDaysOfWeekTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to daysOfWeekType value
func (v DaysOfWeekType) Ptr() *DaysOfWeekType {
	return &v
}

type NullableDaysOfWeekType struct {
	value *DaysOfWeekType
	isSet bool
}

func (v NullableDaysOfWeekType) Get() *DaysOfWeekType {
	return v.value
}

func (v *NullableDaysOfWeekType) Set(val *DaysOfWeekType) {
	v.value = val
	v.isSet = true
}

func (v NullableDaysOfWeekType) IsSet() bool {
	return v.isSet
}

func (v *NullableDaysOfWeekType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaysOfWeekType(val *DaysOfWeekType) *NullableDaysOfWeekType {
	return &NullableDaysOfWeekType{value: val, isSet: true}
}

func (v NullableDaysOfWeekType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaysOfWeekType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

