/*
OPERA Cloud Block Configuration API

APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SalesManagerGoalsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesManagerGoalsInfo{}

// SalesManagerGoalsInfo You can use this API to retrieve individual Sales Manager Goals for a hotel, you can narrow the results using different search criteria
type SalesManagerGoalsInfo struct {
	// Detail Information about Sales Manager's goal.
	SalesManagerGoals []SalesManagerGoalType `json:"salesManagerGoals,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewSalesManagerGoalsInfo instantiates a new SalesManagerGoalsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesManagerGoalsInfo() *SalesManagerGoalsInfo {
	this := SalesManagerGoalsInfo{}
	return &this
}

// NewSalesManagerGoalsInfoWithDefaults instantiates a new SalesManagerGoalsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesManagerGoalsInfoWithDefaults() *SalesManagerGoalsInfo {
	this := SalesManagerGoalsInfo{}
	return &this
}

// GetSalesManagerGoals returns the SalesManagerGoals field value if set, zero value otherwise.
func (o *SalesManagerGoalsInfo) GetSalesManagerGoals() []SalesManagerGoalType {
	if o == nil || IsNil(o.SalesManagerGoals) {
		var ret []SalesManagerGoalType
		return ret
	}
	return o.SalesManagerGoals
}

// GetSalesManagerGoalsOk returns a tuple with the SalesManagerGoals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesManagerGoalsInfo) GetSalesManagerGoalsOk() ([]SalesManagerGoalType, bool) {
	if o == nil || IsNil(o.SalesManagerGoals) {
		return nil, false
	}
	return o.SalesManagerGoals, true
}

// HasSalesManagerGoals returns a boolean if a field has been set.
func (o *SalesManagerGoalsInfo) HasSalesManagerGoals() bool {
	if o != nil && !IsNil(o.SalesManagerGoals) {
		return true
	}

	return false
}

// SetSalesManagerGoals gets a reference to the given []SalesManagerGoalType and assigns it to the SalesManagerGoals field.
func (o *SalesManagerGoalsInfo) SetSalesManagerGoals(v []SalesManagerGoalType) {
	o.SalesManagerGoals = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SalesManagerGoalsInfo) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesManagerGoalsInfo) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SalesManagerGoalsInfo) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *SalesManagerGoalsInfo) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SalesManagerGoalsInfo) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesManagerGoalsInfo) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SalesManagerGoalsInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *SalesManagerGoalsInfo) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o SalesManagerGoalsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesManagerGoalsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SalesManagerGoals) {
		toSerialize["salesManagerGoals"] = o.SalesManagerGoals
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableSalesManagerGoalsInfo struct {
	value *SalesManagerGoalsInfo
	isSet bool
}

func (v NullableSalesManagerGoalsInfo) Get() *SalesManagerGoalsInfo {
	return v.value
}

func (v *NullableSalesManagerGoalsInfo) Set(val *SalesManagerGoalsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesManagerGoalsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesManagerGoalsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesManagerGoalsInfo(val *SalesManagerGoalsInfo) *NullableSalesManagerGoalsInfo {
	return &NullableSalesManagerGoalsInfo{value: val, isSet: true}
}

func (v NullableSalesManagerGoalsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesManagerGoalsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


