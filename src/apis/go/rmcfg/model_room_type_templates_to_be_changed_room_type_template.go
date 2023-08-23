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

// checks if the RoomTypeTemplatesToBeChangedRoomTypeTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomTypeTemplatesToBeChangedRoomTypeTemplate{}

// RoomTypeTemplatesToBeChangedRoomTypeTemplate Room Type template details to be changed.
type RoomTypeTemplatesToBeChangedRoomTypeTemplate struct {
	RoomTypeTemplateDetails *RoomTypeType `json:"roomTypeTemplateDetails,omitempty"`
}

// NewRoomTypeTemplatesToBeChangedRoomTypeTemplate instantiates a new RoomTypeTemplatesToBeChangedRoomTypeTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomTypeTemplatesToBeChangedRoomTypeTemplate() *RoomTypeTemplatesToBeChangedRoomTypeTemplate {
	this := RoomTypeTemplatesToBeChangedRoomTypeTemplate{}
	return &this
}

// NewRoomTypeTemplatesToBeChangedRoomTypeTemplateWithDefaults instantiates a new RoomTypeTemplatesToBeChangedRoomTypeTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomTypeTemplatesToBeChangedRoomTypeTemplateWithDefaults() *RoomTypeTemplatesToBeChangedRoomTypeTemplate {
	this := RoomTypeTemplatesToBeChangedRoomTypeTemplate{}
	return &this
}

// GetRoomTypeTemplateDetails returns the RoomTypeTemplateDetails field value if set, zero value otherwise.
func (o *RoomTypeTemplatesToBeChangedRoomTypeTemplate) GetRoomTypeTemplateDetails() RoomTypeType {
	if o == nil || IsNil(o.RoomTypeTemplateDetails) {
		var ret RoomTypeType
		return ret
	}
	return *o.RoomTypeTemplateDetails
}

// GetRoomTypeTemplateDetailsOk returns a tuple with the RoomTypeTemplateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesToBeChangedRoomTypeTemplate) GetRoomTypeTemplateDetailsOk() (*RoomTypeType, bool) {
	if o == nil || IsNil(o.RoomTypeTemplateDetails) {
		return nil, false
	}
	return o.RoomTypeTemplateDetails, true
}

// HasRoomTypeTemplateDetails returns a boolean if a field has been set.
func (o *RoomTypeTemplatesToBeChangedRoomTypeTemplate) HasRoomTypeTemplateDetails() bool {
	if o != nil && !IsNil(o.RoomTypeTemplateDetails) {
		return true
	}

	return false
}

// SetRoomTypeTemplateDetails gets a reference to the given RoomTypeType and assigns it to the RoomTypeTemplateDetails field.
func (o *RoomTypeTemplatesToBeChangedRoomTypeTemplate) SetRoomTypeTemplateDetails(v RoomTypeType) {
	o.RoomTypeTemplateDetails = &v
}

func (o RoomTypeTemplatesToBeChangedRoomTypeTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomTypeTemplatesToBeChangedRoomTypeTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomTypeTemplateDetails) {
		toSerialize["roomTypeTemplateDetails"] = o.RoomTypeTemplateDetails
	}
	return toSerialize, nil
}

type NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate struct {
	value *RoomTypeTemplatesToBeChangedRoomTypeTemplate
	isSet bool
}

func (v NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate) Get() *RoomTypeTemplatesToBeChangedRoomTypeTemplate {
	return v.value
}

func (v *NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate) Set(val *RoomTypeTemplatesToBeChangedRoomTypeTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomTypeTemplatesToBeChangedRoomTypeTemplate(val *RoomTypeTemplatesToBeChangedRoomTypeTemplate) *NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate {
	return &NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate{value: val, isSet: true}
}

func (v NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomTypeTemplatesToBeChangedRoomTypeTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


