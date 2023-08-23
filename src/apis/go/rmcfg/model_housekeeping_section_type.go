/*
OPERA Cloud Room Configuration API

APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HousekeepingSectionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HousekeepingSectionType{}

// HousekeepingSectionType Information regarding Housekeeping Sections.
type HousekeepingSectionType struct {
	// Property where this section is defined.
	HotelId *string `json:"hotelId,omitempty"`
	// The Housekeeping Section Code.
	Code *string `json:"code,omitempty"`
	// Description of the Section.
	Description *string `json:"description,omitempty"`
	// Housekeeping Section Group.
	SectionGroup *string `json:"sectionGroup,omitempty"`
	// Target Credit for each task sheet created for this section when auto task assignment is broken out by section.
	TargetCredits *int32 `json:"targetCredits,omitempty"`
	// Rooms count for this section code. This is auto populated while creating/fetching record.
	Rooms *int32 `json:"rooms,omitempty"`
	// Number of housekeeping credits. This is auto populated while creating/fetching record.
	RoomCredits *int32 `json:"roomCredits,omitempty"`
	// Display sequence when task assignment is automatically broken out by Section Group.
	Sequence *int32 `json:"sequence,omitempty"`
	// Indicates if the Section Code is active.
	Inactive *bool `json:"inactive,omitempty"`
}

// NewHousekeepingSectionType instantiates a new HousekeepingSectionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHousekeepingSectionType() *HousekeepingSectionType {
	this := HousekeepingSectionType{}
	return &this
}

// NewHousekeepingSectionTypeWithDefaults instantiates a new HousekeepingSectionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHousekeepingSectionTypeWithDefaults() *HousekeepingSectionType {
	this := HousekeepingSectionType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *HousekeepingSectionType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *HousekeepingSectionType) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HousekeepingSectionType) SetDescription(v string) {
	o.Description = &v
}

// GetSectionGroup returns the SectionGroup field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetSectionGroup() string {
	if o == nil || IsNil(o.SectionGroup) {
		var ret string
		return ret
	}
	return *o.SectionGroup
}

// GetSectionGroupOk returns a tuple with the SectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetSectionGroupOk() (*string, bool) {
	if o == nil || IsNil(o.SectionGroup) {
		return nil, false
	}
	return o.SectionGroup, true
}

// HasSectionGroup returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasSectionGroup() bool {
	if o != nil && !IsNil(o.SectionGroup) {
		return true
	}

	return false
}

// SetSectionGroup gets a reference to the given string and assigns it to the SectionGroup field.
func (o *HousekeepingSectionType) SetSectionGroup(v string) {
	o.SectionGroup = &v
}

// GetTargetCredits returns the TargetCredits field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetTargetCredits() int32 {
	if o == nil || IsNil(o.TargetCredits) {
		var ret int32
		return ret
	}
	return *o.TargetCredits
}

// GetTargetCreditsOk returns a tuple with the TargetCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetTargetCreditsOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetCredits) {
		return nil, false
	}
	return o.TargetCredits, true
}

// HasTargetCredits returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasTargetCredits() bool {
	if o != nil && !IsNil(o.TargetCredits) {
		return true
	}

	return false
}

// SetTargetCredits gets a reference to the given int32 and assigns it to the TargetCredits field.
func (o *HousekeepingSectionType) SetTargetCredits(v int32) {
	o.TargetCredits = &v
}

// GetRooms returns the Rooms field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetRooms() int32 {
	if o == nil || IsNil(o.Rooms) {
		var ret int32
		return ret
	}
	return *o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.Rooms) {
		return nil, false
	}
	return o.Rooms, true
}

// HasRooms returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasRooms() bool {
	if o != nil && !IsNil(o.Rooms) {
		return true
	}

	return false
}

// SetRooms gets a reference to the given int32 and assigns it to the Rooms field.
func (o *HousekeepingSectionType) SetRooms(v int32) {
	o.Rooms = &v
}

// GetRoomCredits returns the RoomCredits field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetRoomCredits() int32 {
	if o == nil || IsNil(o.RoomCredits) {
		var ret int32
		return ret
	}
	return *o.RoomCredits
}

// GetRoomCreditsOk returns a tuple with the RoomCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetRoomCreditsOk() (*int32, bool) {
	if o == nil || IsNil(o.RoomCredits) {
		return nil, false
	}
	return o.RoomCredits, true
}

// HasRoomCredits returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasRoomCredits() bool {
	if o != nil && !IsNil(o.RoomCredits) {
		return true
	}

	return false
}

// SetRoomCredits gets a reference to the given int32 and assigns it to the RoomCredits field.
func (o *HousekeepingSectionType) SetRoomCredits(v int32) {
	o.RoomCredits = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetSequence() int32 {
	if o == nil || IsNil(o.Sequence) {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetSequenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasSequence() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *HousekeepingSectionType) SetSequence(v int32) {
	o.Sequence = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *HousekeepingSectionType) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HousekeepingSectionType) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *HousekeepingSectionType) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *HousekeepingSectionType) SetInactive(v bool) {
	o.Inactive = &v
}

func (o HousekeepingSectionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HousekeepingSectionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SectionGroup) {
		toSerialize["sectionGroup"] = o.SectionGroup
	}
	if !IsNil(o.TargetCredits) {
		toSerialize["targetCredits"] = o.TargetCredits
	}
	if !IsNil(o.Rooms) {
		toSerialize["rooms"] = o.Rooms
	}
	if !IsNil(o.RoomCredits) {
		toSerialize["roomCredits"] = o.RoomCredits
	}
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	return toSerialize, nil
}

type NullableHousekeepingSectionType struct {
	value *HousekeepingSectionType
	isSet bool
}

func (v NullableHousekeepingSectionType) Get() *HousekeepingSectionType {
	return v.value
}

func (v *NullableHousekeepingSectionType) Set(val *HousekeepingSectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableHousekeepingSectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableHousekeepingSectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHousekeepingSectionType(val *HousekeepingSectionType) *NullableHousekeepingSectionType {
	return &NullableHousekeepingSectionType{value: val, isSet: true}
}

func (v NullableHousekeepingSectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHousekeepingSectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


