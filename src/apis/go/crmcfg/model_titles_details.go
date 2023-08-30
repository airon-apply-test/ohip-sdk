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

// checks if the TitlesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TitlesDetails{}

// TitlesDetails Response object for fetching Titles.
type TitlesDetails struct {
	// List of Titles.
	Titles []TitleType `json:"titles,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewTitlesDetails instantiates a new TitlesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTitlesDetails() *TitlesDetails {
	this := TitlesDetails{}
	return &this
}

// NewTitlesDetailsWithDefaults instantiates a new TitlesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTitlesDetailsWithDefaults() *TitlesDetails {
	this := TitlesDetails{}
	return &this
}

// GetTitles returns the Titles field value if set, zero value otherwise.
func (o *TitlesDetails) GetTitles() []TitleType {
	if o == nil || IsNil(o.Titles) {
		var ret []TitleType
		return ret
	}
	return o.Titles
}

// GetTitlesOk returns a tuple with the Titles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TitlesDetails) GetTitlesOk() ([]TitleType, bool) {
	if o == nil || IsNil(o.Titles) {
		return nil, false
	}
	return o.Titles, true
}

// HasTitles returns a boolean if a field has been set.
func (o *TitlesDetails) HasTitles() bool {
	if o != nil && !IsNil(o.Titles) {
		return true
	}

	return false
}

// SetTitles gets a reference to the given []TitleType and assigns it to the Titles field.
func (o *TitlesDetails) SetTitles(v []TitleType) {
	o.Titles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TitlesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TitlesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TitlesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *TitlesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *TitlesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TitlesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TitlesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *TitlesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o TitlesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TitlesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Titles) {
		toSerialize["titles"] = o.Titles
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableTitlesDetails struct {
	value *TitlesDetails
	isSet bool
}

func (v NullableTitlesDetails) Get() *TitlesDetails {
	return v.value
}

func (v *NullableTitlesDetails) Set(val *TitlesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTitlesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTitlesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTitlesDetails(val *TitlesDetails) *NullableTitlesDetails {
	return &NullableTitlesDetails{value: val, isSet: true}
}

func (v NullableTitlesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTitlesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


