/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OverrideInstructionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideInstructionType{}

// OverrideInstructionType Type for Overrides. Contains information for the override action performed while booking a reservation.
type OverrideInstructionType struct {
	// The description of the restriction for which the override was done.
	Description *string `json:"description,omitempty"`
	// The date when the override was done.
	Date *string `json:"date,omitempty"`
	// The type of override done. If done for Availability, then it will be AVAILABILITY.
	Type *string `json:"type,omitempty"`
	// Login ID of the user who performed the override.
	UserId *string `json:"userId,omitempty"`
}

// NewOverrideInstructionType instantiates a new OverrideInstructionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideInstructionType() *OverrideInstructionType {
	this := OverrideInstructionType{}
	return &this
}

// NewOverrideInstructionTypeWithDefaults instantiates a new OverrideInstructionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideInstructionTypeWithDefaults() *OverrideInstructionType {
	this := OverrideInstructionType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OverrideInstructionType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideInstructionType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OverrideInstructionType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OverrideInstructionType) SetDescription(v string) {
	o.Description = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *OverrideInstructionType) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideInstructionType) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *OverrideInstructionType) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *OverrideInstructionType) SetDate(v string) {
	o.Date = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OverrideInstructionType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideInstructionType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OverrideInstructionType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OverrideInstructionType) SetType(v string) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OverrideInstructionType) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideInstructionType) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OverrideInstructionType) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OverrideInstructionType) SetUserId(v string) {
	o.UserId = &v
}

func (o OverrideInstructionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideInstructionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableOverrideInstructionType struct {
	value *OverrideInstructionType
	isSet bool
}

func (v NullableOverrideInstructionType) Get() *OverrideInstructionType {
	return v.value
}

func (v *NullableOverrideInstructionType) Set(val *OverrideInstructionType) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideInstructionType) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideInstructionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideInstructionType(val *OverrideInstructionType) *NullableOverrideInstructionType {
	return &NullableOverrideInstructionType{value: val, isSet: true}
}

func (v NullableOverrideInstructionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideInstructionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


