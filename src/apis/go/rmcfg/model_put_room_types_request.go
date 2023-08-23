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

// checks if the PutRoomTypesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutRoomTypesRequest{}

// PutRoomTypesRequest struct for PutRoomTypesRequest
type PutRoomTypesRequest struct {
	RoomType *RoomTypesToBeChangedRoomType `json:"roomType,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutRoomTypesRequest instantiates a new PutRoomTypesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutRoomTypesRequest() *PutRoomTypesRequest {
	this := PutRoomTypesRequest{}
	return &this
}

// NewPutRoomTypesRequestWithDefaults instantiates a new PutRoomTypesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutRoomTypesRequestWithDefaults() *PutRoomTypesRequest {
	this := PutRoomTypesRequest{}
	return &this
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *PutRoomTypesRequest) GetRoomType() RoomTypesToBeChangedRoomType {
	if o == nil || IsNil(o.RoomType) {
		var ret RoomTypesToBeChangedRoomType
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutRoomTypesRequest) GetRoomTypeOk() (*RoomTypesToBeChangedRoomType, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *PutRoomTypesRequest) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given RoomTypesToBeChangedRoomType and assigns it to the RoomType field.
func (o *PutRoomTypesRequest) SetRoomType(v RoomTypesToBeChangedRoomType) {
	o.RoomType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutRoomTypesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutRoomTypesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutRoomTypesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutRoomTypesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutRoomTypesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutRoomTypesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutRoomTypesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutRoomTypesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutRoomTypesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutRoomTypesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutRoomTypesRequest struct {
	value *PutRoomTypesRequest
	isSet bool
}

func (v NullablePutRoomTypesRequest) Get() *PutRoomTypesRequest {
	return v.value
}

func (v *NullablePutRoomTypesRequest) Set(val *PutRoomTypesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutRoomTypesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutRoomTypesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutRoomTypesRequest(val *PutRoomTypesRequest) *NullablePutRoomTypesRequest {
	return &NullablePutRoomTypesRequest{value: val, isSet: true}
}

func (v NullablePutRoomTypesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutRoomTypesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


