/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the CloseFolioCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloseFolioCriteria{}

// CloseFolioCriteria Request to close an Open Folio Reservation. The Property should be using the Open Folio Functionality to use this operation.
type CloseFolioCriteria struct {
	Criteria *CloseFolioCriteriaType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCloseFolioCriteria instantiates a new CloseFolioCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloseFolioCriteria() *CloseFolioCriteria {
	this := CloseFolioCriteria{}
	return &this
}

// NewCloseFolioCriteriaWithDefaults instantiates a new CloseFolioCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloseFolioCriteriaWithDefaults() *CloseFolioCriteria {
	this := CloseFolioCriteria{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *CloseFolioCriteria) GetCriteria() CloseFolioCriteriaType {
	if o == nil || IsNil(o.Criteria) {
		var ret CloseFolioCriteriaType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseFolioCriteria) GetCriteriaOk() (*CloseFolioCriteriaType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *CloseFolioCriteria) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given CloseFolioCriteriaType and assigns it to the Criteria field.
func (o *CloseFolioCriteria) SetCriteria(v CloseFolioCriteriaType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CloseFolioCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseFolioCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CloseFolioCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CloseFolioCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CloseFolioCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseFolioCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CloseFolioCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CloseFolioCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CloseFolioCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloseFolioCriteria) ToMap() (map[string]interface{}, error) {
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

type NullableCloseFolioCriteria struct {
	value *CloseFolioCriteria
	isSet bool
}

func (v NullableCloseFolioCriteria) Get() *CloseFolioCriteria {
	return v.value
}

func (v *NullableCloseFolioCriteria) Set(val *CloseFolioCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseFolioCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseFolioCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseFolioCriteria(val *CloseFolioCriteria) *NullableCloseFolioCriteria {
	return &NullableCloseFolioCriteria{value: val, isSet: true}
}

func (v NullableCloseFolioCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseFolioCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


