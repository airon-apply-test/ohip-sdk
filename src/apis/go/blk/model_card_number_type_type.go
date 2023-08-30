/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
	"fmt"
)

// CardNumberTypeType Simple type for indicating if credit card number is tokenized.
type CardNumberTypeType string

// List of cardNumberTypeType
const (
	CARDNUMBERTYPETYPE_CARD_NUMBER CardNumberTypeType = "CardNumber"
	CARDNUMBERTYPETYPE_TOKEN CardNumberTypeType = "Token"
)

// All allowed values of CardNumberTypeType enum
var AllowedCardNumberTypeTypeEnumValues = []CardNumberTypeType{
	"CardNumber",
	"Token",
}

func (v *CardNumberTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardNumberTypeType(value)
	for _, existing := range AllowedCardNumberTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardNumberTypeType", value)
}

// NewCardNumberTypeTypeFromValue returns a pointer to a valid CardNumberTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardNumberTypeTypeFromValue(v string) (*CardNumberTypeType, error) {
	ev := CardNumberTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardNumberTypeType: valid values are %v", v, AllowedCardNumberTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardNumberTypeType) IsValid() bool {
	for _, existing := range AllowedCardNumberTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cardNumberTypeType value
func (v CardNumberTypeType) Ptr() *CardNumberTypeType {
	return &v
}

type NullableCardNumberTypeType struct {
	value *CardNumberTypeType
	isSet bool
}

func (v NullableCardNumberTypeType) Get() *CardNumberTypeType {
	return v.value
}

func (v *NullableCardNumberTypeType) Set(val *CardNumberTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableCardNumberTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableCardNumberTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardNumberTypeType(val *CardNumberTypeType) *NullableCardNumberTypeType {
	return &NullableCardNumberTypeType{value: val, isSet: true}
}

func (v NullableCardNumberTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardNumberTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

