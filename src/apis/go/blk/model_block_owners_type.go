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

// checks if the BlockOwnersType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockOwnersType{}

// BlockOwnersType Contains a list of block owners.
type BlockOwnersType struct {
	Owner []BlockOwnerType `json:"owner,omitempty"`
	// When this flag is true, the logged in user cannot modify the existing block owners for the current block.
	LockBlockOwners *bool `json:"lockBlockOwners,omitempty"`
	// When this flag is true, the logged in user cannot modify the existing room owners for the current block.
	LockRoomsOwners *bool `json:"lockRoomsOwners,omitempty"`
	// When this flag is true, the logged in user cannot modify the existing catering owners for the current block.
	LockCateringOwners *bool `json:"lockCateringOwners,omitempty"`
}

// NewBlockOwnersType instantiates a new BlockOwnersType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockOwnersType() *BlockOwnersType {
	this := BlockOwnersType{}
	return &this
}

// NewBlockOwnersTypeWithDefaults instantiates a new BlockOwnersType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockOwnersTypeWithDefaults() *BlockOwnersType {
	this := BlockOwnersType{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BlockOwnersType) GetOwner() []BlockOwnerType {
	if o == nil || IsNil(o.Owner) {
		var ret []BlockOwnerType
		return ret
	}
	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockOwnersType) GetOwnerOk() ([]BlockOwnerType, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BlockOwnersType) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given []BlockOwnerType and assigns it to the Owner field.
func (o *BlockOwnersType) SetOwner(v []BlockOwnerType) {
	o.Owner = v
}

// GetLockBlockOwners returns the LockBlockOwners field value if set, zero value otherwise.
func (o *BlockOwnersType) GetLockBlockOwners() bool {
	if o == nil || IsNil(o.LockBlockOwners) {
		var ret bool
		return ret
	}
	return *o.LockBlockOwners
}

// GetLockBlockOwnersOk returns a tuple with the LockBlockOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockOwnersType) GetLockBlockOwnersOk() (*bool, bool) {
	if o == nil || IsNil(o.LockBlockOwners) {
		return nil, false
	}
	return o.LockBlockOwners, true
}

// HasLockBlockOwners returns a boolean if a field has been set.
func (o *BlockOwnersType) HasLockBlockOwners() bool {
	if o != nil && !IsNil(o.LockBlockOwners) {
		return true
	}

	return false
}

// SetLockBlockOwners gets a reference to the given bool and assigns it to the LockBlockOwners field.
func (o *BlockOwnersType) SetLockBlockOwners(v bool) {
	o.LockBlockOwners = &v
}

// GetLockRoomsOwners returns the LockRoomsOwners field value if set, zero value otherwise.
func (o *BlockOwnersType) GetLockRoomsOwners() bool {
	if o == nil || IsNil(o.LockRoomsOwners) {
		var ret bool
		return ret
	}
	return *o.LockRoomsOwners
}

// GetLockRoomsOwnersOk returns a tuple with the LockRoomsOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockOwnersType) GetLockRoomsOwnersOk() (*bool, bool) {
	if o == nil || IsNil(o.LockRoomsOwners) {
		return nil, false
	}
	return o.LockRoomsOwners, true
}

// HasLockRoomsOwners returns a boolean if a field has been set.
func (o *BlockOwnersType) HasLockRoomsOwners() bool {
	if o != nil && !IsNil(o.LockRoomsOwners) {
		return true
	}

	return false
}

// SetLockRoomsOwners gets a reference to the given bool and assigns it to the LockRoomsOwners field.
func (o *BlockOwnersType) SetLockRoomsOwners(v bool) {
	o.LockRoomsOwners = &v
}

// GetLockCateringOwners returns the LockCateringOwners field value if set, zero value otherwise.
func (o *BlockOwnersType) GetLockCateringOwners() bool {
	if o == nil || IsNil(o.LockCateringOwners) {
		var ret bool
		return ret
	}
	return *o.LockCateringOwners
}

// GetLockCateringOwnersOk returns a tuple with the LockCateringOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockOwnersType) GetLockCateringOwnersOk() (*bool, bool) {
	if o == nil || IsNil(o.LockCateringOwners) {
		return nil, false
	}
	return o.LockCateringOwners, true
}

// HasLockCateringOwners returns a boolean if a field has been set.
func (o *BlockOwnersType) HasLockCateringOwners() bool {
	if o != nil && !IsNil(o.LockCateringOwners) {
		return true
	}

	return false
}

// SetLockCateringOwners gets a reference to the given bool and assigns it to the LockCateringOwners field.
func (o *BlockOwnersType) SetLockCateringOwners(v bool) {
	o.LockCateringOwners = &v
}

func (o BlockOwnersType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockOwnersType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.LockBlockOwners) {
		toSerialize["lockBlockOwners"] = o.LockBlockOwners
	}
	if !IsNil(o.LockRoomsOwners) {
		toSerialize["lockRoomsOwners"] = o.LockRoomsOwners
	}
	if !IsNil(o.LockCateringOwners) {
		toSerialize["lockCateringOwners"] = o.LockCateringOwners
	}
	return toSerialize, nil
}

type NullableBlockOwnersType struct {
	value *BlockOwnersType
	isSet bool
}

func (v NullableBlockOwnersType) Get() *BlockOwnersType {
	return v.value
}

func (v *NullableBlockOwnersType) Set(val *BlockOwnersType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockOwnersType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockOwnersType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockOwnersType(val *BlockOwnersType) *NullableBlockOwnersType {
	return &NullableBlockOwnersType{value: val, isSet: true}
}

func (v NullableBlockOwnersType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockOwnersType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


