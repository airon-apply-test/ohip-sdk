/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package act

import (
	"encoding/json"
)

// checks if the ProfileDeliveryMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileDeliveryMethod{}

// ProfileDeliveryMethod Delivery Information type to the profile.
type ProfileDeliveryMethod struct {
	DeliveryId *UniqueIDType `json:"deliveryId,omitempty"`
	// Delivery type can have a value EMAIL, ELECTRONIC etc and it depends on the parameter set in OPERA Control.
	DeliveryType *string `json:"deliveryType,omitempty"`
	// Delivery value holds the corresponding value of the delivery type..
	DeliveryValue *string `json:"deliveryValue,omitempty"`
	// Property that has delivery methods configured.
	HotelId *string `json:"hotelId,omitempty"`
	DeliveryModule *ProfileDeliveryModuleType `json:"deliveryModule,omitempty"`
	// When true, indicates a primary information.
	Primary *bool `json:"primary,omitempty"`
	// Display Order sequence.
	OrderSequence *float32 `json:"orderSequence,omitempty"`
}

// NewProfileDeliveryMethod instantiates a new ProfileDeliveryMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileDeliveryMethod() *ProfileDeliveryMethod {
	this := ProfileDeliveryMethod{}
	return &this
}

// NewProfileDeliveryMethodWithDefaults instantiates a new ProfileDeliveryMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileDeliveryMethodWithDefaults() *ProfileDeliveryMethod {
	this := ProfileDeliveryMethod{}
	return &this
}

// GetDeliveryId returns the DeliveryId field value if set, zero value otherwise.
func (o *ProfileDeliveryMethod) GetDeliveryId() UniqueIDType {
	if o == nil || IsNil(o.DeliveryId) {
		var ret UniqueIDType
		return ret
	}
	return *o.DeliveryId
}

// GetDeliveryIdOk returns a tuple with the DeliveryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileDeliveryMethod) GetDeliveryIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.DeliveryId) {
		return nil, false
	}
	return o.DeliveryId, true
}

// HasDeliveryId returns a boolean if a field has been set.
func (o *ProfileDeliveryMethod) HasDeliveryId() bool {
	if o != nil && !IsNil(o.DeliveryId) {
		return true
	}

	return false
}

// SetDeliveryId gets a reference to the given UniqueIDType and assigns it to the DeliveryId field.
func (o *ProfileDeliveryMethod) SetDeliveryId(v UniqueIDType) {
	o.DeliveryId = &v
}

// GetDeliveryType returns the DeliveryType field value if set, zero value otherwise.
func (o *ProfileDeliveryMethod) GetDeliveryType() string {
	if o == nil || IsNil(o.DeliveryType) {
		var ret string
		return ret
	}
	return *o.DeliveryType
}

// GetDeliveryTypeOk returns a tuple with the DeliveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileDeliveryMethod) GetDeliveryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryType) {
		return nil, false
	}
	return o.DeliveryType, true
}

// HasDeliveryType returns a boolean if a field has been set.
func (o *ProfileDeliveryMethod) HasDeliveryType() bool {
	if o != nil && !IsNil(o.DeliveryType) {
		return true
	}

	return false
}

// SetDeliveryType gets a reference to the given string and assigns it to the DeliveryType field.
func (o *ProfileDeliveryMethod) SetDeliveryType(v string) {
	o.DeliveryType = &v
}

// GetDeliveryValue returns the DeliveryValue field value if set, zero value otherwise.
func (o *ProfileDeliveryMethod) GetDeliveryValue() string {
	if o == nil || IsNil(o.DeliveryValue) {
		var ret string
		return ret
	}
	return *o.DeliveryValue
}

// GetDeliveryValueOk returns a tuple with the DeliveryValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileDeliveryMethod) GetDeliveryValueOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryValue) {
		return nil, false
	}
	return o.DeliveryValue, true
}

