/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the FacilityTaskType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacilityTaskType{}

// FacilityTaskType Information regarding facility task on a reservation.
type FacilityTaskType struct {
	Task *HousekeepingTaskCodeType `json:"task,omitempty"`
	// List of the facility codes.
	Supplies []FacilityCodeType `json:"supplies,omitempty"`
	// The Date on which the task is applicable.
	Date *string `json:"date,omitempty"`
}

// NewFacilityTaskType instantiates a new FacilityTaskType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacilityTaskType() *FacilityTaskType {
	this := FacilityTaskType{}
	return &this
}

// NewFacilityTaskTypeWithDefaults instantiates a new FacilityTaskType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacilityTaskTypeWithDefaults() *FacilityTaskType {
	this := FacilityTaskType{}
	return &this
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *FacilityTaskType) GetTask() HousekeepingTaskCodeType {
	if o == nil || IsNil(o.Task) {
		var ret HousekeepingTaskCodeType
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacilityTaskType) GetTaskOk() (*HousekeepingTaskCodeType, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *FacilityTaskType) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given HousekeepingTaskCodeType and assigns it to the Task field.
func (o *FacilityTaskType) SetTask(v HousekeepingTaskCodeType) {
	o.Task = &v
}

// GetSupplies returns the Supplies field value if set, zero value otherwise.
func (o *FacilityTaskType) GetSupplies() []FacilityCodeType {
	if o == nil || IsNil(o.Supplies) {
		var ret []FacilityCodeType
		return ret
	}
	return o.Supplies
}

// GetSuppliesOk returns a tuple with the Supplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacilityTaskType) GetSuppliesOk() ([]FacilityCodeType, bool) {
	if o == nil || IsNil(o.Supplies) {
		return nil, false
	}
	return o.Supplies, true
}

// HasSupplies returns a boolean if a field has been set.
func (o *FacilityTaskType) HasSupplies() bool {
	if o != nil && !IsNil(o.Supplies) {
		return true
	}

	return false
}

// SetSupplies gets a reference to the given []FacilityCodeType and assigns it to the Supplies field.
func (o *FacilityTaskType) SetSupplies(v []FacilityCodeType) {
	o.Supplies = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *FacilityTaskType) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacilityTaskType) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *FacilityTaskType) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *FacilityTaskType) SetDate(v string) {
	o.Date = &v
}

func (o FacilityTaskType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacilityTaskType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.Supplies) {
		toSerialize["supplies"] = o.Supplies
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

type NullableFacilityTaskType struct {
	value *FacilityTaskType
	isSet bool
}

func (v NullableFacilityTaskType) Get() *FacilityTaskType {
	return v.value
}

func (v *NullableFacilityTaskType) Set(val *FacilityTaskType) {
	v.value = val
	v.isSet = true
}

func (v NullableFacilityTaskType) IsSet() bool {
	return v.isSet
}

func (v *NullableFacilityTaskType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacilityTaskType(val *FacilityTaskType) *NullableFacilityTaskType {
	return &NullableFacilityTaskType{value: val, isSet: true}
}

func (v NullableFacilityTaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacilityTaskType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


