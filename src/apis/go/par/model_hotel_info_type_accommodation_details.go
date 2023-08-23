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

// checks if the HotelInfoTypeAccommodationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelInfoTypeAccommodationDetails{}

// HotelInfoTypeAccommodationDetails Accommodation Details of the hotel.
type HotelInfoTypeAccommodationDetails struct {
	// The number of Single Rooms of the Hotel.
	SingleRooms *int32 `json:"singleRooms,omitempty"`
	// The number of Double Rooms of the Hotel.
	DoubleRooms *int32 `json:"doubleRooms,omitempty"`
	// The number of Twin Rooms of the Hotel.
	TwinRooms *int32 `json:"twinRooms,omitempty"`
	// The number of Family Rooms of the Hotel.
	FamilyRooms *int32 `json:"familyRooms,omitempty"`
	// The number of Connecting Rooms of the Hotel.
	ConnectingRooms *int32 `json:"connectingRooms,omitempty"`
	// The number of Accessible Rooms of the Hotel.
	AccessibleRooms *int32 `json:"accessibleRooms,omitempty"`
	// The number of Non-Smoking Rooms of the Hotel.
	NonSmokingRooms *int32 `json:"nonSmokingRooms,omitempty"`
	// Maximum Adults for Family Room Type.
	MaxAdultsInFamilyRoom *int32 `json:"maxAdultsInFamilyRoom,omitempty"`
	// Maximum Children for Family Room Type.
	MaxChildrenInFamilyRoom *int32 `json:"maxChildrenInFamilyRoom,omitempty"`
	// The total number of the Guest Room Floors.
	GuestRoomFloors *int32 `json:"guestRoomFloors,omitempty"`
	// The number of Guest Room Elevators.
	GuestRoomElevators *int32 `json:"guestRoomElevators,omitempty"`
	// The number of Suites of the Hotel.
	Suites *int32 `json:"suites,omitempty"`
	// The floor number of Executive Floors of the Hotel.
	ExecutiveFloorNo *string `json:"executiveFloorNo,omitempty"`
	// The information about the Room Amenities.
	RoomAmenties *string `json:"roomAmenties,omitempty"`
	// The Description of the shops in the Hotel.
	ShopDescription *string `json:"shopDescription,omitempty"`
}

// NewHotelInfoTypeAccommodationDetails instantiates a new HotelInfoTypeAccommodationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelInfoTypeAccommodationDetails() *HotelInfoTypeAccommodationDetails {
	this := HotelInfoTypeAccommodationDetails{}
	return &this
}

// NewHotelInfoTypeAccommodationDetailsWithDefaults instantiates a new HotelInfoTypeAccommodationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelInfoTypeAccommodationDetailsWithDefaults() *HotelInfoTypeAccommodationDetails {
	this := HotelInfoTypeAccommodationDetails{}
	return &this
}

// GetSingleRooms returns the SingleRooms field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetSingleRooms() int32 {
	if o == nil || IsNil(o.SingleRooms) {
		var ret int32
		return ret
	}
	return *o.SingleRooms
}

// GetSingleRoomsOk returns a tuple with the SingleRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetSingleRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.SingleRooms) {
		return nil, false
	}
	return o.SingleRooms, true
}

// HasSingleRooms returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasSingleRooms() bool {
	if o != nil && !IsNil(o.SingleRooms) {
		return true
	}

	return false
}

// SetSingleRooms gets a reference to the given int32 and assigns it to the SingleRooms field.
func (o *HotelInfoTypeAccommodationDetails) SetSingleRooms(v int32) {
	o.SingleRooms = &v
}

// GetDoubleRooms returns the DoubleRooms field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetDoubleRooms() int32 {
	if o == nil || IsNil(o.DoubleRooms) {
		var ret int32
		return ret
	}
	return *o.DoubleRooms
}

// GetDoubleRoomsOk returns a tuple with the DoubleRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetDoubleRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.DoubleRooms) {
		return nil, false
	}
	return o.DoubleRooms, true
}

// HasDoubleRooms returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasDoubleRooms() bool {
	if o != nil && !IsNil(o.DoubleRooms) {
		return true
	}

	return false
}

// SetDoubleRooms gets a reference to the given int32 and assigns it to the DoubleRooms field.
func (o *HotelInfoTypeAccommodationDetails) SetDoubleRooms(v int32) {
	o.DoubleRooms = &v
}

// GetTwinRooms returns the TwinRooms field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetTwinRooms() int32 {
	if o == nil || IsNil(o.TwinRooms) {
		var ret int32
		return ret
	}
	return *o.TwinRooms
}

// GetTwinRoomsOk returns a tuple with the TwinRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetTwinRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.TwinRooms) {
		return nil, false
	}
	return o.TwinRooms, true
}

