/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MasterInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MasterInfo{}

// MasterInfo The reference to the rate codes selected as best availabe rates
type MasterInfo struct {
	Rate *CodeDescriptionType `json:"rate,omitempty"`
}

// NewMasterInfo instantiates a new MasterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterInfo() *MasterInfo {
	this := MasterInfo{}
	return &this
}

// NewMasterInfoWithDefaults instantiates a new MasterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterInfoWithDefaults() *MasterInfo {
	this := MasterInfo{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *MasterInfo) GetRate() CodeDescriptionType {
	if o == nil || IsNil(o.Rate) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterInfo) GetRateOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *MasterInfo) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given CodeDescriptionType and assigns it to the Rate field.
func (o *MasterInfo) SetRate(v CodeDescriptionType) {
	o.Rate = &v
}

func (o MasterInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MasterInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	return toSerialize, nil
}

type NullableMasterInfo struct {
	value *MasterInfo
	isSet bool
}

func (v NullableMasterInfo) Get() *MasterInfo {
	return v.value
}

func (v *NullableMasterInfo) Set(val *MasterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterInfo(val *MasterInfo) *NullableMasterInfo {
	return &NullableMasterInfo{value: val, isSet: true}
}

func (v NullableMasterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


