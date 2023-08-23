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

// checks if the HotelEventSpaceSummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelEventSpaceSummaryType{}

// HotelEventSpaceSummaryType The summary info of hotel event space
type HotelEventSpaceSummaryType struct {
	// The total event space number which has the same space type, max capacity and setup style.
	No *int32 `json:"no,omitempty"`
	// The type of the event space
	SpaceType *string `json:"spaceType,omitempty"`
	// The max capacity of this event space group
	MaxCapacity *int32 `json:"maxCapacity,omitempty"`
	// List of event space max occupancy.
	MaxOccupancies []int32 `json:"maxOccupancies,omitempty"`
}

// NewHotelEventSpaceSummaryType instantiates a new HotelEventSpaceSummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelEventSpaceSummaryType() *HotelEventSpaceSummaryType {
	this := HotelEventSpaceSummaryType{}
	return &this
}

// NewHotelEventSpaceSummaryTypeWithDefaults instantiates a new HotelEventSpaceSummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelEventSpaceSummaryTypeWithDefaults() *HotelEventSpaceSummaryType {
	this := HotelEventSpaceSummaryType{}
	return &this
}

// GetNo returns the No field value if set, zero value otherwise.
func (o *HotelEventSpaceSummaryType) GetNo() int32 {
	if o == nil || IsNil(o.No) {
		var ret int32
		return ret
	}
	return *o.No
}

// GetNoOk returns a tuple with the No field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelEventSpaceSummaryType) GetNoOk() (*int32, bool) {
	if o == nil || IsNil(o.No) {
		return nil, false
	}
	return o.No, true
}

// HasNo returns a boolean if a field has been set.
func (o *HotelEventSpaceSummaryType) HasNo() bool {
	if o != nil && !IsNil(o.No) {
		return true
	}

	return false
}

// SetNo gets a reference to the given int32 and assigns it to the No field.
func (o *HotelEventSpaceSummaryType) SetNo(v int32) {
	o.No = &v
}

// GetSpaceType returns the SpaceType field value if set, zero value otherwise.
func (o *HotelEventSpaceSummaryType) GetSpaceType() string {
	if o == nil || IsNil(o.SpaceType) {
		var ret string
		return ret
	}
	return *o.SpaceType
}

// GetSpaceTypeOk returns a tuple with the SpaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelEventSpaceSummaryType) GetSpaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SpaceType) {
		return nil, false
	}
	return o.SpaceType, true
}

// HasSpaceType returns a boolean if a field has been set.
func (o *HotelEventSpaceSummaryType) HasSpaceType() bool {
	if o != nil && !IsNil(o.SpaceType) {
		return true
	}

	return false
}

// SetSpaceType gets a reference to the given string and assigns it to the SpaceType field.
func (o *HotelEventSpaceSummaryType) SetSpaceType(v string) {
	o.SpaceType = &v
}

// GetMaxCapacity returns the MaxCapacity field value if set, zero value otherwise.
func (o *HotelEventSpaceSummaryType) GetMaxCapacity() int32 {
	if o == nil || IsNil(o.MaxCapacity) {
		var ret int32
		return ret
	}
	return *o.MaxCapacity
}

// GetMaxCapacityOk returns a tuple with the MaxCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelEventSpaceSummaryType) GetMaxCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxCapacity) {
		return nil, false
	}
	return o.MaxCapacity, true
}

// HasMaxCapacity returns a boolean if a field has been set.
func (o *HotelEventSpaceSummaryType) HasMaxCapacity() bool {
	if o != nil && !IsNil(o.MaxCapacity) {
		return true
	}

	return false
}

// SetMaxCapacity gets a reference to the given int32 and assigns it to the MaxCapacity field.
func (o *HotelEventSpaceSummaryType) SetMaxCapacity(v int32) {
	o.MaxCapacity = &v
}

// GetMaxOccupancies returns the MaxOccupancies field value if set, zero value otherwise.
func (o *HotelEventSpaceSummaryType) GetMaxOccupancies() []int32 {
	if o == nil || IsNil(o.MaxOccupancies) {
		var ret []int32
		return ret
	}
	return o.MaxOccupancies
}

// GetMaxOccupanciesOk returns a tuple with the MaxOccupancies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelEventSpaceSummaryType) GetMaxOccupanciesOk() ([]int32, bool) {
	if o == nil || IsNil(o.MaxOccupancies) {
		return nil, false
	}
	return o.MaxOccupancies, true
}

// HasMaxOccupancies returns a boolean if a field has been set.
func (o *HotelEventSpaceSummaryType) HasMaxOccupancies() bool {
	if o != nil && !IsNil(o.MaxOccupancies) {
		return true
	}

	return false
}

// SetMaxOccupancies gets a reference to the given []int32 and assigns it to the MaxOccupancies field.
func (o *HotelEventSpaceSummaryType) SetMaxOccupancies(v []int32) {
	o.MaxOccupancies = v
}

func (o HotelEventSpaceSummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelEventSpaceSummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.No) {
		toSerialize["no"] = o.No
	}
	if !IsNil(o.SpaceType) {
		toSerialize["spaceType"] = o.SpaceType
	}
	if !IsNil(o.MaxCapacity) {
		toSerialize["maxCapacity"] = o.MaxCapacity
	}
	if !IsNil(o.MaxOccupancies) {
		toSerialize["maxOccupancies"] = o.MaxOccupancies
	}
	return toSerialize, nil
}

type NullableHotelEventSpaceSummaryType struct {
	value *HotelEventSpaceSummaryType
	isSet bool
}

func (v NullableHotelEventSpaceSummaryType) Get() *HotelEventSpaceSummaryType {
	return v.value
}

func (v *NullableHotelEventSpaceSummaryType) Set(val *HotelEventSpaceSummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelEventSpaceSummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelEventSpaceSummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelEventSpaceSummaryType(val *HotelEventSpaceSummaryType) *NullableHotelEventSpaceSummaryType {
	return &NullableHotelEventSpaceSummaryType{value: val, isSet: true}
}

func (v NullableHotelEventSpaceSummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelEventSpaceSummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


