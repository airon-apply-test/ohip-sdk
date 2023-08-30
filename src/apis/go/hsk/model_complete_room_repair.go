/*
OPERA Cloud Housekeeping Service API

APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hsk

import (
	"encoding/json"
)

// checks if the CompleteRoomRepair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteRoomRepair{}

// CompleteRoomRepair struct for CompleteRoomRepair
type CompleteRoomRepair struct {
	Criteria *CompleteRoomRepairCriteria `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCompleteRoomRepair instantiates a new CompleteRoomRepair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteRoomRepair() *CompleteRoomRepair {
	this := CompleteRoomRepair{}
	return &this
}

// NewCompleteRoomRepairWithDefaults instantiates a new CompleteRoomRepair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteRoomRepairWithDefaults() *CompleteRoomRepair {
	this := CompleteRoomRepair{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *CompleteRoomRepair) GetCriteria() CompleteRoomRepairCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret CompleteRoomRepairCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteRoomRepair) GetCriteriaOk() (*CompleteRoomRepairCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *CompleteRoomRepair) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given CompleteRoomRepairCriteria and assigns it to the Criteria field.
func (o *CompleteRoomRepair) SetCriteria(v CompleteRoomRepairCriteria) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CompleteRoomRepair) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteRoomRepair) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CompleteRoomRepair) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CompleteRoomRepair) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CompleteRoomRepair) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteRoomRepair) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CompleteRoomRepair) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CompleteRoomRepair) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CompleteRoomRepair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteRoomRepair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCompleteRoomRepair struct {
	value *CompleteRoomRepair
	isSet bool
}

func (v NullableCompleteRoomRepair) Get() *CompleteRoomRepair {
	return v.value
}

func (v *NullableCompleteRoomRepair) Set(val *CompleteRoomRepair) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteRoomRepair) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteRoomRepair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteRoomRepair(val *CompleteRoomRepair) *NullableCompleteRoomRepair {
	return &NullableCompleteRoomRepair{value: val, isSet: true}
}

func (v NullableCompleteRoomRepair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteRoomRepair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


