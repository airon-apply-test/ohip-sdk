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

// checks if the ConditionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionType{}

// ConditionType struct for ConditionType
type ConditionType struct {
	LeftExpression *ExpressionFieldType `json:"leftExpression,omitempty"`
	Operator *ExpressionOperatorType `json:"operator,omitempty"`
	RightExpression *ExpressionParameterType `json:"rightExpression,omitempty"`
	LogicalOperator *LogicalOperatorType `json:"logicalOperator,omitempty"`
}

// NewConditionType instantiates a new ConditionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionType() *ConditionType {
	this := ConditionType{}
	return &this
}

// NewConditionTypeWithDefaults instantiates a new ConditionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionTypeWithDefaults() *ConditionType {
	this := ConditionType{}
	return &this
}

// GetLeftExpression returns the LeftExpression field value if set, zero value otherwise.
func (o *ConditionType) GetLeftExpression() ExpressionFieldType {
	if o == nil || IsNil(o.LeftExpression) {
		var ret ExpressionFieldType
		return ret
	}
	return *o.LeftExpression
}

// GetLeftExpressionOk returns a tuple with the LeftExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionType) GetLeftExpressionOk() (*ExpressionFieldType, bool) {
	if o == nil || IsNil(o.LeftExpression) {
		return nil, false
	}
	return o.LeftExpression, true
}

// HasLeftExpression returns a boolean if a field has been set.
func (o *ConditionType) HasLeftExpression() bool {
	if o != nil && !IsNil(o.LeftExpression) {
		return true
	}

	return false
}

// SetLeftExpression gets a reference to the given ExpressionFieldType and assigns it to the LeftExpression field.
func (o *ConditionType) SetLeftExpression(v ExpressionFieldType) {
	o.LeftExpression = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ConditionType) GetOperator() ExpressionOperatorType {
	if o == nil || IsNil(o.Operator) {
		var ret ExpressionOperatorType
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionType) GetOperatorOk() (*ExpressionOperatorType, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ConditionType) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given ExpressionOperatorType and assigns it to the Operator field.
func (o *ConditionType) SetOperator(v ExpressionOperatorType) {
	o.Operator = &v
}

// GetRightExpression returns the RightExpression field value if set, zero value otherwise.
func (o *ConditionType) GetRightExpression() ExpressionParameterType {
	if o == nil || IsNil(o.RightExpression) {
		var ret ExpressionParameterType
		return ret
	}
	return *o.RightExpression
}

// GetRightExpressionOk returns a tuple with the RightExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionType) GetRightExpressionOk() (*ExpressionParameterType, bool) {
	if o == nil || IsNil(o.RightExpression) {
		return nil, false
	}
	return o.RightExpression, true
}

// HasRightExpression returns a boolean if a field has been set.
func (o *ConditionType) HasRightExpression() bool {
	if o != nil && !IsNil(o.RightExpression) {
		return true
	}

	return false
}

// SetRightExpression gets a reference to the given ExpressionParameterType and assigns it to the RightExpression field.
func (o *ConditionType) SetRightExpression(v ExpressionParameterType) {
	o.RightExpression = &v
}

// GetLogicalOperator returns the LogicalOperator field value if set, zero value otherwise.
func (o *ConditionType) GetLogicalOperator() LogicalOperatorType {
	if o == nil || IsNil(o.LogicalOperator) {
		var ret LogicalOperatorType
		return ret
	}
	return *o.LogicalOperator
}

// GetLogicalOperatorOk returns a tuple with the LogicalOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionType) GetLogicalOperatorOk() (*LogicalOperatorType, bool) {
	if o == nil || IsNil(o.LogicalOperator) {
		return nil, false
	}
	return o.LogicalOperator, true
}

// HasLogicalOperator returns a boolean if a field has been set.
func (o *ConditionType) HasLogicalOperator() bool {
	if o != nil && !IsNil(o.LogicalOperator) {
		return true
	}

	return false
}

// SetLogicalOperator gets a reference to the given LogicalOperatorType and assigns it to the LogicalOperator field.
func (o *ConditionType) SetLogicalOperator(v LogicalOperatorType) {
	o.LogicalOperator = &v
}

func (o ConditionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LeftExpression) {
		toSerialize["leftExpression"] = o.LeftExpression
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.RightExpression) {
		toSerialize["rightExpression"] = o.RightExpression
	}
	if !IsNil(o.LogicalOperator) {
		toSerialize["logicalOperator"] = o.LogicalOperator
	}
	return toSerialize, nil
}

type NullableConditionType struct {
	value *ConditionType
	isSet bool
}

func (v NullableConditionType) Get() *ConditionType {
	return v.value
}

func (v *NullableConditionType) Set(val *ConditionType) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionType) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionType(val *ConditionType) *NullableConditionType {
	return &NullableConditionType{value: val, isSet: true}
}

func (v NullableConditionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


