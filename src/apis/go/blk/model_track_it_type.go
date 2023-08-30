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

// checks if the TrackItType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackItType{}

// TrackItType Identifies the kind of Parcel, Baggage, or Lost items or Valet-managed vehicles or services.
type TrackItType struct {
	Type *CodeDescriptionType `json:"type,omitempty"`
	Url *URLType `json:"url,omitempty"`
}

// NewTrackItType instantiates a new TrackItType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackItType() *TrackItType {
	this := TrackItType{}
	return &this
}

// NewTrackItTypeWithDefaults instantiates a new TrackItType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackItTypeWithDefaults() *TrackItType {
	this := TrackItType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TrackItType) GetType() CodeDescriptionType {
	if o == nil || IsNil(o.Type) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackItType) GetTypeOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TrackItType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeDescriptionType and assigns it to the Type field.
func (o *TrackItType) SetType(v CodeDescriptionType) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TrackItType) GetUrl() URLType {
	if o == nil || IsNil(o.Url) {
		var ret URLType
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackItType) GetUrlOk() (*URLType, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TrackItType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given URLType and assigns it to the Url field.
func (o *TrackItType) SetUrl(v URLType) {
	o.Url = &v
}

func (o TrackItType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackItType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableTrackItType struct {
	value *TrackItType
	isSet bool
}

func (v NullableTrackItType) Get() *TrackItType {
	return v.value
}

func (v *NullableTrackItType) Set(val *TrackItType) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackItType) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackItType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackItType(val *TrackItType) *NullableTrackItType {
	return &NullableTrackItType{value: val, isSet: true}
}

func (v NullableTrackItType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackItType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


