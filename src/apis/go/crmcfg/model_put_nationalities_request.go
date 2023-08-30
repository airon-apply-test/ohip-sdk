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

// checks if the PutNationalitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutNationalitiesRequest{}

// PutNationalitiesRequest struct for PutNationalitiesRequest
type PutNationalitiesRequest struct {
	// List of Nationalities.
	Nationalities []NationalityType `json:"nationalities,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutNationalitiesRequest instantiates a new PutNationalitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutNationalitiesRequest() *PutNationalitiesRequest {
	this := PutNationalitiesRequest{}
	return &this
}

// NewPutNationalitiesRequestWithDefaults instantiates a new PutNationalitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutNationalitiesRequestWithDefaults() *PutNationalitiesRequest {
	this := PutNationalitiesRequest{}
	return &this
}

// GetNationalities returns the Nationalities field value if set, zero value otherwise.
func (o *PutNationalitiesRequest) GetNationalities() []NationalityType {
	if o == nil || IsNil(o.Nationalities) {
		var ret []NationalityType
		return ret
	}
	return o.Nationalities
}

// GetNationalitiesOk returns a tuple with the Nationalities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutNationalitiesRequest) GetNationalitiesOk() ([]NationalityType, bool) {
	if o == nil || IsNil(o.Nationalities) {
		return nil, false
	}
	return o.Nationalities, true
}

// HasNationalities returns a boolean if a field has been set.
func (o *PutNationalitiesRequest) HasNationalities() bool {
	if o != nil && !IsNil(o.Nationalities) {
		return true
	}

	return false
}

// SetNationalities gets a reference to the given []NationalityType and assigns it to the Nationalities field.
func (o *PutNationalitiesRequest) SetNationalities(v []NationalityType) {
	o.Nationalities = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutNationalitiesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutNationalitiesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutNationalitiesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutNationalitiesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutNationalitiesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutNationalitiesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutNationalitiesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutNationalitiesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutNationalitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutNationalitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nationalities) {
		toSerialize["nationalities"] = o.Nationalities
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutNationalitiesRequest struct {
	value *PutNationalitiesRequest
	isSet bool
}

func (v NullablePutNationalitiesRequest) Get() *PutNationalitiesRequest {
	return v.value
}

func (v *NullablePutNationalitiesRequest) Set(val *PutNationalitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutNationalitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutNationalitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutNationalitiesRequest(val *PutNationalitiesRequest) *NullablePutNationalitiesRequest {
	return &NullablePutNationalitiesRequest{value: val, isSet: true}
}

func (v NullablePutNationalitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutNationalitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


