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

// checks if the ChannelRates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelRates{}

// ChannelRates Request object for distributing channel rates.
type ChannelRates struct {
	DistributeChannelRates *DistributeChannelRatesType `json:"distributeChannelRates,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChannelRates instantiates a new ChannelRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelRates() *ChannelRates {
	this := ChannelRates{}
	return &this
}

// NewChannelRatesWithDefaults instantiates a new ChannelRates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelRatesWithDefaults() *ChannelRates {
	this := ChannelRates{}
	return &this
}

// GetDistributeChannelRates returns the DistributeChannelRates field value if set, zero value otherwise.
func (o *ChannelRates) GetDistributeChannelRates() DistributeChannelRatesType {
	if o == nil || IsNil(o.DistributeChannelRates) {
		var ret DistributeChannelRatesType
		return ret
	}
	return *o.DistributeChannelRates
}

// GetDistributeChannelRatesOk returns a tuple with the DistributeChannelRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRates) GetDistributeChannelRatesOk() (*DistributeChannelRatesType, bool) {
	if o == nil || IsNil(o.DistributeChannelRates) {
		return nil, false
	}
	return o.DistributeChannelRates, true
}

// HasDistributeChannelRates returns a boolean if a field has been set.
func (o *ChannelRates) HasDistributeChannelRates() bool {
	if o != nil && !IsNil(o.DistributeChannelRates) {
		return true
	}

	return false
}

// SetDistributeChannelRates gets a reference to the given DistributeChannelRatesType and assigns it to the DistributeChannelRates field.
func (o *ChannelRates) SetDistributeChannelRates(v DistributeChannelRatesType) {
	o.DistributeChannelRates = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChannelRates) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRates) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChannelRates) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChannelRates) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChannelRates) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRates) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChannelRates) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChannelRates) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChannelRates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelRates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DistributeChannelRates) {
		toSerialize["distributeChannelRates"] = o.DistributeChannelRates
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChannelRates struct {
	value *ChannelRates
	isSet bool
}

func (v NullableChannelRates) Get() *ChannelRates {
	return v.value
}

func (v *NullableChannelRates) Set(val *ChannelRates) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelRates) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelRates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelRates(val *ChannelRates) *NullableChannelRates {
	return &NullableChannelRates{value: val, isSet: true}
}

func (v NullableChannelRates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelRates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


