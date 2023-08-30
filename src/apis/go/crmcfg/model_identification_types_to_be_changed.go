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

// checks if the IdentificationTypesToBeChanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentificationTypesToBeChanged{}

// IdentificationTypesToBeChanged Request object for changing Identification Types.
type IdentificationTypesToBeChanged struct {
	// List of Identification Types.
	IdentificationTypes []IdentificationTypeType `json:"identificationTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewIdentificationTypesToBeChanged instantiates a new IdentificationTypesToBeChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentificationTypesToBeChanged() *IdentificationTypesToBeChanged {
	this := IdentificationTypesToBeChanged{}
	return &this
}

// NewIdentificationTypesToBeChangedWithDefaults instantiates a new IdentificationTypesToBeChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentificationTypesToBeChangedWithDefaults() *IdentificationTypesToBeChanged {
	this := IdentificationTypesToBeChanged{}
	return &this
}

// GetIdentificationTypes returns the IdentificationTypes field value if set, zero value otherwise.
func (o *IdentificationTypesToBeChanged) GetIdentificationTypes() []IdentificationTypeType {
	if o == nil || IsNil(o.IdentificationTypes) {
		var ret []IdentificationTypeType
		return ret
	}
	return o.IdentificationTypes
}

// GetIdentificationTypesOk returns a tuple with the IdentificationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationTypesToBeChanged) GetIdentificationTypesOk() ([]IdentificationTypeType, bool) {
	if o == nil || IsNil(o.IdentificationTypes) {
		return nil, false
	}
	return o.IdentificationTypes, true
}

// HasIdentificationTypes returns a boolean if a field has been set.
func (o *IdentificationTypesToBeChanged) HasIdentificationTypes() bool {
	if o != nil && !IsNil(o.IdentificationTypes) {
		return true
	}

	return false
}

// SetIdentificationTypes gets a reference to the given []IdentificationTypeType and assigns it to the IdentificationTypes field.
func (o *IdentificationTypesToBeChanged) SetIdentificationTypes(v []IdentificationTypeType) {
	o.IdentificationTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IdentificationTypesToBeChanged) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationTypesToBeChanged) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IdentificationTypesToBeChanged) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *IdentificationTypesToBeChanged) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *IdentificationTypesToBeChanged) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationTypesToBeChanged) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *IdentificationTypesToBeChanged) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *IdentificationTypesToBeChanged) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o IdentificationTypesToBeChanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentificationTypesToBeChanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdentificationTypes) {
		toSerialize["identificationTypes"] = o.IdentificationTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableIdentificationTypesToBeChanged struct {
	value *IdentificationTypesToBeChanged
	isSet bool
}

func (v NullableIdentificationTypesToBeChanged) Get() *IdentificationTypesToBeChanged {
	return v.value
}

func (v *NullableIdentificationTypesToBeChanged) Set(val *IdentificationTypesToBeChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationTypesToBeChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationTypesToBeChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationTypesToBeChanged(val *IdentificationTypesToBeChanged) *NullableIdentificationTypesToBeChanged {
	return &NullableIdentificationTypesToBeChanged{value: val, isSet: true}
}

func (v NullableIdentificationTypesToBeChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationTypesToBeChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


