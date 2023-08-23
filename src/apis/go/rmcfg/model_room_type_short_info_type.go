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

// checks if the RoomTypeShortInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomTypeShortInfoType{}

// RoomTypeShortInfoType Basic information of room type.
type RoomTypeShortInfoType struct {
	// Indicates if room is a pseudo. This is read-only.
	Pseudo *bool `json:"pseudo,omitempty"`
	// Indicates if room is a suite. This is read-only.
	Suite *bool `json:"suite,omitempty"`
	// Room Class of the room. This is read-only.
	RoomClass *string `json:"roomClass,omitempty"`
	// Short Description of room type.
	ShortDescription *string `json:"shortDescription,omitempty"`
	// Indicates if room type of the room is flagged as housekeeping. This is read-only.
	HouseKeeping *bool `json:"houseKeeping,omitempty"`
	// Specifies the smoking preference for room type of the room.
	SmokingPreference *string `json:"smokingPreference,omitempty"`
	// Building associated with the room.
	Building *string `json:"building,omitempty"`
	RoomAssignmentRating *RatePlanRatingType `json:"roomAssignmentRating,omitempty"`
	// Minimum occupancy for the room type.
	MinimumOccupancy *int32 `json:"minimumOccupancy,omitempty"`
	// Maximum occupancy for the room type.
	MaximumOccupancy *int32 `json:"maximumOccupancy,omitempty"`
	// A recurring element that identifies the room features.
	RoomFeatures []RoomFeatureType `json:"roomFeatures,omitempty"`
	// Indicates if room type is accessible.
	Accessible *bool `json:"accessible,omitempty"`
	// Room type of the room.
	RoomType *string `json:"roomType,omitempty"`
	// Indicates whether function space can be meeting room.
	MeetingRoom *bool `json:"meetingRoom,omitempty"`
}

// NewRoomTypeShortInfoType instantiates a new RoomTypeShortInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomTypeShortInfoType() *RoomTypeShortInfoType {
	this := RoomTypeShortInfoType{}
	return &this
}

// NewRoomTypeShortInfoTypeWithDefaults instantiates a new RoomTypeShortInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomTypeShortInfoTypeWithDefaults() *RoomTypeShortInfoType {
	this := RoomTypeShortInfoType{}
	return &this
}

// GetPseudo returns the Pseudo field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetPseudo() bool {
	if o == nil || IsNil(o.Pseudo) {
		var ret bool
		return ret
	}
	return *o.Pseudo
}

// GetPseudoOk returns a tuple with the Pseudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetPseudoOk() (*bool, bool) {
	if o == nil || IsNil(o.Pseudo) {
		return nil, false
	}
	return o.Pseudo, true
}

// HasPseudo returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasPseudo() bool {
	if o != nil && !IsNil(o.Pseudo) {
		return true
	}

	return false
}

// SetPseudo gets a reference to the given bool and assigns it to the Pseudo field.
func (o *RoomTypeShortInfoType) SetPseudo(v bool) {
	o.Pseudo = &v
}

// GetSuite returns the Suite field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetSuite() bool {
	if o == nil || IsNil(o.Suite) {
		var ret bool
		return ret
	}
	return *o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetSuiteOk() (*bool, bool) {
	if o == nil || IsNil(o.Suite) {
		return nil, false
	}
	return o.Suite, true
}

// HasSuite returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasSuite() bool {
	if o != nil && !IsNil(o.Suite) {
		return true
	}

	return false
}

// SetSuite gets a reference to the given bool and assigns it to the Suite field.
func (o *RoomTypeShortInfoType) SetSuite(v bool) {
	o.Suite = &v
}

// GetRoomClass returns the RoomClass field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetRoomClass() string {
	if o == nil || IsNil(o.RoomClass) {
		var ret string
		return ret
	}
	return *o.RoomClass
}

// GetRoomClassOk returns a tuple with the RoomClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetRoomClassOk() (*string, bool) {
	if o == nil || IsNil(o.RoomClass) {
		return nil, false
	}
	return o.RoomClass, true
}

// HasRoomClass returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasRoomClass() bool {
	if o != nil && !IsNil(o.RoomClass) {
		return true
	}

	return false
}

// SetRoomClass gets a reference to the given string and assigns it to the RoomClass field.
func (o *RoomTypeShortInfoType) SetRoomClass(v string) {
	o.RoomClass = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *RoomTypeShortInfoType) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetHouseKeeping returns the HouseKeeping field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetHouseKeeping() bool {
	if o == nil || IsNil(o.HouseKeeping) {
		var ret bool
		return ret
	}
	return *o.HouseKeeping
}

