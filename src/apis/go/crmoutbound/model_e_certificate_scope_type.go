/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
	"fmt"
)

// ECertificateScopeType Indicates that OPERA E-Certificate is available for a specific list of hotels.
type ECertificateScopeType string

// List of eCertificateScopeType
const (
	ECERTIFICATESCOPETYPE_GLOBAL ECertificateScopeType = "Global"
	ECERTIFICATESCOPETYPE_HOTEL ECertificateScopeType = "Hotel"
	ECERTIFICATESCOPETYPE_MULTI_HOTEL ECertificateScopeType = "MultiHotel"
)

// All allowed values of ECertificateScopeType enum
var AllowedECertificateScopeTypeEnumValues = []ECertificateScopeType{
	"Global",
	"Hotel",
	"MultiHotel",
}

func (v *ECertificateScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ECertificateScopeType(value)
	for _, existing := range AllowedECertificateScopeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ECertificateScopeType", value)
}

// NewECertificateScopeTypeFromValue returns a pointer to a valid ECertificateScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewECertificateScopeTypeFromValue(v string) (*ECertificateScopeType, error) {
	ev := ECertificateScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ECertificateScopeType: valid values are %v", v, AllowedECertificateScopeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ECertificateScopeType) IsValid() bool {
	for _, existing := range AllowedECertificateScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to eCertificateScopeType value
func (v ECertificateScopeType) Ptr() *ECertificateScopeType {
	return &v
}

type NullableECertificateScopeType struct {
	value *ECertificateScopeType
	isSet bool
}

func (v NullableECertificateScopeType) Get() *ECertificateScopeType {
	return v.value
}

func (v *NullableECertificateScopeType) Set(val *ECertificateScopeType) {
	v.value = val
	v.isSet = true
}

func (v NullableECertificateScopeType) IsSet() bool {
	return v.isSet
}

func (v *NullableECertificateScopeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECertificateScopeType(val *ECertificateScopeType) *NullableECertificateScopeType {
	return &NullableECertificateScopeType{value: val, isSet: true}
}

func (v NullableECertificateScopeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECertificateScopeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

