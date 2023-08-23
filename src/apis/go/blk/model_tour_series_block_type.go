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

// checks if the TourSeriesBlockType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TourSeriesBlockType{}

// TourSeriesBlockType Information about a tour series block.
type TourSeriesBlockType struct {
	// The hotel code of the tour series block.
	HotelId *string `json:"hotelId,omitempty"`
	// The block code of the tour series block.
	BlockCode *string `json:"blockCode,omitempty"`
	// The start date of the tour series block.
	StartDate *string `json:"startDate,omitempty"`
	// The booking status of the tour series block.
	BlockStatus *string `json:"blockStatus,omitempty"`
	// The block name of the tour series block.
	BlockName *string `json:"blockName,omitempty"`
}

// NewTourSeriesBlockType instantiates a new TourSeriesBlockType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTourSeriesBlockType() *TourSeriesBlockType {
	this := TourSeriesBlockType{}
	return &this
}

// NewTourSeriesBlockTypeWithDefaults instantiates a new TourSeriesBlockType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTourSeriesBlockTypeWithDefaults() *TourSeriesBlockType {
	this := TourSeriesBlockType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *TourSeriesBlockType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TourSeriesBlockType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *TourSeriesBlockType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *TourSeriesBlockType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetBlockCode returns the BlockCode field value if set, zero value otherwise.
func (o *TourSeriesBlockType) GetBlockCode() string {
	if o == nil || IsNil(o.BlockCode) {
		var ret string
		return ret
	}
	return *o.BlockCode
}

// GetBlockCodeOk returns a tuple with the BlockCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TourSeriesBlockType) GetBlockCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BlockCode) {
		return nil, false
	}
	return o.BlockCode, true
}

// HasBlockCode returns a boolean if a field has been set.
func (o *TourSeriesBlockType) HasBlockCode() bool {
	if o != nil && !IsNil(o.BlockCode) {
		return true
	}

	return false
}

// SetBlockCode gets a reference to the given string and assigns it to the BlockCode field.
func (o *TourSeriesBlockType) SetBlockCode(v string) {
	o.BlockCode = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TourSeriesBlockType) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TourSeriesBlockType) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TourSeriesBlockType) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *TourSeriesBlockType) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBlockStatus returns the BlockStatus field value if set, zero value otherwise.
func (o *TourSeriesBlockType) GetBlockStatus() string {
	if o == nil || IsNil(o.BlockStatus) {
		var ret string
		return ret
	}
	return *o.BlockStatus
}

// GetBlockStatusOk returns a tuple with the BlockStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TourSeriesBlockType) GetBlockStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BlockStatus) {
		return nil, false
	}
	return o.BlockStatus, true
}

// HasBlockStatus returns a boolean if a field has been set.
func (o *TourSeriesBlockType) HasBlockStatus() bool {
	if o != nil && !IsNil(o.BlockStatus) {
		return true
	}

	return false
}

// SetBlockStatus gets a reference to the given string and assigns it to the BlockStatus field.
func (o *TourSeriesBlockType) SetBlockStatus(v string) {
	o.BlockStatus = &v
}

// GetBlockName returns the BlockName field value if set, zero value otherwise.
func (o *TourSeriesBlockType) GetBlockName() string {
	if o == nil || IsNil(o.BlockName) {
		var ret string
		return ret
	}
	return *o.BlockName
}

// GetBlockNameOk returns a tuple with the BlockName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TourSeriesBlockType) GetBlockNameOk() (*string, bool) {
	if o == nil || IsNil(o.BlockName) {
		return nil, false
	}
	return o.BlockName, true
}

// HasBlockName returns a boolean if a field has been set.
func (o *TourSeriesBlockType) HasBlockName() bool {
	if o != nil && !IsNil(o.BlockName) {
		return true
	}

	return false
}

// SetBlockName gets a reference to the given string and assigns it to the BlockName field.
func (o *TourSeriesBlockType) SetBlockName(v string) {
	o.BlockName = &v
}

func (o TourSeriesBlockType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TourSeriesBlockType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.BlockCode) {
		toSerialize["blockCode"] = o.BlockCode
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.BlockStatus) {
		toSerialize["blockStatus"] = o.BlockStatus
	}
	if !IsNil(o.BlockName) {
		toSerialize["blockName"] = o.BlockName
	}
	return toSerialize, nil
}

type NullableTourSeriesBlockType struct {
	value *TourSeriesBlockType
	isSet bool
}

func (v NullableTourSeriesBlockType) Get() *TourSeriesBlockType {
	return v.value
}

func (v *NullableTourSeriesBlockType) Set(val *TourSeriesBlockType) {
	v.value = val
	v.isSet = true
}

func (v NullableTourSeriesBlockType) IsSet() bool {
	return v.isSet
}

func (v *NullableTourSeriesBlockType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTourSeriesBlockType(val *TourSeriesBlockType) *NullableTourSeriesBlockType {
	return &NullableTourSeriesBlockType{value: val, isSet: true}
}

func (v NullableTourSeriesBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTourSeriesBlockType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