// GetHouseKeepingOk returns a tuple with the HouseKeeping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetHouseKeepingOk() (*bool, bool) {
	if o == nil || IsNil(o.HouseKeeping) {
		return nil, false
	}
	return o.HouseKeeping, true
}

// HasHouseKeeping returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasHouseKeeping() bool {
	if o != nil && !IsNil(o.HouseKeeping) {
		return true
	}

	return false
}

// SetHouseKeeping gets a reference to the given bool and assigns it to the HouseKeeping field.
func (o *RoomTypeShortInfoType) SetHouseKeeping(v bool) {
	o.HouseKeeping = &v
}

// GetSmokingPreference returns the SmokingPreference field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetSmokingPreference() string {
	if o == nil || IsNil(o.SmokingPreference) {
		var ret string
		return ret
	}
	return *o.SmokingPreference
}

// GetSmokingPreferenceOk returns a tuple with the SmokingPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetSmokingPreferenceOk() (*string, bool) {
	if o == nil || IsNil(o.SmokingPreference) {
		return nil, false
	}
	return o.SmokingPreference, true
}

// HasSmokingPreference returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasSmokingPreference() bool {
	if o != nil && !IsNil(o.SmokingPreference) {
		return true
	}

	return false
}

// SetSmokingPreference gets a reference to the given string and assigns it to the SmokingPreference field.
func (o *RoomTypeShortInfoType) SetSmokingPreference(v string) {
	o.SmokingPreference = &v
}

// GetBuilding returns the Building field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetBuilding() string {
	if o == nil || IsNil(o.Building) {
		var ret string
		return ret
	}
	return *o.Building
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetBuildingOk() (*string, bool) {
	if o == nil || IsNil(o.Building) {
		return nil, false
	}
	return o.Building, true
}

// HasBuilding returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasBuilding() bool {
	if o != nil && !IsNil(o.Building) {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given string and assigns it to the Building field.
func (o *RoomTypeShortInfoType) SetBuilding(v string) {
	o.Building = &v
}

// GetRoomAssignmentRating returns the RoomAssignmentRating field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetRoomAssignmentRating() RatePlanRatingType {
	if o == nil || IsNil(o.RoomAssignmentRating) {
		var ret RatePlanRatingType
		return ret
	}
	return *o.RoomAssignmentRating
}

// GetRoomAssignmentRatingOk returns a tuple with the RoomAssignmentRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetRoomAssignmentRatingOk() (*RatePlanRatingType, bool) {
	if o == nil || IsNil(o.RoomAssignmentRating) {
		return nil, false
	}
	return o.RoomAssignmentRating, true
}

// HasRoomAssignmentRating returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasRoomAssignmentRating() bool {
	if o != nil && !IsNil(o.RoomAssignmentRating) {
		return true
	}

	return false
}

// SetRoomAssignmentRating gets a reference to the given RatePlanRatingType and assigns it to the RoomAssignmentRating field.
func (o *RoomTypeShortInfoType) SetRoomAssignmentRating(v RatePlanRatingType) {
	o.RoomAssignmentRating = &v
}

// GetMinimumOccupancy returns the MinimumOccupancy field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetMinimumOccupancy() int32 {
	if o == nil || IsNil(o.MinimumOccupancy) {
		var ret int32
		return ret
	}
	return *o.MinimumOccupancy
}

// GetMinimumOccupancyOk returns a tuple with the MinimumOccupancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetMinimumOccupancyOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumOccupancy) {
		return nil, false
	}
	return o.MinimumOccupancy, true
}

// HasMinimumOccupancy returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasMinimumOccupancy() bool {
	if o != nil && !IsNil(o.MinimumOccupancy) {
		return true
	}

	return false
}

// SetMinimumOccupancy gets a reference to the given int32 and assigns it to the MinimumOccupancy field.
func (o *RoomTypeShortInfoType) SetMinimumOccupancy(v int32) {
	o.MinimumOccupancy = &v
}

// GetMaximumOccupancy returns the MaximumOccupancy field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetMaximumOccupancy() int32 {
	if o == nil || IsNil(o.MaximumOccupancy) {
		var ret int32
		return ret
	}
	return *o.MaximumOccupancy
}

// GetMaximumOccupancyOk returns a tuple with the MaximumOccupancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetMaximumOccupancyOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumOccupancy) {
		return nil, false
	}
	return o.MaximumOccupancy, true
}

// HasMaximumOccupancy returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasMaximumOccupancy() bool {
	if o != nil && !IsNil(o.MaximumOccupancy) {
		return true
	}

	return false
}

