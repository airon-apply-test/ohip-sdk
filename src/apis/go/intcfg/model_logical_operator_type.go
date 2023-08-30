/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intcfg

import (
	"encoding/json"
	"fmt"
)

// LogicalOperatorType the model 'LogicalOperatorType'
type LogicalOperatorType string

// List of logicalOperatorType
const (
	LOGICALOPERATORTYPE_AND LogicalOperatorType = "And"
	LOGICALOPERATORTYPE_OR LogicalOperatorType = "Or"
)

// All allowed values of LogicalOperatorType enum
var AllowedLogicalOperatorTypeEnumValues = []LogicalOperatorType{
	"And",
	"Or",
}

func (v *LogicalOperatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogicalOperatorType(value)
	for _, existing := range AllowedLogicalOperatorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogicalOperatorType", value)
}

// NewLogicalOperatorTypeFromValue returns a pointer to a valid LogicalOperatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogicalOperatorTypeFromValue(v string) (*LogicalOperatorType, error) {
	ev := LogicalOperatorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogicalOperatorType: valid values are %v", v, AllowedLogicalOperatorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogicalOperatorType) IsValid() bool {
	for _, existing := range AllowedLogicalOperatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to logicalOperatorType value
func (v LogicalOperatorType) Ptr() *LogicalOperatorType {
	return &v
}

type NullableLogicalOperatorType struct {
	value *LogicalOperatorType
	isSet bool
}

func (v NullableLogicalOperatorType) Get() *LogicalOperatorType {
	return v.value
}

func (v *NullableLogicalOperatorType) Set(val *LogicalOperatorType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalOperatorType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalOperatorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalOperatorType(val *LogicalOperatorType) *NullableLogicalOperatorType {
	return &NullableLogicalOperatorType{value: val, isSet: true}
}

func (v NullableLogicalOperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalOperatorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

