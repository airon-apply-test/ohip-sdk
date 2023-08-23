/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LinkedActivitiesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkedActivitiesType{}

// LinkedActivitiesType Provides the data to create linked activities to an activity .
type LinkedActivitiesType struct {
	// Hotel Code of the Activity.
	HotelId *string `json:"hotelId,omitempty"`
	ActivityId *ActivityId `json:"activityId,omitempty"`
	OwnerCodeList []string `json:"ownerCodeList,omitempty"`
}

// NewLinkedActivitiesType instantiates a new LinkedActivitiesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedActivitiesType() *LinkedActivitiesType {
	this := LinkedActivitiesType{}
	return &this
}

// NewLinkedActivitiesTypeWithDefaults instantiates a new LinkedActivitiesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedActivitiesTypeWithDefaults() *LinkedActivitiesType {
	this := LinkedActivitiesType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *LinkedActivitiesType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedActivitiesType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *LinkedActivitiesType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *LinkedActivitiesType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *LinkedActivitiesType) GetActivityId() ActivityId {
	if o == nil || IsNil(o.ActivityId) {
		var ret ActivityId
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedActivitiesType) GetActivityIdOk() (*ActivityId, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *LinkedActivitiesType) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given ActivityId and assigns it to the ActivityId field.
func (o *LinkedActivitiesType) SetActivityId(v ActivityId) {
	o.ActivityId = &v
}

// GetOwnerCodeList returns the OwnerCodeList field value if set, zero value otherwise.
func (o *LinkedActivitiesType) GetOwnerCodeList() []string {
	if o == nil || IsNil(o.OwnerCodeList) {
		var ret []string
		return ret
	}
	return o.OwnerCodeList
}

// GetOwnerCodeListOk returns a tuple with the OwnerCodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedActivitiesType) GetOwnerCodeListOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerCodeList) {
		return nil, false
	}
	return o.OwnerCodeList, true
}

// HasOwnerCodeList returns a boolean if a field has been set.
func (o *LinkedActivitiesType) HasOwnerCodeList() bool {
	if o != nil && !IsNil(o.OwnerCodeList) {
		return true
	}

	return false
}

// SetOwnerCodeList gets a reference to the given []string and assigns it to the OwnerCodeList field.
func (o *LinkedActivitiesType) SetOwnerCodeList(v []string) {
	o.OwnerCodeList = v
}

func (o LinkedActivitiesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkedActivitiesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ActivityId) {
		toSerialize["activityId"] = o.ActivityId
	}
	if !IsNil(o.OwnerCodeList) {
		toSerialize["ownerCodeList"] = o.OwnerCodeList
	}
	return toSerialize, nil
}

type NullableLinkedActivitiesType struct {
	value *LinkedActivitiesType
	isSet bool
}

func (v NullableLinkedActivitiesType) Get() *LinkedActivitiesType {
	return v.value
}

func (v *NullableLinkedActivitiesType) Set(val *LinkedActivitiesType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedActivitiesType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedActivitiesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedActivitiesType(val *LinkedActivitiesType) *NullableLinkedActivitiesType {
	return &NullableLinkedActivitiesType{value: val, isSet: true}
}

func (v NullableLinkedActivitiesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedActivitiesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


