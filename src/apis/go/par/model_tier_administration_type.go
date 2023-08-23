/*
OPERA Cloud Price Availability Rate API

APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TierAdministrationType Do not upgrade membership. The membership may be downgraded.
type TierAdministrationType string

// List of tierAdministrationType
const (
	DISABLED TierAdministrationType = "Disabled"
	NO_UPGRADE TierAdministrationType = "NoUpgrade"
)

// All allowed values of TierAdministrationType enum
var AllowedTierAdministrationTypeEnumValues = []TierAdministrationType{
	"Disabled",
	"NoUpgrade",
}

func (v *TierAdministrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TierAdministrationType(value)
	for _, existing := range AllowedTierAdministrationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TierAdministrationType", value)
}

// NewTierAdministrationTypeFromValue returns a pointer to a valid TierAdministrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTierAdministrationTypeFromValue(v string) (*TierAdministrationType, error) {
	ev := TierAdministrationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TierAdministrationType: valid values are %v", v, AllowedTierAdministrationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TierAdministrationType) IsValid() bool {
	for _, existing := range AllowedTierAdministrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tierAdministrationType value
func (v TierAdministrationType) Ptr() *TierAdministrationType {
	return &v
}

type NullableTierAdministrationType struct {
	value *TierAdministrationType
	isSet bool
}

func (v NullableTierAdministrationType) Get() *TierAdministrationType {
	return v.value
}

func (v *NullableTierAdministrationType) Set(val *TierAdministrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableTierAdministrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableTierAdministrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTierAdministrationType(val *TierAdministrationType) *NullableTierAdministrationType {
	return &NullableTierAdministrationType{value: val, isSet: true}
}

func (v NullableTierAdministrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTierAdministrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

