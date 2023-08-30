/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the ParagraphType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParagraphType{}

// ParagraphType An indication of a new paragraph for a sub-section of a formatted text message.
type ParagraphType struct {
	Text *FormattedTextTextType `json:"text,omitempty"`
}

// NewParagraphType instantiates a new ParagraphType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParagraphType() *ParagraphType {
	this := ParagraphType{}
	return &this
}

// NewParagraphTypeWithDefaults instantiates a new ParagraphType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParagraphTypeWithDefaults() *ParagraphType {
	this := ParagraphType{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ParagraphType) GetText() FormattedTextTextType {
	if o == nil || IsNil(o.Text) {
		var ret FormattedTextTextType
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphType) GetTextOk() (*FormattedTextTextType, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ParagraphType) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given FormattedTextTextType and assigns it to the Text field.
func (o *ParagraphType) SetText(v FormattedTextTextType) {
	o.Text = &v
}

func (o ParagraphType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParagraphType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableParagraphType struct {
	value *ParagraphType
	isSet bool
}

func (v NullableParagraphType) Get() *ParagraphType {
	return v.value
}

func (v *NullableParagraphType) Set(val *ParagraphType) {
	v.value = val
	v.isSet = true
}

func (v NullableParagraphType) IsSet() bool {
	return v.isSet
}

func (v *NullableParagraphType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParagraphType(val *ParagraphType) *NullableParagraphType {
	return &NullableParagraphType{value: val, isSet: true}
}

func (v NullableParagraphType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParagraphType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


