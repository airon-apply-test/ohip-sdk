/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GuestMessageDeliveryMethodType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuestMessageDeliveryMethodType{}

// GuestMessageDeliveryMethodType Guest text message configuration, settings for Text Message Delivery of Guest Messages
type GuestMessageDeliveryMethodType struct {
	TextMessage *TextMessageDeliveryConfigurationType `json:"textMessage,omitempty"`
}

// NewGuestMessageDeliveryMethodType instantiates a new GuestMessageDeliveryMethodType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestMessageDeliveryMethodType() *GuestMessageDeliveryMethodType {
	this := GuestMessageDeliveryMethodType{}
	return &this
}

// NewGuestMessageDeliveryMethodTypeWithDefaults instantiates a new GuestMessageDeliveryMethodType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestMessageDeliveryMethodTypeWithDefaults() *GuestMessageDeliveryMethodType {
	this := GuestMessageDeliveryMethodType{}
	return &this
}

// GetTextMessage returns the TextMessage field value if set, zero value otherwise.
func (o *GuestMessageDeliveryMethodType) GetTextMessage() TextMessageDeliveryConfigurationType {
	if o == nil || IsNil(o.TextMessage) {
		var ret TextMessageDeliveryConfigurationType
		return ret
	}
	return *o.TextMessage
}

// GetTextMessageOk returns a tuple with the TextMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestMessageDeliveryMethodType) GetTextMessageOk() (*TextMessageDeliveryConfigurationType, bool) {
	if o == nil || IsNil(o.TextMessage) {
		return nil, false
	}
	return o.TextMessage, true
}

// HasTextMessage returns a boolean if a field has been set.
func (o *GuestMessageDeliveryMethodType) HasTextMessage() bool {
	if o != nil && !IsNil(o.TextMessage) {
		return true
	}

	return false
}

// SetTextMessage gets a reference to the given TextMessageDeliveryConfigurationType and assigns it to the TextMessage field.
func (o *GuestMessageDeliveryMethodType) SetTextMessage(v TextMessageDeliveryConfigurationType) {
	o.TextMessage = &v
}

func (o GuestMessageDeliveryMethodType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuestMessageDeliveryMethodType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TextMessage) {
		toSerialize["textMessage"] = o.TextMessage
	}
	return toSerialize, nil
}

type NullableGuestMessageDeliveryMethodType struct {
	value *GuestMessageDeliveryMethodType
	isSet bool
}

func (v NullableGuestMessageDeliveryMethodType) Get() *GuestMessageDeliveryMethodType {
	return v.value
}

func (v *NullableGuestMessageDeliveryMethodType) Set(val *GuestMessageDeliveryMethodType) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestMessageDeliveryMethodType) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestMessageDeliveryMethodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestMessageDeliveryMethodType(val *GuestMessageDeliveryMethodType) *NullableGuestMessageDeliveryMethodType {
	return &NullableGuestMessageDeliveryMethodType{value: val, isSet: true}
}

func (v NullableGuestMessageDeliveryMethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestMessageDeliveryMethodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


