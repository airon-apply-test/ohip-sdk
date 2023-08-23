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

// checks if the RegionalRateParity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionalRateParity{}

// RegionalRateParity Response object to fetch regional rate parity.
type RegionalRateParity struct {
	RegionalRateParity *RegionalRateParityType `json:"regionalRateParity,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewRegionalRateParity instantiates a new RegionalRateParity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalRateParity() *RegionalRateParity {
	this := RegionalRateParity{}
	return &this
}

// NewRegionalRateParityWithDefaults instantiates a new RegionalRateParity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalRateParityWithDefaults() *RegionalRateParity {
	this := RegionalRateParity{}
	return &this
}

// GetRegionalRateParity returns the RegionalRateParity field value if set, zero value otherwise.
func (o *RegionalRateParity) GetRegionalRateParity() RegionalRateParityType {
	if o == nil || IsNil(o.RegionalRateParity) {
		var ret RegionalRateParityType
		return ret
	}
	return *o.RegionalRateParity
}

// GetRegionalRateParityOk returns a tuple with the RegionalRateParity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalRateParity) GetRegionalRateParityOk() (*RegionalRateParityType, bool) {
	if o == nil || IsNil(o.RegionalRateParity) {
		return nil, false
	}
	return o.RegionalRateParity, true
}

// HasRegionalRateParity returns a boolean if a field has been set.
func (o *RegionalRateParity) HasRegionalRateParity() bool {
	if o != nil && !IsNil(o.RegionalRateParity) {
		return true
	}

	return false
}

// SetRegionalRateParity gets a reference to the given RegionalRateParityType and assigns it to the RegionalRateParity field.
func (o *RegionalRateParity) SetRegionalRateParity(v RegionalRateParityType) {
	o.RegionalRateParity = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RegionalRateParity) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalRateParity) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RegionalRateParity) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *RegionalRateParity) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *RegionalRateParity) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalRateParity) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RegionalRateParity) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *RegionalRateParity) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o RegionalRateParity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionalRateParity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegionalRateParity) {
		toSerialize["regionalRateParity"] = o.RegionalRateParity
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableRegionalRateParity struct {
	value *RegionalRateParity
	isSet bool
}

func (v NullableRegionalRateParity) Get() *RegionalRateParity {
	return v.value
}

func (v *NullableRegionalRateParity) Set(val *RegionalRateParity) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalRateParity) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalRateParity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalRateParity(val *RegionalRateParity) *NullableRegionalRateParity {
	return &NullableRegionalRateParity{value: val, isSet: true}
}

func (v NullableRegionalRateParity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalRateParity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


