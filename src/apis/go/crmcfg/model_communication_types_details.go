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

// checks if the CommunicationTypesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationTypesDetails{}

// CommunicationTypesDetails Response object for fetching Communication Types.
type CommunicationTypesDetails struct {
	// Communication Type details.
	CommunicationTypes []CommunicationTypeType `json:"communicationTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCommunicationTypesDetails instantiates a new CommunicationTypesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationTypesDetails() *CommunicationTypesDetails {
	this := CommunicationTypesDetails{}
	return &this
}

// NewCommunicationTypesDetailsWithDefaults instantiates a new CommunicationTypesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationTypesDetailsWithDefaults() *CommunicationTypesDetails {
	this := CommunicationTypesDetails{}
	return &this
}

// GetCommunicationTypes returns the CommunicationTypes field value if set, zero value otherwise.
func (o *CommunicationTypesDetails) GetCommunicationTypes() []CommunicationTypeType {
	if o == nil || IsNil(o.CommunicationTypes) {
		var ret []CommunicationTypeType
		return ret
	}
	return o.CommunicationTypes
}

// GetCommunicationTypesOk returns a tuple with the CommunicationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationTypesDetails) GetCommunicationTypesOk() ([]CommunicationTypeType, bool) {
	if o == nil || IsNil(o.CommunicationTypes) {
		return nil, false
	}
	return o.CommunicationTypes, true
}

// HasCommunicationTypes returns a boolean if a field has been set.
func (o *CommunicationTypesDetails) HasCommunicationTypes() bool {
	if o != nil && !IsNil(o.CommunicationTypes) {
		return true
	}

	return false
}

// SetCommunicationTypes gets a reference to the given []CommunicationTypeType and assigns it to the CommunicationTypes field.
func (o *CommunicationTypesDetails) SetCommunicationTypes(v []CommunicationTypeType) {
	o.CommunicationTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CommunicationTypesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationTypesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CommunicationTypesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CommunicationTypesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CommunicationTypesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationTypesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CommunicationTypesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CommunicationTypesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CommunicationTypesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationTypesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommunicationTypes) {
		toSerialize["communicationTypes"] = o.CommunicationTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCommunicationTypesDetails struct {
	value *CommunicationTypesDetails
	isSet bool
}

func (v NullableCommunicationTypesDetails) Get() *CommunicationTypesDetails {
	return v.value
}

func (v *NullableCommunicationTypesDetails) Set(val *CommunicationTypesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationTypesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationTypesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationTypesDetails(val *CommunicationTypesDetails) *NullableCommunicationTypesDetails {
	return &NullableCommunicationTypesDetails{value: val, isSet: true}
}

func (v NullableCommunicationTypesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationTypesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


