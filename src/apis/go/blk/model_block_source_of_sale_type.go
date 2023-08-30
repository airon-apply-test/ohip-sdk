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

// checks if the BlockSourceOfSaleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockSourceOfSaleType{}

// BlockSourceOfSaleType Point of Sale of Block. Identifies the entity/channel who made the block reservation.
type BlockSourceOfSaleType struct {
	SourceCode *SourceCodeInfoType `json:"sourceCode,omitempty"`
	// The of entity/channel where this business block originated.
	SourceType *string `json:"sourceType,omitempty"`
	// The of entity/channel where this business block originated.
	SourceTypeDescription *string `json:"sourceTypeDescription,omitempty"`
}

// NewBlockSourceOfSaleType instantiates a new BlockSourceOfSaleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSourceOfSaleType() *BlockSourceOfSaleType {
	this := BlockSourceOfSaleType{}
	return &this
}

// NewBlockSourceOfSaleTypeWithDefaults instantiates a new BlockSourceOfSaleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSourceOfSaleTypeWithDefaults() *BlockSourceOfSaleType {
	this := BlockSourceOfSaleType{}
	return &this
}

// GetSourceCode returns the SourceCode field value if set, zero value otherwise.
func (o *BlockSourceOfSaleType) GetSourceCode() SourceCodeInfoType {
	if o == nil || IsNil(o.SourceCode) {
		var ret SourceCodeInfoType
		return ret
	}
	return *o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSourceOfSaleType) GetSourceCodeOk() (*SourceCodeInfoType, bool) {
	if o == nil || IsNil(o.SourceCode) {
		return nil, false
	}
	return o.SourceCode, true
}

// HasSourceCode returns a boolean if a field has been set.
func (o *BlockSourceOfSaleType) HasSourceCode() bool {
	if o != nil && !IsNil(o.SourceCode) {
		return true
	}

	return false
}

// SetSourceCode gets a reference to the given SourceCodeInfoType and assigns it to the SourceCode field.
func (o *BlockSourceOfSaleType) SetSourceCode(v SourceCodeInfoType) {
	o.SourceCode = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *BlockSourceOfSaleType) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSourceOfSaleType) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *BlockSourceOfSaleType) HasSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *BlockSourceOfSaleType) SetSourceType(v string) {
	o.SourceType = &v
}

// GetSourceTypeDescription returns the SourceTypeDescription field value if set, zero value otherwise.
func (o *BlockSourceOfSaleType) GetSourceTypeDescription() string {
	if o == nil || IsNil(o.SourceTypeDescription) {
		var ret string
		return ret
	}
	return *o.SourceTypeDescription
}

// GetSourceTypeDescriptionOk returns a tuple with the SourceTypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSourceOfSaleType) GetSourceTypeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SourceTypeDescription) {
		return nil, false
	}
	return o.SourceTypeDescription, true
}

// HasSourceTypeDescription returns a boolean if a field has been set.
func (o *BlockSourceOfSaleType) HasSourceTypeDescription() bool {
	if o != nil && !IsNil(o.SourceTypeDescription) {
		return true
	}

	return false
}

// SetSourceTypeDescription gets a reference to the given string and assigns it to the SourceTypeDescription field.
func (o *BlockSourceOfSaleType) SetSourceTypeDescription(v string) {
	o.SourceTypeDescription = &v
}

func (o BlockSourceOfSaleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockSourceOfSaleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceCode) {
		toSerialize["sourceCode"] = o.SourceCode
	}
	if !IsNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !IsNil(o.SourceTypeDescription) {
		toSerialize["sourceTypeDescription"] = o.SourceTypeDescription
	}
	return toSerialize, nil
}

type NullableBlockSourceOfSaleType struct {
	value *BlockSourceOfSaleType
	isSet bool
}

func (v NullableBlockSourceOfSaleType) Get() *BlockSourceOfSaleType {
	return v.value
}

func (v *NullableBlockSourceOfSaleType) Set(val *BlockSourceOfSaleType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSourceOfSaleType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSourceOfSaleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSourceOfSaleType(val *BlockSourceOfSaleType) *NullableBlockSourceOfSaleType {
	return &NullableBlockSourceOfSaleType{value: val, isSet: true}
}

func (v NullableBlockSourceOfSaleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSourceOfSaleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


