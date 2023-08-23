/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ECertificateStatusType Indicates that OPERA E-Certificate is reserved.
type ECertificateStatusType string

// List of eCertificateStatusType
const (
	CANCELLED ECertificateStatusType = "Cancelled"
	CONSUMED ECertificateStatusType = "Consumed"
	DELETED ECertificateStatusType = "Deleted"
	EXPIRED ECertificateStatusType = "Expired"
	ISSUED ECertificateStatusType = "Issued"
	RESERVED ECertificateStatusType = "Reserved"
)

// All allowed values of ECertificateStatusType enum
var AllowedECertificateStatusTypeEnumValues = []ECertificateStatusType{
	"Cancelled",
	"Consumed",
	"Deleted",
	"Expired",
	"Issued",
	"Reserved",
}

func (v *ECertificateStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ECertificateStatusType(value)
	for _, existing := range AllowedECertificateStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ECertificateStatusType", value)
}

// NewECertificateStatusTypeFromValue returns a pointer to a valid ECertificateStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewECertificateStatusTypeFromValue(v string) (*ECertificateStatusType, error) {
	ev := ECertificateStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ECertificateStatusType: valid values are %v", v, AllowedECertificateStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ECertificateStatusType) IsValid() bool {
	for _, existing := range AllowedECertificateStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to eCertificateStatusType value
func (v ECertificateStatusType) Ptr() *ECertificateStatusType {
	return &v
}

type NullableECertificateStatusType struct {
	value *ECertificateStatusType
	isSet bool
}

func (v NullableECertificateStatusType) Get() *ECertificateStatusType {
	return v.value
}

func (v *NullableECertificateStatusType) Set(val *ECertificateStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableECertificateStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableECertificateStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECertificateStatusType(val *ECertificateStatusType) *NullableECertificateStatusType {
	return &NullableECertificateStatusType{value: val, isSet: true}
}

func (v NullableECertificateStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECertificateStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

