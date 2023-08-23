/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PutInterfaceControlsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutInterfaceControlsRequest{}

// PutInterfaceControlsRequest struct for PutInterfaceControlsRequest
type PutInterfaceControlsRequest struct {
	IntegrationSystem *IntegrationSystemType `json:"integrationSystem,omitempty"`
	// OXI Parameters/Settings.
	InterfaceControls []InterfaceControlType `json:"interfaceControls,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutInterfaceControlsRequest instantiates a new PutInterfaceControlsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutInterfaceControlsRequest() *PutInterfaceControlsRequest {
	this := PutInterfaceControlsRequest{}
	return &this
}

// NewPutInterfaceControlsRequestWithDefaults instantiates a new PutInterfaceControlsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutInterfaceControlsRequestWithDefaults() *PutInterfaceControlsRequest {
	this := PutInterfaceControlsRequest{}
	return &this
}

// GetIntegrationSystem returns the IntegrationSystem field value if set, zero value otherwise.
func (o *PutInterfaceControlsRequest) GetIntegrationSystem() IntegrationSystemType {
	if o == nil || IsNil(o.IntegrationSystem) {
		var ret IntegrationSystemType
		return ret
	}
	return *o.IntegrationSystem
}

// GetIntegrationSystemOk returns a tuple with the IntegrationSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutInterfaceControlsRequest) GetIntegrationSystemOk() (*IntegrationSystemType, bool) {
	if o == nil || IsNil(o.IntegrationSystem) {
		return nil, false
	}
	return o.IntegrationSystem, true
}

// HasIntegrationSystem returns a boolean if a field has been set.
func (o *PutInterfaceControlsRequest) HasIntegrationSystem() bool {
	if o != nil && !IsNil(o.IntegrationSystem) {
		return true
	}

	return false
}

// SetIntegrationSystem gets a reference to the given IntegrationSystemType and assigns it to the IntegrationSystem field.
func (o *PutInterfaceControlsRequest) SetIntegrationSystem(v IntegrationSystemType) {
	o.IntegrationSystem = &v
}

// GetInterfaceControls returns the InterfaceControls field value if set, zero value otherwise.
func (o *PutInterfaceControlsRequest) GetInterfaceControls() []InterfaceControlType {
	if o == nil || IsNil(o.InterfaceControls) {
		var ret []InterfaceControlType
		return ret
	}
	return o.InterfaceControls
}

// GetInterfaceControlsOk returns a tuple with the InterfaceControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutInterfaceControlsRequest) GetInterfaceControlsOk() ([]InterfaceControlType, bool) {
	if o == nil || IsNil(o.InterfaceControls) {
		return nil, false
	}
	return o.InterfaceControls, true
}

// HasInterfaceControls returns a boolean if a field has been set.
func (o *PutInterfaceControlsRequest) HasInterfaceControls() bool {
	if o != nil && !IsNil(o.InterfaceControls) {
		return true
	}

	return false
}

// SetInterfaceControls gets a reference to the given []InterfaceControlType and assigns it to the InterfaceControls field.
func (o *PutInterfaceControlsRequest) SetInterfaceControls(v []InterfaceControlType) {
	o.InterfaceControls = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutInterfaceControlsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutInterfaceControlsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutInterfaceControlsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutInterfaceControlsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutInterfaceControlsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutInterfaceControlsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutInterfaceControlsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutInterfaceControlsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutInterfaceControlsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutInterfaceControlsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntegrationSystem) {
		toSerialize["integrationSystem"] = o.IntegrationSystem
	}
	if !IsNil(o.InterfaceControls) {
		toSerialize["interfaceControls"] = o.InterfaceControls
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutInterfaceControlsRequest struct {
	value *PutInterfaceControlsRequest
	isSet bool
}

func (v NullablePutInterfaceControlsRequest) Get() *PutInterfaceControlsRequest {
	return v.value
}

func (v *NullablePutInterfaceControlsRequest) Set(val *PutInterfaceControlsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutInterfaceControlsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutInterfaceControlsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutInterfaceControlsRequest(val *PutInterfaceControlsRequest) *NullablePutInterfaceControlsRequest {
	return &NullablePutInterfaceControlsRequest{value: val, isSet: true}
}

func (v NullablePutInterfaceControlsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutInterfaceControlsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


