/*
OPERA Cloud Activity Management API

APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actcfg

import (
	"encoding/json"
)

// checks if the PostActivityTypesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostActivityTypesRequest{}

// PostActivityTypesRequest struct for PostActivityTypesRequest
type PostActivityTypesRequest struct {
	// Detailed information of activity type.
	ActivityConfigTypes []ActivityConfigTypeDetailType `json:"activityConfigTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostActivityTypesRequest instantiates a new PostActivityTypesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostActivityTypesRequest() *PostActivityTypesRequest {
	this := PostActivityTypesRequest{}
	return &this
}

// NewPostActivityTypesRequestWithDefaults instantiates a new PostActivityTypesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostActivityTypesRequestWithDefaults() *PostActivityTypesRequest {
	this := PostActivityTypesRequest{}
	return &this
}

// GetActivityConfigTypes returns the ActivityConfigTypes field value if set, zero value otherwise.
func (o *PostActivityTypesRequest) GetActivityConfigTypes() []ActivityConfigTypeDetailType {
	if o == nil || IsNil(o.ActivityConfigTypes) {
		var ret []ActivityConfigTypeDetailType
		return ret
	}
	return o.ActivityConfigTypes
}

// GetActivityConfigTypesOk returns a tuple with the ActivityConfigTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostActivityTypesRequest) GetActivityConfigTypesOk() ([]ActivityConfigTypeDetailType, bool) {
	if o == nil || IsNil(o.ActivityConfigTypes) {
		return nil, false
	}
	return o.ActivityConfigTypes, true
}

// HasActivityConfigTypes returns a boolean if a field has been set.
func (o *PostActivityTypesRequest) HasActivityConfigTypes() bool {
	if o != nil && !IsNil(o.ActivityConfigTypes) {
		return true
	}

	return false
}

// SetActivityConfigTypes gets a reference to the given []ActivityConfigTypeDetailType and assigns it to the ActivityConfigTypes field.
func (o *PostActivityTypesRequest) SetActivityConfigTypes(v []ActivityConfigTypeDetailType) {
	o.ActivityConfigTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostActivityTypesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostActivityTypesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostActivityTypesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostActivityTypesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostActivityTypesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostActivityTypesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostActivityTypesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostActivityTypesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostActivityTypesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostActivityTypesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityConfigTypes) {
		toSerialize["activityConfigTypes"] = o.ActivityConfigTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostActivityTypesRequest struct {
	value *PostActivityTypesRequest
	isSet bool
}

func (v NullablePostActivityTypesRequest) Get() *PostActivityTypesRequest {
	return v.value
}

func (v *NullablePostActivityTypesRequest) Set(val *PostActivityTypesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostActivityTypesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostActivityTypesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostActivityTypesRequest(val *PostActivityTypesRequest) *NullablePostActivityTypesRequest {
	return &NullablePostActivityTypesRequest{value: val, isSet: true}
}

func (v NullablePostActivityTypesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostActivityTypesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


