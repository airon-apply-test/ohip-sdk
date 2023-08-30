/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the InterfaceControlGroupType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceControlGroupType{}

// InterfaceControlGroupType Type to group the different OXI Parameters/Settings.
type InterfaceControlGroupType struct {
	// Group Name.
	GroupName *string `json:"groupName,omitempty"`
	// Group Display Name.
	DisplayName *string `json:"displayName,omitempty"`
	// OXI Parameters/Settings.
	InterfaceControls []InterfaceControlType `json:"interfaceControls,omitempty"`
}

// NewInterfaceControlGroupType instantiates a new InterfaceControlGroupType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceControlGroupType() *InterfaceControlGroupType {
	this := InterfaceControlGroupType{}
	return &this
}

// NewInterfaceControlGroupTypeWithDefaults instantiates a new InterfaceControlGroupType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceControlGroupTypeWithDefaults() *InterfaceControlGroupType {
	this := InterfaceControlGroupType{}
	return &this
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *InterfaceControlGroupType) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlGroupType) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *InterfaceControlGroupType) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *InterfaceControlGroupType) SetGroupName(v string) {
	o.GroupName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InterfaceControlGroupType) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlGroupType) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InterfaceControlGroupType) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InterfaceControlGroupType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetInterfaceControls returns the InterfaceControls field value if set, zero value otherwise.
func (o *InterfaceControlGroupType) GetInterfaceControls() []InterfaceControlType {
	if o == nil || IsNil(o.InterfaceControls) {
		var ret []InterfaceControlType
		return ret
	}
	return o.InterfaceControls
}

// GetInterfaceControlsOk returns a tuple with the InterfaceControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlGroupType) GetInterfaceControlsOk() ([]InterfaceControlType, bool) {
	if o == nil || IsNil(o.InterfaceControls) {
		return nil, false
	}
	return o.InterfaceControls, true
}

// HasInterfaceControls returns a boolean if a field has been set.
func (o *InterfaceControlGroupType) HasInterfaceControls() bool {
	if o != nil && !IsNil(o.InterfaceControls) {
		return true
	}

	return false
}

// SetInterfaceControls gets a reference to the given []InterfaceControlType and assigns it to the InterfaceControls field.
func (o *InterfaceControlGroupType) SetInterfaceControls(v []InterfaceControlType) {
	o.InterfaceControls = v
}

func (o InterfaceControlGroupType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceControlGroupType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.InterfaceControls) {
		toSerialize["interfaceControls"] = o.InterfaceControls
	}
	return toSerialize, nil
}

type NullableInterfaceControlGroupType struct {
	value *InterfaceControlGroupType
	isSet bool
}

func (v NullableInterfaceControlGroupType) Get() *InterfaceControlGroupType {
	return v.value
}

func (v *NullableInterfaceControlGroupType) Set(val *InterfaceControlGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceControlGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceControlGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceControlGroupType(val *InterfaceControlGroupType) *NullableInterfaceControlGroupType {
	return &NullableInterfaceControlGroupType{value: val, isSet: true}
}

func (v NullableInterfaceControlGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceControlGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


