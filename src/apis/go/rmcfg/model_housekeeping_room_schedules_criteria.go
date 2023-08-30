/*
OPERA Cloud Room Configuration API

APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rmcfg

import (
	"encoding/json"
)

// checks if the HousekeepingRoomSchedulesCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HousekeepingRoomSchedulesCriteria{}

// HousekeepingRoomSchedulesCriteria Request object for associating housekeeping tasks and housekeeping codes to a room type.
type HousekeepingRoomSchedulesCriteria struct {
	// This type holds a collection of housekeeping tasks attached to a room type.
	HousekeepingRoomSchedules []ConfigHousekeepingRoomScheduleType `json:"housekeepingRoomSchedules,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewHousekeepingRoomSchedulesCriteria instantiates a new HousekeepingRoomSchedulesCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHousekeepingRoomSchedulesCriteria() *HousekeepingRoomSchedulesCriteria {
	this := HousekeepingRoomSchedulesCriteria{}
	return &this
}

// NewHousekeepingRoomSchedulesCriteriaWithDefaults instantiates a new HousekeepingRoomSchedulesCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHousekeepingRoomSchedulesCriteriaWithDefaults() *HousekeepingRoomSchedulesCriteria {
	this := HousekeepingRoomSchedulesCriteria{}
	return &this
}

// GetHousekeepingRoomSchedules returns the HousekeepingRoomSchedules field value if set, zero value otherwise.
func (o *HousekeepingRoomSchedulesCriteria) GetHousekeepingRoomSchedules() []ConfigHousekeepingRoomScheduleType {
	if o == nil || IsNil(o.HousekeepingRoomSchedules) {
		var ret []ConfigHousekeepingRoomScheduleType
		return ret
	}
	return o.HousekeepingRoomSchedules
}

// GetHousekeepingRoomSchedulesOk returns a tuple with the HousekeepingRoomSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingRoomSchedulesCriteria) GetHousekeepingRoomSchedulesOk() ([]ConfigHousekeepingRoomScheduleType, bool) {
	if o == nil || IsNil(o.HousekeepingRoomSchedules) {
		return nil, false
	}
	return o.HousekeepingRoomSchedules, true
}

// HasHousekeepingRoomSchedules returns a boolean if a field has been set.
func (o *HousekeepingRoomSchedulesCriteria) HasHousekeepingRoomSchedules() bool {
	if o != nil && !IsNil(o.HousekeepingRoomSchedules) {
		return true
	}

	return false
}

// SetHousekeepingRoomSchedules gets a reference to the given []ConfigHousekeepingRoomScheduleType and assigns it to the HousekeepingRoomSchedules field.
func (o *HousekeepingRoomSchedulesCriteria) SetHousekeepingRoomSchedules(v []ConfigHousekeepingRoomScheduleType) {
	o.HousekeepingRoomSchedules = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *HousekeepingRoomSchedulesCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingRoomSchedulesCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HousekeepingRoomSchedulesCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *HousekeepingRoomSchedulesCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *HousekeepingRoomSchedulesCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingRoomSchedulesCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *HousekeepingRoomSchedulesCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *HousekeepingRoomSchedulesCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o HousekeepingRoomSchedulesCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HousekeepingRoomSchedulesCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HousekeepingRoomSchedules) {
		toSerialize["housekeepingRoomSchedules"] = o.HousekeepingRoomSchedules
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableHousekeepingRoomSchedulesCriteria struct {
	value *HousekeepingRoomSchedulesCriteria
	isSet bool
}

func (v NullableHousekeepingRoomSchedulesCriteria) Get() *HousekeepingRoomSchedulesCriteria {
	return v.value
}

func (v *NullableHousekeepingRoomSchedulesCriteria) Set(val *HousekeepingRoomSchedulesCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableHousekeepingRoomSchedulesCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableHousekeepingRoomSchedulesCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHousekeepingRoomSchedulesCriteria(val *HousekeepingRoomSchedulesCriteria) *NullableHousekeepingRoomSchedulesCriteria {
	return &NullableHousekeepingRoomSchedulesCriteria{value: val, isSet: true}
}

func (v NullableHousekeepingRoomSchedulesCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHousekeepingRoomSchedulesCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


