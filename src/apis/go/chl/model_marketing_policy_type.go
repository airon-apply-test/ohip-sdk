/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
	"fmt"
)

// MarketingPolicyType Indicates the channel policy type like deposit,cancel etc.
type MarketingPolicyType string

// List of marketingPolicyType
const (
	MARKETINGPOLICYTYPE_DEPOSIT MarketingPolicyType = "Deposit"
	MARKETINGPOLICYTYPE_COMMISSION MarketingPolicyType = "Commission"
	MARKETINGPOLICYTYPE_CANCEL MarketingPolicyType = "Cancel"
	MARKETINGPOLICYTYPE_GUARANTEE MarketingPolicyType = "Guarantee"
	MARKETINGPOLICYTYPE_GENERAL MarketingPolicyType = "General"
	MARKETINGPOLICYTYPE_PENALTY MarketingPolicyType = "Penalty"
	MARKETINGPOLICYTYPE_TAX MarketingPolicyType = "Tax"
	MARKETINGPOLICYTYPE_PROMOTION MarketingPolicyType = "Promotion"
)

// All allowed values of MarketingPolicyType enum
var AllowedMarketingPolicyTypeEnumValues = []MarketingPolicyType{
	"Deposit",
	"Commission",
	"Cancel",
	"Guarantee",
	"General",
	"Penalty",
	"Tax",
	"Promotion",
}

func (v *MarketingPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarketingPolicyType(value)
	for _, existing := range AllowedMarketingPolicyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarketingPolicyType", value)
}

// NewMarketingPolicyTypeFromValue returns a pointer to a valid MarketingPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarketingPolicyTypeFromValue(v string) (*MarketingPolicyType, error) {
	ev := MarketingPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarketingPolicyType: valid values are %v", v, AllowedMarketingPolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarketingPolicyType) IsValid() bool {
	for _, existing := range AllowedMarketingPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marketingPolicyType value
func (v MarketingPolicyType) Ptr() *MarketingPolicyType {
	return &v
}

type NullableMarketingPolicyType struct {
	value *MarketingPolicyType
	isSet bool
}

func (v NullableMarketingPolicyType) Get() *MarketingPolicyType {
	return v.value
}

func (v *NullableMarketingPolicyType) Set(val *MarketingPolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingPolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingPolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingPolicyType(val *MarketingPolicyType) *NullableMarketingPolicyType {
	return &NullableMarketingPolicyType{value: val, isSet: true}
}

func (v NullableMarketingPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingPolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

