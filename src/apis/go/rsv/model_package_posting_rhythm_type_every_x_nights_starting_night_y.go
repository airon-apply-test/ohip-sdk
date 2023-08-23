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

// checks if the PackagePostingRhythmTypeEveryXNightsStartingNightY type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackagePostingRhythmTypeEveryXNightsStartingNightY{}

// PackagePostingRhythmTypeEveryXNightsStartingNightY Post the package every X nights Staring from Night Y.
type PackagePostingRhythmTypeEveryXNightsStartingNightY struct {
	// Starting night for consumption and posting prices.
	StartOnNightY *int32 `json:"startOnNightY,omitempty"`
	// Occurance of nights when the package needs to be posted and consumed.
	PostEveryXNights *int32 `json:"postEveryXNights,omitempty"`
}

// NewPackagePostingRhythmTypeEveryXNightsStartingNightY instantiates a new PackagePostingRhythmTypeEveryXNightsStartingNightY object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagePostingRhythmTypeEveryXNightsStartingNightY() *PackagePostingRhythmTypeEveryXNightsStartingNightY {
	this := PackagePostingRhythmTypeEveryXNightsStartingNightY{}
	return &this
}

// NewPackagePostingRhythmTypeEveryXNightsStartingNightYWithDefaults instantiates a new PackagePostingRhythmTypeEveryXNightsStartingNightY object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagePostingRhythmTypeEveryXNightsStartingNightYWithDefaults() *PackagePostingRhythmTypeEveryXNightsStartingNightY {
	this := PackagePostingRhythmTypeEveryXNightsStartingNightY{}
	return &this
}

// GetStartOnNightY returns the StartOnNightY field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeEveryXNightsStartingNightY) GetStartOnNightY() int32 {
	if o == nil || IsNil(o.StartOnNightY) {
		var ret int32
		return ret
	}
	return *o.StartOnNightY
}

// GetStartOnNightYOk returns a tuple with the StartOnNightY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeEveryXNightsStartingNightY) GetStartOnNightYOk() (*int32, bool) {
	if o == nil || IsNil(o.StartOnNightY) {
		return nil, false
	}
	return o.StartOnNightY, true
}

// HasStartOnNightY returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeEveryXNightsStartingNightY) HasStartOnNightY() bool {
	if o != nil && !IsNil(o.StartOnNightY) {
		return true
	}

	return false
}

// SetStartOnNightY gets a reference to the given int32 and assigns it to the StartOnNightY field.
func (o *PackagePostingRhythmTypeEveryXNightsStartingNightY) SetStartOnNightY(v int32) {
	o.StartOnNightY = &v
}

// GetPostEveryXNights returns the PostEveryXNights field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeEveryXNightsStartingNightY) GetPostEveryXNights() int32 {
	if o == nil || IsNil(o.PostEveryXNights) {
		var ret int32
		return ret
	}
	return *o.PostEveryXNights
}

// GetPostEveryXNightsOk returns a tuple with the PostEveryXNights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeEveryXNightsStartingNightY) GetPostEveryXNightsOk() (*int32, bool) {
	if o == nil || IsNil(o.PostEveryXNights) {
		return nil, false
	}
	return o.PostEveryXNights, true
}

// HasPostEveryXNights returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeEveryXNightsStartingNightY) HasPostEveryXNights() bool {
	if o != nil && !IsNil(o.PostEveryXNights) {
		return true
	}

	return false
}

// SetPostEveryXNights gets a reference to the given int32 and assigns it to the PostEveryXNights field.
func (o *PackagePostingRhythmTypeEveryXNightsStartingNightY) SetPostEveryXNights(v int32) {
	o.PostEveryXNights = &v
}

func (o PackagePostingRhythmTypeEveryXNightsStartingNightY) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackagePostingRhythmTypeEveryXNightsStartingNightY) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartOnNightY) {
		toSerialize["startOnNightY"] = o.StartOnNightY
	}
	if !IsNil(o.PostEveryXNights) {
		toSerialize["postEveryXNights"] = o.PostEveryXNights
	}
	return toSerialize, nil
}

type NullablePackagePostingRhythmTypeEveryXNightsStartingNightY struct {
	value *PackagePostingRhythmTypeEveryXNightsStartingNightY
	isSet bool
}

func (v NullablePackagePostingRhythmTypeEveryXNightsStartingNightY) Get() *PackagePostingRhythmTypeEveryXNightsStartingNightY {
	return v.value
}

func (v *NullablePackagePostingRhythmTypeEveryXNightsStartingNightY) Set(val *PackagePostingRhythmTypeEveryXNightsStartingNightY) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagePostingRhythmTypeEveryXNightsStartingNightY) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagePostingRhythmTypeEveryXNightsStartingNightY) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagePostingRhythmTypeEveryXNightsStartingNightY(val *PackagePostingRhythmTypeEveryXNightsStartingNightY) *NullablePackagePostingRhythmTypeEveryXNightsStartingNightY {
	return &NullablePackagePostingRhythmTypeEveryXNightsStartingNightY{value: val, isSet: true}
}

func (v NullablePackagePostingRhythmTypeEveryXNightsStartingNightY) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagePostingRhythmTypeEveryXNightsStartingNightY) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


