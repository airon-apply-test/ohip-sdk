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

// checks if the RateParity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateParity{}

// RateParity Response object to fetch Rate Parity.
type RateParity struct {
	RateParity *RateParityType `json:"rateParity,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewRateParity instantiates a new RateParity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateParity() *RateParity {
	this := RateParity{}
	return &this
}

// NewRateParityWithDefaults instantiates a new RateParity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateParityWithDefaults() *RateParity {
	this := RateParity{}
	return &this
}

// GetRateParity returns the RateParity field value if set, zero value otherwise.
func (o *RateParity) GetRateParity() RateParityType {
	if o == nil || IsNil(o.RateParity) {
		var ret RateParityType
		return ret
	}
	return *o.RateParity
}

// GetRateParityOk returns a tuple with the RateParity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateParity) GetRateParityOk() (*RateParityType, bool) {
	if o == nil || IsNil(o.RateParity) {
		return nil, false
	}
	return o.RateParity, true
}

// HasRateParity returns a boolean if a field has been set.
func (o *RateParity) HasRateParity() bool {
	if o != nil && !IsNil(o.RateParity) {
		return true
	}

	return false
}

// SetRateParity gets a reference to the given RateParityType and assigns it to the RateParity field.
func (o *RateParity) SetRateParity(v RateParityType) {
	o.RateParity = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RateParity) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateParity) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RateParity) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *RateParity) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *RateParity) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateParity) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RateParity) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *RateParity) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o RateParity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateParity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RateParity) {
		toSerialize["rateParity"] = o.RateParity
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableRateParity struct {
	value *RateParity
	isSet bool
}

func (v NullableRateParity) Get() *RateParity {
	return v.value
}

func (v *NullableRateParity) Set(val *RateParity) {
	v.value = val
	v.isSet = true
}

func (v NullableRateParity) IsSet() bool {
	return v.isSet
}

func (v *NullableRateParity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateParity(val *RateParity) *NullableRateParity {
	return &NullableRateParity{value: val, isSet: true}
}

func (v NullableRateParity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateParity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


