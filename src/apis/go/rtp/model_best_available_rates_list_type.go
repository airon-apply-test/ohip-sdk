/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtp

import (
	"encoding/json"
)

// checks if the BestAvailableRatesListType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BestAvailableRatesListType{}

// BestAvailableRatesListType The list of best availabe rates for rate code(s)
type BestAvailableRatesListType struct {
	// Collection of best available rates for rate code(s)
	BestAvailableRates []BestAvailableRateType `json:"bestAvailableRates,omitempty"`
	// The list of references of rate code selected as best available rates
	MasterInfoList []MasterInfo `json:"masterInfoList,omitempty"`
	// Hotel code for which best available rates will be considered.
	HotelId *string `json:"hotelId,omitempty"`
}

// NewBestAvailableRatesListType instantiates a new BestAvailableRatesListType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBestAvailableRatesListType() *BestAvailableRatesListType {
	this := BestAvailableRatesListType{}
	return &this
}

// NewBestAvailableRatesListTypeWithDefaults instantiates a new BestAvailableRatesListType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBestAvailableRatesListTypeWithDefaults() *BestAvailableRatesListType {
	this := BestAvailableRatesListType{}
	return &this
}

// GetBestAvailableRates returns the BestAvailableRates field value if set, zero value otherwise.
func (o *BestAvailableRatesListType) GetBestAvailableRates() []BestAvailableRateType {
	if o == nil || IsNil(o.BestAvailableRates) {
		var ret []BestAvailableRateType
		return ret
	}
	return o.BestAvailableRates
}

// GetBestAvailableRatesOk returns a tuple with the BestAvailableRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesListType) GetBestAvailableRatesOk() ([]BestAvailableRateType, bool) {
	if o == nil || IsNil(o.BestAvailableRates) {
		return nil, false
	}
	return o.BestAvailableRates, true
}

// HasBestAvailableRates returns a boolean if a field has been set.
func (o *BestAvailableRatesListType) HasBestAvailableRates() bool {
	if o != nil && !IsNil(o.BestAvailableRates) {
		return true
	}

	return false
}

// SetBestAvailableRates gets a reference to the given []BestAvailableRateType and assigns it to the BestAvailableRates field.
func (o *BestAvailableRatesListType) SetBestAvailableRates(v []BestAvailableRateType) {
	o.BestAvailableRates = v
}

// GetMasterInfoList returns the MasterInfoList field value if set, zero value otherwise.
func (o *BestAvailableRatesListType) GetMasterInfoList() []MasterInfo {
	if o == nil || IsNil(o.MasterInfoList) {
		var ret []MasterInfo
		return ret
	}
	return o.MasterInfoList
}

// GetMasterInfoListOk returns a tuple with the MasterInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesListType) GetMasterInfoListOk() ([]MasterInfo, bool) {
	if o == nil || IsNil(o.MasterInfoList) {
		return nil, false
	}
	return o.MasterInfoList, true
}

// HasMasterInfoList returns a boolean if a field has been set.
func (o *BestAvailableRatesListType) HasMasterInfoList() bool {
	if o != nil && !IsNil(o.MasterInfoList) {
		return true
	}

	return false
}

// SetMasterInfoList gets a reference to the given []MasterInfo and assigns it to the MasterInfoList field.
func (o *BestAvailableRatesListType) SetMasterInfoList(v []MasterInfo) {
	o.MasterInfoList = v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *BestAvailableRatesListType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesListType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *BestAvailableRatesListType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *BestAvailableRatesListType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o BestAvailableRatesListType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BestAvailableRatesListType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BestAvailableRates) {
		toSerialize["bestAvailableRates"] = o.BestAvailableRates
	}
	if !IsNil(o.MasterInfoList) {
		toSerialize["masterInfoList"] = o.MasterInfoList
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableBestAvailableRatesListType struct {
	value *BestAvailableRatesListType
	isSet bool
}

func (v NullableBestAvailableRatesListType) Get() *BestAvailableRatesListType {
	return v.value
}

func (v *NullableBestAvailableRatesListType) Set(val *BestAvailableRatesListType) {
	v.value = val
	v.isSet = true
}

func (v NullableBestAvailableRatesListType) IsSet() bool {
	return v.isSet
}

func (v *NullableBestAvailableRatesListType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBestAvailableRatesListType(val *BestAvailableRatesListType) *NullableBestAvailableRatesListType {
	return &NullableBestAvailableRatesListType{value: val, isSet: true}
}

func (v NullableBestAvailableRatesListType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBestAvailableRatesListType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


