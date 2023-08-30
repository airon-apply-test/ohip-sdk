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

// checks if the StagedProfileId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StagedProfileId{}

// StagedProfileId Unique identifier for a staged profile.
type StagedProfileId struct {
	// Hotel code for the staged profile to be reprocessed.
	HotelId *string `json:"hotelId,omitempty"`
	ProfileId *ProfileId `json:"profileId,omitempty"`
}

// NewStagedProfileId instantiates a new StagedProfileId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagedProfileId() *StagedProfileId {
	this := StagedProfileId{}
	return &this
}

// NewStagedProfileIdWithDefaults instantiates a new StagedProfileId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagedProfileIdWithDefaults() *StagedProfileId {
	this := StagedProfileId{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *StagedProfileId) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileId) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *StagedProfileId) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *StagedProfileId) SetHotelId(v string) {
	o.HotelId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *StagedProfileId) GetProfileId() ProfileId {
	if o == nil || IsNil(o.ProfileId) {
		var ret ProfileId
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileId) GetProfileIdOk() (*ProfileId, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *StagedProfileId) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given ProfileId and assigns it to the ProfileId field.
func (o *StagedProfileId) SetProfileId(v ProfileId) {
	o.ProfileId = &v
}

func (o StagedProfileId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StagedProfileId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return toSerialize, nil
}

type NullableStagedProfileId struct {
	value *StagedProfileId
	isSet bool
}

func (v NullableStagedProfileId) Get() *StagedProfileId {
	return v.value
}

func (v *NullableStagedProfileId) Set(val *StagedProfileId) {
	v.value = val
	v.isSet = true
}

func (v NullableStagedProfileId) IsSet() bool {
	return v.isSet
}

func (v *NullableStagedProfileId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagedProfileId(val *StagedProfileId) *NullableStagedProfileId {
	return &NullableStagedProfileId{value: val, isSet: true}
}

func (v NullableStagedProfileId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagedProfileId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


