/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TraceResolveType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceResolveType{}

// TraceResolveType struct for TraceResolveType
type TraceResolveType struct {
	// Date the trace was resolved
	ResolvedOn *string `json:"resolvedOn,omitempty"`
	// User that resolved the trace
	ResolvedBy *string `json:"resolvedBy,omitempty"`
}

// NewTraceResolveType instantiates a new TraceResolveType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceResolveType() *TraceResolveType {
	this := TraceResolveType{}
	return &this
}

// NewTraceResolveTypeWithDefaults instantiates a new TraceResolveType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceResolveTypeWithDefaults() *TraceResolveType {
	this := TraceResolveType{}
	return &this
}

// GetResolvedOn returns the ResolvedOn field value if set, zero value otherwise.
func (o *TraceResolveType) GetResolvedOn() string {
	if o == nil || IsNil(o.ResolvedOn) {
		var ret string
		return ret
	}
	return *o.ResolvedOn
}

// GetResolvedOnOk returns a tuple with the ResolvedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResolveType) GetResolvedOnOk() (*string, bool) {
	if o == nil || IsNil(o.ResolvedOn) {
		return nil, false
	}
	return o.ResolvedOn, true
}

// HasResolvedOn returns a boolean if a field has been set.
func (o *TraceResolveType) HasResolvedOn() bool {
	if o != nil && !IsNil(o.ResolvedOn) {
		return true
	}

	return false
}

// SetResolvedOn gets a reference to the given string and assigns it to the ResolvedOn field.
func (o *TraceResolveType) SetResolvedOn(v string) {
	o.ResolvedOn = &v
}

// GetResolvedBy returns the ResolvedBy field value if set, zero value otherwise.
func (o *TraceResolveType) GetResolvedBy() string {
	if o == nil || IsNil(o.ResolvedBy) {
		var ret string
		return ret
	}
	return *o.ResolvedBy
}

// GetResolvedByOk returns a tuple with the ResolvedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceResolveType) GetResolvedByOk() (*string, bool) {
	if o == nil || IsNil(o.ResolvedBy) {
		return nil, false
	}
	return o.ResolvedBy, true
}

// HasResolvedBy returns a boolean if a field has been set.
func (o *TraceResolveType) HasResolvedBy() bool {
	if o != nil && !IsNil(o.ResolvedBy) {
		return true
	}

	return false
}

// SetResolvedBy gets a reference to the given string and assigns it to the ResolvedBy field.
func (o *TraceResolveType) SetResolvedBy(v string) {
	o.ResolvedBy = &v
}

func (o TraceResolveType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceResolveType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResolvedOn) {
		toSerialize["resolvedOn"] = o.ResolvedOn
	}
	if !IsNil(o.ResolvedBy) {
		toSerialize["resolvedBy"] = o.ResolvedBy
	}
	return toSerialize, nil
}

type NullableTraceResolveType struct {
	value *TraceResolveType
	isSet bool
}

func (v NullableTraceResolveType) Get() *TraceResolveType {
	return v.value
}

func (v *NullableTraceResolveType) Set(val *TraceResolveType) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceResolveType) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceResolveType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceResolveType(val *TraceResolveType) *NullableTraceResolveType {
	return &NullableTraceResolveType{value: val, isSet: true}
}

func (v NullableTraceResolveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceResolveType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


