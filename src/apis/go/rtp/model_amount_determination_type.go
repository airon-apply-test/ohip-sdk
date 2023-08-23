/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AmountDeterminationType Used to indicate if an amount is inclusive or exclusive of other charges, such as taxes.
type AmountDeterminationType string

// List of amountDeterminationType
const (
	INCLUSIVE AmountDeterminationType = "Inclusive"
	EXCLUSIVE AmountDeterminationType = "Exclusive"
)

// All allowed values of AmountDeterminationType enum
var AllowedAmountDeterminationTypeEnumValues = []AmountDeterminationType{
	"Inclusive",
	"Exclusive",
}

func (v *AmountDeterminationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AmountDeterminationType(value)
	for _, existing := range AllowedAmountDeterminationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AmountDeterminationType", value)
}

// NewAmountDeterminationTypeFromValue returns a pointer to a valid AmountDeterminationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAmountDeterminationTypeFromValue(v string) (*AmountDeterminationType, error) {
	ev := AmountDeterminationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AmountDeterminationType: valid values are %v", v, AllowedAmountDeterminationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AmountDeterminationType) IsValid() bool {
	for _, existing := range AllowedAmountDeterminationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to amountDeterminationType value
func (v AmountDeterminationType) Ptr() *AmountDeterminationType {
	return &v
}

type NullableAmountDeterminationType struct {
	value *AmountDeterminationType
	isSet bool
}

func (v NullableAmountDeterminationType) Get() *AmountDeterminationType {
	return v.value
}

func (v *NullableAmountDeterminationType) Set(val *AmountDeterminationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAmountDeterminationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAmountDeterminationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmountDeterminationType(val *AmountDeterminationType) *NullableAmountDeterminationType {
	return &NullableAmountDeterminationType{value: val, isSet: true}
}

func (v NullableAmountDeterminationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmountDeterminationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

