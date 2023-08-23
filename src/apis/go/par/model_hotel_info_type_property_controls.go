/*
OPERA Cloud Price Availability Rate API

APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HotelInfoTypePropertyControls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelInfoTypePropertyControls{}

// HotelInfoTypePropertyControls Property controls information configuration of the hotel.
type HotelInfoTypePropertyControls struct {
	SellControls *HotelInfoTypePropertyControlsSellControls `json:"sellControls,omitempty"`
	CurrencyFormatting *HotelInfoTypePropertyControlsCurrencyFormatting `json:"currencyFormatting,omitempty"`
	CateringCurrencyFormatting *HotelInfoTypePropertyControlsCateringCurrencyFormatting `json:"cateringCurrencyFormatting,omitempty"`
	DateTimeFormatting *HotelInfoTypePropertyControlsDateTimeFormatting `json:"dateTimeFormatting,omitempty"`
	ApplicationMode *HotelInfoTypePropertyControlsApplicationMode `json:"applicationMode,omitempty"`
}

// NewHotelInfoTypePropertyControls instantiates a new HotelInfoTypePropertyControls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelInfoTypePropertyControls() *HotelInfoTypePropertyControls {
	this := HotelInfoTypePropertyControls{}
	return &this
}

// NewHotelInfoTypePropertyControlsWithDefaults instantiates a new HotelInfoTypePropertyControls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelInfoTypePropertyControlsWithDefaults() *HotelInfoTypePropertyControls {
	this := HotelInfoTypePropertyControls{}
	return &this
}

// GetSellControls returns the SellControls field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControls) GetSellControls() HotelInfoTypePropertyControlsSellControls {
	if o == nil || IsNil(o.SellControls) {
		var ret HotelInfoTypePropertyControlsSellControls
		return ret
	}
	return *o.SellControls
}

// GetSellControlsOk returns a tuple with the SellControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControls) GetSellControlsOk() (*HotelInfoTypePropertyControlsSellControls, bool) {
	if o == nil || IsNil(o.SellControls) {
		return nil, false
	}
	return o.SellControls, true
}

// HasSellControls returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControls) HasSellControls() bool {
	if o != nil && !IsNil(o.SellControls) {
		return true
	}

	return false
}

// SetSellControls gets a reference to the given HotelInfoTypePropertyControlsSellControls and assigns it to the SellControls field.
func (o *HotelInfoTypePropertyControls) SetSellControls(v HotelInfoTypePropertyControlsSellControls) {
	o.SellControls = &v
}

// GetCurrencyFormatting returns the CurrencyFormatting field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControls) GetCurrencyFormatting() HotelInfoTypePropertyControlsCurrencyFormatting {
	if o == nil || IsNil(o.CurrencyFormatting) {
		var ret HotelInfoTypePropertyControlsCurrencyFormatting
		return ret
	}
	return *o.CurrencyFormatting
}

// GetCurrencyFormattingOk returns a tuple with the CurrencyFormatting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControls) GetCurrencyFormattingOk() (*HotelInfoTypePropertyControlsCurrencyFormatting, bool) {
	if o == nil || IsNil(o.CurrencyFormatting) {
		return nil, false
	}
	return o.CurrencyFormatting, true
}

// HasCurrencyFormatting returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControls) HasCurrencyFormatting() bool {
	if o != nil && !IsNil(o.CurrencyFormatting) {
		return true
	}

	return false
}

// SetCurrencyFormatting gets a reference to the given HotelInfoTypePropertyControlsCurrencyFormatting and assigns it to the CurrencyFormatting field.
func (o *HotelInfoTypePropertyControls) SetCurrencyFormatting(v HotelInfoTypePropertyControlsCurrencyFormatting) {
	o.CurrencyFormatting = &v
}

// GetCateringCurrencyFormatting returns the CateringCurrencyFormatting field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControls) GetCateringCurrencyFormatting() HotelInfoTypePropertyControlsCateringCurrencyFormatting {
	if o == nil || IsNil(o.CateringCurrencyFormatting) {
		var ret HotelInfoTypePropertyControlsCateringCurrencyFormatting
		return ret
	}
	return *o.CateringCurrencyFormatting
}

// GetCateringCurrencyFormattingOk returns a tuple with the CateringCurrencyFormatting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControls) GetCateringCurrencyFormattingOk() (*HotelInfoTypePropertyControlsCateringCurrencyFormatting, bool) {
	if o == nil || IsNil(o.CateringCurrencyFormatting) {
		return nil, false
	}
	return o.CateringCurrencyFormatting, true
}

// HasCateringCurrencyFormatting returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControls) HasCateringCurrencyFormatting() bool {
	if o != nil && !IsNil(o.CateringCurrencyFormatting) {
		return true
	}

	return false
}

// SetCateringCurrencyFormatting gets a reference to the given HotelInfoTypePropertyControlsCateringCurrencyFormatting and assigns it to the CateringCurrencyFormatting field.
func (o *HotelInfoTypePropertyControls) SetCateringCurrencyFormatting(v HotelInfoTypePropertyControlsCateringCurrencyFormatting) {
	o.CateringCurrencyFormatting = &v
}

// GetDateTimeFormatting returns the DateTimeFormatting field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControls) GetDateTimeFormatting() HotelInfoTypePropertyControlsDateTimeFormatting {
	if o == nil || IsNil(o.DateTimeFormatting) {
		var ret HotelInfoTypePropertyControlsDateTimeFormatting
		return ret
	}
	return *o.DateTimeFormatting
}

// GetDateTimeFormattingOk returns a tuple with the DateTimeFormatting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControls) GetDateTimeFormattingOk() (*HotelInfoTypePropertyControlsDateTimeFormatting, bool) {
	if o == nil || IsNil(o.DateTimeFormatting) {
		return nil, false
	}
	return o.DateTimeFormatting, true
}

// HasDateTimeFormatting returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControls) HasDateTimeFormatting() bool {
	if o != nil && !IsNil(o.DateTimeFormatting) {
		return true
	}

	return false
}

// SetDateTimeFormatting gets a reference to the given HotelInfoTypePropertyControlsDateTimeFormatting and assigns it to the DateTimeFormatting field.
func (o *HotelInfoTypePropertyControls) SetDateTimeFormatting(v HotelInfoTypePropertyControlsDateTimeFormatting) {
	o.DateTimeFormatting = &v
}

// GetApplicationMode returns the ApplicationMode field value if set, zero value otherwise.
func (o *HotelInfoTypePropertyControls) GetApplicationMode() HotelInfoTypePropertyControlsApplicationMode {
	if o == nil || IsNil(o.ApplicationMode) {
		var ret HotelInfoTypePropertyControlsApplicationMode
		return ret
	}
	return *o.ApplicationMode
}

// GetApplicationModeOk returns a tuple with the ApplicationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypePropertyControls) GetApplicationModeOk() (*HotelInfoTypePropertyControlsApplicationMode, bool) {
	if o == nil || IsNil(o.ApplicationMode) {
		return nil, false
	}
	return o.ApplicationMode, true
}

// HasApplicationMode returns a boolean if a field has been set.
func (o *HotelInfoTypePropertyControls) HasApplicationMode() bool {
	if o != nil && !IsNil(o.ApplicationMode) {
		return true
	}

	return false
}

// SetApplicationMode gets a reference to the given HotelInfoTypePropertyControlsApplicationMode and assigns it to the ApplicationMode field.
func (o *HotelInfoTypePropertyControls) SetApplicationMode(v HotelInfoTypePropertyControlsApplicationMode) {
	o.ApplicationMode = &v
}

func (o HotelInfoTypePropertyControls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelInfoTypePropertyControls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SellControls) {
		toSerialize["sellControls"] = o.SellControls
	}
	if !IsNil(o.CurrencyFormatting) {
		toSerialize["currencyFormatting"] = o.CurrencyFormatting
	}
	if !IsNil(o.CateringCurrencyFormatting) {
		toSerialize["cateringCurrencyFormatting"] = o.CateringCurrencyFormatting
	}
	if !IsNil(o.DateTimeFormatting) {
		toSerialize["dateTimeFormatting"] = o.DateTimeFormatting
	}
	if !IsNil(o.ApplicationMode) {
		toSerialize["applicationMode"] = o.ApplicationMode
	}
	return toSerialize, nil
}

type NullableHotelInfoTypePropertyControls struct {
	value *HotelInfoTypePropertyControls
	isSet bool
}

func (v NullableHotelInfoTypePropertyControls) Get() *HotelInfoTypePropertyControls {
	return v.value
}

func (v *NullableHotelInfoTypePropertyControls) Set(val *HotelInfoTypePropertyControls) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelInfoTypePropertyControls) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelInfoTypePropertyControls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelInfoTypePropertyControls(val *HotelInfoTypePropertyControls) *NullableHotelInfoTypePropertyControls {
	return &NullableHotelInfoTypePropertyControls{value: val, isSet: true}
}

func (v NullableHotelInfoTypePropertyControls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelInfoTypePropertyControls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


