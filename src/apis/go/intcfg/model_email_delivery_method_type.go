/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// EmailDeliveryMethodType Email delivery method type
type EmailDeliveryMethodType string

// List of emailDeliveryMethodType
const (
	BILLING EmailDeliveryMethodType = "Billing"
	CONFIRMATION EmailDeliveryMethodType = "Confirmation"
	GENERAL EmailDeliveryMethodType = "General"
	OTHER EmailDeliveryMethodType = "Other"
)

// All allowed values of EmailDeliveryMethodType enum
var AllowedEmailDeliveryMethodTypeEnumValues = []EmailDeliveryMethodType{
	"Billing",
	"Confirmation",
	"General",
	"Other",
}

func (v *EmailDeliveryMethodType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmailDeliveryMethodType(value)
	for _, existing := range AllowedEmailDeliveryMethodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmailDeliveryMethodType", value)
}

// NewEmailDeliveryMethodTypeFromValue returns a pointer to a valid EmailDeliveryMethodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmailDeliveryMethodTypeFromValue(v string) (*EmailDeliveryMethodType, error) {
	ev := EmailDeliveryMethodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmailDeliveryMethodType: valid values are %v", v, AllowedEmailDeliveryMethodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailDeliveryMethodType) IsValid() bool {
	for _, existing := range AllowedEmailDeliveryMethodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to emailDeliveryMethodType value
func (v EmailDeliveryMethodType) Ptr() *EmailDeliveryMethodType {
	return &v
}

type NullableEmailDeliveryMethodType struct {
	value *EmailDeliveryMethodType
	isSet bool
}

func (v NullableEmailDeliveryMethodType) Get() *EmailDeliveryMethodType {
	return v.value
}

func (v *NullableEmailDeliveryMethodType) Set(val *EmailDeliveryMethodType) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDeliveryMethodType) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDeliveryMethodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDeliveryMethodType(val *EmailDeliveryMethodType) *NullableEmailDeliveryMethodType {
	return &NullableEmailDeliveryMethodType{value: val, isSet: true}
}

func (v NullableEmailDeliveryMethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDeliveryMethodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

