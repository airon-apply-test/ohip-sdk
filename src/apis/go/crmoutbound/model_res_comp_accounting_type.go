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

// checks if the ResCompAccountingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResCompAccountingType{}

// ResCompAccountingType Information regarding comp accounting on the reservation.
type ResCompAccountingType struct {
	// Code used to identify the casino comp type and ranking of a guest.
	CompType *string `json:"compType,omitempty"`
	// ID of the employee who will act as the host for this guest.
	Authorizer *string `json:"authorizer,omitempty"`
}

// NewResCompAccountingType instantiates a new ResCompAccountingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResCompAccountingType() *ResCompAccountingType {
	this := ResCompAccountingType{}
	return &this
}

// NewResCompAccountingTypeWithDefaults instantiates a new ResCompAccountingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResCompAccountingTypeWithDefaults() *ResCompAccountingType {
	this := ResCompAccountingType{}
	return &this
}

// GetCompType returns the CompType field value if set, zero value otherwise.
func (o *ResCompAccountingType) GetCompType() string {
	if o == nil || IsNil(o.CompType) {
		var ret string
		return ret
	}
	return *o.CompType
}

// GetCompTypeOk returns a tuple with the CompType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCompAccountingType) GetCompTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CompType) {
		return nil, false
	}
	return o.CompType, true
}

// HasCompType returns a boolean if a field has been set.
func (o *ResCompAccountingType) HasCompType() bool {
	if o != nil && !IsNil(o.CompType) {
		return true
	}

	return false
}

// SetCompType gets a reference to the given string and assigns it to the CompType field.
func (o *ResCompAccountingType) SetCompType(v string) {
	o.CompType = &v
}

// GetAuthorizer returns the Authorizer field value if set, zero value otherwise.
func (o *ResCompAccountingType) GetAuthorizer() string {
	if o == nil || IsNil(o.Authorizer) {
		var ret string
		return ret
	}
	return *o.Authorizer
}

// GetAuthorizerOk returns a tuple with the Authorizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCompAccountingType) GetAuthorizerOk() (*string, bool) {
	if o == nil || IsNil(o.Authorizer) {
		return nil, false
	}
	return o.Authorizer, true
}

// HasAuthorizer returns a boolean if a field has been set.
func (o *ResCompAccountingType) HasAuthorizer() bool {
	if o != nil && !IsNil(o.Authorizer) {
		return true
	}

	return false
}

// SetAuthorizer gets a reference to the given string and assigns it to the Authorizer field.
func (o *ResCompAccountingType) SetAuthorizer(v string) {
	o.Authorizer = &v
}

func (o ResCompAccountingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResCompAccountingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompType) {
		toSerialize["compType"] = o.CompType
	}
	if !IsNil(o.Authorizer) {
		toSerialize["authorizer"] = o.Authorizer
	}
	return toSerialize, nil
}

type NullableResCompAccountingType struct {
	value *ResCompAccountingType
	isSet bool
}

func (v NullableResCompAccountingType) Get() *ResCompAccountingType {
	return v.value
}

func (v *NullableResCompAccountingType) Set(val *ResCompAccountingType) {
	v.value = val
	v.isSet = true
}

func (v NullableResCompAccountingType) IsSet() bool {
	return v.isSet
}

func (v *NullableResCompAccountingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResCompAccountingType(val *ResCompAccountingType) *NullableResCompAccountingType {
	return &NullableResCompAccountingType{value: val, isSet: true}
}

func (v NullableResCompAccountingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResCompAccountingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


