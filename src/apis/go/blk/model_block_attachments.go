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

// checks if the BlockAttachments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockAttachments{}

// BlockAttachments Return object to the request for information regarding block attachments.
type BlockAttachments struct {
	// Attachment List.
	BlockAttachments []AttachmentType `json:"blockAttachments,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBlockAttachments instantiates a new BlockAttachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockAttachments() *BlockAttachments {
	this := BlockAttachments{}
	return &this
}

// NewBlockAttachmentsWithDefaults instantiates a new BlockAttachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockAttachmentsWithDefaults() *BlockAttachments {
	this := BlockAttachments{}
	return &this
}

// GetBlockAttachments returns the BlockAttachments field value if set, zero value otherwise.
func (o *BlockAttachments) GetBlockAttachments() []AttachmentType {
	if o == nil || IsNil(o.BlockAttachments) {
		var ret []AttachmentType
		return ret
	}
	return o.BlockAttachments
}

// GetBlockAttachmentsOk returns a tuple with the BlockAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAttachments) GetBlockAttachmentsOk() ([]AttachmentType, bool) {
	if o == nil || IsNil(o.BlockAttachments) {
		return nil, false
	}
	return o.BlockAttachments, true
}

// HasBlockAttachments returns a boolean if a field has been set.
func (o *BlockAttachments) HasBlockAttachments() bool {
	if o != nil && !IsNil(o.BlockAttachments) {
		return true
	}

	return false
}

// SetBlockAttachments gets a reference to the given []AttachmentType and assigns it to the BlockAttachments field.
func (o *BlockAttachments) SetBlockAttachments(v []AttachmentType) {
	o.BlockAttachments = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BlockAttachments) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAttachments) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BlockAttachments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BlockAttachments) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BlockAttachments) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockAttachments) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BlockAttachments) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BlockAttachments) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o BlockAttachments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockAttachments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockAttachments) {
		toSerialize["blockAttachments"] = o.BlockAttachments
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBlockAttachments struct {
	value *BlockAttachments
	isSet bool
}

func (v NullableBlockAttachments) Get() *BlockAttachments {
	return v.value
}

func (v *NullableBlockAttachments) Set(val *BlockAttachments) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockAttachments) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockAttachments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockAttachments(val *BlockAttachments) *NullableBlockAttachments {
	return &NullableBlockAttachments{value: val, isSet: true}
}

func (v NullableBlockAttachments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockAttachments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


