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

// checks if the PostChannelMarketingTextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostChannelMarketingTextRequest{}

// PostChannelMarketingTextRequest struct for PostChannelMarketingTextRequest
type PostChannelMarketingTextRequest struct {
	// List of channel marketing texts to be created.
	CreateChannelMarketingTexts []MarketingTextType `json:"createChannelMarketingTexts,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostChannelMarketingTextRequest instantiates a new PostChannelMarketingTextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostChannelMarketingTextRequest() *PostChannelMarketingTextRequest {
	this := PostChannelMarketingTextRequest{}
	return &this
}

// NewPostChannelMarketingTextRequestWithDefaults instantiates a new PostChannelMarketingTextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostChannelMarketingTextRequestWithDefaults() *PostChannelMarketingTextRequest {
	this := PostChannelMarketingTextRequest{}
	return &this
}

// GetCreateChannelMarketingTexts returns the CreateChannelMarketingTexts field value if set, zero value otherwise.
func (o *PostChannelMarketingTextRequest) GetCreateChannelMarketingTexts() []MarketingTextType {
	if o == nil || IsNil(o.CreateChannelMarketingTexts) {
		var ret []MarketingTextType
		return ret
	}
	return o.CreateChannelMarketingTexts
}

// GetCreateChannelMarketingTextsOk returns a tuple with the CreateChannelMarketingTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostChannelMarketingTextRequest) GetCreateChannelMarketingTextsOk() ([]MarketingTextType, bool) {
	if o == nil || IsNil(o.CreateChannelMarketingTexts) {
		return nil, false
	}
	return o.CreateChannelMarketingTexts, true
}

// HasCreateChannelMarketingTexts returns a boolean if a field has been set.
func (o *PostChannelMarketingTextRequest) HasCreateChannelMarketingTexts() bool {
	if o != nil && !IsNil(o.CreateChannelMarketingTexts) {
		return true
	}

	return false
}

// SetCreateChannelMarketingTexts gets a reference to the given []MarketingTextType and assigns it to the CreateChannelMarketingTexts field.
func (o *PostChannelMarketingTextRequest) SetCreateChannelMarketingTexts(v []MarketingTextType) {
	o.CreateChannelMarketingTexts = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostChannelMarketingTextRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostChannelMarketingTextRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostChannelMarketingTextRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostChannelMarketingTextRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostChannelMarketingTextRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostChannelMarketingTextRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostChannelMarketingTextRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostChannelMarketingTextRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostChannelMarketingTextRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostChannelMarketingTextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateChannelMarketingTexts) {
		toSerialize["createChannelMarketingTexts"] = o.CreateChannelMarketingTexts
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostChannelMarketingTextRequest struct {
	value *PostChannelMarketingTextRequest
	isSet bool
}

func (v NullablePostChannelMarketingTextRequest) Get() *PostChannelMarketingTextRequest {
	return v.value
}

func (v *NullablePostChannelMarketingTextRequest) Set(val *PostChannelMarketingTextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostChannelMarketingTextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostChannelMarketingTextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostChannelMarketingTextRequest(val *PostChannelMarketingTextRequest) *NullablePostChannelMarketingTextRequest {
	return &NullablePostChannelMarketingTextRequest{value: val, isSet: true}
}

func (v NullablePostChannelMarketingTextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostChannelMarketingTextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


