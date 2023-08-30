/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package act

import (
	"encoding/json"
)

// checks if the AttachmentType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentType{}

// AttachmentType Attached files.
type AttachmentType struct {
	// Name of the file.
	FileName *string `json:"fileName,omitempty"`
	// Size of the file.
	FileSize *int32 `json:"fileSize,omitempty"`
	// Description for the file.
	Description *string `json:"description,omitempty"`
	// Flag to say if attachment is available across properties.
	Global *bool `json:"global,omitempty"`
	// Hotel Code.
	HotelId *string `json:"hotelId,omitempty"`
	History *DateTimeStampGroupType `json:"history,omitempty"`
	// URL that identifies the location associated with the record identified by the UniqueID.
	Url *string `json:"url,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
	// The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
	Instance *string `json:"instance,omitempty"`
	// Used to identify the source of the identifier (e.g., IATA, ABTA).
	IdContext *string `json:"idContext,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
}

// NewAttachmentType instantiates a new AttachmentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentType() *AttachmentType {
	this := AttachmentType{}
	return &this
}

// NewAttachmentTypeWithDefaults instantiates a new AttachmentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentTypeWithDefaults() *AttachmentType {
	this := AttachmentType{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AttachmentType) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AttachmentType) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AttachmentType) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *AttachmentType) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *AttachmentType) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *AttachmentType) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttachmentType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttachmentType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttachmentType) SetDescription(v string) {
	o.Description = &v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *AttachmentType) GetGlobal() bool {
	if o == nil || IsNil(o.Global) {
		var ret bool
		return ret
	}
	return *o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.Global) {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *AttachmentType) HasGlobal() bool {
	if o != nil && !IsNil(o.Global) {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given bool and assigns it to the Global field.
func (o *AttachmentType) SetGlobal(v bool) {
	o.Global = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *AttachmentType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *AttachmentType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *AttachmentType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *AttachmentType) GetHistory() DateTimeStampGroupType {
	if o == nil || IsNil(o.History) {
		var ret DateTimeStampGroupType
		return ret
	}
	return *o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetHistoryOk() (*DateTimeStampGroupType, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *AttachmentType) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given DateTimeStampGroupType and assigns it to the History field.
func (o *AttachmentType) SetHistory(v DateTimeStampGroupType) {
	o.History = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AttachmentType) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AttachmentType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AttachmentType) SetUrl(v string) {
	o.Url = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AttachmentType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AttachmentType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AttachmentType) SetType(v string) {
	o.Type = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *AttachmentType) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *AttachmentType) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *AttachmentType) SetInstance(v string) {
	o.Instance = &v
}

// GetIdContext returns the IdContext field value if set, zero value otherwise.
func (o *AttachmentType) GetIdContext() string {
	if o == nil || IsNil(o.IdContext) {
		var ret string
		return ret
	}
	return *o.IdContext
}

// GetIdContextOk returns a tuple with the IdContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetIdContextOk() (*string, bool) {
	if o == nil || IsNil(o.IdContext) {
		return nil, false
	}
	return o.IdContext, true
}

// HasIdContext returns a boolean if a field has been set.
func (o *AttachmentType) HasIdContext() bool {
	if o != nil && !IsNil(o.IdContext) {
		return true
	}

	return false
}

// SetIdContext gets a reference to the given string and assigns it to the IdContext field.
func (o *AttachmentType) SetIdContext(v string) {
	o.IdContext = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AttachmentType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AttachmentType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AttachmentType) SetId(v string) {
	o.Id = &v
}

func (o AttachmentType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Global) {
		toSerialize["global"] = o.Global
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.IdContext) {
		toSerialize["idContext"] = o.IdContext
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAttachmentType struct {
	value *AttachmentType
	isSet bool
}

func (v NullableAttachmentType) Get() *AttachmentType {
	return v.value
}

func (v *NullableAttachmentType) Set(val *AttachmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentType(val *AttachmentType) *NullableAttachmentType {
	return &NullableAttachmentType{value: val, isSet: true}
}

func (v NullableAttachmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


