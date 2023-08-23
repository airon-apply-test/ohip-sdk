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

// checks if the VIPLevelsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VIPLevelsDetails{}

// VIPLevelsDetails Response object for fetching V I P Levels.
type VIPLevelsDetails struct {
	// List of V I P Levels.
	VIPLevels []VIPLevelType `json:"vIPLevels,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewVIPLevelsDetails instantiates a new VIPLevelsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVIPLevelsDetails() *VIPLevelsDetails {
	this := VIPLevelsDetails{}
	return &this
}

// NewVIPLevelsDetailsWithDefaults instantiates a new VIPLevelsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVIPLevelsDetailsWithDefaults() *VIPLevelsDetails {
	this := VIPLevelsDetails{}
	return &this
}

// GetVIPLevels returns the VIPLevels field value if set, zero value otherwise.
func (o *VIPLevelsDetails) GetVIPLevels() []VIPLevelType {
	if o == nil || IsNil(o.VIPLevels) {
		var ret []VIPLevelType
		return ret
	}
	return o.VIPLevels
}

// GetVIPLevelsOk returns a tuple with the VIPLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VIPLevelsDetails) GetVIPLevelsOk() ([]VIPLevelType, bool) {
	if o == nil || IsNil(o.VIPLevels) {
		return nil, false
	}
	return o.VIPLevels, true
}

// HasVIPLevels returns a boolean if a field has been set.
func (o *VIPLevelsDetails) HasVIPLevels() bool {
	if o != nil && !IsNil(o.VIPLevels) {
		return true
	}

	return false
}

// SetVIPLevels gets a reference to the given []VIPLevelType and assigns it to the VIPLevels field.
func (o *VIPLevelsDetails) SetVIPLevels(v []VIPLevelType) {
	o.VIPLevels = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VIPLevelsDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VIPLevelsDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VIPLevelsDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *VIPLevelsDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *VIPLevelsDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VIPLevelsDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *VIPLevelsDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *VIPLevelsDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o VIPLevelsDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VIPLevelsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VIPLevels) {
		toSerialize["vIPLevels"] = o.VIPLevels
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableVIPLevelsDetails struct {
	value *VIPLevelsDetails
	isSet bool
}

func (v NullableVIPLevelsDetails) Get() *VIPLevelsDetails {
	return v.value
}

func (v *NullableVIPLevelsDetails) Set(val *VIPLevelsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableVIPLevelsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableVIPLevelsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVIPLevelsDetails(val *VIPLevelsDetails) *NullableVIPLevelsDetails {
	return &NullableVIPLevelsDetails{value: val, isSet: true}
}

func (v NullableVIPLevelsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVIPLevelsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


