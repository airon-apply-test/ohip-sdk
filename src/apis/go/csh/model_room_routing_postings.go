/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RoomRoutingPostings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomRoutingPostings{}

// RoomRoutingPostings struct for RoomRoutingPostings
type RoomRoutingPostings struct {
	PostingsForRoomRouting *PostingsInfoType `json:"postingsForRoomRouting,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewRoomRoutingPostings instantiates a new RoomRoutingPostings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomRoutingPostings() *RoomRoutingPostings {
	this := RoomRoutingPostings{}
	return &this
}

// NewRoomRoutingPostingsWithDefaults instantiates a new RoomRoutingPostings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomRoutingPostingsWithDefaults() *RoomRoutingPostings {
	this := RoomRoutingPostings{}
	return &this
}

// GetPostingsForRoomRouting returns the PostingsForRoomRouting field value if set, zero value otherwise.
func (o *RoomRoutingPostings) GetPostingsForRoomRouting() PostingsInfoType {
	if o == nil || IsNil(o.PostingsForRoomRouting) {
		var ret PostingsInfoType
		return ret
	}
	return *o.PostingsForRoomRouting
}

// GetPostingsForRoomRoutingOk returns a tuple with the PostingsForRoomRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomRoutingPostings) GetPostingsForRoomRoutingOk() (*PostingsInfoType, bool) {
	if o == nil || IsNil(o.PostingsForRoomRouting) {
		return nil, false
	}
	return o.PostingsForRoomRouting, true
}

// HasPostingsForRoomRouting returns a boolean if a field has been set.
func (o *RoomRoutingPostings) HasPostingsForRoomRouting() bool {
	if o != nil && !IsNil(o.PostingsForRoomRouting) {
		return true
	}

	return false
}

// SetPostingsForRoomRouting gets a reference to the given PostingsInfoType and assigns it to the PostingsForRoomRouting field.
func (o *RoomRoutingPostings) SetPostingsForRoomRouting(v PostingsInfoType) {
	o.PostingsForRoomRouting = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *RoomRoutingPostings) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomRoutingPostings) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RoomRoutingPostings) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *RoomRoutingPostings) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o RoomRoutingPostings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomRoutingPostings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostingsForRoomRouting) {
		toSerialize["postingsForRoomRouting"] = o.PostingsForRoomRouting
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableRoomRoutingPostings struct {
	value *RoomRoutingPostings
	isSet bool
}

func (v NullableRoomRoutingPostings) Get() *RoomRoutingPostings {
	return v.value
}

func (v *NullableRoomRoutingPostings) Set(val *RoomRoutingPostings) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomRoutingPostings) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomRoutingPostings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomRoutingPostings(val *RoomRoutingPostings) *NullableRoomRoutingPostings {
	return &NullableRoomRoutingPostings{value: val, isSet: true}
}

func (v NullableRoomRoutingPostings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomRoutingPostings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