// HasTwinRooms returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasTwinRooms() bool {
	if o != nil && !IsNil(o.TwinRooms) {
		return true
	}

	return false
}

// SetTwinRooms gets a reference to the given int32 and assigns it to the TwinRooms field.
func (o *HotelInfoTypeAccommodationDetails) SetTwinRooms(v int32) {
	o.TwinRooms = &v
}

// GetFamilyRooms returns the FamilyRooms field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetFamilyRooms() int32 {
	if o == nil || IsNil(o.FamilyRooms) {
		var ret int32
		return ret
	}
	return *o.FamilyRooms
}

// GetFamilyRoomsOk returns a tuple with the FamilyRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetFamilyRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.FamilyRooms) {
		return nil, false
	}
	return o.FamilyRooms, true
}

// HasFamilyRooms returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasFamilyRooms() bool {
	if o != nil && !IsNil(o.FamilyRooms) {
		return true
	}

	return false
}

// SetFamilyRooms gets a reference to the given int32 and assigns it to the FamilyRooms field.
func (o *HotelInfoTypeAccommodationDetails) SetFamilyRooms(v int32) {
	o.FamilyRooms = &v
}

// GetConnectingRooms returns the ConnectingRooms field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetConnectingRooms() int32 {
	if o == nil || IsNil(o.ConnectingRooms) {
		var ret int32
		return ret
	}
	return *o.ConnectingRooms
}

// GetConnectingRoomsOk returns a tuple with the ConnectingRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetConnectingRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectingRooms) {
		return nil, false
	}
	return o.ConnectingRooms, true
}

// HasConnectingRooms returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasConnectingRooms() bool {
	if o != nil && !IsNil(o.ConnectingRooms) {
		return true
	}

	return false
}

// SetConnectingRooms gets a reference to the given int32 and assigns it to the ConnectingRooms field.
func (o *HotelInfoTypeAccommodationDetails) SetConnectingRooms(v int32) {
	o.ConnectingRooms = &v
}

// GetAccessibleRooms returns the AccessibleRooms field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetAccessibleRooms() int32 {
	if o == nil || IsNil(o.AccessibleRooms) {
		var ret int32
		return ret
	}
	return *o.AccessibleRooms
}

// GetAccessibleRoomsOk returns a tuple with the AccessibleRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetAccessibleRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessibleRooms) {
		return nil, false
	}
	return o.AccessibleRooms, true
}

// HasAccessibleRooms returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasAccessibleRooms() bool {
	if o != nil && !IsNil(o.AccessibleRooms) {
		return true
	}

	return false
}

// SetAccessibleRooms gets a reference to the given int32 and assigns it to the AccessibleRooms field.
func (o *HotelInfoTypeAccommodationDetails) SetAccessibleRooms(v int32) {
	o.AccessibleRooms = &v
}

// GetNonSmokingRooms returns the NonSmokingRooms field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetNonSmokingRooms() int32 {
	if o == nil || IsNil(o.NonSmokingRooms) {
		var ret int32
		return ret
	}
	return *o.NonSmokingRooms
}

// GetNonSmokingRoomsOk returns a tuple with the NonSmokingRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetNonSmokingRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.NonSmokingRooms) {
		return nil, false
	}
	return o.NonSmokingRooms, true
}

// HasNonSmokingRooms returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasNonSmokingRooms() bool {
	if o != nil && !IsNil(o.NonSmokingRooms) {
		return true
	}

	return false
}

// SetNonSmokingRooms gets a reference to the given int32 and assigns it to the NonSmokingRooms field.
func (o *HotelInfoTypeAccommodationDetails) SetNonSmokingRooms(v int32) {
	o.NonSmokingRooms = &v
}

// GetMaxAdultsInFamilyRoom returns the MaxAdultsInFamilyRoom field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetMaxAdultsInFamilyRoom() int32 {
	if o == nil || IsNil(o.MaxAdultsInFamilyRoom) {
		var ret int32
		return ret
	}
	return *o.MaxAdultsInFamilyRoom
}

// GetMaxAdultsInFamilyRoomOk returns a tuple with the MaxAdultsInFamilyRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetMaxAdultsInFamilyRoomOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAdultsInFamilyRoom) {
		return nil, false
	}
	return o.MaxAdultsInFamilyRoom, true
}

// HasMaxAdultsInFamilyRoom returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasMaxAdultsInFamilyRoom() bool {
	if o != nil && !IsNil(o.MaxAdultsInFamilyRoom) {
		return true
	}

	return false
}

// SetMaxAdultsInFamilyRoom gets a reference to the given int32 and assigns it to the MaxAdultsInFamilyRoom field.
func (o *HotelInfoTypeAccommodationDetails) SetMaxAdultsInFamilyRoom(v int32) {
	o.MaxAdultsInFamilyRoom = &v
}

