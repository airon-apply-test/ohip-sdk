/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the WelcomeOfferType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WelcomeOfferType{}

// WelcomeOfferType struct for WelcomeOfferType
type WelcomeOfferType struct {
	// Determines the status of the welcome offer.
	Status *string `json:"status,omitempty"`
	Type *WelcomeOfferOptionsType `json:"type,omitempty"`
}

// NewWelcomeOfferType instantiates a new WelcomeOfferType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWelcomeOfferType() *WelcomeOfferType {
	this := WelcomeOfferType{}
	return &this
}

// NewWelcomeOfferTypeWithDefaults instantiates a new WelcomeOfferType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWelcomeOfferTypeWithDefaults() *WelcomeOfferType {
	this := WelcomeOfferType{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WelcomeOfferType) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WelcomeOfferType) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WelcomeOfferType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WelcomeOfferType) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WelcomeOfferType) GetType() WelcomeOfferOptionsType {
	if o == nil || IsNil(o.Type) {
		var ret WelcomeOfferOptionsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WelcomeOfferType) GetTypeOk() (*WelcomeOfferOptionsType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WelcomeOfferType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given WelcomeOfferOptionsType and assigns it to the Type field.
func (o *WelcomeOfferType) SetType(v WelcomeOfferOptionsType) {
	o.Type = &v
}

func (o WelcomeOfferType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WelcomeOfferType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableWelcomeOfferType struct {
	value *WelcomeOfferType
	isSet bool
}

func (v NullableWelcomeOfferType) Get() *WelcomeOfferType {
	return v.value
}

func (v *NullableWelcomeOfferType) Set(val *WelcomeOfferType) {
	v.value = val
	v.isSet = true
}

func (v NullableWelcomeOfferType) IsSet() bool {
	return v.isSet
}

func (v *NullableWelcomeOfferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWelcomeOfferType(val *WelcomeOfferType) *NullableWelcomeOfferType {
	return &NullableWelcomeOfferType{value: val, isSet: true}
}

func (v NullableWelcomeOfferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWelcomeOfferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


