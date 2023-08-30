/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the ResGuestType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResGuestType{}

// ResGuestType A collection of ResGuest objects, identifying the guests associated with this reservation. Which guests are in which room is determined by each RoomStays ResGuestRPHs collection.
type ResGuestType struct {
	ProfileInfo *ResGuestTypeProfileInfo `json:"profileInfo,omitempty"`
	ArrivalTransport *TransportInfoType `json:"arrivalTransport,omitempty"`
	DepartureTransport *TransportInfoType `json:"departureTransport,omitempty"`
	VisaInfo *VisaInfoType `json:"visaInfo,omitempty"`
	// This is a reference placeholder, used as an index for this guest in this reservation. In the ResGuest object it is used like all other RPH attributes to send the delta of a reservation. It is used by the RoomStay and Service objects to indicate which guests are associated with that room stay or service.
	ReservationGuestRPH *string `json:"reservationGuestRPH,omitempty"`
	// When true indicates this is the primary guest.
	Primary *bool `json:"primary,omitempty"`
}

// NewResGuestType instantiates a new ResGuestType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResGuestType() *ResGuestType {
	this := ResGuestType{}
	return &this
}

// NewResGuestTypeWithDefaults instantiates a new ResGuestType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResGuestTypeWithDefaults() *ResGuestType {
	this := ResGuestType{}
	return &this
}

// GetProfileInfo returns the ProfileInfo field value if set, zero value otherwise.
func (o *ResGuestType) GetProfileInfo() ResGuestTypeProfileInfo {
	if o == nil || IsNil(o.ProfileInfo) {
		var ret ResGuestTypeProfileInfo
		return ret
	}
	return *o.ProfileInfo
}

// GetProfileInfoOk returns a tuple with the ProfileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResGuestType) GetProfileInfoOk() (*ResGuestTypeProfileInfo, bool) {
	if o == nil || IsNil(o.ProfileInfo) {
		return nil, false
	}
	return o.ProfileInfo, true
}

// HasProfileInfo returns a boolean if a field has been set.
func (o *ResGuestType) HasProfileInfo() bool {
	if o != nil && !IsNil(o.ProfileInfo) {
		return true
	}

	return false
}

// SetProfileInfo gets a reference to the given ResGuestTypeProfileInfo and assigns it to the ProfileInfo field.
func (o *ResGuestType) SetProfileInfo(v ResGuestTypeProfileInfo) {
	o.ProfileInfo = &v
}

// GetArrivalTransport returns the ArrivalTransport field value if set, zero value otherwise.
func (o *ResGuestType) GetArrivalTransport() TransportInfoType {
	if o == nil || IsNil(o.ArrivalTransport) {
		var ret TransportInfoType
		return ret
	}
	return *o.ArrivalTransport
}

// GetArrivalTransportOk returns a tuple with the ArrivalTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResGuestType) GetArrivalTransportOk() (*TransportInfoType, bool) {
	if o == nil || IsNil(o.ArrivalTransport) {
		return nil, false
	}
	return o.ArrivalTransport, true
}

// HasArrivalTransport returns a boolean if a field has been set.
func (o *ResGuestType) HasArrivalTransport() bool {
	if o != nil && !IsNil(o.ArrivalTransport) {
		return true
	}

	return false
}

// SetArrivalTransport gets a reference to the given TransportInfoType and assigns it to the ArrivalTransport field.
func (o *ResGuestType) SetArrivalTransport(v TransportInfoType) {
	o.ArrivalTransport = &v
}

// GetDepartureTransport returns the DepartureTransport field value if set, zero value otherwise.
func (o *ResGuestType) GetDepartureTransport() TransportInfoType {
	if o == nil || IsNil(o.DepartureTransport) {
		var ret TransportInfoType
		return ret
	}
	return *o.DepartureTransport
}

// GetDepartureTransportOk returns a tuple with the DepartureTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResGuestType) GetDepartureTransportOk() (*TransportInfoType, bool) {
	if o == nil || IsNil(o.DepartureTransport) {
		return nil, false
	}
	return o.DepartureTransport, true
}

