/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PutBlockAllocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBlockAllocationRequest{}

// PutBlockAllocationRequest struct for PutBlockAllocationRequest
type PutBlockAllocationRequest struct {
	Criteria *AllocationGridByRoomTypesType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutBlockAllocationRequest instantiates a new PutBlockAllocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBlockAllocationRequest() *PutBlockAllocationRequest {
	this := PutBlockAllocationRequest{}
	return &this
}

// NewPutBlockAllocationRequestWithDefaults instantiates a new PutBlockAllocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBlockAllocationRequestWithDefaults() *PutBlockAllocationRequest {
	this := PutBlockAllocationRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PutBlockAllocationRequest) GetCriteria() AllocationGridByRoomTypesType {
	if o == nil || IsNil(o.Criteria) {
		var ret AllocationGridByRoomTypesType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBlockAllocationRequest) GetCriteriaOk() (*AllocationGridByRoomTypesType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PutBlockAllocationRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AllocationGridByRoomTypesType and assigns it to the Criteria field.
func (o *PutBlockAllocationRequest) SetCriteria(v AllocationGridByRoomTypesType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutBlockAllocationRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBlockAllocationRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutBlockAllocationRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutBlockAllocationRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutBlockAllocationRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBlockAllocationRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutBlockAllocationRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutBlockAllocationRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutBlockAllocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBlockAllocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutBlockAllocationRequest struct {
	value *PutBlockAllocationRequest
	isSet bool
}

func (v NullablePutBlockAllocationRequest) Get() *PutBlockAllocationRequest {
	return v.value
}

func (v *NullablePutBlockAllocationRequest) Set(val *PutBlockAllocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBlockAllocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBlockAllocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBlockAllocationRequest(val *PutBlockAllocationRequest) *NullablePutBlockAllocationRequest {
	return &NullablePutBlockAllocationRequest{value: val, isSet: true}
}

func (v NullablePutBlockAllocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBlockAllocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


