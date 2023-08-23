/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PostAttractionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAttractionsRequest{}

// PostAttractionsRequest struct for PostAttractionsRequest
type PostAttractionsRequest struct {
	// Collection of hotel level alert codes with attached alert types.
	Attractions []AttractionCodeType `json:"attractions,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostAttractionsRequest instantiates a new PostAttractionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAttractionsRequest() *PostAttractionsRequest {
	this := PostAttractionsRequest{}
	return &this
}

// NewPostAttractionsRequestWithDefaults instantiates a new PostAttractionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAttractionsRequestWithDefaults() *PostAttractionsRequest {
	this := PostAttractionsRequest{}
	return &this
}

// GetAttractions returns the Attractions field value if set, zero value otherwise.
func (o *PostAttractionsRequest) GetAttractions() []AttractionCodeType {
	if o == nil || IsNil(o.Attractions) {
		var ret []AttractionCodeType
		return ret
	}
	return o.Attractions
}

// GetAttractionsOk returns a tuple with the Attractions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAttractionsRequest) GetAttractionsOk() ([]AttractionCodeType, bool) {
	if o == nil || IsNil(o.Attractions) {
		return nil, false
	}
	return o.Attractions, true
}

// HasAttractions returns a boolean if a field has been set.
func (o *PostAttractionsRequest) HasAttractions() bool {
	if o != nil && !IsNil(o.Attractions) {
		return true
	}

	return false
}

// SetAttractions gets a reference to the given []AttractionCodeType and assigns it to the Attractions field.
func (o *PostAttractionsRequest) SetAttractions(v []AttractionCodeType) {
	o.Attractions = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostAttractionsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAttractionsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostAttractionsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostAttractionsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostAttractionsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAttractionsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostAttractionsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostAttractionsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostAttractionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAttractionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attractions) {
		toSerialize["attractions"] = o.Attractions
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostAttractionsRequest struct {
	value *PostAttractionsRequest
	isSet bool
}

func (v NullablePostAttractionsRequest) Get() *PostAttractionsRequest {
	return v.value
}

func (v *NullablePostAttractionsRequest) Set(val *PostAttractionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAttractionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAttractionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAttractionsRequest(val *PostAttractionsRequest) *NullablePostAttractionsRequest {
	return &NullablePostAttractionsRequest{value: val, isSet: true}
}

func (v NullablePostAttractionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAttractionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


