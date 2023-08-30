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

// checks if the StatusStatisticType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusStatisticType{}

// StatusStatisticType List of statistic units for a particular status.
type StatusStatisticType struct {
	// Unit type to hold statistic code and value pair.
	StatisticUnit []StatisticUnitType `json:"statisticUnit,omitempty"`
	// Status type of statistic, e.g. Definite.
	Status *string `json:"status,omitempty"`
}

// NewStatusStatisticType instantiates a new StatusStatisticType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusStatisticType() *StatusStatisticType {
	this := StatusStatisticType{}
	return &this
}

// NewStatusStatisticTypeWithDefaults instantiates a new StatusStatisticType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusStatisticTypeWithDefaults() *StatusStatisticType {
	this := StatusStatisticType{}
	return &this
}

// GetStatisticUnit returns the StatisticUnit field value if set, zero value otherwise.
func (o *StatusStatisticType) GetStatisticUnit() []StatisticUnitType {
	if o == nil || IsNil(o.StatisticUnit) {
		var ret []StatisticUnitType
		return ret
	}
	return o.StatisticUnit
}

// GetStatisticUnitOk returns a tuple with the StatisticUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusStatisticType) GetStatisticUnitOk() ([]StatisticUnitType, bool) {
	if o == nil || IsNil(o.StatisticUnit) {
		return nil, false
	}
	return o.StatisticUnit, true
}

// HasStatisticUnit returns a boolean if a field has been set.
func (o *StatusStatisticType) HasStatisticUnit() bool {
	if o != nil && !IsNil(o.StatisticUnit) {
		return true
	}

	return false
}

// SetStatisticUnit gets a reference to the given []StatisticUnitType and assigns it to the StatisticUnit field.
func (o *StatusStatisticType) SetStatisticUnit(v []StatisticUnitType) {
	o.StatisticUnit = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StatusStatisticType) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusStatisticType) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StatusStatisticType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StatusStatisticType) SetStatus(v string) {
	o.Status = &v
}

func (o StatusStatisticType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusStatisticType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatisticUnit) {
		toSerialize["statisticUnit"] = o.StatisticUnit
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableStatusStatisticType struct {
	value *StatusStatisticType
	isSet bool
}

func (v NullableStatusStatisticType) Get() *StatusStatisticType {
	return v.value
}

func (v *NullableStatusStatisticType) Set(val *StatusStatisticType) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusStatisticType) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusStatisticType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusStatisticType(val *StatusStatisticType) *NullableStatusStatisticType {
	return &NullableStatusStatisticType{value: val, isSet: true}
}

func (v NullableStatusStatisticType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusStatisticType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


