/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intcfg

import (
	"encoding/json"
)

// checks if the ItemizerType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemizerType{}

// ItemizerType struct for ItemizerType
type ItemizerType struct {
	// Split code of the itemizer setup.
	SplitCode *string `json:"splitCode,omitempty"`
	// Split posting of the itemizer setup.
	SplitPost *int32 `json:"splitPost,omitempty"`
	// Split text of the itemizer setup.
	SplitText *string `json:"splitText,omitempty"`
	// Split factor of the itemizer setup.
	SplitFactor *int32 `json:"splitFactor,omitempty"`
}

// NewItemizerType instantiates a new ItemizerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemizerType() *ItemizerType {
	this := ItemizerType{}
	return &this
}

// NewItemizerTypeWithDefaults instantiates a new ItemizerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemizerTypeWithDefaults() *ItemizerType {
	this := ItemizerType{}
	return &this
}

// GetSplitCode returns the SplitCode field value if set, zero value otherwise.
func (o *ItemizerType) GetSplitCode() string {
	if o == nil || IsNil(o.SplitCode) {
		var ret string
		return ret
	}
	return *o.SplitCode
}

// GetSplitCodeOk returns a tuple with the SplitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemizerType) GetSplitCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SplitCode) {
		return nil, false
	}
	return o.SplitCode, true
}

// HasSplitCode returns a boolean if a field has been set.
func (o *ItemizerType) HasSplitCode() bool {
	if o != nil && !IsNil(o.SplitCode) {
		return true
	}

	return false
}

// SetSplitCode gets a reference to the given string and assigns it to the SplitCode field.
func (o *ItemizerType) SetSplitCode(v string) {
	o.SplitCode = &v
}

// GetSplitPost returns the SplitPost field value if set, zero value otherwise.
func (o *ItemizerType) GetSplitPost() int32 {
	if o == nil || IsNil(o.SplitPost) {
		var ret int32
		return ret
	}
	return *o.SplitPost
}

// GetSplitPostOk returns a tuple with the SplitPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemizerType) GetSplitPostOk() (*int32, bool) {
	if o == nil || IsNil(o.SplitPost) {
		return nil, false
	}
	return o.SplitPost, true
}

// HasSplitPost returns a boolean if a field has been set.
func (o *ItemizerType) HasSplitPost() bool {
	if o != nil && !IsNil(o.SplitPost) {
		return true
	}

	return false
}

// SetSplitPost gets a reference to the given int32 and assigns it to the SplitPost field.
func (o *ItemizerType) SetSplitPost(v int32) {
	o.SplitPost = &v
}

// GetSplitText returns the SplitText field value if set, zero value otherwise.
func (o *ItemizerType) GetSplitText() string {
	if o == nil || IsNil(o.SplitText) {
		var ret string
		return ret
	}
	return *o.SplitText
}

// GetSplitTextOk returns a tuple with the SplitText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemizerType) GetSplitTextOk() (*string, bool) {
	if o == nil || IsNil(o.SplitText) {
		return nil, false
	}
	return o.SplitText, true
}

// HasSplitText returns a boolean if a field has been set.
func (o *ItemizerType) HasSplitText() bool {
	if o != nil && !IsNil(o.SplitText) {
		return true
	}

	return false
}

// SetSplitText gets a reference to the given string and assigns it to the SplitText field.
func (o *ItemizerType) SetSplitText(v string) {
	o.SplitText = &v
}

// GetSplitFactor returns the SplitFactor field value if set, zero value otherwise.
func (o *ItemizerType) GetSplitFactor() int32 {
	if o == nil || IsNil(o.SplitFactor) {
		var ret int32
		return ret
	}
	return *o.SplitFactor
}

// GetSplitFactorOk returns a tuple with the SplitFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemizerType) GetSplitFactorOk() (*int32, bool) {
	if o == nil || IsNil(o.SplitFactor) {
		return nil, false
	}
	return o.SplitFactor, true
}

// HasSplitFactor returns a boolean if a field has been set.
func (o *ItemizerType) HasSplitFactor() bool {
	if o != nil && !IsNil(o.SplitFactor) {
		return true
	}

	return false
}

// SetSplitFactor gets a reference to the given int32 and assigns it to the SplitFactor field.
func (o *ItemizerType) SetSplitFactor(v int32) {
	o.SplitFactor = &v
}

func (o ItemizerType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemizerType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SplitCode) {
		toSerialize["splitCode"] = o.SplitCode
	}
	if !IsNil(o.SplitPost) {
		toSerialize["splitPost"] = o.SplitPost
	}
	if !IsNil(o.SplitText) {
		toSerialize["splitText"] = o.SplitText
	}
	if !IsNil(o.SplitFactor) {
		toSerialize["splitFactor"] = o.SplitFactor
	}
	return toSerialize, nil
}

type NullableItemizerType struct {
	value *ItemizerType
	isSet bool
}

func (v NullableItemizerType) Get() *ItemizerType {
	return v.value
}

func (v *NullableItemizerType) Set(val *ItemizerType) {
	v.value = val
	v.isSet = true
}

func (v NullableItemizerType) IsSet() bool {
	return v.isSet
}

func (v *NullableItemizerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemizerType(val *ItemizerType) *NullableItemizerType {
	return &NullableItemizerType{value: val, isSet: true}
}

func (v NullableItemizerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemizerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


