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

// checks if the HotelInfoTypeCommunication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HotelInfoTypeCommunication{}

// HotelInfoTypeCommunication Communication information of the hotel.
type HotelInfoTypeCommunication struct {
	PhoneNumber *TelephoneType `json:"phoneNumber,omitempty"`
	TollFreeNumber *TelephoneType `json:"tollFreeNumber,omitempty"`
	FaxNumber *TelephoneType `json:"faxNumber,omitempty"`
	// Email address
	EmailAddress *string `json:"emailAddress,omitempty"`
	WebPage *URLType `json:"webPage,omitempty"`
}

// NewHotelInfoTypeCommunication instantiates a new HotelInfoTypeCommunication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelInfoTypeCommunication() *HotelInfoTypeCommunication {
	this := HotelInfoTypeCommunication{}
	return &this
}

// NewHotelInfoTypeCommunicationWithDefaults instantiates a new HotelInfoTypeCommunication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelInfoTypeCommunicationWithDefaults() *HotelInfoTypeCommunication {
	this := HotelInfoTypeCommunication{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *HotelInfoTypeCommunication) GetPhoneNumber() TelephoneType {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret TelephoneType
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeCommunication) GetPhoneNumberOk() (*TelephoneType, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *HotelInfoTypeCommunication) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given TelephoneType and assigns it to the PhoneNumber field.
func (o *HotelInfoTypeCommunication) SetPhoneNumber(v TelephoneType) {
	o.PhoneNumber = &v
}

// GetTollFreeNumber returns the TollFreeNumber field value if set, zero value otherwise.
func (o *HotelInfoTypeCommunication) GetTollFreeNumber() TelephoneType {
	if o == nil || IsNil(o.TollFreeNumber) {
		var ret TelephoneType
		return ret
	}
	return *o.TollFreeNumber
}

// GetTollFreeNumberOk returns a tuple with the TollFreeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeCommunication) GetTollFreeNumberOk() (*TelephoneType, bool) {
	if o == nil || IsNil(o.TollFreeNumber) {
		return nil, false
	}
	return o.TollFreeNumber, true
}

// HasTollFreeNumber returns a boolean if a field has been set.
func (o *HotelInfoTypeCommunication) HasTollFreeNumber() bool {
	if o != nil && !IsNil(o.TollFreeNumber) {
		return true
	}

	return false
}

// SetTollFreeNumber gets a reference to the given TelephoneType and assigns it to the TollFreeNumber field.
func (o *HotelInfoTypeCommunication) SetTollFreeNumber(v TelephoneType) {
	o.TollFreeNumber = &v
}

// GetFaxNumber returns the FaxNumber field value if set, zero value otherwise.
func (o *HotelInfoTypeCommunication) GetFaxNumber() TelephoneType {
	if o == nil || IsNil(o.FaxNumber) {
		var ret TelephoneType
		return ret
	}
	return *o.FaxNumber
}

// GetFaxNumberOk returns a tuple with the FaxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeCommunication) GetFaxNumberOk() (*TelephoneType, bool) {
	if o == nil || IsNil(o.FaxNumber) {
		return nil, false
	}
	return o.FaxNumber, true
}

// HasFaxNumber returns a boolean if a field has been set.
func (o *HotelInfoTypeCommunication) HasFaxNumber() bool {
	if o != nil && !IsNil(o.FaxNumber) {
		return true
	}

	return false
}

// SetFaxNumber gets a reference to the given TelephoneType and assigns it to the FaxNumber field.
func (o *HotelInfoTypeCommunication) SetFaxNumber(v TelephoneType) {
	o.FaxNumber = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *HotelInfoTypeCommunication) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeCommunication) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *HotelInfoTypeCommunication) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *HotelInfoTypeCommunication) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetWebPage returns the WebPage field value if set, zero value otherwise.
func (o *HotelInfoTypeCommunication) GetWebPage() URLType {
	if o == nil || IsNil(o.WebPage) {
		var ret URLType
		return ret
	}
	return *o.WebPage
}

// GetWebPageOk returns a tuple with the WebPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelInfoTypeCommunication) GetWebPageOk() (*URLType, bool) {
	if o == nil || IsNil(o.WebPage) {
		return nil, false
	}
	return o.WebPage, true
}

// HasWebPage returns a boolean if a field has been set.
func (o *HotelInfoTypeCommunication) HasWebPage() bool {
	if o != nil && !IsNil(o.WebPage) {
		return true
	}

	return false
}

// SetWebPage gets a reference to the given URLType and assigns it to the WebPage field.
func (o *HotelInfoTypeCommunication) SetWebPage(v URLType) {
	o.WebPage = &v
}

func (o HotelInfoTypeCommunication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HotelInfoTypeCommunication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.TollFreeNumber) {
		toSerialize["tollFreeNumber"] = o.TollFreeNumber
	}
	if !IsNil(o.FaxNumber) {
		toSerialize["faxNumber"] = o.FaxNumber
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.WebPage) {
		toSerialize["webPage"] = o.WebPage
	}
	return toSerialize, nil
}

type NullableHotelInfoTypeCommunication struct {
	value *HotelInfoTypeCommunication
	isSet bool
}

func (v NullableHotelInfoTypeCommunication) Get() *HotelInfoTypeCommunication {
	return v.value
}

func (v *NullableHotelInfoTypeCommunication) Set(val *HotelInfoTypeCommunication) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelInfoTypeCommunication) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelInfoTypeCommunication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelInfoTypeCommunication(val *HotelInfoTypeCommunication) *NullableHotelInfoTypeCommunication {
	return &NullableHotelInfoTypeCommunication{value: val, isSet: true}
}

func (v NullableHotelInfoTypeCommunication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelInfoTypeCommunication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


