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

// checks if the MembershipRateGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipRateGroups{}

// MembershipRateGroups Request object for modifying Membership Rate Groups.
type MembershipRateGroups struct {
	// Details for Membership Rate Group.
	MembershipRateGroups []MembershipRateGroupType `json:"membershipRateGroups,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewMembershipRateGroups instantiates a new MembershipRateGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipRateGroups() *MembershipRateGroups {
	this := MembershipRateGroups{}
	return &this
}

// NewMembershipRateGroupsWithDefaults instantiates a new MembershipRateGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipRateGroupsWithDefaults() *MembershipRateGroups {
	this := MembershipRateGroups{}
	return &this
}

// GetMembershipRateGroups returns the MembershipRateGroups field value if set, zero value otherwise.
func (o *MembershipRateGroups) GetMembershipRateGroups() []MembershipRateGroupType {
	if o == nil || IsNil(o.MembershipRateGroups) {
		var ret []MembershipRateGroupType
		return ret
	}
	return o.MembershipRateGroups
}

// GetMembershipRateGroupsOk returns a tuple with the MembershipRateGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipRateGroups) GetMembershipRateGroupsOk() ([]MembershipRateGroupType, bool) {
	if o == nil || IsNil(o.MembershipRateGroups) {
		return nil, false
	}
	return o.MembershipRateGroups, true
}

// HasMembershipRateGroups returns a boolean if a field has been set.
func (o *MembershipRateGroups) HasMembershipRateGroups() bool {
	if o != nil && !IsNil(o.MembershipRateGroups) {
		return true
	}

	return false
}

// SetMembershipRateGroups gets a reference to the given []MembershipRateGroupType and assigns it to the MembershipRateGroups field.
func (o *MembershipRateGroups) SetMembershipRateGroups(v []MembershipRateGroupType) {
	o.MembershipRateGroups = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MembershipRateGroups) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipRateGroups) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MembershipRateGroups) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *MembershipRateGroups) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MembershipRateGroups) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipRateGroups) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MembershipRateGroups) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *MembershipRateGroups) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o MembershipRateGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipRateGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MembershipRateGroups) {
		toSerialize["membershipRateGroups"] = o.MembershipRateGroups
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableMembershipRateGroups struct {
	value *MembershipRateGroups
	isSet bool
}

func (v NullableMembershipRateGroups) Get() *MembershipRateGroups {
	return v.value
}

func (v *NullableMembershipRateGroups) Set(val *MembershipRateGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipRateGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipRateGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipRateGroups(val *MembershipRateGroups) *NullableMembershipRateGroups {
	return &NullableMembershipRateGroups{value: val, isSet: true}
}

func (v NullableMembershipRateGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipRateGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


