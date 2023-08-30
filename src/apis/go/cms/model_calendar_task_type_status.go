/*
OPERA Cloud API for Customer Management Service

This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cms

import (
	"encoding/json"
)

// checks if the CalendarTaskTypeStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalendarTaskTypeStatus{}

// CalendarTaskTypeStatus Completion status of calendar task.
type CalendarTaskTypeStatus struct {
	// Name of person who completed the task. Ignored when task is not marked as completed.
	CompletedBy *string `json:"completedBy,omitempty"`
	// Date on which the task was completed. Ignored when task is not marked as completed.
	CompletedOn *string `json:"completedOn,omitempty"`
	// Whether the task is completed or not.
	Completed *bool `json:"completed,omitempty"`
}

// NewCalendarTaskTypeStatus instantiates a new CalendarTaskTypeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarTaskTypeStatus() *CalendarTaskTypeStatus {
	this := CalendarTaskTypeStatus{}
	return &this
}

// NewCalendarTaskTypeStatusWithDefaults instantiates a new CalendarTaskTypeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarTaskTypeStatusWithDefaults() *CalendarTaskTypeStatus {
	this := CalendarTaskTypeStatus{}
	return &this
}

// GetCompletedBy returns the CompletedBy field value if set, zero value otherwise.
func (o *CalendarTaskTypeStatus) GetCompletedBy() string {
	if o == nil || IsNil(o.CompletedBy) {
		var ret string
		return ret
	}
	return *o.CompletedBy
}

// GetCompletedByOk returns a tuple with the CompletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTaskTypeStatus) GetCompletedByOk() (*string, bool) {
	if o == nil || IsNil(o.CompletedBy) {
		return nil, false
	}
	return o.CompletedBy, true
}

// HasCompletedBy returns a boolean if a field has been set.
func (o *CalendarTaskTypeStatus) HasCompletedBy() bool {
	if o != nil && !IsNil(o.CompletedBy) {
		return true
	}

	return false
}

// SetCompletedBy gets a reference to the given string and assigns it to the CompletedBy field.
func (o *CalendarTaskTypeStatus) SetCompletedBy(v string) {
	o.CompletedBy = &v
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise.
func (o *CalendarTaskTypeStatus) GetCompletedOn() string {
	if o == nil || IsNil(o.CompletedOn) {
		var ret string
		return ret
	}
	return *o.CompletedOn
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTaskTypeStatus) GetCompletedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CompletedOn) {
		return nil, false
	}
	return o.CompletedOn, true
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *CalendarTaskTypeStatus) HasCompletedOn() bool {
	if o != nil && !IsNil(o.CompletedOn) {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given string and assigns it to the CompletedOn field.
func (o *CalendarTaskTypeStatus) SetCompletedOn(v string) {
	o.CompletedOn = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *CalendarTaskTypeStatus) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTaskTypeStatus) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *CalendarTaskTypeStatus) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *CalendarTaskTypeStatus) SetCompleted(v bool) {
	o.Completed = &v
}

func (o CalendarTaskTypeStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalendarTaskTypeStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompletedBy) {
		toSerialize["completedBy"] = o.CompletedBy
	}
	if !IsNil(o.CompletedOn) {
		toSerialize["completedOn"] = o.CompletedOn
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	return toSerialize, nil
}

type NullableCalendarTaskTypeStatus struct {
	value *CalendarTaskTypeStatus
	isSet bool
}

func (v NullableCalendarTaskTypeStatus) Get() *CalendarTaskTypeStatus {
	return v.value
}

func (v *NullableCalendarTaskTypeStatus) Set(val *CalendarTaskTypeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarTaskTypeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarTaskTypeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarTaskTypeStatus(val *CalendarTaskTypeStatus) *NullableCalendarTaskTypeStatus {
	return &NullableCalendarTaskTypeStatus{value: val, isSet: true}
}

func (v NullableCalendarTaskTypeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarTaskTypeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


