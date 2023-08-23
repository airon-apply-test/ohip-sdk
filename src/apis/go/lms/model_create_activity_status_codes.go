/*
OPERA Cloud Leisure Management API

APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateActivityStatusCodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateActivityStatusCodes{}

// CreateActivityStatusCodes Request object for creating Activity Status Codes.
type CreateActivityStatusCodes struct {
	// Activity Status Codes object.
	ActivityStatusCodes []ActivityStatusCodeType `json:"activityStatusCodes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCreateActivityStatusCodes instantiates a new CreateActivityStatusCodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateActivityStatusCodes() *CreateActivityStatusCodes {
	this := CreateActivityStatusCodes{}
	return &this
}

// NewCreateActivityStatusCodesWithDefaults instantiates a new CreateActivityStatusCodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateActivityStatusCodesWithDefaults() *CreateActivityStatusCodes {
	this := CreateActivityStatusCodes{}
	return &this
}

// GetActivityStatusCodes returns the ActivityStatusCodes field value if set, zero value otherwise.
func (o *CreateActivityStatusCodes) GetActivityStatusCodes() []ActivityStatusCodeType {
	if o == nil || IsNil(o.ActivityStatusCodes) {
		var ret []ActivityStatusCodeType
		return ret
	}
	return o.ActivityStatusCodes
}

// GetActivityStatusCodesOk returns a tuple with the ActivityStatusCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActivityStatusCodes) GetActivityStatusCodesOk() ([]ActivityStatusCodeType, bool) {
	if o == nil || IsNil(o.ActivityStatusCodes) {
		return nil, false
	}
	return o.ActivityStatusCodes, true
}

// HasActivityStatusCodes returns a boolean if a field has been set.
func (o *CreateActivityStatusCodes) HasActivityStatusCodes() bool {
	if o != nil && !IsNil(o.ActivityStatusCodes) {
		return true
	}

	return false
}

// SetActivityStatusCodes gets a reference to the given []ActivityStatusCodeType and assigns it to the ActivityStatusCodes field.
func (o *CreateActivityStatusCodes) SetActivityStatusCodes(v []ActivityStatusCodeType) {
	o.ActivityStatusCodes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateActivityStatusCodes) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActivityStatusCodes) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateActivityStatusCodes) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CreateActivityStatusCodes) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreateActivityStatusCodes) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActivityStatusCodes) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreateActivityStatusCodes) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CreateActivityStatusCodes) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CreateActivityStatusCodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateActivityStatusCodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityStatusCodes) {
		toSerialize["activityStatusCodes"] = o.ActivityStatusCodes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCreateActivityStatusCodes struct {
	value *CreateActivityStatusCodes
	isSet bool
}

func (v NullableCreateActivityStatusCodes) Get() *CreateActivityStatusCodes {
	return v.value
}

func (v *NullableCreateActivityStatusCodes) Set(val *CreateActivityStatusCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateActivityStatusCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateActivityStatusCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateActivityStatusCodes(val *CreateActivityStatusCodes) *NullableCreateActivityStatusCodes {
	return &NullableCreateActivityStatusCodes{value: val, isSet: true}
}

func (v NullableCreateActivityStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateActivityStatusCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


