/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProfileCommissionAccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileCommissionAccountInfo{}

// ProfileCommissionAccountInfo Response object for fetching profile commission detail.
type ProfileCommissionAccountInfo struct {
	// Profile commission info which contains bank account and commission code details.
	ProfileCommissionAccountInfoList []ProfileCommissionAccountInfoType `json:"profileCommissionAccountInfoList,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
}

// NewProfileCommissionAccountInfo instantiates a new ProfileCommissionAccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileCommissionAccountInfo() *ProfileCommissionAccountInfo {
	this := ProfileCommissionAccountInfo{}
	return &this
}

// NewProfileCommissionAccountInfoWithDefaults instantiates a new ProfileCommissionAccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileCommissionAccountInfoWithDefaults() *ProfileCommissionAccountInfo {
	this := ProfileCommissionAccountInfo{}
	return &this
}

// GetProfileCommissionAccountInfoList returns the ProfileCommissionAccountInfoList field value if set, zero value otherwise.
func (o *ProfileCommissionAccountInfo) GetProfileCommissionAccountInfoList() []ProfileCommissionAccountInfoType {
	if o == nil || IsNil(o.ProfileCommissionAccountInfoList) {
		var ret []ProfileCommissionAccountInfoType
		return ret
	}
	return o.ProfileCommissionAccountInfoList
}

// GetProfileCommissionAccountInfoListOk returns a tuple with the ProfileCommissionAccountInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCommissionAccountInfo) GetProfileCommissionAccountInfoListOk() ([]ProfileCommissionAccountInfoType, bool) {
	if o == nil || IsNil(o.ProfileCommissionAccountInfoList) {
		return nil, false
	}
	return o.ProfileCommissionAccountInfoList, true
}

// HasProfileCommissionAccountInfoList returns a boolean if a field has been set.
func (o *ProfileCommissionAccountInfo) HasProfileCommissionAccountInfoList() bool {
	if o != nil && !IsNil(o.ProfileCommissionAccountInfoList) {
		return true
	}

	return false
}

// SetProfileCommissionAccountInfoList gets a reference to the given []ProfileCommissionAccountInfoType and assigns it to the ProfileCommissionAccountInfoList field.
func (o *ProfileCommissionAccountInfo) SetProfileCommissionAccountInfoList(v []ProfileCommissionAccountInfoType) {
	o.ProfileCommissionAccountInfoList = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProfileCommissionAccountInfo) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCommissionAccountInfo) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileCommissionAccountInfo) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ProfileCommissionAccountInfo) SetLinks(v []InstanceLink) {
	o.Links = v
}

func (o ProfileCommissionAccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileCommissionAccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileCommissionAccountInfoList) {
		toSerialize["profileCommissionAccountInfoList"] = o.ProfileCommissionAccountInfoList
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableProfileCommissionAccountInfo struct {
	value *ProfileCommissionAccountInfo
	isSet bool
}

func (v NullableProfileCommissionAccountInfo) Get() *ProfileCommissionAccountInfo {
	return v.value
}

func (v *NullableProfileCommissionAccountInfo) Set(val *ProfileCommissionAccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileCommissionAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileCommissionAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileCommissionAccountInfo(val *ProfileCommissionAccountInfo) *NullableProfileCommissionAccountInfo {
	return &NullableProfileCommissionAccountInfo{value: val, isSet: true}
}

func (v NullableProfileCommissionAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileCommissionAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


