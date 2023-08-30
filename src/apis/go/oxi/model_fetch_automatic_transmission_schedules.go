/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the FetchAutomaticTransmissionSchedules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchAutomaticTransmissionSchedules{}

// FetchAutomaticTransmissionSchedules struct for FetchAutomaticTransmissionSchedules
type FetchAutomaticTransmissionSchedules struct {
	// List of automatic transmission schedules
	AutomaticTransmissonSchedules []AutomaticTransmissionScheduleType `json:"automaticTransmissonSchedules,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewFetchAutomaticTransmissionSchedules instantiates a new FetchAutomaticTransmissionSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchAutomaticTransmissionSchedules() *FetchAutomaticTransmissionSchedules {
	this := FetchAutomaticTransmissionSchedules{}
	return &this
}

// NewFetchAutomaticTransmissionSchedulesWithDefaults instantiates a new FetchAutomaticTransmissionSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchAutomaticTransmissionSchedulesWithDefaults() *FetchAutomaticTransmissionSchedules {
	this := FetchAutomaticTransmissionSchedules{}
	return &this
}

// GetAutomaticTransmissonSchedules returns the AutomaticTransmissonSchedules field value if set, zero value otherwise.
func (o *FetchAutomaticTransmissionSchedules) GetAutomaticTransmissonSchedules() []AutomaticTransmissionScheduleType {
	if o == nil || IsNil(o.AutomaticTransmissonSchedules) {
		var ret []AutomaticTransmissionScheduleType
		return ret
	}
	return o.AutomaticTransmissonSchedules
}

// GetAutomaticTransmissonSchedulesOk returns a tuple with the AutomaticTransmissonSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAutomaticTransmissionSchedules) GetAutomaticTransmissonSchedulesOk() ([]AutomaticTransmissionScheduleType, bool) {
	if o == nil || IsNil(o.AutomaticTransmissonSchedules) {
		return nil, false
	}
	return o.AutomaticTransmissonSchedules, true
}

// HasAutomaticTransmissonSchedules returns a boolean if a field has been set.
func (o *FetchAutomaticTransmissionSchedules) HasAutomaticTransmissonSchedules() bool {
	if o != nil && !IsNil(o.AutomaticTransmissonSchedules) {
		return true
	}

	return false
}

// SetAutomaticTransmissonSchedules gets a reference to the given []AutomaticTransmissionScheduleType and assigns it to the AutomaticTransmissonSchedules field.
func (o *FetchAutomaticTransmissionSchedules) SetAutomaticTransmissonSchedules(v []AutomaticTransmissionScheduleType) {
	o.AutomaticTransmissonSchedules = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FetchAutomaticTransmissionSchedules) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAutomaticTransmissionSchedules) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FetchAutomaticTransmissionSchedules) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *FetchAutomaticTransmissionSchedules) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FetchAutomaticTransmissionSchedules) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAutomaticTransmissionSchedules) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FetchAutomaticTransmissionSchedules) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *FetchAutomaticTransmissionSchedules) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o FetchAutomaticTransmissionSchedules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchAutomaticTransmissionSchedules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutomaticTransmissonSchedules) {
		toSerialize["automaticTransmissonSchedules"] = o.AutomaticTransmissonSchedules
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableFetchAutomaticTransmissionSchedules struct {
	value *FetchAutomaticTransmissionSchedules
	isSet bool
}

func (v NullableFetchAutomaticTransmissionSchedules) Get() *FetchAutomaticTransmissionSchedules {
	return v.value
}

func (v *NullableFetchAutomaticTransmissionSchedules) Set(val *FetchAutomaticTransmissionSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchAutomaticTransmissionSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchAutomaticTransmissionSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchAutomaticTransmissionSchedules(val *FetchAutomaticTransmissionSchedules) *NullableFetchAutomaticTransmissionSchedules {
	return &NullableFetchAutomaticTransmissionSchedules{value: val, isSet: true}
}

func (v NullableFetchAutomaticTransmissionSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchAutomaticTransmissionSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


