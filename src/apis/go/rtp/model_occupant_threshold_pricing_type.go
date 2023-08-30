/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtp

import (
	"encoding/json"
)

// checks if the OccupantThresholdPricingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OccupantThresholdPricingType{}

// OccupantThresholdPricingType Definition for creating pricing for a rate schedule, based on occupants threshold level.
type OccupantThresholdPricingType struct {
	Adults *OccupantThresholdPricingTypeAdults `json:"adults,omitempty"`
	Children *OccupantThresholdPricingTypeChildren `json:"children,omitempty"`
	Occupants *OccupantThresholdPricingTypeOccupants `json:"occupants,omitempty"`
}

// NewOccupantThresholdPricingType instantiates a new OccupantThresholdPricingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOccupantThresholdPricingType() *OccupantThresholdPricingType {
	this := OccupantThresholdPricingType{}
	return &this
}

// NewOccupantThresholdPricingTypeWithDefaults instantiates a new OccupantThresholdPricingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOccupantThresholdPricingTypeWithDefaults() *OccupantThresholdPricingType {
	this := OccupantThresholdPricingType{}
	return &this
}

// GetAdults returns the Adults field value if set, zero value otherwise.
func (o *OccupantThresholdPricingType) GetAdults() OccupantThresholdPricingTypeAdults {
	if o == nil || IsNil(o.Adults) {
		var ret OccupantThresholdPricingTypeAdults
		return ret
	}
	return *o.Adults
}

// GetAdultsOk returns a tuple with the Adults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OccupantThresholdPricingType) GetAdultsOk() (*OccupantThresholdPricingTypeAdults, bool) {
	if o == nil || IsNil(o.Adults) {
		return nil, false
	}
	return o.Adults, true
}

// HasAdults returns a boolean if a field has been set.
func (o *OccupantThresholdPricingType) HasAdults() bool {
	if o != nil && !IsNil(o.Adults) {
		return true
	}

	return false
}

// SetAdults gets a reference to the given OccupantThresholdPricingTypeAdults and assigns it to the Adults field.
func (o *OccupantThresholdPricingType) SetAdults(v OccupantThresholdPricingTypeAdults) {
	o.Adults = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *OccupantThresholdPricingType) GetChildren() OccupantThresholdPricingTypeChildren {
	if o == nil || IsNil(o.Children) {
		var ret OccupantThresholdPricingTypeChildren
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OccupantThresholdPricingType) GetChildrenOk() (*OccupantThresholdPricingTypeChildren, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *OccupantThresholdPricingType) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given OccupantThresholdPricingTypeChildren and assigns it to the Children field.
func (o *OccupantThresholdPricingType) SetChildren(v OccupantThresholdPricingTypeChildren) {
	o.Children = &v
}

// GetOccupants returns the Occupants field value if set, zero value otherwise.
func (o *OccupantThresholdPricingType) GetOccupants() OccupantThresholdPricingTypeOccupants {
	if o == nil || IsNil(o.Occupants) {
		var ret OccupantThresholdPricingTypeOccupants
		return ret
	}
	return *o.Occupants
}

// GetOccupantsOk returns a tuple with the Occupants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OccupantThresholdPricingType) GetOccupantsOk() (*OccupantThresholdPricingTypeOccupants, bool) {
	if o == nil || IsNil(o.Occupants) {
		return nil, false
	}
	return o.Occupants, true
}

// HasOccupants returns a boolean if a field has been set.
func (o *OccupantThresholdPricingType) HasOccupants() bool {
	if o != nil && !IsNil(o.Occupants) {
		return true
	}

	return false
}

// SetOccupants gets a reference to the given OccupantThresholdPricingTypeOccupants and assigns it to the Occupants field.
func (o *OccupantThresholdPricingType) SetOccupants(v OccupantThresholdPricingTypeOccupants) {
	o.Occupants = &v
}

func (o OccupantThresholdPricingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OccupantThresholdPricingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adults) {
		toSerialize["adults"] = o.Adults
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.Occupants) {
		toSerialize["occupants"] = o.Occupants
	}
	return toSerialize, nil
}

type NullableOccupantThresholdPricingType struct {
	value *OccupantThresholdPricingType
	isSet bool
}

func (v NullableOccupantThresholdPricingType) Get() *OccupantThresholdPricingType {
	return v.value
}

func (v *NullableOccupantThresholdPricingType) Set(val *OccupantThresholdPricingType) {
	v.value = val
	v.isSet = true
}

func (v NullableOccupantThresholdPricingType) IsSet() bool {
	return v.isSet
}

func (v *NullableOccupantThresholdPricingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOccupantThresholdPricingType(val *OccupantThresholdPricingType) *NullableOccupantThresholdPricingType {
	return &NullableOccupantThresholdPricingType{value: val, isSet: true}
}

func (v NullableOccupantThresholdPricingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOccupantThresholdPricingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


