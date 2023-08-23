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

// checks if the ConfigRoomType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigRoomType{}

// ConfigRoomType This type represents the primary room attributes.
type ConfigRoomType struct {
	RoomType *RoomTypeShortInfoType `json:"roomType,omitempty"`
	// Floor of the Room.
	Floor *string `json:"floor,omitempty"`
	// Description for the Floor of the Room.
	FloorDescription *string `json:"floorDescription,omitempty"`
	// A recurring element that identifies the room features.
	RoomFeatures []RoomFeatureType `json:"roomFeatures,omitempty"`
	// Detail Long Description Of The Room.
	RoomDescription *string `json:"roomDescription,omitempty"`
	Description *TranslationTextType2000 `json:"description,omitempty"`
	// This indicates room smoking preference.
	SmokingPreference *string `json:"smokingPreference,omitempty"`
	// This indicates the description of the room smoking preference.
	SmokingPreferenceDescription *string `json:"smokingPreferenceDescription,omitempty"`
	// Building associated with the room.
	Building *string `json:"building,omitempty"`
	RoomAssignmentRating *RatePlanRatingType `json:"roomAssignmentRating,omitempty"`
	// Indicates whether the room is accessibility compliant.
	Accessible *bool `json:"accessible,omitempty"`
	// Code of the room.
	RoomId *string `json:"roomId,omitempty"`
	// Indicates whether the room is a Meeting Room
	MeetingRoom *bool `json:"meetingRoom,omitempty"`
	// Component of a room.
	RoomComponents []RoomComponentType `json:"roomComponents,omitempty"`
	// Collection of rooms.
	ConnectingRooms []RoomRoomType `json:"connectingRooms,omitempty"`
	// Published rate code of a room.
	RateCode *string `json:"rateCode,omitempty"`
	RateAmount *CurrencyAmountType `json:"rateAmount,omitempty"`
	// Maximum occupancy of a room.
	MaximumOccupancy *int32 `json:"maximumOccupancy,omitempty"`
	// Display sequence of a room.
	SellSequence *float32 `json:"sellSequence,omitempty"`
	// This indicates if room is marked as a owner room
	OwnerRoom *bool `json:"ownerRoom,omitempty"`
	// The Unit Grade Code attached to the room
	UnitGradeCode *string `json:"unitGradeCode,omitempty"`
	// This indicates if smoking is allowed in the room.
	Smoking *bool `json:"smoking,omitempty"`
	// Key code of a room.
	KeyCode *string `json:"keyCode,omitempty"`
	KeyOptions []string `json:"keyOptions,omitempty"`
	// Square units of a room.
	SquareUnits *float32 `json:"squareUnits,omitempty"`
	// Turndown service flag of a room.
	TurndownService *bool `json:"turndownService,omitempty"`
	// Square unit measurement of a room.
	SquareUnitsMeasurement *string `json:"squareUnitsMeasurement,omitempty"`
	// Square unit measurement of a room.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Housekeeping credits of the room. This may include stayover, departure, pickup and turndown credits.
	HousekeepingCredit []HousekeepingCreditsType `json:"housekeepingCredit,omitempty"`
	RoomSection *RoomSectionType `json:"roomSection,omitempty"`
	// Number of beds of the room.
	NoOfBeds *int32 `json:"noOfBeds,omitempty"`
}

// NewConfigRoomType instantiates a new ConfigRoomType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigRoomType() *ConfigRoomType {
	this := ConfigRoomType{}
	return &this
}

