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

// checks if the PutProfileInactiveReasonsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutProfileInactiveReasonsRequest{}

// PutProfileInactiveReasonsRequest struct for PutProfileInactiveReasonsRequest
type PutProfileInactiveReasonsRequest struct {
	// List of Profile Inactive Reasons.
	ProfileInactiveReasons []ProfileInactiveReasonType `json:"profileInactiveReasons,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutProfileInactiveReasonsRequest instantiates a new PutProfileInactiveReasonsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutProfileInactiveReasonsRequest() *PutProfileInactiveReasonsRequest {
	this := PutProfileInactiveReasonsRequest{}
	return &this
}

// NewPutProfileInactiveReasonsRequestWithDefaults instantiates a new PutProfileInactiveReasonsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutProfileInactiveReasonsRequestWithDefaults() *PutProfileInactiveReasonsRequest {
	this := PutProfileInactiveReasonsRequest{}
	return &this
}

// GetProfileInactiveReasons returns the ProfileInactiveReasons field value if set, zero value otherwise.
func (o *PutProfileInactiveReasonsRequest) GetProfileInactiveReasons() []ProfileInactiveReasonType {
	if o == nil || IsNil(o.ProfileInactiveReasons) {
		var ret []ProfileInactiveReasonType
		return ret
	}
	return o.ProfileInactiveReasons
}

// GetProfileInactiveReasonsOk returns a tuple with the ProfileInactiveReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutProfileInactiveReasonsRequest) GetProfileInactiveReasonsOk() ([]ProfileInactiveReasonType, bool) {
	if o == nil || IsNil(o.ProfileInactiveReasons) {
		return nil, false
	}
	return o.ProfileInactiveReasons, true
}

// HasProfileInactiveReasons returns a boolean if a field has been set.
func (o *PutProfileInactiveReasonsRequest) HasProfileInactiveReasons() bool {
	if o != nil && !IsNil(o.ProfileInactiveReasons) {
		return true
	}

	return false
}

// SetProfileInactiveReasons gets a reference to the given []ProfileInactiveReasonType and assigns it to the ProfileInactiveReasons field.
func (o *PutProfileInactiveReasonsRequest) SetProfileInactiveReasons(v []ProfileInactiveReasonType) {
	o.ProfileInactiveReasons = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutProfileInactiveReasonsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutProfileInactiveReasonsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutProfileInactiveReasonsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutProfileInactiveReasonsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutProfileInactiveReasonsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutProfileInactiveReasonsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutProfileInactiveReasonsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutProfileInactiveReasonsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutProfileInactiveReasonsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutProfileInactiveReasonsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileInactiveReasons) {
		toSerialize["profileInactiveReasons"] = o.ProfileInactiveReasons
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutProfileInactiveReasonsRequest struct {
	value *PutProfileInactiveReasonsRequest
	isSet bool
}

func (v NullablePutProfileInactiveReasonsRequest) Get() *PutProfileInactiveReasonsRequest {
	return v.value
}

func (v *NullablePutProfileInactiveReasonsRequest) Set(val *PutProfileInactiveReasonsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutProfileInactiveReasonsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutProfileInactiveReasonsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutProfileInactiveReasonsRequest(val *PutProfileInactiveReasonsRequest) *NullablePutProfileInactiveReasonsRequest {
	return &NullablePutProfileInactiveReasonsRequest{value: val, isSet: true}
}

func (v NullablePutProfileInactiveReasonsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutProfileInactiveReasonsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


