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

// checks if the BlockReservationsToChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockReservationsToChange{}

// BlockReservationsToChange The request object to implement batch changes to block reservations based upon changes made to a reference reservation.
type BlockReservationsToChange struct {
	Criteria *ChangeBlockReservationsType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBlockReservationsToChange instantiates a new BlockReservationsToChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockReservationsToChange() *BlockReservationsToChange {
	this := BlockReservationsToChange{}
	return &this
}

// NewBlockReservationsToChangeWithDefaults instantiates a new BlockReservationsToChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockReservationsToChangeWithDefaults() *BlockReservationsToChange {
	this := BlockReservationsToChange{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *BlockReservationsToChange) GetCriteria() ChangeBlockReservationsType {
	if o == nil || IsNil(o.Criteria) {
		var ret ChangeBlockReservationsType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockReservationsToChange) GetCriteriaOk() (*ChangeBlockReservationsType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *BlockReservationsToChange) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ChangeBlockReservationsType and assigns it to the Criteria field.
func (o *BlockReservationsToChange) SetCriteria(v ChangeBlockReservationsType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BlockReservationsToChange) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockReservationsToChange) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BlockReservationsToChange) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BlockReservationsToChange) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BlockReservationsToChange) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockReservationsToChange) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BlockReservationsToChange) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BlockReservationsToChange) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o BlockReservationsToChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockReservationsToChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBlockReservationsToChange struct {
	value *BlockReservationsToChange
	isSet bool
}

func (v NullableBlockReservationsToChange) Get() *BlockReservationsToChange {
	return v.value
}

func (v *NullableBlockReservationsToChange) Set(val *BlockReservationsToChange) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockReservationsToChange) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockReservationsToChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockReservationsToChange(val *BlockReservationsToChange) *NullableBlockReservationsToChange {
	return &NullableBlockReservationsToChange{value: val, isSet: true}
}

func (v NullableBlockReservationsToChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockReservationsToChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