// NewConfigRoomTypeWithDefaults instantiates a new ConfigRoomType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigRoomTypeWithDefaults() *ConfigRoomType {
	this := ConfigRoomType{}
	return &this
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRoomType() RoomTypeShortInfoType {
	if o == nil || IsNil(o.RoomType) {
		var ret RoomTypeShortInfoType
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRoomTypeOk() (*RoomTypeShortInfoType, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given RoomTypeShortInfoType and assigns it to the RoomType field.
func (o *ConfigRoomType) SetRoomType(v RoomTypeShortInfoType) {
	o.RoomType = &v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *ConfigRoomType) GetFloor() string {
	if o == nil || IsNil(o.Floor) {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetFloorOk() (*string, bool) {
	if o == nil || IsNil(o.Floor) {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *ConfigRoomType) HasFloor() bool {
	if o != nil && !IsNil(o.Floor) {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *ConfigRoomType) SetFloor(v string) {
	o.Floor = &v
}

// GetFloorDescription returns the FloorDescription field value if set, zero value otherwise.
func (o *ConfigRoomType) GetFloorDescription() string {
	if o == nil || IsNil(o.FloorDescription) {
		var ret string
		return ret
	}
	return *o.FloorDescription
}

// GetFloorDescriptionOk returns a tuple with the FloorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetFloorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.FloorDescription) {
		return nil, false
	}
	return o.FloorDescription, true
}

// HasFloorDescription returns a boolean if a field has been set.
func (o *ConfigRoomType) HasFloorDescription() bool {
	if o != nil && !IsNil(o.FloorDescription) {
		return true
	}

	return false
}

// SetFloorDescription gets a reference to the given string and assigns it to the FloorDescription field.
func (o *ConfigRoomType) SetFloorDescription(v string) {
	o.FloorDescription = &v
}

// GetRoomFeatures returns the RoomFeatures field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRoomFeatures() []RoomFeatureType {
	if o == nil || IsNil(o.RoomFeatures) {
		var ret []RoomFeatureType
		return ret
	}
	return o.RoomFeatures
}

// GetRoomFeaturesOk returns a tuple with the RoomFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRoomFeaturesOk() ([]RoomFeatureType, bool) {
	if o == nil || IsNil(o.RoomFeatures) {
		return nil, false
	}
	return o.RoomFeatures, true
}

// HasRoomFeatures returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRoomFeatures() bool {
	if o != nil && !IsNil(o.RoomFeatures) {
		return true
	}

	return false
}

// SetRoomFeatures gets a reference to the given []RoomFeatureType and assigns it to the RoomFeatures field.
func (o *ConfigRoomType) SetRoomFeatures(v []RoomFeatureType) {
	o.RoomFeatures = v
}

// GetRoomDescription returns the RoomDescription field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRoomDescription() string {
	if o == nil || IsNil(o.RoomDescription) {
		var ret string
		return ret
	}
	return *o.RoomDescription
}

// GetRoomDescriptionOk returns a tuple with the RoomDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRoomDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.RoomDescription) {
		return nil, false
	}
	return o.RoomDescription, true
}

// HasRoomDescription returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRoomDescription() bool {
	if o != nil && !IsNil(o.RoomDescription) {
		return true
	}

	return false
}

// SetRoomDescription gets a reference to the given string and assigns it to the RoomDescription field.
func (o *ConfigRoomType) SetRoomDescription(v string) {
	o.RoomDescription = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigRoomType) GetDescription() TranslationTextType2000 {
	if o == nil || IsNil(o.Description) {
		var ret TranslationTextType2000
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetDescriptionOk() (*TranslationTextType2000, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigRoomType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given TranslationTextType2000 and assigns it to the Description field.
func (o *ConfigRoomType) SetDescription(v TranslationTextType2000) {
	o.Description = &v
}

// GetSmokingPreference returns the SmokingPreference field value if set, zero value otherwise.
func (o *ConfigRoomType) GetSmokingPreference() string {
	if o == nil || IsNil(o.SmokingPreference) {
		var ret string
		return ret
	}
	return *o.SmokingPreference
}

// GetSmokingPreferenceOk returns a tuple with the SmokingPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetSmokingPreferenceOk() (*string, bool) {
	if o == nil || IsNil(o.SmokingPreference) {
		return nil, false
	}
	return o.SmokingPreference, true
}

// HasSmokingPreference returns a boolean if a field has been set.
func (o *ConfigRoomType) HasSmokingPreference() bool {
	if o != nil && !IsNil(o.SmokingPreference) {
		return true
	}

	return false
}

// SetSmokingPreference gets a reference to the given string and assigns it to the SmokingPreference field.
func (o *ConfigRoomType) SetSmokingPreference(v string) {
	o.SmokingPreference = &v
}

// GetSmokingPreferenceDescription returns the SmokingPreferenceDescription field value if set, zero value otherwise.
func (o *ConfigRoomType) GetSmokingPreferenceDescription() string {
	if o == nil || IsNil(o.SmokingPreferenceDescription) {
		var ret string
		return ret
	}
	return *o.SmokingPreferenceDescription
}

// GetSmokingPreferenceDescriptionOk returns a tuple with the SmokingPreferenceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetSmokingPreferenceDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SmokingPreferenceDescription) {
		return nil, false
	}
	return o.SmokingPreferenceDescription, true
}

