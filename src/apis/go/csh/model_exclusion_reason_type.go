/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
	"fmt"
)

// ExclusionReasonType Custom Charge has already been Processed for this date within the reservation stay.
type ExclusionReasonType string

// List of exclusionReasonType
const (
	EXCLUSIONREASONTYPE_NOT_CONFIGURED ExclusionReasonType = "NotConfigured"
	EXCLUSIONREASONTYPE_CUSTOM_CHARGE_PROCESSED ExclusionReasonType = "CustomChargeProcessed"
)

// All allowed values of ExclusionReasonType enum
var AllowedExclusionReasonTypeEnumValues = []ExclusionReasonType{
	"NotConfigured",
	"CustomChargeProcessed",
}

func (v *ExclusionReasonType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExclusionReasonType(value)
	for _, existing := range AllowedExclusionReasonTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExclusionReasonType", value)
}

// NewExclusionReasonTypeFromValue returns a pointer to a valid ExclusionReasonType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExclusionReasonTypeFromValue(v string) (*ExclusionReasonType, error) {
	ev := ExclusionReasonType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExclusionReasonType: valid values are %v", v, AllowedExclusionReasonTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExclusionReasonType) IsValid() bool {
	for _, existing := range AllowedExclusionReasonTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to exclusionReasonType value
func (v ExclusionReasonType) Ptr() *ExclusionReasonType {
	return &v
}

type NullableExclusionReasonType struct {
	value *ExclusionReasonType
	isSet bool
}

func (v NullableExclusionReasonType) Get() *ExclusionReasonType {
	return v.value
}

func (v *NullableExclusionReasonType) Set(val *ExclusionReasonType) {
	v.value = val
	v.isSet = true
}

func (v NullableExclusionReasonType) IsSet() bool {
	return v.isSet
}

func (v *NullableExclusionReasonType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExclusionReasonType(val *ExclusionReasonType) *NullableExclusionReasonType {
	return &NullableExclusionReasonType{value: val, isSet: true}
}

func (v NullableExclusionReasonType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExclusionReasonType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

