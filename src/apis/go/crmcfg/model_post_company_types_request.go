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

// checks if the PostCompanyTypesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCompanyTypesRequest{}

// PostCompanyTypesRequest struct for PostCompanyTypesRequest
type PostCompanyTypesRequest struct {
	// List of Company Types.
	CompanyTypes []CompanyTypeType `json:"companyTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostCompanyTypesRequest instantiates a new PostCompanyTypesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCompanyTypesRequest() *PostCompanyTypesRequest {
	this := PostCompanyTypesRequest{}
	return &this
}

// NewPostCompanyTypesRequestWithDefaults instantiates a new PostCompanyTypesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCompanyTypesRequestWithDefaults() *PostCompanyTypesRequest {
	this := PostCompanyTypesRequest{}
	return &this
}

// GetCompanyTypes returns the CompanyTypes field value if set, zero value otherwise.
func (o *PostCompanyTypesRequest) GetCompanyTypes() []CompanyTypeType {
	if o == nil || IsNil(o.CompanyTypes) {
		var ret []CompanyTypeType
		return ret
	}
	return o.CompanyTypes
}

// GetCompanyTypesOk returns a tuple with the CompanyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompanyTypesRequest) GetCompanyTypesOk() ([]CompanyTypeType, bool) {
	if o == nil || IsNil(o.CompanyTypes) {
		return nil, false
	}
	return o.CompanyTypes, true
}

// HasCompanyTypes returns a boolean if a field has been set.
func (o *PostCompanyTypesRequest) HasCompanyTypes() bool {
	if o != nil && !IsNil(o.CompanyTypes) {
		return true
	}

	return false
}

// SetCompanyTypes gets a reference to the given []CompanyTypeType and assigns it to the CompanyTypes field.
func (o *PostCompanyTypesRequest) SetCompanyTypes(v []CompanyTypeType) {
	o.CompanyTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostCompanyTypesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompanyTypesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostCompanyTypesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostCompanyTypesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostCompanyTypesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompanyTypesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostCompanyTypesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostCompanyTypesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostCompanyTypesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCompanyTypesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyTypes) {
		toSerialize["companyTypes"] = o.CompanyTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostCompanyTypesRequest struct {
	value *PostCompanyTypesRequest
	isSet bool
}

func (v NullablePostCompanyTypesRequest) Get() *PostCompanyTypesRequest {
	return v.value
}

func (v *NullablePostCompanyTypesRequest) Set(val *PostCompanyTypesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCompanyTypesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCompanyTypesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCompanyTypesRequest(val *PostCompanyTypesRequest) *NullablePostCompanyTypesRequest {
	return &NullablePostCompanyTypesRequest{value: val, isSet: true}
}

func (v NullablePostCompanyTypesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCompanyTypesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


