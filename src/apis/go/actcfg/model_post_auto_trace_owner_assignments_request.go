/*
OPERA Cloud Activity Management API

APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PostAutoTraceOwnerAssignmentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAutoTraceOwnerAssignmentsRequest{}

// PostAutoTraceOwnerAssignmentsRequest struct for PostAutoTraceOwnerAssignmentsRequest
type PostAutoTraceOwnerAssignmentsRequest struct {
	// Detailed information of trace owner assignment.
	AutoTraceOwnerAssignments []AutoTraceOwnerAssignmentType `json:"autoTraceOwnerAssignments,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostAutoTraceOwnerAssignmentsRequest instantiates a new PostAutoTraceOwnerAssignmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAutoTraceOwnerAssignmentsRequest() *PostAutoTraceOwnerAssignmentsRequest {
	this := PostAutoTraceOwnerAssignmentsRequest{}
	return &this
}

// NewPostAutoTraceOwnerAssignmentsRequestWithDefaults instantiates a new PostAutoTraceOwnerAssignmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAutoTraceOwnerAssignmentsRequestWithDefaults() *PostAutoTraceOwnerAssignmentsRequest {
	this := PostAutoTraceOwnerAssignmentsRequest{}
	return &this
}

// GetAutoTraceOwnerAssignments returns the AutoTraceOwnerAssignments field value if set, zero value otherwise.
func (o *PostAutoTraceOwnerAssignmentsRequest) GetAutoTraceOwnerAssignments() []AutoTraceOwnerAssignmentType {
	if o == nil || IsNil(o.AutoTraceOwnerAssignments) {
		var ret []AutoTraceOwnerAssignmentType
		return ret
	}
	return o.AutoTraceOwnerAssignments
}

// GetAutoTraceOwnerAssignmentsOk returns a tuple with the AutoTraceOwnerAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAutoTraceOwnerAssignmentsRequest) GetAutoTraceOwnerAssignmentsOk() ([]AutoTraceOwnerAssignmentType, bool) {
	if o == nil || IsNil(o.AutoTraceOwnerAssignments) {
		return nil, false
	}
	return o.AutoTraceOwnerAssignments, true
}

// HasAutoTraceOwnerAssignments returns a boolean if a field has been set.
func (o *PostAutoTraceOwnerAssignmentsRequest) HasAutoTraceOwnerAssignments() bool {
	if o != nil && !IsNil(o.AutoTraceOwnerAssignments) {
		return true
	}

	return false
}

// SetAutoTraceOwnerAssignments gets a reference to the given []AutoTraceOwnerAssignmentType and assigns it to the AutoTraceOwnerAssignments field.
func (o *PostAutoTraceOwnerAssignmentsRequest) SetAutoTraceOwnerAssignments(v []AutoTraceOwnerAssignmentType) {
	o.AutoTraceOwnerAssignments = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostAutoTraceOwnerAssignmentsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAutoTraceOwnerAssignmentsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostAutoTraceOwnerAssignmentsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostAutoTraceOwnerAssignmentsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostAutoTraceOwnerAssignmentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAutoTraceOwnerAssignmentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoTraceOwnerAssignments) {
		toSerialize["autoTraceOwnerAssignments"] = o.AutoTraceOwnerAssignments
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostAutoTraceOwnerAssignmentsRequest struct {
	value *PostAutoTraceOwnerAssignmentsRequest
	isSet bool
}

func (v NullablePostAutoTraceOwnerAssignmentsRequest) Get() *PostAutoTraceOwnerAssignmentsRequest {
	return v.value
}

func (v *NullablePostAutoTraceOwnerAssignmentsRequest) Set(val *PostAutoTraceOwnerAssignmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAutoTraceOwnerAssignmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAutoTraceOwnerAssignmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAutoTraceOwnerAssignmentsRequest(val *PostAutoTraceOwnerAssignmentsRequest) *NullablePostAutoTraceOwnerAssignmentsRequest {
	return &NullablePostAutoTraceOwnerAssignmentsRequest{value: val, isSet: true}
}

func (v NullablePostAutoTraceOwnerAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAutoTraceOwnerAssignmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