// HasSmokingPreferenceDescription returns a boolean if a field has been set.
func (o *ConfigRoomType) HasSmokingPreferenceDescription() bool {
	if o != nil && !IsNil(o.SmokingPreferenceDescription) {
		return true
	}

	return false
}

// SetSmokingPreferenceDescription gets a reference to the given string and assigns it to the SmokingPreferenceDescription field.
func (o *ConfigRoomType) SetSmokingPreferenceDescription(v string) {
	o.SmokingPreferenceDescription = &v
}

// GetBuilding returns the Building field value if set, zero value otherwise.
func (o *ConfigRoomType) GetBuilding() string {
	if o == nil || IsNil(o.Building) {
		var ret string
		return ret
	}
	return *o.Building
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetBuildingOk() (*string, bool) {
	if o == nil || IsNil(o.Building) {
		return nil, false
	}
	return o.Building, true
}

// HasBuilding returns a boolean if a field has been set.
func (o *ConfigRoomType) HasBuilding() bool {
	if o != nil && !IsNil(o.Building) {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given string and assigns it to the Building field.
func (o *ConfigRoomType) SetBuilding(v string) {
	o.Building = &v
}

// GetRoomAssignmentRating returns the RoomAssignmentRating field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRoomAssignmentRating() RatePlanRatingType {
	if o == nil || IsNil(o.RoomAssignmentRating) {
		var ret RatePlanRatingType
		return ret
	}
	return *o.RoomAssignmentRating
}

// GetRoomAssignmentRatingOk returns a tuple with the RoomAssignmentRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRoomAssignmentRatingOk() (*RatePlanRatingType, bool) {
	if o == nil || IsNil(o.RoomAssignmentRating) {
		return nil, false
	}
	return o.RoomAssignmentRating, true
}

// HasRoomAssignmentRating returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRoomAssignmentRating() bool {
	if o != nil && !IsNil(o.RoomAssignmentRating) {
		return true
	}

	return false
}

// SetRoomAssignmentRating gets a reference to the given RatePlanRatingType and assigns it to the RoomAssignmentRating field.
func (o *ConfigRoomType) SetRoomAssignmentRating(v RatePlanRatingType) {
	o.RoomAssignmentRating = &v
}

// GetAccessible returns the Accessible field value if set, zero value otherwise.
func (o *ConfigRoomType) GetAccessible() bool {
	if o == nil || IsNil(o.Accessible) {
		var ret bool
		return ret
	}
	return *o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetAccessibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Accessible) {
		return nil, false
	}
	return o.Accessible, true
}

// HasAccessible returns a boolean if a field has been set.
func (o *ConfigRoomType) HasAccessible() bool {
	if o != nil && !IsNil(o.Accessible) {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given bool and assigns it to the Accessible field.
func (o *ConfigRoomType) SetAccessible(v bool) {
	o.Accessible = &v
}

// GetRoomId returns the RoomId field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRoomId() string {
	if o == nil || IsNil(o.RoomId) {
		var ret string
		return ret
	}
	return *o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRoomIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoomId) {
		return nil, false
	}
	return o.RoomId, true
}

