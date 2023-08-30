/*
OPERA Cloud Room Configuration API

APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rmcfg

import (
	"encoding/json"
)

// checks if the ChangeHousekeepingSectionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeHousekeepingSectionsRequest{}

// ChangeHousekeepingSectionsRequest struct for ChangeHousekeepingSectionsRequest
type ChangeHousekeepingSectionsRequest struct {
	// List of the housekeeping sections to be configured or fetched
	HousekeepingSections []HousekeepingSectionType `json:"housekeepingSections,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChangeHousekeepingSectionsRequest instantiates a new ChangeHousekeepingSectionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeHousekeepingSectionsRequest() *ChangeHousekeepingSectionsRequest {
	this := ChangeHousekeepingSectionsRequest{}
	return &this
}

// NewChangeHousekeepingSectionsRequestWithDefaults instantiates a new ChangeHousekeepingSectionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeHousekeepingSectionsRequestWithDefaults() *ChangeHousekeepingSectionsRequest {
	this := ChangeHousekeepingSectionsRequest{}
	return &this
}

// GetHousekeepingSections returns the HousekeepingSections field value if set, zero value otherwise.
func (o *ChangeHousekeepingSectionsRequest) GetHousekeepingSections() []HousekeepingSectionType {
	if o == nil || IsNil(o.HousekeepingSections) {
		var ret []HousekeepingSectionType
		return ret
	}
	return o.HousekeepingSections
}

// GetHousekeepingSectionsOk returns a tuple with the HousekeepingSections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeHousekeepingSectionsRequest) GetHousekeepingSectionsOk() ([]HousekeepingSectionType, bool) {
	if o == nil || IsNil(o.HousekeepingSections) {
		return nil, false
	}
	return o.HousekeepingSections, true
}

// HasHousekeepingSections returns a boolean if a field has been set.
func (o *ChangeHousekeepingSectionsRequest) HasHousekeepingSections() bool {
	if o != nil && !IsNil(o.HousekeepingSections) {
		return true
	}

	return false
}

// SetHousekeepingSections gets a reference to the given []HousekeepingSectionType and assigns it to the HousekeepingSections field.
func (o *ChangeHousekeepingSectionsRequest) SetHousekeepingSections(v []HousekeepingSectionType) {
	o.HousekeepingSections = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChangeHousekeepingSectionsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeHousekeepingSectionsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChangeHousekeepingSectionsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChangeHousekeepingSectionsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChangeHousekeepingSectionsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeHousekeepingSectionsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChangeHousekeepingSectionsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChangeHousekeepingSectionsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChangeHousekeepingSectionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeHousekeepingSectionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HousekeepingSections) {
		toSerialize["housekeepingSections"] = o.HousekeepingSections
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChangeHousekeepingSectionsRequest struct {
	value *ChangeHousekeepingSectionsRequest
	isSet bool
}

func (v NullableChangeHousekeepingSectionsRequest) Get() *ChangeHousekeepingSectionsRequest {
	return v.value
}

func (v *NullableChangeHousekeepingSectionsRequest) Set(val *ChangeHousekeepingSectionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeHousekeepingSectionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeHousekeepingSectionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeHousekeepingSectionsRequest(val *ChangeHousekeepingSectionsRequest) *NullableChangeHousekeepingSectionsRequest {
	return &NullableChangeHousekeepingSectionsRequest{value: val, isSet: true}
}

func (v NullableChangeHousekeepingSectionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeHousekeepingSectionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


