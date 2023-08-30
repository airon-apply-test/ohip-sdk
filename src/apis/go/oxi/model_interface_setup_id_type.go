/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the InterfaceSetupIDType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceSetupIDType{}

// InterfaceSetupIDType Type represents ID one Interface Setup
type InterfaceSetupIDType struct {
	// ID of the Interface Setup
	InterfaceId *string `json:"interfaceId,omitempty"`
	// Property for which the Interface Setup is defined.
	HotelId *string `json:"hotelId,omitempty"`
}

// NewInterfaceSetupIDType instantiates a new InterfaceSetupIDType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceSetupIDType() *InterfaceSetupIDType {
	this := InterfaceSetupIDType{}
	return &this
}

// NewInterfaceSetupIDTypeWithDefaults instantiates a new InterfaceSetupIDType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceSetupIDTypeWithDefaults() *InterfaceSetupIDType {
	this := InterfaceSetupIDType{}
	return &this
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *InterfaceSetupIDType) GetInterfaceId() string {
	if o == nil || IsNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupIDType) GetInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *InterfaceSetupIDType) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *InterfaceSetupIDType) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *InterfaceSetupIDType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupIDType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *InterfaceSetupIDType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *InterfaceSetupIDType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o InterfaceSetupIDType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceSetupIDType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableInterfaceSetupIDType struct {
	value *InterfaceSetupIDType
	isSet bool
}

func (v NullableInterfaceSetupIDType) Get() *InterfaceSetupIDType {
	return v.value
}

func (v *NullableInterfaceSetupIDType) Set(val *InterfaceSetupIDType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceSetupIDType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceSetupIDType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceSetupIDType(val *InterfaceSetupIDType) *NullableInterfaceSetupIDType {
	return &NullableInterfaceSetupIDType{value: val, isSet: true}
}

func (v NullableInterfaceSetupIDType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceSetupIDType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


