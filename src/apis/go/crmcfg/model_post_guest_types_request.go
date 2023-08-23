/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PostGuestTypesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostGuestTypesRequest{}

// PostGuestTypesRequest struct for PostGuestTypesRequest
type PostGuestTypesRequest struct {
	// List of Guest Types.
	GuestTypes []GuestTypeType `json:"guestTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostGuestTypesRequest instantiates a new PostGuestTypesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostGuestTypesRequest() *PostGuestTypesRequest {
	this := PostGuestTypesRequest{}
	return &this
}

// NewPostGuestTypesRequestWithDefaults instantiates a new PostGuestTypesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostGuestTypesRequestWithDefaults() *PostGuestTypesRequest {
	this := PostGuestTypesRequest{}
	return &this
}

// GetGuestTypes returns the GuestTypes field value if set, zero value otherwise.
func (o *PostGuestTypesRequest) GetGuestTypes() []GuestTypeType {
	if o == nil || IsNil(o.GuestTypes) {
		var ret []GuestTypeType
		return ret
	}
	return o.GuestTypes
}

// GetGuestTypesOk returns a tuple with the GuestTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGuestTypesRequest) GetGuestTypesOk() ([]GuestTypeType, bool) {
	if o == nil || IsNil(o.GuestTypes) {
		return nil, false
	}
	return o.GuestTypes, true
}

// HasGuestTypes returns a boolean if a field has been set.
func (o *PostGuestTypesRequest) HasGuestTypes() bool {
	if o != nil && !IsNil(o.GuestTypes) {
		return true
	}

	return false
}

// SetGuestTypes gets a reference to the given []GuestTypeType and assigns it to the GuestTypes field.
func (o *PostGuestTypesRequest) SetGuestTypes(v []GuestTypeType) {
	o.GuestTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostGuestTypesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGuestTypesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostGuestTypesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostGuestTypesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostGuestTypesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGuestTypesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostGuestTypesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostGuestTypesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostGuestTypesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostGuestTypesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuestTypes) {
		toSerialize["guestTypes"] = o.GuestTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostGuestTypesRequest struct {
	value *PostGuestTypesRequest
	isSet bool
}

func (v NullablePostGuestTypesRequest) Get() *PostGuestTypesRequest {
	return v.value
}

func (v *NullablePostGuestTypesRequest) Set(val *PostGuestTypesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostGuestTypesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostGuestTypesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostGuestTypesRequest(val *PostGuestTypesRequest) *NullablePostGuestTypesRequest {
	return &NullablePostGuestTypesRequest{value: val, isSet: true}
}

func (v NullablePostGuestTypesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostGuestTypesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


