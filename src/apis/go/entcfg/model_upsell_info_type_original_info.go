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

// checks if the UpsellInfoTypeOriginalInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsellInfoTypeOriginalInfo{}

// UpsellInfoTypeOriginalInfo Original Values of the reservation prior to upgrade.
type UpsellInfoTypeOriginalInfo struct {
	// The Rate Code of the Reservation before it was upgraded
	RateCode *string `json:"rateCode,omitempty"`
	TotalAmount *CurrencyAmountType `json:"totalAmount,omitempty"`
	RoomType *CodeDescriptionType `json:"roomType,omitempty"`
	// The number of nights of the reservation before being upgraded.
	Nights *int32 `json:"nights,omitempty"`
}

// NewUpsellInfoTypeOriginalInfo instantiates a new UpsellInfoTypeOriginalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsellInfoTypeOriginalInfo() *UpsellInfoTypeOriginalInfo {
	this := UpsellInfoTypeOriginalInfo{}
	return &this
}

// NewUpsellInfoTypeOriginalInfoWithDefaults instantiates a new UpsellInfoTypeOriginalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsellInfoTypeOriginalInfoWithDefaults() *UpsellInfoTypeOriginalInfo {
	this := UpsellInfoTypeOriginalInfo{}
	return &this
}

// GetRateCode returns the RateCode field value if set, zero value otherwise.
func (o *UpsellInfoTypeOriginalInfo) GetRateCode() string {
	if o == nil || IsNil(o.RateCode) {
		var ret string
		return ret
	}
	return *o.RateCode
}

// GetRateCodeOk returns a tuple with the RateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsellInfoTypeOriginalInfo) GetRateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RateCode) {
		return nil, false
	}
	return o.RateCode, true
}

// HasRateCode returns a boolean if a field has been set.
func (o *UpsellInfoTypeOriginalInfo) HasRateCode() bool {
	if o != nil && !IsNil(o.RateCode) {
		return true
	}

	return false
}

// SetRateCode gets a reference to the given string and assigns it to the RateCode field.
func (o *UpsellInfoTypeOriginalInfo) SetRateCode(v string) {
	o.RateCode = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *UpsellInfoTypeOriginalInfo) GetTotalAmount() CurrencyAmountType {
	if o == nil || IsNil(o.TotalAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsellInfoTypeOriginalInfo) GetTotalAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *UpsellInfoTypeOriginalInfo) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given CurrencyAmountType and assigns it to the TotalAmount field.
func (o *UpsellInfoTypeOriginalInfo) SetTotalAmount(v CurrencyAmountType) {
	o.TotalAmount = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *UpsellInfoTypeOriginalInfo) GetRoomType() CodeDescriptionType {
	if o == nil || IsNil(o.RoomType) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsellInfoTypeOriginalInfo) GetRoomTypeOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *UpsellInfoTypeOriginalInfo) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given CodeDescriptionType and assigns it to the RoomType field.
func (o *UpsellInfoTypeOriginalInfo) SetRoomType(v CodeDescriptionType) {
	o.RoomType = &v
}

// GetNights returns the Nights field value if set, zero value otherwise.
func (o *UpsellInfoTypeOriginalInfo) GetNights() int32 {
	if o == nil || IsNil(o.Nights) {
		var ret int32
		return ret
	}
	return *o.Nights
}

// GetNightsOk returns a tuple with the Nights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsellInfoTypeOriginalInfo) GetNightsOk() (*int32, bool) {
	if o == nil || IsNil(o.Nights) {
		return nil, false
	}
	return o.Nights, true
}

// HasNights returns a boolean if a field has been set.
func (o *UpsellInfoTypeOriginalInfo) HasNights() bool {
	if o != nil && !IsNil(o.Nights) {
		return true
	}

	return false
}

// SetNights gets a reference to the given int32 and assigns it to the Nights field.
func (o *UpsellInfoTypeOriginalInfo) SetNights(v int32) {
	o.Nights = &v
}

func (o UpsellInfoTypeOriginalInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsellInfoTypeOriginalInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RateCode) {
		toSerialize["rateCode"] = o.RateCode
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.Nights) {
		toSerialize["nights"] = o.Nights
	}
	return toSerialize, nil
}

type NullableUpsellInfoTypeOriginalInfo struct {
	value *UpsellInfoTypeOriginalInfo
	isSet bool
}

func (v NullableUpsellInfoTypeOriginalInfo) Get() *UpsellInfoTypeOriginalInfo {
	return v.value
}

func (v *NullableUpsellInfoTypeOriginalInfo) Set(val *UpsellInfoTypeOriginalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsellInfoTypeOriginalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsellInfoTypeOriginalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsellInfoTypeOriginalInfo(val *UpsellInfoTypeOriginalInfo) *NullableUpsellInfoTypeOriginalInfo {
	return &NullableUpsellInfoTypeOriginalInfo{value: val, isSet: true}
}

func (v NullableUpsellInfoTypeOriginalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsellInfoTypeOriginalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