// GetMaxChildrenInFamilyRoom returns the MaxChildrenInFamilyRoom field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetMaxChildrenInFamilyRoom() int32 {
	if o == nil || IsNil(o.MaxChildrenInFamilyRoom) {
		var ret int32
		return ret
	}
	return *o.MaxChildrenInFamilyRoom
}

// GetMaxChildrenInFamilyRoomOk returns a tuple with the MaxChildrenInFamilyRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetMaxChildrenInFamilyRoomOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxChildrenInFamilyRoom) {
		return nil, false
	}
	return o.MaxChildrenInFamilyRoom, true
}

// HasMaxChildrenInFamilyRoom returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasMaxChildrenInFamilyRoom() bool {
	if o != nil && !IsNil(o.MaxChildrenInFamilyRoom) {
		return true
	}

	return false
}

// SetMaxChildrenInFamilyRoom gets a reference to the given int32 and assigns it to the MaxChildrenInFamilyRoom field.
func (o *HotelInfoTypeAccommodationDetails) SetMaxChildrenInFamilyRoom(v int32) {
	o.MaxChildrenInFamilyRoom = &v
}

// GetGuestRoomFloors returns the GuestRoomFloors field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetGuestRoomFloors() int32 {
	if o == nil || IsNil(o.GuestRoomFloors) {
		var ret int32
		return ret
	}
	return *o.GuestRoomFloors
}

// GetGuestRoomFloorsOk returns a tuple with the GuestRoomFloors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetGuestRoomFloorsOk() (*int32, bool) {
	if o == nil || IsNil(o.GuestRoomFloors) {
		return nil, false
	}
	return o.GuestRoomFloors, true
}

// HasGuestRoomFloors returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasGuestRoomFloors() bool {
	if o != nil && !IsNil(o.GuestRoomFloors) {
		return true
	}

	return false
}

// SetGuestRoomFloors gets a reference to the given int32 and assigns it to the GuestRoomFloors field.
func (o *HotelInfoTypeAccommodationDetails) SetGuestRoomFloors(v int32) {
	o.GuestRoomFloors = &v
}

// GetGuestRoomElevators returns the GuestRoomElevators field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetGuestRoomElevators() int32 {
	if o == nil || IsNil(o.GuestRoomElevators) {
		var ret int32
		return ret
	}
	return *o.GuestRoomElevators
}

// GetGuestRoomElevatorsOk returns a tuple with the GuestRoomElevators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetGuestRoomElevatorsOk() (*int32, bool) {
	if o == nil || IsNil(o.GuestRoomElevators) {
		return nil, false
	}
	return o.GuestRoomElevators, true
}

// HasGuestRoomElevators returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasGuestRoomElevators() bool {
	if o != nil && !IsNil(o.GuestRoomElevators) {
		return true
	}

	return false
}

// SetGuestRoomElevators gets a reference to the given int32 and assigns it to the GuestRoomElevators field.
func (o *HotelInfoTypeAccommodationDetails) SetGuestRoomElevators(v int32) {
	o.GuestRoomElevators = &v
}

// GetSuites returns the Suites field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetSuites() int32 {
	if o == nil || IsNil(o.Suites) {
		var ret int32
		return ret
	}
	return *o.Suites
}

// GetSuitesOk returns a tuple with the Suites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetSuitesOk() (*int32, bool) {
	if o == nil || IsNil(o.Suites) {
		return nil, false
	}
	return o.Suites, true
}

// HasSuites returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasSuites() bool {
	if o != nil && !IsNil(o.Suites) {
		return true
	}

	return false
}

// SetSuites gets a reference to the given int32 and assigns it to the Suites field.
func (o *HotelInfoTypeAccommodationDetails) SetSuites(v int32) {
	o.Suites = &v
}

// GetExecutiveFloorNo returns the ExecutiveFloorNo field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetExecutiveFloorNo() string {
	if o == nil || IsNil(o.ExecutiveFloorNo) {
		var ret string
		return ret
	}
	return *o.ExecutiveFloorNo
}

// GetExecutiveFloorNoOk returns a tuple with the ExecutiveFloorNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetExecutiveFloorNoOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutiveFloorNo) {
		return nil, false
	}
	return o.ExecutiveFloorNo, true
}

// HasExecutiveFloorNo returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasExecutiveFloorNo() bool {
	if o != nil && !IsNil(o.ExecutiveFloorNo) {
		return true
	}

	return false
}

// SetExecutiveFloorNo gets a reference to the given string and assigns it to the ExecutiveFloorNo field.
func (o *HotelInfoTypeAccommodationDetails) SetExecutiveFloorNo(v string) {
	o.ExecutiveFloorNo = &v
}

