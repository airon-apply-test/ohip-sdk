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

// checks if the PostChannelAccountsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostChannelAccountsRequest{}

// PostChannelAccountsRequest struct for PostChannelAccountsRequest
type PostChannelAccountsRequest struct {
	// Channel account information object to hold details of channel account.
	ChannelAccounts []ChannelAccountInformationType `json:"channelAccounts,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostChannelAccountsRequest instantiates a new PostChannelAccountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostChannelAccountsRequest() *PostChannelAccountsRequest {
	this := PostChannelAccountsRequest{}
	return &this
}

// NewPostChannelAccountsRequestWithDefaults instantiates a new PostChannelAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostChannelAccountsRequestWithDefaults() *PostChannelAccountsRequest {
	this := PostChannelAccountsRequest{}
	return &this
}

// GetChannelAccounts returns the ChannelAccounts field value if set, zero value otherwise.
func (o *PostChannelAccountsRequest) GetChannelAccounts() []ChannelAccountInformationType {
	if o == nil || IsNil(o.ChannelAccounts) {
		var ret []ChannelAccountInformationType
		return ret
	}
	return o.ChannelAccounts
}

// GetChannelAccountsOk returns a tuple with the ChannelAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostChannelAccountsRequest) GetChannelAccountsOk() ([]ChannelAccountInformationType, bool) {
	if o == nil || IsNil(o.ChannelAccounts) {
		return nil, false
	}
	return o.ChannelAccounts, true
}

// HasChannelAccounts returns a boolean if a field has been set.
func (o *PostChannelAccountsRequest) HasChannelAccounts() bool {
	if o != nil && !IsNil(o.ChannelAccounts) {
		return true
	}

	return false
}

// SetChannelAccounts gets a reference to the given []ChannelAccountInformationType and assigns it to the ChannelAccounts field.
func (o *PostChannelAccountsRequest) SetChannelAccounts(v []ChannelAccountInformationType) {
	o.ChannelAccounts = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostChannelAccountsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostChannelAccountsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostChannelAccountsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostChannelAccountsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostChannelAccountsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostChannelAccountsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostChannelAccountsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostChannelAccountsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostChannelAccountsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostChannelAccountsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelAccounts) {
		toSerialize["channelAccounts"] = o.ChannelAccounts
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostChannelAccountsRequest struct {
	value *PostChannelAccountsRequest
	isSet bool
}

func (v NullablePostChannelAccountsRequest) Get() *PostChannelAccountsRequest {
	return v.value
}

func (v *NullablePostChannelAccountsRequest) Set(val *PostChannelAccountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostChannelAccountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostChannelAccountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostChannelAccountsRequest(val *PostChannelAccountsRequest) *NullablePostChannelAccountsRequest {
	return &NullablePostChannelAccountsRequest{value: val, isSet: true}
}

func (v NullablePostChannelAccountsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostChannelAccountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


