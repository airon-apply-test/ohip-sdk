/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
)

// checks if the ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings{}

// ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings Paging information and collection of channel-hotel guarantee code mappings.
type ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings struct {
	// Channel-hotel guarantee code mapping and its details.
	ChannelGuaranteeCodeMapping []ChannelGuaranteeCodeMappingType `json:"channelGuaranteeCodeMapping,omitempty"`
}

// NewChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings instantiates a new ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings() *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings {
	this := ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings{}
	return &this
}

// NewChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappingsWithDefaults instantiates a new ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappingsWithDefaults() *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings {
	this := ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings{}
	return &this
}

// GetChannelGuaranteeCodeMapping returns the ChannelGuaranteeCodeMapping field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) GetChannelGuaranteeCodeMapping() []ChannelGuaranteeCodeMappingType {
	if o == nil || IsNil(o.ChannelGuaranteeCodeMapping) {
		var ret []ChannelGuaranteeCodeMappingType
		return ret
	}
	return o.ChannelGuaranteeCodeMapping
}

// GetChannelGuaranteeCodeMappingOk returns a tuple with the ChannelGuaranteeCodeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) GetChannelGuaranteeCodeMappingOk() ([]ChannelGuaranteeCodeMappingType, bool) {
	if o == nil || IsNil(o.ChannelGuaranteeCodeMapping) {
		return nil, false
	}
	return o.ChannelGuaranteeCodeMapping, true
}

// HasChannelGuaranteeCodeMapping returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) HasChannelGuaranteeCodeMapping() bool {
	if o != nil && !IsNil(o.ChannelGuaranteeCodeMapping) {
		return true
	}

	return false
}

// SetChannelGuaranteeCodeMapping gets a reference to the given []ChannelGuaranteeCodeMappingType and assigns it to the ChannelGuaranteeCodeMapping field.
func (o *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) SetChannelGuaranteeCodeMapping(v []ChannelGuaranteeCodeMappingType) {
	o.ChannelGuaranteeCodeMapping = v
}

func (o ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelGuaranteeCodeMapping) {
		toSerialize["channelGuaranteeCodeMapping"] = o.ChannelGuaranteeCodeMapping
	}
	return toSerialize, nil
}

type NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings struct {
	value *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings
	isSet bool
}

func (v NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) Get() *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings {
	return v.value
}

func (v *NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) Set(val *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings(val *ChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) *NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings {
	return &NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings{value: val, isSet: true}
}

func (v NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelGuaranteeCodeMappingDetailsChannelGuaranteeCodeMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


