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

// checks if the MembershipAwards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipAwards{}

// MembershipAwards Request object for modifying membership awards.
type MembershipAwards struct {
	// Membership Award details.
	MembershipAwards []MembershipAwardType `json:"membershipAwards,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewMembershipAwards instantiates a new MembershipAwards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipAwards() *MembershipAwards {
	this := MembershipAwards{}
	return &this
}

// NewMembershipAwardsWithDefaults instantiates a new MembershipAwards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipAwardsWithDefaults() *MembershipAwards {
	this := MembershipAwards{}
	return &this
}

// GetMembershipAwards returns the MembershipAwards field value if set, zero value otherwise.
func (o *MembershipAwards) GetMembershipAwards() []MembershipAwardType {
	if o == nil || IsNil(o.MembershipAwards) {
		var ret []MembershipAwardType
		return ret
	}
	return o.MembershipAwards
}

// GetMembershipAwardsOk returns a tuple with the MembershipAwards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwards) GetMembershipAwardsOk() ([]MembershipAwardType, bool) {
	if o == nil || IsNil(o.MembershipAwards) {
		return nil, false
	}
	return o.MembershipAwards, true
}

// HasMembershipAwards returns a boolean if a field has been set.
func (o *MembershipAwards) HasMembershipAwards() bool {
	if o != nil && !IsNil(o.MembershipAwards) {
		return true
	}

	return false
}

// SetMembershipAwards gets a reference to the given []MembershipAwardType and assigns it to the MembershipAwards field.
func (o *MembershipAwards) SetMembershipAwards(v []MembershipAwardType) {
	o.MembershipAwards = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MembershipAwards) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwards) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MembershipAwards) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *MembershipAwards) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MembershipAwards) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipAwards) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MembershipAwards) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *MembershipAwards) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o MembershipAwards) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipAwards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MembershipAwards) {
		toSerialize["membershipAwards"] = o.MembershipAwards
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableMembershipAwards struct {
	value *MembershipAwards
	isSet bool
}

func (v NullableMembershipAwards) Get() *MembershipAwards {
	return v.value
}

func (v *NullableMembershipAwards) Set(val *MembershipAwards) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipAwards) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipAwards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipAwards(val *MembershipAwards) *NullableMembershipAwards {
	return &NullableMembershipAwards{value: val, isSet: true}
}

func (v NullableMembershipAwards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipAwards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


