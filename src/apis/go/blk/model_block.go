/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Block type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Block{}

// Block Request object for creation of blocks. This object contains block details with unique identifiers for each block. The standard optional Opera Context element is also included.
type Block struct {
	Blocks *BlocksType `json:"blocks,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBlock instantiates a new Block object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlock() *Block {
	this := Block{}
	return &this
}

// NewBlockWithDefaults instantiates a new Block object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockWithDefaults() *Block {
	this := Block{}
	return &this
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *Block) GetBlocks() BlocksType {
	if o == nil || IsNil(o.Blocks) {
		var ret BlocksType
		return ret
	}
	return *o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Block) GetBlocksOk() (*BlocksType, bool) {
	if o == nil || IsNil(o.Blocks) {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *Block) HasBlocks() bool {
	if o != nil && !IsNil(o.Blocks) {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given BlocksType and assigns it to the Blocks field.
func (o *Block) SetBlocks(v BlocksType) {
	o.Blocks = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Block) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Block) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Block) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *Block) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Block) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Block) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Block) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *Block) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o Block) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Block) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blocks) {
		toSerialize["blocks"] = o.Blocks
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBlock struct {
	value *Block
	isSet bool
}

func (v NullableBlock) Get() *Block {
	return v.value
}

func (v *NullableBlock) Set(val *Block) {
	v.value = val
	v.isSet = true
}

func (v NullableBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlock(val *Block) *NullableBlock {
	return &NullableBlock{value: val, isSet: true}
}

func (v NullableBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


