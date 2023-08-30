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

// checks if the ChannelGuaranteeCodeMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelGuaranteeCodeMappings{}

// ChannelGuaranteeCodeMappings Request object for creating a new channel-hotel guarantee code mapping.
type ChannelGuaranteeCodeMappings struct {
	ChannelGuaranteeCodeMapping *ChannelGuaranteeCodeMappingType `json:"channelGuaranteeCodeMapping,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChannelGuaranteeCodeMappings instantiates a new ChannelGuaranteeCodeMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelGuaranteeCodeMappings() *ChannelGuaranteeCodeMappings {
	this := ChannelGuaranteeCodeMappings{}
	return &this
}

// NewChannelGuaranteeCodeMappingsWithDefaults instantiates a new ChannelGuaranteeCodeMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelGuaranteeCodeMappingsWithDefaults() *ChannelGuaranteeCodeMappings {
	this := ChannelGuaranteeCodeMappings{}
	return &this
}

// GetChannelGuaranteeCodeMapping returns the ChannelGuaranteeCodeMapping field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappings) GetChannelGuaranteeCodeMapping() ChannelGuaranteeCodeMappingType {
	if o == nil || IsNil(o.ChannelGuaranteeCodeMapping) {
		var ret ChannelGuaranteeCodeMappingType
		return ret
	}
	return *o.ChannelGuaranteeCodeMapping
}

// GetChannelGuaranteeCodeMappingOk returns a tuple with the ChannelGuaranteeCodeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappings) GetChannelGuaranteeCodeMappingOk() (*ChannelGuaranteeCodeMappingType, bool) {
	if o == nil || IsNil(o.ChannelGuaranteeCodeMapping) {
		return nil, false
	}
	return o.ChannelGuaranteeCodeMapping, true
}

// HasChannelGuaranteeCodeMapping returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappings) HasChannelGuaranteeCodeMapping() bool {
	if o != nil && !IsNil(o.ChannelGuaranteeCodeMapping) {
		return true
	}

	return false
}

// SetChannelGuaranteeCodeMapping gets a reference to the given ChannelGuaranteeCodeMappingType and assigns it to the ChannelGuaranteeCodeMapping field.
func (o *ChannelGuaranteeCodeMappings) SetChannelGuaranteeCodeMapping(v ChannelGuaranteeCodeMappingType) {
	o.ChannelGuaranteeCodeMapping = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappings) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappings) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappings) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChannelGuaranteeCodeMappings) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappings) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappings) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappings) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChannelGuaranteeCodeMappings) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChannelGuaranteeCodeMappings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelGuaranteeCodeMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelGuaranteeCodeMapping) {
		toSerialize["channelGuaranteeCodeMapping"] = o.ChannelGuaranteeCodeMapping
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChannelGuaranteeCodeMappings struct {
	value *ChannelGuaranteeCodeMappings
	isSet bool
}

func (v NullableChannelGuaranteeCodeMappings) Get() *ChannelGuaranteeCodeMappings {
	return v.value
}

func (v *NullableChannelGuaranteeCodeMappings) Set(val *ChannelGuaranteeCodeMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelGuaranteeCodeMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelGuaranteeCodeMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelGuaranteeCodeMappings(val *ChannelGuaranteeCodeMappings) *NullableChannelGuaranteeCodeMappings {
	return &NullableChannelGuaranteeCodeMappings{value: val, isSet: true}
}

func (v NullableChannelGuaranteeCodeMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelGuaranteeCodeMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


