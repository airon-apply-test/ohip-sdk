/*
OPERA Cloud Content Service

Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FileToUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileToUpload{}

// FileToUpload Request to upload a file attachment.
type FileToUpload struct {
	// Used for Character Strings, length 0 to 200.
	LinkId *string `json:"linkId,omitempty"`
	// Used for Character Strings, length 0 to 200.
	LinkType *string `json:"linkType,omitempty"`
	// Used for Character Strings, length 0 to 1000.
	FileName *string `json:"fileName,omitempty"`
	// Used for Character Strings, length 0 to 2000.
	Description *string `json:"description,omitempty"`
	// Used for Character Strings, length 0 to 200.
	UserName *string `json:"userName,omitempty"`
	// Used for Character Strings, length 0 to 200.
	HotelId *string `json:"hotelId,omitempty"`
	// Used for Character Strings, length 0 to 10.
	GlobalYN *string `json:"globalYN,omitempty"`
	// Attachment file in base64 binary format.
	FileAttachment *string `json:"fileAttachment,omitempty"`
	// Used for Character Strings, length 0 to 10.
	OverwriteExistingFileYN *string `json:"overwriteExistingFileYN,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewFileToUpload instantiates a new FileToUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileToUpload() *FileToUpload {
	this := FileToUpload{}
	return &this
}

// NewFileToUploadWithDefaults instantiates a new FileToUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileToUploadWithDefaults() *FileToUpload {
	this := FileToUpload{}
	return &this
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *FileToUpload) GetLinkId() string {
	if o == nil || IsNil(o.LinkId) {
		var ret string
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *FileToUpload) HasLinkId() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given string and assigns it to the LinkId field.
func (o *FileToUpload) SetLinkId(v string) {
	o.LinkId = &v
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *FileToUpload) GetLinkType() string {
	if o == nil || IsNil(o.LinkType) {
		var ret string
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetLinkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LinkType) {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *FileToUpload) HasLinkType() bool {
	if o != nil && !IsNil(o.LinkType) {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given string and assigns it to the LinkType field.
func (o *FileToUpload) SetLinkType(v string) {
	o.LinkType = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *FileToUpload) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *FileToUpload) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *FileToUpload) SetFileName(v string) {
	o.FileName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileToUpload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileToUpload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileToUpload) SetDescription(v string) {
	o.Description = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *FileToUpload) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *FileToUpload) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *FileToUpload) SetUserName(v string) {
	o.UserName = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *FileToUpload) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *FileToUpload) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *FileToUpload) SetHotelId(v string) {
	o.HotelId = &v
}

// GetGlobalYN returns the GlobalYN field value if set, zero value otherwise.
func (o *FileToUpload) GetGlobalYN() string {
	if o == nil || IsNil(o.GlobalYN) {
		var ret string
		return ret
	}
	return *o.GlobalYN
}

// GetGlobalYNOk returns a tuple with the GlobalYN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetGlobalYNOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalYN) {
		return nil, false
	}
	return o.GlobalYN, true
}

// HasGlobalYN returns a boolean if a field has been set.
func (o *FileToUpload) HasGlobalYN() bool {
	if o != nil && !IsNil(o.GlobalYN) {
		return true
	}

	return false
}

// SetGlobalYN gets a reference to the given string and assigns it to the GlobalYN field.
func (o *FileToUpload) SetGlobalYN(v string) {
	o.GlobalYN = &v
}

// GetFileAttachment returns the FileAttachment field value if set, zero value otherwise.
func (o *FileToUpload) GetFileAttachment() string {
	if o == nil || IsNil(o.FileAttachment) {
		var ret string
		return ret
	}
	return *o.FileAttachment
}

// GetFileAttachmentOk returns a tuple with the FileAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetFileAttachmentOk() (*string, bool) {
	if o == nil || IsNil(o.FileAttachment) {
		return nil, false
	}
	return o.FileAttachment, true
}

// HasFileAttachment returns a boolean if a field has been set.
func (o *FileToUpload) HasFileAttachment() bool {
	if o != nil && !IsNil(o.FileAttachment) {
		return true
	}

	return false
}

// SetFileAttachment gets a reference to the given string and assigns it to the FileAttachment field.
func (o *FileToUpload) SetFileAttachment(v string) {
	o.FileAttachment = &v
}

// GetOverwriteExistingFileYN returns the OverwriteExistingFileYN field value if set, zero value otherwise.
func (o *FileToUpload) GetOverwriteExistingFileYN() string {
	if o == nil || IsNil(o.OverwriteExistingFileYN) {
		var ret string
		return ret
	}
	return *o.OverwriteExistingFileYN
}

// GetOverwriteExistingFileYNOk returns a tuple with the OverwriteExistingFileYN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetOverwriteExistingFileYNOk() (*string, bool) {
	if o == nil || IsNil(o.OverwriteExistingFileYN) {
		return nil, false
	}
	return o.OverwriteExistingFileYN, true
}

// HasOverwriteExistingFileYN returns a boolean if a field has been set.
func (o *FileToUpload) HasOverwriteExistingFileYN() bool {
	if o != nil && !IsNil(o.OverwriteExistingFileYN) {
		return true
	}

	return false
}

// SetOverwriteExistingFileYN gets a reference to the given string and assigns it to the OverwriteExistingFileYN field.
func (o *FileToUpload) SetOverwriteExistingFileYN(v string) {
	o.OverwriteExistingFileYN = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FileToUpload) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FileToUpload) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *FileToUpload) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FileToUpload) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileToUpload) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FileToUpload) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *FileToUpload) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o FileToUpload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileToUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkId) {
		toSerialize["linkId"] = o.LinkId
	}
	if !IsNil(o.LinkType) {
		toSerialize["linkType"] = o.LinkType
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.GlobalYN) {
		toSerialize["globalYN"] = o.GlobalYN
	}
	if !IsNil(o.FileAttachment) {
		toSerialize["fileAttachment"] = o.FileAttachment
	}
	if !IsNil(o.OverwriteExistingFileYN) {
		toSerialize["overwriteExistingFileYN"] = o.OverwriteExistingFileYN
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableFileToUpload struct {
	value *FileToUpload
	isSet bool
}

func (v NullableFileToUpload) Get() *FileToUpload {
	return v.value
}

func (v *NullableFileToUpload) Set(val *FileToUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableFileToUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableFileToUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileToUpload(val *FileToUpload) *NullableFileToUpload {
	return &NullableFileToUpload{value: val, isSet: true}
}

func (v NullableFileToUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileToUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


