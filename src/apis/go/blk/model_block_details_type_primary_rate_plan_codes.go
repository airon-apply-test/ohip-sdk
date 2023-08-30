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

// checks if the BlockDetailsTypePrimaryRatePlanCodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockDetailsTypePrimaryRatePlanCodes{}

// BlockDetailsTypePrimaryRatePlanCodes Primary rates for the block.
type BlockDetailsTypePrimaryRatePlanCodes struct {
	BlockRatePlanCode *BlockRatePlanInfoType `json:"blockRatePlanCode,omitempty"`
	ShoulderStartRatePlanCode *BlockRatePlanInfoType `json:"shoulderStartRatePlanCode,omitempty"`
	ShoulderEndRatePlanCode *BlockRatePlanInfoType `json:"shoulderEndRatePlanCode,omitempty"`
	// Total number of Block Rates.
	BlockRatePlanCodeCount *float32 `json:"blockRatePlanCodeCount,omitempty"`
	// Total number of Shoulder Start Rates.
	ShoulderStartRatePlanCodeCount *float32 `json:"shoulderStartRatePlanCodeCount,omitempty"`
	// Total number of Shoulder End Rates.
	ShoulderEndRatePlanCodeCount *float32 `json:"shoulderEndRatePlanCodeCount,omitempty"`
}

// NewBlockDetailsTypePrimaryRatePlanCodes instantiates a new BlockDetailsTypePrimaryRatePlanCodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockDetailsTypePrimaryRatePlanCodes() *BlockDetailsTypePrimaryRatePlanCodes {
	this := BlockDetailsTypePrimaryRatePlanCodes{}
	return &this
}

// NewBlockDetailsTypePrimaryRatePlanCodesWithDefaults instantiates a new BlockDetailsTypePrimaryRatePlanCodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockDetailsTypePrimaryRatePlanCodesWithDefaults() *BlockDetailsTypePrimaryRatePlanCodes {
	this := BlockDetailsTypePrimaryRatePlanCodes{}
	return &this
}

// GetBlockRatePlanCode returns the BlockRatePlanCode field value if set, zero value otherwise.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetBlockRatePlanCode() BlockRatePlanInfoType {
	if o == nil || IsNil(o.BlockRatePlanCode) {
		var ret BlockRatePlanInfoType
		return ret
	}
	return *o.BlockRatePlanCode
}

// GetBlockRatePlanCodeOk returns a tuple with the BlockRatePlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetBlockRatePlanCodeOk() (*BlockRatePlanInfoType, bool) {
	if o == nil || IsNil(o.BlockRatePlanCode) {
		return nil, false
	}
	return o.BlockRatePlanCode, true
}

// HasBlockRatePlanCode returns a boolean if a field has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) HasBlockRatePlanCode() bool {
	if o != nil && !IsNil(o.BlockRatePlanCode) {
		return true
	}

	return false
}

// SetBlockRatePlanCode gets a reference to the given BlockRatePlanInfoType and assigns it to the BlockRatePlanCode field.
func (o *BlockDetailsTypePrimaryRatePlanCodes) SetBlockRatePlanCode(v BlockRatePlanInfoType) {
	o.BlockRatePlanCode = &v
}

// GetShoulderStartRatePlanCode returns the ShoulderStartRatePlanCode field value if set, zero value otherwise.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetShoulderStartRatePlanCode() BlockRatePlanInfoType {
	if o == nil || IsNil(o.ShoulderStartRatePlanCode) {
		var ret BlockRatePlanInfoType
		return ret
	}
	return *o.ShoulderStartRatePlanCode
}

// GetShoulderStartRatePlanCodeOk returns a tuple with the ShoulderStartRatePlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetShoulderStartRatePlanCodeOk() (*BlockRatePlanInfoType, bool) {
	if o == nil || IsNil(o.ShoulderStartRatePlanCode) {
		return nil, false
	}
	return o.ShoulderStartRatePlanCode, true
}

