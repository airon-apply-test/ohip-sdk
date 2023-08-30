/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
)

// checks if the PutChannelAccountContractsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutChannelAccountContractsRequest{}

// PutChannelAccountContractsRequest struct for PutChannelAccountContractsRequest
type PutChannelAccountContractsRequest struct {
	ChannelAccountContracts *ChannelAccountContractsType `json:"channelAccountContracts,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutChannelAccountContractsRequest instantiates a new PutChannelAccountContractsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutChannelAccountContractsRequest() *PutChannelAccountContractsRequest {
	this := PutChannelAccountContractsRequest{}
	return &this
}

// NewPutChannelAccountContractsRequestWithDefaults instantiates a new PutChannelAccountContractsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutChannelAccountContractsRequestWithDefaults() *PutChannelAccountContractsRequest {
	this := PutChannelAccountContractsRequest{}
	return &this
}

// GetChannelAccountContracts returns the ChannelAccountContracts field value if set, zero value otherwise.
func (o *PutChannelAccountContractsRequest) GetChannelAccountContracts() ChannelAccountContractsType {
	if o == nil || IsNil(o.ChannelAccountContracts) {
		var ret ChannelAccountContractsType
		return ret
	}
	return *o.ChannelAccountContracts
}

// GetChannelAccountContractsOk returns a tuple with the ChannelAccountContracts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelAccountContractsRequest) GetChannelAccountContractsOk() (*ChannelAccountContractsType, bool) {
	if o == nil || IsNil(o.ChannelAccountContracts) {
		return nil, false
	}
	return o.ChannelAccountContracts, true
}

// HasChannelAccountContracts returns a boolean if a field has been set.
func (o *PutChannelAccountContractsRequest) HasChannelAccountContracts() bool {
	if o != nil && !IsNil(o.ChannelAccountContracts) {
		return true
	}

	return false
}

// SetChannelAccountContracts gets a reference to the given ChannelAccountContractsType and assigns it to the ChannelAccountContracts field.
func (o *PutChannelAccountContractsRequest) SetChannelAccountContracts(v ChannelAccountContractsType) {
	o.ChannelAccountContracts = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutChannelAccountContractsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelAccountContractsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutChannelAccountContractsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutChannelAccountContractsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutChannelAccountContractsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelAccountContractsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutChannelAccountContractsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutChannelAccountContractsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutChannelAccountContractsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutChannelAccountContractsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelAccountContracts) {
		toSerialize["channelAccountContracts"] = o.ChannelAccountContracts
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutChannelAccountContractsRequest struct {
	value *PutChannelAccountContractsRequest
	isSet bool
}

func (v NullablePutChannelAccountContractsRequest) Get() *PutChannelAccountContractsRequest {
	return v.value
}

func (v *NullablePutChannelAccountContractsRequest) Set(val *PutChannelAccountContractsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutChannelAccountContractsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutChannelAccountContractsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutChannelAccountContractsRequest(val *PutChannelAccountContractsRequest) *NullablePutChannelAccountContractsRequest {
	return &NullablePutChannelAccountContractsRequest{value: val, isSet: true}
}

func (v NullablePutChannelAccountContractsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutChannelAccountContractsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


