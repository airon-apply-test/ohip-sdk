/*
OPERA Cloud DataValueMapping Service API

APIs which offer external systems to config and use values different than what are configured in opera<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ConversionHotelInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConversionHotelInfoType{}

// ConversionHotelInfoType Generic information of the resort , external resort and external system.
type ConversionHotelInfoType struct {
	// Opera Resort for which conversion is needed.
	OperaHotelCode *string `json:"operaHotelCode,omitempty"`
	// External value of resort for which conversion is needed.
	ExternalHotelCode *string `json:"externalHotelCode,omitempty"`
}

// NewConversionHotelInfoType instantiates a new ConversionHotelInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversionHotelInfoType() *ConversionHotelInfoType {
	this := ConversionHotelInfoType{}
	return &this
}

// NewConversionHotelInfoTypeWithDefaults instantiates a new ConversionHotelInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversionHotelInfoTypeWithDefaults() *ConversionHotelInfoType {
	this := ConversionHotelInfoType{}
	return &this
}

// GetOperaHotelCode returns the OperaHotelCode field value if set, zero value otherwise.
func (o *ConversionHotelInfoType) GetOperaHotelCode() string {
	if o == nil || IsNil(o.OperaHotelCode) {
		var ret string
		return ret
	}
	return *o.OperaHotelCode
}

// GetOperaHotelCodeOk returns a tuple with the OperaHotelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionHotelInfoType) GetOperaHotelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OperaHotelCode) {
		return nil, false
	}
	return o.OperaHotelCode, true
}

// HasOperaHotelCode returns a boolean if a field has been set.
func (o *ConversionHotelInfoType) HasOperaHotelCode() bool {
	if o != nil && !IsNil(o.OperaHotelCode) {
		return true
	}

	return false
}

// SetOperaHotelCode gets a reference to the given string and assigns it to the OperaHotelCode field.
func (o *ConversionHotelInfoType) SetOperaHotelCode(v string) {
	o.OperaHotelCode = &v
}

// GetExternalHotelCode returns the ExternalHotelCode field value if set, zero value otherwise.
func (o *ConversionHotelInfoType) GetExternalHotelCode() string {
	if o == nil || IsNil(o.ExternalHotelCode) {
		var ret string
		return ret
	}
	return *o.ExternalHotelCode
}

// GetExternalHotelCodeOk returns a tuple with the ExternalHotelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionHotelInfoType) GetExternalHotelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalHotelCode) {
		return nil, false
	}
	return o.ExternalHotelCode, true
}

// HasExternalHotelCode returns a boolean if a field has been set.
func (o *ConversionHotelInfoType) HasExternalHotelCode() bool {
	if o != nil && !IsNil(o.ExternalHotelCode) {
		return true
	}

	return false
}

// SetExternalHotelCode gets a reference to the given string and assigns it to the ExternalHotelCode field.
func (o *ConversionHotelInfoType) SetExternalHotelCode(v string) {
	o.ExternalHotelCode = &v
}

func (o ConversionHotelInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConversionHotelInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperaHotelCode) {
		toSerialize["operaHotelCode"] = o.OperaHotelCode
	}
	if !IsNil(o.ExternalHotelCode) {
		toSerialize["externalHotelCode"] = o.ExternalHotelCode
	}
	return toSerialize, nil
}

type NullableConversionHotelInfoType struct {
	value *ConversionHotelInfoType
	isSet bool
}

func (v NullableConversionHotelInfoType) Get() *ConversionHotelInfoType {
	return v.value
}

func (v *NullableConversionHotelInfoType) Set(val *ConversionHotelInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableConversionHotelInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableConversionHotelInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversionHotelInfoType(val *ConversionHotelInfoType) *NullableConversionHotelInfoType {
	return &NullableConversionHotelInfoType{value: val, isSet: true}
}

func (v NullableConversionHotelInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversionHotelInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


