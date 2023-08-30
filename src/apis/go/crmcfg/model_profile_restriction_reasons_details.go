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

// checks if the ProfileRestrictionReasonsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileRestrictionReasonsDetails{}

// ProfileRestrictionReasonsDetails Response object for fetching Profile Restriction Reasons.
type ProfileRestrictionReasonsDetails struct {
	// List of Profile Restriction Reasons.
	ProfileRestrictionReasons []ProfileRestrictionReasonType `json:"profileRestrictionReasons,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewProfileRestrictionReasonsDetails instantiates a new ProfileRestrictionReasonsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileRestrictionReasonsDetails() *ProfileRestrictionReasonsDetails {
	this := ProfileRestrictionReasonsDetails{}
	return &this
}

// NewProfileRestrictionReasonsDetailsWithDefaults instantiates a new ProfileRestrictionReasonsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileRestrictionReasonsDetailsWithDefaults() *ProfileRestrictionReasonsDetails {
	this := ProfileRestrictionReasonsDetails{}
	return &this
}

// GetProfileRestrictionReasons returns the ProfileRestrictionReasons field value if set, zero value otherwise.
func (o *ProfileRestrictionReasonsDetails) GetProfileRestrictionReasons() []ProfileRestrictionReasonType {
	if o == nil || IsNil(o.ProfileRestrictionReasons) {
		var ret []ProfileRestrictionReasonType
		return ret
	}
	return o.ProfileRestrictionReasons
}

// GetProfileRestrictionReasonsOk returns a tuple with the ProfileRestrictionReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRestrictionReasonsDetails) GetProfileRestrictionReasonsOk() ([]ProfileRestrictionReasonType, bool) {
	if o == nil || IsNil(o.ProfileRestrictionReasons) {
		return nil, false
	}
	return o.ProfileRestrictionReasons, true
}

// HasProfileRestrictionReasons returns a boolean if a field has been set.
func (o *ProfileRestrictionReasonsDetails) HasProfileRestrictionReasons() bool {
	if o != nil && !IsNil(o.ProfileRestrictionReasons) {
		return true
	}

	return false
}

// SetProfileRestrictionReasons gets a reference to the given []ProfileRestrictionReasonType and assigns it to the ProfileRestrictionReasons field.
func (o *ProfileRestrictionReasonsDetails) SetProfileRestrictionReasons(v []ProfileRestrictionReasonType) {
	o.ProfileRestrictionReasons = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProfileRestrictionReasonsDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRestrictionReasonsDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileRestrictionReasonsDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ProfileRestrictionReasonsDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ProfileRestrictionReasonsDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRestrictionReasonsDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ProfileRestrictionReasonsDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ProfileRestrictionReasonsDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ProfileRestrictionReasonsDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileRestrictionReasonsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileRestrictionReasons) {
		toSerialize["profileRestrictionReasons"] = o.ProfileRestrictionReasons
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableProfileRestrictionReasonsDetails struct {
	value *ProfileRestrictionReasonsDetails
	isSet bool
}

func (v NullableProfileRestrictionReasonsDetails) Get() *ProfileRestrictionReasonsDetails {
	return v.value
}

func (v *NullableProfileRestrictionReasonsDetails) Set(val *ProfileRestrictionReasonsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileRestrictionReasonsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileRestrictionReasonsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileRestrictionReasonsDetails(val *ProfileRestrictionReasonsDetails) *NullableProfileRestrictionReasonsDetails {
	return &NullableProfileRestrictionReasonsDetails{value: val, isSet: true}
}

func (v NullableProfileRestrictionReasonsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileRestrictionReasonsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


