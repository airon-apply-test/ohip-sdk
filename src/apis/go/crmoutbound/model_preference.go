/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the Preference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Preference{}

// Preference Response object for fetch preference. This object contains collection of preferences,Success,Warnings and Errors related to this operation.
type Preference struct {
	// Detailed information of preferences related to the profile
	PreferenceCollections []PreferenceTypeType `json:"preferenceCollections,omitempty"`
	// Unique Id that references an object uniquely in the system.
	ProfileIdList []UniqueIDType `json:"profileIdList,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
}

// NewPreference instantiates a new Preference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreference() *Preference {
	this := Preference{}
	return &this
}

// NewPreferenceWithDefaults instantiates a new Preference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferenceWithDefaults() *Preference {
	this := Preference{}
	return &this
}

// GetPreferenceCollections returns the PreferenceCollections field value if set, zero value otherwise.
func (o *Preference) GetPreferenceCollections() []PreferenceTypeType {
	if o == nil || IsNil(o.PreferenceCollections) {
		var ret []PreferenceTypeType
		return ret
	}
	return o.PreferenceCollections
}

// GetPreferenceCollectionsOk returns a tuple with the PreferenceCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preference) GetPreferenceCollectionsOk() ([]PreferenceTypeType, bool) {
	if o == nil || IsNil(o.PreferenceCollections) {
		return nil, false
	}
	return o.PreferenceCollections, true
}

// HasPreferenceCollections returns a boolean if a field has been set.
func (o *Preference) HasPreferenceCollections() bool {
	if o != nil && !IsNil(o.PreferenceCollections) {
		return true
	}

	return false
}

// SetPreferenceCollections gets a reference to the given []PreferenceTypeType and assigns it to the PreferenceCollections field.
func (o *Preference) SetPreferenceCollections(v []PreferenceTypeType) {
	o.PreferenceCollections = v
}

// GetProfileIdList returns the ProfileIdList field value if set, zero value otherwise.
func (o *Preference) GetProfileIdList() []UniqueIDType {
	if o == nil || IsNil(o.ProfileIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ProfileIdList
}

// GetProfileIdListOk returns a tuple with the ProfileIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preference) GetProfileIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileIdList) {
		return nil, false
	}
	return o.ProfileIdList, true
}

// HasProfileIdList returns a boolean if a field has been set.
func (o *Preference) HasProfileIdList() bool {
	if o != nil && !IsNil(o.ProfileIdList) {
		return true
	}

	return false
}

// SetProfileIdList gets a reference to the given []UniqueIDType and assigns it to the ProfileIdList field.
func (o *Preference) SetProfileIdList(v []UniqueIDType) {
	o.ProfileIdList = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Preference) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preference) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Preference) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *Preference) SetLinks(v []InstanceLink) {
	o.Links = v
}

func (o Preference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Preference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreferenceCollections) {
		toSerialize["preferenceCollections"] = o.PreferenceCollections
	}
	if !IsNil(o.ProfileIdList) {
		toSerialize["profileIdList"] = o.ProfileIdList
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullablePreference struct {
	value *Preference
	isSet bool
}

func (v NullablePreference) Get() *Preference {
	return v.value
}

func (v *NullablePreference) Set(val *Preference) {
	v.value = val
	v.isSet = true
}

func (v NullablePreference) IsSet() bool {
	return v.isSet
}

func (v *NullablePreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreference(val *Preference) *NullablePreference {
	return &NullablePreference{value: val, isSet: true}
}

func (v NullablePreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


