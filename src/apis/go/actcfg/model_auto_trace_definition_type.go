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

// checks if the AutoTraceDefinitionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTraceDefinitionType{}

// AutoTraceDefinitionType Auto Trace Definition information.
type AutoTraceDefinitionType struct {
	AutoTraceDefinitionDetail *AutoTraceDefinitionDetailType `json:"autoTraceDefinitionDetail,omitempty"`
	AutoTraceDefinitionActivityInfo *AutoTraceDefinitionActivityInfoType `json:"autoTraceDefinitionActivityInfo,omitempty"`
	AutoTraceDefinitionOwnerInfo *AutoTraceDefinitionOwnerInfoType `json:"autoTraceDefinitionOwnerInfo,omitempty"`
	// Determines whether to fetch inactive records or not.
	Inactive *bool `json:"inactive,omitempty"`
}

// NewAutoTraceDefinitionType instantiates a new AutoTraceDefinitionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTraceDefinitionType() *AutoTraceDefinitionType {
	this := AutoTraceDefinitionType{}
	return &this
}

// NewAutoTraceDefinitionTypeWithDefaults instantiates a new AutoTraceDefinitionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTraceDefinitionTypeWithDefaults() *AutoTraceDefinitionType {
	this := AutoTraceDefinitionType{}
	return &this
}

// GetAutoTraceDefinitionDetail returns the AutoTraceDefinitionDetail field value if set, zero value otherwise.
func (o *AutoTraceDefinitionType) GetAutoTraceDefinitionDetail() AutoTraceDefinitionDetailType {
	if o == nil || IsNil(o.AutoTraceDefinitionDetail) {
		var ret AutoTraceDefinitionDetailType
		return ret
	}
	return *o.AutoTraceDefinitionDetail
}

// GetAutoTraceDefinitionDetailOk returns a tuple with the AutoTraceDefinitionDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionType) GetAutoTraceDefinitionDetailOk() (*AutoTraceDefinitionDetailType, bool) {
	if o == nil || IsNil(o.AutoTraceDefinitionDetail) {
		return nil, false
	}
	return o.AutoTraceDefinitionDetail, true
}

// HasAutoTraceDefinitionDetail returns a boolean if a field has been set.
func (o *AutoTraceDefinitionType) HasAutoTraceDefinitionDetail() bool {
	if o != nil && !IsNil(o.AutoTraceDefinitionDetail) {
		return true
	}

	return false
}

// SetAutoTraceDefinitionDetail gets a reference to the given AutoTraceDefinitionDetailType and assigns it to the AutoTraceDefinitionDetail field.
func (o *AutoTraceDefinitionType) SetAutoTraceDefinitionDetail(v AutoTraceDefinitionDetailType) {
	o.AutoTraceDefinitionDetail = &v
}

// GetAutoTraceDefinitionActivityInfo returns the AutoTraceDefinitionActivityInfo field value if set, zero value otherwise.
func (o *AutoTraceDefinitionType) GetAutoTraceDefinitionActivityInfo() AutoTraceDefinitionActivityInfoType {
	if o == nil || IsNil(o.AutoTraceDefinitionActivityInfo) {
		var ret AutoTraceDefinitionActivityInfoType
		return ret
	}
	return *o.AutoTraceDefinitionActivityInfo
}

// GetAutoTraceDefinitionActivityInfoOk returns a tuple with the AutoTraceDefinitionActivityInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionType) GetAutoTraceDefinitionActivityInfoOk() (*AutoTraceDefinitionActivityInfoType, bool) {
	if o == nil || IsNil(o.AutoTraceDefinitionActivityInfo) {
		return nil, false
	}
	return o.AutoTraceDefinitionActivityInfo, true
}

// HasAutoTraceDefinitionActivityInfo returns a boolean if a field has been set.
func (o *AutoTraceDefinitionType) HasAutoTraceDefinitionActivityInfo() bool {
	if o != nil && !IsNil(o.AutoTraceDefinitionActivityInfo) {
		return true
	}

	return false
}

// SetAutoTraceDefinitionActivityInfo gets a reference to the given AutoTraceDefinitionActivityInfoType and assigns it to the AutoTraceDefinitionActivityInfo field.
func (o *AutoTraceDefinitionType) SetAutoTraceDefinitionActivityInfo(v AutoTraceDefinitionActivityInfoType) {
	o.AutoTraceDefinitionActivityInfo = &v
}

// GetAutoTraceDefinitionOwnerInfo returns the AutoTraceDefinitionOwnerInfo field value if set, zero value otherwise.
func (o *AutoTraceDefinitionType) GetAutoTraceDefinitionOwnerInfo() AutoTraceDefinitionOwnerInfoType {
	if o == nil || IsNil(o.AutoTraceDefinitionOwnerInfo) {
		var ret AutoTraceDefinitionOwnerInfoType
		return ret
	}
	return *o.AutoTraceDefinitionOwnerInfo
}

// GetAutoTraceDefinitionOwnerInfoOk returns a tuple with the AutoTraceDefinitionOwnerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionType) GetAutoTraceDefinitionOwnerInfoOk() (*AutoTraceDefinitionOwnerInfoType, bool) {
	if o == nil || IsNil(o.AutoTraceDefinitionOwnerInfo) {
		return nil, false
	}
	return o.AutoTraceDefinitionOwnerInfo, true
}

// HasAutoTraceDefinitionOwnerInfo returns a boolean if a field has been set.
func (o *AutoTraceDefinitionType) HasAutoTraceDefinitionOwnerInfo() bool {
	if o != nil && !IsNil(o.AutoTraceDefinitionOwnerInfo) {
		return true
	}

	return false
}

// SetAutoTraceDefinitionOwnerInfo gets a reference to the given AutoTraceDefinitionOwnerInfoType and assigns it to the AutoTraceDefinitionOwnerInfo field.
func (o *AutoTraceDefinitionType) SetAutoTraceDefinitionOwnerInfo(v AutoTraceDefinitionOwnerInfoType) {
	o.AutoTraceDefinitionOwnerInfo = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *AutoTraceDefinitionType) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionType) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *AutoTraceDefinitionType) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *AutoTraceDefinitionType) SetInactive(v bool) {
	o.Inactive = &v
}

func (o AutoTraceDefinitionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTraceDefinitionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoTraceDefinitionDetail) {
		toSerialize["autoTraceDefinitionDetail"] = o.AutoTraceDefinitionDetail
	}
	if !IsNil(o.AutoTraceDefinitionActivityInfo) {
		toSerialize["autoTraceDefinitionActivityInfo"] = o.AutoTraceDefinitionActivityInfo
	}
	if !IsNil(o.AutoTraceDefinitionOwnerInfo) {
		toSerialize["autoTraceDefinitionOwnerInfo"] = o.AutoTraceDefinitionOwnerInfo
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	return toSerialize, nil
}

type NullableAutoTraceDefinitionType struct {
	value *AutoTraceDefinitionType
	isSet bool
}

func (v NullableAutoTraceDefinitionType) Get() *AutoTraceDefinitionType {
	return v.value
}

func (v *NullableAutoTraceDefinitionType) Set(val *AutoTraceDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTraceDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTraceDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTraceDefinitionType(val *AutoTraceDefinitionType) *NullableAutoTraceDefinitionType {
	return &NullableAutoTraceDefinitionType{value: val, isSet: true}
}

func (v NullableAutoTraceDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTraceDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


