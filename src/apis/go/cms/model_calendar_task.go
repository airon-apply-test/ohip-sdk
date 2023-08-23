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
)

// checks if the CalendarTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalendarTask{}

// CalendarTask Response for fetching calendar task.
type CalendarTask struct {
	CalendarTaskDetails *CalendarTaskType `json:"calendarTaskDetails,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCalendarTask instantiates a new CalendarTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarTask() *CalendarTask {
	this := CalendarTask{}
	return &this
}

// NewCalendarTaskWithDefaults instantiates a new CalendarTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarTaskWithDefaults() *CalendarTask {
	this := CalendarTask{}
	return &this
}

// GetCalendarTaskDetails returns the CalendarTaskDetails field value if set, zero value otherwise.
func (o *CalendarTask) GetCalendarTaskDetails() CalendarTaskType {
	if o == nil || IsNil(o.CalendarTaskDetails) {
		var ret CalendarTaskType
		return ret
	}
	return *o.CalendarTaskDetails
}

// GetCalendarTaskDetailsOk returns a tuple with the CalendarTaskDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTask) GetCalendarTaskDetailsOk() (*CalendarTaskType, bool) {
	if o == nil || IsNil(o.CalendarTaskDetails) {
		return nil, false
	}
	return o.CalendarTaskDetails, true
}

// HasCalendarTaskDetails returns a boolean if a field has been set.
func (o *CalendarTask) HasCalendarTaskDetails() bool {
	if o != nil && !IsNil(o.CalendarTaskDetails) {
		return true
	}

	return false
}

// SetCalendarTaskDetails gets a reference to the given CalendarTaskType and assigns it to the CalendarTaskDetails field.
func (o *CalendarTask) SetCalendarTaskDetails(v CalendarTaskType) {
	o.CalendarTaskDetails = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CalendarTask) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTask) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CalendarTask) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CalendarTask) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CalendarTask) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTask) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CalendarTask) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CalendarTask) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CalendarTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalendarTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CalendarTaskDetails) {
		toSerialize["calendarTaskDetails"] = o.CalendarTaskDetails
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCalendarTask struct {
	value *CalendarTask
	isSet bool
}

func (v NullableCalendarTask) Get() *CalendarTask {
	return v.value
}

func (v *NullableCalendarTask) Set(val *CalendarTask) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarTask) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarTask(val *CalendarTask) *NullableCalendarTask {
	return &NullableCalendarTask{value: val, isSet: true}
}

func (v NullableCalendarTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