// HasDepartureTransport returns a boolean if a field has been set.
func (o *ResGuestType) HasDepartureTransport() bool {
	if o != nil && !IsNil(o.DepartureTransport) {
		return true
	}

	return false
}

// SetDepartureTransport gets a reference to the given TransportInfoType and assigns it to the DepartureTransport field.
func (o *ResGuestType) SetDepartureTransport(v TransportInfoType) {
	o.DepartureTransport = &v
}

// GetVisaInfo returns the VisaInfo field value if set, zero value otherwise.
func (o *ResGuestType) GetVisaInfo() VisaInfoType {
	if o == nil || IsNil(o.VisaInfo) {
		var ret VisaInfoType
		return ret
	}
	return *o.VisaInfo
}

// GetVisaInfoOk returns a tuple with the VisaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResGuestType) GetVisaInfoOk() (*VisaInfoType, bool) {
	if o == nil || IsNil(o.VisaInfo) {
		return nil, false
	}
	return o.VisaInfo, true
}

// HasVisaInfo returns a boolean if a field has been set.
func (o *ResGuestType) HasVisaInfo() bool {
	if o != nil && !IsNil(o.VisaInfo) {
		return true
	}

	return false
}

// SetVisaInfo gets a reference to the given VisaInfoType and assigns it to the VisaInfo field.
func (o *ResGuestType) SetVisaInfo(v VisaInfoType) {
	o.VisaInfo = &v
}

// GetReservationGuestRPH returns the ReservationGuestRPH field value if set, zero value otherwise.
func (o *ResGuestType) GetReservationGuestRPH() string {
	if o == nil || IsNil(o.ReservationGuestRPH) {
		var ret string
		return ret
	}
	return *o.ReservationGuestRPH
}

// GetReservationGuestRPHOk returns a tuple with the ReservationGuestRPH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResGuestType) GetReservationGuestRPHOk() (*string, bool) {
	if o == nil || IsNil(o.ReservationGuestRPH) {
		return nil, false
	}
	return o.ReservationGuestRPH, true
}

// HasReservationGuestRPH returns a boolean if a field has been set.
func (o *ResGuestType) HasReservationGuestRPH() bool {
	if o != nil && !IsNil(o.ReservationGuestRPH) {
		return true
	}

	return false
}

// SetReservationGuestRPH gets a reference to the given string and assigns it to the ReservationGuestRPH field.
func (o *ResGuestType) SetReservationGuestRPH(v string) {
	o.ReservationGuestRPH = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *ResGuestType) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResGuestType) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *ResGuestType) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *ResGuestType) SetPrimary(v bool) {
	o.Primary = &v
}

func (o ResGuestType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResGuestType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileInfo) {
		toSerialize["profileInfo"] = o.ProfileInfo
	}
	if !IsNil(o.ArrivalTransport) {
		toSerialize["arrivalTransport"] = o.ArrivalTransport
	}
	if !IsNil(o.DepartureTransport) {
		toSerialize["departureTransport"] = o.DepartureTransport
	}
	if !IsNil(o.VisaInfo) {
		toSerialize["visaInfo"] = o.VisaInfo
	}
	if !IsNil(o.ReservationGuestRPH) {
		toSerialize["reservationGuestRPH"] = o.ReservationGuestRPH
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	return toSerialize, nil
}

type NullableResGuestType struct {
	value *ResGuestType
	isSet bool
}

func (v NullableResGuestType) Get() *ResGuestType {
	return v.value
}

func (v *NullableResGuestType) Set(val *ResGuestType) {
	v.value = val
	v.isSet = true
}

func (v NullableResGuestType) IsSet() bool {
	return v.isSet
}

func (v *NullableResGuestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResGuestType(val *ResGuestType) *NullableResGuestType {
	return &NullableResGuestType{value: val, isSet: true}
}

func (v NullableResGuestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResGuestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


