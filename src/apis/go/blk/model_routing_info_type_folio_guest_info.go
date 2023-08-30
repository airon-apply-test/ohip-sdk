/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the RoutingInfoTypeFolioGuestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingInfoTypeFolioGuestInfo{}

// RoutingInfoTypeFolioGuestInfo Guest details
type RoutingInfoTypeFolioGuestInfo struct {
	// Unique Id that references an object uniquely in the system.
	ProfileIdList []UniqueIDType `json:"profileIdList,omitempty"`
}

// NewRoutingInfoTypeFolioGuestInfo instantiates a new RoutingInfoTypeFolioGuestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingInfoTypeFolioGuestInfo() *RoutingInfoTypeFolioGuestInfo {
	this := RoutingInfoTypeFolioGuestInfo{}
	return &this
}

// NewRoutingInfoTypeFolioGuestInfoWithDefaults instantiates a new RoutingInfoTypeFolioGuestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingInfoTypeFolioGuestInfoWithDefaults() *RoutingInfoTypeFolioGuestInfo {
	this := RoutingInfoTypeFolioGuestInfo{}
	return &this
}

// GetProfileIdList returns the ProfileIdList field value if set, zero value otherwise.
func (o *RoutingInfoTypeFolioGuestInfo) GetProfileIdList() []UniqueIDType {
	if o == nil || IsNil(o.ProfileIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ProfileIdList
}

// GetProfileIdListOk returns a tuple with the ProfileIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoTypeFolioGuestInfo) GetProfileIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileIdList) {
		return nil, false
	}
	return o.ProfileIdList, true
}

// HasProfileIdList returns a boolean if a field has been set.
func (o *RoutingInfoTypeFolioGuestInfo) HasProfileIdList() bool {
	if o != nil && !IsNil(o.ProfileIdList) {
		return true
	}

	return false
}

// SetProfileIdList gets a reference to the given []UniqueIDType and assigns it to the ProfileIdList field.
func (o *RoutingInfoTypeFolioGuestInfo) SetProfileIdList(v []UniqueIDType) {
	o.ProfileIdList = v
}

func (o RoutingInfoTypeFolioGuestInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingInfoTypeFolioGuestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileIdList) {
		toSerialize["profileIdList"] = o.ProfileIdList
	}
	return toSerialize, nil
}

type NullableRoutingInfoTypeFolioGuestInfo struct {
	value *RoutingInfoTypeFolioGuestInfo
	isSet bool
}

func (v NullableRoutingInfoTypeFolioGuestInfo) Get() *RoutingInfoTypeFolioGuestInfo {
	return v.value
}

func (v *NullableRoutingInfoTypeFolioGuestInfo) Set(val *RoutingInfoTypeFolioGuestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingInfoTypeFolioGuestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingInfoTypeFolioGuestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingInfoTypeFolioGuestInfo(val *RoutingInfoTypeFolioGuestInfo) *NullableRoutingInfoTypeFolioGuestInfo {
	return &NullableRoutingInfoTypeFolioGuestInfo{value: val, isSet: true}
}

func (v NullableRoutingInfoTypeFolioGuestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingInfoTypeFolioGuestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


