/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ParameterType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterType{}

// ParameterType Name value pair type that will hold generic parameter information. Only use this type when the parameters being passed are too dynamic to be defined.
type ParameterType struct {
	// Name of the parameter.
	ParameterName *string `json:"parameterName,omitempty"`
	// Value of the parameter.
	ParameterValue *string `json:"parameterValue,omitempty"`
}

// NewParameterType instantiates a new ParameterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterType() *ParameterType {
	this := ParameterType{}
	return &this
}

// NewParameterTypeWithDefaults instantiates a new ParameterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterTypeWithDefaults() *ParameterType {
	this := ParameterType{}
	return &this
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *ParameterType) GetParameterName() string {
	if o == nil || IsNil(o.ParameterName) {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterType) GetParameterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterName) {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *ParameterType) HasParameterName() bool {
	if o != nil && !IsNil(o.ParameterName) {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *ParameterType) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterValue returns the ParameterValue field value if set, zero value otherwise.
func (o *ParameterType) GetParameterValue() string {
	if o == nil || IsNil(o.ParameterValue) {
		var ret string
		return ret
	}
	return *o.ParameterValue
}

// GetParameterValueOk returns a tuple with the ParameterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterType) GetParameterValueOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterValue) {
		return nil, false
	}
	return o.ParameterValue, true
}

// HasParameterValue returns a boolean if a field has been set.
func (o *ParameterType) HasParameterValue() bool {
	if o != nil && !IsNil(o.ParameterValue) {
		return true
	}

	return false
}

// SetParameterValue gets a reference to the given string and assigns it to the ParameterValue field.
func (o *ParameterType) SetParameterValue(v string) {
	o.ParameterValue = &v
}

func (o ParameterType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParameterName) {
		toSerialize["parameterName"] = o.ParameterName
	}
	if !IsNil(o.ParameterValue) {
		toSerialize["parameterValue"] = o.ParameterValue
	}
	return toSerialize, nil
}

type NullableParameterType struct {
	value *ParameterType
	isSet bool
}

func (v NullableParameterType) Get() *ParameterType {
	return v.value
}

func (v *NullableParameterType) Set(val *ParameterType) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterType) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterType(val *ParameterType) *NullableParameterType {
	return &NullableParameterType{value: val, isSet: true}
}

func (v NullableParameterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


