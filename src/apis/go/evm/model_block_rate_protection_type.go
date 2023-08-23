/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BlockRateProtectionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockRateProtectionType{}

// BlockRateProtectionType Block Rate Protection code information.
type BlockRateProtectionType struct {
	Criteria *RateProtectionType `json:"criteria,omitempty"`
	// Specifies a single date.
	ProtectedDates []string `json:"protectedDates,omitempty"`
}

// NewBlockRateProtectionType instantiates a new BlockRateProtectionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockRateProtectionType() *BlockRateProtectionType {
	this := BlockRateProtectionType{}
	return &this
}

// NewBlockRateProtectionTypeWithDefaults instantiates a new BlockRateProtectionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockRateProtectionTypeWithDefaults() *BlockRateProtectionType {
	this := BlockRateProtectionType{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *BlockRateProtectionType) GetCriteria() RateProtectionType {
	if o == nil || IsNil(o.Criteria) {
		var ret RateProtectionType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockRateProtectionType) GetCriteriaOk() (*RateProtectionType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *BlockRateProtectionType) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given RateProtectionType and assigns it to the Criteria field.
func (o *BlockRateProtectionType) SetCriteria(v RateProtectionType) {
	o.Criteria = &v
}

// GetProtectedDates returns the ProtectedDates field value if set, zero value otherwise.
func (o *BlockRateProtectionType) GetProtectedDates() []string {
	if o == nil || IsNil(o.ProtectedDates) {
		var ret []string
		return ret
	}
	return o.ProtectedDates
}

// GetProtectedDatesOk returns a tuple with the ProtectedDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockRateProtectionType) GetProtectedDatesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProtectedDates) {
		return nil, false
	}
	return o.ProtectedDates, true
}

// HasProtectedDates returns a boolean if a field has been set.
func (o *BlockRateProtectionType) HasProtectedDates() bool {
	if o != nil && !IsNil(o.ProtectedDates) {
		return true
	}

	return false
}

// SetProtectedDates gets a reference to the given []string and assigns it to the ProtectedDates field.
func (o *BlockRateProtectionType) SetProtectedDates(v []string) {
	o.ProtectedDates = v
}

func (o BlockRateProtectionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockRateProtectionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.ProtectedDates) {
		toSerialize["protectedDates"] = o.ProtectedDates
	}
	return toSerialize, nil
}

type NullableBlockRateProtectionType struct {
	value *BlockRateProtectionType
	isSet bool
}

func (v NullableBlockRateProtectionType) Get() *BlockRateProtectionType {
	return v.value
}

func (v *NullableBlockRateProtectionType) Set(val *BlockRateProtectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockRateProtectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockRateProtectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockRateProtectionType(val *BlockRateProtectionType) *NullableBlockRateProtectionType {
	return &NullableBlockRateProtectionType{value: val, isSet: true}
}

func (v NullableBlockRateProtectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockRateProtectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


