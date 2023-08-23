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

// checks if the ItemizerShortInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemizerShortInfoType{}

// ItemizerShortInfoType struct for ItemizerShortInfoType
type ItemizerShortInfoType struct {
	// Code of the itemizer on a transaction code setup.
	Code *string `json:"code,omitempty"`
	// Itemizer name on a transaction code setup.
	ItemizerName *string `json:"itemizerName,omitempty"`
	// Itemizer number which is the split post on a transaction code setup.
	ItemNumber *int32 `json:"itemNumber,omitempty"`
}

// NewItemizerShortInfoType instantiates a new ItemizerShortInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemizerShortInfoType() *ItemizerShortInfoType {
	this := ItemizerShortInfoType{}
	return &this
}

// NewItemizerShortInfoTypeWithDefaults instantiates a new ItemizerShortInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemizerShortInfoTypeWithDefaults() *ItemizerShortInfoType {
	this := ItemizerShortInfoType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ItemizerShortInfoType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemizerShortInfoType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ItemizerShortInfoType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ItemizerShortInfoType) SetCode(v string) {
	o.Code = &v
}

// GetItemizerName returns the ItemizerName field value if set, zero value otherwise.
func (o *ItemizerShortInfoType) GetItemizerName() string {
	if o == nil || IsNil(o.ItemizerName) {
		var ret string
		return ret
	}
	return *o.ItemizerName
}

// GetItemizerNameOk returns a tuple with the ItemizerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemizerShortInfoType) GetItemizerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ItemizerName) {
		return nil, false
	}
	return o.ItemizerName, true
}

// HasItemizerName returns a boolean if a field has been set.
func (o *ItemizerShortInfoType) HasItemizerName() bool {
	if o != nil && !IsNil(o.ItemizerName) {
		return true
	}

	return false
}

// SetItemizerName gets a reference to the given string and assigns it to the ItemizerName field.
func (o *ItemizerShortInfoType) SetItemizerName(v string) {
	o.ItemizerName = &v
}

// GetItemNumber returns the ItemNumber field value if set, zero value otherwise.
func (o *ItemizerShortInfoType) GetItemNumber() int32 {
	if o == nil || IsNil(o.ItemNumber) {
		var ret int32
		return ret
	}
	return *o.ItemNumber
}

// GetItemNumberOk returns a tuple with the ItemNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemizerShortInfoType) GetItemNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemNumber) {
		return nil, false
	}
	return o.ItemNumber, true
}

// HasItemNumber returns a boolean if a field has been set.
func (o *ItemizerShortInfoType) HasItemNumber() bool {
	if o != nil && !IsNil(o.ItemNumber) {
		return true
	}

	return false
}

// SetItemNumber gets a reference to the given int32 and assigns it to the ItemNumber field.
func (o *ItemizerShortInfoType) SetItemNumber(v int32) {
	o.ItemNumber = &v
}

func (o ItemizerShortInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemizerShortInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ItemizerName) {
		toSerialize["itemizerName"] = o.ItemizerName
	}
	if !IsNil(o.ItemNumber) {
		toSerialize["itemNumber"] = o.ItemNumber
	}
	return toSerialize, nil
}

type NullableItemizerShortInfoType struct {
	value *ItemizerShortInfoType
	isSet bool
}

func (v NullableItemizerShortInfoType) Get() *ItemizerShortInfoType {
	return v.value
}

func (v *NullableItemizerShortInfoType) Set(val *ItemizerShortInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableItemizerShortInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableItemizerShortInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemizerShortInfoType(val *ItemizerShortInfoType) *NullableItemizerShortInfoType {
	return &NullableItemizerShortInfoType{value: val, isSet: true}
}

func (v NullableItemizerShortInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemizerShortInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


