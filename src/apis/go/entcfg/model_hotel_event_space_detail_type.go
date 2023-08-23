/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HotelEventSpaceDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelEventSpaceDetailType{}

// HotelEventSpaceDetailType The detail info of hotel event space
type HotelEventSpaceDetailType struct {
	// The code of hotel event space.
	Code *string `json:"code,omitempty"`
	// Th description of the hotel event space
	Description *string `json:"description,omitempty"`
	// Th max capacity of the hotel event space
	MaxCapacity *int32 `json:"maxCapacity,omitempty"`
	// List of event space max occupancy.
	MaxOccupancies []int32 `json:"maxOccupancies,omitempty"`
}

// NewHotelEventSpaceDetailType instantiates a new HotelEventSpaceDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelEventSpaceDetailType() *HotelEventSpaceDetailType {
	this := HotelEventSpaceDetailType{}
	return &this
}

// NewHotelEventSpaceDetailTypeWithDefaults instantiates a new HotelEventSpaceDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelEventSpaceDetailTypeWithDefaults() *HotelEventSpaceDetailType {
	this := HotelEventSpaceDetailType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *HotelEventSpaceDetailType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelEventSpaceDetailType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *HotelEventSpaceDetailType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *HotelEventSpaceDetailType) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HotelEventSpaceDetailType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelEventSpaceDetailType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HotelEventSpaceDetailType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HotelEventSpaceDetailType) SetDescription(v string) {
	o.Description = &v
}

// GetMaxCapacity returns the MaxCapacity field value if set, zero value otherwise.
func (o *HotelEventSpaceDetailType) GetMaxCapacity() int32 {
	if o == nil || IsNil(o.MaxCapacity) {
		var ret int32
		return ret
	}
	return *o.MaxCapacity
}

// GetMaxCapacityOk returns a tuple with the MaxCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelEventSpaceDetailType) GetMaxCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxCapacity) {
		return nil, false
	}
	return o.MaxCapacity, true
}

// HasMaxCapacity returns a boolean if a field has been set.
func (o *HotelEventSpaceDetailType) HasMaxCapacity() bool {
	if o != nil && !IsNil(o.MaxCapacity) {
		return true
	}

	return false
}

// SetMaxCapacity gets a reference to the given int32 and assigns it to the MaxCapacity field.
func (o *HotelEventSpaceDetailType) SetMaxCapacity(v int32) {
	o.MaxCapacity = &v
}

// GetMaxOccupancies returns the MaxOccupancies field value if set, zero value otherwise.
func (o *HotelEventSpaceDetailType) GetMaxOccupancies() []int32 {
	if o == nil || IsNil(o.MaxOccupancies) {
		var ret []int32
		return ret
	}
	return o.MaxOccupancies
}

// GetMaxOccupanciesOk returns a tuple with the MaxOccupancies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelEventSpaceDetailType) GetMaxOccupanciesOk() ([]int32, bool) {
	if o == nil || IsNil(o.MaxOccupancies) {
		return nil, false
	}
	return o.MaxOccupancies, true
}

// HasMaxOccupancies returns a boolean if a field has been set.
func (o *HotelEventSpaceDetailType) HasMaxOccupancies() bool {
	if o != nil && !IsNil(o.MaxOccupancies) {
		return true
	}

	return false
}

// SetMaxOccupancies gets a reference to the given []int32 and assigns it to the MaxOccupancies field.
func (o *HotelEventSpaceDetailType) SetMaxOccupancies(v []int32) {
	o.MaxOccupancies = v
}

func (o HotelEventSpaceDetailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelEventSpaceDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MaxCapacity) {
		toSerialize["maxCapacity"] = o.MaxCapacity
	}
	if !IsNil(o.MaxOccupancies) {
		toSerialize["maxOccupancies"] = o.MaxOccupancies
	}
	return toSerialize, nil
}

type NullableHotelEventSpaceDetailType struct {
	value *HotelEventSpaceDetailType
	isSet bool
}

func (v NullableHotelEventSpaceDetailType) Get() *HotelEventSpaceDetailType {
	return v.value
}

func (v *NullableHotelEventSpaceDetailType) Set(val *HotelEventSpaceDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelEventSpaceDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelEventSpaceDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelEventSpaceDetailType(val *HotelEventSpaceDetailType) *NullableHotelEventSpaceDetailType {
	return &NullableHotelEventSpaceDetailType{value: val, isSet: true}
}

func (v NullableHotelEventSpaceDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelEventSpaceDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


