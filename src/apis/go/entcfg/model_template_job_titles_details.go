/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TemplateJobTitlesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateJobTitlesDetails{}

// TemplateJobTitlesDetails Response object for fetching template job titles.
type TemplateJobTitlesDetails struct {
	// Template job title details.
	TemplateJobTitles []TemplateJobTitleType `json:"templateJobTitles,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewTemplateJobTitlesDetails instantiates a new TemplateJobTitlesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateJobTitlesDetails() *TemplateJobTitlesDetails {
	this := TemplateJobTitlesDetails{}
	return &this
}

// NewTemplateJobTitlesDetailsWithDefaults instantiates a new TemplateJobTitlesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateJobTitlesDetailsWithDefaults() *TemplateJobTitlesDetails {
	this := TemplateJobTitlesDetails{}
	return &this
}

// GetTemplateJobTitles returns the TemplateJobTitles field value if set, zero value otherwise.
func (o *TemplateJobTitlesDetails) GetTemplateJobTitles() []TemplateJobTitleType {
	if o == nil || IsNil(o.TemplateJobTitles) {
		var ret []TemplateJobTitleType
		return ret
	}
	return o.TemplateJobTitles
}

// GetTemplateJobTitlesOk returns a tuple with the TemplateJobTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateJobTitlesDetails) GetTemplateJobTitlesOk() ([]TemplateJobTitleType, bool) {
	if o == nil || IsNil(o.TemplateJobTitles) {
		return nil, false
	}
	return o.TemplateJobTitles, true
}

// HasTemplateJobTitles returns a boolean if a field has been set.
func (o *TemplateJobTitlesDetails) HasTemplateJobTitles() bool {
	if o != nil && !IsNil(o.TemplateJobTitles) {
		return true
	}

	return false
}

// SetTemplateJobTitles gets a reference to the given []TemplateJobTitleType and assigns it to the TemplateJobTitles field.
func (o *TemplateJobTitlesDetails) SetTemplateJobTitles(v []TemplateJobTitleType) {
	o.TemplateJobTitles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TemplateJobTitlesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateJobTitlesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TemplateJobTitlesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *TemplateJobTitlesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *TemplateJobTitlesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateJobTitlesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TemplateJobTitlesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *TemplateJobTitlesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o TemplateJobTitlesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateJobTitlesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateJobTitles) {
		toSerialize["templateJobTitles"] = o.TemplateJobTitles
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableTemplateJobTitlesDetails struct {
	value *TemplateJobTitlesDetails
	isSet bool
}

func (v NullableTemplateJobTitlesDetails) Get() *TemplateJobTitlesDetails {
	return v.value
}

func (v *NullableTemplateJobTitlesDetails) Set(val *TemplateJobTitlesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateJobTitlesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateJobTitlesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateJobTitlesDetails(val *TemplateJobTitlesDetails) *NullableTemplateJobTitlesDetails {
	return &NullableTemplateJobTitlesDetails{value: val, isSet: true}
}

func (v NullableTemplateJobTitlesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateJobTitlesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


