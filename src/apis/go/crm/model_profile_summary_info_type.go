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

// checks if the ProfileSummaryInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileSummaryInfoType{}

// ProfileSummaryInfoType Summary information about a profile and the associated Unique IDs.
type ProfileSummaryInfoType struct {
	// Unique Id that references an object uniquely in the system.
	ProfileIdList []UniqueIDType `json:"profileIdList,omitempty"`
	Profile *ProfileSummaryType `json:"profile,omitempty"`
}

// NewProfileSummaryInfoType instantiates a new ProfileSummaryInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileSummaryInfoType() *ProfileSummaryInfoType {
	this := ProfileSummaryInfoType{}
	return &this
}

// NewProfileSummaryInfoTypeWithDefaults instantiates a new ProfileSummaryInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileSummaryInfoTypeWithDefaults() *ProfileSummaryInfoType {
	this := ProfileSummaryInfoType{}
	return &this
}

// GetProfileIdList returns the ProfileIdList field value if set, zero value otherwise.
func (o *ProfileSummaryInfoType) GetProfileIdList() []UniqueIDType {
	if o == nil || IsNil(o.ProfileIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ProfileIdList
}

// GetProfileIdListOk returns a tuple with the ProfileIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryInfoType) GetProfileIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileIdList) {
		return nil, false
	}
	return o.ProfileIdList, true
}

// HasProfileIdList returns a boolean if a field has been set.
func (o *ProfileSummaryInfoType) HasProfileIdList() bool {
	if o != nil && !IsNil(o.ProfileIdList) {
		return true
	}

	return false
}

// SetProfileIdList gets a reference to the given []UniqueIDType and assigns it to the ProfileIdList field.
func (o *ProfileSummaryInfoType) SetProfileIdList(v []UniqueIDType) {
	o.ProfileIdList = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ProfileSummaryInfoType) GetProfile() ProfileSummaryType {
	if o == nil || IsNil(o.Profile) {
		var ret ProfileSummaryType
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryInfoType) GetProfileOk() (*ProfileSummaryType, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ProfileSummaryInfoType) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ProfileSummaryType and assigns it to the Profile field.
func (o *ProfileSummaryInfoType) SetProfile(v ProfileSummaryType) {
	o.Profile = &v
}

func (o ProfileSummaryInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileSummaryInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileIdList) {
		toSerialize["profileIdList"] = o.ProfileIdList
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	return toSerialize, nil
}

type NullableProfileSummaryInfoType struct {
	value *ProfileSummaryInfoType
	isSet bool
}

func (v NullableProfileSummaryInfoType) Get() *ProfileSummaryInfoType {
	return v.value
}

func (v *NullableProfileSummaryInfoType) Set(val *ProfileSummaryInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSummaryInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSummaryInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSummaryInfoType(val *ProfileSummaryInfoType) *NullableProfileSummaryInfoType {
	return &NullableProfileSummaryInfoType{value: val, isSet: true}
}

func (v NullableProfileSummaryInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSummaryInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


