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

// PMSResStatusType This reservation is in checked in status and the business date is past departure date. This could occur when ORS and PMS are in same environment.
type PMSResStatusType string

// List of pMS_ResStatusType
const (
	RESERVED PMSResStatusType = "Reserved"
	REQUESTED PMSResStatusType = "Requested"
	NO_SHOW PMSResStatusType = "NoShow"
	CANCELLED PMSResStatusType = "Cancelled"
	IN_HOUSE PMSResStatusType = "InHouse"
	CHECKED_OUT PMSResStatusType = "CheckedOut"
	WAITLISTED PMSResStatusType = "Waitlisted"
	DUE_IN PMSResStatusType = "DueIn"
	DUE_OUT PMSResStatusType = "DueOut"
	WALKIN PMSResStatusType = "Walkin"
	PENDING_CHECKOUT PMSResStatusType = "PendingCheckout"
)

// All allowed values of PMSResStatusType enum
var AllowedPMSResStatusTypeEnumValues = []PMSResStatusType{
	"Reserved",
	"Requested",
	"NoShow",
	"Cancelled",
	"InHouse",
	"CheckedOut",
	"Waitlisted",
	"DueIn",
	"DueOut",
	"Walkin",
	"PendingCheckout",
}

func (v *PMSResStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PMSResStatusType(value)
	for _, existing := range AllowedPMSResStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PMSResStatusType", value)
}

// NewPMSResStatusTypeFromValue returns a pointer to a valid PMSResStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPMSResStatusTypeFromValue(v string) (*PMSResStatusType, error) {
	ev := PMSResStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PMSResStatusType: valid values are %v", v, AllowedPMSResStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PMSResStatusType) IsValid() bool {
	for _, existing := range AllowedPMSResStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to pMS_ResStatusType value
func (v PMSResStatusType) Ptr() *PMSResStatusType {
	return &v
}

type NullablePMSResStatusType struct {
	value *PMSResStatusType
	isSet bool
}

func (v NullablePMSResStatusType) Get() *PMSResStatusType {
	return v.value
}

func (v *NullablePMSResStatusType) Set(val *PMSResStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullablePMSResStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullablePMSResStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePMSResStatusType(val *PMSResStatusType) *NullablePMSResStatusType {
	return &NullablePMSResStatusType{value: val, isSet: true}
}

func (v NullablePMSResStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePMSResStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

