/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the CharacterUDFType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharacterUDFType{}

// CharacterUDFType Used to hold user defined field of Character/String Type.
type CharacterUDFType struct {
	// Used to hold user defined field of Character Type. It is highly recommended to use UDFC01, UDFC02,...UDFC40 (Total 40) as Character/String UDF names(commonly used on Reservation, Profile etc.). Name is not restricted using enumeration, to provide flexibility of different name usage if required.
	Name *string `json:"name,omitempty"`
	// Value of user defined field.
	Value *string `json:"value,omitempty"`
	// Label of user defined field used by vendors or customers.
	AlternateName *string `json:"alternateName,omitempty"`
}

// NewCharacterUDFType instantiates a new CharacterUDFType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacterUDFType() *CharacterUDFType {
	this := CharacterUDFType{}
	return &this
}

// NewCharacterUDFTypeWithDefaults instantiates a new CharacterUDFType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacterUDFTypeWithDefaults() *CharacterUDFType {
	this := CharacterUDFType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CharacterUDFType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterUDFType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CharacterUDFType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CharacterUDFType) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CharacterUDFType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterUDFType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CharacterUDFType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CharacterUDFType) SetValue(v string) {
	o.Value = &v
}

// GetAlternateName returns the AlternateName field value if set, zero value otherwise.
func (o *CharacterUDFType) GetAlternateName() string {
	if o == nil || IsNil(o.AlternateName) {
		var ret string
		return ret
	}
	return *o.AlternateName
}

// GetAlternateNameOk returns a tuple with the AlternateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterUDFType) GetAlternateNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateName) {
		return nil, false
	}
	return o.AlternateName, true
}

// HasAlternateName returns a boolean if a field has been set.
func (o *CharacterUDFType) HasAlternateName() bool {
	if o != nil && !IsNil(o.AlternateName) {
		return true
	}

	return false
}

// SetAlternateName gets a reference to the given string and assigns it to the AlternateName field.
func (o *CharacterUDFType) SetAlternateName(v string) {
	o.AlternateName = &v
}

func (o CharacterUDFType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharacterUDFType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.AlternateName) {
		toSerialize["alternateName"] = o.AlternateName
	}
	return toSerialize, nil
}

type NullableCharacterUDFType struct {
	value *CharacterUDFType
	isSet bool
}

func (v NullableCharacterUDFType) Get() *CharacterUDFType {
	return v.value
}

func (v *NullableCharacterUDFType) Set(val *CharacterUDFType) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacterUDFType) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacterUDFType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacterUDFType(val *CharacterUDFType) *NullableCharacterUDFType {
	return &NullableCharacterUDFType{value: val, isSet: true}
}

func (v NullableCharacterUDFType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacterUDFType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


