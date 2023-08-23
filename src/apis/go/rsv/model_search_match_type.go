/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SearchMatchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchMatchType{}

// SearchMatchType Search match indicating attribute and the matching value.
type SearchMatchType struct {
	// Search match attribute.
	Attribute *string `json:"attribute,omitempty"`
	// Search match value.
	Value *string `json:"value,omitempty"`
}

// NewSearchMatchType instantiates a new SearchMatchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchMatchType() *SearchMatchType {
	this := SearchMatchType{}
	return &this
}

// NewSearchMatchTypeWithDefaults instantiates a new SearchMatchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchMatchTypeWithDefaults() *SearchMatchType {
	this := SearchMatchType{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *SearchMatchType) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMatchType) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *SearchMatchType) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *SearchMatchType) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SearchMatchType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMatchType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SearchMatchType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SearchMatchType) SetValue(v string) {
	o.Value = &v
}

func (o SearchMatchType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchMatchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSearchMatchType struct {
	value *SearchMatchType
	isSet bool
}

func (v NullableSearchMatchType) Get() *SearchMatchType {
	return v.value
}

func (v *NullableSearchMatchType) Set(val *SearchMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchMatchType(val *SearchMatchType) *NullableSearchMatchType {
	return &NullableSearchMatchType{value: val, isSet: true}
}

func (v NullableSearchMatchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