// HasRoomId returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRoomId() bool {
	if o != nil && !IsNil(o.RoomId) {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given string and assigns it to the RoomId field.
func (o *ConfigRoomType) SetRoomId(v string) {
	o.RoomId = &v
}

// GetMeetingRoom returns the MeetingRoom field value if set, zero value otherwise.
func (o *ConfigRoomType) GetMeetingRoom() bool {
	if o == nil || IsNil(o.MeetingRoom) {
		var ret bool
		return ret
	}
	return *o.MeetingRoom
}

// GetMeetingRoomOk returns a tuple with the MeetingRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetMeetingRoomOk() (*bool, bool) {
	if o == nil || IsNil(o.MeetingRoom) {
		return nil, false
	}
	return o.MeetingRoom, true
}

// HasMeetingRoom returns a boolean if a field has been set.
func (o *ConfigRoomType) HasMeetingRoom() bool {
	if o != nil && !IsNil(o.MeetingRoom) {
		return true
	}

	return false
}

// SetMeetingRoom gets a reference to the given bool and assigns it to the MeetingRoom field.
func (o *ConfigRoomType) SetMeetingRoom(v bool) {
	o.MeetingRoom = &v
}

// GetRoomComponents returns the RoomComponents field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRoomComponents() []RoomComponentType {
	if o == nil || IsNil(o.RoomComponents) {
		var ret []RoomComponentType
		return ret
	}
	return o.RoomComponents
}

// GetRoomComponentsOk returns a tuple with the RoomComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRoomComponentsOk() ([]RoomComponentType, bool) {
	if o == nil || IsNil(o.RoomComponents) {
		return nil, false
	}
	return o.RoomComponents, true
}

// HasRoomComponents returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRoomComponents() bool {
	if o != nil && !IsNil(o.RoomComponents) {
		return true
	}

	return false
}

// SetRoomComponents gets a reference to the given []RoomComponentType and assigns it to the RoomComponents field.
func (o *ConfigRoomType) SetRoomComponents(v []RoomComponentType) {
	o.RoomComponents = v
}

// GetConnectingRooms returns the ConnectingRooms field value if set, zero value otherwise.
func (o *ConfigRoomType) GetConnectingRooms() []RoomRoomType {
	if o == nil || IsNil(o.ConnectingRooms) {
		var ret []RoomRoomType
		return ret
	}
	return o.ConnectingRooms
}

// GetConnectingRoomsOk returns a tuple with the ConnectingRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetConnectingRoomsOk() ([]RoomRoomType, bool) {
	if o == nil || IsNil(o.ConnectingRooms) {
		return nil, false
	}
	return o.ConnectingRooms, true
}

// HasConnectingRooms returns a boolean if a field has been set.
func (o *ConfigRoomType) HasConnectingRooms() bool {
	if o != nil && !IsNil(o.ConnectingRooms) {
		return true
	}

	return false
}

// SetConnectingRooms gets a reference to the given []RoomRoomType and assigns it to the ConnectingRooms field.
func (o *ConfigRoomType) SetConnectingRooms(v []RoomRoomType) {
	o.ConnectingRooms = v
}

// GetRateCode returns the RateCode field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRateCode() string {
	if o == nil || IsNil(o.RateCode) {
		var ret string
		return ret
	}
	return *o.RateCode
}

// GetRateCodeOk returns a tuple with the RateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RateCode) {
		return nil, false
	}
	return o.RateCode, true
}

// HasRateCode returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRateCode() bool {
	if o != nil && !IsNil(o.RateCode) {
		return true
	}

	return false
}

// SetRateCode gets a reference to the given string and assigns it to the RateCode field.
func (o *ConfigRoomType) SetRateCode(v string) {
	o.RateCode = &v
}

// GetRateAmount returns the RateAmount field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRateAmount() CurrencyAmountType {
	if o == nil || IsNil(o.RateAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.RateAmount
}

// GetRateAmountOk returns a tuple with the RateAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRateAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.RateAmount) {
		return nil, false
	}
	return o.RateAmount, true
}

// HasRateAmount returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRateAmount() bool {
	if o != nil && !IsNil(o.RateAmount) {
		return true
	}

	return false
}

