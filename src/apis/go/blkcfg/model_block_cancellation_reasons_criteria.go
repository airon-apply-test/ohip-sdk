/*
OPERA Cloud Block Configuration API

APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BlockCancellationReasonsCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockCancellationReasonsCriteria{}

// BlockCancellationReasonsCriteria Request object for creating Block Cancellation Reasons.
type BlockCancellationReasonsCriteria struct {
	// List of Block Cancellation Reasons.
	BlockCancellationReasons []BlockCancellationReasonType `json:"blockCancellationReasons,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBlockCancellationReasonsCriteria instantiates a new BlockCancellationReasonsCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockCancellationReasonsCriteria() *BlockCancellationReasonsCriteria {
	this := BlockCancellationReasonsCriteria{}
	return &this
}

// NewBlockCancellationReasonsCriteriaWithDefaults instantiates a new BlockCancellationReasonsCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockCancellationReasonsCriteriaWithDefaults() *BlockCancellationReasonsCriteria {
	this := BlockCancellationReasonsCriteria{}
	return &this
}

// GetBlockCancellationReasons returns the BlockCancellationReasons field value if set, zero value otherwise.
func (o *BlockCancellationReasonsCriteria) GetBlockCancellationReasons() []BlockCancellationReasonType {
	if o == nil || IsNil(o.BlockCancellationReasons) {
		var ret []BlockCancellationReasonType
		return ret
	}
	return o.BlockCancellationReasons
}

// GetBlockCancellationReasonsOk returns a tuple with the BlockCancellationReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCancellationReasonsCriteria) GetBlockCancellationReasonsOk() ([]BlockCancellationReasonType, bool) {
	if o == nil || IsNil(o.BlockCancellationReasons) {
		return nil, false
	}
	return o.BlockCancellationReasons, true
}

// HasBlockCancellationReasons returns a boolean if a field has been set.
func (o *BlockCancellationReasonsCriteria) HasBlockCancellationReasons() bool {
	if o != nil && !IsNil(o.BlockCancellationReasons) {
		return true
	}

	return false
}

// SetBlockCancellationReasons gets a reference to the given []BlockCancellationReasonType and assigns it to the BlockCancellationReasons field.
func (o *BlockCancellationReasonsCriteria) SetBlockCancellationReasons(v []BlockCancellationReasonType) {
	o.BlockCancellationReasons = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BlockCancellationReasonsCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCancellationReasonsCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BlockCancellationReasonsCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BlockCancellationReasonsCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BlockCancellationReasonsCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCancellationReasonsCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BlockCancellationReasonsCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BlockCancellationReasonsCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o BlockCancellationReasonsCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockCancellationReasonsCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockCancellationReasons) {
		toSerialize["blockCancellationReasons"] = o.BlockCancellationReasons
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBlockCancellationReasonsCriteria struct {
	value *BlockCancellationReasonsCriteria
	isSet bool
}

func (v NullableBlockCancellationReasonsCriteria) Get() *BlockCancellationReasonsCriteria {
	return v.value
}

func (v *NullableBlockCancellationReasonsCriteria) Set(val *BlockCancellationReasonsCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockCancellationReasonsCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockCancellationReasonsCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockCancellationReasonsCriteria(val *BlockCancellationReasonsCriteria) *NullableBlockCancellationReasonsCriteria {
	return &NullableBlockCancellationReasonsCriteria{value: val, isSet: true}
}

func (v NullableBlockCancellationReasonsCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockCancellationReasonsCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


