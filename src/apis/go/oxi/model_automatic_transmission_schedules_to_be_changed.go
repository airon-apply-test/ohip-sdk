/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AutomaticTransmissionSchedulesToBeChanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutomaticTransmissionSchedulesToBeChanged{}

// AutomaticTransmissionSchedulesToBeChanged struct for AutomaticTransmissionSchedulesToBeChanged
type AutomaticTransmissionSchedulesToBeChanged struct {
	// List of automatic transmission schedules
	AutomaticTransmissionSchedules []AutomaticTransmissionScheduleType `json:"automaticTransmissionSchedules,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewAutomaticTransmissionSchedulesToBeChanged instantiates a new AutomaticTransmissionSchedulesToBeChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomaticTransmissionSchedulesToBeChanged() *AutomaticTransmissionSchedulesToBeChanged {
	this := AutomaticTransmissionSchedulesToBeChanged{}
	return &this
}

// NewAutomaticTransmissionSchedulesToBeChangedWithDefaults instantiates a new AutomaticTransmissionSchedulesToBeChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomaticTransmissionSchedulesToBeChangedWithDefaults() *AutomaticTransmissionSchedulesToBeChanged {
	this := AutomaticTransmissionSchedulesToBeChanged{}
	return &this
}

// GetAutomaticTransmissionSchedules returns the AutomaticTransmissionSchedules field value if set, zero value otherwise.
func (o *AutomaticTransmissionSchedulesToBeChanged) GetAutomaticTransmissionSchedules() []AutomaticTransmissionScheduleType {
	if o == nil || IsNil(o.AutomaticTransmissionSchedules) {
		var ret []AutomaticTransmissionScheduleType
		return ret
	}
	return o.AutomaticTransmissionSchedules
}

// GetAutomaticTransmissionSchedulesOk returns a tuple with the AutomaticTransmissionSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticTransmissionSchedulesToBeChanged) GetAutomaticTransmissionSchedulesOk() ([]AutomaticTransmissionScheduleType, bool) {
	if o == nil || IsNil(o.AutomaticTransmissionSchedules) {
		return nil, false
	}
	return o.AutomaticTransmissionSchedules, true
}

// HasAutomaticTransmissionSchedules returns a boolean if a field has been set.
func (o *AutomaticTransmissionSchedulesToBeChanged) HasAutomaticTransmissionSchedules() bool {
	if o != nil && !IsNil(o.AutomaticTransmissionSchedules) {
		return true
	}

	return false
}

// SetAutomaticTransmissionSchedules gets a reference to the given []AutomaticTransmissionScheduleType and assigns it to the AutomaticTransmissionSchedules field.
func (o *AutomaticTransmissionSchedulesToBeChanged) SetAutomaticTransmissionSchedules(v []AutomaticTransmissionScheduleType) {
	o.AutomaticTransmissionSchedules = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AutomaticTransmissionSchedulesToBeChanged) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticTransmissionSchedulesToBeChanged) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AutomaticTransmissionSchedulesToBeChanged) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *AutomaticTransmissionSchedulesToBeChanged) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AutomaticTransmissionSchedulesToBeChanged) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticTransmissionSchedulesToBeChanged) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AutomaticTransmissionSchedulesToBeChanged) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *AutomaticTransmissionSchedulesToBeChanged) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o AutomaticTransmissionSchedulesToBeChanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutomaticTransmissionSchedulesToBeChanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutomaticTransmissionSchedules) {
		toSerialize["automaticTransmissionSchedules"] = o.AutomaticTransmissionSchedules
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAutomaticTransmissionSchedulesToBeChanged struct {
	value *AutomaticTransmissionSchedulesToBeChanged
	isSet bool
}

func (v NullableAutomaticTransmissionSchedulesToBeChanged) Get() *AutomaticTransmissionSchedulesToBeChanged {
	return v.value
}

func (v *NullableAutomaticTransmissionSchedulesToBeChanged) Set(val *AutomaticTransmissionSchedulesToBeChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticTransmissionSchedulesToBeChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticTransmissionSchedulesToBeChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticTransmissionSchedulesToBeChanged(val *AutomaticTransmissionSchedulesToBeChanged) *NullableAutomaticTransmissionSchedulesToBeChanged {
	return &NullableAutomaticTransmissionSchedulesToBeChanged{value: val, isSet: true}
}

func (v NullableAutomaticTransmissionSchedulesToBeChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomaticTransmissionSchedulesToBeChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


