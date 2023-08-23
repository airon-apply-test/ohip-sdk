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

// checks if the ChannelAccountContractCopyType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelAccountContractCopyType{}

// ChannelAccountContractCopyType Channel account contract type to hold account and contract details to copy.
type ChannelAccountContractCopyType struct {
	TargetProfileId *UniqueIDType `json:"targetProfileId,omitempty"`
	// Target Account code to which contract to be copied. This is utilised to find account when TargetProfileID is not provided.
	TargetAccountCode *string `json:"targetAccountCode,omitempty"`
	// Contract details to copy.
	ChannelAccountContractCopyDetails []ChannelAccountContractCopyDetailsType `json:"channelAccountContractCopyDetails,omitempty"`
}

// NewChannelAccountContractCopyType instantiates a new ChannelAccountContractCopyType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelAccountContractCopyType() *ChannelAccountContractCopyType {
	this := ChannelAccountContractCopyType{}
	return &this
}

// NewChannelAccountContractCopyTypeWithDefaults instantiates a new ChannelAccountContractCopyType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelAccountContractCopyTypeWithDefaults() *ChannelAccountContractCopyType {
	this := ChannelAccountContractCopyType{}
	return &this
}

// GetTargetProfileId returns the TargetProfileId field value if set, zero value otherwise.
func (o *ChannelAccountContractCopyType) GetTargetProfileId() UniqueIDType {
	if o == nil || IsNil(o.TargetProfileId) {
		var ret UniqueIDType
		return ret
	}
	return *o.TargetProfileId
}

// GetTargetProfileIdOk returns a tuple with the TargetProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContractCopyType) GetTargetProfileIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.TargetProfileId) {
		return nil, false
	}
	return o.TargetProfileId, true
}

// HasTargetProfileId returns a boolean if a field has been set.
func (o *ChannelAccountContractCopyType) HasTargetProfileId() bool {
	if o != nil && !IsNil(o.TargetProfileId) {
		return true
	}

	return false
}

// SetTargetProfileId gets a reference to the given UniqueIDType and assigns it to the TargetProfileId field.
func (o *ChannelAccountContractCopyType) SetTargetProfileId(v UniqueIDType) {
	o.TargetProfileId = &v
}

// GetTargetAccountCode returns the TargetAccountCode field value if set, zero value otherwise.
func (o *ChannelAccountContractCopyType) GetTargetAccountCode() string {
	if o == nil || IsNil(o.TargetAccountCode) {
		var ret string
		return ret
	}
	return *o.TargetAccountCode
}

// GetTargetAccountCodeOk returns a tuple with the TargetAccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContractCopyType) GetTargetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAccountCode) {
		return nil, false
	}
	return o.TargetAccountCode, true
}

// HasTargetAccountCode returns a boolean if a field has been set.
func (o *ChannelAccountContractCopyType) HasTargetAccountCode() bool {
	if o != nil && !IsNil(o.TargetAccountCode) {
		return true
	}

	return false
}

// SetTargetAccountCode gets a reference to the given string and assigns it to the TargetAccountCode field.
func (o *ChannelAccountContractCopyType) SetTargetAccountCode(v string) {
	o.TargetAccountCode = &v
}

// GetChannelAccountContractCopyDetails returns the ChannelAccountContractCopyDetails field value if set, zero value otherwise.
func (o *ChannelAccountContractCopyType) GetChannelAccountContractCopyDetails() []ChannelAccountContractCopyDetailsType {
	if o == nil || IsNil(o.ChannelAccountContractCopyDetails) {
		var ret []ChannelAccountContractCopyDetailsType
		return ret
	}
	return o.ChannelAccountContractCopyDetails
}

// GetChannelAccountContractCopyDetailsOk returns a tuple with the ChannelAccountContractCopyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContractCopyType) GetChannelAccountContractCopyDetailsOk() ([]ChannelAccountContractCopyDetailsType, bool) {
	if o == nil || IsNil(o.ChannelAccountContractCopyDetails) {
		return nil, false
	}
	return o.ChannelAccountContractCopyDetails, true
}

// HasChannelAccountContractCopyDetails returns a boolean if a field has been set.
func (o *ChannelAccountContractCopyType) HasChannelAccountContractCopyDetails() bool {
	if o != nil && !IsNil(o.ChannelAccountContractCopyDetails) {
		return true
	}

	return false
}

// SetChannelAccountContractCopyDetails gets a reference to the given []ChannelAccountContractCopyDetailsType and assigns it to the ChannelAccountContractCopyDetails field.
func (o *ChannelAccountContractCopyType) SetChannelAccountContractCopyDetails(v []ChannelAccountContractCopyDetailsType) {
	o.ChannelAccountContractCopyDetails = v
}

func (o ChannelAccountContractCopyType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelAccountContractCopyType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetProfileId) {
		toSerialize["targetProfileId"] = o.TargetProfileId
	}
	if !IsNil(o.TargetAccountCode) {
		toSerialize["targetAccountCode"] = o.TargetAccountCode
	}
	if !IsNil(o.ChannelAccountContractCopyDetails) {
		toSerialize["channelAccountContractCopyDetails"] = o.ChannelAccountContractCopyDetails
	}
	return toSerialize, nil
}

type NullableChannelAccountContractCopyType struct {
	value *ChannelAccountContractCopyType
	isSet bool
}

func (v NullableChannelAccountContractCopyType) Get() *ChannelAccountContractCopyType {
	return v.value
}

func (v *NullableChannelAccountContractCopyType) Set(val *ChannelAccountContractCopyType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelAccountContractCopyType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelAccountContractCopyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelAccountContractCopyType(val *ChannelAccountContractCopyType) *NullableChannelAccountContractCopyType {
	return &NullableChannelAccountContractCopyType{value: val, isSet: true}
}

func (v NullableChannelAccountContractCopyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelAccountContractCopyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