// SetRateAmount gets a reference to the given CurrencyAmountType and assigns it to the RateAmount field.
func (o *ConfigRoomType) SetRateAmount(v CurrencyAmountType) {
	o.RateAmount = &v
}

// GetMaximumOccupancy returns the MaximumOccupancy field value if set, zero value otherwise.
func (o *ConfigRoomType) GetMaximumOccupancy() int32 {
	if o == nil || IsNil(o.MaximumOccupancy) {
		var ret int32
		return ret
	}
	return *o.MaximumOccupancy
}

// GetMaximumOccupancyOk returns a tuple with the MaximumOccupancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetMaximumOccupancyOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumOccupancy) {
		return nil, false
	}
	return o.MaximumOccupancy, true
}

// HasMaximumOccupancy returns a boolean if a field has been set.
func (o *ConfigRoomType) HasMaximumOccupancy() bool {
	if o != nil && !IsNil(o.MaximumOccupancy) {
		return true
	}

	return false
}

// SetMaximumOccupancy gets a reference to the given int32 and assigns it to the MaximumOccupancy field.
func (o *ConfigRoomType) SetMaximumOccupancy(v int32) {
	o.MaximumOccupancy = &v
}

// GetSellSequence returns the SellSequence field value if set, zero value otherwise.
func (o *ConfigRoomType) GetSellSequence() float32 {
	if o == nil || IsNil(o.SellSequence) {
		var ret float32
		return ret
	}
	return *o.SellSequence
}

// GetSellSequenceOk returns a tuple with the SellSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetSellSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.SellSequence) {
		return nil, false
	}
	return o.SellSequence, true
}

// HasSellSequence returns a boolean if a field has been set.
func (o *ConfigRoomType) HasSellSequence() bool {
	if o != nil && !IsNil(o.SellSequence) {
		return true
	}

	return false
}

// SetSellSequence gets a reference to the given float32 and assigns it to the SellSequence field.
func (o *ConfigRoomType) SetSellSequence(v float32) {
	o.SellSequence = &v
}

// GetOwnerRoom returns the OwnerRoom field value if set, zero value otherwise.
func (o *ConfigRoomType) GetOwnerRoom() bool {
	if o == nil || IsNil(o.OwnerRoom) {
		var ret bool
		return ret
	}
	return *o.OwnerRoom
}

// GetOwnerRoomOk returns a tuple with the OwnerRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetOwnerRoomOk() (*bool, bool) {
	if o == nil || IsNil(o.OwnerRoom) {
		return nil, false
	}
	return o.OwnerRoom, true
}

// HasOwnerRoom returns a boolean if a field has been set.
func (o *ConfigRoomType) HasOwnerRoom() bool {
	if o != nil && !IsNil(o.OwnerRoom) {
		return true
	}

	return false
}

// SetOwnerRoom gets a reference to the given bool and assigns it to the OwnerRoom field.
func (o *ConfigRoomType) SetOwnerRoom(v bool) {
	o.OwnerRoom = &v
}

// GetUnitGradeCode returns the UnitGradeCode field value if set, zero value otherwise.
func (o *ConfigRoomType) GetUnitGradeCode() string {
	if o == nil || IsNil(o.UnitGradeCode) {
		var ret string
		return ret
	}
	return *o.UnitGradeCode
}

// GetUnitGradeCodeOk returns a tuple with the UnitGradeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetUnitGradeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UnitGradeCode) {
		return nil, false
	}
	return o.UnitGradeCode, true
}

// HasUnitGradeCode returns a boolean if a field has been set.
func (o *ConfigRoomType) HasUnitGradeCode() bool {
	if o != nil && !IsNil(o.UnitGradeCode) {
		return true
	}

	return false
}

// SetUnitGradeCode gets a reference to the given string and assigns it to the UnitGradeCode field.
func (o *ConfigRoomType) SetUnitGradeCode(v string) {
	o.UnitGradeCode = &v
}

// GetSmoking returns the Smoking field value if set, zero value otherwise.
func (o *ConfigRoomType) GetSmoking() bool {
	if o == nil || IsNil(o.Smoking) {
		var ret bool
		return ret
	}
	return *o.Smoking
}

