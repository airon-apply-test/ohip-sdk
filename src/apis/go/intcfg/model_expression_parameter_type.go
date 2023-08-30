/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intcfg

import (
	"encoding/json"
)

// checks if the ExpressionParameterType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpressionParameterType{}

// ExpressionParameterType struct for ExpressionParameterType
type ExpressionParameterType struct {
	Parameter []string `json:"parameter,omitempty"`
	FunctionIdOne *float32 `json:"functionIdOne,omitempty"`
	FunctionIdTwo *float32 `json:"functionIdTwo,omitempty"`
}

// NewExpressionParameterType instantiates a new ExpressionParameterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpressionParameterType() *ExpressionParameterType {
	this := ExpressionParameterType{}
	return &this
}

// NewExpressionParameterTypeWithDefaults instantiates a new ExpressionParameterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionParameterTypeWithDefaults() *ExpressionParameterType {
	this := ExpressionParameterType{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ExpressionParameterType) GetParameter() []string {
	if o == nil || IsNil(o.Parameter) {
		var ret []string
		return ret
	}
	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionParameterType) GetParameterOk() ([]string, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ExpressionParameterType) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given []string and assigns it to the Parameter field.
func (o *ExpressionParameterType) SetParameter(v []string) {
	o.Parameter = v
}

// GetFunctionIdOne returns the FunctionIdOne field value if set, zero value otherwise.
func (o *ExpressionParameterType) GetFunctionIdOne() float32 {
	if o == nil || IsNil(o.FunctionIdOne) {
		var ret float32
		return ret
	}
	return *o.FunctionIdOne
}

// GetFunctionIdOneOk returns a tuple with the FunctionIdOne field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionParameterType) GetFunctionIdOneOk() (*float32, bool) {
	if o == nil || IsNil(o.FunctionIdOne) {
		return nil, false
	}
	return o.FunctionIdOne, true
}

// HasFunctionIdOne returns a boolean if a field has been set.
func (o *ExpressionParameterType) HasFunctionIdOne() bool {
	if o != nil && !IsNil(o.FunctionIdOne) {
		return true
	}

	return false
}

// SetFunctionIdOne gets a reference to the given float32 and assigns it to the FunctionIdOne field.
func (o *ExpressionParameterType) SetFunctionIdOne(v float32) {
	o.FunctionIdOne = &v
}

// GetFunctionIdTwo returns the FunctionIdTwo field value if set, zero value otherwise.
func (o *ExpressionParameterType) GetFunctionIdTwo() float32 {
	if o == nil || IsNil(o.FunctionIdTwo) {
		var ret float32
		return ret
	}
	return *o.FunctionIdTwo
}

// GetFunctionIdTwoOk returns a tuple with the FunctionIdTwo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionParameterType) GetFunctionIdTwoOk() (*float32, bool) {
	if o == nil || IsNil(o.FunctionIdTwo) {
		return nil, false
	}
	return o.FunctionIdTwo, true
}

// HasFunctionIdTwo returns a boolean if a field has been set.
func (o *ExpressionParameterType) HasFunctionIdTwo() bool {
	if o != nil && !IsNil(o.FunctionIdTwo) {
		return true
	}

	return false
}

// SetFunctionIdTwo gets a reference to the given float32 and assigns it to the FunctionIdTwo field.
func (o *ExpressionParameterType) SetFunctionIdTwo(v float32) {
	o.FunctionIdTwo = &v
}

func (o ExpressionParameterType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpressionParameterType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.FunctionIdOne) {
		toSerialize["functionIdOne"] = o.FunctionIdOne
	}
	if !IsNil(o.FunctionIdTwo) {
		toSerialize["functionIdTwo"] = o.FunctionIdTwo
	}
	return toSerialize, nil
}

type NullableExpressionParameterType struct {
	value *ExpressionParameterType
	isSet bool
}

func (v NullableExpressionParameterType) Get() *ExpressionParameterType {
	return v.value
}

func (v *NullableExpressionParameterType) Set(val *ExpressionParameterType) {
	v.value = val
	v.isSet = true
}

func (v NullableExpressionParameterType) IsSet() bool {
	return v.isSet
}

func (v *NullableExpressionParameterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpressionParameterType(val *ExpressionParameterType) *NullableExpressionParameterType {
	return &NullableExpressionParameterType{value: val, isSet: true}
}

func (v NullableExpressionParameterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpressionParameterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


