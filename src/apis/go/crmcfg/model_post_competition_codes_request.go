/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmcfg

import (
	"encoding/json"
)

// checks if the PostCompetitionCodesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCompetitionCodesRequest{}

// PostCompetitionCodesRequest struct for PostCompetitionCodesRequest
type PostCompetitionCodesRequest struct {
	// List of Competition Codes.
	CompetitionCodes []CompetitionCodeType `json:"competitionCodes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostCompetitionCodesRequest instantiates a new PostCompetitionCodesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCompetitionCodesRequest() *PostCompetitionCodesRequest {
	this := PostCompetitionCodesRequest{}
	return &this
}

// NewPostCompetitionCodesRequestWithDefaults instantiates a new PostCompetitionCodesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCompetitionCodesRequestWithDefaults() *PostCompetitionCodesRequest {
	this := PostCompetitionCodesRequest{}
	return &this
}

// GetCompetitionCodes returns the CompetitionCodes field value if set, zero value otherwise.
func (o *PostCompetitionCodesRequest) GetCompetitionCodes() []CompetitionCodeType {
	if o == nil || IsNil(o.CompetitionCodes) {
		var ret []CompetitionCodeType
		return ret
	}
	return o.CompetitionCodes
}

// GetCompetitionCodesOk returns a tuple with the CompetitionCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompetitionCodesRequest) GetCompetitionCodesOk() ([]CompetitionCodeType, bool) {
	if o == nil || IsNil(o.CompetitionCodes) {
		return nil, false
	}
	return o.CompetitionCodes, true
}

// HasCompetitionCodes returns a boolean if a field has been set.
func (o *PostCompetitionCodesRequest) HasCompetitionCodes() bool {
	if o != nil && !IsNil(o.CompetitionCodes) {
		return true
	}

	return false
}

// SetCompetitionCodes gets a reference to the given []CompetitionCodeType and assigns it to the CompetitionCodes field.
func (o *PostCompetitionCodesRequest) SetCompetitionCodes(v []CompetitionCodeType) {
	o.CompetitionCodes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostCompetitionCodesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompetitionCodesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostCompetitionCodesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostCompetitionCodesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostCompetitionCodesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompetitionCodesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostCompetitionCodesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostCompetitionCodesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostCompetitionCodesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCompetitionCodesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompetitionCodes) {
		toSerialize["competitionCodes"] = o.CompetitionCodes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostCompetitionCodesRequest struct {
	value *PostCompetitionCodesRequest
	isSet bool
}

func (v NullablePostCompetitionCodesRequest) Get() *PostCompetitionCodesRequest {
	return v.value
}

func (v *NullablePostCompetitionCodesRequest) Set(val *PostCompetitionCodesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCompetitionCodesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCompetitionCodesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCompetitionCodesRequest(val *PostCompetitionCodesRequest) *NullablePostCompetitionCodesRequest {
	return &NullablePostCompetitionCodesRequest{value: val, isSet: true}
}

func (v NullablePostCompetitionCodesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCompetitionCodesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


