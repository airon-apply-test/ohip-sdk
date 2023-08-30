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

// checks if the HotelRevenueDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelRevenueDetails{}

// HotelRevenueDetails Response object for hotel room revenue for all HHonors redemption reservation.
type HotelRevenueDetails struct {
	HotelRevenue *HotelRevenueType `json:"hotelRevenue,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewHotelRevenueDetails instantiates a new HotelRevenueDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelRevenueDetails() *HotelRevenueDetails {
	this := HotelRevenueDetails{}
	return &this
}

// NewHotelRevenueDetailsWithDefaults instantiates a new HotelRevenueDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelRevenueDetailsWithDefaults() *HotelRevenueDetails {
	this := HotelRevenueDetails{}
	return &this
}

// GetHotelRevenue returns the HotelRevenue field value if set, zero value otherwise.
func (o *HotelRevenueDetails) GetHotelRevenue() HotelRevenueType {
	if o == nil || IsNil(o.HotelRevenue) {
		var ret HotelRevenueType
		return ret
	}
	return *o.HotelRevenue
}

// GetHotelRevenueOk returns a tuple with the HotelRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelRevenueDetails) GetHotelRevenueOk() (*HotelRevenueType, bool) {
	if o == nil || IsNil(o.HotelRevenue) {
		return nil, false
	}
	return o.HotelRevenue, true
}

// HasHotelRevenue returns a boolean if a field has been set.
func (o *HotelRevenueDetails) HasHotelRevenue() bool {
	if o != nil && !IsNil(o.HotelRevenue) {
		return true
	}

	return false
}

// SetHotelRevenue gets a reference to the given HotelRevenueType and assigns it to the HotelRevenue field.
func (o *HotelRevenueDetails) SetHotelRevenue(v HotelRevenueType) {
	o.HotelRevenue = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *HotelRevenueDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelRevenueDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HotelRevenueDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *HotelRevenueDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *HotelRevenueDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelRevenueDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *HotelRevenueDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *HotelRevenueDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o HotelRevenueDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelRevenueDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelRevenue) {
		toSerialize["hotelRevenue"] = o.HotelRevenue
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableHotelRevenueDetails struct {
	value *HotelRevenueDetails
	isSet bool
}

func (v NullableHotelRevenueDetails) Get() *HotelRevenueDetails {
	return v.value
}

func (v *NullableHotelRevenueDetails) Set(val *HotelRevenueDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelRevenueDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelRevenueDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelRevenueDetails(val *HotelRevenueDetails) *NullableHotelRevenueDetails {
	return &NullableHotelRevenueDetails{value: val, isSet: true}
}

func (v NullableHotelRevenueDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelRevenueDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


