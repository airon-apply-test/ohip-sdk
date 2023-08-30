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

// ContractPriceFrequencyType To specify bills will be generated on Yearly basis.
type ContractPriceFrequencyType string

// List of contractPriceFrequencyType
const (
	CONTRACTPRICEFREQUENCYTYPE_MONTHLY ContractPriceFrequencyType = "Monthly"
	CONTRACTPRICEFREQUENCYTYPE_QUARTERLY ContractPriceFrequencyType = "Quarterly"
	CONTRACTPRICEFREQUENCYTYPE_HALF_YEARLY ContractPriceFrequencyType = "HalfYearly"
	CONTRACTPRICEFREQUENCYTYPE_YEARLY ContractPriceFrequencyType = "Yearly"
)

// All allowed values of ContractPriceFrequencyType enum
var AllowedContractPriceFrequencyTypeEnumValues = []ContractPriceFrequencyType{
	"Monthly",
	"Quarterly",
	"HalfYearly",
	"Yearly",
}

func (v *ContractPriceFrequencyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContractPriceFrequencyType(value)
	for _, existing := range AllowedContractPriceFrequencyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContractPriceFrequencyType", value)
}

// NewContractPriceFrequencyTypeFromValue returns a pointer to a valid ContractPriceFrequencyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContractPriceFrequencyTypeFromValue(v string) (*ContractPriceFrequencyType, error) {
	ev := ContractPriceFrequencyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContractPriceFrequencyType: valid values are %v", v, AllowedContractPriceFrequencyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContractPriceFrequencyType) IsValid() bool {
	for _, existing := range AllowedContractPriceFrequencyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to contractPriceFrequencyType value
func (v ContractPriceFrequencyType) Ptr() *ContractPriceFrequencyType {
	return &v
}

type NullableContractPriceFrequencyType struct {
	value *ContractPriceFrequencyType
	isSet bool
}

func (v NullableContractPriceFrequencyType) Get() *ContractPriceFrequencyType {
	return v.value
}

func (v *NullableContractPriceFrequencyType) Set(val *ContractPriceFrequencyType) {
	v.value = val
	v.isSet = true
}

func (v NullableContractPriceFrequencyType) IsSet() bool {
	return v.isSet
}

func (v *NullableContractPriceFrequencyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractPriceFrequencyType(val *ContractPriceFrequencyType) *NullableContractPriceFrequencyType {
	return &NullableContractPriceFrequencyType{value: val, isSet: true}
}

func (v NullableContractPriceFrequencyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractPriceFrequencyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