// GetSmokingOk returns a tuple with the Smoking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetSmokingOk() (*bool, bool) {
	if o == nil || IsNil(o.Smoking) {
		return nil, false
	}
	return o.Smoking, true
}

// HasSmoking returns a boolean if a field has been set.
func (o *ConfigRoomType) HasSmoking() bool {
	if o != nil && !IsNil(o.Smoking) {
		return true
	}

	return false
}

// SetSmoking gets a reference to the given bool and assigns it to the Smoking field.
func (o *ConfigRoomType) SetSmoking(v bool) {
	o.Smoking = &v
}

// GetKeyCode returns the KeyCode field value if set, zero value otherwise.
func (o *ConfigRoomType) GetKeyCode() string {
	if o == nil || IsNil(o.KeyCode) {
		var ret string
		return ret
	}
	return *o.KeyCode
}

// GetKeyCodeOk returns a tuple with the KeyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetKeyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyCode) {
		return nil, false
	}
	return o.KeyCode, true
}

// HasKeyCode returns a boolean if a field has been set.
func (o *ConfigRoomType) HasKeyCode() bool {
	if o != nil && !IsNil(o.KeyCode) {
		return true
	}

	return false
}

// SetKeyCode gets a reference to the given string and assigns it to the KeyCode field.
func (o *ConfigRoomType) SetKeyCode(v string) {
	o.KeyCode = &v
}

// GetKeyOptions returns the KeyOptions field value if set, zero value otherwise.
func (o *ConfigRoomType) GetKeyOptions() []string {
	if o == nil || IsNil(o.KeyOptions) {
		var ret []string
		return ret
	}
	return o.KeyOptions
}

// GetKeyOptionsOk returns a tuple with the KeyOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetKeyOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.KeyOptions) {
		return nil, false
	}
	return o.KeyOptions, true
}

// HasKeyOptions returns a boolean if a field has been set.
func (o *ConfigRoomType) HasKeyOptions() bool {
	if o != nil && !IsNil(o.KeyOptions) {
		return true
	}

	return false
}

// SetKeyOptions gets a reference to the given []string and assigns it to the KeyOptions field.
func (o *ConfigRoomType) SetKeyOptions(v []string) {
	o.KeyOptions = v
}

// GetSquareUnits returns the SquareUnits field value if set, zero value otherwise.
func (o *ConfigRoomType) GetSquareUnits() float32 {
	if o == nil || IsNil(o.SquareUnits) {
		var ret float32
		return ret
	}
	return *o.SquareUnits
}

// GetSquareUnitsOk returns a tuple with the SquareUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetSquareUnitsOk() (*float32, bool) {
	if o == nil || IsNil(o.SquareUnits) {
		return nil, false
	}
	return o.SquareUnits, true
}

// HasSquareUnits returns a boolean if a field has been set.
func (o *ConfigRoomType) HasSquareUnits() bool {
	if o != nil && !IsNil(o.SquareUnits) {
		return true
	}

	return false
}

// SetSquareUnits gets a reference to the given float32 and assigns it to the SquareUnits field.
func (o *ConfigRoomType) SetSquareUnits(v float32) {
	o.SquareUnits = &v
}

// GetTurndownService returns the TurndownService field value if set, zero value otherwise.
func (o *ConfigRoomType) GetTurndownService() bool {
	if o == nil || IsNil(o.TurndownService) {
		var ret bool
		return ret
	}
	return *o.TurndownService
}

// GetTurndownServiceOk returns a tuple with the TurndownService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetTurndownServiceOk() (*bool, bool) {
	if o == nil || IsNil(o.TurndownService) {
		return nil, false
	}
	return o.TurndownService, true
}

// HasTurndownService returns a boolean if a field has been set.
func (o *ConfigRoomType) HasTurndownService() bool {
	if o != nil && !IsNil(o.TurndownService) {
		return true
	}

	return false
}

// SetTurndownService gets a reference to the given bool and assigns it to the TurndownService field.
func (o *ConfigRoomType) SetTurndownService(v bool) {
	o.TurndownService = &v
}

