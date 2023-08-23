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

// checks if the ChannelGuaranteeCodeMappingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelGuaranteeCodeMappingType{}

// ChannelGuaranteeCodeMappingType Channel-hotel guarantee code mapping information.
type ChannelGuaranteeCodeMappingType struct {
	// New channel guarantee code for the existing guarantee code mapping.
	NewChannelGuaranteeCode *string `json:"newChannelGuaranteeCode,omitempty"`
	// Payment type of guarantee code mapping.
	PaymentType *string `json:"paymentType,omitempty"`
	// Inactive date of guarantee code mapping.
	InactiveDate *string `json:"inactiveDate,omitempty"`
	// Hotel code of guarantee code mapping.
	HotelId *string `json:"hotelId,omitempty"`
	// Booking channel code of guarantee code mapping.
	BookingChannelCode *string `json:"bookingChannelCode,omitempty"`
	// Guarantee code of guarantee code mapping.
	GuaranteeCode *string `json:"guaranteeCode,omitempty"`
	// Channel guarantee code of guarantee code mapping.
	ChannelGuaranteeCode *string `json:"channelGuaranteeCode,omitempty"`
}

// NewChannelGuaranteeCodeMappingType instantiates a new ChannelGuaranteeCodeMappingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelGuaranteeCodeMappingType() *ChannelGuaranteeCodeMappingType {
	this := ChannelGuaranteeCodeMappingType{}
	return &this
}

// NewChannelGuaranteeCodeMappingTypeWithDefaults instantiates a new ChannelGuaranteeCodeMappingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelGuaranteeCodeMappingTypeWithDefaults() *ChannelGuaranteeCodeMappingType {
	this := ChannelGuaranteeCodeMappingType{}
	return &this
}

// GetNewChannelGuaranteeCode returns the NewChannelGuaranteeCode field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappingType) GetNewChannelGuaranteeCode() string {
	if o == nil || IsNil(o.NewChannelGuaranteeCode) {
		var ret string
		return ret
	}
	return *o.NewChannelGuaranteeCode
}

// GetNewChannelGuaranteeCodeOk returns a tuple with the NewChannelGuaranteeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappingType) GetNewChannelGuaranteeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.NewChannelGuaranteeCode) {
		return nil, false
	}
	return o.NewChannelGuaranteeCode, true
}

// HasNewChannelGuaranteeCode returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappingType) HasNewChannelGuaranteeCode() bool {
	if o != nil && !IsNil(o.NewChannelGuaranteeCode) {
		return true
	}

	return false
}

// SetNewChannelGuaranteeCode gets a reference to the given string and assigns it to the NewChannelGuaranteeCode field.
func (o *ChannelGuaranteeCodeMappingType) SetNewChannelGuaranteeCode(v string) {
	o.NewChannelGuaranteeCode = &v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappingType) GetPaymentType() string {
	if o == nil || IsNil(o.PaymentType) {
		var ret string
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappingType) GetPaymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentType) {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappingType) HasPaymentType() bool {
	if o != nil && !IsNil(o.PaymentType) {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given string and assigns it to the PaymentType field.
func (o *ChannelGuaranteeCodeMappingType) SetPaymentType(v string) {
	o.PaymentType = &v
}

// GetInactiveDate returns the InactiveDate field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappingType) GetInactiveDate() string {
	if o == nil || IsNil(o.InactiveDate) {
		var ret string
		return ret
	}
	return *o.InactiveDate
}

// GetInactiveDateOk returns a tuple with the InactiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappingType) GetInactiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.InactiveDate) {
		return nil, false
	}
	return o.InactiveDate, true
}

// HasInactiveDate returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappingType) HasInactiveDate() bool {
	if o != nil && !IsNil(o.InactiveDate) {
		return true
	}

	return false
}

// SetInactiveDate gets a reference to the given string and assigns it to the InactiveDate field.
func (o *ChannelGuaranteeCodeMappingType) SetInactiveDate(v string) {
	o.InactiveDate = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappingType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappingType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappingType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ChannelGuaranteeCodeMappingType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetBookingChannelCode returns the BookingChannelCode field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappingType) GetBookingChannelCode() string {
	if o == nil || IsNil(o.BookingChannelCode) {
		var ret string
		return ret
	}
	return *o.BookingChannelCode
}

