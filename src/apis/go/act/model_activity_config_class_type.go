/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package act

import (
	"encoding/json"
	"fmt"
)

// ActivityConfigClassType A collection of supported list of activity classes.
type ActivityConfigClassType string

// List of activityConfigClassType
const (
	ACTIVITYCONFIGCLASSTYPE_APPOINTMENT ActivityConfigClassType = "Appointment"
	ACTIVITYCONFIGCLASSTYPE_TODO ActivityConfigClassType = "Todo"
)

// All allowed values of ActivityConfigClassType enum
var AllowedActivityConfigClassTypeEnumValues = []ActivityConfigClassType{
	"Appointment",
	"Todo",
}

func (v *ActivityConfigClassType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActivityConfigClassType(value)
	for _, existing := range AllowedActivityConfigClassTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActivityConfigClassType", value)
}

// NewActivityConfigClassTypeFromValue returns a pointer to a valid ActivityConfigClassType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActivityConfigClassTypeFromValue(v string) (*ActivityConfigClassType, error) {
	ev := ActivityConfigClassType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActivityConfigClassType: valid values are %v", v, AllowedActivityConfigClassTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActivityConfigClassType) IsValid() bool {
	for _, existing := range AllowedActivityConfigClassTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to activityConfigClassType value
func (v ActivityConfigClassType) Ptr() *ActivityConfigClassType {
	return &v
}

type NullableActivityConfigClassType struct {
	value *ActivityConfigClassType
	isSet bool
}

func (v NullableActivityConfigClassType) Get() *ActivityConfigClassType {
	return v.value
}

func (v *NullableActivityConfigClassType) Set(val *ActivityConfigClassType) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityConfigClassType) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityConfigClassType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityConfigClassType(val *ActivityConfigClassType) *NullableActivityConfigClassType {
	return &NullableActivityConfigClassType{value: val, isSet: true}
}

func (v NullableActivityConfigClassType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityConfigClassType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

