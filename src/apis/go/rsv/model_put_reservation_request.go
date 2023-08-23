/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PutReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutReservationRequest{}

// PutReservationRequest struct for PutReservationRequest
type PutReservationRequest struct {
	// A collection of Reservations with information that needs to be changed.
	Reservations []HotelReservationInstructionType `json:"reservations,omitempty"`
	ReservationsInstructionsType *ReservationsInstructionsType `json:"reservationsInstructionsType,omitempty"`
	ChannelInformation *ChannelResvRQInfoType `json:"channelInformation,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutReservationRequest instantiates a new PutReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutReservationRequest() *PutReservationRequest {
	this := PutReservationRequest{}
	return &this
}

// NewPutReservationRequestWithDefaults instantiates a new PutReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutReservationRequestWithDefaults() *PutReservationRequest {
	this := PutReservationRequest{}
	return &this
}

// GetReservations returns the Reservations field value if set, zero value otherwise.
func (o *PutReservationRequest) GetReservations() []HotelReservationInstructionType {
	if o == nil || IsNil(o.Reservations) {
		var ret []HotelReservationInstructionType
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutReservationRequest) GetReservationsOk() ([]HotelReservationInstructionType, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *PutReservationRequest) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []HotelReservationInstructionType and assigns it to the Reservations field.
func (o *PutReservationRequest) SetReservations(v []HotelReservationInstructionType) {
	o.Reservations = v
}

// GetReservationsInstructionsType returns the ReservationsInstructionsType field value if set, zero value otherwise.
func (o *PutReservationRequest) GetReservationsInstructionsType() ReservationsInstructionsType {
	if o == nil || IsNil(o.ReservationsInstructionsType) {
		var ret ReservationsInstructionsType
		return ret
	}
	return *o.ReservationsInstructionsType
}

// GetReservationsInstructionsTypeOk returns a tuple with the ReservationsInstructionsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutReservationRequest) GetReservationsInstructionsTypeOk() (*ReservationsInstructionsType, bool) {
	if o == nil || IsNil(o.ReservationsInstructionsType) {
		return nil, false
	}
	return o.ReservationsInstructionsType, true
}

// HasReservationsInstructionsType returns a boolean if a field has been set.
func (o *PutReservationRequest) HasReservationsInstructionsType() bool {
	if o != nil && !IsNil(o.ReservationsInstructionsType) {
		return true
	}

	return false
}

// SetReservationsInstructionsType gets a reference to the given ReservationsInstructionsType and assigns it to the ReservationsInstructionsType field.
func (o *PutReservationRequest) SetReservationsInstructionsType(v ReservationsInstructionsType) {
	o.ReservationsInstructionsType = &v
}

// GetChannelInformation returns the ChannelInformation field value if set, zero value otherwise.
func (o *PutReservationRequest) GetChannelInformation() ChannelResvRQInfoType {
	if o == nil || IsNil(o.ChannelInformation) {
		var ret ChannelResvRQInfoType
		return ret
	}
	return *o.ChannelInformation
}

// GetChannelInformationOk returns a tuple with the ChannelInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutReservationRequest) GetChannelInformationOk() (*ChannelResvRQInfoType, bool) {
	if o == nil || IsNil(o.ChannelInformation) {
		return nil, false
	}
	return o.ChannelInformation, true
}

// HasChannelInformation returns a boolean if a field has been set.
func (o *PutReservationRequest) HasChannelInformation() bool {
	if o != nil && !IsNil(o.ChannelInformation) {
		return true
	}

	return false
}

// SetChannelInformation gets a reference to the given ChannelResvRQInfoType and assigns it to the ChannelInformation field.
func (o *PutReservationRequest) SetChannelInformation(v ChannelResvRQInfoType) {
	o.ChannelInformation = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutReservationRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutReservationRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutReservationRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutReservationRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutReservationRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutReservationRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutReservationRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutReservationRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reservations) {
		toSerialize["reservations"] = o.Reservations
	}
	if !IsNil(o.ReservationsInstructionsType) {
		toSerialize["reservationsInstructionsType"] = o.ReservationsInstructionsType
	}
	if !IsNil(o.ChannelInformation) {
		toSerialize["channelInformation"] = o.ChannelInformation
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutReservationRequest struct {
	value *PutReservationRequest
	isSet bool
}

func (v NullablePutReservationRequest) Get() *PutReservationRequest {
	return v.value
}

func (v *NullablePutReservationRequest) Set(val *PutReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutReservationRequest(val *PutReservationRequest) *NullablePutReservationRequest {
	return &NullablePutReservationRequest{value: val, isSet: true}
}

func (v NullablePutReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


