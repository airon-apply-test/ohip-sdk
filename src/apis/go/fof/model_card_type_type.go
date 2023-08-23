/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CardTypeType This is required for Credit Card Payment Methods. This indicates the type of Credit Card associated with this payment method.
type CardTypeType string

// List of cardTypeType
const (
	AB CardTypeType = "Ab"
	AM CardTypeType = "Am"
	AX CardTypeType = "Ax"
	CB CardTypeType = "Cb"
	DC CardTypeType = "Dc"
	DS CardTypeType = "Ds"
	DT CardTypeType = "Dt"
	EC CardTypeType = "Ec"
	ER CardTypeType = "Er"
	JC CardTypeType = "Jc"
	JL CardTypeType = "Jl"
	MC CardTypeType = "Mc"
	NB CardTypeType = "Nb"
	SO CardTypeType = "So"
	ST CardTypeType = "St"
	SW CardTypeType = "Sw"
	VA CardTypeType = "Va"
	XY CardTypeType = "Xy"
	ZZ CardTypeType = "Zz"
	CP CardTypeType = "Cp"
	CU CardTypeType = "Cu"
)

// All allowed values of CardTypeType enum
var AllowedCardTypeTypeEnumValues = []CardTypeType{
	"Ab",
	"Am",
	"Ax",
	"Cb",
	"Dc",
	"Ds",
	"Dt",
	"Ec",
	"Er",
	"Jc",
	"Jl",
	"Mc",
	"Nb",
	"So",
	"St",
	"Sw",
	"Va",
	"Xy",
	"Zz",
	"Cp",
	"Cu",
}

func (v *CardTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardTypeType(value)
	for _, existing := range AllowedCardTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardTypeType", value)
}

// NewCardTypeTypeFromValue returns a pointer to a valid CardTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardTypeTypeFromValue(v string) (*CardTypeType, error) {
	ev := CardTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardTypeType: valid values are %v", v, AllowedCardTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardTypeType) IsValid() bool {
	for _, existing := range AllowedCardTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cardTypeType value
func (v CardTypeType) Ptr() *CardTypeType {
	return &v
}

type NullableCardTypeType struct {
	value *CardTypeType
	isSet bool
}

func (v NullableCardTypeType) Get() *CardTypeType {
	return v.value
}

func (v *NullableCardTypeType) Set(val *CardTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTypeType(val *CardTypeType) *NullableCardTypeType {
	return &NullableCardTypeType{value: val, isSet: true}
}

func (v NullableCardTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

