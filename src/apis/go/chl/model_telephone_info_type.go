/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
)

// checks if the TelephoneInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelephoneInfoType{}

// TelephoneInfoType Information on a telephone number for the customer.
type TelephoneInfoType struct {
	Telephone *TelephoneType `json:"telephone,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
}

// NewTelephoneInfoType instantiates a new TelephoneInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephoneInfoType() *TelephoneInfoType {
	this := TelephoneInfoType{}
	return &this
}

// NewTelephoneInfoTypeWithDefaults instantiates a new TelephoneInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephoneInfoTypeWithDefaults() *TelephoneInfoType {
	this := TelephoneInfoType{}
	return &this
}

// GetTelephone returns the Telephone field value if set, zero value otherwise.
func (o *TelephoneInfoType) GetTelephone() TelephoneType {
	if o == nil || IsNil(o.Telephone) {
		var ret TelephoneType
		return ret
	}
	return *o.Telephone
}

// GetTelephoneOk returns a tuple with the Telephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneInfoType) GetTelephoneOk() (*TelephoneType, bool) {
	if o == nil || IsNil(o.Telephone) {
		return nil, false
	}
	return o.Telephone, true
}

// HasTelephone returns a boolean if a field has been set.
func (o *TelephoneInfoType) HasTelephone() bool {
	if o != nil && !IsNil(o.Telephone) {
		return true
	}

	return false
}

// SetTelephone gets a reference to the given TelephoneType and assigns it to the Telephone field.
func (o *TelephoneInfoType) SetTelephone(v TelephoneType) {
	o.Telephone = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TelephoneInfoType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneInfoType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TelephoneInfoType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TelephoneInfoType) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TelephoneInfoType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneInfoType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TelephoneInfoType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TelephoneInfoType) SetType(v string) {
	o.Type = &v
}

func (o TelephoneInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelephoneInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Telephone) {
		toSerialize["telephone"] = o.Telephone
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTelephoneInfoType struct {
	value *TelephoneInfoType
	isSet bool
}

func (v NullableTelephoneInfoType) Get() *TelephoneInfoType {
	return v.value
}

func (v *NullableTelephoneInfoType) Set(val *TelephoneInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephoneInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephoneInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephoneInfoType(val *TelephoneInfoType) *NullableTelephoneInfoType {
	return &NullableTelephoneInfoType{value: val, isSet: true}
}

func (v NullableTelephoneInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephoneInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


