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

// checks if the ChannelRateMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelRateMapping{}

// ChannelRateMapping Request object for changing channel rate mappings.
type ChannelRateMapping struct {
	ChannelRateMappings *ChannelRateMappingChannelRateMappings `json:"channelRateMappings,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChannelRateMapping instantiates a new ChannelRateMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelRateMapping() *ChannelRateMapping {
	this := ChannelRateMapping{}
	return &this
}

// NewChannelRateMappingWithDefaults instantiates a new ChannelRateMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelRateMappingWithDefaults() *ChannelRateMapping {
	this := ChannelRateMapping{}
	return &this
}

// GetChannelRateMappings returns the ChannelRateMappings field value if set, zero value otherwise.
func (o *ChannelRateMapping) GetChannelRateMappings() ChannelRateMappingChannelRateMappings {
	if o == nil || IsNil(o.ChannelRateMappings) {
		var ret ChannelRateMappingChannelRateMappings
		return ret
	}
	return *o.ChannelRateMappings
}

// GetChannelRateMappingsOk returns a tuple with the ChannelRateMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRateMapping) GetChannelRateMappingsOk() (*ChannelRateMappingChannelRateMappings, bool) {
	if o == nil || IsNil(o.ChannelRateMappings) {
		return nil, false
	}
	return o.ChannelRateMappings, true
}

// HasChannelRateMappings returns a boolean if a field has been set.
func (o *ChannelRateMapping) HasChannelRateMappings() bool {
	if o != nil && !IsNil(o.ChannelRateMappings) {
		return true
	}

	return false
}

// SetChannelRateMappings gets a reference to the given ChannelRateMappingChannelRateMappings and assigns it to the ChannelRateMappings field.
func (o *ChannelRateMapping) SetChannelRateMappings(v ChannelRateMappingChannelRateMappings) {
	o.ChannelRateMappings = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChannelRateMapping) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRateMapping) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChannelRateMapping) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChannelRateMapping) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChannelRateMapping) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRateMapping) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChannelRateMapping) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChannelRateMapping) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChannelRateMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelRateMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelRateMappings) {
		toSerialize["channelRateMappings"] = o.ChannelRateMappings
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChannelRateMapping struct {
	value *ChannelRateMapping
	isSet bool
}

func (v NullableChannelRateMapping) Get() *ChannelRateMapping {
	return v.value
}

func (v *NullableChannelRateMapping) Set(val *ChannelRateMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelRateMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelRateMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelRateMapping(val *ChannelRateMapping) *NullableChannelRateMapping {
	return &NullableChannelRateMapping{value: val, isSet: true}
}

func (v NullableChannelRateMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelRateMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


