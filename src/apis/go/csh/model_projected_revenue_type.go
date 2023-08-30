/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the ProjectedRevenueType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectedRevenueType{}

// ProjectedRevenueType Information about projected catering and room revenue for the reservation.
type ProjectedRevenueType struct {
	ProjectedRoomRevenue *CurrencyAmountType `json:"projectedRoomRevenue,omitempty"`
	ProjectedCateringRevenue *CurrencyAmountType `json:"projectedCateringRevenue,omitempty"`
	// Number of room nights for the reservation.
	RoomNights *int32 `json:"roomNights,omitempty"`
}

// NewProjectedRevenueType instantiates a new ProjectedRevenueType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectedRevenueType() *ProjectedRevenueType {
	this := ProjectedRevenueType{}
	return &this
}

// NewProjectedRevenueTypeWithDefaults instantiates a new ProjectedRevenueType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectedRevenueTypeWithDefaults() *ProjectedRevenueType {
	this := ProjectedRevenueType{}
	return &this
}

// GetProjectedRoomRevenue returns the ProjectedRoomRevenue field value if set, zero value otherwise.
func (o *ProjectedRevenueType) GetProjectedRoomRevenue() CurrencyAmountType {
	if o == nil || IsNil(o.ProjectedRoomRevenue) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.ProjectedRoomRevenue
}

// GetProjectedRoomRevenueOk returns a tuple with the ProjectedRoomRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedRevenueType) GetProjectedRoomRevenueOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.ProjectedRoomRevenue) {
		return nil, false
	}
	return o.ProjectedRoomRevenue, true
}

// HasProjectedRoomRevenue returns a boolean if a field has been set.
func (o *ProjectedRevenueType) HasProjectedRoomRevenue() bool {
	if o != nil && !IsNil(o.ProjectedRoomRevenue) {
		return true
	}

	return false
}

// SetProjectedRoomRevenue gets a reference to the given CurrencyAmountType and assigns it to the ProjectedRoomRevenue field.
func (o *ProjectedRevenueType) SetProjectedRoomRevenue(v CurrencyAmountType) {
	o.ProjectedRoomRevenue = &v
}

// GetProjectedCateringRevenue returns the ProjectedCateringRevenue field value if set, zero value otherwise.
func (o *ProjectedRevenueType) GetProjectedCateringRevenue() CurrencyAmountType {
	if o == nil || IsNil(o.ProjectedCateringRevenue) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.ProjectedCateringRevenue
}

// GetProjectedCateringRevenueOk returns a tuple with the ProjectedCateringRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedRevenueType) GetProjectedCateringRevenueOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.ProjectedCateringRevenue) {
		return nil, false
	}
	return o.ProjectedCateringRevenue, true
}

// HasProjectedCateringRevenue returns a boolean if a field has been set.
func (o *ProjectedRevenueType) HasProjectedCateringRevenue() bool {
	if o != nil && !IsNil(o.ProjectedCateringRevenue) {
		return true
	}

	return false
}

// SetProjectedCateringRevenue gets a reference to the given CurrencyAmountType and assigns it to the ProjectedCateringRevenue field.
func (o *ProjectedRevenueType) SetProjectedCateringRevenue(v CurrencyAmountType) {
	o.ProjectedCateringRevenue = &v
}

// GetRoomNights returns the RoomNights field value if set, zero value otherwise.
func (o *ProjectedRevenueType) GetRoomNights() int32 {
	if o == nil || IsNil(o.RoomNights) {
		var ret int32
		return ret
	}
	return *o.RoomNights
}

// GetRoomNightsOk returns a tuple with the RoomNights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedRevenueType) GetRoomNightsOk() (*int32, bool) {
	if o == nil || IsNil(o.RoomNights) {
		return nil, false
	}
	return o.RoomNights, true
}

// HasRoomNights returns a boolean if a field has been set.
func (o *ProjectedRevenueType) HasRoomNights() bool {
	if o != nil && !IsNil(o.RoomNights) {
		return true
	}

	return false
}

// SetRoomNights gets a reference to the given int32 and assigns it to the RoomNights field.
func (o *ProjectedRevenueType) SetRoomNights(v int32) {
	o.RoomNights = &v
}

func (o ProjectedRevenueType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectedRevenueType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectedRoomRevenue) {
		toSerialize["projectedRoomRevenue"] = o.ProjectedRoomRevenue
	}
	if !IsNil(o.ProjectedCateringRevenue) {
		toSerialize["projectedCateringRevenue"] = o.ProjectedCateringRevenue
	}
	if !IsNil(o.RoomNights) {
		toSerialize["roomNights"] = o.RoomNights
	}
	return toSerialize, nil
}

type NullableProjectedRevenueType struct {
	value *ProjectedRevenueType
	isSet bool
}

func (v NullableProjectedRevenueType) Get() *ProjectedRevenueType {
	return v.value
}

func (v *NullableProjectedRevenueType) Set(val *ProjectedRevenueType) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectedRevenueType) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectedRevenueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectedRevenueType(val *ProjectedRevenueType) *NullableProjectedRevenueType {
	return &NullableProjectedRevenueType{value: val, isSet: true}
}

func (v NullableProjectedRevenueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectedRevenueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


