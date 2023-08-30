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

// checks if the PropertyTypesCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyTypesCriteria{}

// PropertyTypesCriteria Request object for creating Property Types.
type PropertyTypesCriteria struct {
	// List of Property Types.
	PropertyTypes []PropertyTypeType `json:"propertyTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPropertyTypesCriteria instantiates a new PropertyTypesCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyTypesCriteria() *PropertyTypesCriteria {
	this := PropertyTypesCriteria{}
	return &this
}

// NewPropertyTypesCriteriaWithDefaults instantiates a new PropertyTypesCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyTypesCriteriaWithDefaults() *PropertyTypesCriteria {
	this := PropertyTypesCriteria{}
	return &this
}

// GetPropertyTypes returns the PropertyTypes field value if set, zero value otherwise.
func (o *PropertyTypesCriteria) GetPropertyTypes() []PropertyTypeType {
	if o == nil || IsNil(o.PropertyTypes) {
		var ret []PropertyTypeType
		return ret
	}
	return o.PropertyTypes
}

// GetPropertyTypesOk returns a tuple with the PropertyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyTypesCriteria) GetPropertyTypesOk() ([]PropertyTypeType, bool) {
	if o == nil || IsNil(o.PropertyTypes) {
		return nil, false
	}
	return o.PropertyTypes, true
}

// HasPropertyTypes returns a boolean if a field has been set.
func (o *PropertyTypesCriteria) HasPropertyTypes() bool {
	if o != nil && !IsNil(o.PropertyTypes) {
		return true
	}

	return false
}

// SetPropertyTypes gets a reference to the given []PropertyTypeType and assigns it to the PropertyTypes field.
func (o *PropertyTypesCriteria) SetPropertyTypes(v []PropertyTypeType) {
	o.PropertyTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PropertyTypesCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyTypesCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PropertyTypesCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PropertyTypesCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PropertyTypesCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyTypesCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PropertyTypesCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PropertyTypesCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PropertyTypesCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyTypesCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyTypes) {
		toSerialize["propertyTypes"] = o.PropertyTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePropertyTypesCriteria struct {
	value *PropertyTypesCriteria
	isSet bool
}

func (v NullablePropertyTypesCriteria) Get() *PropertyTypesCriteria {
	return v.value
}

func (v *NullablePropertyTypesCriteria) Set(val *PropertyTypesCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyTypesCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyTypesCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyTypesCriteria(val *PropertyTypesCriteria) *NullablePropertyTypesCriteria {
	return &NullablePropertyTypesCriteria{value: val, isSet: true}
}

func (v NullablePropertyTypesCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyTypesCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


