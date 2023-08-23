/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// StagedAddressCleansStatus Simple type for status of staged address cleansing.
type StagedAddressCleansStatus string

// List of stagedAddressCleansStatus
const (
	NOT_CLEANSED StagedAddressCleansStatus = "NotCleansed"
	CLEANSED StagedAddressCleansStatus = "Cleansed"
	FAILURE StagedAddressCleansStatus = "Failure"
)

// All allowed values of StagedAddressCleansStatus enum
var AllowedStagedAddressCleansStatusEnumValues = []StagedAddressCleansStatus{
	"NotCleansed",
	"Cleansed",
	"Failure",
}

func (v *StagedAddressCleansStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StagedAddressCleansStatus(value)
	for _, existing := range AllowedStagedAddressCleansStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StagedAddressCleansStatus", value)
}

// NewStagedAddressCleansStatusFromValue returns a pointer to a valid StagedAddressCleansStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStagedAddressCleansStatusFromValue(v string) (*StagedAddressCleansStatus, error) {
	ev := StagedAddressCleansStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StagedAddressCleansStatus: valid values are %v", v, AllowedStagedAddressCleansStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StagedAddressCleansStatus) IsValid() bool {
	for _, existing := range AllowedStagedAddressCleansStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stagedAddressCleansStatus value
func (v StagedAddressCleansStatus) Ptr() *StagedAddressCleansStatus {
	return &v
}

type NullableStagedAddressCleansStatus struct {
	value *StagedAddressCleansStatus
	isSet bool
}

func (v NullableStagedAddressCleansStatus) Get() *StagedAddressCleansStatus {
	return v.value
}

func (v *NullableStagedAddressCleansStatus) Set(val *StagedAddressCleansStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStagedAddressCleansStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStagedAddressCleansStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagedAddressCleansStatus(val *StagedAddressCleansStatus) *NullableStagedAddressCleansStatus {
	return &NullableStagedAddressCleansStatus{value: val, isSet: true}
}

func (v NullableStagedAddressCleansStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagedAddressCleansStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

