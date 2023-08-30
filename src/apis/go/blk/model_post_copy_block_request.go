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

// checks if the PostCopyBlockRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCopyBlockRequest{}

// PostCopyBlockRequest struct for PostCopyBlockRequest
type PostCopyBlockRequest struct {
	Criteria *CopyBlockType `json:"criteria,omitempty"`
	ResponseInstruction *ResponseInstructionType `json:"responseInstruction,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostCopyBlockRequest instantiates a new PostCopyBlockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCopyBlockRequest() *PostCopyBlockRequest {
	this := PostCopyBlockRequest{}
	return &this
}

// NewPostCopyBlockRequestWithDefaults instantiates a new PostCopyBlockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCopyBlockRequestWithDefaults() *PostCopyBlockRequest {
	this := PostCopyBlockRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PostCopyBlockRequest) GetCriteria() CopyBlockType {
	if o == nil || IsNil(o.Criteria) {
		var ret CopyBlockType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCopyBlockRequest) GetCriteriaOk() (*CopyBlockType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PostCopyBlockRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given CopyBlockType and assigns it to the Criteria field.
func (o *PostCopyBlockRequest) SetCriteria(v CopyBlockType) {
	o.Criteria = &v
}

// GetResponseInstruction returns the ResponseInstruction field value if set, zero value otherwise.
func (o *PostCopyBlockRequest) GetResponseInstruction() ResponseInstructionType {
	if o == nil || IsNil(o.ResponseInstruction) {
		var ret ResponseInstructionType
		return ret
	}
	return *o.ResponseInstruction
}

// GetResponseInstructionOk returns a tuple with the ResponseInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCopyBlockRequest) GetResponseInstructionOk() (*ResponseInstructionType, bool) {
	if o == nil || IsNil(o.ResponseInstruction) {
		return nil, false
	}
	return o.ResponseInstruction, true
}

// HasResponseInstruction returns a boolean if a field has been set.
func (o *PostCopyBlockRequest) HasResponseInstruction() bool {
	if o != nil && !IsNil(o.ResponseInstruction) {
		return true
	}

	return false
}

// SetResponseInstruction gets a reference to the given ResponseInstructionType and assigns it to the ResponseInstruction field.
func (o *PostCopyBlockRequest) SetResponseInstruction(v ResponseInstructionType) {
	o.ResponseInstruction = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostCopyBlockRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCopyBlockRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostCopyBlockRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostCopyBlockRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostCopyBlockRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCopyBlockRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostCopyBlockRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostCopyBlockRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostCopyBlockRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCopyBlockRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.ResponseInstruction) {
		toSerialize["responseInstruction"] = o.ResponseInstruction
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostCopyBlockRequest struct {
	value *PostCopyBlockRequest
	isSet bool
}

func (v NullablePostCopyBlockRequest) Get() *PostCopyBlockRequest {
	return v.value
}

func (v *NullablePostCopyBlockRequest) Set(val *PostCopyBlockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCopyBlockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCopyBlockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCopyBlockRequest(val *PostCopyBlockRequest) *NullablePostCopyBlockRequest {
	return &NullablePostCopyBlockRequest{value: val, isSet: true}
}

func (v NullablePostCopyBlockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCopyBlockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


