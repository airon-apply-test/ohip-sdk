/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the IdentificationInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentificationInfoType{}

// IdentificationInfoType Information on the identification of the customer.
type IdentificationInfoType struct {
	Identification *IdentificationType `json:"identification,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
}

// NewIdentificationInfoType instantiates a new IdentificationInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentificationInfoType() *IdentificationInfoType {
	this := IdentificationInfoType{}
	return &this
}

// NewIdentificationInfoTypeWithDefaults instantiates a new IdentificationInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentificationInfoTypeWithDefaults() *IdentificationInfoType {
	this := IdentificationInfoType{}
	return &this
}

// GetIdentification returns the Identification field value if set, zero value otherwise.
func (o *IdentificationInfoType) GetIdentification() IdentificationType {
	if o == nil || IsNil(o.Identification) {
		var ret IdentificationType
		return ret
	}
	return *o.Identification
}

// GetIdentificationOk returns a tuple with the Identification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationInfoType) GetIdentificationOk() (*IdentificationType, bool) {
	if o == nil || IsNil(o.Identification) {
		return nil, false
	}
	return o.Identification, true
}

// HasIdentification returns a boolean if a field has been set.
func (o *IdentificationInfoType) HasIdentification() bool {
	if o != nil && !IsNil(o.Identification) {
		return true
	}

	return false
}

// SetIdentification gets a reference to the given IdentificationType and assigns it to the Identification field.
func (o *IdentificationInfoType) SetIdentification(v IdentificationType) {
	o.Identification = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentificationInfoType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationInfoType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentificationInfoType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentificationInfoType) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentificationInfoType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationInfoType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentificationInfoType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentificationInfoType) SetType(v string) {
	o.Type = &v
}

func (o IdentificationInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentificationInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identification) {
		toSerialize["identification"] = o.Identification
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIdentificationInfoType struct {
	value *IdentificationInfoType
	isSet bool
}

func (v NullableIdentificationInfoType) Get() *IdentificationInfoType {
	return v.value
}

func (v *NullableIdentificationInfoType) Set(val *IdentificationInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationInfoType(val *IdentificationInfoType) *NullableIdentificationInfoType {
	return &NullableIdentificationInfoType{value: val, isSet: true}
}

func (v NullableIdentificationInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


