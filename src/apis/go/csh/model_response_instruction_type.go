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

// ResponseInstructionType Response types used in the Apply Final Postings request. Based on the request the folio summary , folio details or no details will be returned in the response.
type ResponseInstructionType string

// List of responseInstructionType
const (
	NONE ResponseInstructionType = "None"
	SUMMARY ResponseInstructionType = "Summary"
	DETAILS ResponseInstructionType = "Details"
)

// All allowed values of ResponseInstructionType enum
var AllowedResponseInstructionTypeEnumValues = []ResponseInstructionType{
	"None",
	"Summary",
	"Details",
}

func (v *ResponseInstructionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResponseInstructionType(value)
	for _, existing := range AllowedResponseInstructionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResponseInstructionType", value)
}

// NewResponseInstructionTypeFromValue returns a pointer to a valid ResponseInstructionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResponseInstructionTypeFromValue(v string) (*ResponseInstructionType, error) {
	ev := ResponseInstructionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResponseInstructionType: valid values are %v", v, AllowedResponseInstructionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResponseInstructionType) IsValid() bool {
	for _, existing := range AllowedResponseInstructionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to responseInstructionType value
func (v ResponseInstructionType) Ptr() *ResponseInstructionType {
	return &v
}

type NullableResponseInstructionType struct {
	value *ResponseInstructionType
	isSet bool
}

func (v NullableResponseInstructionType) Get() *ResponseInstructionType {
	return v.value
}

func (v *NullableResponseInstructionType) Set(val *ResponseInstructionType) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseInstructionType) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseInstructionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseInstructionType(val *ResponseInstructionType) *NullableResponseInstructionType {
	return &NullableResponseInstructionType{value: val, isSet: true}
}

func (v NullableResponseInstructionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseInstructionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

