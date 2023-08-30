/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
	"fmt"
)

// BucketTypeType Fiscal Bucket Type.
type BucketTypeType string

// List of bucketTypeType
const (
	BUCKETTYPETYPE_DEFAULT BucketTypeType = "Default"
	BUCKETTYPETYPE_FISCAL BucketTypeType = "Fiscal"
)

// All allowed values of BucketTypeType enum
var AllowedBucketTypeTypeEnumValues = []BucketTypeType{
	"Default",
	"Fiscal",
}

func (v *BucketTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BucketTypeType(value)
	for _, existing := range AllowedBucketTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BucketTypeType", value)
}

// NewBucketTypeTypeFromValue returns a pointer to a valid BucketTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBucketTypeTypeFromValue(v string) (*BucketTypeType, error) {
	ev := BucketTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BucketTypeType: valid values are %v", v, AllowedBucketTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BucketTypeType) IsValid() bool {
	for _, existing := range AllowedBucketTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bucketTypeType value
func (v BucketTypeType) Ptr() *BucketTypeType {
	return &v
}

type NullableBucketTypeType struct {
	value *BucketTypeType
	isSet bool
}

func (v NullableBucketTypeType) Get() *BucketTypeType {
	return v.value
}

func (v *NullableBucketTypeType) Set(val *BucketTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketTypeType(val *BucketTypeType) *NullableBucketTypeType {
	return &NullableBucketTypeType{value: val, isSet: true}
}

func (v NullableBucketTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

