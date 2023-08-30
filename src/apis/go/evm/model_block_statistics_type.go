/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
)

// checks if the BlockStatisticsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockStatisticsType{}

// BlockStatisticsType struct for BlockStatisticsType
type BlockStatisticsType struct {
	AllocatedRoomStatistics *RoomStatisticsType `json:"allocatedRoomStatistics,omitempty"`
	ActualRoomStatistics *RoomStatisticsType `json:"actualRoomStatistics,omitempty"`
	// Indicates the catering revenue on the books.
	CateringRevenueOnBooks *float32 `json:"cateringRevenueOnBooks,omitempty"`
	// Indicates actualized catering revenue.
	ActualCateringRevenue *float32 `json:"actualCateringRevenue,omitempty"`
}

// NewBlockStatisticsType instantiates a new BlockStatisticsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockStatisticsType() *BlockStatisticsType {
	this := BlockStatisticsType{}
	return &this
}

// NewBlockStatisticsTypeWithDefaults instantiates a new BlockStatisticsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockStatisticsTypeWithDefaults() *BlockStatisticsType {
	this := BlockStatisticsType{}
	return &this
}

// GetAllocatedRoomStatistics returns the AllocatedRoomStatistics field value if set, zero value otherwise.
func (o *BlockStatisticsType) GetAllocatedRoomStatistics() RoomStatisticsType {
	if o == nil || IsNil(o.AllocatedRoomStatistics) {
		var ret RoomStatisticsType
		return ret
	}
	return *o.AllocatedRoomStatistics
}

// GetAllocatedRoomStatisticsOk returns a tuple with the AllocatedRoomStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsType) GetAllocatedRoomStatisticsOk() (*RoomStatisticsType, bool) {
	if o == nil || IsNil(o.AllocatedRoomStatistics) {
		return nil, false
	}
	return o.AllocatedRoomStatistics, true
}

// HasAllocatedRoomStatistics returns a boolean if a field has been set.
func (o *BlockStatisticsType) HasAllocatedRoomStatistics() bool {
	if o != nil && !IsNil(o.AllocatedRoomStatistics) {
		return true
	}

	return false
}

// SetAllocatedRoomStatistics gets a reference to the given RoomStatisticsType and assigns it to the AllocatedRoomStatistics field.
func (o *BlockStatisticsType) SetAllocatedRoomStatistics(v RoomStatisticsType) {
	o.AllocatedRoomStatistics = &v
}

// GetActualRoomStatistics returns the ActualRoomStatistics field value if set, zero value otherwise.
func (o *BlockStatisticsType) GetActualRoomStatistics() RoomStatisticsType {
	if o == nil || IsNil(o.ActualRoomStatistics) {
		var ret RoomStatisticsType
		return ret
	}
	return *o.ActualRoomStatistics
}

// GetActualRoomStatisticsOk returns a tuple with the ActualRoomStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsType) GetActualRoomStatisticsOk() (*RoomStatisticsType, bool) {
	if o == nil || IsNil(o.ActualRoomStatistics) {
		return nil, false
	}
	return o.ActualRoomStatistics, true
}

// HasActualRoomStatistics returns a boolean if a field has been set.
func (o *BlockStatisticsType) HasActualRoomStatistics() bool {
	if o != nil && !IsNil(o.ActualRoomStatistics) {
		return true
	}

	return false
}

// SetActualRoomStatistics gets a reference to the given RoomStatisticsType and assigns it to the ActualRoomStatistics field.
func (o *BlockStatisticsType) SetActualRoomStatistics(v RoomStatisticsType) {
	o.ActualRoomStatistics = &v
}

// GetCateringRevenueOnBooks returns the CateringRevenueOnBooks field value if set, zero value otherwise.
func (o *BlockStatisticsType) GetCateringRevenueOnBooks() float32 {
	if o == nil || IsNil(o.CateringRevenueOnBooks) {
		var ret float32
		return ret
	}
	return *o.CateringRevenueOnBooks
}

// GetCateringRevenueOnBooksOk returns a tuple with the CateringRevenueOnBooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsType) GetCateringRevenueOnBooksOk() (*float32, bool) {
	if o == nil || IsNil(o.CateringRevenueOnBooks) {
		return nil, false
	}
	return o.CateringRevenueOnBooks, true
}

// HasCateringRevenueOnBooks returns a boolean if a field has been set.
func (o *BlockStatisticsType) HasCateringRevenueOnBooks() bool {
	if o != nil && !IsNil(o.CateringRevenueOnBooks) {
		return true
	}

	return false
}

// SetCateringRevenueOnBooks gets a reference to the given float32 and assigns it to the CateringRevenueOnBooks field.
func (o *BlockStatisticsType) SetCateringRevenueOnBooks(v float32) {
	o.CateringRevenueOnBooks = &v
}

// GetActualCateringRevenue returns the ActualCateringRevenue field value if set, zero value otherwise.
func (o *BlockStatisticsType) GetActualCateringRevenue() float32 {
	if o == nil || IsNil(o.ActualCateringRevenue) {
		var ret float32
		return ret
	}
	return *o.ActualCateringRevenue
}

// GetActualCateringRevenueOk returns a tuple with the ActualCateringRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsType) GetActualCateringRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.ActualCateringRevenue) {
		return nil, false
	}
	return o.ActualCateringRevenue, true
}

// HasActualCateringRevenue returns a boolean if a field has been set.
func (o *BlockStatisticsType) HasActualCateringRevenue() bool {
	if o != nil && !IsNil(o.ActualCateringRevenue) {
		return true
	}

	return false
}

// SetActualCateringRevenue gets a reference to the given float32 and assigns it to the ActualCateringRevenue field.
func (o *BlockStatisticsType) SetActualCateringRevenue(v float32) {
	o.ActualCateringRevenue = &v
}

func (o BlockStatisticsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockStatisticsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllocatedRoomStatistics) {
		toSerialize["allocatedRoomStatistics"] = o.AllocatedRoomStatistics
	}
	if !IsNil(o.ActualRoomStatistics) {
		toSerialize["actualRoomStatistics"] = o.ActualRoomStatistics
	}
	if !IsNil(o.CateringRevenueOnBooks) {
		toSerialize["cateringRevenueOnBooks"] = o.CateringRevenueOnBooks
	}
	if !IsNil(o.ActualCateringRevenue) {
		toSerialize["actualCateringRevenue"] = o.ActualCateringRevenue
	}
	return toSerialize, nil
}

type NullableBlockStatisticsType struct {
	value *BlockStatisticsType
	isSet bool
}

func (v NullableBlockStatisticsType) Get() *BlockStatisticsType {
	return v.value
}

func (v *NullableBlockStatisticsType) Set(val *BlockStatisticsType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockStatisticsType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockStatisticsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockStatisticsType(val *BlockStatisticsType) *NullableBlockStatisticsType {
	return &NullableBlockStatisticsType{value: val, isSet: true}
}

func (v NullableBlockStatisticsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockStatisticsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


