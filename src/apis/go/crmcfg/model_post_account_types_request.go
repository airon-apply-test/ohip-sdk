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

// checks if the PostAccountTypesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAccountTypesRequest{}

// PostAccountTypesRequest struct for PostAccountTypesRequest
type PostAccountTypesRequest struct {
	// List of Account Types.
	AccountTypes []AccountTypeType `json:"accountTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostAccountTypesRequest instantiates a new PostAccountTypesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAccountTypesRequest() *PostAccountTypesRequest {
	this := PostAccountTypesRequest{}
	return &this
}

// NewPostAccountTypesRequestWithDefaults instantiates a new PostAccountTypesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAccountTypesRequestWithDefaults() *PostAccountTypesRequest {
	this := PostAccountTypesRequest{}
	return &this
}

// GetAccountTypes returns the AccountTypes field value if set, zero value otherwise.
func (o *PostAccountTypesRequest) GetAccountTypes() []AccountTypeType {
	if o == nil || IsNil(o.AccountTypes) {
		var ret []AccountTypeType
		return ret
	}
	return o.AccountTypes
}

// GetAccountTypesOk returns a tuple with the AccountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAccountTypesRequest) GetAccountTypesOk() ([]AccountTypeType, bool) {
	if o == nil || IsNil(o.AccountTypes) {
		return nil, false
	}
	return o.AccountTypes, true
}

// HasAccountTypes returns a boolean if a field has been set.
func (o *PostAccountTypesRequest) HasAccountTypes() bool {
	if o != nil && !IsNil(o.AccountTypes) {
		return true
	}

	return false
}

// SetAccountTypes gets a reference to the given []AccountTypeType and assigns it to the AccountTypes field.
func (o *PostAccountTypesRequest) SetAccountTypes(v []AccountTypeType) {
	o.AccountTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostAccountTypesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAccountTypesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostAccountTypesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostAccountTypesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostAccountTypesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAccountTypesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostAccountTypesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostAccountTypesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostAccountTypesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAccountTypesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountTypes) {
		toSerialize["accountTypes"] = o.AccountTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostAccountTypesRequest struct {
	value *PostAccountTypesRequest
	isSet bool
}

func (v NullablePostAccountTypesRequest) Get() *PostAccountTypesRequest {
	return v.value
}

func (v *NullablePostAccountTypesRequest) Set(val *PostAccountTypesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAccountTypesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAccountTypesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAccountTypesRequest(val *PostAccountTypesRequest) *NullablePostAccountTypesRequest {
	return &NullablePostAccountTypesRequest{value: val, isSet: true}
}

func (v NullablePostAccountTypesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAccountTypesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


