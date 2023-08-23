/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RewardCertificateType Point Saver reward.
type RewardCertificateType string

// List of rewardCertificateType
const (
	REGULAR RewardCertificateType = "Regular"
	STAY_ANYTIME RewardCertificateType = "StayAnytime"
	POINT_SAVER RewardCertificateType = "PointSaver"
)

// All allowed values of RewardCertificateType enum
var AllowedRewardCertificateTypeEnumValues = []RewardCertificateType{
	"Regular",
	"StayAnytime",
	"PointSaver",
}

func (v *RewardCertificateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RewardCertificateType(value)
	for _, existing := range AllowedRewardCertificateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RewardCertificateType", value)
}

// NewRewardCertificateTypeFromValue returns a pointer to a valid RewardCertificateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRewardCertificateTypeFromValue(v string) (*RewardCertificateType, error) {
	ev := RewardCertificateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RewardCertificateType: valid values are %v", v, AllowedRewardCertificateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RewardCertificateType) IsValid() bool {
	for _, existing := range AllowedRewardCertificateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to rewardCertificateType value
func (v RewardCertificateType) Ptr() *RewardCertificateType {
	return &v
}

type NullableRewardCertificateType struct {
	value *RewardCertificateType
	isSet bool
}

func (v NullableRewardCertificateType) Get() *RewardCertificateType {
	return v.value
}

func (v *NullableRewardCertificateType) Set(val *RewardCertificateType) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardCertificateType) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardCertificateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardCertificateType(val *RewardCertificateType) *NullableRewardCertificateType {
	return &NullableRewardCertificateType{value: val, isSet: true}
}

func (v NullableRewardCertificateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardCertificateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

