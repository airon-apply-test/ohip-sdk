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

// checks if the PostMembershipPropertyGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostMembershipPropertyGroupsRequest{}

// PostMembershipPropertyGroupsRequest struct for PostMembershipPropertyGroupsRequest
type PostMembershipPropertyGroupsRequest struct {
	// Details for Membership Property Group along with associated property codes.
	MembershipPropertyGroups []MembershipPropertyGroupType `json:"membershipPropertyGroups,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostMembershipPropertyGroupsRequest instantiates a new PostMembershipPropertyGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostMembershipPropertyGroupsRequest() *PostMembershipPropertyGroupsRequest {
	this := PostMembershipPropertyGroupsRequest{}
	return &this
}

// NewPostMembershipPropertyGroupsRequestWithDefaults instantiates a new PostMembershipPropertyGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostMembershipPropertyGroupsRequestWithDefaults() *PostMembershipPropertyGroupsRequest {
	this := PostMembershipPropertyGroupsRequest{}
	return &this
}

// GetMembershipPropertyGroups returns the MembershipPropertyGroups field value if set, zero value otherwise.
func (o *PostMembershipPropertyGroupsRequest) GetMembershipPropertyGroups() []MembershipPropertyGroupType {
	if o == nil || IsNil(o.MembershipPropertyGroups) {
		var ret []MembershipPropertyGroupType
		return ret
	}
	return o.MembershipPropertyGroups
}

// GetMembershipPropertyGroupsOk returns a tuple with the MembershipPropertyGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMembershipPropertyGroupsRequest) GetMembershipPropertyGroupsOk() ([]MembershipPropertyGroupType, bool) {
	if o == nil || IsNil(o.MembershipPropertyGroups) {
		return nil, false
	}
	return o.MembershipPropertyGroups, true
}

// HasMembershipPropertyGroups returns a boolean if a field has been set.
func (o *PostMembershipPropertyGroupsRequest) HasMembershipPropertyGroups() bool {
	if o != nil && !IsNil(o.MembershipPropertyGroups) {
		return true
	}

	return false
}

// SetMembershipPropertyGroups gets a reference to the given []MembershipPropertyGroupType and assigns it to the MembershipPropertyGroups field.
func (o *PostMembershipPropertyGroupsRequest) SetMembershipPropertyGroups(v []MembershipPropertyGroupType) {
	o.MembershipPropertyGroups = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostMembershipPropertyGroupsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMembershipPropertyGroupsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostMembershipPropertyGroupsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostMembershipPropertyGroupsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostMembershipPropertyGroupsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMembershipPropertyGroupsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostMembershipPropertyGroupsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostMembershipPropertyGroupsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostMembershipPropertyGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostMembershipPropertyGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MembershipPropertyGroups) {
		toSerialize["membershipPropertyGroups"] = o.MembershipPropertyGroups
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostMembershipPropertyGroupsRequest struct {
	value *PostMembershipPropertyGroupsRequest
	isSet bool
}

func (v NullablePostMembershipPropertyGroupsRequest) Get() *PostMembershipPropertyGroupsRequest {
	return v.value
}

func (v *NullablePostMembershipPropertyGroupsRequest) Set(val *PostMembershipPropertyGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostMembershipPropertyGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostMembershipPropertyGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostMembershipPropertyGroupsRequest(val *PostMembershipPropertyGroupsRequest) *NullablePostMembershipPropertyGroupsRequest {
	return &NullablePostMembershipPropertyGroupsRequest{value: val, isSet: true}
}

func (v NullablePostMembershipPropertyGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostMembershipPropertyGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


