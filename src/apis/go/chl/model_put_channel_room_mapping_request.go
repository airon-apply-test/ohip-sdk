/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PutChannelRoomMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutChannelRoomMappingRequest{}

// PutChannelRoomMappingRequest struct for PutChannelRoomMappingRequest
type PutChannelRoomMappingRequest struct {
	ChannelRoomMapping *ChannelRoomMappingType `json:"channelRoomMapping,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutChannelRoomMappingRequest instantiates a new PutChannelRoomMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutChannelRoomMappingRequest() *PutChannelRoomMappingRequest {
	this := PutChannelRoomMappingRequest{}
	return &this
}

// NewPutChannelRoomMappingRequestWithDefaults instantiates a new PutChannelRoomMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutChannelRoomMappingRequestWithDefaults() *PutChannelRoomMappingRequest {
	this := PutChannelRoomMappingRequest{}
	return &this
}

// GetChannelRoomMapping returns the ChannelRoomMapping field value if set, zero value otherwise.
func (o *PutChannelRoomMappingRequest) GetChannelRoomMapping() ChannelRoomMappingType {
	if o == nil || IsNil(o.ChannelRoomMapping) {
		var ret ChannelRoomMappingType
		return ret
	}
	return *o.ChannelRoomMapping
}

// GetChannelRoomMappingOk returns a tuple with the ChannelRoomMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRoomMappingRequest) GetChannelRoomMappingOk() (*ChannelRoomMappingType, bool) {
	if o == nil || IsNil(o.ChannelRoomMapping) {
		return nil, false
	}
	return o.ChannelRoomMapping, true
}

// HasChannelRoomMapping returns a boolean if a field has been set.
func (o *PutChannelRoomMappingRequest) HasChannelRoomMapping() bool {
	if o != nil && !IsNil(o.ChannelRoomMapping) {
		return true
	}

	return false
}

// SetChannelRoomMapping gets a reference to the given ChannelRoomMappingType and assigns it to the ChannelRoomMapping field.
func (o *PutChannelRoomMappingRequest) SetChannelRoomMapping(v ChannelRoomMappingType) {
	o.ChannelRoomMapping = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutChannelRoomMappingRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRoomMappingRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutChannelRoomMappingRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutChannelRoomMappingRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutChannelRoomMappingRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRoomMappingRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutChannelRoomMappingRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutChannelRoomMappingRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutChannelRoomMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutChannelRoomMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelRoomMapping) {
		toSerialize["channelRoomMapping"] = o.ChannelRoomMapping
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutChannelRoomMappingRequest struct {
	value *PutChannelRoomMappingRequest
	isSet bool
}

func (v NullablePutChannelRoomMappingRequest) Get() *PutChannelRoomMappingRequest {
	return v.value
}

func (v *NullablePutChannelRoomMappingRequest) Set(val *PutChannelRoomMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutChannelRoomMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutChannelRoomMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutChannelRoomMappingRequest(val *PutChannelRoomMappingRequest) *NullablePutChannelRoomMappingRequest {
	return &NullablePutChannelRoomMappingRequest{value: val, isSet: true}
}

func (v NullablePutChannelRoomMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutChannelRoomMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


