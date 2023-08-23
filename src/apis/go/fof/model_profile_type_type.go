/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ProfileTypeType The types of Profile handled by the web service.
type ProfileTypeType string

// List of profileTypeType
const (
	GUEST ProfileTypeType = "Guest"
	AGENT ProfileTypeType = "Agent"
	COMPANY ProfileTypeType = "Company"
	GROUP ProfileTypeType = "Group"
	SOURCE ProfileTypeType = "Source"
	EMPLOYEE ProfileTypeType = "Employee"
	HOTEL ProfileTypeType = "Hotel"
	VENDOR ProfileTypeType = "Vendor"
	CONTACT ProfileTypeType = "Contact"
	PURGE ProfileTypeType = "Purge"
	BUSINESS_HEADER ProfileTypeType = "BusinessHeader"
	BILLING_ACCOUNT ProfileTypeType = "BillingAccount"
	ACTIVITY ProfileTypeType = "Activity"
	POTENTIAL ProfileTypeType = "Potential"
	ACCOUNT ProfileTypeType = "Account"
)

// All allowed values of ProfileTypeType enum
var AllowedProfileTypeTypeEnumValues = []ProfileTypeType{
	"Guest",
	"Agent",
	"Company",
	"Group",
	"Source",
	"Employee",
	"Hotel",
	"Vendor",
	"Contact",
	"Purge",
	"BusinessHeader",
	"BillingAccount",
	"Activity",
	"Potential",
	"Account",
}

func (v *ProfileTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProfileTypeType(value)
	for _, existing := range AllowedProfileTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProfileTypeType", value)
}

// NewProfileTypeTypeFromValue returns a pointer to a valid ProfileTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileTypeTypeFromValue(v string) (*ProfileTypeType, error) {
	ev := ProfileTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProfileTypeType: valid values are %v", v, AllowedProfileTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileTypeType) IsValid() bool {
	for _, existing := range AllowedProfileTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to profileTypeType value
func (v ProfileTypeType) Ptr() *ProfileTypeType {
	return &v
}

type NullableProfileTypeType struct {
	value *ProfileTypeType
	isSet bool
}

func (v NullableProfileTypeType) Get() *ProfileTypeType {
	return v.value
}

func (v *NullableProfileTypeType) Set(val *ProfileTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileTypeType(val *ProfileTypeType) *NullableProfileTypeType {
	return &NullableProfileTypeType{value: val, isSet: true}
}

func (v NullableProfileTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