// GetSquareUnitsMeasurement returns the SquareUnitsMeasurement field value if set, zero value otherwise.
func (o *ConfigRoomType) GetSquareUnitsMeasurement() string {
	if o == nil || IsNil(o.SquareUnitsMeasurement) {
		var ret string
		return ret
	}
	return *o.SquareUnitsMeasurement
}

// GetSquareUnitsMeasurementOk returns a tuple with the SquareUnitsMeasurement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetSquareUnitsMeasurementOk() (*string, bool) {
	if o == nil || IsNil(o.SquareUnitsMeasurement) {
		return nil, false
	}
	return o.SquareUnitsMeasurement, true
}

// HasSquareUnitsMeasurement returns a boolean if a field has been set.
func (o *ConfigRoomType) HasSquareUnitsMeasurement() bool {
	if o != nil && !IsNil(o.SquareUnitsMeasurement) {
		return true
	}

	return false
}

// SetSquareUnitsMeasurement gets a reference to the given string and assigns it to the SquareUnitsMeasurement field.
func (o *ConfigRoomType) SetSquareUnitsMeasurement(v string) {
	o.SquareUnitsMeasurement = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ConfigRoomType) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ConfigRoomType) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ConfigRoomType) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetHousekeepingCredit returns the HousekeepingCredit field value if set, zero value otherwise.
func (o *ConfigRoomType) GetHousekeepingCredit() []HousekeepingCreditsType {
	if o == nil || IsNil(o.HousekeepingCredit) {
		var ret []HousekeepingCreditsType
		return ret
	}
	return o.HousekeepingCredit
}

// GetHousekeepingCreditOk returns a tuple with the HousekeepingCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetHousekeepingCreditOk() ([]HousekeepingCreditsType, bool) {
	if o == nil || IsNil(o.HousekeepingCredit) {
		return nil, false
	}
	return o.HousekeepingCredit, true
}

// HasHousekeepingCredit returns a boolean if a field has been set.
func (o *ConfigRoomType) HasHousekeepingCredit() bool {
	if o != nil && !IsNil(o.HousekeepingCredit) {
		return true
	}

	return false
}

// SetHousekeepingCredit gets a reference to the given []HousekeepingCreditsType and assigns it to the HousekeepingCredit field.
func (o *ConfigRoomType) SetHousekeepingCredit(v []HousekeepingCreditsType) {
	o.HousekeepingCredit = v
}

// GetRoomSection returns the RoomSection field value if set, zero value otherwise.
func (o *ConfigRoomType) GetRoomSection() RoomSectionType {
	if o == nil || IsNil(o.RoomSection) {
		var ret RoomSectionType
		return ret
	}
	return *o.RoomSection
}

// GetRoomSectionOk returns a tuple with the RoomSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetRoomSectionOk() (*RoomSectionType, bool) {
	if o == nil || IsNil(o.RoomSection) {
		return nil, false
	}
	return o.RoomSection, true
}

// HasRoomSection returns a boolean if a field has been set.
func (o *ConfigRoomType) HasRoomSection() bool {
	if o != nil && !IsNil(o.RoomSection) {
		return true
	}

	return false
}

// SetRoomSection gets a reference to the given RoomSectionType and assigns it to the RoomSection field.
func (o *ConfigRoomType) SetRoomSection(v RoomSectionType) {
	o.RoomSection = &v
}

// GetNoOfBeds returns the NoOfBeds field value if set, zero value otherwise.
func (o *ConfigRoomType) GetNoOfBeds() int32 {
	if o == nil || IsNil(o.NoOfBeds) {
		var ret int32
		return ret
	}
	return *o.NoOfBeds
}

// GetNoOfBedsOk returns a tuple with the NoOfBeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigRoomType) GetNoOfBedsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfBeds) {
		return nil, false
	}
	return o.NoOfBeds, true
}

// HasNoOfBeds returns a boolean if a field has been set.
func (o *ConfigRoomType) HasNoOfBeds() bool {
	if o != nil && !IsNil(o.NoOfBeds) {
		return true
	}

	return false
}

