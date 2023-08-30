/*
OPERA Cloud Room Configuration API

APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rmcfg

import (
	"encoding/json"
)

// checks if the RoomTypeSummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomTypeSummaryType{}

// RoomTypeSummaryType This type represents the summary of room type attributes.
type RoomTypeSummaryType struct {
	// Room class for the room type code.
	RoomClass *string `json:"roomClass,omitempty"`
	// Short Description of room type.
	ShortDescription *string `json:"shortDescription,omitempty"`
	// Active date of the room type.
	ActiveDate *string `json:"activeDate,omitempty"`
	// Indicates if room type is pseudo.
	Pseudo *bool `json:"pseudo,omitempty"`
	// Indicates if room type is accessible.
	Accessible *bool `json:"accessible,omitempty"`
	// Indicates if room type is sent to interface.
	SendToInterface *bool `json:"sendToInterface,omitempty"`
	// Indicates room types sell sequence.
	SellSequence *float32 `json:"sellSequence,omitempty"`
	// Indicates room type is a suite.
	Suite *bool `json:"suite,omitempty"`
	// Indicates room type is meeting room. This Can be Meeting room flag cannot be unmarked at the property level. Can only be marked for non pseudo room types.
	MeetingRoom *bool `json:"meetingRoom,omitempty"`
	RoomType *string `json:"roomType,omitempty"`
	// Number of rooms for this room type.
	NumberOfRooms *int32 `json:"numberOfRooms,omitempty"`
	// Indicates the room type is inactive or not.
	Inactive *bool `json:"inactive,omitempty"`
}

// NewRoomTypeSummaryType instantiates a new RoomTypeSummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomTypeSummaryType() *RoomTypeSummaryType {
	this := RoomTypeSummaryType{}
	return &this
}

// NewRoomTypeSummaryTypeWithDefaults instantiates a new RoomTypeSummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomTypeSummaryTypeWithDefaults() *RoomTypeSummaryType {
	this := RoomTypeSummaryType{}
	return &this
}

// GetRoomClass returns the RoomClass field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetRoomClass() string {
	if o == nil || IsNil(o.RoomClass) {
		var ret string
		return ret
	}
	return *o.RoomClass
}

// GetRoomClassOk returns a tuple with the RoomClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetRoomClassOk() (*string, bool) {
	if o == nil || IsNil(o.RoomClass) {
		return nil, false
	}
	return o.RoomClass, true
}

// HasRoomClass returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasRoomClass() bool {
	if o != nil && !IsNil(o.RoomClass) {
		return true
	}

	return false
}

// SetRoomClass gets a reference to the given string and assigns it to the RoomClass field.
func (o *RoomTypeSummaryType) SetRoomClass(v string) {
	o.RoomClass = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *RoomTypeSummaryType) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetActiveDate returns the ActiveDate field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetActiveDate() string {
	if o == nil || IsNil(o.ActiveDate) {
		var ret string
		return ret
	}
	return *o.ActiveDate
}

// GetActiveDateOk returns a tuple with the ActiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetActiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveDate) {
		return nil, false
	}
	return o.ActiveDate, true
}

// HasActiveDate returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasActiveDate() bool {
	if o != nil && !IsNil(o.ActiveDate) {
		return true
	}

	return false
}

// SetActiveDate gets a reference to the given string and assigns it to the ActiveDate field.
func (o *RoomTypeSummaryType) SetActiveDate(v string) {
	o.ActiveDate = &v
}

// GetPseudo returns the Pseudo field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetPseudo() bool {
	if o == nil || IsNil(o.Pseudo) {
		var ret bool
		return ret
	}
	return *o.Pseudo
}

// GetPseudoOk returns a tuple with the Pseudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetPseudoOk() (*bool, bool) {
	if o == nil || IsNil(o.Pseudo) {
		return nil, false
	}
	return o.Pseudo, true
}

// HasPseudo returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasPseudo() bool {
	if o != nil && !IsNil(o.Pseudo) {
		return true
	}

	return false
}

// SetPseudo gets a reference to the given bool and assigns it to the Pseudo field.
func (o *RoomTypeSummaryType) SetPseudo(v bool) {
	o.Pseudo = &v
}

// GetAccessible returns the Accessible field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetAccessible() bool {
	if o == nil || IsNil(o.Accessible) {
		var ret bool
		return ret
	}
	return *o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetAccessibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Accessible) {
		return nil, false
	}
	return o.Accessible, true
}

// HasAccessible returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasAccessible() bool {
	if o != nil && !IsNil(o.Accessible) {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given bool and assigns it to the Accessible field.
func (o *RoomTypeSummaryType) SetAccessible(v bool) {
	o.Accessible = &v
}

// GetSendToInterface returns the SendToInterface field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetSendToInterface() bool {
	if o == nil || IsNil(o.SendToInterface) {
		var ret bool
		return ret
	}
	return *o.SendToInterface
}

// GetSendToInterfaceOk returns a tuple with the SendToInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetSendToInterfaceOk() (*bool, bool) {
	if o == nil || IsNil(o.SendToInterface) {
		return nil, false
	}
	return o.SendToInterface, true
}

// HasSendToInterface returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasSendToInterface() bool {
	if o != nil && !IsNil(o.SendToInterface) {
		return true
	}

	return false
}

// SetSendToInterface gets a reference to the given bool and assigns it to the SendToInterface field.
func (o *RoomTypeSummaryType) SetSendToInterface(v bool) {
	o.SendToInterface = &v
}

// GetSellSequence returns the SellSequence field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetSellSequence() float32 {
	if o == nil || IsNil(o.SellSequence) {
		var ret float32
		return ret
	}
	return *o.SellSequence
}

// GetSellSequenceOk returns a tuple with the SellSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetSellSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.SellSequence) {
		return nil, false
	}
	return o.SellSequence, true
}

// HasSellSequence returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasSellSequence() bool {
	if o != nil && !IsNil(o.SellSequence) {
		return true
	}

	return false
}

// SetSellSequence gets a reference to the given float32 and assigns it to the SellSequence field.
func (o *RoomTypeSummaryType) SetSellSequence(v float32) {
	o.SellSequence = &v
}

// GetSuite returns the Suite field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetSuite() bool {
	if o == nil || IsNil(o.Suite) {
		var ret bool
		return ret
	}
	return *o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetSuiteOk() (*bool, bool) {
	if o == nil || IsNil(o.Suite) {
		return nil, false
	}
	return o.Suite, true
}

// HasSuite returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasSuite() bool {
	if o != nil && !IsNil(o.Suite) {
		return true
	}

	return false
}

// SetSuite gets a reference to the given bool and assigns it to the Suite field.
func (o *RoomTypeSummaryType) SetSuite(v bool) {
	o.Suite = &v
}

// GetMeetingRoom returns the MeetingRoom field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetMeetingRoom() bool {
	if o == nil || IsNil(o.MeetingRoom) {
		var ret bool
		return ret
	}
	return *o.MeetingRoom
}

// GetMeetingRoomOk returns a tuple with the MeetingRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetMeetingRoomOk() (*bool, bool) {
	if o == nil || IsNil(o.MeetingRoom) {
		return nil, false
	}
	return o.MeetingRoom, true
}

// HasMeetingRoom returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasMeetingRoom() bool {
	if o != nil && !IsNil(o.MeetingRoom) {
		return true
	}

	return false
}

// SetMeetingRoom gets a reference to the given bool and assigns it to the MeetingRoom field.
func (o *RoomTypeSummaryType) SetMeetingRoom(v bool) {
	o.MeetingRoom = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *RoomTypeSummaryType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetNumberOfRooms returns the NumberOfRooms field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetNumberOfRooms() int32 {
	if o == nil || IsNil(o.NumberOfRooms) {
		var ret int32
		return ret
	}
	return *o.NumberOfRooms
}

// GetNumberOfRoomsOk returns a tuple with the NumberOfRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetNumberOfRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfRooms) {
		return nil, false
	}
	return o.NumberOfRooms, true
}

// HasNumberOfRooms returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasNumberOfRooms() bool {
	if o != nil && !IsNil(o.NumberOfRooms) {
		return true
	}

	return false
}

// SetNumberOfRooms gets a reference to the given int32 and assigns it to the NumberOfRooms field.
func (o *RoomTypeSummaryType) SetNumberOfRooms(v int32) {
	o.NumberOfRooms = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *RoomTypeSummaryType) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeSummaryType) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *RoomTypeSummaryType) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *RoomTypeSummaryType) SetInactive(v bool) {
	o.Inactive = &v
}

func (o RoomTypeSummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomTypeSummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomClass) {
		toSerialize["roomClass"] = o.RoomClass
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.ActiveDate) {
		toSerialize["activeDate"] = o.ActiveDate
	}
	if !IsNil(o.Pseudo) {
		toSerialize["pseudo"] = o.Pseudo
	}
	if !IsNil(o.Accessible) {
		toSerialize["accessible"] = o.Accessible
	}
	if !IsNil(o.SendToInterface) {
		toSerialize["sendToInterface"] = o.SendToInterface
	}
	if !IsNil(o.SellSequence) {
		toSerialize["sellSequence"] = o.SellSequence
	}
	if !IsNil(o.Suite) {
		toSerialize["suite"] = o.Suite
	}
	if !IsNil(o.MeetingRoom) {
		toSerialize["meetingRoom"] = o.MeetingRoom
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.NumberOfRooms) {
		toSerialize["numberOfRooms"] = o.NumberOfRooms
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	return toSerialize, nil
}

type NullableRoomTypeSummaryType struct {
	value *RoomTypeSummaryType
	isSet bool
}

func (v NullableRoomTypeSummaryType) Get() *RoomTypeSummaryType {
	return v.value
}

func (v *NullableRoomTypeSummaryType) Set(val *RoomTypeSummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomTypeSummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomTypeSummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomTypeSummaryType(val *RoomTypeSummaryType) *NullableRoomTypeSummaryType {
	return &NullableRoomTypeSummaryType{value: val, isSet: true}
}

func (v NullableRoomTypeSummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomTypeSummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


