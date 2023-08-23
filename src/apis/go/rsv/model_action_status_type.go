/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ActionStatusType Action status of the Track It item(Open, Closed).
type ActionStatusType string

// List of actionStatusType
const (
	OPEN ActionStatusType = "Open"
	CLOSED ActionStatusType = "Closed"
)

// All allowed values of ActionStatusType enum
var AllowedActionStatusTypeEnumValues = []ActionStatusType{
	"Open",
	"Closed",
}

func (v *ActionStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionStatusType(value)
	for _, existing := range AllowedActionStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionStatusType", value)
}

// NewActionStatusTypeFromValue returns a pointer to a valid ActionStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionStatusTypeFromValue(v string) (*ActionStatusType, error) {
	ev := ActionStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionStatusType: valid values are %v", v, AllowedActionStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionStatusType) IsValid() bool {
	for _, existing := range AllowedActionStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to actionStatusType value
func (v ActionStatusType) Ptr() *ActionStatusType {
	return &v
}

type NullableActionStatusType struct {
	value *ActionStatusType
	isSet bool
}

func (v NullableActionStatusType) Get() *ActionStatusType {
	return v.value
}

func (v *NullableActionStatusType) Set(val *ActionStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableActionStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableActionStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionStatusType(val *ActionStatusType) *NullableActionStatusType {
	return &NullableActionStatusType{value: val, isSet: true}
}

func (v NullableActionStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

