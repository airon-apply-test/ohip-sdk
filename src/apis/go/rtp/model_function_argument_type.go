/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FunctionArgumentType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionArgumentType{}

// FunctionArgumentType This type provided details of a function argument and the value it holds.
type FunctionArgumentType struct {
	// Specifies the name of the function argument.
	Name *string `json:"name,omitempty"`
	// Specifies the value held by the function argument.
	Value *string `json:"value,omitempty"`
	// Specifies the position of the function argument in the argument list.
	Position *int32 `json:"position,omitempty"`
	// Specifies the datatype of the function argument.
	DataType *string `json:"dataType,omitempty"`
	// Argument of the function mandatory or not.
	Required *bool `json:"required,omitempty"`
}

// NewFunctionArgumentType instantiates a new FunctionArgumentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionArgumentType() *FunctionArgumentType {
	this := FunctionArgumentType{}
	return &this
}

// NewFunctionArgumentTypeWithDefaults instantiates a new FunctionArgumentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionArgumentTypeWithDefaults() *FunctionArgumentType {
	this := FunctionArgumentType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FunctionArgumentType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionArgumentType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FunctionArgumentType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FunctionArgumentType) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FunctionArgumentType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionArgumentType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FunctionArgumentType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FunctionArgumentType) SetValue(v string) {
	o.Value = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *FunctionArgumentType) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionArgumentType) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *FunctionArgumentType) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *FunctionArgumentType) SetPosition(v int32) {
	o.Position = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *FunctionArgumentType) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionArgumentType) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *FunctionArgumentType) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *FunctionArgumentType) SetDataType(v string) {
	o.DataType = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FunctionArgumentType) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionArgumentType) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FunctionArgumentType) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FunctionArgumentType) SetRequired(v bool) {
	o.Required = &v
}

func (o FunctionArgumentType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionArgumentType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableFunctionArgumentType struct {
	value *FunctionArgumentType
	isSet bool
}

func (v NullableFunctionArgumentType) Get() *FunctionArgumentType {
	return v.value
}

func (v *NullableFunctionArgumentType) Set(val *FunctionArgumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionArgumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionArgumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionArgumentType(val *FunctionArgumentType) *NullableFunctionArgumentType {
	return &NullableFunctionArgumentType{value: val, isSet: true}
}

func (v NullableFunctionArgumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionArgumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


