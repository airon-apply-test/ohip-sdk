/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MasterAccountInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MasterAccountInfoType{}

// MasterAccountInfoType struct for MasterAccountInfoType
type MasterAccountInfoType struct {
	MasterAccountId *UniqueIDType `json:"masterAccountId,omitempty"`
	// Name of the Master account.
	MasterAccountName *string `json:"masterAccountName,omitempty"`
}

// NewMasterAccountInfoType instantiates a new MasterAccountInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterAccountInfoType() *MasterAccountInfoType {
	this := MasterAccountInfoType{}
	return &this
}

// NewMasterAccountInfoTypeWithDefaults instantiates a new MasterAccountInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterAccountInfoTypeWithDefaults() *MasterAccountInfoType {
	this := MasterAccountInfoType{}
	return &this
}

// GetMasterAccountId returns the MasterAccountId field value if set, zero value otherwise.
func (o *MasterAccountInfoType) GetMasterAccountId() UniqueIDType {
	if o == nil || IsNil(o.MasterAccountId) {
		var ret UniqueIDType
		return ret
	}
	return *o.MasterAccountId
}

// GetMasterAccountIdOk returns a tuple with the MasterAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterAccountInfoType) GetMasterAccountIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.MasterAccountId) {
		return nil, false
	}
	return o.MasterAccountId, true
}

// HasMasterAccountId returns a boolean if a field has been set.
func (o *MasterAccountInfoType) HasMasterAccountId() bool {
	if o != nil && !IsNil(o.MasterAccountId) {
		return true
	}

	return false
}

// SetMasterAccountId gets a reference to the given UniqueIDType and assigns it to the MasterAccountId field.
func (o *MasterAccountInfoType) SetMasterAccountId(v UniqueIDType) {
	o.MasterAccountId = &v
}

// GetMasterAccountName returns the MasterAccountName field value if set, zero value otherwise.
func (o *MasterAccountInfoType) GetMasterAccountName() string {
	if o == nil || IsNil(o.MasterAccountName) {
		var ret string
		return ret
	}
	return *o.MasterAccountName
}

// GetMasterAccountNameOk returns a tuple with the MasterAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterAccountInfoType) GetMasterAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.MasterAccountName) {
		return nil, false
	}
	return o.MasterAccountName, true
}

// HasMasterAccountName returns a boolean if a field has been set.
func (o *MasterAccountInfoType) HasMasterAccountName() bool {
	if o != nil && !IsNil(o.MasterAccountName) {
		return true
	}

	return false
}

// SetMasterAccountName gets a reference to the given string and assigns it to the MasterAccountName field.
func (o *MasterAccountInfoType) SetMasterAccountName(v string) {
	o.MasterAccountName = &v
}

func (o MasterAccountInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MasterAccountInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MasterAccountId) {
		toSerialize["masterAccountId"] = o.MasterAccountId
	}
	if !IsNil(o.MasterAccountName) {
		toSerialize["masterAccountName"] = o.MasterAccountName
	}
	return toSerialize, nil
}

type NullableMasterAccountInfoType struct {
	value *MasterAccountInfoType
	isSet bool
}

func (v NullableMasterAccountInfoType) Get() *MasterAccountInfoType {
	return v.value
}

func (v *NullableMasterAccountInfoType) Set(val *MasterAccountInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterAccountInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterAccountInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterAccountInfoType(val *MasterAccountInfoType) *NullableMasterAccountInfoType {
	return &NullableMasterAccountInfoType{value: val, isSet: true}
}

func (v NullableMasterAccountInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterAccountInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


