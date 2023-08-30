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

// checks if the BlockRestrictionsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockRestrictionsType{}

// BlockRestrictionsType List of restrictions belonging to a block.
type BlockRestrictionsType struct {
	BlockId *BlockId `json:"blockId,omitempty"`
	// Block restriction details.
	BlockRestriction []BlockRestrictionType `json:"blockRestriction,omitempty"`
	// Hotel to which the block belongs to.
	HotelId *string `json:"hotelId,omitempty"`
}

// NewBlockRestrictionsType instantiates a new BlockRestrictionsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockRestrictionsType() *BlockRestrictionsType {
	this := BlockRestrictionsType{}
	return &this
}

// NewBlockRestrictionsTypeWithDefaults instantiates a new BlockRestrictionsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockRestrictionsTypeWithDefaults() *BlockRestrictionsType {
	this := BlockRestrictionsType{}
	return &this
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *BlockRestrictionsType) GetBlockId() BlockId {
	if o == nil || IsNil(o.BlockId) {
		var ret BlockId
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockRestrictionsType) GetBlockIdOk() (*BlockId, bool) {
	if o == nil || IsNil(o.BlockId) {
		return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *BlockRestrictionsType) HasBlockId() bool {
	if o != nil && !IsNil(o.BlockId) {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given BlockId and assigns it to the BlockId field.
func (o *BlockRestrictionsType) SetBlockId(v BlockId) {
	o.BlockId = &v
}

// GetBlockRestriction returns the BlockRestriction field value if set, zero value otherwise.
func (o *BlockRestrictionsType) GetBlockRestriction() []BlockRestrictionType {
	if o == nil || IsNil(o.BlockRestriction) {
		var ret []BlockRestrictionType
		return ret
	}
	return o.BlockRestriction
}

// GetBlockRestrictionOk returns a tuple with the BlockRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockRestrictionsType) GetBlockRestrictionOk() ([]BlockRestrictionType, bool) {
	if o == nil || IsNil(o.BlockRestriction) {
		return nil, false
	}
	return o.BlockRestriction, true
}

// HasBlockRestriction returns a boolean if a field has been set.
func (o *BlockRestrictionsType) HasBlockRestriction() bool {
	if o != nil && !IsNil(o.BlockRestriction) {
		return true
	}

	return false
}

// SetBlockRestriction gets a reference to the given []BlockRestrictionType and assigns it to the BlockRestriction field.
func (o *BlockRestrictionsType) SetBlockRestriction(v []BlockRestrictionType) {
	o.BlockRestriction = v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *BlockRestrictionsType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockRestrictionsType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *BlockRestrictionsType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *BlockRestrictionsType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o BlockRestrictionsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockRestrictionsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockId) {
		toSerialize["blockId"] = o.BlockId
	}
	if !IsNil(o.BlockRestriction) {
		toSerialize["blockRestriction"] = o.BlockRestriction
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableBlockRestrictionsType struct {
	value *BlockRestrictionsType
	isSet bool
}

func (v NullableBlockRestrictionsType) Get() *BlockRestrictionsType {
	return v.value
}

func (v *NullableBlockRestrictionsType) Set(val *BlockRestrictionsType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockRestrictionsType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockRestrictionsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockRestrictionsType(val *BlockRestrictionsType) *NullableBlockRestrictionsType {
	return &NullableBlockRestrictionsType{value: val, isSet: true}
}

func (v NullableBlockRestrictionsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockRestrictionsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


