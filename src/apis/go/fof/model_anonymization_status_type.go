/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
	"fmt"
)

// AnonymizationStatusType Guest has been anonymized.
type AnonymizationStatusType string

// List of anonymizationStatusType
const (
	ANONYMIZATIONSTATUSTYPE_REQUESTED AnonymizationStatusType = "Requested"
	ANONYMIZATIONSTATUSTYPE_ANONYMIZED AnonymizationStatusType = "Anonymized"
)

// All allowed values of AnonymizationStatusType enum
var AllowedAnonymizationStatusTypeEnumValues = []AnonymizationStatusType{
	"Requested",
	"Anonymized",
}

func (v *AnonymizationStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnonymizationStatusType(value)
	for _, existing := range AllowedAnonymizationStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnonymizationStatusType", value)
}

// NewAnonymizationStatusTypeFromValue returns a pointer to a valid AnonymizationStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnonymizationStatusTypeFromValue(v string) (*AnonymizationStatusType, error) {
	ev := AnonymizationStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnonymizationStatusType: valid values are %v", v, AllowedAnonymizationStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnonymizationStatusType) IsValid() bool {
	for _, existing := range AllowedAnonymizationStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to anonymizationStatusType value
func (v AnonymizationStatusType) Ptr() *AnonymizationStatusType {
	return &v
}

type NullableAnonymizationStatusType struct {
	value *AnonymizationStatusType
	isSet bool
}

func (v NullableAnonymizationStatusType) Get() *AnonymizationStatusType {
	return v.value
}

func (v *NullableAnonymizationStatusType) Set(val *AnonymizationStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableAnonymizationStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableAnonymizationStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnonymizationStatusType(val *AnonymizationStatusType) *NullableAnonymizationStatusType {
	return &NullableAnonymizationStatusType{value: val, isSet: true}
}

func (v NullableAnonymizationStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnonymizationStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