// SetNoOfBeds gets a reference to the given int32 and assigns it to the NoOfBeds field.
func (o *ConfigRoomType) SetNoOfBeds(v int32) {
	o.NoOfBeds = &v
}

func (o ConfigRoomType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigRoomType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.Floor) {
		toSerialize["floor"] = o.Floor
	}
	if !IsNil(o.FloorDescription) {
		toSerialize["floorDescription"] = o.FloorDescription
	}
	if !IsNil(o.RoomFeatures) {
		toSerialize["roomFeatures"] = o.RoomFeatures
	}
	if !IsNil(o.RoomDescription) {
		toSerialize["roomDescription"] = o.RoomDescription
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SmokingPreference) {
		toSerialize["smokingPreference"] = o.SmokingPreference
	}
	if !IsNil(o.SmokingPreferenceDescription) {
		toSerialize["smokingPreferenceDescription"] = o.SmokingPreferenceDescription
	}
	if !IsNil(o.Building) {
		toSerialize["building"] = o.Building
	}
	if !IsNil(o.RoomAssignmentRating) {
		toSerialize["roomAssignmentRating"] = o.RoomAssignmentRating
	}
	if !IsNil(o.Accessible) {
		toSerialize["accessible"] = o.Accessible
	}
	if !IsNil(o.RoomId) {
		toSerialize["roomId"] = o.RoomId
	}
	if !IsNil(o.MeetingRoom) {
		toSerialize["meetingRoom"] = o.MeetingRoom
	}
	if !IsNil(o.RoomComponents) {
		toSerialize["roomComponents"] = o.RoomComponents
	}
	if !IsNil(o.ConnectingRooms) {
		toSerialize["connectingRooms"] = o.ConnectingRooms
	}
	if !IsNil(o.RateCode) {
		toSerialize["rateCode"] = o.RateCode
	}
	if !IsNil(o.RateAmount) {
		toSerialize["rateAmount"] = o.RateAmount
	}
	if !IsNil(o.MaximumOccupancy) {
		toSerialize["maximumOccupancy"] = o.MaximumOccupancy
	}
	if !IsNil(o.SellSequence) {
		toSerialize["sellSequence"] = o.SellSequence
	}
	if !IsNil(o.OwnerRoom) {
		toSerialize["ownerRoom"] = o.OwnerRoom
	}
	if !IsNil(o.UnitGradeCode) {
		toSerialize["unitGradeCode"] = o.UnitGradeCode
	}
	if !IsNil(o.Smoking) {
		toSerialize["smoking"] = o.Smoking
	}
	if !IsNil(o.KeyCode) {
		toSerialize["keyCode"] = o.KeyCode
	}
	if !IsNil(o.KeyOptions) {
		toSerialize["keyOptions"] = o.KeyOptions
	}
	if !IsNil(o.SquareUnits) {
		toSerialize["squareUnits"] = o.SquareUnits
	}
	if !IsNil(o.TurndownService) {
		toSerialize["turndownService"] = o.TurndownService
	}
	if !IsNil(o.SquareUnitsMeasurement) {
		toSerialize["squareUnitsMeasurement"] = o.SquareUnitsMeasurement
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.HousekeepingCredit) {
		toSerialize["housekeepingCredit"] = o.HousekeepingCredit
	}
	if !IsNil(o.RoomSection) {
		toSerialize["roomSection"] = o.RoomSection
	}
	if !IsNil(o.NoOfBeds) {
		toSerialize["noOfBeds"] = o.NoOfBeds
	}
	return toSerialize, nil
}

type NullableConfigRoomType struct {
	value *ConfigRoomType
	isSet bool
}

func (v NullableConfigRoomType) Get() *ConfigRoomType {
	return v.value
}

func (v *NullableConfigRoomType) Set(val *ConfigRoomType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigRoomType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigRoomType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigRoomType(val *ConfigRoomType) *NullableConfigRoomType {
	return &NullableConfigRoomType{value: val, isSet: true}
}

func (v NullableConfigRoomType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigRoomType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


