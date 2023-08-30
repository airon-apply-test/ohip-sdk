/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmcfg

import (
	"encoding/json"
)

// checks if the PreferenceGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreferenceGroups{}

// PreferenceGroups Response object for fetching preference groups.
type PreferenceGroups struct {
	// Collection of preference groups.
	PreferenceGroups []PreferenceGroupType `json:"preferenceGroups,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPreferenceGroups instantiates a new PreferenceGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferenceGroups() *PreferenceGroups {
	this := PreferenceGroups{}
	return &this
}

// NewPreferenceGroupsWithDefaults instantiates a new PreferenceGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferenceGroupsWithDefaults() *PreferenceGroups {
	this := PreferenceGroups{}
	return &this
}

// GetPreferenceGroups returns the PreferenceGroups field value if set, zero value otherwise.
func (o *PreferenceGroups) GetPreferenceGroups() []PreferenceGroupType {
	if o == nil || IsNil(o.PreferenceGroups) {
		var ret []PreferenceGroupType
		return ret
	}
	return o.PreferenceGroups
}

// GetPreferenceGroupsOk returns a tuple with the PreferenceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceGroups) GetPreferenceGroupsOk() ([]PreferenceGroupType, bool) {
	if o == nil || IsNil(o.PreferenceGroups) {
		return nil, false
	}
	return o.PreferenceGroups, true
}

// HasPreferenceGroups returns a boolean if a field has been set.
func (o *PreferenceGroups) HasPreferenceGroups() bool {
	if o != nil && !IsNil(o.PreferenceGroups) {
		return true
	}

	return false
}

// SetPreferenceGroups gets a reference to the given []PreferenceGroupType and assigns it to the PreferenceGroups field.
func (o *PreferenceGroups) SetPreferenceGroups(v []PreferenceGroupType) {
	o.PreferenceGroups = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PreferenceGroups) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceGroups) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PreferenceGroups) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PreferenceGroups) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PreferenceGroups) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceGroups) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PreferenceGroups) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PreferenceGroups) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PreferenceGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreferenceGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreferenceGroups) {
		toSerialize["preferenceGroups"] = o.PreferenceGroups
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePreferenceGroups struct {
	value *PreferenceGroups
	isSet bool
}

func (v NullablePreferenceGroups) Get() *PreferenceGroups {
	return v.value
}

func (v *NullablePreferenceGroups) Set(val *PreferenceGroups) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferenceGroups) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferenceGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferenceGroups(val *PreferenceGroups) *NullablePreferenceGroups {
	return &NullablePreferenceGroups{value: val, isSet: true}
}

func (v NullablePreferenceGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferenceGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


