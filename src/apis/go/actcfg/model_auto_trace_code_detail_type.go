/*
OPERA Cloud Activity Management API

APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actcfg

import (
	"encoding/json"
)

// checks if the AutoTraceCodeDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTraceCodeDetailType{}

// AutoTraceCodeDetailType Auto Trace Code detail information.
type AutoTraceCodeDetailType struct {
	// Trace Code.
	TraceCode *string `json:"traceCode,omitempty"`
	// Description for the Trace Code.
	Description *string `json:"description,omitempty"`
	TraceGroup *AutoTraceGroupConfigType `json:"traceGroup,omitempty"`
}

// NewAutoTraceCodeDetailType instantiates a new AutoTraceCodeDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTraceCodeDetailType() *AutoTraceCodeDetailType {
	this := AutoTraceCodeDetailType{}
	return &this
}

// NewAutoTraceCodeDetailTypeWithDefaults instantiates a new AutoTraceCodeDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTraceCodeDetailTypeWithDefaults() *AutoTraceCodeDetailType {
	this := AutoTraceCodeDetailType{}
	return &this
}

// GetTraceCode returns the TraceCode field value if set, zero value otherwise.
func (o *AutoTraceCodeDetailType) GetTraceCode() string {
	if o == nil || IsNil(o.TraceCode) {
		var ret string
		return ret
	}
	return *o.TraceCode
}

// GetTraceCodeOk returns a tuple with the TraceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceCodeDetailType) GetTraceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TraceCode) {
		return nil, false
	}
	return o.TraceCode, true
}

// HasTraceCode returns a boolean if a field has been set.
func (o *AutoTraceCodeDetailType) HasTraceCode() bool {
	if o != nil && !IsNil(o.TraceCode) {
		return true
	}

	return false
}

// SetTraceCode gets a reference to the given string and assigns it to the TraceCode field.
func (o *AutoTraceCodeDetailType) SetTraceCode(v string) {
	o.TraceCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AutoTraceCodeDetailType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceCodeDetailType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AutoTraceCodeDetailType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AutoTraceCodeDetailType) SetDescription(v string) {
	o.Description = &v
}

// GetTraceGroup returns the TraceGroup field value if set, zero value otherwise.
func (o *AutoTraceCodeDetailType) GetTraceGroup() AutoTraceGroupConfigType {
	if o == nil || IsNil(o.TraceGroup) {
		var ret AutoTraceGroupConfigType
		return ret
	}
	return *o.TraceGroup
}

// GetTraceGroupOk returns a tuple with the TraceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceCodeDetailType) GetTraceGroupOk() (*AutoTraceGroupConfigType, bool) {
	if o == nil || IsNil(o.TraceGroup) {
		return nil, false
	}
	return o.TraceGroup, true
}

// HasTraceGroup returns a boolean if a field has been set.
func (o *AutoTraceCodeDetailType) HasTraceGroup() bool {
	if o != nil && !IsNil(o.TraceGroup) {
		return true
	}

	return false
}

// SetTraceGroup gets a reference to the given AutoTraceGroupConfigType and assigns it to the TraceGroup field.
func (o *AutoTraceCodeDetailType) SetTraceGroup(v AutoTraceGroupConfigType) {
	o.TraceGroup = &v
}

func (o AutoTraceCodeDetailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTraceCodeDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TraceCode) {
		toSerialize["traceCode"] = o.TraceCode
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TraceGroup) {
		toSerialize["traceGroup"] = o.TraceGroup
	}
	return toSerialize, nil
}

type NullableAutoTraceCodeDetailType struct {
	value *AutoTraceCodeDetailType
	isSet bool
}

func (v NullableAutoTraceCodeDetailType) Get() *AutoTraceCodeDetailType {
	return v.value
}

func (v *NullableAutoTraceCodeDetailType) Set(val *AutoTraceCodeDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTraceCodeDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTraceCodeDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTraceCodeDetailType(val *AutoTraceCodeDetailType) *NullableAutoTraceCodeDetailType {
	return &NullableAutoTraceCodeDetailType{value: val, isSet: true}
}

func (v NullableAutoTraceCodeDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTraceCodeDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


