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

// checks if the AlternateLanguageGuestTitlesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternateLanguageGuestTitlesDetails{}

// AlternateLanguageGuestTitlesDetails Response object for fetching Alternate Language Guest Titles.
type AlternateLanguageGuestTitlesDetails struct {
	// List of Guest Titles.
	AlternateLanguageGuestTitles []GuestTitleType `json:"alternateLanguageGuestTitles,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewAlternateLanguageGuestTitlesDetails instantiates a new AlternateLanguageGuestTitlesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternateLanguageGuestTitlesDetails() *AlternateLanguageGuestTitlesDetails {
	this := AlternateLanguageGuestTitlesDetails{}
	return &this
}

// NewAlternateLanguageGuestTitlesDetailsWithDefaults instantiates a new AlternateLanguageGuestTitlesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternateLanguageGuestTitlesDetailsWithDefaults() *AlternateLanguageGuestTitlesDetails {
	this := AlternateLanguageGuestTitlesDetails{}
	return &this
}

// GetAlternateLanguageGuestTitles returns the AlternateLanguageGuestTitles field value if set, zero value otherwise.
func (o *AlternateLanguageGuestTitlesDetails) GetAlternateLanguageGuestTitles() []GuestTitleType {
	if o == nil || IsNil(o.AlternateLanguageGuestTitles) {
		var ret []GuestTitleType
		return ret
	}
	return o.AlternateLanguageGuestTitles
}

// GetAlternateLanguageGuestTitlesOk returns a tuple with the AlternateLanguageGuestTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternateLanguageGuestTitlesDetails) GetAlternateLanguageGuestTitlesOk() ([]GuestTitleType, bool) {
	if o == nil || IsNil(o.AlternateLanguageGuestTitles) {
		return nil, false
	}
	return o.AlternateLanguageGuestTitles, true
}

// HasAlternateLanguageGuestTitles returns a boolean if a field has been set.
func (o *AlternateLanguageGuestTitlesDetails) HasAlternateLanguageGuestTitles() bool {
	if o != nil && !IsNil(o.AlternateLanguageGuestTitles) {
		return true
	}

	return false
}

// SetAlternateLanguageGuestTitles gets a reference to the given []GuestTitleType and assigns it to the AlternateLanguageGuestTitles field.
func (o *AlternateLanguageGuestTitlesDetails) SetAlternateLanguageGuestTitles(v []GuestTitleType) {
	o.AlternateLanguageGuestTitles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlternateLanguageGuestTitlesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternateLanguageGuestTitlesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlternateLanguageGuestTitlesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *AlternateLanguageGuestTitlesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AlternateLanguageGuestTitlesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternateLanguageGuestTitlesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AlternateLanguageGuestTitlesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *AlternateLanguageGuestTitlesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o AlternateLanguageGuestTitlesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternateLanguageGuestTitlesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlternateLanguageGuestTitles) {
		toSerialize["alternateLanguageGuestTitles"] = o.AlternateLanguageGuestTitles
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAlternateLanguageGuestTitlesDetails struct {
	value *AlternateLanguageGuestTitlesDetails
	isSet bool
}

func (v NullableAlternateLanguageGuestTitlesDetails) Get() *AlternateLanguageGuestTitlesDetails {
	return v.value
}

func (v *NullableAlternateLanguageGuestTitlesDetails) Set(val *AlternateLanguageGuestTitlesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternateLanguageGuestTitlesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternateLanguageGuestTitlesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternateLanguageGuestTitlesDetails(val *AlternateLanguageGuestTitlesDetails) *NullableAlternateLanguageGuestTitlesDetails {
	return &NullableAlternateLanguageGuestTitlesDetails{value: val, isSet: true}
}

func (v NullableAlternateLanguageGuestTitlesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternateLanguageGuestTitlesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


