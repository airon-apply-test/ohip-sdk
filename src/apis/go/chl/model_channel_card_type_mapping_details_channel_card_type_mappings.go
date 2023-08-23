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

// checks if the ChannelCardTypeMappingDetailsChannelCardTypeMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelCardTypeMappingDetailsChannelCardTypeMappings{}

// ChannelCardTypeMappingDetailsChannelCardTypeMappings Paging information and collection of channel-hotel card type mappings.
type ChannelCardTypeMappingDetailsChannelCardTypeMappings struct {
	// Channel-hotel card type mapping and its details.
	ChannelCardTypeMapping []ChannelCardTypeMappingType `json:"channelCardTypeMapping,omitempty"`
}

// NewChannelCardTypeMappingDetailsChannelCardTypeMappings instantiates a new ChannelCardTypeMappingDetailsChannelCardTypeMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelCardTypeMappingDetailsChannelCardTypeMappings() *ChannelCardTypeMappingDetailsChannelCardTypeMappings {
	this := ChannelCardTypeMappingDetailsChannelCardTypeMappings{}
	return &this
}

// NewChannelCardTypeMappingDetailsChannelCardTypeMappingsWithDefaults instantiates a new ChannelCardTypeMappingDetailsChannelCardTypeMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelCardTypeMappingDetailsChannelCardTypeMappingsWithDefaults() *ChannelCardTypeMappingDetailsChannelCardTypeMappings {
	this := ChannelCardTypeMappingDetailsChannelCardTypeMappings{}
	return &this
}

// GetChannelCardTypeMapping returns the ChannelCardTypeMapping field value if set, zero value otherwise.
func (o *ChannelCardTypeMappingDetailsChannelCardTypeMappings) GetChannelCardTypeMapping() []ChannelCardTypeMappingType {
	if o == nil || IsNil(o.ChannelCardTypeMapping) {
		var ret []ChannelCardTypeMappingType
		return ret
	}
	return o.ChannelCardTypeMapping
}

// GetChannelCardTypeMappingOk returns a tuple with the ChannelCardTypeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelCardTypeMappingDetailsChannelCardTypeMappings) GetChannelCardTypeMappingOk() ([]ChannelCardTypeMappingType, bool) {
	if o == nil || IsNil(o.ChannelCardTypeMapping) {
		return nil, false
	}
	return o.ChannelCardTypeMapping, true
}

// HasChannelCardTypeMapping returns a boolean if a field has been set.
func (o *ChannelCardTypeMappingDetailsChannelCardTypeMappings) HasChannelCardTypeMapping() bool {
	if o != nil && !IsNil(o.ChannelCardTypeMapping) {
		return true
	}

	return false
}

// SetChannelCardTypeMapping gets a reference to the given []ChannelCardTypeMappingType and assigns it to the ChannelCardTypeMapping field.
func (o *ChannelCardTypeMappingDetailsChannelCardTypeMappings) SetChannelCardTypeMapping(v []ChannelCardTypeMappingType) {
	o.ChannelCardTypeMapping = v
}

func (o ChannelCardTypeMappingDetailsChannelCardTypeMappings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelCardTypeMappingDetailsChannelCardTypeMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelCardTypeMapping) {
		toSerialize["channelCardTypeMapping"] = o.ChannelCardTypeMapping
	}
	return toSerialize, nil
}

type NullableChannelCardTypeMappingDetailsChannelCardTypeMappings struct {
	value *ChannelCardTypeMappingDetailsChannelCardTypeMappings
	isSet bool
}

func (v NullableChannelCardTypeMappingDetailsChannelCardTypeMappings) Get() *ChannelCardTypeMappingDetailsChannelCardTypeMappings {
	return v.value
}

func (v *NullableChannelCardTypeMappingDetailsChannelCardTypeMappings) Set(val *ChannelCardTypeMappingDetailsChannelCardTypeMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelCardTypeMappingDetailsChannelCardTypeMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelCardTypeMappingDetailsChannelCardTypeMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelCardTypeMappingDetailsChannelCardTypeMappings(val *ChannelCardTypeMappingDetailsChannelCardTypeMappings) *NullableChannelCardTypeMappingDetailsChannelCardTypeMappings {
	return &NullableChannelCardTypeMappingDetailsChannelCardTypeMappings{value: val, isSet: true}
}

func (v NullableChannelCardTypeMappingDetailsChannelCardTypeMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelCardTypeMappingDetailsChannelCardTypeMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


