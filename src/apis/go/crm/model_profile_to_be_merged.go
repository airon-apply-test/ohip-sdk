/*
OPERA Cloud Customer Relationship Management API

APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm

import (
	"encoding/json"
)

// checks if the ProfileToBeMerged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileToBeMerged{}

// ProfileToBeMerged Profile Resource will be merged.
type ProfileToBeMerged struct {
	// Unique Id that references an object uniquely in the system.
	VictimProfileId []UniqueIDType `json:"victimProfileId,omitempty"`
	ProfileDetails *ProfileType `json:"profileDetails,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewProfileToBeMerged instantiates a new ProfileToBeMerged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileToBeMerged() *ProfileToBeMerged {
	this := ProfileToBeMerged{}
	return &this
}

// NewProfileToBeMergedWithDefaults instantiates a new ProfileToBeMerged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileToBeMergedWithDefaults() *ProfileToBeMerged {
	this := ProfileToBeMerged{}
	return &this
}

// GetVictimProfileId returns the VictimProfileId field value if set, zero value otherwise.
func (o *ProfileToBeMerged) GetVictimProfileId() []UniqueIDType {
	if o == nil || IsNil(o.VictimProfileId) {
		var ret []UniqueIDType
		return ret
	}
	return o.VictimProfileId
}

// GetVictimProfileIdOk returns a tuple with the VictimProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileToBeMerged) GetVictimProfileIdOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.VictimProfileId) {
		return nil, false
	}
	return o.VictimProfileId, true
}

// HasVictimProfileId returns a boolean if a field has been set.
func (o *ProfileToBeMerged) HasVictimProfileId() bool {
	if o != nil && !IsNil(o.VictimProfileId) {
		return true
	}

	return false
}

// SetVictimProfileId gets a reference to the given []UniqueIDType and assigns it to the VictimProfileId field.
func (o *ProfileToBeMerged) SetVictimProfileId(v []UniqueIDType) {
	o.VictimProfileId = v
}

// GetProfileDetails returns the ProfileDetails field value if set, zero value otherwise.
func (o *ProfileToBeMerged) GetProfileDetails() ProfileType {
	if o == nil || IsNil(o.ProfileDetails) {
		var ret ProfileType
		return ret
	}
	return *o.ProfileDetails
}

// GetProfileDetailsOk returns a tuple with the ProfileDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileToBeMerged) GetProfileDetailsOk() (*ProfileType, bool) {
	if o == nil || IsNil(o.ProfileDetails) {
		return nil, false
	}
	return o.ProfileDetails, true
}

// HasProfileDetails returns a boolean if a field has been set.
func (o *ProfileToBeMerged) HasProfileDetails() bool {
	if o != nil && !IsNil(o.ProfileDetails) {
		return true
	}

	return false
}

// SetProfileDetails gets a reference to the given ProfileType and assigns it to the ProfileDetails field.
func (o *ProfileToBeMerged) SetProfileDetails(v ProfileType) {
	o.ProfileDetails = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProfileToBeMerged) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileToBeMerged) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileToBeMerged) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ProfileToBeMerged) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ProfileToBeMerged) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileToBeMerged) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ProfileToBeMerged) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ProfileToBeMerged) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ProfileToBeMerged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileToBeMerged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VictimProfileId) {
		toSerialize["victimProfileId"] = o.VictimProfileId
	}
	if !IsNil(o.ProfileDetails) {
		toSerialize["profileDetails"] = o.ProfileDetails
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableProfileToBeMerged struct {
	value *ProfileToBeMerged
	isSet bool
}

func (v NullableProfileToBeMerged) Get() *ProfileToBeMerged {
	return v.value
}

func (v *NullableProfileToBeMerged) Set(val *ProfileToBeMerged) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileToBeMerged) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileToBeMerged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileToBeMerged(val *ProfileToBeMerged) *NullableProfileToBeMerged {
	return &NullableProfileToBeMerged{value: val, isSet: true}
}

func (v NullableProfileToBeMerged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileToBeMerged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


