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

// LinkedReservationType Enumeration of the different Linked Reservation Types.
type LinkedReservationType string

// List of linkedReservationType
const (
	LINKED LinkedReservationType = "Linked"
	SHARED LinkedReservationType = "Shared"
	DEFAULT LinkedReservationType = "Default"
)

// All allowed values of LinkedReservationType enum
var AllowedLinkedReservationTypeEnumValues = []LinkedReservationType{
	"Linked",
	"Shared",
	"Default",
}

func (v *LinkedReservationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkedReservationType(value)
	for _, existing := range AllowedLinkedReservationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkedReservationType", value)
}

// NewLinkedReservationTypeFromValue returns a pointer to a valid LinkedReservationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkedReservationTypeFromValue(v string) (*LinkedReservationType, error) {
	ev := LinkedReservationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LinkedReservationType: valid values are %v", v, AllowedLinkedReservationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkedReservationType) IsValid() bool {
	for _, existing := range AllowedLinkedReservationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to linkedReservationType value
func (v LinkedReservationType) Ptr() *LinkedReservationType {
	return &v
}

type NullableLinkedReservationType struct {
	value *LinkedReservationType
	isSet bool
}

func (v NullableLinkedReservationType) Get() *LinkedReservationType {
	return v.value
}

func (v *NullableLinkedReservationType) Set(val *LinkedReservationType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedReservationType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedReservationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedReservationType(val *LinkedReservationType) *NullableLinkedReservationType {
	return &NullableLinkedReservationType{value: val, isSet: true}
}

func (v NullableLinkedReservationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedReservationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

