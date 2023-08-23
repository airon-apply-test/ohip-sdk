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

// checks if the ExpressionOperatorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpressionOperatorType{}

// ExpressionOperatorType struct for ExpressionOperatorType
type ExpressionOperatorType struct {
	OperatorName *string `json:"operatorName,omitempty"`
	OperatorDescription *string `json:"operatorDescription,omitempty"`
	ParameterSize *float32 `json:"parameterSize,omitempty"`
}

// NewExpressionOperatorType instantiates a new ExpressionOperatorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpressionOperatorType() *ExpressionOperatorType {
	this := ExpressionOperatorType{}
	return &this
}

// NewExpressionOperatorTypeWithDefaults instantiates a new ExpressionOperatorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionOperatorTypeWithDefaults() *ExpressionOperatorType {
	this := ExpressionOperatorType{}
	return &this
}

// GetOperatorName returns the OperatorName field value if set, zero value otherwise.
func (o *ExpressionOperatorType) GetOperatorName() string {
	if o == nil || IsNil(o.OperatorName) {
		var ret string
		return ret
	}
	return *o.OperatorName
}

// GetOperatorNameOk returns a tuple with the OperatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionOperatorType) GetOperatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.OperatorName) {
		return nil, false
	}
	return o.OperatorName, true
}

// HasOperatorName returns a boolean if a field has been set.
func (o *ExpressionOperatorType) HasOperatorName() bool {
	if o != nil && !IsNil(o.OperatorName) {
		return true
	}

	return false
}

// SetOperatorName gets a reference to the given string and assigns it to the OperatorName field.
func (o *ExpressionOperatorType) SetOperatorName(v string) {
	o.OperatorName = &v
}

// GetOperatorDescription returns the OperatorDescription field value if set, zero value otherwise.
func (o *ExpressionOperatorType) GetOperatorDescription() string {
	if o == nil || IsNil(o.OperatorDescription) {
		var ret string
		return ret
	}
	return *o.OperatorDescription
}

// GetOperatorDescriptionOk returns a tuple with the OperatorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionOperatorType) GetOperatorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.OperatorDescription) {
		return nil, false
	}
	return o.OperatorDescription, true
}

// HasOperatorDescription returns a boolean if a field has been set.
func (o *ExpressionOperatorType) HasOperatorDescription() bool {
	if o != nil && !IsNil(o.OperatorDescription) {
		return true
	}

	return false
}

// SetOperatorDescription gets a reference to the given string and assigns it to the OperatorDescription field.
func (o *ExpressionOperatorType) SetOperatorDescription(v string) {
	o.OperatorDescription = &v
}

// GetParameterSize returns the ParameterSize field value if set, zero value otherwise.
func (o *ExpressionOperatorType) GetParameterSize() float32 {
	if o == nil || IsNil(o.ParameterSize) {
		var ret float32
		return ret
	}
	return *o.ParameterSize
}

// GetParameterSizeOk returns a tuple with the ParameterSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionOperatorType) GetParameterSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.ParameterSize) {
		return nil, false
	}
	return o.ParameterSize, true
}

// HasParameterSize returns a boolean if a field has been set.
func (o *ExpressionOperatorType) HasParameterSize() bool {
	if o != nil && !IsNil(o.ParameterSize) {
		return true
	}

	return false
}

// SetParameterSize gets a reference to the given float32 and assigns it to the ParameterSize field.
func (o *ExpressionOperatorType) SetParameterSize(v float32) {
	o.ParameterSize = &v
}

func (o ExpressionOperatorType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpressionOperatorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperatorName) {
		toSerialize["operatorName"] = o.OperatorName
	}
	if !IsNil(o.OperatorDescription) {
		toSerialize["operatorDescription"] = o.OperatorDescription
	}
	if !IsNil(o.ParameterSize) {
		toSerialize["parameterSize"] = o.ParameterSize
	}
	return toSerialize, nil
}

type NullableExpressionOperatorType struct {
	value *ExpressionOperatorType
	isSet bool
}

func (v NullableExpressionOperatorType) Get() *ExpressionOperatorType {
	return v.value
}

func (v *NullableExpressionOperatorType) Set(val *ExpressionOperatorType) {
	v.value = val
	v.isSet = true
}

func (v NullableExpressionOperatorType) IsSet() bool {
	return v.isSet
}

func (v *NullableExpressionOperatorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpressionOperatorType(val *ExpressionOperatorType) *NullableExpressionOperatorType {
	return &NullableExpressionOperatorType{value: val, isSet: true}
}

func (v NullableExpressionOperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpressionOperatorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


