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

// checks if the ChangedBlockReservations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangedBlockReservations{}

// ChangedBlockReservations The response object that contains the reservations that were changed for a block along with the status of each change.
type ChangedBlockReservations struct {
	// Contains details of the changed reservation along with a success or error message.
	Reservations []ChangeBlockReservationType `json:"reservations,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChangedBlockReservations instantiates a new ChangedBlockReservations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangedBlockReservations() *ChangedBlockReservations {
	this := ChangedBlockReservations{}
	return &this
}

// NewChangedBlockReservationsWithDefaults instantiates a new ChangedBlockReservations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangedBlockReservationsWithDefaults() *ChangedBlockReservations {
	this := ChangedBlockReservations{}
	return &this
}

// GetReservations returns the Reservations field value if set, zero value otherwise.
func (o *ChangedBlockReservations) GetReservations() []ChangeBlockReservationType {
	if o == nil || IsNil(o.Reservations) {
		var ret []ChangeBlockReservationType
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedBlockReservations) GetReservationsOk() ([]ChangeBlockReservationType, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *ChangedBlockReservations) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []ChangeBlockReservationType and assigns it to the Reservations field.
func (o *ChangedBlockReservations) SetReservations(v []ChangeBlockReservationType) {
	o.Reservations = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChangedBlockReservations) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedBlockReservations) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChangedBlockReservations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChangedBlockReservations) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChangedBlockReservations) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedBlockReservations) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChangedBlockReservations) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChangedBlockReservations) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChangedBlockReservations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangedBlockReservations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reservations) {
		toSerialize["reservations"] = o.Reservations
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChangedBlockReservations struct {
	value *ChangedBlockReservations
	isSet bool
}

func (v NullableChangedBlockReservations) Get() *ChangedBlockReservations {
	return v.value
}

func (v *NullableChangedBlockReservations) Set(val *ChangedBlockReservations) {
	v.value = val
	v.isSet = true
}

func (v NullableChangedBlockReservations) IsSet() bool {
	return v.isSet
}

func (v *NullableChangedBlockReservations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangedBlockReservations(val *ChangedBlockReservations) *NullableChangedBlockReservations {
	return &NullableChangedBlockReservations{value: val, isSet: true}
}

func (v NullableChangedBlockReservations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangedBlockReservations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


