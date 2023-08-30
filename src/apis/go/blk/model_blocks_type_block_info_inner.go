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

// checks if the BlocksTypeBlockInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlocksTypeBlockInfoInner{}

// BlocksTypeBlockInfoInner struct for BlocksTypeBlockInfoInner
type BlocksTypeBlockInfoInner struct {
	Block *BlockType `json:"block,omitempty"`
}

// NewBlocksTypeBlockInfoInner instantiates a new BlocksTypeBlockInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlocksTypeBlockInfoInner() *BlocksTypeBlockInfoInner {
	this := BlocksTypeBlockInfoInner{}
	return &this
}

// NewBlocksTypeBlockInfoInnerWithDefaults instantiates a new BlocksTypeBlockInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlocksTypeBlockInfoInnerWithDefaults() *BlocksTypeBlockInfoInner {
	this := BlocksTypeBlockInfoInner{}
	return &this
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *BlocksTypeBlockInfoInner) GetBlock() BlockType {
	if o == nil || IsNil(o.Block) {
		var ret BlockType
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlocksTypeBlockInfoInner) GetBlockOk() (*BlockType, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *BlocksTypeBlockInfoInner) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given BlockType and assigns it to the Block field.
func (o *BlocksTypeBlockInfoInner) SetBlock(v BlockType) {
	o.Block = &v
}

func (o BlocksTypeBlockInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlocksTypeBlockInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	return toSerialize, nil
}

type NullableBlocksTypeBlockInfoInner struct {
	value *BlocksTypeBlockInfoInner
	isSet bool
}

func (v NullableBlocksTypeBlockInfoInner) Get() *BlocksTypeBlockInfoInner {
	return v.value
}

func (v *NullableBlocksTypeBlockInfoInner) Set(val *BlocksTypeBlockInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBlocksTypeBlockInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBlocksTypeBlockInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlocksTypeBlockInfoInner(val *BlocksTypeBlockInfoInner) *NullableBlocksTypeBlockInfoInner {
	return &NullableBlocksTypeBlockInfoInner{value: val, isSet: true}
}

func (v NullableBlocksTypeBlockInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlocksTypeBlockInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


