/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ImageStyleType An enumeration of image styles.
type ImageStyleType string

// List of imageStyleType
const (
	FULL ImageStyleType = "Full"
	THUMBNAIL ImageStyleType = "Thumbnail"
	ICON ImageStyleType = "Icon"
	SQUARE ImageStyleType = "Square"
	BANNER ImageStyleType = "Banner"
	BUTTON ImageStyleType = "Button"
	LOWRES ImageStyleType = "Lowres"
	HIGHRES ImageStyleType = "Highres"
)

// All allowed values of ImageStyleType enum
var AllowedImageStyleTypeEnumValues = []ImageStyleType{
	"Full",
	"Thumbnail",
	"Icon",
	"Square",
	"Banner",
	"Button",
	"Lowres",
	"Highres",
}

func (v *ImageStyleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageStyleType(value)
	for _, existing := range AllowedImageStyleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageStyleType", value)
}

// NewImageStyleTypeFromValue returns a pointer to a valid ImageStyleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageStyleTypeFromValue(v string) (*ImageStyleType, error) {
	ev := ImageStyleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageStyleType: valid values are %v", v, AllowedImageStyleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageStyleType) IsValid() bool {
	for _, existing := range AllowedImageStyleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to imageStyleType value
func (v ImageStyleType) Ptr() *ImageStyleType {
	return &v
}

type NullableImageStyleType struct {
	value *ImageStyleType
	isSet bool
}

func (v NullableImageStyleType) Get() *ImageStyleType {
	return v.value
}

func (v *NullableImageStyleType) Set(val *ImageStyleType) {
	v.value = val
	v.isSet = true
}

func (v NullableImageStyleType) IsSet() bool {
	return v.isSet
}

func (v *NullableImageStyleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageStyleType(val *ImageStyleType) *NullableImageStyleType {
	return &NullableImageStyleType{value: val, isSet: true}
}

func (v NullableImageStyleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageStyleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

