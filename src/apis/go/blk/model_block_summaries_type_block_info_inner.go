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

// checks if the BlockSummariesTypeBlockInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockSummariesTypeBlockInfoInner{}

// BlockSummariesTypeBlockInfoInner struct for BlockSummariesTypeBlockInfoInner
type BlockSummariesTypeBlockInfoInner struct {
	// Unique Id that references an object uniquely in the system.
	BlockIdList []UniqueIDType `json:"blockIdList,omitempty"`
	Block *BlockSummaryType `json:"block,omitempty"`
}

// NewBlockSummariesTypeBlockInfoInner instantiates a new BlockSummariesTypeBlockInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSummariesTypeBlockInfoInner() *BlockSummariesTypeBlockInfoInner {
	this := BlockSummariesTypeBlockInfoInner{}
	return &this
}

// NewBlockSummariesTypeBlockInfoInnerWithDefaults instantiates a new BlockSummariesTypeBlockInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSummariesTypeBlockInfoInnerWithDefaults() *BlockSummariesTypeBlockInfoInner {
	this := BlockSummariesTypeBlockInfoInner{}
	return &this
}

// GetBlockIdList returns the BlockIdList field value if set, zero value otherwise.
func (o *BlockSummariesTypeBlockInfoInner) GetBlockIdList() []UniqueIDType {
	if o == nil || IsNil(o.BlockIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.BlockIdList
}

// GetBlockIdListOk returns a tuple with the BlockIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSummariesTypeBlockInfoInner) GetBlockIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.BlockIdList) {
		return nil, false
	}
	return o.BlockIdList, true
}

// HasBlockIdList returns a boolean if a field has been set.
func (o *BlockSummariesTypeBlockInfoInner) HasBlockIdList() bool {
	if o != nil && !IsNil(o.BlockIdList) {
		return true
	}

	return false
}

// SetBlockIdList gets a reference to the given []UniqueIDType and assigns it to the BlockIdList field.
func (o *BlockSummariesTypeBlockInfoInner) SetBlockIdList(v []UniqueIDType) {
	o.BlockIdList = v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *BlockSummariesTypeBlockInfoInner) GetBlock() BlockSummaryType {
	if o == nil || IsNil(o.Block) {
		var ret BlockSummaryType
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSummariesTypeBlockInfoInner) GetBlockOk() (*BlockSummaryType, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *BlockSummariesTypeBlockInfoInner) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given BlockSummaryType and assigns it to the Block field.
func (o *BlockSummariesTypeBlockInfoInner) SetBlock(v BlockSummaryType) {
	o.Block = &v
}

func (o BlockSummariesTypeBlockInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockSummariesTypeBlockInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockIdList) {
		toSerialize["blockIdList"] = o.BlockIdList
	}
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	return toSerialize, nil
}

type NullableBlockSummariesTypeBlockInfoInner struct {
	value *BlockSummariesTypeBlockInfoInner
	isSet bool
}

func (v NullableBlockSummariesTypeBlockInfoInner) Get() *BlockSummariesTypeBlockInfoInner {
	return v.value
}

func (v *NullableBlockSummariesTypeBlockInfoInner) Set(val *BlockSummariesTypeBlockInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSummariesTypeBlockInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSummariesTypeBlockInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSummariesTypeBlockInfoInner(val *BlockSummariesTypeBlockInfoInner) *NullableBlockSummariesTypeBlockInfoInner {
	return &NullableBlockSummariesTypeBlockInfoInner{value: val, isSet: true}
}

func (v NullableBlockSummariesTypeBlockInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSummariesTypeBlockInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


