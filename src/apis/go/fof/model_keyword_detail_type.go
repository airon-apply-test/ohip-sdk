/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the KeywordDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordDetailType{}

// KeywordDetailType struct for KeywordDetailType
type KeywordDetailType struct {
	// The new keyword value.
	NewKeyword *string `json:"newKeyword,omitempty"`
}

// NewKeywordDetailType instantiates a new KeywordDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordDetailType() *KeywordDetailType {
	this := KeywordDetailType{}
	return &this
}

// NewKeywordDetailTypeWithDefaults instantiates a new KeywordDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordDetailTypeWithDefaults() *KeywordDetailType {
	this := KeywordDetailType{}
	return &this
}

// GetNewKeyword returns the NewKeyword field value if set, zero value otherwise.
func (o *KeywordDetailType) GetNewKeyword() string {
	if o == nil || IsNil(o.NewKeyword) {
		var ret string
		return ret
	}
	return *o.NewKeyword
}

// GetNewKeywordOk returns a tuple with the NewKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordDetailType) GetNewKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.NewKeyword) {
		return nil, false
	}
	return o.NewKeyword, true
}

// HasNewKeyword returns a boolean if a field has been set.
func (o *KeywordDetailType) HasNewKeyword() bool {
	if o != nil && !IsNil(o.NewKeyword) {
		return true
	}

	return false
}

// SetNewKeyword gets a reference to the given string and assigns it to the NewKeyword field.
func (o *KeywordDetailType) SetNewKeyword(v string) {
	o.NewKeyword = &v
}

func (o KeywordDetailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeywordDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewKeyword) {
		toSerialize["newKeyword"] = o.NewKeyword
	}
	return toSerialize, nil
}

type NullableKeywordDetailType struct {
	value *KeywordDetailType
	isSet bool
}

func (v NullableKeywordDetailType) Get() *KeywordDetailType {
	return v.value
}

func (v *NullableKeywordDetailType) Set(val *KeywordDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordDetailType(val *KeywordDetailType) *NullableKeywordDetailType {
	return &NullableKeywordDetailType{value: val, isSet: true}
}

func (v NullableKeywordDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeywordDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


