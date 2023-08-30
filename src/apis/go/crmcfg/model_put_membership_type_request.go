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

// checks if the PutMembershipTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutMembershipTypeRequest{}

// PutMembershipTypeRequest struct for PutMembershipTypeRequest
type PutMembershipTypeRequest struct {
	// A collection of MembershipTypes with information that needs to be changed.
	MembershipTypeChangeInstructions []MembershipTypeChangeInstructionType `json:"membershipTypeChangeInstructions,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutMembershipTypeRequest instantiates a new PutMembershipTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutMembershipTypeRequest() *PutMembershipTypeRequest {
	this := PutMembershipTypeRequest{}
	return &this
}

// NewPutMembershipTypeRequestWithDefaults instantiates a new PutMembershipTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutMembershipTypeRequestWithDefaults() *PutMembershipTypeRequest {
	this := PutMembershipTypeRequest{}
	return &this
}

// GetMembershipTypeChangeInstructions returns the MembershipTypeChangeInstructions field value if set, zero value otherwise.
func (o *PutMembershipTypeRequest) GetMembershipTypeChangeInstructions() []MembershipTypeChangeInstructionType {
	if o == nil || IsNil(o.MembershipTypeChangeInstructions) {
		var ret []MembershipTypeChangeInstructionType
		return ret
	}
	return o.MembershipTypeChangeInstructions
}

// GetMembershipTypeChangeInstructionsOk returns a tuple with the MembershipTypeChangeInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutMembershipTypeRequest) GetMembershipTypeChangeInstructionsOk() ([]MembershipTypeChangeInstructionType, bool) {
	if o == nil || IsNil(o.MembershipTypeChangeInstructions) {
		return nil, false
	}
	return o.MembershipTypeChangeInstructions, true
}

// HasMembershipTypeChangeInstructions returns a boolean if a field has been set.
func (o *PutMembershipTypeRequest) HasMembershipTypeChangeInstructions() bool {
	if o != nil && !IsNil(o.MembershipTypeChangeInstructions) {
		return true
	}

	return false
}

// SetMembershipTypeChangeInstructions gets a reference to the given []MembershipTypeChangeInstructionType and assigns it to the MembershipTypeChangeInstructions field.
func (o *PutMembershipTypeRequest) SetMembershipTypeChangeInstructions(v []MembershipTypeChangeInstructionType) {
	o.MembershipTypeChangeInstructions = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutMembershipTypeRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutMembershipTypeRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutMembershipTypeRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutMembershipTypeRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutMembershipTypeRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutMembershipTypeRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutMembershipTypeRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutMembershipTypeRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutMembershipTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutMembershipTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MembershipTypeChangeInstructions) {
		toSerialize["membershipTypeChangeInstructions"] = o.MembershipTypeChangeInstructions
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutMembershipTypeRequest struct {
	value *PutMembershipTypeRequest
	isSet bool
}

func (v NullablePutMembershipTypeRequest) Get() *PutMembershipTypeRequest {
	return v.value
}

func (v *NullablePutMembershipTypeRequest) Set(val *PutMembershipTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutMembershipTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutMembershipTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutMembershipTypeRequest(val *PutMembershipTypeRequest) *NullablePutMembershipTypeRequest {
	return &NullablePutMembershipTypeRequest{value: val, isSet: true}
}

func (v NullablePutMembershipTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutMembershipTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


