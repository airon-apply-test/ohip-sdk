/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
	"fmt"
)

// DepositProcessType The Deposit is transferred to an internal account for returned deposits to be handled by the property. Users can select to keep all (not exceeding total), a portion of, or none of the deposit.
type DepositProcessType string

// List of depositProcessType
const (
	DEPOSITPROCESSTYPE_KEEP DepositProcessType = "Keep"
	DEPOSITPROCESSTYPE_RETURN DepositProcessType = "Return"
	DEPOSITPROCESSTYPE_MATURE DepositProcessType = "Mature"
)

// All allowed values of DepositProcessType enum
var AllowedDepositProcessTypeEnumValues = []DepositProcessType{
	"Keep",
	"Return",
	"Mature",
}

func (v *DepositProcessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DepositProcessType(value)
	for _, existing := range AllowedDepositProcessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DepositProcessType", value)
}

// NewDepositProcessTypeFromValue returns a pointer to a valid DepositProcessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDepositProcessTypeFromValue(v string) (*DepositProcessType, error) {
	ev := DepositProcessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DepositProcessType: valid values are %v", v, AllowedDepositProcessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DepositProcessType) IsValid() bool {
	for _, existing := range AllowedDepositProcessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to depositProcessType value
func (v DepositProcessType) Ptr() *DepositProcessType {
	return &v
}

type NullableDepositProcessType struct {
	value *DepositProcessType
	isSet bool
}

func (v NullableDepositProcessType) Get() *DepositProcessType {
	return v.value
}

func (v *NullableDepositProcessType) Set(val *DepositProcessType) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositProcessType) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositProcessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositProcessType(val *DepositProcessType) *NullableDepositProcessType {
	return &NullableDepositProcessType{value: val, isSet: true}
}

func (v NullableDepositProcessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositProcessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

