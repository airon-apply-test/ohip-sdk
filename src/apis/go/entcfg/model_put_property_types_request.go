/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PutPropertyTypesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutPropertyTypesRequest{}

// PutPropertyTypesRequest struct for PutPropertyTypesRequest
type PutPropertyTypesRequest struct {
	// List of Property Types.
	PropertyTypes []PropertyTypeType `json:"propertyTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutPropertyTypesRequest instantiates a new PutPropertyTypesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutPropertyTypesRequest() *PutPropertyTypesRequest {
	this := PutPropertyTypesRequest{}
	return &this
}

// NewPutPropertyTypesRequestWithDefaults instantiates a new PutPropertyTypesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutPropertyTypesRequestWithDefaults() *PutPropertyTypesRequest {
	this := PutPropertyTypesRequest{}
	return &this
}

// GetPropertyTypes returns the PropertyTypes field value if set, zero value otherwise.
func (o *PutPropertyTypesRequest) GetPropertyTypes() []PropertyTypeType {
	if o == nil || IsNil(o.PropertyTypes) {
		var ret []PropertyTypeType
		return ret
	}
	return o.PropertyTypes
}

// GetPropertyTypesOk returns a tuple with the PropertyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutPropertyTypesRequest) GetPropertyTypesOk() ([]PropertyTypeType, bool) {
	if o == nil || IsNil(o.PropertyTypes) {
		return nil, false
	}
	return o.PropertyTypes, true
}

// HasPropertyTypes returns a boolean if a field has been set.
func (o *PutPropertyTypesRequest) HasPropertyTypes() bool {
	if o != nil && !IsNil(o.PropertyTypes) {
		return true
	}

	return false
}

// SetPropertyTypes gets a reference to the given []PropertyTypeType and assigns it to the PropertyTypes field.
func (o *PutPropertyTypesRequest) SetPropertyTypes(v []PropertyTypeType) {
	o.PropertyTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutPropertyTypesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutPropertyTypesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutPropertyTypesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutPropertyTypesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutPropertyTypesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutPropertyTypesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutPropertyTypesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutPropertyTypesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutPropertyTypesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutPropertyTypesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyTypes) {
		toSerialize["propertyTypes"] = o.PropertyTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutPropertyTypesRequest struct {
	value *PutPropertyTypesRequest
	isSet bool
}

func (v NullablePutPropertyTypesRequest) Get() *PutPropertyTypesRequest {
	return v.value
}

func (v *NullablePutPropertyTypesRequest) Set(val *PutPropertyTypesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutPropertyTypesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutPropertyTypesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutPropertyTypesRequest(val *PutPropertyTypesRequest) *NullablePutPropertyTypesRequest {
	return &NullablePutPropertyTypesRequest{value: val, isSet: true}
}

func (v NullablePutPropertyTypesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutPropertyTypesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


