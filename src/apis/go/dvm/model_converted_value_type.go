/*
OPERA Cloud DataValueMapping Service API

APIs which offer external systems to config and use values different than what are configured in opera<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dvm

import (
	"encoding/json"
)

// checks if the ConvertedValueType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertedValueType{}

// ConvertedValueType Details of the converted value.
type ConvertedValueType struct {
	// Opera Value
	OperaValue *string `json:"operaValue,omitempty"`
	// Value used by the external vendors.
	ExternalValue *string `json:"externalValue,omitempty"`
	ConversionCode *DataValueMappingCodeType `json:"conversionCode,omitempty"`
	// Opera Master Value
	MasterValue *string `json:"masterValue,omitempty"`
	// The flag will be true for all the conversions that are pms defaults. In case there are more than one conversions available in opera.
	PmsDefaultConversion *bool `json:"pmsDefaultConversion,omitempty"`
}

// NewConvertedValueType instantiates a new ConvertedValueType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertedValueType() *ConvertedValueType {
	this := ConvertedValueType{}
	return &this
}

// NewConvertedValueTypeWithDefaults instantiates a new ConvertedValueType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertedValueTypeWithDefaults() *ConvertedValueType {
	this := ConvertedValueType{}
	return &this
}

// GetOperaValue returns the OperaValue field value if set, zero value otherwise.
func (o *ConvertedValueType) GetOperaValue() string {
	if o == nil || IsNil(o.OperaValue) {
		var ret string
		return ret
	}
	return *o.OperaValue
}

// GetOperaValueOk returns a tuple with the OperaValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedValueType) GetOperaValueOk() (*string, bool) {
	if o == nil || IsNil(o.OperaValue) {
		return nil, false
	}
	return o.OperaValue, true
}

// HasOperaValue returns a boolean if a field has been set.
func (o *ConvertedValueType) HasOperaValue() bool {
	if o != nil && !IsNil(o.OperaValue) {
		return true
	}

	return false
}

// SetOperaValue gets a reference to the given string and assigns it to the OperaValue field.
func (o *ConvertedValueType) SetOperaValue(v string) {
	o.OperaValue = &v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *ConvertedValueType) GetExternalValue() string {
	if o == nil || IsNil(o.ExternalValue) {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedValueType) GetExternalValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalValue) {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *ConvertedValueType) HasExternalValue() bool {
	if o != nil && !IsNil(o.ExternalValue) {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *ConvertedValueType) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetConversionCode returns the ConversionCode field value if set, zero value otherwise.
func (o *ConvertedValueType) GetConversionCode() DataValueMappingCodeType {
	if o == nil || IsNil(o.ConversionCode) {
		var ret DataValueMappingCodeType
		return ret
	}
	return *o.ConversionCode
}

// GetConversionCodeOk returns a tuple with the ConversionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedValueType) GetConversionCodeOk() (*DataValueMappingCodeType, bool) {
	if o == nil || IsNil(o.ConversionCode) {
		return nil, false
	}
	return o.ConversionCode, true
}

// HasConversionCode returns a boolean if a field has been set.
func (o *ConvertedValueType) HasConversionCode() bool {
	if o != nil && !IsNil(o.ConversionCode) {
		return true
	}

	return false
}

// SetConversionCode gets a reference to the given DataValueMappingCodeType and assigns it to the ConversionCode field.
func (o *ConvertedValueType) SetConversionCode(v DataValueMappingCodeType) {
	o.ConversionCode = &v
}

// GetMasterValue returns the MasterValue field value if set, zero value otherwise.
func (o *ConvertedValueType) GetMasterValue() string {
	if o == nil || IsNil(o.MasterValue) {
		var ret string
		return ret
	}
	return *o.MasterValue
}

// GetMasterValueOk returns a tuple with the MasterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedValueType) GetMasterValueOk() (*string, bool) {
	if o == nil || IsNil(o.MasterValue) {
		return nil, false
	}
	return o.MasterValue, true
}

// HasMasterValue returns a boolean if a field has been set.
func (o *ConvertedValueType) HasMasterValue() bool {
	if o != nil && !IsNil(o.MasterValue) {
		return true
	}

	return false
}

// SetMasterValue gets a reference to the given string and assigns it to the MasterValue field.
func (o *ConvertedValueType) SetMasterValue(v string) {
	o.MasterValue = &v
}

// GetPmsDefaultConversion returns the PmsDefaultConversion field value if set, zero value otherwise.
func (o *ConvertedValueType) GetPmsDefaultConversion() bool {
	if o == nil || IsNil(o.PmsDefaultConversion) {
		var ret bool
		return ret
	}
	return *o.PmsDefaultConversion
}

// GetPmsDefaultConversionOk returns a tuple with the PmsDefaultConversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedValueType) GetPmsDefaultConversionOk() (*bool, bool) {
	if o == nil || IsNil(o.PmsDefaultConversion) {
		return nil, false
	}
	return o.PmsDefaultConversion, true
}

// HasPmsDefaultConversion returns a boolean if a field has been set.
func (o *ConvertedValueType) HasPmsDefaultConversion() bool {
	if o != nil && !IsNil(o.PmsDefaultConversion) {
		return true
	}

	return false
}

// SetPmsDefaultConversion gets a reference to the given bool and assigns it to the PmsDefaultConversion field.
func (o *ConvertedValueType) SetPmsDefaultConversion(v bool) {
	o.PmsDefaultConversion = &v
}

func (o ConvertedValueType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertedValueType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperaValue) {
		toSerialize["operaValue"] = o.OperaValue
	}
	if !IsNil(o.ExternalValue) {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if !IsNil(o.ConversionCode) {
		toSerialize["conversionCode"] = o.ConversionCode
	}
	if !IsNil(o.MasterValue) {
		toSerialize["masterValue"] = o.MasterValue
	}
	if !IsNil(o.PmsDefaultConversion) {
		toSerialize["pmsDefaultConversion"] = o.PmsDefaultConversion
	}
	return toSerialize, nil
}

type NullableConvertedValueType struct {
	value *ConvertedValueType
	isSet bool
}

func (v NullableConvertedValueType) Get() *ConvertedValueType {
	return v.value
}

func (v *NullableConvertedValueType) Set(val *ConvertedValueType) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertedValueType) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertedValueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertedValueType(val *ConvertedValueType) *NullableConvertedValueType {
	return &NullableConvertedValueType{value: val, isSet: true}
}

func (v NullableConvertedValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertedValueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