// HasShoulderStartRatePlanCode returns a boolean if a field has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) HasShoulderStartRatePlanCode() bool {
	if o != nil && !IsNil(o.ShoulderStartRatePlanCode) {
		return true
	}

	return false
}

// SetShoulderStartRatePlanCode gets a reference to the given BlockRatePlanInfoType and assigns it to the ShoulderStartRatePlanCode field.
func (o *BlockDetailsTypePrimaryRatePlanCodes) SetShoulderStartRatePlanCode(v BlockRatePlanInfoType) {
	o.ShoulderStartRatePlanCode = &v
}

// GetShoulderEndRatePlanCode returns the ShoulderEndRatePlanCode field value if set, zero value otherwise.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetShoulderEndRatePlanCode() BlockRatePlanInfoType {
	if o == nil || IsNil(o.ShoulderEndRatePlanCode) {
		var ret BlockRatePlanInfoType
		return ret
	}
	return *o.ShoulderEndRatePlanCode
}

// GetShoulderEndRatePlanCodeOk returns a tuple with the ShoulderEndRatePlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetShoulderEndRatePlanCodeOk() (*BlockRatePlanInfoType, bool) {
	if o == nil || IsNil(o.ShoulderEndRatePlanCode) {
		return nil, false
	}
	return o.ShoulderEndRatePlanCode, true
}

// HasShoulderEndRatePlanCode returns a boolean if a field has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) HasShoulderEndRatePlanCode() bool {
	if o != nil && !IsNil(o.ShoulderEndRatePlanCode) {
		return true
	}

	return false
}

// SetShoulderEndRatePlanCode gets a reference to the given BlockRatePlanInfoType and assigns it to the ShoulderEndRatePlanCode field.
func (o *BlockDetailsTypePrimaryRatePlanCodes) SetShoulderEndRatePlanCode(v BlockRatePlanInfoType) {
	o.ShoulderEndRatePlanCode = &v
}

// GetBlockRatePlanCodeCount returns the BlockRatePlanCodeCount field value if set, zero value otherwise.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetBlockRatePlanCodeCount() float32 {
	if o == nil || IsNil(o.BlockRatePlanCodeCount) {
		var ret float32
		return ret
	}
	return *o.BlockRatePlanCodeCount
}

// GetBlockRatePlanCodeCountOk returns a tuple with the BlockRatePlanCodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetBlockRatePlanCodeCountOk() (*float32, bool) {
	if o == nil || IsNil(o.BlockRatePlanCodeCount) {
		return nil, false
	}
	return o.BlockRatePlanCodeCount, true
}

// HasBlockRatePlanCodeCount returns a boolean if a field has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) HasBlockRatePlanCodeCount() bool {
	if o != nil && !IsNil(o.BlockRatePlanCodeCount) {
		return true
	}

	return false
}

// SetBlockRatePlanCodeCount gets a reference to the given float32 and assigns it to the BlockRatePlanCodeCount field.
func (o *BlockDetailsTypePrimaryRatePlanCodes) SetBlockRatePlanCodeCount(v float32) {
	o.BlockRatePlanCodeCount = &v
}

// GetShoulderStartRatePlanCodeCount returns the ShoulderStartRatePlanCodeCount field value if set, zero value otherwise.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetShoulderStartRatePlanCodeCount() float32 {
	if o == nil || IsNil(o.ShoulderStartRatePlanCodeCount) {
		var ret float32
		return ret
	}
	return *o.ShoulderStartRatePlanCodeCount
}

// GetShoulderStartRatePlanCodeCountOk returns a tuple with the ShoulderStartRatePlanCodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetShoulderStartRatePlanCodeCountOk() (*float32, bool) {
	if o == nil || IsNil(o.ShoulderStartRatePlanCodeCount) {
		return nil, false
	}
	return o.ShoulderStartRatePlanCodeCount, true
}

// HasShoulderStartRatePlanCodeCount returns a boolean if a field has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) HasShoulderStartRatePlanCodeCount() bool {
	if o != nil && !IsNil(o.ShoulderStartRatePlanCodeCount) {
		return true
	}

	return false
}

