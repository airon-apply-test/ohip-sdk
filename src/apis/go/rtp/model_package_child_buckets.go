/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PackageChildBuckets Simple type for package child buckets.
type PackageChildBuckets string

// List of packageChildBuckets
const (
	BUCKET1 PackageChildBuckets = "Bucket1"
	BUCKET2 PackageChildBuckets = "Bucket2"
	BUCKET3 PackageChildBuckets = "Bucket3"
)

// All allowed values of PackageChildBuckets enum
var AllowedPackageChildBucketsEnumValues = []PackageChildBuckets{
	"Bucket1",
	"Bucket2",
	"Bucket3",
}

func (v *PackageChildBuckets) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PackageChildBuckets(value)
	for _, existing := range AllowedPackageChildBucketsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PackageChildBuckets", value)
}

// NewPackageChildBucketsFromValue returns a pointer to a valid PackageChildBuckets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPackageChildBucketsFromValue(v string) (*PackageChildBuckets, error) {
	ev := PackageChildBuckets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PackageChildBuckets: valid values are %v", v, AllowedPackageChildBucketsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PackageChildBuckets) IsValid() bool {
	for _, existing := range AllowedPackageChildBucketsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to packageChildBuckets value
func (v PackageChildBuckets) Ptr() *PackageChildBuckets {
	return &v
}

type NullablePackageChildBuckets struct {
	value *PackageChildBuckets
	isSet bool
}

func (v NullablePackageChildBuckets) Get() *PackageChildBuckets {
	return v.value
}

func (v *NullablePackageChildBuckets) Set(val *PackageChildBuckets) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageChildBuckets) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageChildBuckets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageChildBuckets(val *PackageChildBuckets) *NullablePackageChildBuckets {
	return &NullablePackageChildBuckets{value: val, isSet: true}
}

func (v NullablePackageChildBuckets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageChildBuckets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

