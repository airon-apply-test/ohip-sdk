/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the ProcessRoomRoutingRefreshRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessRoomRoutingRefreshRequest{}

// ProcessRoomRoutingRefreshRequest struct for ProcessRoomRoutingRefreshRequest
type ProcessRoomRoutingRefreshRequest struct {
	RoomRoutingRefreshCriteria *RoomRoutingRefreshCriteriaType `json:"roomRoutingRefreshCriteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewProcessRoomRoutingRefreshRequest instantiates a new ProcessRoomRoutingRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessRoomRoutingRefreshRequest() *ProcessRoomRoutingRefreshRequest {
	this := ProcessRoomRoutingRefreshRequest{}
	return &this
}

// NewProcessRoomRoutingRefreshRequestWithDefaults instantiates a new ProcessRoomRoutingRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessRoomRoutingRefreshRequestWithDefaults() *ProcessRoomRoutingRefreshRequest {
	this := ProcessRoomRoutingRefreshRequest{}
	return &this
}

// GetRoomRoutingRefreshCriteria returns the RoomRoutingRefreshCriteria field value if set, zero value otherwise.
func (o *ProcessRoomRoutingRefreshRequest) GetRoomRoutingRefreshCriteria() RoomRoutingRefreshCriteriaType {
	if o == nil || IsNil(o.RoomRoutingRefreshCriteria) {
		var ret RoomRoutingRefreshCriteriaType
		return ret
	}
	return *o.RoomRoutingRefreshCriteria
}

// GetRoomRoutingRefreshCriteriaOk returns a tuple with the RoomRoutingRefreshCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessRoomRoutingRefreshRequest) GetRoomRoutingRefreshCriteriaOk() (*RoomRoutingRefreshCriteriaType, bool) {
	if o == nil || IsNil(o.RoomRoutingRefreshCriteria) {
		return nil, false
	}
	return o.RoomRoutingRefreshCriteria, true
}

// HasRoomRoutingRefreshCriteria returns a boolean if a field has been set.
func (o *ProcessRoomRoutingRefreshRequest) HasRoomRoutingRefreshCriteria() bool {
	if o != nil && !IsNil(o.RoomRoutingRefreshCriteria) {
		return true
	}

	return false
}

// SetRoomRoutingRefreshCriteria gets a reference to the given RoomRoutingRefreshCriteriaType and assigns it to the RoomRoutingRefreshCriteria field.
func (o *ProcessRoomRoutingRefreshRequest) SetRoomRoutingRefreshCriteria(v RoomRoutingRefreshCriteriaType) {
	o.RoomRoutingRefreshCriteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProcessRoomRoutingRefreshRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessRoomRoutingRefreshRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProcessRoomRoutingRefreshRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ProcessRoomRoutingRefreshRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ProcessRoomRoutingRefreshRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessRoomRoutingRefreshRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ProcessRoomRoutingRefreshRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ProcessRoomRoutingRefreshRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ProcessRoomRoutingRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessRoomRoutingRefreshRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomRoutingRefreshCriteria) {
		toSerialize["roomRoutingRefreshCriteria"] = o.RoomRoutingRefreshCriteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableProcessRoomRoutingRefreshRequest struct {
	value *ProcessRoomRoutingRefreshRequest
	isSet bool
}

func (v NullableProcessRoomRoutingRefreshRequest) Get() *ProcessRoomRoutingRefreshRequest {
	return v.value
}

func (v *NullableProcessRoomRoutingRefreshRequest) Set(val *ProcessRoomRoutingRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessRoomRoutingRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessRoomRoutingRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessRoomRoutingRefreshRequest(val *ProcessRoomRoutingRefreshRequest) *NullableProcessRoomRoutingRefreshRequest {
	return &NullableProcessRoomRoutingRefreshRequest{value: val, isSet: true}
}

func (v NullableProcessRoomRoutingRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessRoomRoutingRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


