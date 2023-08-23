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

// checks if the AddressTypesCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressTypesCriteria{}

// AddressTypesCriteria Request object for creating a new Address Type.
type AddressTypesCriteria struct {
	// Communication Role Enumeration element.
	AddressTypes []AddressTypeType `json:"addressTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewAddressTypesCriteria instantiates a new AddressTypesCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTypesCriteria() *AddressTypesCriteria {
	this := AddressTypesCriteria{}
	return &this
}

// NewAddressTypesCriteriaWithDefaults instantiates a new AddressTypesCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTypesCriteriaWithDefaults() *AddressTypesCriteria {
	this := AddressTypesCriteria{}
	return &this
}

// GetAddressTypes returns the AddressTypes field value if set, zero value otherwise.
func (o *AddressTypesCriteria) GetAddressTypes() []AddressTypeType {
	if o == nil || IsNil(o.AddressTypes) {
		var ret []AddressTypeType
		return ret
	}
	return o.AddressTypes
}

// GetAddressTypesOk returns a tuple with the AddressTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTypesCriteria) GetAddressTypesOk() ([]AddressTypeType, bool) {
	if o == nil || IsNil(o.AddressTypes) {
		return nil, false
	}
	return o.AddressTypes, true
}

// HasAddressTypes returns a boolean if a field has been set.
func (o *AddressTypesCriteria) HasAddressTypes() bool {
	if o != nil && !IsNil(o.AddressTypes) {
		return true
	}

	return false
}

// SetAddressTypes gets a reference to the given []AddressTypeType and assigns it to the AddressTypes field.
func (o *AddressTypesCriteria) SetAddressTypes(v []AddressTypeType) {
	o.AddressTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AddressTypesCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTypesCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AddressTypesCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *AddressTypesCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AddressTypesCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTypesCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AddressTypesCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *AddressTypesCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o AddressTypesCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTypesCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressTypes) {
		toSerialize["addressTypes"] = o.AddressTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAddressTypesCriteria struct {
	value *AddressTypesCriteria
	isSet bool
}

func (v NullableAddressTypesCriteria) Get() *AddressTypesCriteria {
	return v.value
}

func (v *NullableAddressTypesCriteria) Set(val *AddressTypesCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTypesCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTypesCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTypesCriteria(val *AddressTypesCriteria) *NullableAddressTypesCriteria {
	return &NullableAddressTypesCriteria{value: val, isSet: true}
}

func (v NullableAddressTypesCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTypesCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


