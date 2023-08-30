/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the CurrentRoomInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentRoomInfoType{}

// CurrentRoomInfoType Room information of the reservation for the current day.
type CurrentRoomInfoType struct {
	// Current room type.
	RoomType *string `json:"roomType,omitempty"`
	// Current room number.
	RoomId *string `json:"roomId,omitempty"`
	SuggestedRoomNumbers []string `json:"suggestedRoomNumbers,omitempty"`
	// Current room description.
	RoomDescription *string `json:"roomDescription,omitempty"`
	// Represents the room view code like City view, River view, Ocean view etc.
	RoomViewCode *string `json:"roomViewCode,omitempty"`
	// Represents the room was assigned by AI Room Assignment.
	AssignedByAI *bool `json:"assignedByAI,omitempty"`
	// Represents the room was upgraded by AI Room Assignment.
	UpgradedByAI *bool `json:"upgradedByAI,omitempty"`
}

// NewCurrentRoomInfoType instantiates a new CurrentRoomInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentRoomInfoType() *CurrentRoomInfoType {
	this := CurrentRoomInfoType{}
	return &this
}

// NewCurrentRoomInfoTypeWithDefaults instantiates a new CurrentRoomInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentRoomInfoTypeWithDefaults() *CurrentRoomInfoType {
	this := CurrentRoomInfoType{}
	return &this
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *CurrentRoomInfoType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentRoomInfoType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *CurrentRoomInfoType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *CurrentRoomInfoType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetRoomId returns the RoomId field value if set, zero value otherwise.
func (o *CurrentRoomInfoType) GetRoomId() string {
	if o == nil || IsNil(o.RoomId) {
		var ret string
		return ret
	}
	return *o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentRoomInfoType) GetRoomIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoomId) {
		return nil, false
	}
	return o.RoomId, true
}

// HasRoomId returns a boolean if a field has been set.
func (o *CurrentRoomInfoType) HasRoomId() bool {
	if o != nil && !IsNil(o.RoomId) {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given string and assigns it to the RoomId field.
func (o *CurrentRoomInfoType) SetRoomId(v string) {
	o.RoomId = &v
}

// GetSuggestedRoomNumbers returns the SuggestedRoomNumbers field value if set, zero value otherwise.
func (o *CurrentRoomInfoType) GetSuggestedRoomNumbers() []string {
	if o == nil || IsNil(o.SuggestedRoomNumbers) {
		var ret []string
		return ret
	}
	return o.SuggestedRoomNumbers
}

// GetSuggestedRoomNumbersOk returns a tuple with the SuggestedRoomNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentRoomInfoType) GetSuggestedRoomNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.SuggestedRoomNumbers) {
		return nil, false
	}
	return o.SuggestedRoomNumbers, true
}

// HasSuggestedRoomNumbers returns a boolean if a field has been set.
func (o *CurrentRoomInfoType) HasSuggestedRoomNumbers() bool {
	if o != nil && !IsNil(o.SuggestedRoomNumbers) {
		return true
	}

	return false
}

// SetSuggestedRoomNumbers gets a reference to the given []string and assigns it to the SuggestedRoomNumbers field.
func (o *CurrentRoomInfoType) SetSuggestedRoomNumbers(v []string) {
	o.SuggestedRoomNumbers = v
}

// GetRoomDescription returns the RoomDescription field value if set, zero value otherwise.
func (o *CurrentRoomInfoType) GetRoomDescription() string {
	if o == nil || IsNil(o.RoomDescription) {
		var ret string
		return ret
	}
	return *o.RoomDescription
}

// GetRoomDescriptionOk returns a tuple with the RoomDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentRoomInfoType) GetRoomDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.RoomDescription) {
		return nil, false
	}
	return o.RoomDescription, true
}

// HasRoomDescription returns a boolean if a field has been set.
func (o *CurrentRoomInfoType) HasRoomDescription() bool {
	if o != nil && !IsNil(o.RoomDescription) {
		return true
	}

	return false
}

// SetRoomDescription gets a reference to the given string and assigns it to the RoomDescription field.
func (o *CurrentRoomInfoType) SetRoomDescription(v string) {
	o.RoomDescription = &v
}

