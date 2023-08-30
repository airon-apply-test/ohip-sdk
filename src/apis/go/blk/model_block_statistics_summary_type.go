/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the BlockStatisticsSummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockStatisticsSummaryType{}

// BlockStatisticsSummaryType Type to store summary detail like rooms sold, average room rate and revenue for a statistic type.
type BlockStatisticsSummaryType struct {
	// Total Number of Rooms Picked Up.
	RoomsSold *int32 `json:"roomsSold,omitempty"`
	RevenueSummary *RevenueSummaryType `json:"revenueSummary,omitempty"`
	AvgRoomRate *float32 `json:"avgRoomRate,omitempty"`
	StatisticType *StatisticsSummaryType `json:"statisticType,omitempty"`
}

// NewBlockStatisticsSummaryType instantiates a new BlockStatisticsSummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockStatisticsSummaryType() *BlockStatisticsSummaryType {
	this := BlockStatisticsSummaryType{}
	return &this
}

// NewBlockStatisticsSummaryTypeWithDefaults instantiates a new BlockStatisticsSummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockStatisticsSummaryTypeWithDefaults() *BlockStatisticsSummaryType {
	this := BlockStatisticsSummaryType{}
	return &this
}

// GetRoomsSold returns the RoomsSold field value if set, zero value otherwise.
func (o *BlockStatisticsSummaryType) GetRoomsSold() int32 {
	if o == nil || IsNil(o.RoomsSold) {
		var ret int32
		return ret
	}
	return *o.RoomsSold
}

// GetRoomsSoldOk returns a tuple with the RoomsSold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsSummaryType) GetRoomsSoldOk() (*int32, bool) {
	if o == nil || IsNil(o.RoomsSold) {
		return nil, false
	}
	return o.RoomsSold, true
}

// HasRoomsSold returns a boolean if a field has been set.
func (o *BlockStatisticsSummaryType) HasRoomsSold() bool {
	if o != nil && !IsNil(o.RoomsSold) {
		return true
	}

	return false
}

// SetRoomsSold gets a reference to the given int32 and assigns it to the RoomsSold field.
func (o *BlockStatisticsSummaryType) SetRoomsSold(v int32) {
	o.RoomsSold = &v
}

// GetRevenueSummary returns the RevenueSummary field value if set, zero value otherwise.
func (o *BlockStatisticsSummaryType) GetRevenueSummary() RevenueSummaryType {
	if o == nil || IsNil(o.RevenueSummary) {
		var ret RevenueSummaryType
		return ret
	}
	return *o.RevenueSummary
}

// GetRevenueSummaryOk returns a tuple with the RevenueSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsSummaryType) GetRevenueSummaryOk() (*RevenueSummaryType, bool) {
	if o == nil || IsNil(o.RevenueSummary) {
		return nil, false
	}
	return o.RevenueSummary, true
}

// HasRevenueSummary returns a boolean if a field has been set.
func (o *BlockStatisticsSummaryType) HasRevenueSummary() bool {
	if o != nil && !IsNil(o.RevenueSummary) {
		return true
	}

	return false
}

// SetRevenueSummary gets a reference to the given RevenueSummaryType and assigns it to the RevenueSummary field.
func (o *BlockStatisticsSummaryType) SetRevenueSummary(v RevenueSummaryType) {
	o.RevenueSummary = &v
}

// GetAvgRoomRate returns the AvgRoomRate field value if set, zero value otherwise.
func (o *BlockStatisticsSummaryType) GetAvgRoomRate() float32 {
	if o == nil || IsNil(o.AvgRoomRate) {
		var ret float32
		return ret
	}
	return *o.AvgRoomRate
}

// GetAvgRoomRateOk returns a tuple with the AvgRoomRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsSummaryType) GetAvgRoomRateOk() (*float32, bool) {
	if o == nil || IsNil(o.AvgRoomRate) {
		return nil, false
	}
	return o.AvgRoomRate, true
}

// HasAvgRoomRate returns a boolean if a field has been set.
func (o *BlockStatisticsSummaryType) HasAvgRoomRate() bool {
	if o != nil && !IsNil(o.AvgRoomRate) {
		return true
	}

	return false
}

// SetAvgRoomRate gets a reference to the given float32 and assigns it to the AvgRoomRate field.
func (o *BlockStatisticsSummaryType) SetAvgRoomRate(v float32) {
	o.AvgRoomRate = &v
}

// GetStatisticType returns the StatisticType field value if set, zero value otherwise.
func (o *BlockStatisticsSummaryType) GetStatisticType() StatisticsSummaryType {
	if o == nil || IsNil(o.StatisticType) {
		var ret StatisticsSummaryType
		return ret
	}
	return *o.StatisticType
}

// GetStatisticTypeOk returns a tuple with the StatisticType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockStatisticsSummaryType) GetStatisticTypeOk() (*StatisticsSummaryType, bool) {
	if o == nil || IsNil(o.StatisticType) {
		return nil, false
	}
	return o.StatisticType, true
}

// HasStatisticType returns a boolean if a field has been set.
func (o *BlockStatisticsSummaryType) HasStatisticType() bool {
	if o != nil && !IsNil(o.StatisticType) {
		return true
	}

	return false
}

// SetStatisticType gets a reference to the given StatisticsSummaryType and assigns it to the StatisticType field.
func (o *BlockStatisticsSummaryType) SetStatisticType(v StatisticsSummaryType) {
	o.StatisticType = &v
}

func (o BlockStatisticsSummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockStatisticsSummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomsSold) {
		toSerialize["roomsSold"] = o.RoomsSold
	}
	if !IsNil(o.RevenueSummary) {
		toSerialize["revenueSummary"] = o.RevenueSummary
	}
	if !IsNil(o.AvgRoomRate) {
		toSerialize["avgRoomRate"] = o.AvgRoomRate
	}
	if !IsNil(o.StatisticType) {
		toSerialize["statisticType"] = o.StatisticType
	}
	return toSerialize, nil
}

type NullableBlockStatisticsSummaryType struct {
	value *BlockStatisticsSummaryType
	isSet bool
}

func (v NullableBlockStatisticsSummaryType) Get() *BlockStatisticsSummaryType {
	return v.value
}

func (v *NullableBlockStatisticsSummaryType) Set(val *BlockStatisticsSummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockStatisticsSummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockStatisticsSummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockStatisticsSummaryType(val *BlockStatisticsSummaryType) *NullableBlockStatisticsSummaryType {
	return &NullableBlockStatisticsSummaryType{value: val, isSet: true}
}

func (v NullableBlockStatisticsSummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockStatisticsSummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