// SetShoulderStartRatePlanCodeCount gets a reference to the given float32 and assigns it to the ShoulderStartRatePlanCodeCount field.
func (o *BlockDetailsTypePrimaryRatePlanCodes) SetShoulderStartRatePlanCodeCount(v float32) {
	o.ShoulderStartRatePlanCodeCount = &v
}

// GetShoulderEndRatePlanCodeCount returns the ShoulderEndRatePlanCodeCount field value if set, zero value otherwise.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetShoulderEndRatePlanCodeCount() float32 {
	if o == nil || IsNil(o.ShoulderEndRatePlanCodeCount) {
		var ret float32
		return ret
	}
	return *o.ShoulderEndRatePlanCodeCount
}

// GetShoulderEndRatePlanCodeCountOk returns a tuple with the ShoulderEndRatePlanCodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) GetShoulderEndRatePlanCodeCountOk() (*float32, bool) {
	if o == nil || IsNil(o.ShoulderEndRatePlanCodeCount) {
		return nil, false
	}
	return o.ShoulderEndRatePlanCodeCount, true
}

// HasShoulderEndRatePlanCodeCount returns a boolean if a field has been set.
func (o *BlockDetailsTypePrimaryRatePlanCodes) HasShoulderEndRatePlanCodeCount() bool {
	if o != nil && !IsNil(o.ShoulderEndRatePlanCodeCount) {
		return true
	}

	return false
}

// SetShoulderEndRatePlanCodeCount gets a reference to the given float32 and assigns it to the ShoulderEndRatePlanCodeCount field.
func (o *BlockDetailsTypePrimaryRatePlanCodes) SetShoulderEndRatePlanCodeCount(v float32) {
	o.ShoulderEndRatePlanCodeCount = &v
}

func (o BlockDetailsTypePrimaryRatePlanCodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockDetailsTypePrimaryRatePlanCodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockRatePlanCode) {
		toSerialize["blockRatePlanCode"] = o.BlockRatePlanCode
	}
	if !IsNil(o.ShoulderStartRatePlanCode) {
		toSerialize["shoulderStartRatePlanCode"] = o.ShoulderStartRatePlanCode
	}
	if !IsNil(o.ShoulderEndRatePlanCode) {
		toSerialize["shoulderEndRatePlanCode"] = o.ShoulderEndRatePlanCode
	}
	if !IsNil(o.BlockRatePlanCodeCount) {
		toSerialize["blockRatePlanCodeCount"] = o.BlockRatePlanCodeCount
	}
	if !IsNil(o.ShoulderStartRatePlanCodeCount) {
		toSerialize["shoulderStartRatePlanCodeCount"] = o.ShoulderStartRatePlanCodeCount
	}
	if !IsNil(o.ShoulderEndRatePlanCodeCount) {
		toSerialize["shoulderEndRatePlanCodeCount"] = o.ShoulderEndRatePlanCodeCount
	}
	return toSerialize, nil
}

type NullableBlockDetailsTypePrimaryRatePlanCodes struct {
	value *BlockDetailsTypePrimaryRatePlanCodes
	isSet bool
}

func (v NullableBlockDetailsTypePrimaryRatePlanCodes) Get() *BlockDetailsTypePrimaryRatePlanCodes {
	return v.value
}

func (v *NullableBlockDetailsTypePrimaryRatePlanCodes) Set(val *BlockDetailsTypePrimaryRatePlanCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockDetailsTypePrimaryRatePlanCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockDetailsTypePrimaryRatePlanCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockDetailsTypePrimaryRatePlanCodes(val *BlockDetailsTypePrimaryRatePlanCodes) *NullableBlockDetailsTypePrimaryRatePlanCodes {
	return &NullableBlockDetailsTypePrimaryRatePlanCodes{value: val, isSet: true}
}

func (v NullableBlockDetailsTypePrimaryRatePlanCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockDetailsTypePrimaryRatePlanCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


