/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HotelDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelDetails{}

// HotelDetails Response object for fetching Property configuration information.
type HotelDetails struct {
	HotelConfigInfo *HotelInfoType `json:"hotelConfigInfo,omitempty"`
	// Refer to Generic common types document.
	MasterInfoList []MasterInfoType `json:"masterInfoList,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewHotelDetails instantiates a new HotelDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelDetails() *HotelDetails {
	this := HotelDetails{}
	return &this
}

// NewHotelDetailsWithDefaults instantiates a new HotelDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelDetailsWithDefaults() *HotelDetails {
	this := HotelDetails{}
	return &this
}

// GetHotelConfigInfo returns the HotelConfigInfo field value if set, zero value otherwise.
func (o *HotelDetails) GetHotelConfigInfo() HotelInfoType {
	if o == nil || IsNil(o.HotelConfigInfo) {
		var ret HotelInfoType
		return ret
	}
	return *o.HotelConfigInfo
}

// GetHotelConfigInfoOk returns a tuple with the HotelConfigInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelDetails) GetHotelConfigInfoOk() (*HotelInfoType, bool) {
	if o == nil || IsNil(o.HotelConfigInfo) {
		return nil, false
	}
	return o.HotelConfigInfo, true
}

// HasHotelConfigInfo returns a boolean if a field has been set.
func (o *HotelDetails) HasHotelConfigInfo() bool {
	if o != nil && !IsNil(o.HotelConfigInfo) {
		return true
	}

	return false
}

// SetHotelConfigInfo gets a reference to the given HotelInfoType and assigns it to the HotelConfigInfo field.
func (o *HotelDetails) SetHotelConfigInfo(v HotelInfoType) {
	o.HotelConfigInfo = &v
}

// GetMasterInfoList returns the MasterInfoList field value if set, zero value otherwise.
func (o *HotelDetails) GetMasterInfoList() []MasterInfoType {
	if o == nil || IsNil(o.MasterInfoList) {
		var ret []MasterInfoType
		return ret
	}
	return o.MasterInfoList
}

// GetMasterInfoListOk returns a tuple with the MasterInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelDetails) GetMasterInfoListOk() ([]MasterInfoType, bool) {
	if o == nil || IsNil(o.MasterInfoList) {
		return nil, false
	}
	return o.MasterInfoList, true
}

// HasMasterInfoList returns a boolean if a field has been set.
func (o *HotelDetails) HasMasterInfoList() bool {
	if o != nil && !IsNil(o.MasterInfoList) {
		return true
	}

	return false
}

// SetMasterInfoList gets a reference to the given []MasterInfoType and assigns it to the MasterInfoList field.
func (o *HotelDetails) SetMasterInfoList(v []MasterInfoType) {
	o.MasterInfoList = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *HotelDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HotelDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *HotelDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *HotelDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *HotelDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *HotelDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o HotelDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelConfigInfo) {
		toSerialize["hotelConfigInfo"] = o.HotelConfigInfo
	}
	if !IsNil(o.MasterInfoList) {
		toSerialize["masterInfoList"] = o.MasterInfoList
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableHotelDetails struct {
	value *HotelDetails
	isSet bool
}

func (v NullableHotelDetails) Get() *HotelDetails {
	return v.value
}

func (v *NullableHotelDetails) Set(val *HotelDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelDetails(val *HotelDetails) *NullableHotelDetails {
	return &NullableHotelDetails{value: val, isSet: true}
}

func (v NullableHotelDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


