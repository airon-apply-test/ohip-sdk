/*
OPERA Cloud Content Service

Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package medcfg

import (
	"encoding/json"
)

// checks if the FileAttachmentType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileAttachmentType{}

// FileAttachmentType ID that uniquely determines file attachment
type FileAttachmentType struct {
	AttachmentId *UniqueIDType `json:"attachmentId,omitempty"`
	// File attachment in base64 binary format
	FileAttachment *string `json:"fileAttachment,omitempty"`
	// Name of the attachment
	FileName *string `json:"fileName,omitempty"`
}

// NewFileAttachmentType instantiates a new FileAttachmentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileAttachmentType() *FileAttachmentType {
	this := FileAttachmentType{}
	return &this
}

// NewFileAttachmentTypeWithDefaults instantiates a new FileAttachmentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileAttachmentTypeWithDefaults() *FileAttachmentType {
	this := FileAttachmentType{}
	return &this
}

// GetAttachmentId returns the AttachmentId field value if set, zero value otherwise.
func (o *FileAttachmentType) GetAttachmentId() UniqueIDType {
	if o == nil || IsNil(o.AttachmentId) {
		var ret UniqueIDType
		return ret
	}
	return *o.AttachmentId
}

// GetAttachmentIdOk returns a tuple with the AttachmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAttachmentType) GetAttachmentIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.AttachmentId) {
		return nil, false
	}
	return o.AttachmentId, true
}

// HasAttachmentId returns a boolean if a field has been set.
func (o *FileAttachmentType) HasAttachmentId() bool {
	if o != nil && !IsNil(o.AttachmentId) {
		return true
	}

	return false
}

// SetAttachmentId gets a reference to the given UniqueIDType and assigns it to the AttachmentId field.
func (o *FileAttachmentType) SetAttachmentId(v UniqueIDType) {
	o.AttachmentId = &v
}

// GetFileAttachment returns the FileAttachment field value if set, zero value otherwise.
func (o *FileAttachmentType) GetFileAttachment() string {
	if o == nil || IsNil(o.FileAttachment) {
		var ret string
		return ret
	}
	return *o.FileAttachment
}

// GetFileAttachmentOk returns a tuple with the FileAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAttachmentType) GetFileAttachmentOk() (*string, bool) {
	if o == nil || IsNil(o.FileAttachment) {
		return nil, false
	}
	return o.FileAttachment, true
}

// HasFileAttachment returns a boolean if a field has been set.
func (o *FileAttachmentType) HasFileAttachment() bool {
	if o != nil && !IsNil(o.FileAttachment) {
		return true
	}

	return false
}

// SetFileAttachment gets a reference to the given string and assigns it to the FileAttachment field.
func (o *FileAttachmentType) SetFileAttachment(v string) {
	o.FileAttachment = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *FileAttachmentType) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAttachmentType) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *FileAttachmentType) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *FileAttachmentType) SetFileName(v string) {
	o.FileName = &v
}

func (o FileAttachmentType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileAttachmentType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttachmentId) {
		toSerialize["attachmentId"] = o.AttachmentId
	}
	if !IsNil(o.FileAttachment) {
		toSerialize["fileAttachment"] = o.FileAttachment
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	return toSerialize, nil
}

type NullableFileAttachmentType struct {
	value *FileAttachmentType
	isSet bool
}

func (v NullableFileAttachmentType) Get() *FileAttachmentType {
	return v.value
}

func (v *NullableFileAttachmentType) Set(val *FileAttachmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileAttachmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileAttachmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileAttachmentType(val *FileAttachmentType) *NullableFileAttachmentType {
	return &NullableFileAttachmentType{value: val, isSet: true}
}

func (v NullableFileAttachmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileAttachmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


