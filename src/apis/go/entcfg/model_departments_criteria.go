/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
)

// checks if the DepartmentsCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepartmentsCriteria{}

// DepartmentsCriteria Request object for creating new departments.
type DepartmentsCriteria struct {
	// Collection of departments.
	Departments []DepartmentType `json:"departments,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewDepartmentsCriteria instantiates a new DepartmentsCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepartmentsCriteria() *DepartmentsCriteria {
	this := DepartmentsCriteria{}
	return &this
}

// NewDepartmentsCriteriaWithDefaults instantiates a new DepartmentsCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepartmentsCriteriaWithDefaults() *DepartmentsCriteria {
	this := DepartmentsCriteria{}
	return &this
}

// GetDepartments returns the Departments field value if set, zero value otherwise.
func (o *DepartmentsCriteria) GetDepartments() []DepartmentType {
	if o == nil || IsNil(o.Departments) {
		var ret []DepartmentType
		return ret
	}
	return o.Departments
}

// GetDepartmentsOk returns a tuple with the Departments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepartmentsCriteria) GetDepartmentsOk() ([]DepartmentType, bool) {
	if o == nil || IsNil(o.Departments) {
		return nil, false
	}
	return o.Departments, true
}

// HasDepartments returns a boolean if a field has been set.
func (o *DepartmentsCriteria) HasDepartments() bool {
	if o != nil && !IsNil(o.Departments) {
		return true
	}

	return false
}

// SetDepartments gets a reference to the given []DepartmentType and assigns it to the Departments field.
func (o *DepartmentsCriteria) SetDepartments(v []DepartmentType) {
	o.Departments = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DepartmentsCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepartmentsCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DepartmentsCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *DepartmentsCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DepartmentsCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepartmentsCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DepartmentsCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *DepartmentsCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o DepartmentsCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepartmentsCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Departments) {
		toSerialize["departments"] = o.Departments
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableDepartmentsCriteria struct {
	value *DepartmentsCriteria
	isSet bool
}

func (v NullableDepartmentsCriteria) Get() *DepartmentsCriteria {
	return v.value
}

func (v *NullableDepartmentsCriteria) Set(val *DepartmentsCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableDepartmentsCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableDepartmentsCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepartmentsCriteria(val *DepartmentsCriteria) *NullableDepartmentsCriteria {
	return &NullableDepartmentsCriteria{value: val, isSet: true}
}

func (v NullableDepartmentsCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepartmentsCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


