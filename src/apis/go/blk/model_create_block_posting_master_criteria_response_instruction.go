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

// checks if the CreateBlockPostingMasterCriteriaResponseInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBlockPostingMasterCriteriaResponseInstruction{}

// CreateBlockPostingMasterCriteriaResponseInstruction Dictates the response structure of the returned posting master reservation object.
type CreateBlockPostingMasterCriteriaResponseInstruction struct {
	// Whether the response should contain only the ReservationID or full posting master reservation info.
	FetchFullReservation *bool `json:"fetchFullReservation,omitempty"`
}

// NewCreateBlockPostingMasterCriteriaResponseInstruction instantiates a new CreateBlockPostingMasterCriteriaResponseInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBlockPostingMasterCriteriaResponseInstruction() *CreateBlockPostingMasterCriteriaResponseInstruction {
	this := CreateBlockPostingMasterCriteriaResponseInstruction{}
	return &this
}

// NewCreateBlockPostingMasterCriteriaResponseInstructionWithDefaults instantiates a new CreateBlockPostingMasterCriteriaResponseInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBlockPostingMasterCriteriaResponseInstructionWithDefaults() *CreateBlockPostingMasterCriteriaResponseInstruction {
	this := CreateBlockPostingMasterCriteriaResponseInstruction{}
	return &this
}

// GetFetchFullReservation returns the FetchFullReservation field value if set, zero value otherwise.
func (o *CreateBlockPostingMasterCriteriaResponseInstruction) GetFetchFullReservation() bool {
	if o == nil || IsNil(o.FetchFullReservation) {
		var ret bool
		return ret
	}
	return *o.FetchFullReservation
}

// GetFetchFullReservationOk returns a tuple with the FetchFullReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBlockPostingMasterCriteriaResponseInstruction) GetFetchFullReservationOk() (*bool, bool) {
	if o == nil || IsNil(o.FetchFullReservation) {
		return nil, false
	}
	return o.FetchFullReservation, true
}

// HasFetchFullReservation returns a boolean if a field has been set.
func (o *CreateBlockPostingMasterCriteriaResponseInstruction) HasFetchFullReservation() bool {
	if o != nil && !IsNil(o.FetchFullReservation) {
		return true
	}

	return false
}

// SetFetchFullReservation gets a reference to the given bool and assigns it to the FetchFullReservation field.
func (o *CreateBlockPostingMasterCriteriaResponseInstruction) SetFetchFullReservation(v bool) {
	o.FetchFullReservation = &v
}

func (o CreateBlockPostingMasterCriteriaResponseInstruction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBlockPostingMasterCriteriaResponseInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FetchFullReservation) {
		toSerialize["fetchFullReservation"] = o.FetchFullReservation
	}
	return toSerialize, nil
}

type NullableCreateBlockPostingMasterCriteriaResponseInstruction struct {
	value *CreateBlockPostingMasterCriteriaResponseInstruction
	isSet bool
}

func (v NullableCreateBlockPostingMasterCriteriaResponseInstruction) Get() *CreateBlockPostingMasterCriteriaResponseInstruction {
	return v.value
}

func (v *NullableCreateBlockPostingMasterCriteriaResponseInstruction) Set(val *CreateBlockPostingMasterCriteriaResponseInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBlockPostingMasterCriteriaResponseInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBlockPostingMasterCriteriaResponseInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBlockPostingMasterCriteriaResponseInstruction(val *CreateBlockPostingMasterCriteriaResponseInstruction) *NullableCreateBlockPostingMasterCriteriaResponseInstruction {
	return &NullableCreateBlockPostingMasterCriteriaResponseInstruction{value: val, isSet: true}
}

func (v NullableCreateBlockPostingMasterCriteriaResponseInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBlockPostingMasterCriteriaResponseInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