// GetRoomViewCode returns the RoomViewCode field value if set, zero value otherwise.
func (o *CurrentRoomInfoType) GetRoomViewCode() string {
	if o == nil || IsNil(o.RoomViewCode) {
		var ret string
		return ret
	}
	return *o.RoomViewCode
}

// GetRoomViewCodeOk returns a tuple with the RoomViewCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentRoomInfoType) GetRoomViewCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomViewCode) {
		return nil, false
	}
	return o.RoomViewCode, true
}

// HasRoomViewCode returns a boolean if a field has been set.
func (o *CurrentRoomInfoType) HasRoomViewCode() bool {
	if o != nil && !IsNil(o.RoomViewCode) {
		return true
	}

	return false
}

// SetRoomViewCode gets a reference to the given string and assigns it to the RoomViewCode field.
func (o *CurrentRoomInfoType) SetRoomViewCode(v string) {
	o.RoomViewCode = &v
}

// GetAssignedByAI returns the AssignedByAI field value if set, zero value otherwise.
func (o *CurrentRoomInfoType) GetAssignedByAI() bool {
	if o == nil || IsNil(o.AssignedByAI) {
		var ret bool
		return ret
	}
	return *o.AssignedByAI
}

// GetAssignedByAIOk returns a tuple with the AssignedByAI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentRoomInfoType) GetAssignedByAIOk() (*bool, bool) {
	if o == nil || IsNil(o.AssignedByAI) {
		return nil, false
	}
	return o.AssignedByAI, true
}

// HasAssignedByAI returns a boolean if a field has been set.
func (o *CurrentRoomInfoType) HasAssignedByAI() bool {
	if o != nil && !IsNil(o.AssignedByAI) {
		return true
	}

	return false
}

// SetAssignedByAI gets a reference to the given bool and assigns it to the AssignedByAI field.
func (o *CurrentRoomInfoType) SetAssignedByAI(v bool) {
	o.AssignedByAI = &v
}

// GetUpgradedByAI returns the UpgradedByAI field value if set, zero value otherwise.
func (o *CurrentRoomInfoType) GetUpgradedByAI() bool {
	if o == nil || IsNil(o.UpgradedByAI) {
		var ret bool
		return ret
	}
	return *o.UpgradedByAI
}

// GetUpgradedByAIOk returns a tuple with the UpgradedByAI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentRoomInfoType) GetUpgradedByAIOk() (*bool, bool) {
	if o == nil || IsNil(o.UpgradedByAI) {
		return nil, false
	}
	return o.UpgradedByAI, true
}

// HasUpgradedByAI returns a boolean if a field has been set.
func (o *CurrentRoomInfoType) HasUpgradedByAI() bool {
	if o != nil && !IsNil(o.UpgradedByAI) {
		return true
	}

	return false
}

// SetUpgradedByAI gets a reference to the given bool and assigns it to the UpgradedByAI field.
func (o *CurrentRoomInfoType) SetUpgradedByAI(v bool) {
	o.UpgradedByAI = &v
}

func (o CurrentRoomInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentRoomInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.RoomId) {
		toSerialize["roomId"] = o.RoomId
	}
	if !IsNil(o.SuggestedRoomNumbers) {
		toSerialize["suggestedRoomNumbers"] = o.SuggestedRoomNumbers
	}
	if !IsNil(o.RoomDescription) {
		toSerialize["roomDescription"] = o.RoomDescription
	}
	if !IsNil(o.RoomViewCode) {
		toSerialize["roomViewCode"] = o.RoomViewCode
	}
	if !IsNil(o.AssignedByAI) {
		toSerialize["assignedByAI"] = o.AssignedByAI
	}
	if !IsNil(o.UpgradedByAI) {
		toSerialize["upgradedByAI"] = o.UpgradedByAI
	}
	return toSerialize, nil
}

type NullableCurrentRoomInfoType struct {
	value *CurrentRoomInfoType
	isSet bool
}

func (v NullableCurrentRoomInfoType) Get() *CurrentRoomInfoType {
	return v.value
}

func (v *NullableCurrentRoomInfoType) Set(val *CurrentRoomInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentRoomInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentRoomInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentRoomInfoType(val *CurrentRoomInfoType) *NullableCurrentRoomInfoType {
	return &NullableCurrentRoomInfoType{value: val, isSet: true}
}

func (v NullableCurrentRoomInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentRoomInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


