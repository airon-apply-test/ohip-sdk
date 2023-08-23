/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PostLinkedActivitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLinkedActivitiesRequest{}

// PostLinkedActivitiesRequest struct for PostLinkedActivitiesRequest
type PostLinkedActivitiesRequest struct {
	LinkedActivityDetails *LinkedActivitiesType `json:"linkedActivityDetails,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostLinkedActivitiesRequest instantiates a new PostLinkedActivitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLinkedActivitiesRequest() *PostLinkedActivitiesRequest {
	this := PostLinkedActivitiesRequest{}
	return &this
}

// NewPostLinkedActivitiesRequestWithDefaults instantiates a new PostLinkedActivitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLinkedActivitiesRequestWithDefaults() *PostLinkedActivitiesRequest {
	this := PostLinkedActivitiesRequest{}
	return &this
}

// GetLinkedActivityDetails returns the LinkedActivityDetails field value if set, zero value otherwise.
func (o *PostLinkedActivitiesRequest) GetLinkedActivityDetails() LinkedActivitiesType {
	if o == nil || IsNil(o.LinkedActivityDetails) {
		var ret LinkedActivitiesType
		return ret
	}
	return *o.LinkedActivityDetails
}

// GetLinkedActivityDetailsOk returns a tuple with the LinkedActivityDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinkedActivitiesRequest) GetLinkedActivityDetailsOk() (*LinkedActivitiesType, bool) {
	if o == nil || IsNil(o.LinkedActivityDetails) {
		return nil, false
	}
	return o.LinkedActivityDetails, true
}

// HasLinkedActivityDetails returns a boolean if a field has been set.
func (o *PostLinkedActivitiesRequest) HasLinkedActivityDetails() bool {
	if o != nil && !IsNil(o.LinkedActivityDetails) {
		return true
	}

	return false
}

// SetLinkedActivityDetails gets a reference to the given LinkedActivitiesType and assigns it to the LinkedActivityDetails field.
func (o *PostLinkedActivitiesRequest) SetLinkedActivityDetails(v LinkedActivitiesType) {
	o.LinkedActivityDetails = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostLinkedActivitiesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinkedActivitiesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostLinkedActivitiesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostLinkedActivitiesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostLinkedActivitiesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinkedActivitiesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostLinkedActivitiesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostLinkedActivitiesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostLinkedActivitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLinkedActivitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkedActivityDetails) {
		toSerialize["linkedActivityDetails"] = o.LinkedActivityDetails
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostLinkedActivitiesRequest struct {
	value *PostLinkedActivitiesRequest
	isSet bool
}

func (v NullablePostLinkedActivitiesRequest) Get() *PostLinkedActivitiesRequest {
	return v.value
}

func (v *NullablePostLinkedActivitiesRequest) Set(val *PostLinkedActivitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLinkedActivitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLinkedActivitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLinkedActivitiesRequest(val *PostLinkedActivitiesRequest) *NullablePostLinkedActivitiesRequest {
	return &NullablePostLinkedActivitiesRequest{value: val, isSet: true}
}

func (v NullablePostLinkedActivitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLinkedActivitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


