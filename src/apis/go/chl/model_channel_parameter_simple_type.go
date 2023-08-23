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

// checks if the ChannelParameterSimpleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelParameterSimpleType{}

// ChannelParameterSimpleType Basic channel parameter info used in changing channel parameter values.
type ChannelParameterSimpleType struct {
	ParameterName *string `json:"parameterName,omitempty"`
	ParameterValue *string `json:"parameterValue,omitempty"`
	HotelId *string `json:"hotelId,omitempty"`
	ParameterValueType *ChannelParameterValueType `json:"parameterValueType,omitempty"`
}

// NewChannelParameterSimpleType instantiates a new ChannelParameterSimpleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelParameterSimpleType() *ChannelParameterSimpleType {
	this := ChannelParameterSimpleType{}
	return &this
}

// NewChannelParameterSimpleTypeWithDefaults instantiates a new ChannelParameterSimpleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelParameterSimpleTypeWithDefaults() *ChannelParameterSimpleType {
	this := ChannelParameterSimpleType{}
	return &this
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *ChannelParameterSimpleType) GetParameterName() string {
	if o == nil || IsNil(o.ParameterName) {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterSimpleType) GetParameterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterName) {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *ChannelParameterSimpleType) HasParameterName() bool {
	if o != nil && !IsNil(o.ParameterName) {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *ChannelParameterSimpleType) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterValue returns the ParameterValue field value if set, zero value otherwise.
func (o *ChannelParameterSimpleType) GetParameterValue() string {
	if o == nil || IsNil(o.ParameterValue) {
		var ret string
		return ret
	}
	return *o.ParameterValue
}

// GetParameterValueOk returns a tuple with the ParameterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterSimpleType) GetParameterValueOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterValue) {
		return nil, false
	}
	return o.ParameterValue, true
}

// HasParameterValue returns a boolean if a field has been set.
func (o *ChannelParameterSimpleType) HasParameterValue() bool {
	if o != nil && !IsNil(o.ParameterValue) {
		return true
	}

	return false
}

// SetParameterValue gets a reference to the given string and assigns it to the ParameterValue field.
func (o *ChannelParameterSimpleType) SetParameterValue(v string) {
	o.ParameterValue = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ChannelParameterSimpleType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterSimpleType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ChannelParameterSimpleType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ChannelParameterSimpleType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetParameterValueType returns the ParameterValueType field value if set, zero value otherwise.
func (o *ChannelParameterSimpleType) GetParameterValueType() ChannelParameterValueType {
	if o == nil || IsNil(o.ParameterValueType) {
		var ret ChannelParameterValueType
		return ret
	}
	return *o.ParameterValueType
}

// GetParameterValueTypeOk returns a tuple with the ParameterValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterSimpleType) GetParameterValueTypeOk() (*ChannelParameterValueType, bool) {
	if o == nil || IsNil(o.ParameterValueType) {
		return nil, false
	}
	return o.ParameterValueType, true
}

// HasParameterValueType returns a boolean if a field has been set.
func (o *ChannelParameterSimpleType) HasParameterValueType() bool {
	if o != nil && !IsNil(o.ParameterValueType) {
		return true
	}

	return false
}

// SetParameterValueType gets a reference to the given ChannelParameterValueType and assigns it to the ParameterValueType field.
func (o *ChannelParameterSimpleType) SetParameterValueType(v ChannelParameterValueType) {
	o.ParameterValueType = &v
}

func (o ChannelParameterSimpleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelParameterSimpleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParameterName) {
		toSerialize["parameterName"] = o.ParameterName
	}
	if !IsNil(o.ParameterValue) {
		toSerialize["parameterValue"] = o.ParameterValue
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ParameterValueType) {
		toSerialize["parameterValueType"] = o.ParameterValueType
	}
	return toSerialize, nil
}

type NullableChannelParameterSimpleType struct {
	value *ChannelParameterSimpleType
	isSet bool
}

func (v NullableChannelParameterSimpleType) Get() *ChannelParameterSimpleType {
	return v.value
}

func (v *NullableChannelParameterSimpleType) Set(val *ChannelParameterSimpleType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelParameterSimpleType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelParameterSimpleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelParameterSimpleType(val *ChannelParameterSimpleType) *NullableChannelParameterSimpleType {
	return &NullableChannelParameterSimpleType{value: val, isSet: true}
}

func (v NullableChannelParameterSimpleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelParameterSimpleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


