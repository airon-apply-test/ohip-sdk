/*
OPERA Cloud API for Customer Management Service

This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TrackItGroup Group to which the Track It items belong.
type TrackItGroup string

// List of trackItGroup
const (
	PARCEL TrackItGroup = "Parcel"
	VALET TrackItGroup = "Valet"
	BAGGAGE TrackItGroup = "Baggage"
	LOST TrackItGroup = "Lost"
)

// All allowed values of TrackItGroup enum
var AllowedTrackItGroupEnumValues = []TrackItGroup{
	"Parcel",
	"Valet",
	"Baggage",
	"Lost",
}

func (v *TrackItGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrackItGroup(value)
	for _, existing := range AllowedTrackItGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrackItGroup", value)
}

// NewTrackItGroupFromValue returns a pointer to a valid TrackItGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrackItGroupFromValue(v string) (*TrackItGroup, error) {
	ev := TrackItGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrackItGroup: valid values are %v", v, AllowedTrackItGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrackItGroup) IsValid() bool {
	for _, existing := range AllowedTrackItGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to trackItGroup value
func (v TrackItGroup) Ptr() *TrackItGroup {
	return &v
}

type NullableTrackItGroup struct {
	value *TrackItGroup
	isSet bool
}

func (v NullableTrackItGroup) Get() *TrackItGroup {
	return v.value
}

func (v *NullableTrackItGroup) Set(val *TrackItGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackItGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackItGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackItGroup(val *TrackItGroup) *NullableTrackItGroup {
	return &NullableTrackItGroup{value: val, isSet: true}
}

func (v NullableTrackItGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackItGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

