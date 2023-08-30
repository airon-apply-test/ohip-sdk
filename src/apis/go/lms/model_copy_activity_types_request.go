/*
OPERA Cloud Leisure Management API

APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lms

import (
	"encoding/json"
)

// checks if the CopyActivityTypesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyActivityTypesRequest{}

// CopyActivityTypesRequest struct for CopyActivityTypesRequest
type CopyActivityTypesRequest struct {
	// Information needed to copy configuration code from one property to the other.
	CopyInstructions []CopyConfigurationCodeType `json:"copyInstructions,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCopyActivityTypesRequest instantiates a new CopyActivityTypesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyActivityTypesRequest() *CopyActivityTypesRequest {
	this := CopyActivityTypesRequest{}
	return &this
}

// NewCopyActivityTypesRequestWithDefaults instantiates a new CopyActivityTypesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyActivityTypesRequestWithDefaults() *CopyActivityTypesRequest {
	this := CopyActivityTypesRequest{}
	return &this
}

// GetCopyInstructions returns the CopyInstructions field value if set, zero value otherwise.
func (o *CopyActivityTypesRequest) GetCopyInstructions() []CopyConfigurationCodeType {
	if o == nil || IsNil(o.CopyInstructions) {
		var ret []CopyConfigurationCodeType
		return ret
	}
	return o.CopyInstructions
}

// GetCopyInstructionsOk returns a tuple with the CopyInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyActivityTypesRequest) GetCopyInstructionsOk() ([]CopyConfigurationCodeType, bool) {
	if o == nil || IsNil(o.CopyInstructions) {
		return nil, false
	}
	return o.CopyInstructions, true
}

// HasCopyInstructions returns a boolean if a field has been set.
func (o *CopyActivityTypesRequest) HasCopyInstructions() bool {
	if o != nil && !IsNil(o.CopyInstructions) {
		return true
	}

	return false
}

// SetCopyInstructions gets a reference to the given []CopyConfigurationCodeType and assigns it to the CopyInstructions field.
func (o *CopyActivityTypesRequest) SetCopyInstructions(v []CopyConfigurationCodeType) {
	o.CopyInstructions = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CopyActivityTypesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyActivityTypesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CopyActivityTypesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CopyActivityTypesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CopyActivityTypesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyActivityTypesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CopyActivityTypesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CopyActivityTypesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CopyActivityTypesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyActivityTypesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CopyInstructions) {
		toSerialize["copyInstructions"] = o.CopyInstructions
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCopyActivityTypesRequest struct {
	value *CopyActivityTypesRequest
	isSet bool
}

func (v NullableCopyActivityTypesRequest) Get() *CopyActivityTypesRequest {
	return v.value
}

func (v *NullableCopyActivityTypesRequest) Set(val *CopyActivityTypesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyActivityTypesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyActivityTypesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyActivityTypesRequest(val *CopyActivityTypesRequest) *NullableCopyActivityTypesRequest {
	return &NullableCopyActivityTypesRequest{value: val, isSet: true}
}

func (v NullableCopyActivityTypesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyActivityTypesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


