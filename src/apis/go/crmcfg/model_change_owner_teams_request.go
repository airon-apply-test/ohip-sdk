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

// checks if the ChangeOwnerTeamsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeOwnerTeamsRequest{}

// ChangeOwnerTeamsRequest struct for ChangeOwnerTeamsRequest
type ChangeOwnerTeamsRequest struct {
	// List of Owner Teams.
	OwnerTeams []OwnerTeamType `json:"ownerTeams,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChangeOwnerTeamsRequest instantiates a new ChangeOwnerTeamsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeOwnerTeamsRequest() *ChangeOwnerTeamsRequest {
	this := ChangeOwnerTeamsRequest{}
	return &this
}

// NewChangeOwnerTeamsRequestWithDefaults instantiates a new ChangeOwnerTeamsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeOwnerTeamsRequestWithDefaults() *ChangeOwnerTeamsRequest {
	this := ChangeOwnerTeamsRequest{}
	return &this
}

// GetOwnerTeams returns the OwnerTeams field value if set, zero value otherwise.
func (o *ChangeOwnerTeamsRequest) GetOwnerTeams() []OwnerTeamType {
	if o == nil || IsNil(o.OwnerTeams) {
		var ret []OwnerTeamType
		return ret
	}
	return o.OwnerTeams
}

// GetOwnerTeamsOk returns a tuple with the OwnerTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeOwnerTeamsRequest) GetOwnerTeamsOk() ([]OwnerTeamType, bool) {
	if o == nil || IsNil(o.OwnerTeams) {
		return nil, false
	}
	return o.OwnerTeams, true
}

// HasOwnerTeams returns a boolean if a field has been set.
func (o *ChangeOwnerTeamsRequest) HasOwnerTeams() bool {
	if o != nil && !IsNil(o.OwnerTeams) {
		return true
	}

	return false
}

// SetOwnerTeams gets a reference to the given []OwnerTeamType and assigns it to the OwnerTeams field.
func (o *ChangeOwnerTeamsRequest) SetOwnerTeams(v []OwnerTeamType) {
	o.OwnerTeams = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChangeOwnerTeamsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeOwnerTeamsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChangeOwnerTeamsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChangeOwnerTeamsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChangeOwnerTeamsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeOwnerTeamsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChangeOwnerTeamsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChangeOwnerTeamsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChangeOwnerTeamsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeOwnerTeamsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OwnerTeams) {
		toSerialize["ownerTeams"] = o.OwnerTeams
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChangeOwnerTeamsRequest struct {
	value *ChangeOwnerTeamsRequest
	isSet bool
}

func (v NullableChangeOwnerTeamsRequest) Get() *ChangeOwnerTeamsRequest {
	return v.value
}

func (v *NullableChangeOwnerTeamsRequest) Set(val *ChangeOwnerTeamsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeOwnerTeamsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeOwnerTeamsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeOwnerTeamsRequest(val *ChangeOwnerTeamsRequest) *NullableChangeOwnerTeamsRequest {
	return &NullableChangeOwnerTeamsRequest{value: val, isSet: true}
}

func (v NullableChangeOwnerTeamsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeOwnerTeamsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


