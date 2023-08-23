/*
OPERA Cloud Price Availability Rate API

APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MembershipSearchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipSearchType{}

// MembershipSearchType Identifies criteria for searching frequent customer reward program.
type MembershipSearchType struct {
	// Membership ID criteria.
	MembershipId *string `json:"membershipId,omitempty"`
	MembershipLevel []string `json:"membershipLevel,omitempty"`
	MembershipType []string `json:"membershipType,omitempty"`
	// If this is true,the reservations which has membership information associated will be resulted .
	AssociatedReservationsOnly *bool `json:"associatedReservationsOnly,omitempty"`
}

// NewMembershipSearchType instantiates a new MembershipSearchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipSearchType() *MembershipSearchType {
	this := MembershipSearchType{}
	return &this
}

// NewMembershipSearchTypeWithDefaults instantiates a new MembershipSearchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipSearchTypeWithDefaults() *MembershipSearchType {
	this := MembershipSearchType{}
	return &this
}

// GetMembershipId returns the MembershipId field value if set, zero value otherwise.
func (o *MembershipSearchType) GetMembershipId() string {
	if o == nil || IsNil(o.MembershipId) {
		var ret string
		return ret
	}
	return *o.MembershipId
}

// GetMembershipIdOk returns a tuple with the MembershipId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipSearchType) GetMembershipIdOk() (*string, bool) {
	if o == nil || IsNil(o.MembershipId) {
		return nil, false
	}
	return o.MembershipId, true
}

// HasMembershipId returns a boolean if a field has been set.
func (o *MembershipSearchType) HasMembershipId() bool {
	if o != nil && !IsNil(o.MembershipId) {
		return true
	}

	return false
}

// SetMembershipId gets a reference to the given string and assigns it to the MembershipId field.
func (o *MembershipSearchType) SetMembershipId(v string) {
	o.MembershipId = &v
}

// GetMembershipLevel returns the MembershipLevel field value if set, zero value otherwise.
func (o *MembershipSearchType) GetMembershipLevel() []string {
	if o == nil || IsNil(o.MembershipLevel) {
		var ret []string
		return ret
	}
	return o.MembershipLevel
}

// GetMembershipLevelOk returns a tuple with the MembershipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipSearchType) GetMembershipLevelOk() ([]string, bool) {
	if o == nil || IsNil(o.MembershipLevel) {
		return nil, false
	}
	return o.MembershipLevel, true
}

// HasMembershipLevel returns a boolean if a field has been set.
func (o *MembershipSearchType) HasMembershipLevel() bool {
	if o != nil && !IsNil(o.MembershipLevel) {
		return true
	}

	return false
}

// SetMembershipLevel gets a reference to the given []string and assigns it to the MembershipLevel field.
func (o *MembershipSearchType) SetMembershipLevel(v []string) {
	o.MembershipLevel = v
}

// GetMembershipType returns the MembershipType field value if set, zero value otherwise.
func (o *MembershipSearchType) GetMembershipType() []string {
	if o == nil || IsNil(o.MembershipType) {
		var ret []string
		return ret
	}
	return o.MembershipType
}

// GetMembershipTypeOk returns a tuple with the MembershipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipSearchType) GetMembershipTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.MembershipType) {
		return nil, false
	}
	return o.MembershipType, true
}

// HasMembershipType returns a boolean if a field has been set.
func (o *MembershipSearchType) HasMembershipType() bool {
	if o != nil && !IsNil(o.MembershipType) {
		return true
	}

	return false
}

// SetMembershipType gets a reference to the given []string and assigns it to the MembershipType field.
func (o *MembershipSearchType) SetMembershipType(v []string) {
	o.MembershipType = v
}

// GetAssociatedReservationsOnly returns the AssociatedReservationsOnly field value if set, zero value otherwise.
func (o *MembershipSearchType) GetAssociatedReservationsOnly() bool {
	if o == nil || IsNil(o.AssociatedReservationsOnly) {
		var ret bool
		return ret
	}
	return *o.AssociatedReservationsOnly
}

// GetAssociatedReservationsOnlyOk returns a tuple with the AssociatedReservationsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipSearchType) GetAssociatedReservationsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.AssociatedReservationsOnly) {
		return nil, false
	}
	return o.AssociatedReservationsOnly, true
}

// HasAssociatedReservationsOnly returns a boolean if a field has been set.
func (o *MembershipSearchType) HasAssociatedReservationsOnly() bool {
	if o != nil && !IsNil(o.AssociatedReservationsOnly) {
		return true
	}

	return false
}

// SetAssociatedReservationsOnly gets a reference to the given bool and assigns it to the AssociatedReservationsOnly field.
func (o *MembershipSearchType) SetAssociatedReservationsOnly(v bool) {
	o.AssociatedReservationsOnly = &v
}

func (o MembershipSearchType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipSearchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MembershipId) {
		toSerialize["membershipId"] = o.MembershipId
	}
	if !IsNil(o.MembershipLevel) {
		toSerialize["membershipLevel"] = o.MembershipLevel
	}
	if !IsNil(o.MembershipType) {
		toSerialize["membershipType"] = o.MembershipType
	}
	if !IsNil(o.AssociatedReservationsOnly) {
		toSerialize["associatedReservationsOnly"] = o.AssociatedReservationsOnly
	}
	return toSerialize, nil
}

type NullableMembershipSearchType struct {
	value *MembershipSearchType
	isSet bool
}

func (v NullableMembershipSearchType) Get() *MembershipSearchType {
	return v.value
}

func (v *NullableMembershipSearchType) Set(val *MembershipSearchType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipSearchType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipSearchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipSearchType(val *MembershipSearchType) *NullableMembershipSearchType {
	return &NullableMembershipSearchType{value: val, isSet: true}
}

func (v NullableMembershipSearchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipSearchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


