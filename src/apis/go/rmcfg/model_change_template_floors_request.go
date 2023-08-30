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

// checks if the ChangeTemplateFloorsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeTemplateFloorsRequest{}

// ChangeTemplateFloorsRequest struct for ChangeTemplateFloorsRequest
type ChangeTemplateFloorsRequest struct {
	// This type holds a collection of floors at the template level.
	TemplateFloors []TemplateFloorType `json:"templateFloors,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChangeTemplateFloorsRequest instantiates a new ChangeTemplateFloorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeTemplateFloorsRequest() *ChangeTemplateFloorsRequest {
	this := ChangeTemplateFloorsRequest{}
	return &this
}

// NewChangeTemplateFloorsRequestWithDefaults instantiates a new ChangeTemplateFloorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeTemplateFloorsRequestWithDefaults() *ChangeTemplateFloorsRequest {
	this := ChangeTemplateFloorsRequest{}
	return &this
}

// GetTemplateFloors returns the TemplateFloors field value if set, zero value otherwise.
func (o *ChangeTemplateFloorsRequest) GetTemplateFloors() []TemplateFloorType {
	if o == nil || IsNil(o.TemplateFloors) {
		var ret []TemplateFloorType
		return ret
	}
	return o.TemplateFloors
}

// GetTemplateFloorsOk returns a tuple with the TemplateFloors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeTemplateFloorsRequest) GetTemplateFloorsOk() ([]TemplateFloorType, bool) {
	if o == nil || IsNil(o.TemplateFloors) {
		return nil, false
	}
	return o.TemplateFloors, true
}

// HasTemplateFloors returns a boolean if a field has been set.
func (o *ChangeTemplateFloorsRequest) HasTemplateFloors() bool {
	if o != nil && !IsNil(o.TemplateFloors) {
		return true
	}

	return false
}

// SetTemplateFloors gets a reference to the given []TemplateFloorType and assigns it to the TemplateFloors field.
func (o *ChangeTemplateFloorsRequest) SetTemplateFloors(v []TemplateFloorType) {
	o.TemplateFloors = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChangeTemplateFloorsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeTemplateFloorsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChangeTemplateFloorsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChangeTemplateFloorsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChangeTemplateFloorsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeTemplateFloorsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChangeTemplateFloorsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChangeTemplateFloorsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChangeTemplateFloorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeTemplateFloorsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateFloors) {
		toSerialize["templateFloors"] = o.TemplateFloors
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChangeTemplateFloorsRequest struct {
	value *ChangeTemplateFloorsRequest
	isSet bool
}

func (v NullableChangeTemplateFloorsRequest) Get() *ChangeTemplateFloorsRequest {
	return v.value
}

func (v *NullableChangeTemplateFloorsRequest) Set(val *ChangeTemplateFloorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeTemplateFloorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeTemplateFloorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeTemplateFloorsRequest(val *ChangeTemplateFloorsRequest) *NullableChangeTemplateFloorsRequest {
	return &NullableChangeTemplateFloorsRequest{value: val, isSet: true}
}

func (v NullableChangeTemplateFloorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeTemplateFloorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


