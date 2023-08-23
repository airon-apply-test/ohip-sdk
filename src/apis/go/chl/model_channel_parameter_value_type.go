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

// ChannelParameterValueType The value types of parameter, possible types are Boolean, String, Date, and Number.
type ChannelParameterValueType string

// List of channelParameterValueType
const (
	BOOLEAN ChannelParameterValueType = "Boolean"
	STRING ChannelParameterValueType = "String"
	DATE ChannelParameterValueType = "Date"
	TIME ChannelParameterValueType = "Time"
	NUMBER ChannelParameterValueType = "Number"
	SINGLE_SELECT_LOV ChannelParameterValueType = "SingleSelectLOV"
	MULTI_SELECT_LOV ChannelParameterValueType = "MultiSelectLOV"
)

// All allowed values of ChannelParameterValueType enum
var AllowedChannelParameterValueTypeEnumValues = []ChannelParameterValueType{
	"Boolean",
	"String",
	"Date",
	"Time",
	"Number",
	"SingleSelectLOV",
	"MultiSelectLOV",
}

func (v *ChannelParameterValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelParameterValueType(value)
	for _, existing := range AllowedChannelParameterValueTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelParameterValueType", value)
}

// NewChannelParameterValueTypeFromValue returns a pointer to a valid ChannelParameterValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelParameterValueTypeFromValue(v string) (*ChannelParameterValueType, error) {
	ev := ChannelParameterValueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelParameterValueType: valid values are %v", v, AllowedChannelParameterValueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelParameterValueType) IsValid() bool {
	for _, existing := range AllowedChannelParameterValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to channelParameterValueType value
func (v ChannelParameterValueType) Ptr() *ChannelParameterValueType {
	return &v
}

type NullableChannelParameterValueType struct {
	value *ChannelParameterValueType
	isSet bool
}

func (v NullableChannelParameterValueType) Get() *ChannelParameterValueType {
	return v.value
}

func (v *NullableChannelParameterValueType) Set(val *ChannelParameterValueType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelParameterValueType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelParameterValueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelParameterValueType(val *ChannelParameterValueType) *NullableChannelParameterValueType {
	return &NullableChannelParameterValueType{value: val, isSet: true}
}

func (v NullableChannelParameterValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelParameterValueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

