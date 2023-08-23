/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SourcesCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourcesCriteria{}

// SourcesCriteria Request object for creating Sources.
type SourcesCriteria struct {
	// List of Sources.
	Sources []SourceType `json:"sources,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewSourcesCriteria instantiates a new SourcesCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcesCriteria() *SourcesCriteria {
	this := SourcesCriteria{}
	return &this
}

// NewSourcesCriteriaWithDefaults instantiates a new SourcesCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcesCriteriaWithDefaults() *SourcesCriteria {
	this := SourcesCriteria{}
	return &this
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *SourcesCriteria) GetSources() []SourceType {
	if o == nil || IsNil(o.Sources) {
		var ret []SourceType
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesCriteria) GetSourcesOk() ([]SourceType, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *SourcesCriteria) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []SourceType and assigns it to the Sources field.
func (o *SourcesCriteria) SetSources(v []SourceType) {
	o.Sources = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SourcesCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SourcesCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *SourcesCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SourcesCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SourcesCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *SourcesCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o SourcesCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourcesCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableSourcesCriteria struct {
	value *SourcesCriteria
	isSet bool
}

func (v NullableSourcesCriteria) Get() *SourcesCriteria {
	return v.value
}

func (v *NullableSourcesCriteria) Set(val *SourcesCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcesCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcesCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcesCriteria(val *SourcesCriteria) *NullableSourcesCriteria {
	return &NullableSourcesCriteria{value: val, isSet: true}
}

func (v NullableSourcesCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcesCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


