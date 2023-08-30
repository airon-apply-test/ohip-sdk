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

// checks if the EcertificateLocationTypesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EcertificateLocationTypesDetails{}

// EcertificateLocationTypesDetails Response object for fetching Ecertificate Location Types.
type EcertificateLocationTypesDetails struct {
	// List of Ecertificate Location Types.
	EcertificateLocationTypes []EcertificateLocationTypeType `json:"ecertificateLocationTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewEcertificateLocationTypesDetails instantiates a new EcertificateLocationTypesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcertificateLocationTypesDetails() *EcertificateLocationTypesDetails {
	this := EcertificateLocationTypesDetails{}
	return &this
}

// NewEcertificateLocationTypesDetailsWithDefaults instantiates a new EcertificateLocationTypesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcertificateLocationTypesDetailsWithDefaults() *EcertificateLocationTypesDetails {
	this := EcertificateLocationTypesDetails{}
	return &this
}

// GetEcertificateLocationTypes returns the EcertificateLocationTypes field value if set, zero value otherwise.
func (o *EcertificateLocationTypesDetails) GetEcertificateLocationTypes() []EcertificateLocationTypeType {
	if o == nil || IsNil(o.EcertificateLocationTypes) {
		var ret []EcertificateLocationTypeType
		return ret
	}
	return o.EcertificateLocationTypes
}

// GetEcertificateLocationTypesOk returns a tuple with the EcertificateLocationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcertificateLocationTypesDetails) GetEcertificateLocationTypesOk() ([]EcertificateLocationTypeType, bool) {
	if o == nil || IsNil(o.EcertificateLocationTypes) {
		return nil, false
	}
	return o.EcertificateLocationTypes, true
}

// HasEcertificateLocationTypes returns a boolean if a field has been set.
func (o *EcertificateLocationTypesDetails) HasEcertificateLocationTypes() bool {
	if o != nil && !IsNil(o.EcertificateLocationTypes) {
		return true
	}

	return false
}

// SetEcertificateLocationTypes gets a reference to the given []EcertificateLocationTypeType and assigns it to the EcertificateLocationTypes field.
func (o *EcertificateLocationTypesDetails) SetEcertificateLocationTypes(v []EcertificateLocationTypeType) {
	o.EcertificateLocationTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EcertificateLocationTypesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcertificateLocationTypesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EcertificateLocationTypesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *EcertificateLocationTypesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *EcertificateLocationTypesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcertificateLocationTypesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *EcertificateLocationTypesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *EcertificateLocationTypesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o EcertificateLocationTypesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcertificateLocationTypesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EcertificateLocationTypes) {
		toSerialize["ecertificateLocationTypes"] = o.EcertificateLocationTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableEcertificateLocationTypesDetails struct {
	value *EcertificateLocationTypesDetails
	isSet bool
}

func (v NullableEcertificateLocationTypesDetails) Get() *EcertificateLocationTypesDetails {
	return v.value
}

func (v *NullableEcertificateLocationTypesDetails) Set(val *EcertificateLocationTypesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEcertificateLocationTypesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEcertificateLocationTypesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcertificateLocationTypesDetails(val *EcertificateLocationTypesDetails) *NullableEcertificateLocationTypesDetails {
	return &NullableEcertificateLocationTypesDetails{value: val, isSet: true}
}

func (v NullableEcertificateLocationTypesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcertificateLocationTypesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


