/*
Opera Cloud Rate Plan Asynchronous Service API

APIs catering to the Rate Plan asynchronous related functionality in a hotel.  This includes adding/updating daily rates&apos; pricing schedules and best available rates by day or length of stay. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtpasync

import (
	"encoding/json"
	"fmt"
)

// RateCodeLockStatusType Indicates that the rate code is locked by the central system and cannot be edited at the property level.
type RateCodeLockStatusType string

// List of rateCodeLockStatusType
const (
	RATECODELOCKSTATUSTYPE_UNLOCKED RateCodeLockStatusType = "Unlocked"
	RATECODELOCKSTATUSTYPE_EXTERNAL RateCodeLockStatusType = "External"
	RATECODELOCKSTATUSTYPE_PROPERTY RateCodeLockStatusType = "Property"
	RATECODELOCKSTATUSTYPE_CENTRAL RateCodeLockStatusType = "Central"
)

// All allowed values of RateCodeLockStatusType enum
var AllowedRateCodeLockStatusTypeEnumValues = []RateCodeLockStatusType{
	"Unlocked",
	"External",
	"Property",
	"Central",
}

func (v *RateCodeLockStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RateCodeLockStatusType(value)
	for _, existing := range AllowedRateCodeLockStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RateCodeLockStatusType", value)
}

// NewRateCodeLockStatusTypeFromValue returns a pointer to a valid RateCodeLockStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRateCodeLockStatusTypeFromValue(v string) (*RateCodeLockStatusType, error) {
	ev := RateCodeLockStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RateCodeLockStatusType: valid values are %v", v, AllowedRateCodeLockStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RateCodeLockStatusType) IsValid() bool {
	for _, existing := range AllowedRateCodeLockStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to rateCodeLockStatusType value
func (v RateCodeLockStatusType) Ptr() *RateCodeLockStatusType {
	return &v
}

type NullableRateCodeLockStatusType struct {
	value *RateCodeLockStatusType
	isSet bool
}

func (v NullableRateCodeLockStatusType) Get() *RateCodeLockStatusType {
	return v.value
}

func (v *NullableRateCodeLockStatusType) Set(val *RateCodeLockStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableRateCodeLockStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableRateCodeLockStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateCodeLockStatusType(val *RateCodeLockStatusType) *NullableRateCodeLockStatusType {
	return &NullableRateCodeLockStatusType{value: val, isSet: true}
}

func (v NullableRateCodeLockStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateCodeLockStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