// GetRoomAmenties returns the RoomAmenties field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetRoomAmenties() string {
	if o == nil || IsNil(o.RoomAmenties) {
		var ret string
		return ret
	}
	return *o.RoomAmenties
}

// GetRoomAmentiesOk returns a tuple with the RoomAmenties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetRoomAmentiesOk() (*string, bool) {
	if o == nil || IsNil(o.RoomAmenties) {
		return nil, false
	}
	return o.RoomAmenties, true
}

// HasRoomAmenties returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasRoomAmenties() bool {
	if o != nil && !IsNil(o.RoomAmenties) {
		return true
	}

	return false
}

// SetRoomAmenties gets a reference to the given string and assigns it to the RoomAmenties field.
func (o *HotelInfoTypeAccommodationDetails) SetRoomAmenties(v string) {
	o.RoomAmenties = &v
}

// GetShopDescription returns the ShopDescription field value if set, zero value otherwise.
func (o *HotelInfoTypeAccommodationDetails) GetShopDescription() string {
	if o == nil || IsNil(o.ShopDescription) {
		var ret string
		return ret
	}
	return *o.ShopDescription
}

// GetShopDescriptionOk returns a tuple with the ShopDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeAccommodationDetails) GetShopDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShopDescription) {
		return nil, false
	}
	return o.ShopDescription, true
}

// HasShopDescription returns a boolean if a field has been set.
func (o *HotelInfoTypeAccommodationDetails) HasShopDescription() bool {
	if o != nil && !IsNil(o.ShopDescription) {
		return true
	}

	return false
}

// SetShopDescription gets a reference to the given string and assigns it to the ShopDescription field.
func (o *HotelInfoTypeAccommodationDetails) SetShopDescription(v string) {
	o.ShopDescription = &v
}

func (o HotelInfoTypeAccommodationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelInfoTypeAccommodationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SingleRooms) {
		toSerialize["singleRooms"] = o.SingleRooms
	}
	if !IsNil(o.DoubleRooms) {
		toSerialize["doubleRooms"] = o.DoubleRooms
	}
	if !IsNil(o.TwinRooms) {
		toSerialize["twinRooms"] = o.TwinRooms
	}
	if !IsNil(o.FamilyRooms) {
		toSerialize["familyRooms"] = o.FamilyRooms
	}
	if !IsNil(o.ConnectingRooms) {
		toSerialize["connectingRooms"] = o.ConnectingRooms
	}
	if !IsNil(o.AccessibleRooms) {
		toSerialize["accessibleRooms"] = o.AccessibleRooms
	}
	if !IsNil(o.NonSmokingRooms) {
		toSerialize["nonSmokingRooms"] = o.NonSmokingRooms
	}
	if !IsNil(o.MaxAdultsInFamilyRoom) {
		toSerialize["maxAdultsInFamilyRoom"] = o.MaxAdultsInFamilyRoom
	}
	if !IsNil(o.MaxChildrenInFamilyRoom) {
		toSerialize["maxChildrenInFamilyRoom"] = o.MaxChildrenInFamilyRoom
	}
	if !IsNil(o.GuestRoomFloors) {
		toSerialize["guestRoomFloors"] = o.GuestRoomFloors
	}
	if !IsNil(o.GuestRoomElevators) {
		toSerialize["guestRoomElevators"] = o.GuestRoomElevators
	}
	if !IsNil(o.Suites) {
		toSerialize["suites"] = o.Suites
	}
	if !IsNil(o.ExecutiveFloorNo) {
		toSerialize["executiveFloorNo"] = o.ExecutiveFloorNo
	}
	if !IsNil(o.RoomAmenties) {
		toSerialize["roomAmenties"] = o.RoomAmenties
	}
	if !IsNil(o.ShopDescription) {
		toSerialize["shopDescription"] = o.ShopDescription
	}
	return toSerialize, nil
}

type NullableHotelInfoTypeAccommodationDetails struct {
	value *HotelInfoTypeAccommodationDetails
	isSet bool
}

func (v NullableHotelInfoTypeAccommodationDetails) Get() *HotelInfoTypeAccommodationDetails {
	return v.value
}

func (v *NullableHotelInfoTypeAccommodationDetails) Set(val *HotelInfoTypeAccommodationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelInfoTypeAccommodationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelInfoTypeAccommodationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelInfoTypeAccommodationDetails(val *HotelInfoTypeAccommodationDetails) *NullableHotelInfoTypeAccommodationDetails {
	return &NullableHotelInfoTypeAccommodationDetails{value: val, isSet: true}
}

func (v NullableHotelInfoTypeAccommodationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelInfoTypeAccommodationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


