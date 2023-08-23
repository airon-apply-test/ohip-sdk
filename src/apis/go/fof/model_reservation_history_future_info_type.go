/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ReservationHistoryFutureInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationHistoryFutureInfoType{}

// ReservationHistoryFutureInfoType Information of History and Future Reservation details attached to Profiles.
type ReservationHistoryFutureInfoType struct {
	HistoryList *HistoryListType `json:"historyList,omitempty"`
	FutureList *FutureListType `json:"futureList,omitempty"`
}

// NewReservationHistoryFutureInfoType instantiates a new ReservationHistoryFutureInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationHistoryFutureInfoType() *ReservationHistoryFutureInfoType {
	this := ReservationHistoryFutureInfoType{}
	return &this
}

// NewReservationHistoryFutureInfoTypeWithDefaults instantiates a new ReservationHistoryFutureInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationHistoryFutureInfoTypeWithDefaults() *ReservationHistoryFutureInfoType {
	this := ReservationHistoryFutureInfoType{}
	return &this
}

// GetHistoryList returns the HistoryList field value if set, zero value otherwise.
func (o *ReservationHistoryFutureInfoType) GetHistoryList() HistoryListType {
	if o == nil || IsNil(o.HistoryList) {
		var ret HistoryListType
		return ret
	}
	return *o.HistoryList
}

// GetHistoryListOk returns a tuple with the HistoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationHistoryFutureInfoType) GetHistoryListOk() (*HistoryListType, bool) {
	if o == nil || IsNil(o.HistoryList) {
		return nil, false
	}
	return o.HistoryList, true
}

// HasHistoryList returns a boolean if a field has been set.
func (o *ReservationHistoryFutureInfoType) HasHistoryList() bool {
	if o != nil && !IsNil(o.HistoryList) {
		return true
	}

	return false
}

// SetHistoryList gets a reference to the given HistoryListType and assigns it to the HistoryList field.
func (o *ReservationHistoryFutureInfoType) SetHistoryList(v HistoryListType) {
	o.HistoryList = &v
}

// GetFutureList returns the FutureList field value if set, zero value otherwise.
func (o *ReservationHistoryFutureInfoType) GetFutureList() FutureListType {
	if o == nil || IsNil(o.FutureList) {
		var ret FutureListType
		return ret
	}
	return *o.FutureList
}

// GetFutureListOk returns a tuple with the FutureList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationHistoryFutureInfoType) GetFutureListOk() (*FutureListType, bool) {
	if o == nil || IsNil(o.FutureList) {
		return nil, false
	}
	return o.FutureList, true
}

// HasFutureList returns a boolean if a field has been set.
func (o *ReservationHistoryFutureInfoType) HasFutureList() bool {
	if o != nil && !IsNil(o.FutureList) {
		return true
	}

	return false
}

// SetFutureList gets a reference to the given FutureListType and assigns it to the FutureList field.
func (o *ReservationHistoryFutureInfoType) SetFutureList(v FutureListType) {
	o.FutureList = &v
}

func (o ReservationHistoryFutureInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationHistoryFutureInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HistoryList) {
		toSerialize["historyList"] = o.HistoryList
	}
	if !IsNil(o.FutureList) {
		toSerialize["futureList"] = o.FutureList
	}
	return toSerialize, nil
}

type NullableReservationHistoryFutureInfoType struct {
	value *ReservationHistoryFutureInfoType
	isSet bool
}

func (v NullableReservationHistoryFutureInfoType) Get() *ReservationHistoryFutureInfoType {
	return v.value
}

func (v *NullableReservationHistoryFutureInfoType) Set(val *ReservationHistoryFutureInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationHistoryFutureInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationHistoryFutureInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationHistoryFutureInfoType(val *ReservationHistoryFutureInfoType) *NullableReservationHistoryFutureInfoType {
	return &NullableReservationHistoryFutureInfoType{value: val, isSet: true}
}

func (v NullableReservationHistoryFutureInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationHistoryFutureInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


