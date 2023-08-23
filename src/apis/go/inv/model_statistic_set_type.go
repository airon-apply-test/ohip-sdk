/*
OPERA Cloud Inventory API

APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StatisticSetType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatisticSetType{}

// StatisticSetType An instance of a statistic, which is a set containing revenue category and number category summaries.
type StatisticSetType struct {
	// Collection of RevenueCategorySummary elements. Used if revenue values reported.
	Revenue []RevenueCategorySummaryType `json:"revenue,omitempty"`
	// Collection of CountCategorySummary elements. Used if count values reported.
	Inventory []NumericCategorySummaryType `json:"inventory,omitempty"`
	// Date of the statistic.
	StatisticDate *string `json:"statisticDate,omitempty"`
	// Determines whether statistic date is a weekend date.
	WeekendDate *bool `json:"weekendDate,omitempty"`
}

// NewStatisticSetType instantiates a new StatisticSetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticSetType() *StatisticSetType {
	this := StatisticSetType{}
	return &this
}

// NewStatisticSetTypeWithDefaults instantiates a new StatisticSetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticSetTypeWithDefaults() *StatisticSetType {
	this := StatisticSetType{}
	return &this
}

// GetRevenue returns the Revenue field value if set, zero value otherwise.
func (o *StatisticSetType) GetRevenue() []RevenueCategorySummaryType {
	if o == nil || IsNil(o.Revenue) {
		var ret []RevenueCategorySummaryType
		return ret
	}
	return o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticSetType) GetRevenueOk() ([]RevenueCategorySummaryType, bool) {
	if o == nil || IsNil(o.Revenue) {
		return nil, false
	}
	return o.Revenue, true
}

// HasRevenue returns a boolean if a field has been set.
func (o *StatisticSetType) HasRevenue() bool {
	if o != nil && !IsNil(o.Revenue) {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given []RevenueCategorySummaryType and assigns it to the Revenue field.
func (o *StatisticSetType) SetRevenue(v []RevenueCategorySummaryType) {
	o.Revenue = v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *StatisticSetType) GetInventory() []NumericCategorySummaryType {
	if o == nil || IsNil(o.Inventory) {
		var ret []NumericCategorySummaryType
		return ret
	}
	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticSetType) GetInventoryOk() ([]NumericCategorySummaryType, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *StatisticSetType) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given []NumericCategorySummaryType and assigns it to the Inventory field.
func (o *StatisticSetType) SetInventory(v []NumericCategorySummaryType) {
	o.Inventory = v
}

// GetStatisticDate returns the StatisticDate field value if set, zero value otherwise.
func (o *StatisticSetType) GetStatisticDate() string {
	if o == nil || IsNil(o.StatisticDate) {
		var ret string
		return ret
	}
	return *o.StatisticDate
}

// GetStatisticDateOk returns a tuple with the StatisticDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticSetType) GetStatisticDateOk() (*string, bool) {
	if o == nil || IsNil(o.StatisticDate) {
		return nil, false
	}
	return o.StatisticDate, true
}

// HasStatisticDate returns a boolean if a field has been set.
func (o *StatisticSetType) HasStatisticDate() bool {
	if o != nil && !IsNil(o.StatisticDate) {
		return true
	}

	return false
}

// SetStatisticDate gets a reference to the given string and assigns it to the StatisticDate field.
func (o *StatisticSetType) SetStatisticDate(v string) {
	o.StatisticDate = &v
}

// GetWeekendDate returns the WeekendDate field value if set, zero value otherwise.
func (o *StatisticSetType) GetWeekendDate() bool {
	if o == nil || IsNil(o.WeekendDate) {
		var ret bool
		return ret
	}
	return *o.WeekendDate
}

// GetWeekendDateOk returns a tuple with the WeekendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticSetType) GetWeekendDateOk() (*bool, bool) {
	if o == nil || IsNil(o.WeekendDate) {
		return nil, false
	}
	return o.WeekendDate, true
}

// HasWeekendDate returns a boolean if a field has been set.
func (o *StatisticSetType) HasWeekendDate() bool {
	if o != nil && !IsNil(o.WeekendDate) {
		return true
	}

	return false
}

// SetWeekendDate gets a reference to the given bool and assigns it to the WeekendDate field.
func (o *StatisticSetType) SetWeekendDate(v bool) {
	o.WeekendDate = &v
}

func (o StatisticSetType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatisticSetType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Revenue) {
		toSerialize["revenue"] = o.Revenue
	}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.StatisticDate) {
		toSerialize["statisticDate"] = o.StatisticDate
	}
	if !IsNil(o.WeekendDate) {
		toSerialize["weekendDate"] = o.WeekendDate
	}
	return toSerialize, nil
}

type NullableStatisticSetType struct {
	value *StatisticSetType
	isSet bool
}

func (v NullableStatisticSetType) Get() *StatisticSetType {
	return v.value
}

func (v *NullableStatisticSetType) Set(val *StatisticSetType) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticSetType) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticSetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticSetType(val *StatisticSetType) *NullableStatisticSetType {
	return &NullableStatisticSetType{value: val, isSet: true}
}

func (v NullableStatisticSetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticSetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


