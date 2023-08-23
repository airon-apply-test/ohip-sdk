/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RoutingInstructionsToChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingInstructionsToChange{}

// RoutingInstructionsToChange Request when changing a routing instruction.
type RoutingInstructionsToChange struct {
	Criteria *RoutingInstructionsToChangeCriteria `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewRoutingInstructionsToChange instantiates a new RoutingInstructionsToChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingInstructionsToChange() *RoutingInstructionsToChange {
	this := RoutingInstructionsToChange{}
	return &this
}

// NewRoutingInstructionsToChangeWithDefaults instantiates a new RoutingInstructionsToChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingInstructionsToChangeWithDefaults() *RoutingInstructionsToChange {
	this := RoutingInstructionsToChange{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *RoutingInstructionsToChange) GetCriteria() RoutingInstructionsToChangeCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret RoutingInstructionsToChangeCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChange) GetCriteriaOk() (*RoutingInstructionsToChangeCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *RoutingInstructionsToChange) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given RoutingInstructionsToChangeCriteria and assigns it to the Criteria field.
func (o *RoutingInstructionsToChange) SetCriteria(v RoutingInstructionsToChangeCriteria) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoutingInstructionsToChange) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChange) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoutingInstructionsToChange) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *RoutingInstructionsToChange) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *RoutingInstructionsToChange) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChange) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RoutingInstructionsToChange) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *RoutingInstructionsToChange) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o RoutingInstructionsToChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingInstructionsToChange) ToMap() (map[string]interface{}, error) {
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

type NullableRoutingInstructionsToChange struct {
	value *RoutingInstructionsToChange
	isSet bool
}

func (v NullableRoutingInstructionsToChange) Get() *RoutingInstructionsToChange {
	return v.value
}

func (v *NullableRoutingInstructionsToChange) Set(val *RoutingInstructionsToChange) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingInstructionsToChange) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingInstructionsToChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingInstructionsToChange(val *RoutingInstructionsToChange) *NullableRoutingInstructionsToChange {
	return &NullableRoutingInstructionsToChange{value: val, isSet: true}
}

func (v NullableRoutingInstructionsToChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingInstructionsToChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


