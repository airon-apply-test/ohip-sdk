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

// checks if the ResAttachedProfileType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResAttachedProfileType{}

// ResAttachedProfileType struct for ResAttachedProfileType
type ResAttachedProfileType struct {
	// Attached profile name
	Name *string `json:"name,omitempty"`
	// Unique Id that references an object uniquely in the system.
	ProfileIdList []UniqueIDType `json:"profileIdList,omitempty"`
	ReservationProfileType *ResProfileTypeType `json:"reservationProfileType,omitempty"`
}

// NewResAttachedProfileType instantiates a new ResAttachedProfileType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResAttachedProfileType() *ResAttachedProfileType {
	this := ResAttachedProfileType{}
	return &this
}

// NewResAttachedProfileTypeWithDefaults instantiates a new ResAttachedProfileType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResAttachedProfileTypeWithDefaults() *ResAttachedProfileType {
	this := ResAttachedProfileType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResAttachedProfileType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResAttachedProfileType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResAttachedProfileType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResAttachedProfileType) SetName(v string) {
	o.Name = &v
}

// GetProfileIdList returns the ProfileIdList field value if set, zero value otherwise.
func (o *ResAttachedProfileType) GetProfileIdList() []UniqueIDType {
	if o == nil || IsNil(o.ProfileIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ProfileIdList
}

// GetProfileIdListOk returns a tuple with the ProfileIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResAttachedProfileType) GetProfileIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileIdList) {
		return nil, false
	}
	return o.ProfileIdList, true
}

// HasProfileIdList returns a boolean if a field has been set.
func (o *ResAttachedProfileType) HasProfileIdList() bool {
	if o != nil && !IsNil(o.ProfileIdList) {
		return true
	}

	return false
}

// SetProfileIdList gets a reference to the given []UniqueIDType and assigns it to the ProfileIdList field.
func (o *ResAttachedProfileType) SetProfileIdList(v []UniqueIDType) {
	o.ProfileIdList = v
}

// GetReservationProfileType returns the ReservationProfileType field value if set, zero value otherwise.
func (o *ResAttachedProfileType) GetReservationProfileType() ResProfileTypeType {
	if o == nil || IsNil(o.ReservationProfileType) {
		var ret ResProfileTypeType
		return ret
	}
	return *o.ReservationProfileType
}

// GetReservationProfileTypeOk returns a tuple with the ReservationProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResAttachedProfileType) GetReservationProfileTypeOk() (*ResProfileTypeType, bool) {
	if o == nil || IsNil(o.ReservationProfileType) {
		return nil, false
	}
	return o.ReservationProfileType, true
}

// HasReservationProfileType returns a boolean if a field has been set.
func (o *ResAttachedProfileType) HasReservationProfileType() bool {
	if o != nil && !IsNil(o.ReservationProfileType) {
		return true
	}

	return false
}

// SetReservationProfileType gets a reference to the given ResProfileTypeType and assigns it to the ReservationProfileType field.
func (o *ResAttachedProfileType) SetReservationProfileType(v ResProfileTypeType) {
	o.ReservationProfileType = &v
}

func (o ResAttachedProfileType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResAttachedProfileType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProfileIdList) {
		toSerialize["profileIdList"] = o.ProfileIdList
	}
	if !IsNil(o.ReservationProfileType) {
		toSerialize["reservationProfileType"] = o.ReservationProfileType
	}
	return toSerialize, nil
}

type NullableResAttachedProfileType struct {
	value *ResAttachedProfileType
	isSet bool
}

func (v NullableResAttachedProfileType) Get() *ResAttachedProfileType {
	return v.value
}

func (v *NullableResAttachedProfileType) Set(val *ResAttachedProfileType) {
	v.value = val
	v.isSet = true
}

func (v NullableResAttachedProfileType) IsSet() bool {
	return v.isSet
}

func (v *NullableResAttachedProfileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResAttachedProfileType(val *ResAttachedProfileType) *NullableResAttachedProfileType {
	return &NullableResAttachedProfileType{value: val, isSet: true}
}

func (v NullableResAttachedProfileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResAttachedProfileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


