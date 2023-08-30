/*
OPERA Cloud List of Values Management API

APIs to cater for List of Value functionality in OPERA Cloud. A List of Values in the OPERA Application can be configured by a property.  Then by using these APIs you can retrieve all configured codes.  As an example, Titles is a configurable ListOfValues.  A hotel can specify what titles they wish to use, and thus fetching the LOV for title, you can view the codes that are configured for a property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lov

import (
	"encoding/json"
)

// checks if the ListOfValuesCriteriaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOfValuesCriteriaType{}

// ListOfValuesCriteriaType struct for ListOfValuesCriteriaType
type ListOfValuesCriteriaType struct {
	// Collection of generic Name-Value-Pair parameters.
	Parameters []ParameterType `json:"parameters,omitempty"`
	ItemCodes []string `json:"itemCodes,omitempty"`
	ExcludeCodeList []string `json:"excludeCodeList,omitempty"`
	// The name of the LOV to fetch.
	LovName *string `json:"lovName,omitempty"`
	// Only useful for LOVs that support toggle of inactive records inclusion. When set to true, inactive records will be included.
	IncludeInactive *bool `json:"includeInactive,omitempty"`
}

// NewListOfValuesCriteriaType instantiates a new ListOfValuesCriteriaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOfValuesCriteriaType() *ListOfValuesCriteriaType {
	this := ListOfValuesCriteriaType{}
	return &this
}

// NewListOfValuesCriteriaTypeWithDefaults instantiates a new ListOfValuesCriteriaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOfValuesCriteriaTypeWithDefaults() *ListOfValuesCriteriaType {
	this := ListOfValuesCriteriaType{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ListOfValuesCriteriaType) GetParameters() []ParameterType {
	if o == nil || IsNil(o.Parameters) {
		var ret []ParameterType
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOfValuesCriteriaType) GetParametersOk() ([]ParameterType, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ListOfValuesCriteriaType) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ParameterType and assigns it to the Parameters field.
func (o *ListOfValuesCriteriaType) SetParameters(v []ParameterType) {
	o.Parameters = v
}

// GetItemCodes returns the ItemCodes field value if set, zero value otherwise.
func (o *ListOfValuesCriteriaType) GetItemCodes() []string {
	if o == nil || IsNil(o.ItemCodes) {
		var ret []string
		return ret
	}
	return o.ItemCodes
}

// GetItemCodesOk returns a tuple with the ItemCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOfValuesCriteriaType) GetItemCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemCodes) {
		return nil, false
	}
	return o.ItemCodes, true
}

// HasItemCodes returns a boolean if a field has been set.
func (o *ListOfValuesCriteriaType) HasItemCodes() bool {
	if o != nil && !IsNil(o.ItemCodes) {
		return true
	}

	return false
}

// SetItemCodes gets a reference to the given []string and assigns it to the ItemCodes field.
func (o *ListOfValuesCriteriaType) SetItemCodes(v []string) {
	o.ItemCodes = v
}

// GetExcludeCodeList returns the ExcludeCodeList field value if set, zero value otherwise.
func (o *ListOfValuesCriteriaType) GetExcludeCodeList() []string {
	if o == nil || IsNil(o.ExcludeCodeList) {
		var ret []string
		return ret
	}
	return o.ExcludeCodeList
}

// GetExcludeCodeListOk returns a tuple with the ExcludeCodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOfValuesCriteriaType) GetExcludeCodeListOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeCodeList) {
		return nil, false
	}
	return o.ExcludeCodeList, true
}

// HasExcludeCodeList returns a boolean if a field has been set.
func (o *ListOfValuesCriteriaType) HasExcludeCodeList() bool {
	if o != nil && !IsNil(o.ExcludeCodeList) {
		return true
	}

	return false
}

// SetExcludeCodeList gets a reference to the given []string and assigns it to the ExcludeCodeList field.
func (o *ListOfValuesCriteriaType) SetExcludeCodeList(v []string) {
	o.ExcludeCodeList = v
}

// GetLovName returns the LovName field value if set, zero value otherwise.
func (o *ListOfValuesCriteriaType) GetLovName() string {
	if o == nil || IsNil(o.LovName) {
		var ret string
		return ret
	}
	return *o.LovName
}

// GetLovNameOk returns a tuple with the LovName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOfValuesCriteriaType) GetLovNameOk() (*string, bool) {
	if o == nil || IsNil(o.LovName) {
		return nil, false
	}
	return o.LovName, true
}

// HasLovName returns a boolean if a field has been set.
func (o *ListOfValuesCriteriaType) HasLovName() bool {
	if o != nil && !IsNil(o.LovName) {
		return true
	}

	return false
}

// SetLovName gets a reference to the given string and assigns it to the LovName field.
func (o *ListOfValuesCriteriaType) SetLovName(v string) {
	o.LovName = &v
}

// GetIncludeInactive returns the IncludeInactive field value if set, zero value otherwise.
func (o *ListOfValuesCriteriaType) GetIncludeInactive() bool {
	if o == nil || IsNil(o.IncludeInactive) {
		var ret bool
		return ret
	}
	return *o.IncludeInactive
}

// GetIncludeInactiveOk returns a tuple with the IncludeInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOfValuesCriteriaType) GetIncludeInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeInactive) {
		return nil, false
	}
	return o.IncludeInactive, true
}

// HasIncludeInactive returns a boolean if a field has been set.
func (o *ListOfValuesCriteriaType) HasIncludeInactive() bool {
	if o != nil && !IsNil(o.IncludeInactive) {
		return true
	}

	return false
}

// SetIncludeInactive gets a reference to the given bool and assigns it to the IncludeInactive field.
func (o *ListOfValuesCriteriaType) SetIncludeInactive(v bool) {
	o.IncludeInactive = &v
}

func (o ListOfValuesCriteriaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOfValuesCriteriaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.ItemCodes) {
		toSerialize["itemCodes"] = o.ItemCodes
	}
	if !IsNil(o.ExcludeCodeList) {
		toSerialize["excludeCodeList"] = o.ExcludeCodeList
	}
	if !IsNil(o.LovName) {
		toSerialize["lovName"] = o.LovName
	}
	if !IsNil(o.IncludeInactive) {
		toSerialize["includeInactive"] = o.IncludeInactive
	}
	return toSerialize, nil
}

type NullableListOfValuesCriteriaType struct {
	value *ListOfValuesCriteriaType
	isSet bool
}

func (v NullableListOfValuesCriteriaType) Get() *ListOfValuesCriteriaType {
	return v.value
}

func (v *NullableListOfValuesCriteriaType) Set(val *ListOfValuesCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableListOfValuesCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableListOfValuesCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOfValuesCriteriaType(val *ListOfValuesCriteriaType) *NullableListOfValuesCriteriaType {
	return &NullableListOfValuesCriteriaType{value: val, isSet: true}
}

func (v NullableListOfValuesCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOfValuesCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


