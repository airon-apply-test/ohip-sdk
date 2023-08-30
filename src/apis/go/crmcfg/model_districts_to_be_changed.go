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

// checks if the DistrictsToBeChanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistrictsToBeChanged{}

// DistrictsToBeChanged Request object for changing Districts.
type DistrictsToBeChanged struct {
	// List of Districts.
	Districts []DistrictType `json:"districts,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewDistrictsToBeChanged instantiates a new DistrictsToBeChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistrictsToBeChanged() *DistrictsToBeChanged {
	this := DistrictsToBeChanged{}
	return &this
}

// NewDistrictsToBeChangedWithDefaults instantiates a new DistrictsToBeChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistrictsToBeChangedWithDefaults() *DistrictsToBeChanged {
	this := DistrictsToBeChanged{}
	return &this
}

// GetDistricts returns the Districts field value if set, zero value otherwise.
func (o *DistrictsToBeChanged) GetDistricts() []DistrictType {
	if o == nil || IsNil(o.Districts) {
		var ret []DistrictType
		return ret
	}
	return o.Districts
}

// GetDistrictsOk returns a tuple with the Districts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistrictsToBeChanged) GetDistrictsOk() ([]DistrictType, bool) {
	if o == nil || IsNil(o.Districts) {
		return nil, false
	}
	return o.Districts, true
}

// HasDistricts returns a boolean if a field has been set.
func (o *DistrictsToBeChanged) HasDistricts() bool {
	if o != nil && !IsNil(o.Districts) {
		return true
	}

	return false
}

// SetDistricts gets a reference to the given []DistrictType and assigns it to the Districts field.
func (o *DistrictsToBeChanged) SetDistricts(v []DistrictType) {
	o.Districts = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DistrictsToBeChanged) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistrictsToBeChanged) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DistrictsToBeChanged) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *DistrictsToBeChanged) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DistrictsToBeChanged) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistrictsToBeChanged) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DistrictsToBeChanged) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *DistrictsToBeChanged) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o DistrictsToBeChanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistrictsToBeChanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Districts) {
		toSerialize["districts"] = o.Districts
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableDistrictsToBeChanged struct {
	value *DistrictsToBeChanged
	isSet bool
}

func (v NullableDistrictsToBeChanged) Get() *DistrictsToBeChanged {
	return v.value
}

func (v *NullableDistrictsToBeChanged) Set(val *DistrictsToBeChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableDistrictsToBeChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableDistrictsToBeChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistrictsToBeChanged(val *DistrictsToBeChanged) *NullableDistrictsToBeChanged {
	return &NullableDistrictsToBeChanged{value: val, isSet: true}
}

func (v NullableDistrictsToBeChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistrictsToBeChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


