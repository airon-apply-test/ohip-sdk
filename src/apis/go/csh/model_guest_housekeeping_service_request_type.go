/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// GuestHousekeepingServiceRequestType Possible values for the Guest Service Status.
type GuestHousekeepingServiceRequestType string

// List of guestHousekeepingServiceRequestType
const (
	DO_NOT_DISTURB GuestHousekeepingServiceRequestType = "DoNotDisturb"
	MAKE_UP_ROOM GuestHousekeepingServiceRequestType = "MakeUpRoom"
	NO_STATUS_SELECTED GuestHousekeepingServiceRequestType = "NoStatusSelected"
)

// All allowed values of GuestHousekeepingServiceRequestType enum
var AllowedGuestHousekeepingServiceRequestTypeEnumValues = []GuestHousekeepingServiceRequestType{
	"DoNotDisturb",
	"MakeUpRoom",
	"NoStatusSelected",
}

func (v *GuestHousekeepingServiceRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GuestHousekeepingServiceRequestType(value)
	for _, existing := range AllowedGuestHousekeepingServiceRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GuestHousekeepingServiceRequestType", value)
}

// NewGuestHousekeepingServiceRequestTypeFromValue returns a pointer to a valid GuestHousekeepingServiceRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGuestHousekeepingServiceRequestTypeFromValue(v string) (*GuestHousekeepingServiceRequestType, error) {
	ev := GuestHousekeepingServiceRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GuestHousekeepingServiceRequestType: valid values are %v", v, AllowedGuestHousekeepingServiceRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GuestHousekeepingServiceRequestType) IsValid() bool {
	for _, existing := range AllowedGuestHousekeepingServiceRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to guestHousekeepingServiceRequestType value
func (v GuestHousekeepingServiceRequestType) Ptr() *GuestHousekeepingServiceRequestType {
	return &v
}

type NullableGuestHousekeepingServiceRequestType struct {
	value *GuestHousekeepingServiceRequestType
	isSet bool
}

func (v NullableGuestHousekeepingServiceRequestType) Get() *GuestHousekeepingServiceRequestType {
	return v.value
}

func (v *NullableGuestHousekeepingServiceRequestType) Set(val *GuestHousekeepingServiceRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestHousekeepingServiceRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestHousekeepingServiceRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestHousekeepingServiceRequestType(val *GuestHousekeepingServiceRequestType) *NullableGuestHousekeepingServiceRequestType {
	return &NullableGuestHousekeepingServiceRequestType{value: val, isSet: true}
}

func (v NullableGuestHousekeepingServiceRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestHousekeepingServiceRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