// GetBookingChannelCodeOk returns a tuple with the BookingChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappingType) GetBookingChannelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BookingChannelCode) {
		return nil, false
	}
	return o.BookingChannelCode, true
}

// HasBookingChannelCode returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappingType) HasBookingChannelCode() bool {
	if o != nil && !IsNil(o.BookingChannelCode) {
		return true
	}

	return false
}

// SetBookingChannelCode gets a reference to the given string and assigns it to the BookingChannelCode field.
func (o *ChannelGuaranteeCodeMappingType) SetBookingChannelCode(v string) {
	o.BookingChannelCode = &v
}

// GetGuaranteeCode returns the GuaranteeCode field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappingType) GetGuaranteeCode() string {
	if o == nil || IsNil(o.GuaranteeCode) {
		var ret string
		return ret
	}
	return *o.GuaranteeCode
}

// GetGuaranteeCodeOk returns a tuple with the GuaranteeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappingType) GetGuaranteeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.GuaranteeCode) {
		return nil, false
	}
	return o.GuaranteeCode, true
}

// HasGuaranteeCode returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappingType) HasGuaranteeCode() bool {
	if o != nil && !IsNil(o.GuaranteeCode) {
		return true
	}

	return false
}

// SetGuaranteeCode gets a reference to the given string and assigns it to the GuaranteeCode field.
func (o *ChannelGuaranteeCodeMappingType) SetGuaranteeCode(v string) {
	o.GuaranteeCode = &v
}

// GetChannelGuaranteeCode returns the ChannelGuaranteeCode field value if set, zero value otherwise.
func (o *ChannelGuaranteeCodeMappingType) GetChannelGuaranteeCode() string {
	if o == nil || IsNil(o.ChannelGuaranteeCode) {
		var ret string
		return ret
	}
	return *o.ChannelGuaranteeCode
}

// GetChannelGuaranteeCodeOk returns a tuple with the ChannelGuaranteeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelGuaranteeCodeMappingType) GetChannelGuaranteeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelGuaranteeCode) {
		return nil, false
	}
	return o.ChannelGuaranteeCode, true
}

// HasChannelGuaranteeCode returns a boolean if a field has been set.
func (o *ChannelGuaranteeCodeMappingType) HasChannelGuaranteeCode() bool {
	if o != nil && !IsNil(o.ChannelGuaranteeCode) {
		return true
	}

	return false
}

// SetChannelGuaranteeCode gets a reference to the given string and assigns it to the ChannelGuaranteeCode field.
func (o *ChannelGuaranteeCodeMappingType) SetChannelGuaranteeCode(v string) {
	o.ChannelGuaranteeCode = &v
}

func (o ChannelGuaranteeCodeMappingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelGuaranteeCodeMappingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewChannelGuaranteeCode) {
		toSerialize["newChannelGuaranteeCode"] = o.NewChannelGuaranteeCode
	}
	if !IsNil(o.PaymentType) {
		toSerialize["paymentType"] = o.PaymentType
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
	if !IsNil(o.GuaranteeCode) {
		toSerialize["guaranteeCode"] = o.GuaranteeCode
	}
	if !IsNil(o.ChannelGuaranteeCode) {
		toSerialize["channelGuaranteeCode"] = o.ChannelGuaranteeCode
	}
	return toSerialize, nil
}

type NullableChannelGuaranteeCodeMappingType struct {
	value *ChannelGuaranteeCodeMappingType
	isSet bool
}

func (v NullableChannelGuaranteeCodeMappingType) Get() *ChannelGuaranteeCodeMappingType {
	return v.value
}

func (v *NullableChannelGuaranteeCodeMappingType) Set(val *ChannelGuaranteeCodeMappingType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelGuaranteeCodeMappingType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelGuaranteeCodeMappingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelGuaranteeCodeMappingType(val *ChannelGuaranteeCodeMappingType) *NullableChannelGuaranteeCodeMappingType {
	return &NullableChannelGuaranteeCodeMappingType{value: val, isSet: true}
}

func (v NullableChannelGuaranteeCodeMappingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelGuaranteeCodeMappingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


