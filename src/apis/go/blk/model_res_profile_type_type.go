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

// ResProfileTypeType the model 'ResProfileTypeType'
type ResProfileTypeType string

// List of resProfileTypeType
const (
	GUEST ResProfileTypeType = "Guest"
	COMPANY ResProfileTypeType = "Company"
	GROUP ResProfileTypeType = "Group"
	TRAVEL_AGENT ResProfileTypeType = "TravelAgent"
	SOURCE ResProfileTypeType = "Source"
	RESERVATION_CONTACT ResProfileTypeType = "ReservationContact"
	BILLING_CONTACT ResProfileTypeType = "BillingContact"
	ADDRESSEE ResProfileTypeType = "Addressee"
)

// All allowed values of ResProfileTypeType enum
var AllowedResProfileTypeTypeEnumValues = []ResProfileTypeType{
	"Guest",
	"Company",
	"Group",
	"TravelAgent",
	"Source",
	"ReservationContact",
	"BillingContact",
	"Addressee",
}

func (v *ResProfileTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResProfileTypeType(value)
	for _, existing := range AllowedResProfileTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResProfileTypeType", value)
}

// NewResProfileTypeTypeFromValue returns a pointer to a valid ResProfileTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResProfileTypeTypeFromValue(v string) (*ResProfileTypeType, error) {
	ev := ResProfileTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResProfileTypeType: valid values are %v", v, AllowedResProfileTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResProfileTypeType) IsValid() bool {
	for _, existing := range AllowedResProfileTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resProfileTypeType value
func (v ResProfileTypeType) Ptr() *ResProfileTypeType {
	return &v
}

type NullableResProfileTypeType struct {
	value *ResProfileTypeType
	isSet bool
}

func (v NullableResProfileTypeType) Get() *ResProfileTypeType {
	return v.value
}

func (v *NullableResProfileTypeType) Set(val *ResProfileTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableResProfileTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableResProfileTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResProfileTypeType(val *ResProfileTypeType) *NullableResProfileTypeType {
	return &NullableResProfileTypeType{value: val, isSet: true}
}

func (v NullableResProfileTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResProfileTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