// SetMaximumOccupancy gets a reference to the given int32 and assigns it to the MaximumOccupancy field.
func (o *RoomTypeShortInfoType) SetMaximumOccupancy(v int32) {
	o.MaximumOccupancy = &v
}

// GetRoomFeatures returns the RoomFeatures field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetRoomFeatures() []RoomFeatureType {
	if o == nil || IsNil(o.RoomFeatures) {
		var ret []RoomFeatureType
		return ret
	}
	return o.RoomFeatures
}

// GetRoomFeaturesOk returns a tuple with the RoomFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetRoomFeaturesOk() ([]RoomFeatureType, bool) {
	if o == nil || IsNil(o.RoomFeatures) {
		return nil, false
	}
	return o.RoomFeatures, true
}

// HasRoomFeatures returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasRoomFeatures() bool {
	if o != nil && !IsNil(o.RoomFeatures) {
		return true
	}

	return false
}

// SetRoomFeatures gets a reference to the given []RoomFeatureType and assigns it to the RoomFeatures field.
func (o *RoomTypeShortInfoType) SetRoomFeatures(v []RoomFeatureType) {
	o.RoomFeatures = v
}

// GetAccessible returns the Accessible field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetAccessible() bool {
	if o == nil || IsNil(o.Accessible) {
		var ret bool
		return ret
	}
	return *o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetAccessibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Accessible) {
		return nil, false
	}
	return o.Accessible, true
}

// HasAccessible returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasAccessible() bool {
	if o != nil && !IsNil(o.Accessible) {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given bool and assigns it to the Accessible field.
func (o *RoomTypeShortInfoType) SetAccessible(v bool) {
	o.Accessible = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *RoomTypeShortInfoType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetMeetingRoom returns the MeetingRoom field value if set, zero value otherwise.
func (o *RoomTypeShortInfoType) GetMeetingRoom() bool {
	if o == nil || IsNil(o.MeetingRoom) {
		var ret bool
		return ret
	}
	return *o.MeetingRoom
}

// GetMeetingRoomOk returns a tuple with the MeetingRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeShortInfoType) GetMeetingRoomOk() (*bool, bool) {
	if o == nil || IsNil(o.MeetingRoom) {
		return nil, false
	}
	return o.MeetingRoom, true
}

// HasMeetingRoom returns a boolean if a field has been set.
func (o *RoomTypeShortInfoType) HasMeetingRoom() bool {
	if o != nil && !IsNil(o.MeetingRoom) {
		return true
	}

	return false
}

// SetMeetingRoom gets a reference to the given bool and assigns it to the MeetingRoom field.
func (o *RoomTypeShortInfoType) SetMeetingRoom(v bool) {
	o.MeetingRoom = &v
}

func (o RoomTypeShortInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomTypeShortInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pseudo) {
		toSerialize["pseudo"] = o.Pseudo
	}
	if !IsNil(o.Suite) {
		toSerialize["suite"] = o.Suite
	}
	if !IsNil(o.RoomClass) {
		toSerialize["roomClass"] = o.RoomClass
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.HouseKeeping) {
		toSerialize["houseKeeping"] = o.HouseKeeping
	}
	if !IsNil(o.SmokingPreference) {
		toSerialize["smokingPreference"] = o.SmokingPreference
	}
	if !IsNil(o.Building) {
		toSerialize["building"] = o.Building
	}
	if !IsNil(o.RoomAssignmentRating) {
		toSerialize["roomAssignmentRating"] = o.RoomAssignmentRating
	}
	if !IsNil(o.MinimumOccupancy) {
		toSerialize["minimumOccupancy"] = o.MinimumOccupancy
	}
	if !IsNil(o.MaximumOccupancy) {
		toSerialize["maximumOccupancy"] = o.MaximumOccupancy
	}
	if !IsNil(o.RoomFeatures) {
		toSerialize["roomFeatures"] = o.RoomFeatures
	}
	if !IsNil(o.Accessible) {
		toSerialize["accessible"] = o.Accessible
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.MeetingRoom) {
		toSerialize["meetingRoom"] = o.MeetingRoom
	}
	return toSerialize, nil
}

type NullableRoomTypeShortInfoType struct {
	value *RoomTypeShortInfoType
	isSet bool
}

func (v NullableRoomTypeShortInfoType) Get() *RoomTypeShortInfoType {
	return v.value
}

func (v *NullableRoomTypeShortInfoType) Set(val *RoomTypeShortInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomTypeShortInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomTypeShortInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomTypeShortInfoType(val *RoomTypeShortInfoType) *NullableRoomTypeShortInfoType {
	return &NullableRoomTypeShortInfoType{value: val, isSet: true}
}

func (v NullableRoomTypeShortInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomTypeShortInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


