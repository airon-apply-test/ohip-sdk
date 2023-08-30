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

// ChannelParameterLevelType the model 'ChannelParameterLevelType'
type ChannelParameterLevelType string

// List of channelParameterLevelType
const (
	CHANNELPARAMETERLEVELTYPE_GLOBAL ChannelParameterLevelType = "Global"
	CHANNELPARAMETERLEVELTYPE_HOTEL ChannelParameterLevelType = "Hotel"
)

// All allowed values of ChannelParameterLevelType enum
var AllowedChannelParameterLevelTypeEnumValues = []ChannelParameterLevelType{
	"Global",
	"Hotel",
}

func (v *ChannelParameterLevelType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelParameterLevelType(value)
	for _, existing := range AllowedChannelParameterLevelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelParameterLevelType", value)
}

// NewChannelParameterLevelTypeFromValue returns a pointer to a valid ChannelParameterLevelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelParameterLevelTypeFromValue(v string) (*ChannelParameterLevelType, error) {
	ev := ChannelParameterLevelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelParameterLevelType: valid values are %v", v, AllowedChannelParameterLevelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelParameterLevelType) IsValid() bool {
	for _, existing := range AllowedChannelParameterLevelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to channelParameterLevelType value
func (v ChannelParameterLevelType) Ptr() *ChannelParameterLevelType {
	return &v
}

type NullableChannelParameterLevelType struct {
	value *ChannelParameterLevelType
	isSet bool
}

func (v NullableChannelParameterLevelType) Get() *ChannelParameterLevelType {
	return v.value
}

func (v *NullableChannelParameterLevelType) Set(val *ChannelParameterLevelType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelParameterLevelType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelParameterLevelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelParameterLevelType(val *ChannelParameterLevelType) *NullableChannelParameterLevelType {
	return &NullableChannelParameterLevelType{value: val, isSet: true}
}

func (v NullableChannelParameterLevelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelParameterLevelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

