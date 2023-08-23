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

// StatisticsSummaryType Type of statistic for which summary information is populated.
type StatisticsSummaryType string

// List of statisticsSummaryType
const (
	POTENTIAL StatisticsSummaryType = "Potential"
	ACTUAL StatisticsSummaryType = "Actual"
	VARIANCE StatisticsSummaryType = "Variance"
)

// All allowed values of StatisticsSummaryType enum
var AllowedStatisticsSummaryTypeEnumValues = []StatisticsSummaryType{
	"Potential",
	"Actual",
	"Variance",
}

func (v *StatisticsSummaryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatisticsSummaryType(value)
	for _, existing := range AllowedStatisticsSummaryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatisticsSummaryType", value)
}

// NewStatisticsSummaryTypeFromValue returns a pointer to a valid StatisticsSummaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatisticsSummaryTypeFromValue(v string) (*StatisticsSummaryType, error) {
	ev := StatisticsSummaryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatisticsSummaryType: valid values are %v", v, AllowedStatisticsSummaryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatisticsSummaryType) IsValid() bool {
	for _, existing := range AllowedStatisticsSummaryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to statisticsSummaryType value
func (v StatisticsSummaryType) Ptr() *StatisticsSummaryType {
	return &v
}

type NullableStatisticsSummaryType struct {
	value *StatisticsSummaryType
	isSet bool
}

func (v NullableStatisticsSummaryType) Get() *StatisticsSummaryType {
	return v.value
}

func (v *NullableStatisticsSummaryType) Set(val *StatisticsSummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticsSummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticsSummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticsSummaryType(val *StatisticsSummaryType) *NullableStatisticsSummaryType {
	return &NullableStatisticsSummaryType{value: val, isSet: true}
}

func (v NullableStatisticsSummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticsSummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

