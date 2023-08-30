/*
OPERA Cloud Leisure Management API

APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lms

import (
	"encoding/json"
)

// checks if the CreateActivityLocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateActivityLocations{}

// CreateActivityLocations Request object for creating Activity Locations.
type CreateActivityLocations struct {
	// Collection of Activity Locations.
	ActivityLocations []ActivityLocationType `json:"activityLocations,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCreateActivityLocations instantiates a new CreateActivityLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateActivityLocations() *CreateActivityLocations {
	this := CreateActivityLocations{}
	return &this
}

// NewCreateActivityLocationsWithDefaults instantiates a new CreateActivityLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateActivityLocationsWithDefaults() *CreateActivityLocations {
	this := CreateActivityLocations{}
	return &this
}

// GetActivityLocations returns the ActivityLocations field value if set, zero value otherwise.
func (o *CreateActivityLocations) GetActivityLocations() []ActivityLocationType {
	if o == nil || IsNil(o.ActivityLocations) {
		var ret []ActivityLocationType
		return ret
	}
	return o.ActivityLocations
}

// GetActivityLocationsOk returns a tuple with the ActivityLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActivityLocations) GetActivityLocationsOk() ([]ActivityLocationType, bool) {
	if o == nil || IsNil(o.ActivityLocations) {
		return nil, false
	}
	return o.ActivityLocations, true
}

// HasActivityLocations returns a boolean if a field has been set.
func (o *CreateActivityLocations) HasActivityLocations() bool {
	if o != nil && !IsNil(o.ActivityLocations) {
		return true
	}

	return false
}

// SetActivityLocations gets a reference to the given []ActivityLocationType and assigns it to the ActivityLocations field.
func (o *CreateActivityLocations) SetActivityLocations(v []ActivityLocationType) {
	o.ActivityLocations = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateActivityLocations) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActivityLocations) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateActivityLocations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CreateActivityLocations) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreateActivityLocations) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActivityLocations) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreateActivityLocations) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CreateActivityLocations) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CreateActivityLocations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateActivityLocations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityLocations) {
		toSerialize["activityLocations"] = o.ActivityLocations
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCreateActivityLocations struct {
	value *CreateActivityLocations
	isSet bool
}

func (v NullableCreateActivityLocations) Get() *CreateActivityLocations {
	return v.value
}

func (v *NullableCreateActivityLocations) Set(val *CreateActivityLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateActivityLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateActivityLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateActivityLocations(val *CreateActivityLocations) *NullableCreateActivityLocations {
	return &NullableCreateActivityLocations{value: val, isSet: true}
}

func (v NullableCreateActivityLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateActivityLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


