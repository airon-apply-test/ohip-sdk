/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChannelCardTypeMappingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelCardTypeMappingType{}

// ChannelCardTypeMappingType Channel-hotel card type mapping information.
type ChannelCardTypeMappingType struct {
	// New channel card type for the existing card type mapping.
	NewChannelCardType *string `json:"newChannelCardType,omitempty"`
	// Inactive date of card type mapping.
	InactiveDate *string `json:"inactiveDate,omitempty"`
	// Hotel code of card type mapping.
	HotelId *string `json:"hotelId,omitempty"`
	// Booking channel code of card type mapping.
	BookingChannelCode *string `json:"bookingChannelCode,omitempty"`
	// Card type of card type mapping.
	CardType *string `json:"cardType,omitempty"`
	// Channel card type of card type mapping.
	ChannelCardType *string `json:"channelCardType,omitempty"`
}

// NewChannelCardTypeMappingType instantiates a new ChannelCardTypeMappingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelCardTypeMappingType() *ChannelCardTypeMappingType {
	this := ChannelCardTypeMappingType{}
	return &this
}

// NewChannelCardTypeMappingTypeWithDefaults instantiates a new ChannelCardTypeMappingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelCardTypeMappingTypeWithDefaults() *ChannelCardTypeMappingType {
	this := ChannelCardTypeMappingType{}
	return &this
}

// GetNewChannelCardType returns the NewChannelCardType field value if set, zero value otherwise.
func (o *ChannelCardTypeMappingType) GetNewChannelCardType() string {
	if o == nil || IsNil(o.NewChannelCardType) {
		var ret string
		return ret
	}
	return *o.NewChannelCardType
}

// GetNewChannelCardTypeOk returns a tuple with the NewChannelCardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelCardTypeMappingType) GetNewChannelCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NewChannelCardType) {
		return nil, false
	}
	return o.NewChannelCardType, true
}

// HasNewChannelCardType returns a boolean if a field has been set.
func (o *ChannelCardTypeMappingType) HasNewChannelCardType() bool {
	if o != nil && !IsNil(o.NewChannelCardType) {
		return true
	}

	return false
}

// SetNewChannelCardType gets a reference to the given string and assigns it to the NewChannelCardType field.
func (o *ChannelCardTypeMappingType) SetNewChannelCardType(v string) {
	o.NewChannelCardType = &v
}

// GetInactiveDate returns the InactiveDate field value if set, zero value otherwise.
func (o *ChannelCardTypeMappingType) GetInactiveDate() string {
	if o == nil || IsNil(o.InactiveDate) {
		var ret string
		return ret
	}
	return *o.InactiveDate
}

// GetInactiveDateOk returns a tuple with the InactiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelCardTypeMappingType) GetInactiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.InactiveDate) {
		return nil, false
	}
	return o.InactiveDate, true
}

// HasInactiveDate returns a boolean if a field has been set.
func (o *ChannelCardTypeMappingType) HasInactiveDate() bool {
	if o != nil && !IsNil(o.InactiveDate) {
		return true
	}

	return false
}

// SetInactiveDate gets a reference to the given string and assigns it to the InactiveDate field.
func (o *ChannelCardTypeMappingType) SetInactiveDate(v string) {
	o.InactiveDate = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ChannelCardTypeMappingType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelCardTypeMappingType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ChannelCardTypeMappingType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ChannelCardTypeMappingType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetBookingChannelCode returns the BookingChannelCode field value if set, zero value otherwise.
func (o *ChannelCardTypeMappingType) GetBookingChannelCode() string {
	if o == nil || IsNil(o.BookingChannelCode) {
		var ret string
		return ret
	}
	return *o.BookingChannelCode
}

// GetBookingChannelCodeOk returns a tuple with the BookingChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelCardTypeMappingType) GetBookingChannelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BookingChannelCode) {
		return nil, false
	}
	return o.BookingChannelCode, true
}

// HasBookingChannelCode returns a boolean if a field has been set.
func (o *ChannelCardTypeMappingType) HasBookingChannelCode() bool {
	if o != nil && !IsNil(o.BookingChannelCode) {
		return true
	}

	return false
}

// SetBookingChannelCode gets a reference to the given string and assigns it to the BookingChannelCode field.
func (o *ChannelCardTypeMappingType) SetBookingChannelCode(v string) {
	o.BookingChannelCode = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *ChannelCardTypeMappingType) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelCardTypeMappingType) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *ChannelCardTypeMappingType) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *ChannelCardTypeMappingType) SetCardType(v string) {
	o.CardType = &v
}

// GetChannelCardType returns the ChannelCardType field value if set, zero value otherwise.
func (o *ChannelCardTypeMappingType) GetChannelCardType() string {
	if o == nil || IsNil(o.ChannelCardType) {
		var ret string
		return ret
	}
	return *o.ChannelCardType
}

// GetChannelCardTypeOk returns a tuple with the ChannelCardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelCardTypeMappingType) GetChannelCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelCardType) {
		return nil, false
	}
	return o.ChannelCardType, true
}

// HasChannelCardType returns a boolean if a field has been set.
func (o *ChannelCardTypeMappingType) HasChannelCardType() bool {
	if o != nil && !IsNil(o.ChannelCardType) {
		return true
	}

	return false
}

// SetChannelCardType gets a reference to the given string and assigns it to the ChannelCardType field.
func (o *ChannelCardTypeMappingType) SetChannelCardType(v string) {
	o.ChannelCardType = &v
}

func (o ChannelCardTypeMappingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelCardTypeMappingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewChannelCardType) {
		toSerialize["newChannelCardType"] = o.NewChannelCardType
	}
	if !IsNil(o.InactiveDate) {
		toSerialize["inactiveDate"] = o.InactiveDate
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.BookingChannelCode) {
		toSerialize["bookingChannelCode"] = o.BookingChannelCode
	}
	if !IsNil(o.CardType) {
		toSerialize["cardType"] = o.CardType
	}
	if !IsNil(o.ChannelCardType) {
		toSerialize["channelCardType"] = o.ChannelCardType
	}
	return toSerialize, nil
}

type NullableChannelCardTypeMappingType struct {
	value *ChannelCardTypeMappingType
	isSet bool
}

func (v NullableChannelCardTypeMappingType) Get() *ChannelCardTypeMappingType {
	return v.value
}

func (v *NullableChannelCardTypeMappingType) Set(val *ChannelCardTypeMappingType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelCardTypeMappingType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelCardTypeMappingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelCardTypeMappingType(val *ChannelCardTypeMappingType) *NullableChannelCardTypeMappingType {
	return &NullableChannelCardTypeMappingType{value: val, isSet: true}
}

func (v NullableChannelCardTypeMappingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelCardTypeMappingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


