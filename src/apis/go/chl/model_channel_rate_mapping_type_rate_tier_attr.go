/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChannelRateMappingTypeRateTierAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelRateMappingTypeRateTierAttr{}

// ChannelRateMappingTypeRateTierAttr Rate Tier.
type ChannelRateMappingTypeRateTierAttr struct {
	RateTier *int32 `json:"rateTier,omitempty"`
}

// NewChannelRateMappingTypeRateTierAttr instantiates a new ChannelRateMappingTypeRateTierAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelRateMappingTypeRateTierAttr() *ChannelRateMappingTypeRateTierAttr {
	this := ChannelRateMappingTypeRateTierAttr{}
	return &this
}

// NewChannelRateMappingTypeRateTierAttrWithDefaults instantiates a new ChannelRateMappingTypeRateTierAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelRateMappingTypeRateTierAttrWithDefaults() *ChannelRateMappingTypeRateTierAttr {
	this := ChannelRateMappingTypeRateTierAttr{}
	return &this
}

// GetRateTier returns the RateTier field value if set, zero value otherwise.
func (o *ChannelRateMappingTypeRateTierAttr) GetRateTier() int32 {
	if o == nil || IsNil(o.RateTier) {
		var ret int32
		return ret
	}
	return *o.RateTier
}

// GetRateTierOk returns a tuple with the RateTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRateMappingTypeRateTierAttr) GetRateTierOk() (*int32, bool) {
	if o == nil || IsNil(o.RateTier) {
		return nil, false
	}
	return o.RateTier, true
}

// HasRateTier returns a boolean if a field has been set.
func (o *ChannelRateMappingTypeRateTierAttr) HasRateTier() bool {
	if o != nil && !IsNil(o.RateTier) {
		return true
	}

	return false
}

// SetRateTier gets a reference to the given int32 and assigns it to the RateTier field.
func (o *ChannelRateMappingTypeRateTierAttr) SetRateTier(v int32) {
	o.RateTier = &v
}

func (o ChannelRateMappingTypeRateTierAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelRateMappingTypeRateTierAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RateTier) {
		toSerialize["rateTier"] = o.RateTier
	}
	return toSerialize, nil
}

type NullableChannelRateMappingTypeRateTierAttr struct {
	value *ChannelRateMappingTypeRateTierAttr
	isSet bool
}

func (v NullableChannelRateMappingTypeRateTierAttr) Get() *ChannelRateMappingTypeRateTierAttr {
	return v.value
}

func (v *NullableChannelRateMappingTypeRateTierAttr) Set(val *ChannelRateMappingTypeRateTierAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelRateMappingTypeRateTierAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelRateMappingTypeRateTierAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelRateMappingTypeRateTierAttr(val *ChannelRateMappingTypeRateTierAttr) *NullableChannelRateMappingTypeRateTierAttr {
	return &NullableChannelRateMappingTypeRateTierAttr{value: val, isSet: true}
}

func (v NullableChannelRateMappingTypeRateTierAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelRateMappingTypeRateTierAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


