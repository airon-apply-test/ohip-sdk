/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
)

// checks if the HotelCategoriesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelCategoriesDetails{}

// HotelCategoriesDetails Response object for fetching Hotel Categories.
type HotelCategoriesDetails struct {
	// List of Hotel Categories.
	HotelCategories []HotelCategoryType `json:"hotelCategories,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewHotelCategoriesDetails instantiates a new HotelCategoriesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelCategoriesDetails() *HotelCategoriesDetails {
	this := HotelCategoriesDetails{}
	return &this
}

// NewHotelCategoriesDetailsWithDefaults instantiates a new HotelCategoriesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelCategoriesDetailsWithDefaults() *HotelCategoriesDetails {
	this := HotelCategoriesDetails{}
	return &this
}

// GetHotelCategories returns the HotelCategories field value if set, zero value otherwise.
func (o *HotelCategoriesDetails) GetHotelCategories() []HotelCategoryType {
	if o == nil || IsNil(o.HotelCategories) {
		var ret []HotelCategoryType
		return ret
	}
	return o.HotelCategories
}

// GetHotelCategoriesOk returns a tuple with the HotelCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelCategoriesDetails) GetHotelCategoriesOk() ([]HotelCategoryType, bool) {
	if o == nil || IsNil(o.HotelCategories) {
		return nil, false
	}
	return o.HotelCategories, true
}

// HasHotelCategories returns a boolean if a field has been set.
func (o *HotelCategoriesDetails) HasHotelCategories() bool {
	if o != nil && !IsNil(o.HotelCategories) {
		return true
	}

	return false
}

// SetHotelCategories gets a reference to the given []HotelCategoryType and assigns it to the HotelCategories field.
func (o *HotelCategoriesDetails) SetHotelCategories(v []HotelCategoryType) {
	o.HotelCategories = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *HotelCategoriesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelCategoriesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HotelCategoriesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *HotelCategoriesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *HotelCategoriesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelCategoriesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *HotelCategoriesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *HotelCategoriesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o HotelCategoriesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelCategoriesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelCategories) {
		toSerialize["hotelCategories"] = o.HotelCategories
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableHotelCategoriesDetails struct {
	value *HotelCategoriesDetails
	isSet bool
}

func (v NullableHotelCategoriesDetails) Get() *HotelCategoriesDetails {
	return v.value
}

func (v *NullableHotelCategoriesDetails) Set(val *HotelCategoriesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelCategoriesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelCategoriesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelCategoriesDetails(val *HotelCategoriesDetails) *NullableHotelCategoriesDetails {
	return &NullableHotelCategoriesDetails{value: val, isSet: true}
}

func (v NullableHotelCategoriesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelCategoriesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


