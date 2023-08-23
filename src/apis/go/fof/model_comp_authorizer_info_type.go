/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CompAuthorizerInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompAuthorizerInfoType{}

// CompAuthorizerInfoType Authorizer Information
type CompAuthorizerInfoType struct {
	AuthorizerId *UniqueIDType `json:"authorizerId,omitempty"`
	// Application user name of the authorizer
	AuthorizerUserName *string `json:"authorizerUserName,omitempty"`
	// Full name of the authorizer.
	AuthorizerName *string `json:"authorizerName,omitempty"`
}

// NewCompAuthorizerInfoType instantiates a new CompAuthorizerInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompAuthorizerInfoType() *CompAuthorizerInfoType {
	this := CompAuthorizerInfoType{}
	return &this
}

// NewCompAuthorizerInfoTypeWithDefaults instantiates a new CompAuthorizerInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompAuthorizerInfoTypeWithDefaults() *CompAuthorizerInfoType {
	this := CompAuthorizerInfoType{}
	return &this
}

// GetAuthorizerId returns the AuthorizerId field value if set, zero value otherwise.
func (o *CompAuthorizerInfoType) GetAuthorizerId() UniqueIDType {
	if o == nil || IsNil(o.AuthorizerId) {
		var ret UniqueIDType
		return ret
	}
	return *o.AuthorizerId
}

// GetAuthorizerIdOk returns a tuple with the AuthorizerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompAuthorizerInfoType) GetAuthorizerIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.AuthorizerId) {
		return nil, false
	}
	return o.AuthorizerId, true
}

// HasAuthorizerId returns a boolean if a field has been set.
func (o *CompAuthorizerInfoType) HasAuthorizerId() bool {
	if o != nil && !IsNil(o.AuthorizerId) {
		return true
	}

	return false
}

// SetAuthorizerId gets a reference to the given UniqueIDType and assigns it to the AuthorizerId field.
func (o *CompAuthorizerInfoType) SetAuthorizerId(v UniqueIDType) {
	o.AuthorizerId = &v
}

// GetAuthorizerUserName returns the AuthorizerUserName field value if set, zero value otherwise.
func (o *CompAuthorizerInfoType) GetAuthorizerUserName() string {
	if o == nil || IsNil(o.AuthorizerUserName) {
		var ret string
		return ret
	}
	return *o.AuthorizerUserName
}

// GetAuthorizerUserNameOk returns a tuple with the AuthorizerUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompAuthorizerInfoType) GetAuthorizerUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizerUserName) {
		return nil, false
	}
	return o.AuthorizerUserName, true
}

// HasAuthorizerUserName returns a boolean if a field has been set.
func (o *CompAuthorizerInfoType) HasAuthorizerUserName() bool {
	if o != nil && !IsNil(o.AuthorizerUserName) {
		return true
	}

	return false
}

// SetAuthorizerUserName gets a reference to the given string and assigns it to the AuthorizerUserName field.
func (o *CompAuthorizerInfoType) SetAuthorizerUserName(v string) {
	o.AuthorizerUserName = &v
}

// GetAuthorizerName returns the AuthorizerName field value if set, zero value otherwise.
func (o *CompAuthorizerInfoType) GetAuthorizerName() string {
	if o == nil || IsNil(o.AuthorizerName) {
		var ret string
		return ret
	}
	return *o.AuthorizerName
}

// GetAuthorizerNameOk returns a tuple with the AuthorizerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompAuthorizerInfoType) GetAuthorizerNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizerName) {
		return nil, false
	}
	return o.AuthorizerName, true
}

// HasAuthorizerName returns a boolean if a field has been set.
func (o *CompAuthorizerInfoType) HasAuthorizerName() bool {
	if o != nil && !IsNil(o.AuthorizerName) {
		return true
	}

	return false
}

// SetAuthorizerName gets a reference to the given string and assigns it to the AuthorizerName field.
func (o *CompAuthorizerInfoType) SetAuthorizerName(v string) {
	o.AuthorizerName = &v
}

func (o CompAuthorizerInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompAuthorizerInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizerId) {
		toSerialize["authorizerId"] = o.AuthorizerId
	}
	if !IsNil(o.AuthorizerUserName) {
		toSerialize["authorizerUserName"] = o.AuthorizerUserName
	}
	if !IsNil(o.AuthorizerName) {
		toSerialize["authorizerName"] = o.AuthorizerName
	}
	return toSerialize, nil
}

type NullableCompAuthorizerInfoType struct {
	value *CompAuthorizerInfoType
	isSet bool
}

func (v NullableCompAuthorizerInfoType) Get() *CompAuthorizerInfoType {
	return v.value
}

func (v *NullableCompAuthorizerInfoType) Set(val *CompAuthorizerInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableCompAuthorizerInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableCompAuthorizerInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompAuthorizerInfoType(val *CompAuthorizerInfoType) *NullableCompAuthorizerInfoType {
	return &NullableCompAuthorizerInfoType{value: val, isSet: true}
}

func (v NullableCompAuthorizerInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompAuthorizerInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


