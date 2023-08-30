/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
	"fmt"
)

// NameValueTypeType the model 'NameValueTypeType'
type NameValueTypeType string

// List of nameValueTypeType
const (
	NAMEVALUETYPETYPE_RESERVATION_OPTIONS NameValueTypeType = "ReservationOptions"
	NAMEVALUETYPETYPE_AFTER_SETTLEMENT NameValueTypeType = "AfterSettlement"
	NAMEVALUETYPETYPE_BEFORE_SETTLEMENT NameValueTypeType = "BeforeSettlement"
	NAMEVALUETYPETYPE_AFTER_PAYMENT NameValueTypeType = "AfterPayment"
	NAMEVALUETYPETYPE_PROFILE_OPTIONS NameValueTypeType = "ProfileOptions"
)

// All allowed values of NameValueTypeType enum
var AllowedNameValueTypeTypeEnumValues = []NameValueTypeType{
	"ReservationOptions",
	"AfterSettlement",
	"BeforeSettlement",
	"AfterPayment",
	"ProfileOptions",
}

func (v *NameValueTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NameValueTypeType(value)
	for _, existing := range AllowedNameValueTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NameValueTypeType", value)
}

// NewNameValueTypeTypeFromValue returns a pointer to a valid NameValueTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNameValueTypeTypeFromValue(v string) (*NameValueTypeType, error) {
	ev := NameValueTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NameValueTypeType: valid values are %v", v, AllowedNameValueTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NameValueTypeType) IsValid() bool {
	for _, existing := range AllowedNameValueTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to nameValueTypeType value
func (v NameValueTypeType) Ptr() *NameValueTypeType {
	return &v
}

type NullableNameValueTypeType struct {
	value *NameValueTypeType
	isSet bool
}

func (v NullableNameValueTypeType) Get() *NameValueTypeType {
	return v.value
}

func (v *NullableNameValueTypeType) Set(val *NameValueTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableNameValueTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableNameValueTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameValueTypeType(val *NameValueTypeType) *NullableNameValueTypeType {
	return &NullableNameValueTypeType{value: val, isSet: true}
}

func (v NullableNameValueTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameValueTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

