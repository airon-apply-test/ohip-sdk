/*
OPERA Cloud Integration Processor API

APIs to get Business Events generated in OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package int

import (
	"encoding/json"
	"fmt"
)

// BusinessEventModuleType the model 'BusinessEventModuleType'
type BusinessEventModuleType string

// List of businessEventModuleType
const (
	BUSINESSEVENTMODULETYPE_ACTIVITY BusinessEventModuleType = "Activity"
	BUSINESSEVENTMODULETYPE_AVAILABILITY BusinessEventModuleType = "Availability"
	BUSINESSEVENTMODULETYPE_BLOCK BusinessEventModuleType = "Block"
	BUSINESSEVENTMODULETYPE_BLOCK_OFFSETS BusinessEventModuleType = "BlockOffsets"
	BUSINESSEVENTMODULETYPE_CASHIERING BusinessEventModuleType = "Cashiering"
	BUSINESSEVENTMODULETYPE_CATERING BusinessEventModuleType = "Catering"
	BUSINESSEVENTMODULETYPE_CATERING_NOTES BusinessEventModuleType = "CateringNotes"
	BUSINESSEVENTMODULETYPE_CONFIGURATION BusinessEventModuleType = "Configuration"
	BUSINESSEVENTMODULETYPE_E_CERTIFICATE BusinessEventModuleType = "ECertificate"
	BUSINESSEVENTMODULETYPE_ENROLLMENT BusinessEventModuleType = "Enrollment"
	BUSINESSEVENTMODULETYPE_HOUSEKEEPING BusinessEventModuleType = "Housekeeping"
	BUSINESSEVENTMODULETYPE_MR_ENROLLMENT BusinessEventModuleType = "MrEnrollment"
	BUSINESSEVENTMODULETYPE_NIGHT_AUDIT BusinessEventModuleType = "NightAudit"
	BUSINESSEVENTMODULETYPE_OWNER_CONTRACT BusinessEventModuleType = "OwnerContract"
	BUSINESSEVENTMODULETYPE_PROFILE BusinessEventModuleType = "Profile"
	BUSINESSEVENTMODULETYPE_RATE BusinessEventModuleType = "Rate"
	BUSINESSEVENTMODULETYPE_RESERVATION BusinessEventModuleType = "Reservation"
	BUSINESSEVENTMODULETYPE_STAY_RECORDS BusinessEventModuleType = "StayRecords"
)

// All allowed values of BusinessEventModuleType enum
var AllowedBusinessEventModuleTypeEnumValues = []BusinessEventModuleType{
	"Activity",
	"Availability",
	"Block",
	"BlockOffsets",
	"Cashiering",
	"Catering",
	"CateringNotes",
	"Configuration",
	"ECertificate",
	"Enrollment",
	"Housekeeping",
	"MrEnrollment",
	"NightAudit",
	"OwnerContract",
	"Profile",
	"Rate",
	"Reservation",
	"StayRecords",
}

func (v *BusinessEventModuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BusinessEventModuleType(value)
	for _, existing := range AllowedBusinessEventModuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BusinessEventModuleType", value)
}

// NewBusinessEventModuleTypeFromValue returns a pointer to a valid BusinessEventModuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBusinessEventModuleTypeFromValue(v string) (*BusinessEventModuleType, error) {
	ev := BusinessEventModuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BusinessEventModuleType: valid values are %v", v, AllowedBusinessEventModuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessEventModuleType) IsValid() bool {
	for _, existing := range AllowedBusinessEventModuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to businessEventModuleType value
func (v BusinessEventModuleType) Ptr() *BusinessEventModuleType {
	return &v
}

type NullableBusinessEventModuleType struct {
	value *BusinessEventModuleType
	isSet bool
}

func (v NullableBusinessEventModuleType) Get() *BusinessEventModuleType {
	return v.value
}

func (v *NullableBusinessEventModuleType) Set(val *BusinessEventModuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessEventModuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessEventModuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessEventModuleType(val *BusinessEventModuleType) *NullableBusinessEventModuleType {
	return &NullableBusinessEventModuleType{value: val, isSet: true}
}

func (v NullableBusinessEventModuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessEventModuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

