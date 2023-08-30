/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the PutConversionCodeMappingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutConversionCodeMappingsRequest{}

// PutConversionCodeMappingsRequest struct for PutConversionCodeMappingsRequest
type PutConversionCodeMappingsRequest struct {
	IntegrationSystem *IntegrationSystemType `json:"integrationSystem,omitempty"`
	// List of Conversion Code Mappings.
	ConversionCodeMappings []ConversionCodeMappingType `json:"conversionCodeMappings,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutConversionCodeMappingsRequest instantiates a new PutConversionCodeMappingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutConversionCodeMappingsRequest() *PutConversionCodeMappingsRequest {
	this := PutConversionCodeMappingsRequest{}
	return &this
}

// NewPutConversionCodeMappingsRequestWithDefaults instantiates a new PutConversionCodeMappingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutConversionCodeMappingsRequestWithDefaults() *PutConversionCodeMappingsRequest {
	this := PutConversionCodeMappingsRequest{}
	return &this
}

// GetIntegrationSystem returns the IntegrationSystem field value if set, zero value otherwise.
func (o *PutConversionCodeMappingsRequest) GetIntegrationSystem() IntegrationSystemType {
	if o == nil || IsNil(o.IntegrationSystem) {
		var ret IntegrationSystemType
		return ret
	}
	return *o.IntegrationSystem
}

// GetIntegrationSystemOk returns a tuple with the IntegrationSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutConversionCodeMappingsRequest) GetIntegrationSystemOk() (*IntegrationSystemType, bool) {
	if o == nil || IsNil(o.IntegrationSystem) {
		return nil, false
	}
	return o.IntegrationSystem, true
}

// HasIntegrationSystem returns a boolean if a field has been set.
func (o *PutConversionCodeMappingsRequest) HasIntegrationSystem() bool {
	if o != nil && !IsNil(o.IntegrationSystem) {
		return true
	}

	return false
}

// SetIntegrationSystem gets a reference to the given IntegrationSystemType and assigns it to the IntegrationSystem field.
func (o *PutConversionCodeMappingsRequest) SetIntegrationSystem(v IntegrationSystemType) {
	o.IntegrationSystem = &v
}

// GetConversionCodeMappings returns the ConversionCodeMappings field value if set, zero value otherwise.
func (o *PutConversionCodeMappingsRequest) GetConversionCodeMappings() []ConversionCodeMappingType {
	if o == nil || IsNil(o.ConversionCodeMappings) {
		var ret []ConversionCodeMappingType
		return ret
	}
	return o.ConversionCodeMappings
}

// GetConversionCodeMappingsOk returns a tuple with the ConversionCodeMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutConversionCodeMappingsRequest) GetConversionCodeMappingsOk() ([]ConversionCodeMappingType, bool) {
	if o == nil || IsNil(o.ConversionCodeMappings) {
		return nil, false
	}
	return o.ConversionCodeMappings, true
}

// HasConversionCodeMappings returns a boolean if a field has been set.
func (o *PutConversionCodeMappingsRequest) HasConversionCodeMappings() bool {
	if o != nil && !IsNil(o.ConversionCodeMappings) {
		return true
	}

	return false
}

// SetConversionCodeMappings gets a reference to the given []ConversionCodeMappingType and assigns it to the ConversionCodeMappings field.
func (o *PutConversionCodeMappingsRequest) SetConversionCodeMappings(v []ConversionCodeMappingType) {
	o.ConversionCodeMappings = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutConversionCodeMappingsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutConversionCodeMappingsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutConversionCodeMappingsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutConversionCodeMappingsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutConversionCodeMappingsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutConversionCodeMappingsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutConversionCodeMappingsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutConversionCodeMappingsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutConversionCodeMappingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutConversionCodeMappingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntegrationSystem) {
		toSerialize["integrationSystem"] = o.IntegrationSystem
	}
	if !IsNil(o.ConversionCodeMappings) {
		toSerialize["conversionCodeMappings"] = o.ConversionCodeMappings
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutConversionCodeMappingsRequest struct {
	value *PutConversionCodeMappingsRequest
	isSet bool
}

func (v NullablePutConversionCodeMappingsRequest) Get() *PutConversionCodeMappingsRequest {
	return v.value
}

func (v *NullablePutConversionCodeMappingsRequest) Set(val *PutConversionCodeMappingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutConversionCodeMappingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutConversionCodeMappingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutConversionCodeMappingsRequest(val *PutConversionCodeMappingsRequest) *NullablePutConversionCodeMappingsRequest {
	return &NullablePutConversionCodeMappingsRequest{value: val, isSet: true}
}

func (v NullablePutConversionCodeMappingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutConversionCodeMappingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


