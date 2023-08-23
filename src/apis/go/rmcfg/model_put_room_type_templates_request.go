/*
OPERA Cloud Room Configuration API

APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PutRoomTypeTemplatesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutRoomTypeTemplatesRequest{}

// PutRoomTypeTemplatesRequest struct for PutRoomTypeTemplatesRequest
type PutRoomTypeTemplatesRequest struct {
	RoomTypeTemplate *RoomTypeTemplatesToBeChangedRoomTypeTemplate `json:"roomTypeTemplate,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutRoomTypeTemplatesRequest instantiates a new PutRoomTypeTemplatesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutRoomTypeTemplatesRequest() *PutRoomTypeTemplatesRequest {
	this := PutRoomTypeTemplatesRequest{}
	return &this
}

// NewPutRoomTypeTemplatesRequestWithDefaults instantiates a new PutRoomTypeTemplatesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutRoomTypeTemplatesRequestWithDefaults() *PutRoomTypeTemplatesRequest {
	this := PutRoomTypeTemplatesRequest{}
	return &this
}

// GetRoomTypeTemplate returns the RoomTypeTemplate field value if set, zero value otherwise.
func (o *PutRoomTypeTemplatesRequest) GetRoomTypeTemplate() RoomTypeTemplatesToBeChangedRoomTypeTemplate {
	if o == nil || IsNil(o.RoomTypeTemplate) {
		var ret RoomTypeTemplatesToBeChangedRoomTypeTemplate
		return ret
	}
	return *o.RoomTypeTemplate
}

// GetRoomTypeTemplateOk returns a tuple with the RoomTypeTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutRoomTypeTemplatesRequest) GetRoomTypeTemplateOk() (*RoomTypeTemplatesToBeChangedRoomTypeTemplate, bool) {
	if o == nil || IsNil(o.RoomTypeTemplate) {
		return nil, false
	}
	return o.RoomTypeTemplate, true
}

// HasRoomTypeTemplate returns a boolean if a field has been set.
func (o *PutRoomTypeTemplatesRequest) HasRoomTypeTemplate() bool {
	if o != nil && !IsNil(o.RoomTypeTemplate) {
		return true
	}

	return false
}

// SetRoomTypeTemplate gets a reference to the given RoomTypeTemplatesToBeChangedRoomTypeTemplate and assigns it to the RoomTypeTemplate field.
func (o *PutRoomTypeTemplatesRequest) SetRoomTypeTemplate(v RoomTypeTemplatesToBeChangedRoomTypeTemplate) {
	o.RoomTypeTemplate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutRoomTypeTemplatesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutRoomTypeTemplatesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutRoomTypeTemplatesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutRoomTypeTemplatesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutRoomTypeTemplatesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutRoomTypeTemplatesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutRoomTypeTemplatesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutRoomTypeTemplatesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutRoomTypeTemplatesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutRoomTypeTemplatesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomTypeTemplate) {
		toSerialize["roomTypeTemplate"] = o.RoomTypeTemplate
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutRoomTypeTemplatesRequest struct {
	value *PutRoomTypeTemplatesRequest
	isSet bool
}

func (v NullablePutRoomTypeTemplatesRequest) Get() *PutRoomTypeTemplatesRequest {
	return v.value
}

func (v *NullablePutRoomTypeTemplatesRequest) Set(val *PutRoomTypeTemplatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutRoomTypeTemplatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutRoomTypeTemplatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutRoomTypeTemplatesRequest(val *PutRoomTypeTemplatesRequest) *NullablePutRoomTypeTemplatesRequest {
	return &NullablePutRoomTypeTemplatesRequest{value: val, isSet: true}
}

func (v NullablePutRoomTypeTemplatesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutRoomTypeTemplatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


