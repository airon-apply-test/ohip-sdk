/*
OPERA Cloud Price Availability Rate Async API

APIs to cater for Price and Rate Availability Asynchronous functionality in OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BlockId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockId{}

// BlockId An identifier used to uniquely reference an object in a system (e.g. an airline reservation reference, customer profile reference, booking confirmation number, or a reference to a previous availability quote).
type BlockId struct {
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
}

// NewBlockId instantiates a new BlockId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockId() *BlockId {
	this := BlockId{}
	return &this
}

// NewBlockIdWithDefaults instantiates a new BlockId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockIdWithDefaults() *BlockId {
	this := BlockId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlockId) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockId) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlockId) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BlockId) SetId(v string) {
	o.Id = &v
}

func (o BlockId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableBlockId struct {
	value *BlockId
	isSet bool
}

func (v NullableBlockId) Get() *BlockId {
	return v.value
}

func (v *NullableBlockId) Set(val *BlockId) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockId) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockId(val *BlockId) *NullableBlockId {
	return &NullableBlockId{value: val, isSet: true}
}

func (v NullableBlockId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


