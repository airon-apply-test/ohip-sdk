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

// checks if the MasterInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MasterInfoType{}

// MasterInfoType struct for MasterInfoType
type MasterInfoType struct {
	CodeInfo []CodeInfoType `json:"codeInfo,omitempty"`
	CodeType *MasterType `json:"codeType,omitempty"`
}

// NewMasterInfoType instantiates a new MasterInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterInfoType() *MasterInfoType {
	this := MasterInfoType{}
	return &this
}

// NewMasterInfoTypeWithDefaults instantiates a new MasterInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterInfoTypeWithDefaults() *MasterInfoType {
	this := MasterInfoType{}
	return &this
}

// GetCodeInfo returns the CodeInfo field value if set, zero value otherwise.
func (o *MasterInfoType) GetCodeInfo() []CodeInfoType {
	if o == nil || IsNil(o.CodeInfo) {
		var ret []CodeInfoType
		return ret
	}
	return o.CodeInfo
}

// GetCodeInfoOk returns a tuple with the CodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterInfoType) GetCodeInfoOk() ([]CodeInfoType, bool) {
	if o == nil || IsNil(o.CodeInfo) {
		return nil, false
	}
	return o.CodeInfo, true
}

// HasCodeInfo returns a boolean if a field has been set.
func (o *MasterInfoType) HasCodeInfo() bool {
	if o != nil && !IsNil(o.CodeInfo) {
		return true
	}

	return false
}

// SetCodeInfo gets a reference to the given []CodeInfoType and assigns it to the CodeInfo field.
func (o *MasterInfoType) SetCodeInfo(v []CodeInfoType) {
	o.CodeInfo = v
}

// GetCodeType returns the CodeType field value if set, zero value otherwise.
func (o *MasterInfoType) GetCodeType() MasterType {
	if o == nil || IsNil(o.CodeType) {
		var ret MasterType
		return ret
	}
	return *o.CodeType
}

// GetCodeTypeOk returns a tuple with the CodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterInfoType) GetCodeTypeOk() (*MasterType, bool) {
	if o == nil || IsNil(o.CodeType) {
		return nil, false
	}
	return o.CodeType, true
}

// HasCodeType returns a boolean if a field has been set.
func (o *MasterInfoType) HasCodeType() bool {
	if o != nil && !IsNil(o.CodeType) {
		return true
	}

	return false
}

// SetCodeType gets a reference to the given MasterType and assigns it to the CodeType field.
func (o *MasterInfoType) SetCodeType(v MasterType) {
	o.CodeType = &v
}

func (o MasterInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MasterInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeInfo) {
		toSerialize["codeInfo"] = o.CodeInfo
	}
	if !IsNil(o.CodeType) {
		toSerialize["codeType"] = o.CodeType
	}
	return toSerialize, nil
}

type NullableMasterInfoType struct {
	value *MasterInfoType
	isSet bool
}

func (v NullableMasterInfoType) Get() *MasterInfoType {
	return v.value
}

func (v *NullableMasterInfoType) Set(val *MasterInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterInfoType(val *MasterInfoType) *NullableMasterInfoType {
	return &NullableMasterInfoType{value: val, isSet: true}
}

func (v NullableMasterInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


