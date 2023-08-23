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

// OrderByAscDescType the model 'OrderByAscDescType'
type OrderByAscDescType string

// List of orderByAscDescType
const (
	ASC OrderByAscDescType = "Asc"
	DESC OrderByAscDescType = "Desc"
)

// All allowed values of OrderByAscDescType enum
var AllowedOrderByAscDescTypeEnumValues = []OrderByAscDescType{
	"Asc",
	"Desc",
}

func (v *OrderByAscDescType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderByAscDescType(value)
	for _, existing := range AllowedOrderByAscDescTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderByAscDescType", value)
}

// NewOrderByAscDescTypeFromValue returns a pointer to a valid OrderByAscDescType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderByAscDescTypeFromValue(v string) (*OrderByAscDescType, error) {
	ev := OrderByAscDescType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderByAscDescType: valid values are %v", v, AllowedOrderByAscDescTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderByAscDescType) IsValid() bool {
	for _, existing := range AllowedOrderByAscDescTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderByAscDescType value
func (v OrderByAscDescType) Ptr() *OrderByAscDescType {
	return &v
}

type NullableOrderByAscDescType struct {
	value *OrderByAscDescType
	isSet bool
}

func (v NullableOrderByAscDescType) Get() *OrderByAscDescType {
	return v.value
}

func (v *NullableOrderByAscDescType) Set(val *OrderByAscDescType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderByAscDescType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderByAscDescType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderByAscDescType(val *OrderByAscDescType) *NullableOrderByAscDescType {
	return &NullableOrderByAscDescType{value: val, isSet: true}
}

func (v NullableOrderByAscDescType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderByAscDescType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

