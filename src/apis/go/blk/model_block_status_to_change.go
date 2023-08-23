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

// checks if the BlockStatusToChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockStatusToChange{}

// BlockStatusToChange Request object for changing the booking status of the business block.
type BlockStatusToChange struct {
	// Contains the booking status of the business block.
	BlocksStatus map[string]interface{} `json:"blocksStatus,omitempty"`
	ChangeBlockStatus *ChangeBlockStatusType `json:"changeBlockStatus,omitempty"`
	// Indicator if the request is a verification on whether the block status can be changed.
	VerificationOnly *bool `json:"verificationOnly,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBlockStatusToChange instantiates a new BlockStatusToChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockStatusToChange() *BlockStatusToChange {
	this := BlockStatusToChange{}
	return &this
}

// NewBlockStatusToChangeWithDefaults instantiates a new BlockStatusToChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockStatusToChangeWithDefaults() *BlockStatusToChange {
	this := BlockStatusToChange{}
	return &this
}

// GetBlocksStatus returns the BlocksStatus field value if set, zero value otherwise.
func (o *BlockStatusToChange) GetBlocksStatus() map[string]interface{} {
	if o == nil || IsNil(o.BlocksStatus) {
		var ret map[string]interface{}
		return ret
	}
	return o.BlocksStatus
}

// GetBlocksStatusOk returns a tuple with the BlocksStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatusToChange) GetBlocksStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.BlocksStatus) {
		return map[string]interface{}{}, false
	}
	return o.BlocksStatus, true
}

// HasBlocksStatus returns a boolean if a field has been set.
func (o *BlockStatusToChange) HasBlocksStatus() bool {
	if o != nil && !IsNil(o.BlocksStatus) {
		return true
	}

	return false
}

// SetBlocksStatus gets a reference to the given map[string]interface{} and assigns it to the BlocksStatus field.
func (o *BlockStatusToChange) SetBlocksStatus(v map[string]interface{}) {
	o.BlocksStatus = v
}

// GetChangeBlockStatus returns the ChangeBlockStatus field value if set, zero value otherwise.
func (o *BlockStatusToChange) GetChangeBlockStatus() ChangeBlockStatusType {
	if o == nil || IsNil(o.ChangeBlockStatus) {
		var ret ChangeBlockStatusType
		return ret
	}
	return *o.ChangeBlockStatus
}

// GetChangeBlockStatusOk returns a tuple with the ChangeBlockStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatusToChange) GetChangeBlockStatusOk() (*ChangeBlockStatusType, bool) {
	if o == nil || IsNil(o.ChangeBlockStatus) {
		return nil, false
	}
	return o.ChangeBlockStatus, true
}

// HasChangeBlockStatus returns a boolean if a field has been set.
func (o *BlockStatusToChange) HasChangeBlockStatus() bool {
	if o != nil && !IsNil(o.ChangeBlockStatus) {
		return true
	}

	return false
}

// SetChangeBlockStatus gets a reference to the given ChangeBlockStatusType and assigns it to the ChangeBlockStatus field.
func (o *BlockStatusToChange) SetChangeBlockStatus(v ChangeBlockStatusType) {
	o.ChangeBlockStatus = &v
}

// GetVerificationOnly returns the VerificationOnly field value if set, zero value otherwise.
func (o *BlockStatusToChange) GetVerificationOnly() bool {
	if o == nil || IsNil(o.VerificationOnly) {
		var ret bool
		return ret
	}
	return *o.VerificationOnly
}

// GetVerificationOnlyOk returns a tuple with the VerificationOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatusToChange) GetVerificationOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.VerificationOnly) {
		return nil, false
	}
	return o.VerificationOnly, true
}

// HasVerificationOnly returns a boolean if a field has been set.
func (o *BlockStatusToChange) HasVerificationOnly() bool {
	if o != nil && !IsNil(o.VerificationOnly) {
		return true
	}

	return false
}

// SetVerificationOnly gets a reference to the given bool and assigns it to the VerificationOnly field.
func (o *BlockStatusToChange) SetVerificationOnly(v bool) {
	o.VerificationOnly = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BlockStatusToChange) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatusToChange) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BlockStatusToChange) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BlockStatusToChange) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BlockStatusToChange) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatusToChange) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BlockStatusToChange) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BlockStatusToChange) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o BlockStatusToChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockStatusToChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlocksStatus) {
		toSerialize["blocksStatus"] = o.BlocksStatus
	}
	if !IsNil(o.ChangeBlockStatus) {
		toSerialize["changeBlockStatus"] = o.ChangeBlockStatus
	}
	if !IsNil(o.VerificationOnly) {
		toSerialize["verificationOnly"] = o.VerificationOnly
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBlockStatusToChange struct {
	value *BlockStatusToChange
	isSet bool
}

func (v NullableBlockStatusToChange) Get() *BlockStatusToChange {
	return v.value
}

func (v *NullableBlockStatusToChange) Set(val *BlockStatusToChange) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockStatusToChange) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockStatusToChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockStatusToChange(val *BlockStatusToChange) *NullableBlockStatusToChange {
	return &NullableBlockStatusToChange{value: val, isSet: true}
}

func (v NullableBlockStatusToChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockStatusToChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