// HasDeliveryValue returns a boolean if a field has been set.
func (o *ProfileDeliveryMethod) HasDeliveryValue() bool {
	if o != nil && !IsNil(o.DeliveryValue) {
		return true
	}

	return false
}

// SetDeliveryValue gets a reference to the given string and assigns it to the DeliveryValue field.
func (o *ProfileDeliveryMethod) SetDeliveryValue(v string) {
	o.DeliveryValue = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ProfileDeliveryMethod) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileDeliveryMethod) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ProfileDeliveryMethod) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ProfileDeliveryMethod) SetHotelId(v string) {
	o.HotelId = &v
}

// GetDeliveryModule returns the DeliveryModule field value if set, zero value otherwise.
func (o *ProfileDeliveryMethod) GetDeliveryModule() ProfileDeliveryModuleType {
	if o == nil || IsNil(o.DeliveryModule) {
		var ret ProfileDeliveryModuleType
		return ret
	}
	return *o.DeliveryModule
}

// GetDeliveryModuleOk returns a tuple with the DeliveryModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileDeliveryMethod) GetDeliveryModuleOk() (*ProfileDeliveryModuleType, bool) {
	if o == nil || IsNil(o.DeliveryModule) {
		return nil, false
	}
	return o.DeliveryModule, true
}

// HasDeliveryModule returns a boolean if a field has been set.
func (o *ProfileDeliveryMethod) HasDeliveryModule() bool {
	if o != nil && !IsNil(o.DeliveryModule) {
		return true
	}

	return false
}

// SetDeliveryModule gets a reference to the given ProfileDeliveryModuleType and assigns it to the DeliveryModule field.
func (o *ProfileDeliveryMethod) SetDeliveryModule(v ProfileDeliveryModuleType) {
	o.DeliveryModule = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *ProfileDeliveryMethod) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileDeliveryMethod) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *ProfileDeliveryMethod) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *ProfileDeliveryMethod) SetPrimary(v bool) {
	o.Primary = &v
}

// GetOrderSequence returns the OrderSequence field value if set, zero value otherwise.
func (o *ProfileDeliveryMethod) GetOrderSequence() float32 {
	if o == nil || IsNil(o.OrderSequence) {
		var ret float32
		return ret
	}
	return *o.OrderSequence
}

// GetOrderSequenceOk returns a tuple with the OrderSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileDeliveryMethod) GetOrderSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderSequence) {
		return nil, false
	}
	return o.OrderSequence, true
}

// HasOrderSequence returns a boolean if a field has been set.
func (o *ProfileDeliveryMethod) HasOrderSequence() bool {
	if o != nil && !IsNil(o.OrderSequence) {
		return true
	}

	return false
}

// SetOrderSequence gets a reference to the given float32 and assigns it to the OrderSequence field.
func (o *ProfileDeliveryMethod) SetOrderSequence(v float32) {
	o.OrderSequence = &v
}

func (o ProfileDeliveryMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileDeliveryMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryId) {
		toSerialize["deliveryId"] = o.DeliveryId
	}
	if !IsNil(o.DeliveryType) {
		toSerialize["deliveryType"] = o.DeliveryType
	}
	if !IsNil(o.DeliveryValue) {
		toSerialize["deliveryValue"] = o.DeliveryValue
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.DeliveryModule) {
		toSerialize["deliveryModule"] = o.DeliveryModule
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.OrderSequence) {
		toSerialize["orderSequence"] = o.OrderSequence
	}
	return toSerialize, nil
}

type NullableProfileDeliveryMethod struct {
	value *ProfileDeliveryMethod
	isSet bool
}

func (v NullableProfileDeliveryMethod) Get() *ProfileDeliveryMethod {
	return v.value
}

func (v *NullableProfileDeliveryMethod) Set(val *ProfileDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileDeliveryMethod(val *ProfileDeliveryMethod) *NullableProfileDeliveryMethod {
	return &NullableProfileDeliveryMethod{value: val, isSet: true}
}

func (v NullableProfileDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


