/*
OPERA Cloud Inventory API

APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inv

import (
	"encoding/json"
)

// checks if the BlockSummaryStatisticType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockSummaryStatisticType{}

// BlockSummaryStatisticType Statistic units group by status, booking and daily targets.
type BlockSummaryStatisticType struct {
	// Statistic summary for a particular status.
	StatusSummaryStatistic []StatusStatisticType `json:"statusSummaryStatistic,omitempty"`
	// Unit type to hold statistic code and value pair.
	BookingSummaryStatistic []StatisticUnitType `json:"bookingSummaryStatistic,omitempty"`
	// Unit type to hold statistic code and value pair.
	DailyTargetsSummaryStatistic []StatisticUnitType `json:"dailyTargetsSummaryStatistic,omitempty"`
	// Date of the block inventory statistic.
	StatisticDate *string `json:"statisticDate,omitempty"`
}

// NewBlockSummaryStatisticType instantiates a new BlockSummaryStatisticType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSummaryStatisticType() *BlockSummaryStatisticType {
	this := BlockSummaryStatisticType{}
	return &this
}

// NewBlockSummaryStatisticTypeWithDefaults instantiates a new BlockSummaryStatisticType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSummaryStatisticTypeWithDefaults() *BlockSummaryStatisticType {
	this := BlockSummaryStatisticType{}
	return &this
}

// GetStatusSummaryStatistic returns the StatusSummaryStatistic field value if set, zero value otherwise.
func (o *BlockSummaryStatisticType) GetStatusSummaryStatistic() []StatusStatisticType {
	if o == nil || IsNil(o.StatusSummaryStatistic) {
		var ret []StatusStatisticType
		return ret
	}
	return o.StatusSummaryStatistic
}

// GetStatusSummaryStatisticOk returns a tuple with the StatusSummaryStatistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSummaryStatisticType) GetStatusSummaryStatisticOk() ([]StatusStatisticType, bool) {
	if o == nil || IsNil(o.StatusSummaryStatistic) {
		return nil, false
	}
	return o.StatusSummaryStatistic, true
}

// HasStatusSummaryStatistic returns a boolean if a field has been set.
func (o *BlockSummaryStatisticType) HasStatusSummaryStatistic() bool {
	if o != nil && !IsNil(o.StatusSummaryStatistic) {
		return true
	}

	return false
}

// SetStatusSummaryStatistic gets a reference to the given []StatusStatisticType and assigns it to the StatusSummaryStatistic field.
func (o *BlockSummaryStatisticType) SetStatusSummaryStatistic(v []StatusStatisticType) {
	o.StatusSummaryStatistic = v
}

// GetBookingSummaryStatistic returns the BookingSummaryStatistic field value if set, zero value otherwise.
func (o *BlockSummaryStatisticType) GetBookingSummaryStatistic() []StatisticUnitType {
	if o == nil || IsNil(o.BookingSummaryStatistic) {
		var ret []StatisticUnitType
		return ret
	}
	return o.BookingSummaryStatistic
}

// GetBookingSummaryStatisticOk returns a tuple with the BookingSummaryStatistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSummaryStatisticType) GetBookingSummaryStatisticOk() ([]StatisticUnitType, bool) {
	if o == nil || IsNil(o.BookingSummaryStatistic) {
		return nil, false
	}
	return o.BookingSummaryStatistic, true
}

// HasBookingSummaryStatistic returns a boolean if a field has been set.
func (o *BlockSummaryStatisticType) HasBookingSummaryStatistic() bool {
	if o != nil && !IsNil(o.BookingSummaryStatistic) {
		return true
	}

	return false
}

// SetBookingSummaryStatistic gets a reference to the given []StatisticUnitType and assigns it to the BookingSummaryStatistic field.
func (o *BlockSummaryStatisticType) SetBookingSummaryStatistic(v []StatisticUnitType) {
	o.BookingSummaryStatistic = v
}

// GetDailyTargetsSummaryStatistic returns the DailyTargetsSummaryStatistic field value if set, zero value otherwise.
func (o *BlockSummaryStatisticType) GetDailyTargetsSummaryStatistic() []StatisticUnitType {
	if o == nil || IsNil(o.DailyTargetsSummaryStatistic) {
		var ret []StatisticUnitType
		return ret
	}
	return o.DailyTargetsSummaryStatistic
}

// GetDailyTargetsSummaryStatisticOk returns a tuple with the DailyTargetsSummaryStatistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSummaryStatisticType) GetDailyTargetsSummaryStatisticOk() ([]StatisticUnitType, bool) {
	if o == nil || IsNil(o.DailyTargetsSummaryStatistic) {
		return nil, false
	}
	return o.DailyTargetsSummaryStatistic, true
}

// HasDailyTargetsSummaryStatistic returns a boolean if a field has been set.
func (o *BlockSummaryStatisticType) HasDailyTargetsSummaryStatistic() bool {
	if o != nil && !IsNil(o.DailyTargetsSummaryStatistic) {
		return true
	}

	return false
}

// SetDailyTargetsSummaryStatistic gets a reference to the given []StatisticUnitType and assigns it to the DailyTargetsSummaryStatistic field.
func (o *BlockSummaryStatisticType) SetDailyTargetsSummaryStatistic(v []StatisticUnitType) {
	o.DailyTargetsSummaryStatistic = v
}

// GetStatisticDate returns the StatisticDate field value if set, zero value otherwise.
func (o *BlockSummaryStatisticType) GetStatisticDate() string {
	if o == nil || IsNil(o.StatisticDate) {
		var ret string
		return ret
	}
	return *o.StatisticDate
}

// GetStatisticDateOk returns a tuple with the StatisticDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSummaryStatisticType) GetStatisticDateOk() (*string, bool) {
	if o == nil || IsNil(o.StatisticDate) {
		return nil, false
	}
	return o.StatisticDate, true
}

// HasStatisticDate returns a boolean if a field has been set.
func (o *BlockSummaryStatisticType) HasStatisticDate() bool {
	if o != nil && !IsNil(o.StatisticDate) {
		return true
	}

	return false
}

// SetStatisticDate gets a reference to the given string and assigns it to the StatisticDate field.
func (o *BlockSummaryStatisticType) SetStatisticDate(v string) {
	o.StatisticDate = &v
}

func (o BlockSummaryStatisticType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockSummaryStatisticType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusSummaryStatistic) {
		toSerialize["statusSummaryStatistic"] = o.StatusSummaryStatistic
	}
	if !IsNil(o.BookingSummaryStatistic) {
		toSerialize["bookingSummaryStatistic"] = o.BookingSummaryStatistic
	}
	if !IsNil(o.DailyTargetsSummaryStatistic) {
		toSerialize["dailyTargetsSummaryStatistic"] = o.DailyTargetsSummaryStatistic
	}
	if !IsNil(o.StatisticDate) {
		toSerialize["statisticDate"] = o.StatisticDate
	}
	return toSerialize, nil
}

type NullableBlockSummaryStatisticType struct {
	value *BlockSummaryStatisticType
	isSet bool
}

func (v NullableBlockSummaryStatisticType) Get() *BlockSummaryStatisticType {
	return v.value
}

func (v *NullableBlockSummaryStatisticType) Set(val *BlockSummaryStatisticType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSummaryStatisticType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSummaryStatisticType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSummaryStatisticType(val *BlockSummaryStatisticType) *NullableBlockSummaryStatisticType {
	return &NullableBlockSummaryStatisticType{value: val, isSet: true}
}

func (v NullableBlockSummaryStatisticType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSummaryStatisticType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


