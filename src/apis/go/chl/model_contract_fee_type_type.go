/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ContractFeeTypeType To specify contract element is for Miscellaneous fee.
type ContractFeeTypeType string

// List of contractFeeTypeType
const (
	SETUP ContractFeeTypeType = "Setup"
	MAINTENANCE ContractFeeTypeType = "Maintenance"
	TRANSACTION ContractFeeTypeType = "Transaction"
	FAX ContractFeeTypeType = "Fax"
	MINIMUM_FEE ContractFeeTypeType = "MinimumFee"
	MISCELLANEOUS ContractFeeTypeType = "Miscellaneous"
)

// All allowed values of ContractFeeTypeType enum
var AllowedContractFeeTypeTypeEnumValues = []ContractFeeTypeType{
	"Setup",
	"Maintenance",
	"Transaction",
	"Fax",
	"MinimumFee",
	"Miscellaneous",
}

func (v *ContractFeeTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContractFeeTypeType(value)
	for _, existing := range AllowedContractFeeTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContractFeeTypeType", value)
}

// NewContractFeeTypeTypeFromValue returns a pointer to a valid ContractFeeTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContractFeeTypeTypeFromValue(v string) (*ContractFeeTypeType, error) {
	ev := ContractFeeTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContractFeeTypeType: valid values are %v", v, AllowedContractFeeTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContractFeeTypeType) IsValid() bool {
	for _, existing := range AllowedContractFeeTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to contractFeeTypeType value
func (v ContractFeeTypeType) Ptr() *ContractFeeTypeType {
	return &v
}

type NullableContractFeeTypeType struct {
	value *ContractFeeTypeType
	isSet bool
}

func (v NullableContractFeeTypeType) Get() *ContractFeeTypeType {
	return v.value
}

func (v *NullableContractFeeTypeType) Set(val *ContractFeeTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableContractFeeTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableContractFeeTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractFeeTypeType(val *ContractFeeTypeType) *NullableContractFeeTypeType {
	return &NullableContractFeeTypeType{value: val, isSet: true}
}

func (v NullableContractFeeTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractFeeTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

